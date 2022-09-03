package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/rebuschain/rebus.core/v1/x/claim/types"
)

// Keeper struct
type Keeper struct {
	cdc           codec.Codec
	storeKey      sdk.StoreKey
	paramSpace    paramtypes.Subspace
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
	distrKeeper   types.DistrKeeper
}

// NewKeeper returns keeper
func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace,
	ak types.AccountKeeper, bk types.BankKeeper, sk types.StakingKeeper, dk types.DistrKeeper) *Keeper {

	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		accountKeeper: ak,
		bankKeeper:    bk,
		stakingKeeper: sk,
		distrKeeper:   dk,
	}
}

// Logger returns logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetModuleAccount returns the module account for the claim module
func (k Keeper) GetModuleAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
}

// GetModuleAccountAddress gets the airdrop coin balance of module account
func (k Keeper) GetModuleAccountAddress() sdk.AccAddress {
	return k.accountKeeper.GetModuleAddress(types.ModuleName)
}

// GetModuleAccountBalances gets the balances of module account that escrows the
// airdrop tokens
func (k Keeper) GetModuleAccountBalances(ctx sdk.Context) sdk.Coins {
	moduleAccAddr := k.GetModuleAccountAddress()
	return k.bankKeeper.GetAllBalances(ctx, moduleAccAddr)
}

func (k Keeper) GetModuleAccountBalance(ctx sdk.Context) sdk.Coin {
	moduleAccAddr := k.GetModuleAccountAddress()
	params, _ := k.GetParams(ctx)
	return k.bankKeeper.GetBalance(ctx, moduleAccAddr, params.ClaimDenom)
}
