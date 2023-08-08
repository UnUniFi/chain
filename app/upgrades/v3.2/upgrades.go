package v3_2

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

		backedloanParam := keepers.NftbackedloanKeeper.GetParamSet(ctx)
		backedloanParam.BidTokens = []string{"uguu"}
		backedloanParam.NftListingCancelRequiredSeconds = 20
		backedloanParam.BidCancelRequiredSeconds = 20
		backedloanParam.NftListingFullPaymentPeriod = 60 * 60 * 24 * 2           // 2days
		backedloanParam.NftListingNftDeliveryPeriod = 60 * 60 * 24 * 1           // 1day
		backedloanParam.NftListingCommissionRate = sdk.MustNewDecFromStr("0.05") // 5%
		keepers.NftbackedloanKeeper.SetParamSet(ctx, backedloanParam)

		factoryParam := keepers.NftfactoryKeeper.GetParamSet(ctx)
		factoryParam.MaxNFTSupplyCap = 100000
		factoryParam.MinClassNameLen = 3
		factoryParam.MaxClassNameLen = 128
		factoryParam.MinUriLen = 8
		factoryParam.MaxUriLen = 512
		factoryParam.MaxDescriptionLen = 16
		factoryParam.MaxDescriptionLen = 1024
		keepers.NftfactoryKeeper.SetParamSet(ctx, factoryParam)

		return vm, nil
	}
}
