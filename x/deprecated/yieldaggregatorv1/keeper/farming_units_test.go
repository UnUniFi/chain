package keeper_test

// import (
// 	"github.com/cometbft/cometbft/crypto/ed25519"
// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/UnUniFi/chain/x/deprecated/yieldaggregatorv1/types"
// )

// func (suite *KeeperTestSuite) TestFarmingUnitGetSet() {
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	addr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

// 	// get not available unit
// 	farmingUnit := suite.app.YieldaggregatorKeeper.GetFarmingUnit(suite.ctx, addr1.String(), "OsmoFarm", "OsmoGUUFarm")
// 	suite.Require().Equal(farmingUnit, types.FarmingUnit{})

// 	// set farming units
// 	unit1 := types.FarmingUnit{
// 		AccountId: "OsmoFarm",
// 		TargetId:  "OsmoGUUFarm",
// 		Owner:     addr1.String(),
// 	}
// 	unit2 := types.FarmingUnit{
// 		AccountId: "OsmoFarm",
// 		TargetId:  "OsmoGUUFarm",
// 		Owner:     addr2.String(),
// 	}
// 	suite.app.YieldaggregatorKeeper.SetFarmingUnit(suite.ctx, unit1)
// 	suite.app.YieldaggregatorKeeper.SetFarmingUnit(suite.ctx, unit2)

// 	// check farming units
// 	farmingUnit = suite.app.YieldaggregatorKeeper.GetFarmingUnit(suite.ctx, addr1.String(), "OsmoFarm", "OsmoGUUFarm")
// 	suite.Require().Equal(farmingUnit, unit1)
// 	farmingUnit = suite.app.YieldaggregatorKeeper.GetFarmingUnit(suite.ctx, addr2.String(), "OsmoFarm", "OsmoGUUFarm")
// 	suite.Require().Equal(farmingUnit, unit2)
// 	farmingUnits := suite.app.YieldaggregatorKeeper.GetAllFarmingUnits(suite.ctx)
// 	suite.Require().Len(farmingUnits, 2)
// 	farmingUnits = suite.app.YieldaggregatorKeeper.GetFarmingUnitsOfAddress(suite.ctx, addr1)
// 	suite.Require().Len(farmingUnits, 1)

// 	// delete farming unit and check
// 	suite.app.YieldaggregatorKeeper.DeleteFarmingUnit(suite.ctx, unit2)
// 	farmingUnit = suite.app.YieldaggregatorKeeper.GetFarmingUnit(suite.ctx, addr2.String(), "OsmoFarm", "OsmoGUUFarm")
// 	suite.Require().Equal(farmingUnit, types.FarmingUnit{})
// 	farmingUnits = suite.app.YieldaggregatorKeeper.GetAllFarmingUnits(suite.ctx)
// 	suite.Require().Len(farmingUnits, 1)

// 	// add farming unit and check
// 	err := suite.app.YieldaggregatorKeeper.AddFarmingUnit(suite.ctx, unit2)
// 	suite.Require().NoError(err)
// 	farmingUnits = suite.app.YieldaggregatorKeeper.GetAllFarmingUnits(suite.ctx)
// 	suite.Require().Len(farmingUnits, 2)

// 	// try adding once more to check error
// 	err = suite.app.YieldaggregatorKeeper.AddFarmingUnit(suite.ctx, unit2)
// 	suite.Require().Error(err)
// }
