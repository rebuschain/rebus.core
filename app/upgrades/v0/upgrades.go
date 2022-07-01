package v0

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ethermint "github.com/tharsis/ethermint/types"
	feemarkettypes "github.com/tharsis/ethermint/x/feemarket/types"
	rebuscfg "github.com/tharsis/evmos/v4/cmd/config"
	mintkeeper "github.com/tharsis/evmos/v4/x/mint/keeper"
	types "github.com/tharsis/evmos/v4/x/mint/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v4
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bk types.BankKeeper,
	mk mintkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger()
		// Refs:
		// - https://docs.cosmos.network/master/building-modules/upgrade.html#registering-migrations
		// - https://docs.cosmos.network/master/migrations/chain-upgrade-guide-044.html#chain-upgrade

		logger.Info("upgrading: ", ctx.BlockHeight())

		ResetCoinDistribution(ctx, bk, mk)

		vm[feemarkettypes.ModuleName] = 2

		// Leave modules are as-is to avoid running InitGenesis.
		return mm.RunMigrations(ctx, configurator, vm)
	}
}

func getAddress(listAddress []string) (sdk.AccAddress, error) {
	strAddress := listAddress[0]
	address, err := sdk.AccAddressFromBech32(strAddress)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func ResetCoinDistribution(ctx sdk.Context, bk types.BankKeeper, mk mintkeeper.Keeper) {

	logger := ctx.Logger()

	var OriginAddressList = [1]string{OriginAddress}

	originAddress, err := getAddress(OriginAddressList[:])
	if err != nil {
		panic(fmt.Errorf("failed to upgrade getting originaddress: %w", err))
	}

	supply := bk.GetSupply(ctx, rebuscfg.BaseDenom)

	logger.Info("supply before burning event: ", supply)

	amt := sdk.NewInt(OriginAmt).Mul(ethermint.PowerReduction)
	coin := sdk.NewCoin(rebuscfg.BaseDenom, amt.ToDec().TruncateInt())
	originCoins := sdk.NewCoins(coin)
	err = bk.SendCoinsFromAccountToModule(ctx, originAddress, types.ModuleName, originCoins)
	if err != nil {
		panic(fmt.Errorf("failed to upgrade sending coins from account to mint module: %w", err))
	}
	err = bk.BurnCoins(ctx, types.ModuleName, originCoins)
	if err != nil {
		panic(fmt.Errorf("failed to upgrade burning coins from mint module: %w", err))
	}
	supply = bk.GetSupply(ctx, rebuscfg.BaseDenom)

	logger.Info("supply after burning event: ", supply)

	logger.Info("enabling minting ")

	minter := types.DefaultInitialMinter()
	mk.SetMinter(ctx, minter)
	logger.Info("minting enabled")
}
