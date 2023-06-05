package keeper_test

// import (
// 	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
// 	tmtime "github.com/cometbft/cometbft/types/time"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/stretchr/testify/suite"

// 	"github.com/UnUniFi/chain/app"
// 	"github.com/UnUniFi/chain/deprecated/cdp/keeper"
// )

// type KeeperTestSuite struct {
// 	suite.Suite

// 	keeper keeper.Keeper
// 	app    app.TestApp
// 	ctx    sdk.Context
// }

// func (suite *KeeperTestSuite) SetupTest() {
// 	suite.ResetChain()
// }

// func (suite *KeeperTestSuite) ResetChain() {
// 	tApp := app.NewTestApp()
// 	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
// 	keeper := tApp.GetCDPKeeper()

// 	suite.app = tApp
// 	suite.ctx = ctx
// 	suite.keeper = keeper
// }

// /*
// func (suite *KeeperTestSuite) TestGetSetSavingsRateDistributed() {
// 	suite.ResetChain()

// 	// Set savings rate distributed value
// 	savingsRateDist := sdk.NewInt(555000555000)
// 	suite.keeper.SetSavingsRateDistributed(suite.ctx, savingsRateDist)

// 	// Check store's savings rate distributed value
// 	s := suite.keeper.GetSavingsRateDistributed(suite.ctx)
// 	suite.Equal(savingsRateDist, s)
// }
// */
