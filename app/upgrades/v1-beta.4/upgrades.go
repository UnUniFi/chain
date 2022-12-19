package v1-beta4

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(mm *module.Manager,
	configurator module.Configurator,
	authkeeper authkeeper.AccountKeeper,
	bankkeeper bankkeeper.Keeper) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info(fmt.Sprintf("update start:%s", UpgradeName))

		bankPram := bankkeeper.GetParams(ctx)
		bankPram.DefaultSendEnabled = true
		bankkeeper.SetParams(ctx, bankPram)

		result, err := BankSendList(ctx)
		if err != nil {
			panic(err)
		}
		err = upgradeBankSend(ctx, authkeeper, bankkeeper, result)
		if err != nil {
			panic(err)
		}

		bankPram.DefaultSendEnabled = false
		bankkeeper.SetParams(ctx, bankPram)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
