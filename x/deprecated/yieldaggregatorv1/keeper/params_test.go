package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)

	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	params.RewardRateFeeders = addr.String()

	suite.app.YieldaggregatorKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)
	suite.Require().Equal(params, newParams)
}
