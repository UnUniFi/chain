package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/app"
	auctiontypes "github.com/lcnem/jpyx/x/auction/types"

	cdpkeeper "github.com/lcnem/jpyx/x/cdp/keeper"
	cdptypes "github.com/lcnem/jpyx/x/cdp/types"

	"github.com/stretchr/testify/suite"

	"github.com/tendermint/tendermint/crypto"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

type AuctionTestSuite struct {
	suite.Suite

	keeper cdpkeeper.Keeper
	app    app.TestApp
	ctx    sdk.Context
	addrs  []sdk.AccAddress
}

func (suite *AuctionTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	taddr := sdk.AccAddress(crypto.AddressHash([]byte("KavaTestUser1")))
	authGS := app.NewAuthGenState(tApp, []sdk.AccAddress{taddr}, []sdk.Coins{cs(c("jsmn", 21000000000))})
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	tApp.InitializeFromGenesisStates(
		authGS,
		NewPricefeedGenStateMulti(),
		NewCDPGenStateMulti(),
	)
	keeper := tApp.GetCDPKeeper()
	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = keeper
	suite.addrs = []sdk.AccAddress{taddr}
}

func (suite *AuctionTestSuite) TestNetDebtSurplus() {
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 100)))
	suite.NoError(err)
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jsmn", 10)))
	suite.NoError(err)
	suite.NotPanics(func() { suite.keeper.NetSurplusAndDebt(suite.ctx) })
	acc := ak.GetModuleAccount(suite.ctx, cdptypes.LiquidatorMacc)
	suite.Equal(cs(c("debt", 90)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
}

func (suite *AuctionTestSuite) TestCollateralAuction() {
	sk := suite.app.GetBankKeeper()
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 21000000000), c("bnb", 190000000000)))
	suite.Require().NoError(err)
	testDeposit := cdptypes.NewDeposit(1, suite.addrs[0], c("bnb", 190000000000))
	err = suite.keeper.AuctionCollateral(suite.ctx, cdptypes.Deposits{testDeposit}, "bnb-a", i(21000000000), "jsmn")
	suite.Require().NoError(err)
}

func (suite *AuctionTestSuite) TestSurplusAuction() {
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jsmn", 600000000000)))
	suite.NoError(err)
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 100000000000)))
	suite.NoError(err)
	suite.keeper.RunSurplusAndDebtAuctions(suite.ctx)
	acc := ak.GetModuleAccount(suite.ctx, auctiontypes.ModuleName)
	suite.Equal(cs(c("jsmn", 10000000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
	acc = ak.GetModuleAccount(suite.ctx, cdptypes.LiquidatorMacc)
	suite.Equal(cs(c("jsmn", 490000000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
}

func (suite *AuctionTestSuite) TestDebtAuction() {
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jsmn", 100000000000)))
	suite.NoError(err)
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 200000000000)))
	suite.NoError(err)
	suite.keeper.RunSurplusAndDebtAuctions(suite.ctx)
	acc := ak.GetModuleAccount(suite.ctx, auctiontypes.ModuleName)
	suite.Equal(cs(c("debt", 10000000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
	acc = ak.GetModuleAccount(suite.ctx, cdptypes.LiquidatorMacc)
	suite.Equal(cs(c("debt", 90000000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
}

func (suite *AuctionTestSuite) TestGetTotalSurplus() {
	sk := suite.app.GetBankKeeper()

	// liquidator account has zero coins
	suite.Require().Equal(sdk.NewInt(0), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc))

	// mint some coins
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jsmn", 100e6)))
	suite.Require().NoError(err)
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jsmn", 200e6)))
	suite.Require().NoError(err)

	// liquidator account has 300e6 total jsmn
	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc))

	// mint some debt
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 500e6)))
	suite.Require().NoError(err)

	// liquidator account still has 300e6 total jsmn -- debt balance is ignored
	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc))

	// burn some jsmn
	err = sk.BurnCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jsmn", 50e6)))
	suite.Require().NoError(err)

	// liquidator jsmn decreases
	suite.Require().Equal(sdk.NewInt(250e6), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc))
}

func (suite *AuctionTestSuite) TestGetTotalDebt() {
	sk := suite.app.GetBankKeeper()

	// liquidator account has zero debt
	suite.Require().Equal(sdk.NewInt(0), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc))

	// mint some debt
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 100e6)))
	suite.Require().NoError(err)
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 200e6)))
	suite.Require().NoError(err)

	// liquidator account has 300e6 total debt
	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalDebt(suite.ctx, cdptypes.LiquidatorMacc))

	// mint some jsmn
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jsmn", 500e6)))
	suite.Require().NoError(err)

	// liquidator account still has 300e6 total debt -- jsmn balance is ignored
	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalDebt(suite.ctx, cdptypes.LiquidatorMacc))

	// burn some debt
	err = sk.BurnCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 50e6)))
	suite.Require().NoError(err)

	// liquidator debt decreases
	suite.Require().Equal(sdk.NewInt(250e6), suite.keeper.GetTotalDebt(suite.ctx, cdptypes.LiquidatorMacc))
}

func TestAuctionTestSuite(t *testing.T) {
	suite.Run(t, new(AuctionTestSuite))
}
