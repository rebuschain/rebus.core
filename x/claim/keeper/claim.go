package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/rebuschain/rebus.core/v1/x/claim/types"
)

func (k Keeper) EndAirdrop(ctx sdk.Context) error {
	err := k.fundRemainingsToCommunity(ctx)
	if err != nil {
		return err
	}
	k.clearInitialClaimables(ctx)
	return nil
}

// ClearClaimables clear claimable amounts
func (k Keeper) clearInitialClaimables(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ClaimRecordsStorePrefix))
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		store.Delete(key)
	}
	iterator.Close()
}

// SetClaimables set claimable amount from balances object
func (k Keeper) SetClaimRecords(ctx sdk.Context, claimRecords []types.ClaimRecord) error {
	for _, claimRecord := range claimRecords {
		err := k.SetClaimRecord(ctx, claimRecord)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetClaimables get claimables for genesis export
func (k Keeper) GetClaimRecords(ctx sdk.Context) []types.ClaimRecord {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.ClaimRecordsStorePrefix))

	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	claimRecords := []types.ClaimRecord{}
	for ; iterator.Valid(); iterator.Next() {

		claimRecord := types.ClaimRecord{}

		err := proto.Unmarshal(iterator.Value(), &claimRecord)
		if err != nil {
			panic(err)
		}

		claimRecords = append(claimRecords, claimRecord)
	}
	return claimRecords
}

// GetClaimRecord returns the claim record for a specific address
func (k Keeper) GetClaimRecord(ctx sdk.Context, addr sdk.AccAddress) (types.ClaimRecord, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ClaimRecordsStorePrefix))

	bz := store.Get(addr)
	if len(bz) == 0 {
		return types.ClaimRecord{}, nil
	}

	claimRecord := types.ClaimRecord{}
	err := proto.Unmarshal(bz, &claimRecord)
	if err != nil {
		return types.ClaimRecord{}, err
	}

	return claimRecord, nil
}

// SetClaimRecord sets a claim record for an address in store
func (k Keeper) SetClaimRecord(ctx sdk.Context, claimRecord types.ClaimRecord) error {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.ClaimRecordsStorePrefix))

	bz, err := proto.Marshal(&claimRecord)
	if err != nil {
		return err
	}

	addr, err := sdk.AccAddressFromBech32(claimRecord.Address)
	if err != nil {
		return err
	}

	prefixStore.Set(addr, bz)
	return nil
}

// GetClaimable returns claimable amount for a specific action done by an address
func (k Keeper) GetClaimableAmountForAction(ctx sdk.Context, claimRecord types.ClaimRecord, addr sdk.AccAddress, action types.Action) (sdk.Coins, error) {
	if claimRecord.Address == "" {
		return sdk.Coins{}, nil
	}

	// if action already completed, nothing is claimable
	if claimRecord.ActionCompleted[action] {
		return sdk.Coins{}, nil
	}

	params, err := k.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	// If we are before the start time, do nothing.
	// This case _shouldn't_ occur on chain, since the
	// start time ought to be chain start time.
	if ctx.BlockTime().Before(params.AirdropStartTime) {
		return sdk.Coins{}, nil
	}

	InitialClaimablePerAction := sdk.Coins{}
	for _, coin := range claimRecord.InitialClaimableAmount {
		InitialClaimablePerAction = InitialClaimablePerAction.Add(
			sdk.NewCoin(coin.Denom,
				coin.Amount.QuoRaw(int64(len(types.Action_name)+1)),
			),
		)
	}

	elapsedAirdropTime := ctx.BlockTime().Sub(params.AirdropStartTime)
	if elapsedAirdropTime <= params.AirdropDuration {
		return InitialClaimablePerAction, nil
	}

	return sdk.Coins{}, nil
}

// GetClaimable returns the total claimable amount by an address
func (k Keeper) GetUserTotalClaimable(ctx sdk.Context, addr sdk.AccAddress) (sdk.Coins, error) {
	claimRecord, err := k.GetClaimRecord(ctx, addr)
	if err != nil {
		return sdk.Coins{}, err
	}
	if claimRecord.Address == "" {
		return sdk.Coins{}, nil
	}

	totalClaimable := sdk.Coins{}

	for action := range types.Action_name {
		claimableForAction, err := k.GetClaimableAmountForAction(ctx, claimRecord, addr, types.Action(action))
		if err != nil {
			return sdk.Coins{}, err
		}
		totalClaimable = totalClaimable.Add(claimableForAction...)
	}
	return totalClaimable, nil
}

// ClaimCoins remove claimable amount entry and transfer it to user's account
func (k Keeper) ClaimCoinsForAction(ctx sdk.Context, claimRecord types.ClaimRecord, addr sdk.AccAddress, action types.Action) (sdk.Coins, error) {
	claimableAmount, err := k.GetClaimableAmountForAction(ctx, claimRecord, addr, action)
	if err != nil {
		return claimableAmount, err
	}

	if claimableAmount.Empty() {
		return claimableAmount, nil
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, claimableAmount)
	if err != nil {
		return nil, err
	}

	claimRecord.ActionCompleted[action] = true

	err = k.SetClaimRecord(ctx, claimRecord)
	if err != nil {
		return claimableAmount, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeClaim,
			sdk.NewAttribute(sdk.AttributeKeySender, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, claimableAmount.String()),
		),
	})

	return claimableAmount, nil
}

// FundRemainingsToCommunity fund remainings to the community when airdrop period end
func (k Keeper) fundRemainingsToCommunity(ctx sdk.Context) error {
	moduleAccAddr := k.GetModuleAccountAddress()
	amt := k.GetModuleAccountBalance(ctx)
	return k.distrKeeper.FundCommunityPool(ctx, sdk.NewCoins(amt), moduleAccAddr)
}
