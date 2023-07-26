package v3

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

		iyaParam := keepers.YieldaggregatorKeeper.GetParams(ctx)
		iyaParam.CommissionRate = sdk.NewDecWithPrec(1, 3)
		iyaParam.VaultCreationFee = sdk.NewCoin("uguu", sdk.NewInt(10000000))
		iyaParam.VaultCreationDeposit = sdk.NewCoin("uguu", sdk.NewInt(1000000))

		keepers.YieldaggregatorKeeper.SetParams(ctx, iyaParam)
		return vm, nil
	}
}
