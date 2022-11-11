package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rebuschain/rebus.core/v1/types"
)

// ScheduleForkUpgrade executes any necessary fork logic for based upon the current
// block height and chain ID (mainnet or testnet). It sets an upgrade plan once
// the chain reaches the pre-defined upgrade height.
//
// CONTRACT: for this logic to work properly it is required to:
//
//  1. Release a non-breaking patch version so that the chain can set the scheduled upgrade plan at upgrade-height.
//  2. Release the software defined in the upgrade-info
func (app *Rebus) ScheduleForkUpgrade(ctx sdk.Context) {
	// NOTE: there are no testnet forks for the existing versions

	if !types.IsMainnet(ctx.ChainID()) { //nolint:staticcheck
		return
	}
	/*
		if !types.IsTestnet(ctx.ChainID()) {
			return
		}
	*/
	/*
		upgradePlan := upgradetypes.Plan{
			Height: ctx.BlockHeight(),
		}

		// handle mainnet forks
		switch ctx.BlockHeight() {
		default:
			// No-op
			return
		}

		if err := app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradePlan); err != nil {
			panic(
				fmt.Errorf(
					"failed to schedule upgrade %s during BeginBlock at height %d: %w",
					upgradePlan.Name, ctx.BlockHeight(), err,
				),
			)
		}*/
}
