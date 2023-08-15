package v3_2_1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

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

		iyaParams, err := keepers.YieldaggregatorKeeper.GetParams(ctx)
		if err != nil {
			return vm, err
		}
		iyaParams.FeeCollectorAddress = keepers.AccountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName).GetAddress().String()
		_ = keepers.YieldaggregatorKeeper.SetParams(ctx, iyaParams)

		return vm, nil
	}
}
