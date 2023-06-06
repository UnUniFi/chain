package keeper_test

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/UnUniFi/chain/app"
// 	auctiontypes "github.com/UnUniFi/chain/deprecated/x/auction/types"

// 	cdpkeeper "github.com/UnUniFi/chain/deprecated/x/cdp/keeper"
// 	cdptypes "github.com/UnUniFi/chain/deprecated/x/cdp/types"

// 	"github.com/stretchr/testify/suite"

// 	"github.com/cometbft/cometbft/crypto"
// 	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
// 	tmtime "github.com/cometbft/cometbft/types/time"
// )

// type AuctionTestSuite struct {
// 	suite.Suite

// 	keeper cdpkeeper.Keeper
// 	app    app.TestApp
// 	ctx    sdk.Context
// 	addrs  []sdk.AccAddress
// }

// func (suite *AuctionTestSuite) SetupTest() {
// 	tApp := app.NewTestApp()
// 	taddr := sdk.AccAddress(crypto.AddressHash([]byte("KavaTestUser1")))
// 	authGS := app.NewAuthGenState(tApp, []sdk.AccAddress{taddr}, []sdk.Coins{cs(c("jpu", 51000000000), c("euu", 21000000000))})
// 	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
// 	tApp.InitializeFromGenesisStates(
// 		authGS,
// 		NewPricefeedGenStateMulti(tApp),
// 		NewCDPGenStateMulti(tApp),
// 	)
// 	keeper := tApp.GetCDPKeeper()
// 	suite.app = tApp
// 	suite.ctx = ctx
// 	suite.keeper = keeper
// 	suite.addrs = []sdk.AccAddress{taddr}
// }

// func (suite *AuctionTestSuite) TestNetDebtSurplus() {
// 	ak := suite.app.GetAccountKeeper()
// 	sk := suite.app.GetBankKeeper()
// 	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", 100)))
// 	suite.NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debteuu", 100)))
// 	suite.NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 10)))
// 	suite.NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("euu", 20)))
// 	suite.NoError(err)
// 	suite.NotPanics(func() { suite.keeper.NetSurplusAndDebt(suite.ctx) })
// 	acc := ak.GetModuleAccount(suite.ctx, cdptypes.LiquidatorMacc)
// 	suite.Equal(cs(c("debtjpu", 90), c("debteuu", 80)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
// }

// func (suite *AuctionTestSuite) TestCollateralAuction() {
// 	sk := suite.app.GetBankKeeper()
// 	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", 21000000000), c("bnb", 190000000000)))
// 	suite.Require().NoError(err)
// 	testDeposit := cdptypes.NewDeposit(1, suite.addrs[0], c("bnb", 190000000000))
// 	err = suite.keeper.AuctionCollateral(suite.ctx, cdptypes.Deposits{testDeposit}, "bnb-a", i(21000000000), "jpu")
// 	suite.Require().NoError(err)
// }

// func (suite *AuctionTestSuite) TestSurplusAuction() {
// 	ak := suite.app.GetAccountKeeper()
// 	sk := suite.app.GetBankKeeper()
// 	var jpuNum int64 = 600_000_000_000
// 	var euuNum int64 = 900_000_000_000
// 	var debtJpuNum int64 = 100_000_000_000
// 	var debtEuuNum int64 = 200_000_000_000
// 	var err error
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", jpuNum)))
// 	suite.NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", debtJpuNum)))
// 	suite.NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("euu", euuNum)))
// 	suite.NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debteuu", debtEuuNum)))
// 	suite.NoError(err)
// 	suite.keeper.RunSurplusAndDebtAuctions(suite.ctx)
// 	acc := ak.GetModuleAccount(suite.ctx, auctiontypes.ModuleName)
// 	suite.Equal(cs(c("jpu", 10_000_000_000), c("euu", 10_000_000_000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
// 	acc = ak.GetModuleAccount(suite.ctx, cdptypes.LiquidatorMacc)
// 	suite.Equal(cs(c("jpu", 490_000_000_000), c("euu", 690_000_000_000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
// }

