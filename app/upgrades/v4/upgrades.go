package v4

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

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

		// yieldaggregator params upgrade
		iyaParam, err := keepers.YieldaggregatorKeeper.GetParams(ctx)
		if err != nil {
			return vm, err
		}
		iyaParam.IbcTransferTimeoutNanos = 1800000000000 // 3min
		_ = keepers.YieldaggregatorKeeper.SetParams(ctx, iyaParam)

		// initialize DenomInfos, SymbolInfos, IntermediaryAccountInfo
		denomInfos := []yieldaggregatortypes.DenomInfo{}
		symbolInfos := []yieldaggregatortypes.SymbolInfo{}
		interAcc := yieldaggregatortypes.IntermediaryAccountInfo{}
		if ctx.ChainID() == "ununifi-beta-v1" { // mainnet
			denomInfos = []yieldaggregatortypes.DenomInfo{
				{ // ATOM.osmosis
					Denom:  "ibc/20D06D04E1BC1FAC482FECC06C2E2879A596904D64D8BA3285B4A3789DEAF910",
					Symbol: "ATOM",
					Channels: []yieldaggregatortypes.TransferChannel{
						{
							RecvChainId: "cosmoshub-4",
							SendChainId: "osmosis-1",
							ChannelId:   "channel-0",
						},
						{
							RecvChainId: "osmosis-1",
							SendChainId: "ununifi-beta-v1",
							ChannelId:   "channel-4",
						},
					},
				},
				{ // ATOM.cosmoshub
					Denom:  "ibc/25418646C017D377ADF3202FF1E43590D0DAE3346E594E8D78176A139A928F88",
					Symbol: "ATOM",
					Channels: []yieldaggregatortypes.TransferChannel{
						{
							RecvChainId: "cosmoshub-4",
							SendChainId: "ununifi-beta-v1",
							ChannelId:   "channel-7",
						},
					},
				},
				{ // OSMO.osmosis
					Denom:  "ibc/05AC4BBA78C5951339A47DD1BC1E7FC922A9311DF81C85745B1C162F516FF2F1",
					Symbol: "OSMO",
					Channels: []yieldaggregatortypes.TransferChannel{
						{
							RecvChainId: "osmosis-1",
							SendChainId: "ununifi-beta-v1",
							ChannelId:   "channel-4",
						},
					},
				},
			}

			symbolInfos = []yieldaggregatortypes.SymbolInfo{
				{
					Symbol:        "ATOM",
					NativeChainId: "cosmoshub-4",
					Channels: []yieldaggregatortypes.TransferChannel{
						{
							SendChainId: "cosmoshub-4",
							RecvChainId: "osmosis-1",
							ChannelId:   "channel-141",
						},
						{
							SendChainId: "cosmoshub-4",
							RecvChainId: "ununifi-beta-v1",
							ChannelId:   "channel-683",
						},
					},
				},
				{
					Symbol:        "OSMO",
					NativeChainId: "osmosis-1",
					Channels: []yieldaggregatortypes.TransferChannel{
						{
							SendChainId: "cosmoshub-4",
							RecvChainId: "ununifi-beta-v1",
							ChannelId:   "channel-683",
						},
					},
				},
			}

			interAcc = yieldaggregatortypes.IntermediaryAccountInfo{
				Addrs: []yieldaggregatortypes.ChainAddress{
					{
						ChainId: "cosmoshub-4",
						Address: "cosmos1fvhcnyddukcqfnt7nlwv3thm5we22lyxyxylr9h77cvgkcn43xfs60ggw8",
					},
					{
						ChainId: "osmosis-1",
						Address: "osmo1fvhcnyddukcqfnt7nlwv3thm5we22lyxyxylr9h77cvgkcn43xfs0jssep",
					},
				},
			}
		}

		for _, denomInfo := range denomInfos {
			keepers.YieldaggregatorKeeper.SetDenomInfo(ctx, denomInfo)
		}

		for _, symbolInfo := range symbolInfos {
			keepers.YieldaggregatorKeeper.SetSymbolInfo(ctx, symbolInfo)
		}

		keepers.YieldaggregatorKeeper.SetIntermediaryAccountInfo(ctx, interAcc.Addrs)

		// migrate vaults
		keepers.YieldaggregatorKeeper.MigrateAllLegacyVaults(ctx)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
