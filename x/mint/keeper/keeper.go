package keeper

import (
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tharsis/evmos/v4/x/mint/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         sdk.StoreKey
	paramSpace       paramtypes.Subspace
	stakingKeeper    types.StakingKeeper
	bankKeeper       types.BankKeeper
	accountKeeper    types.AccountKeeper
	distrKeeper      types.DistrKeeper
	feeCollectorName string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	sk types.StakingKeeper, ak types.AccountKeeper, bk types.BankKeeper,
	dk types.DistrKeeper,
	feeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:              cdc,
		storeKey:         key,
		paramSpace:       paramSpace,
		stakingKeeper:    sk,
		bankKeeper:       bk,
		accountKeeper:    ak,
		distrKeeper:      dk,
		feeCollectorName: feeCollectorName,
	}
}

//______________________________________________________________________

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// get the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

// set the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.MinterKey, b)
}

//______________________________________________________________________

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

//______________________________________________________________________

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// StakingTokenSupply to be used in BeginBlocker.
func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdk.Int {
	return k.stakingKeeper.StakingTokenSupply(ctx)
}

// TokenSupply implements an alias call to the underlying bank keeper's
// TokenSupply to be used in BeginBlocker.
func (k Keeper) TokenSupply(ctx sdk.Context, denom string) sdk.Int {
	return k.bankKeeper.GetSupply(ctx, denom).Amount
}

// BondedRatio implements an alias call to the underlying staking keeper's
// BondedRatio to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	return k.stakingKeeper.BondedRatio(ctx)
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

//
func (k Keeper) getAddress(ctx sdk.Context, listAddress []string) (sdk.AccAddress, error) {
	strAddress := listAddress[0]
	address, err := sdk.AccAddressFromBech32(strAddress)
	if err != nil {
		return nil, err
	}
	return address, nil
}

// GetProportions gets the balance of the `MintedDenom` from minted coins and returns coins according to the `AllocationRatio`.
func (k Keeper) GetProportions(ctx sdk.Context, mintedCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(mintedCoin.Denom, mintedCoin.Amount.ToDec().Mul(ratio).TruncateInt())
}

func (k Keeper) DistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin) error {
	posRewardCoins := sdk.NewCoins(k.GetProportions(ctx, mintedCoin, types.PosRewardProportion))
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, posRewardCoins)
	if err != nil {
		return err
	}

	communityRewardCoins := sdk.NewCoins(k.GetProportions(ctx, mintedCoin, types.CommunityProportion))
	k.distrKeeper.FundCommunityPool(ctx, communityRewardCoins, k.accountKeeper.GetModuleAddress(types.ModuleName))
	if err != nil {
		return err
	}

	incentiveAddress, err := k.getAddress(ctx, types.IncentiveAddressList[:])
	if err != nil {
		return err
	}
	incentiveCoins := sdk.NewCoins(k.GetProportions(ctx, mintedCoin, types.IncentiveProportion))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, incentiveAddress, incentiveCoins)
	if err != nil {
		return err
	}

	ethicalAddress, err := k.getAddress(ctx, types.EthicalAddressList[:])
	if err != nil {
		return err
	}
	ethicalCoins := sdk.NewCoins(k.GetProportions(ctx, mintedCoin, types.EthicalProportion))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ethicalAddress, ethicalCoins)
	if err != nil {
		return err
	}

	treasuryAddress, err := k.getAddress(ctx, types.TreasuryAddressList[:])
	if err != nil {
		return err
	}
	treasuryCoins := sdk.NewCoins(k.GetProportions(ctx, mintedCoin, types.TreasuryProportion))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, treasuryAddress, treasuryCoins)
	if err != nil {
		return err
	}

	// change := sdk.NewCoins(mintedCoin).Sub(posRewardCoins).Sub(communityRewardCoins).Sub(incentiveCoins).Sub(ethicalCoins).Sub(treasuryCoins)
	//

	return err
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) AddCollectedFees(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, fees)
}
