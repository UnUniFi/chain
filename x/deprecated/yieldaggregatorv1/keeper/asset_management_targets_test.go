package keeper_test

import "github.com/UnUniFi/chain/x/yieldaggregatorv1/types"

func (suite *KeeperTestSuite) TestAssetManagementTargetGetSet() {
	// get not available target
	assetManagementTarget := suite.app.YieldaggregatorKeeper.GetAssetManagementTarget(suite.ctx, "OsmoFarm", "OsmoGUUFarm")
	suite.Require().Equal(assetManagementTarget, types.AssetManagementTarget{})

	// set asset management target
	osmoManagementTarget := types.AssetManagementTarget{
		Id:                       "OsmoGUUFarm",
		AssetManagementAccountId: "OsmoFarm",
		Enabled:                  true,
		AssetConditions: []types.AssetCondition{
			{
				Denom: "uguu",
				Ratio: 100,
				Min:   "",
			},
		},
	}
	evmosManagementTarget := types.AssetManagementTarget{
		Id:                       "EvmosGUUFarm",
		AssetManagementAccountId: "EvmosFarm",
		Enabled:                  false,
		AssetConditions: []types.AssetCondition{
			{
				Denom: "uguu",
				Ratio: 100,
				Min:   "",
			},
		},
	}
	suite.app.YieldaggregatorKeeper.SetAssetManagementTarget(suite.ctx, osmoManagementTarget)
	suite.app.YieldaggregatorKeeper.SetAssetManagementTarget(suite.ctx, evmosManagementTarget)

	// check asset management targets
	assetManagementTarget = suite.app.YieldaggregatorKeeper.GetAssetManagementTarget(suite.ctx, "OsmoFarm", "OsmoGUUFarm")
	suite.Require().Equal(assetManagementTarget, osmoManagementTarget)
	assetManagementTarget = suite.app.YieldaggregatorKeeper.GetAssetManagementTarget(suite.ctx, "EvmosFarm", "EvmosGUUFarm")
	suite.Require().Equal(assetManagementTarget, evmosManagementTarget)
	assetManagementTargets := suite.app.YieldaggregatorKeeper.GetAllAssetManagementTargets(suite.ctx)
	suite.Require().Len(assetManagementTargets, 2)
	assetManagementTargets = suite.app.YieldaggregatorKeeper.GetAssetManagementTargetsOfAccount(suite.ctx, "OsmoFarm")
	suite.Require().Len(assetManagementTargets, 1)
	assetManagementTargets = suite.app.YieldaggregatorKeeper.GetAssetManagementTargetsOfDenom(suite.ctx, "OsmoFarm", "uguu")
	suite.Require().Len(assetManagementTargets, 1)
	assetManagementTargets = suite.app.YieldaggregatorKeeper.GetAssetManagementTargetsOfDenom(suite.ctx, "OsmoFarm", "axxx")
	suite.Require().Len(assetManagementTargets, 0)

	// delete asset management target and check
	suite.app.YieldaggregatorKeeper.DeleteAssetManagementTarget(suite.ctx, "EvmosFarm", "EvmosGUUFarm")
	assetManagementTarget = suite.app.YieldaggregatorKeeper.GetAssetManagementTarget(suite.ctx, "EvmosFarm", "EvmosGUUFarm")
	suite.Require().Equal(assetManagementTarget, types.AssetManagementTarget{})
	assetManagementTargets = suite.app.YieldaggregatorKeeper.GetAllAssetManagementTargets(suite.ctx)
	suite.Require().Len(assetManagementTargets, 1)

	// update asset management account target and check
	osmoManagementTarget.AssetConditions[0].Denom = "aguu"
	suite.app.YieldaggregatorKeeper.UpdateAssetManagementTargetsOfAccount(suite.ctx, "OsmoFarm", []types.AssetManagementTarget{osmoManagementTarget})
	assetManagementTargets = suite.app.YieldaggregatorKeeper.GetAssetManagementTargetsOfDenom(suite.ctx, "OsmoFarm", "aguu")
	suite.Require().Len(assetManagementTargets, 1)
	assetManagementTargets = suite.app.YieldaggregatorKeeper.GetAssetManagementTargetsOfDenom(suite.ctx, "OsmoFarm", "uguu")
	suite.Require().Len(assetManagementTargets, 0)

	// delete asset management targets of an account and check
	suite.app.YieldaggregatorKeeper.DeleteAssetManagementTargetsOfAccount(suite.ctx, "OsmoFarm")
	assetManagementTargets = suite.app.YieldaggregatorKeeper.GetAssetManagementTargetsOfAccount(suite.ctx, "OsmoFarm")
	suite.Require().Len(assetManagementTargets, 0)
}
