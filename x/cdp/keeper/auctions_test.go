package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/app"
	auctiontypes "github.com/UnUniFi/chain/x/auction/types"

	cdpkeeper "github.com/UnUniFi/chain/x/cdp/keeper"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"

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
	// todo: add euu
	authGS := app.NewAuthGenState(tApp, []sdk.AccAddress{taddr}, []sdk.Coins{cs(c("jpu", 21000000000))})
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	tApp.InitializeFromGenesisStates(
		authGS,
		NewPricefeedGenStateMulti(tApp),
		NewCDPGenStateMulti(tApp),
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
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 10)))
	suite.NoError(err)
	// todo: add euu test
	suite.NotPanics(func() { suite.keeper.NetSurplusAndDebt(suite.ctx) })
	acc := ak.GetModuleAccount(suite.ctx, cdptypes.LiquidatorMacc)
	suite.Equal(cs(c("debt", 90)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
}

func (suite *AuctionTestSuite) TestCollateralAuction() {
	sk := suite.app.GetBankKeeper()
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 21000000000), c("bnb", 190000000000)))
	suite.Require().NoError(err)
	testDeposit := cdptypes.NewDeposit(1, suite.addrs[0], c("bnb", 190000000000))
	err = suite.keeper.AuctionCollateral(suite.ctx, cdptypes.Deposits{testDeposit}, "bnb-a", i(21000000000), "jpu")
	suite.Require().NoError(err)
}

func (suite *AuctionTestSuite) TestSurplusAuction() {
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 600000000000)))
	suite.NoError(err)
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 100000000000)))
	suite.NoError(err)
	suite.keeper.RunSurplusAndDebtAuctions(suite.ctx)
	acc := ak.GetModuleAccount(suite.ctx, auctiontypes.ModuleName)
	suite.Equal(cs(c("jpu", 10000000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
	acc = ak.GetModuleAccount(suite.ctx, cdptypes.LiquidatorMacc)
	suite.Equal(cs(c("jpu", 490000000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))
}

func (suite *AuctionTestSuite) TestDebtAuction() {
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 100000000000)))
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
	suite.Require().Equal(sdk.NewInt(0), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "jpu"))

	// mint some coins
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 100e6)))
	suite.Require().NoError(err)
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 200e6)))
	suite.Require().NoError(err)

	// liquidator account has 300e6 total jpu
	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "jpu"))

	// mint some debt
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 500e6)))
	suite.Require().NoError(err)

	// liquidator account still has 300e6 total jpu -- debt balance is ignored
	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "jpu"))

	// burn some jpu
	err = sk.BurnCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 50e6)))
	suite.Require().NoError(err)

	// liquidator jpu decreases
	suite.Require().Equal(sdk.NewInt(250e6), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "jpu"))
}

func (suite *AuctionTestSuite) TestGetTotalDebt() {
	sk := suite.app.GetBankKeeper()

	// liquidator account has zero debt
	suite.Require().Equal(sdk.NewInt(0), suite.keeper.GetTotalSurplus(suite.ctx, cdptypes.LiquidatorMacc, "debt"))

	// mint some debt
	err := sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 100e6)))
	suite.Require().NoError(err)
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("debt", 200e6)))
	suite.Require().NoError(err)

	// liquidator account has 300e6 total debt
	suite.Require().Equal(sdk.NewInt(300e6), suite.keeper.GetTotalDebt(suite.ctx, cdptypes.LiquidatorMacc))

	// mint some jpu
	err = sk.MintCoins(suite.ctx, cdptypes.LiquidatorMacc, cs(c("jpu", 500e6)))
	suite.Require().NoError(err)

	// liquidator account still has 300e6 total debt -- jpu balance is ignored
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
