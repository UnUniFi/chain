package keeper_test

import (
	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestAnnualizeYieldRate() {
	annualizedYieldRate := suite.keeper.AnnualizeYieldRate(suite.ctx, sdk.NewDec(4), 22, 44)

	// Check if the annualizedYieldRate was calculated
	suite.Require().Equal(annualizedYieldRate, sdk.NewDec(100))
}
