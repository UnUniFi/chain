package keeper_test

import (
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (suite *KeeperTestSuite) TestAssetManagementAccountGetSet() {
	// get not available account
	assetManagementAccount := suite.app.YieldaggregatorKeeper.GetAssetManagementAccount(suite.ctx, "OsmoFarm")
	suite.Require().Equal(assetManagementAccount, types.AssetManagementAccount{})

	// set asset management account
	osmoManagementAccount := types.AssetManagementAccount{
		Id:      "OsmoFarm",
		Name:    "Osmosis Farm",
		Enabled: true,
	}
	evmosManagementAccount := types.AssetManagementAccount{
		Id:      "EvmosFarm",
		Name:    "Evmos Farm",
		Enabled: false,
	}
	suite.app.YieldaggregatorKeeper.SetAssetManagementAccount(suite.ctx, osmoManagementAccount)
	suite.app.YieldaggregatorKeeper.SetAssetManagementAccount(suite.ctx, evmosManagementAccount)

	// check asset management accounts
	assetManagementAccount = suite.app.YieldaggregatorKeeper.GetAssetManagementAccount(suite.ctx, "OsmoFarm")
	suite.Require().Equal(assetManagementAccount, osmoManagementAccount)
	assetManagementAccount = suite.app.YieldaggregatorKeeper.GetAssetManagementAccount(suite.ctx, "EvmosFarm")
	suite.Require().Equal(assetManagementAccount, evmosManagementAccount)
	assetManagementAccounts := suite.app.YieldaggregatorKeeper.GetAllAssetManagementAccounts(suite.ctx)
	suite.Require().Len(assetManagementAccounts, 2)

	// delete user deposit and check
	suite.app.YieldaggregatorKeeper.DeleteAssetManagementAccount(suite.ctx, "EvmosFarm")
	assetManagementAccount = suite.app.YieldaggregatorKeeper.GetAssetManagementAccount(suite.ctx, "EvmosFarm")
	suite.Require().Equal(assetManagementAccount, types.AssetManagementAccount{})
}

func (suite *KeeperTestSuite) TestAddAssetManagementAccount() {
	// get not available account
	assetManagementAccount := suite.app.YieldaggregatorKeeper.GetAssetManagementAccount(suite.ctx, "OsmoFarm")
	suite.Require().Equal(assetManagementAccount, types.AssetManagementAccount{})

	// set asset management account
	err := suite.app.YieldaggregatorKeeper.AddAssetManagementAccount(suite.ctx, "OsmoFarm", "Osmosis Farm")
	suite.Require().NoError(err)
	err = suite.app.YieldaggregatorKeeper.AddAssetManagementAccount(suite.ctx, "EvmosFarm", "Evmos Farm")
	suite.Require().NoError(err)

	// check asset management accounts
	osmoManagementAccount := types.AssetManagementAccount{
		Id:      "OsmoFarm",
		Name:    "Osmosis Farm",
		Enabled: true,
	}
	evmosManagementAccount := types.AssetManagementAccount{
		Id:      "EvmosFarm",
		Name:    "Evmos Farm",
		Enabled: true,
	}
	assetManagementAccount = suite.app.YieldaggregatorKeeper.GetAssetManagementAccount(suite.ctx, "OsmoFarm")
	suite.Require().Equal(assetManagementAccount, osmoManagementAccount)
	assetManagementAccount = suite.app.YieldaggregatorKeeper.GetAssetManagementAccount(suite.ctx, "EvmosFarm")
	suite.Require().Equal(assetManagementAccount, evmosManagementAccount)
	assetManagementAccounts := suite.app.YieldaggregatorKeeper.GetAllAssetManagementAccounts(suite.ctx)
	suite.Require().Len(assetManagementAccounts, 2)

	// try adding same account
	err = suite.app.YieldaggregatorKeeper.AddAssetManagementAccount(suite.ctx, "OsmoFarm", "Osmosis Farm")
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestUpdateAssetManagementAccount() {
	// try updating when not available
	osmoManagementAccount := types.AssetManagementAccount{
		Id:      "OsmoFarm",
		Name:    "Osmosis Farm",
		Enabled: true,
	}
	err := suite.app.YieldaggregatorKeeper.UpdateAssetManagementAccount(suite.ctx, osmoManagementAccount)
	suite.Require().Error(err)

	// check asset management accounts
	err = suite.app.YieldaggregatorKeeper.AddAssetManagementAccount(suite.ctx, "OsmoFarm", "")
	suite.Require().NoError(err)
	assetManagementAccount := suite.app.YieldaggregatorKeeper.GetAssetManagementAccount(suite.ctx, "OsmoFarm")
	suite.Require().Equal(assetManagementAccount.Name, "")

	// update after addition
	err = suite.app.YieldaggregatorKeeper.UpdateAssetManagementAccount(suite.ctx, osmoManagementAccount)
	suite.Require().NoError(err)
	assetManagementAccount = suite.app.YieldaggregatorKeeper.GetAssetManagementAccount(suite.ctx, "OsmoFarm")
	suite.Require().Equal(assetManagementAccount, osmoManagementAccount)
}
