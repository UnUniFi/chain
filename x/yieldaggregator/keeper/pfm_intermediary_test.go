package keeper_test

import "github.com/UnUniFi/chain/x/yieldaggregator/types"

func (suite *KeeperTestSuite) TestIntermediaryAccountStore() {
	interAcc := types.IntermediaryAccountInfo{
		Addrs: []types.ChainAddress{
			{
				ChainId: "osmosis-1",
				Address: "osmo1aqvlxpk8dc4m2nkmxkf63a5zez9jkzgm6amkgddhfk0qj9j4rw3q662wuk",
			},
			{
				ChainId: "cosmoshub-4",
				Address: "cosmos1q7q90qsl9g0gl2zz0njxwv2a649yqrtyxtnv3v",
			},
		},
	}

	suite.app.YieldaggregatorKeeper.SetIntermediaryAccountInfo(suite.ctx, interAcc.Addrs)

	r := suite.app.YieldaggregatorKeeper.GetIntermediaryAccountInfo(suite.ctx)
	suite.Require().Equal(r, interAcc)

	hubAddress := suite.app.YieldaggregatorKeeper.GetIntermediaryReceiver(suite.ctx, "cosmoshub-4")
	suite.Require().Equal(hubAddress, "cosmos1q7q90qsl9g0gl2zz0njxwv2a649yqrtyxtnv3v")
}
