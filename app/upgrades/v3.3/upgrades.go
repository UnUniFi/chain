package v3_1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

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

		params := yieldaggregatortypes.Params{}
		keepers.GetSubspace(yieldaggregatortypes.ModuleName).GetParamSet(ctx, &params)

		_ = keepers.YieldaggregatorKeeper.SetParams(ctx, &params)

		return vm, nil
	}
}
