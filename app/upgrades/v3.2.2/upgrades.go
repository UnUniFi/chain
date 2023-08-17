package v3_2_2

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/app/keepers"
	"github.com/UnUniFi/chain/app/upgrades"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func CreateUpgradeHandler(mm *module.Manager,
	configurator module.Configurator,
	_ upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info(fmt.Sprintf("update start:%s", UpgradeName))

		// migrate vaults
		keepers.YieldaggregatorKeeper.MigrateAllLegacyVaults(ctx)
		// migrate strategies
		keepers.YieldaggregatorKeeper.MigrateAllLegacyStrategies(ctx)
		// migrate denoms
		balances := keepers.BankKeeper.GetAccountsBalances(ctx)
		for _, balance := range balances {
			for _, coin := range balance.Coins {
				denomParts := strings.Split(coin.Denom, "/")

				if len(denomParts) != 3 {
					continue
				}
				if denomParts[0] != "yield-aggregator" {
					continue
				}
				err := keepers.BankKeeper.SendCoinsFromAccountToModule(ctx, balance.GetAddress(), yieldaggregatortypes.ModuleName, sdk.NewCoins(coin))
				if err != nil {
					return vm, err
				}
				err = keepers.BankKeeper.BurnCoins(ctx, yieldaggregatortypes.ModuleName, sdk.NewCoins(coin))
				if err != nil {
					return vm, err
				}

				denomParts[0] = yieldaggregatortypes.ModuleName
				migratedCoin := sdk.NewCoin(strings.Join(denomParts, "/"), coin.Amount)
				err = keepers.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.Coins{migratedCoin})
				if err != nil {
					return vm, err
				}
				err = keepers.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, balance.GetAddress(), sdk.Coins{migratedCoin})
				if err != nil {
					return vm, err
				}
			}
		}
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
