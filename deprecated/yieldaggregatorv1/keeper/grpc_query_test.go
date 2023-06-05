package keeper_test

// import (
// 	"github.com/cometbft/cometbft/crypto/ed25519"
// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/UnUniFi/chain/deprecated/yieldaggregatorv1/types"
// )

// func (suite *KeeperTestSuite) TestGRPCQueryAssetManagementAccount() {
// 	// get not available account
// 	_, err := suite.app.YieldaggregatorKeeper.AssetManagementAccount(sdk.WrapSDKContext(suite.ctx), &types.QueryAssetManagementAccountRequest{Id: "OsmoFarm"})
// 	suite.Require().NoError(err)

// 	// set asset management account
// 	osmoManagementAccount := types.AssetManagementAccount{
// 		Id:      "OsmoFarm",
// 		Name:    "Osmosis Farm",
// 		Enabled: true,
// 	}
// 	evmosManagementAccount := types.AssetManagementAccount{
// 		Id:      "EvmosFarm",
// 		Name:    "Evmos Farm",
// 		Enabled: false,
// 	}
// 	suite.app.YieldaggregatorKeeper.SetAssetManagementAccount(suite.ctx, osmoManagementAccount)
// 	suite.app.YieldaggregatorKeeper.SetAssetManagementAccount(suite.ctx, evmosManagementAccount)

// 	// check asset management accounts
// 	resp, err := suite.app.YieldaggregatorKeeper.AssetManagementAccount(sdk.WrapSDKContext(suite.ctx), &types.QueryAssetManagementAccountRequest{Id: "OsmoFarm"})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(resp.Account.Id, osmoManagementAccount.Id)
// 	suite.Require().Equal(resp.Account.Name, osmoManagementAccount.Name)
// 	resp, err = suite.app.YieldaggregatorKeeper.AssetManagementAccount(sdk.WrapSDKContext(suite.ctx), &types.QueryAssetManagementAccountRequest{Id: "EvmosFarm"})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(resp.Account.Id, evmosManagementAccount.Id)
// 	suite.Require().Equal(resp.Account.Name, evmosManagementAccount.Name)
// }

// func (suite *KeeperTestSuite) TestGRPCQueryAllAssetManagementAccounts() {
// 	// get all asset management accounts at initial
// 	resp, err := suite.app.YieldaggregatorKeeper.AllAssetManagementAccounts(sdk.WrapSDKContext(suite.ctx), &types.QueryAllAssetManagementAccountsRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(resp.Accounts, 0)

// 	// set asset management account
// 	osmoManagementAccount := types.AssetManagementAccount{
// 		Id:      "OsmoFarm",
// 		Name:    "Osmosis Farm",
// 		Enabled: true,
// 	}
// 	evmosManagementAccount := types.AssetManagementAccount{
// 		Id:      "EvmosFarm",
// 		Name:    "Evmos Farm",
// 		Enabled: false,
// 	}
// 	suite.app.YieldaggregatorKeeper.SetAssetManagementAccount(suite.ctx, osmoManagementAccount)
// 	suite.app.YieldaggregatorKeeper.SetAssetManagementAccount(suite.ctx, evmosManagementAccount)

// 	// check all asset management accounts
// 	resp, err = suite.app.YieldaggregatorKeeper.AllAssetManagementAccounts(sdk.WrapSDKContext(suite.ctx), &types.QueryAllAssetManagementAccountsRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(resp.Accounts, 2)
// }

// func (suite *KeeperTestSuite) TestGRPCQueryUserInfo() {
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

// 	// get initial user deposit
// 	resp, err := suite.app.YieldaggregatorKeeper.UserInfo(sdk.WrapSDKContext(suite.ctx), &types.QueryUserInfoRequest{
// 		Address: addr1.String(),
// 	})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(resp.UserInfo.Amount, []sdk.Coin{})
// 	suite.Require().Len(resp.UserInfo.FarmingOrders, 0)

// 	// set user deposit
// 	coins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 10000))
// 	suite.app.YieldaggregatorKeeper.SetUserDeposit(suite.ctx, addr1, coins)

// 	// check user deposit
// 	resp, err = suite.app.YieldaggregatorKeeper.UserInfo(sdk.WrapSDKContext(suite.ctx), &types.QueryUserInfoRequest{
// 		Address: addr1.String(),
// 	})
// 	suite.Require().NoError(err)
// 	suite.Require().Equal(resp.UserInfo.Amount, []sdk.Coin(coins))
// 	suite.Require().Len(resp.UserInfo.FarmingOrders, 0)
// }

// func (suite *KeeperTestSuite) TestGRPCQueryAllFarmingUnits() {
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	addr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

// 	// get not available unit
// 	resp, err := suite.app.YieldaggregatorKeeper.AllFarmingUnits(sdk.WrapSDKContext(suite.ctx), &types.QueryAllFarmingUnitsRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(resp.Units, 0)

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
// 	resp, err = suite.app.YieldaggregatorKeeper.AllFarmingUnits(sdk.WrapSDKContext(suite.ctx), &types.QueryAllFarmingUnitsRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(resp.Units, 2)
// }

// func (suite *KeeperTestSuite) TestGRPCDailyRewardPercents() {
// 	// get not available rate
// 	resp, err := suite.app.YieldaggregatorKeeper.DailyRewardPercents(sdk.WrapSDKContext(suite.ctx), &types.QueryDailyRewardPercentsRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(resp.DailyPercents, 0)

// 	// set rates
// 	percent1 := types.DailyPercent{
// 		AccountId: "OsmoFarm",
// 		TargetId:  "OsmoGUUFarm",
// 		Rate:      sdk.NewDecWithPrec(1, 1),
// 	}
// 	percent2 := types.DailyPercent{
// 		AccountId: "EvmosFarm",
// 		TargetId:  "EvmosGUUFarm",
// 		Rate:      sdk.NewDecWithPrec(1, 1),
// 	}
// 	suite.app.YieldaggregatorKeeper.SetDailyRewardPercent(suite.ctx, percent1)
// 	suite.app.YieldaggregatorKeeper.SetDailyRewardPercent(suite.ctx, percent2)

// 	// check rates
// 	resp, err = suite.app.YieldaggregatorKeeper.DailyRewardPercents(sdk.WrapSDKContext(suite.ctx), &types.QueryDailyRewardPercentsRequest{})
// 	suite.Require().NoError(err)
// 	suite.Require().Len(resp.DailyPercents, 2)
// }
