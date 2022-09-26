package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v1
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	sk stakingkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		FixMinCommissionRate(ctx, sk)
		return mm.RunMigrations(ctx, configurator, vm)
	}
}

func FixMinCommissionRate(ctx sdk.Context, staking stakingkeeper.Keeper) {
	// Upgrade every validators min-commission rate
	validators := staking.GetAllValidators(ctx)
	minCommissionRate := sdk.NewDecWithPrec(5, 2) // 5%
	blockTime := ctx.BlockHeader().Time
	for _, v := range validators {
		if v.Commission.Rate.LT(minCommissionRate) {
			v.Commission.Rate = minCommissionRate
			v.Commission.UpdateTime = blockTime

			// call the before-modification hook since we're about to update the commission
			staking.BeforeValidatorModified(ctx, v.GetOperator())
			staking.SetValidator(ctx, v)
		}
	}
}
