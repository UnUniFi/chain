package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)

	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	params.RewardRateFeeders = []string{addr.String()}

	suite.app.YieldaggregatorKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)
	suite.Require().Equal(params, newParams)
}
