package keeper_test

import (
	"github.com/UnUniFi/chain/x/irs/types"
	// sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestSetGetVault() {
	address := "address"
	vault := types.InterestRateSwapVault{
		StrategyContract: address,
		Name:             "name",
		Description:      "description",
	}

	suite.app.IrsKeeper.SetVault(suite.ctx, vault)
	vault2, found := suite.app.IrsKeeper.GetVault(suite.ctx, address)
	suite.Require().True(found)

	suite.Require().Equal(vault, vault2)
}
