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

		fromAddr, err := sdk.AccAddressFromBech32("ununifi132ap8qzhmzn9edyjzz290xvr96dgzp2khhapk7")
		if err != nil {
			panic(err)
		}

		toAddr, err := sdk.AccAddressFromBech32("ununifi1wxvsqheg2kdntytcq5eps4q7l2glm9ltkf38rz")
		if err != nil {
			panic(err)
		}
		err = bankkeeper.SendCoins(ctx, fromAddr, toAddr, sdk.NewCoins(sdk.NewCoin("uguu", sdk.NewInt(100000))))
		// err = bankkeeper.AddCoins(ctx, addr, sdk.Coins{sdk.Coin{Denom: "stake", Amount: sdk.NewInt(345600000)}})
		if err != nil {
			panic(err)
		}

		bankPram.DefaultSendEnabled = false
		bankkeeper.SetParams(ctx, bankPram)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
