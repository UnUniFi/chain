package keeper_test

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)
	suite.app.YieldaggregatorKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)
	suite.Require().Equal(params, newParams)
}
