package keeper_test

import (
	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestSetRemainingMargin() {
	// owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	margin := sdk.Coin{
		Denom:  "uusd",
		Amount: sdk.NewInt(100),
	}

	suite.keeper.SetRemainingMargin(suite.ctx, "positionId", margin)

	// Check if the margin was set
	remainingMargin := suite.keeper.GetRemainingMargin(suite.ctx, "positionId")

	suite.Require().Equal(remainingMargin, &margin)
}
