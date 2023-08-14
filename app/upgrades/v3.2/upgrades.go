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

		backedloanParam := keepers.NftbackedloanKeeper.GetParamSet(ctx)
		backedloanParam.BidTokens = []string{"uguu"}
		backedloanParam.NftListingCancelRequiredSeconds = 20
		backedloanParam.BidCancelRequiredSeconds = 20
		backedloanParam.NftListingFullPaymentPeriod = 60 * 60 * 24 * 2           // 2days
		backedloanParam.NftListingNftDeliveryPeriod = 60 * 60 * 24 * 1           // 1day
		backedloanParam.NftListingCommissionRate = sdk.MustNewDecFromStr("0.05") // 5%
		keepers.NftbackedloanKeeper.SetParamSet(ctx, backedloanParam)

		factoryParam := keepers.NftfactoryKeeper.GetParams(ctx)
		factoryParam.ClassCreationFee = sdk.Coins{}
		factoryParam.FeeCollectorAddress = ""
		keepers.NftfactoryKeeper.SetParams(ctx, factoryParam)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