// func (suite *AuctionTestSuite) TestDebtAuction() {
// 	ak := suite.app.GetAccountKeeper()
// 	sk := suite.app.GetBankKeeper()
// 	var jpuNum int64 = 100_000_000_000
// 	var euuNum int64 = 200_000_000_000
// 	var debtJpuNum int64 = 200_000_000_000
// 	var debtEuuNum int64 = 900_000_000_000
// 	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", jpuNum)))
// 	suite.NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", debtJpuNum)))
// 	suite.NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("euu", euuNum)))
// 	suite.NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debteuu", debtEuuNum)))
// 	suite.NoError(err)
// 	suite.keeper.RunSurplusAndDebtAuctions(suite.ctx)
// 	acc := ak.GetModuleAccount(suite.ctx, auctiontypes.ModuleName)
// 	suite.Equal(cs(c("debtjpu", 10_000_000_000), c("debteuu", 10_000_000_000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
// 	acc = ak.GetModuleAccount(suite.ctx, cdptypes.LiquidatorMacc)
// 	suite.Equal(cs(c("debtjpu", 90_000_000_000), c("debteuu", 690_000_000_000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
// }

// func (suite *AuctionTestSuite) TestGetTotalSurplus() {
// 	sk := suite.app.GetBankKeeper()

// 	// liquidator account has zero coins
// 	suite.Require().Equal(sdk.NewInt(0), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "jpu"))

// 	// mint some coins
// 	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 100e6)))
// 	suite.Require().NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 200e6)))
// 	suite.Require().NoError(err)

// 	// liquidator account has 300e6 total jpu
// 	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "jpu"))

// 	// mint some debt
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", 500e6)))
// 	suite.Require().NoError(err)

// 	// liquidator account still has 300e6 total jpu -- debt balance is ignored
// 	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "jpu"))

// 	// burn some jpu
// 	err = sk.BurnCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 50e6)))
// 	suite.Require().NoError(err)

// 	// liquidator jpu decreases
// 	suite.Require().Equal(sdk.NewInt(250e6), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "jpu"))
// }

// func (suite *AuctionTestSuite) TestGetTotalDebt() {
// 	sk := suite.app.GetBankKeeper()

// 	// liquidator account has zero debt
// 	suite.Require().Equal(sdk.NewInt(0), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "debtjpu"))

// 	// mint some debt
// 	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", 100e6)))
// 	suite.Require().NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", 200e6)))
// 	suite.Require().NoError(err)

// 	// liquidator account has 300e6 total debt
// 	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalDebt(suite.ctx, cdptypes.LiquidatorMacc, "debtjpu"))

// 	// mint some jpu
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 500e6)))
// 	suite.Require().NoError(err)

// 	// liquidator account still has 300e6 total debt -- jpu balance is ignored
// 	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalDebt(suite.ctx, cdptypes.LiquidatorMacc, "debtjpu"))

// 	// burn some debt
// 	err = sk.BurnCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", 50e6)))
// 	suite.Require().NoError(err)

// 	// liquidator debt decreases
// 	suite.Require().Equal(sdk.NewInt(250e6), suite.keeper.GetTotalDebt(suite.ctx, cdptypes.LiquidatorMacc, "debtjpu"))
// }

// func (suite *AuctionTestSuite) TestGetTotalDenom() {
// 	sk := suite.app.GetBankKeeper()

// 	// liquidator account has zero debt
// 	suite.Require().Equal(sdk.NewInt(0), suite.keeper.TestGetTotalDenom(suite.ctx, cdptypes.LiquidatorMacc, "debtjpu"))

// 	// mint some debt
// 	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", 100e6)))
// 	suite.Require().NoError(err)
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", 200e6)))
// 	suite.Require().NoError(err)

// 	// liquidator account has 300e6 total debt
// 	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.TestGetTotalDenom(suite.ctx, cdptypes.LiquidatorMacc, "debtjpu"))

// 	// mint some jpu
// 	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 500e6)))
// 	suite.Require().NoError(err)

// 	// liquidator account still has 300e6 total debt -- jpu balance is ignored
// 	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.TestGetTotalDenom(suite.ctx, cdptypes.LiquidatorMacc, "debtjpu"))

// 	// burn some debt
// 	err = sk.BurnCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debtjpu", 50e6)))
// 	suite.Require().NoError(err)

// 	// liquidator debt decreases
// 	suite.Require().Equal(sdk.NewInt(250e6), suite.keeper.TestGetTotalDenom(suite.ctx, cdptypes.LiquidatorMacc, "debtjpu"))
// }

// // func TestAuctionTestSuite(t *testing.T) {
// // 	suite.Run(t, new(AuctionTestSuite))
// // }
