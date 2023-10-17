package keeper_test

import "github.com/UnUniFi/chain/x/yieldaggregator/types"

func (suite *KeeperTestSuite) TestDenomInfoStore() {
	denomInfos := []types.DenomInfo{
		{
			Denom:  "ibc/01AAFF",
			Symbol: "ATOM",
			Channels: []types.TransferChannel{
				{
					ChainId:   "cosmoshub-4",
					ChannelId: "channel-1",
				},
			},
		},
		{
			Denom:  "ibc/11AAFF",
			Symbol: "ATOM",
			Channels: []types.TransferChannel{
				{
					ChainId:   "cosmoshub-4",
					ChannelId: "channel-2",
				},
			},
		},
		{
			Denom:  "ibc/21AAFF",
			Symbol: "ATOM",
			Channels: []types.TransferChannel{
				{
					ChainId:   "osmosis-1",
					ChannelId: "channel-1",
				},
				{
					ChainId:   "cosmoshub-4",
					ChannelId: "channel-2",
				},
			},
		},
	}

	for _, denomInfo := range denomInfos {
		suite.app.YieldaggregatorKeeper.SetDenomInfo(suite.ctx, denomInfo)
	}

	for _, denomInfo := range denomInfos {
		r := suite.app.YieldaggregatorKeeper.GetDenomInfo(suite.ctx, denomInfo.Denom)
		suite.Require().Equal(r, denomInfo)
	}

	storedInfos := suite.app.YieldaggregatorKeeper.GetAllDenomInfo(suite.ctx)
	suite.Require().Len(storedInfos, 3)
}
