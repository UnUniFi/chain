package v3_1_rc0

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/UnUniFi/chain/app/keepers"
	"github.com/UnUniFi/chain/app/upgrades"
	epochtypes "github.com/UnUniFi/chain/x/epochs/types"
)

func CreateUpgradeHandler(mm *module.Manager,
	configurator module.Configurator,
	_ upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info(fmt.Sprintf("update start:%s", UpgradeName))

		epochInfo, found := keepers.EpochsKeeper.GetEpochInfo(ctx, epochtypes.BASE_EPOCH)
		if !found {
			return vm, fmt.Errorf("epoch %s not found", epochtypes.BASE_EPOCH)
		}
		epochInfo.Duration = time.Minute * 15 // 15min
		keepers.EpochsKeeper.SetEpochInfo(ctx, epochInfo)

		return vm, nil
	}
}
