package v0

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	feemarkettypes "github.com/tharsis/ethermint/x/feemarket/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v4
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger()
		// Refs:
		// - https://docs.cosmos.network/master/building-modules/upgrade.html#registering-migrations
		// - https://docs.cosmos.network/master/migrations/chain-upgrade-guide-044.html#chain-upgrade

		logger.Info("upgrading: ", ctx.BlockHeight())

		vm[feemarkettypes.ModuleName] = 2

		// Leave modules are as-is to avoid running InitGenesis.
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
