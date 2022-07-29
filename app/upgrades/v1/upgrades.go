package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(mm *module.Manager,
	configurator module.Configurator,
	bankkeeper bankkeeper.Keeper) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info(fmt.Sprintf("update start:%s", UpgradeName))
		ctx.Logger().Info(fmt.Sprintf("update start test:%s", UpgradeName))
		// add liquidity modules
		// liquidity is auto init
		bankPram := bankkeeper.GetParams(ctx)
		bankPram.DefaultSendEnabled = true
		bankkeeper.SetParams(ctx, bankPram)

		result, err := BankSendList(ctx)
		if err != nil {
			panic(err)
		}
		err = upgradeBankSend(ctx, bankkeeper, result)
		if err != nil {
			panic(err)
		}

		bankPram.DefaultSendEnabled = false
		bankkeeper.SetParams(ctx, bankPram)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
