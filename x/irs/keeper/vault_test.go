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

func (suite *KeeperTestSuite) TestRemoveVault() {
	address := "address"
	vault := types.InterestRateSwapVault{
		StrategyContract: address,
		Name:             "name",
		Description:      "description",
	}

	suite.app.IrsKeeper.SetVault(suite.ctx, vault)
	suite.app.IrsKeeper.RemoveVault(suite.ctx, address)
	_, found := suite.app.IrsKeeper.GetVault(suite.ctx, address)
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestGetAllVault() {
	vaults := []types.InterestRateSwapVault{
		{
			StrategyContract: "address01",
			Name:             "name01",
			Description:      "description01",
		},
		{
			StrategyContract: "address02",
			Name:             "name02",
			Description:      "description02",
		},
	}

	for _, vault := range vaults {
		suite.app.IrsKeeper.SetVault(suite.ctx, vault)
	}

	vaults2 := suite.app.IrsKeeper.GetAllVault(suite.ctx)
	suite.Require().Equal(vaults, vaults2)
}
