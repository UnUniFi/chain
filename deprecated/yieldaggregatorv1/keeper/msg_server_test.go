package keeper_test

// import (
// 	"github.com/cometbft/cometbft/crypto/ed25519"
// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/UnUniFi/chain/deprecated/yieldaggregatorv1/keeper"
// 	"github.com/UnUniFi/chain/deprecated/yieldaggregatorv1/types"
// )

// func (suite *KeeperTestSuite) TestMsgServerSetDailyRewardPercent() {
// 	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

// 	// when set by not a feeder
// 	msgServer := keeper.NewMsgServerImpl(suite.app.YieldaggregatorKeeper)
// 	msg := types.NewMsgSetDailyRewardPercent(addr, "acc1", "tar1", sdk.NewDec(1), 1662433360)
// 	_, err := msgServer.SetDailyRewardPercent(sdk.WrapSDKContext(suite.ctx), &msg)
// 	suite.Require().Error(err)

// 	// when set by a feeder
// 	params := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)
// 	params.RewardRateFeeders = addr.String()
// 	suite.app.YieldaggregatorKeeper.SetParams(suite.ctx, params)
// 	_, err = msgServer.SetDailyRewardPercent(sdk.WrapSDKContext(suite.ctx), &msg)
// 	suite.Require().NoError(err)

// 	// check result
// 	percent := suite.app.YieldaggregatorKeeper.GetDailyRewardPercent(suite.ctx, "acc1", "tar1")
// 	suite.Require().Equal(percent.Rate, msg.Rate)
// 	suite.Require().Equal(percent.Date, msg.Date)
// }
