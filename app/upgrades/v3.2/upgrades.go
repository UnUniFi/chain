package v3_2

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/UnUniFi/chain/app/keepers"
	"github.com/UnUniFi/chain/app/upgrades"
)

func CreateUpgradeHandler(mm *module.Manager,
	configurator module.Configurator,
	_ upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info(fmt.Sprintf("update start:%s", UpgradeName))

		vm, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vm, err
		}

		factoryParam, err := keepers.NftfactoryKeeper.GetParams(ctx)
		if err != nil {
			return vm, err
		}
		factoryParam.ClassCreationFee = []sdk.Coin{}
		factoryParam.FeeCollectorAddress = ""
		_ = keepers.NftfactoryKeeper.SetParams(ctx, factoryParam)

		iyaParam, err := keepers.YieldaggregatorKeeper.GetParams(ctx)
		if err != nil {
			return vm, err
		}
		_ = keepers.YieldaggregatorKeeper.SetParams(ctx, iyaParam)

		return vm, nil
	}
}
