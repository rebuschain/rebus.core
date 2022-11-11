package claim

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rebuschain/rebus.core/v1/x/claim/keeper"
	"github.com/rebuschain/rebus.core/v1/x/claim/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// ensure claim module account is set on genesis
	if acc := k.GetModuleAccount(ctx); acc == nil {
		panic("the claim module account has not been set")
	}

	// set the start time to the current block time by default
	if genState.Params.AirdropStartTime.IsZero() {
		genState.Params.AirdropStartTime = ctx.BlockTime()
	}
	/*
		if genState.Params.AirdropStartTime.Equal(time.Time{}) {
			genState.Params.AirdropStartTime = ctx.BlockTime()
		}
	*/

	k.SetParams(ctx, genState.Params)
	_ = k.SetClaimRecords(ctx, genState.ClaimRecords)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params, _ := k.GetParams(ctx)
	genesis := types.DefaultGenesis()
	genesis.Params = params
	genesis.ClaimRecords = k.GetClaimRecords(ctx)
	return genesis
}
