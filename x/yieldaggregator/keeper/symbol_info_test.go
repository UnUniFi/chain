package keeper_test

import "github.com/UnUniFi/chain/x/yieldaggregator/types"

func (suite *KeeperTestSuite) TestSymbolInfoStore() {
	symbolInfos := []types.SymbolInfo{
		{
			Symbol:        "ATOM",
			NativeChainId: "cosmoshub-4",
			Channels: []types.TransferChannel{
				{
					SendChainId: "cosmoshut-4",
					RecvChainId: "osmosis-1",
					ChannelId:   "channel-1",
				},
			},
		},
		{
			Symbol:        "OSMO",
			NativeChainId: "osmosis-1",
			Channels: []types.TransferChannel{
				{
					SendChainId: "osmosis-1",
					RecvChainId: "cosmoshub-4",
					ChannelId:   "channel-2",
				},
			},
		},
	}

	for _, symbolInfo := range symbolInfos {
		suite.app.YieldaggregatorKeeper.SetSymbolInfo(suite.ctx, symbolInfo)
	}

	for _, symbolInfo := range symbolInfos {
		r := suite.app.YieldaggregatorKeeper.GetSymbolInfo(suite.ctx, symbolInfo.Symbol)
		suite.Require().Equal(r, symbolInfo)
	}

	storedInfos := suite.app.YieldaggregatorKeeper.GetAllSymbolInfo(suite.ctx)
	suite.Require().Len(storedInfos, 2)
}
