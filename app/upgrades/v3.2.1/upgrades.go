package v3_2_1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/UnUniFi/chain/app/keepers"
	"github.com/UnUniFi/chain/app/upgrades"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func CreateUpgradeHandler(mm *module.Manager,
	configurator module.Configurator,
	_ upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info(fmt.Sprintf("update start:%s", UpgradeName))

		iyaParams := yieldaggregatortypes.Params{}
		paramtypes.NewKeyTable().RegisterParamSet(&yieldaggregatortypes.Params{})
		keepers.GetSubspace(yieldaggregatortypes.ModuleName).WithKeyTable(yieldaggregatortypes.ParamKeyTable()).GetParamSet(ctx, &iyaParams)

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

		iyaParams.FeeCollectorAddress = keepers.AccountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName).String()
		_ = keepers.YieldaggregatorKeeper.SetParams(ctx, &iyaParams)

		return vm, nil
	}
}
