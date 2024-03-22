package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params, err := suite.app.IrsKeeper.GetParams(suite.ctx)
	suite.Require().NoError(err)
	params.TradeFeeRate = sdk.NewDecWithPrec(1, 1)
	suite.app.IrsKeeper.SetParams(suite.ctx, params)
	newParams, err := suite.app.IrsKeeper.GetParams(suite.ctx)
	suite.Require().NoError(err)

	suite.Require().Equal(params, newParams)
}
