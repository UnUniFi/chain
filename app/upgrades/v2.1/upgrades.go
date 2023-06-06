package v2_1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	"github.com/UnUniFi/chain/app/keepers"
	"github.com/UnUniFi/chain/app/upgrades"
)

func CreateUpgradeHandler(mm *module.Manager,
	configurator module.Configurator,
	_ upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info(fmt.Sprintf("update start:%s", UpgradeName))

		// update 1 change wasm permission to everybody
		// Add wasmStack on ibcRouter
		// ibc patch version v7.0.1
		wasmParam := keepers.WasmKeeper.GetParams(ctx)
		wasmParam.CodeUploadAccess.Permission = wasmtypes.AccessTypeEverybody
		wasmParam.InstantiateDefaultPermission = wasmtypes.AccessTypeEverybody

		keepers.WasmKeeper.SetParams(ctx, wasmParam)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
