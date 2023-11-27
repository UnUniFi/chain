package v4_0_1

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

		keepers.BankKeeper.SetSendEnabled(ctx, Denom, true)

		result, err := BankSendList(ctx)
		if err != nil {
			panic(err)
		}
		err = upgradeBankSend(ctx, keepers.AccountKeeper, keepers.BankKeeper, result)
		if err != nil {
			panic(err)
		}

		keepers.BankKeeper.SetSendEnabled(ctx, Denom, false)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
