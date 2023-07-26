package v1_beta3

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
		ctx.Logger().Info(fmt.Sprintf("update start test:%s", UpgradeName))

		bankPram := keepers.BankKeeper.GetParams(ctx)
		bankPram.DefaultSendEnabled = true
		keepers.BankKeeper.SetParams(ctx, bankPram)

		result, err := BankSendList(ctx)
		if err != nil {
			panic(err)
		}
		err = upgradeBankSend(ctx, *keepers.AccountKeeper, *keepers.BankKeeper, result)
		if err != nil {
			panic(err)
		}

		bankPram.DefaultSendEnabled = false
		keepers.BankKeeper.SetParams(ctx, bankPram)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
