package keeper_test

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/lcnem/jpyx/x/issuance/keeper"
	"github.com/lcnem/jpyx/x/issuance/types"
)

func (suite *KeeperTestSuite) TestQuerierGetParams() {
	querier := keeper.NewQuerier(suite.keeper)
	bz, err := querier(suite.ctx, []string{types.QueryGetParams}, abci.RequestQuery{})
	suite.Require().NoError(err)
	suite.NotNil(bz)

	var p types.Params
	suite.Nil(types.ModuleCdc.UnmarshalJSON(bz, &p))
	suite.Require().Equal(types.Params{Assets: types.Assets(nil)}, p)
}
