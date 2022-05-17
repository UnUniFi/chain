package keeper_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmtime "github.com/tendermint/tendermint/libs/time"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/cdp/keeper"
	"github.com/UnUniFi/chain/x/cdp/types"
)

type DrawTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	app    app.TestApp
	ctx    sdk.Context
	addrs  []sdk.AccAddress
}

func (suite *DrawTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	_, addrs := app.GeneratePrivKeyAddressPairs(3)
	authGS := app.NewAuthGenState(
		tApp,
		addrs,
		[]sdk.Coins{
			cs(c("xrp", 500000000), c("btc", 500000000), c("jpu", 10000000000)),
			cs(c("xrp", 200000000)),
			cs(c("xrp", 10000000000000), c("jpu", 100000000000))})
	tApp.InitializeFromGenesisStates(
		authGS,
		NewPricefeedGenStateMulti(tApp),
		NewCDPGenStateMulti(tApp),
	)
	keeper := tApp.GetCDPKeeper()
	suite.app = tApp
	suite.keeper = keeper
	suite.ctx = ctx
	suite.addrs = addrs
	err := suite.keeper.AddCdp(suite.ctx, addrs[0], c("xrp", 400000000), c("jpu", 10000000), "xrp-a")
	suite.NoError(err)
}

func (suite *DrawTestSuite) TestAddRepayPrincipal() {

	err := suite.keeper.AddPrincipal(suite.ctx, suite.addrs[0], "xrp-a", c("jpu", 10000000))
	suite.NoError(err)

	t, found := suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(1))
	suite.True(found)
	suite.Equal(c("jpu", 20000000), t.Principal)
	ctd := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, t.Collateral, "xrp-a", t.Principal.Add(t.AccumulatedFees))
	suite.Equal(d("20.0"), ctd)
	ts := suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", d("20.0"))
	suite.Equal(0, len(ts))
	ts = suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", d("20.0").Add(sdk.SmallestDec()))
	suite.Equal(ts[0], t)
	tp := suite.keeper.GetTotalPrincipal(suite.ctx, "xrp-a", "jpu")
	suite.Equal(i(20000000), tp)
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	acc := ak.GetModuleAccount(suite.ctx, types.ModuleName)
	suite.Equal(cs(c("xrp", 400000000), c("debtjpu", 20000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))

	err = suite.keeper.AddPrincipal(suite.ctx, suite.addrs[0], "xrp-a", c("sjpy", 10000000))
	suite.Require().True(errors.Is(err, types.ErrInvalidDebtRequest))

	err = suite.keeper.AddPrincipal(suite.ctx, suite.addrs[1], "xrp-a", c("jpu", 10000000))
	suite.Require().True(errors.Is(err, types.ErrCdpNotFound))
	err = suite.keeper.AddPrincipal(suite.ctx, suite.addrs[0], "xrp-a", c("xjpy", 10000000))
	suite.Require().True(errors.Is(err, types.ErrInvalidDebtRequest))
	err = suite.keeper.AddPrincipal(suite.ctx, suite.addrs[0], "xrp-a", c("jpu", 311000000))
	suite.Require().True(errors.Is(err, types.ErrInvalidCollateralRatio))

	err = suite.keeper.RepayPrincipal(suite.ctx, suite.addrs[0], "xrp-a", c("jpu", 10000000))
	suite.NoError(err)

	t, found = suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(1))
	suite.True(found)
	suite.Equal(c("jpu", 10000000), t.Principal)

	ctd = suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, t.Collateral, "xrp-a", t.Principal.Add(t.AccumulatedFees))
	suite.Equal(d("40.0"), ctd)
	ts = suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", d("40.0"))
	suite.Equal(0, len(ts))
	ts = suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", d("40.0").Add(sdk.SmallestDec()))
	suite.Equal(ts[0], t)
	ak = suite.app.GetAccountKeeper()
	sk = suite.app.GetBankKeeper()
	acc = ak.GetModuleAccount(suite.ctx, types.ModuleName)
	suite.Equal(cs(c("xrp", 400000000), c("debtjpu", 10000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))

	err = suite.keeper.RepayPrincipal(suite.ctx, suite.addrs[0], "xrp-a", c("xjpy", 10000000))
	suite.Require().True(errors.Is(err, types.ErrInvalidPayment))
	err = suite.keeper.RepayPrincipal(suite.ctx, suite.addrs[1], "xrp-a", c("xjpy", 10000000))
	suite.Require().True(errors.Is(err, types.ErrCdpNotFound))

	err = suite.keeper.RepayPrincipal(suite.ctx, suite.addrs[0], "xrp-a", c("jpu", 9000000))
	suite.Require().True(errors.Is(err, types.ErrBelowDebtFloor))
	err = suite.keeper.RepayPrincipal(suite.ctx, suite.addrs[0], "xrp-a", c("jpu", 10000000))
	suite.NoError(err)

	_, found = suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(1))
	suite.False(found)
	ts = suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", types.MaxSortableDec)
	suite.Equal(0, len(ts))
	ts = suite.keeper.GetAllCdpsByCollateralType(suite.ctx, "xrp-a")
	suite.Equal(0, len(ts))
	ak = suite.app.GetAccountKeeper()
	sk = suite.app.GetBankKeeper()
	acc = ak.GetModuleAccount(suite.ctx, types.ModuleName)
	suite.Equal(sdk.Coins{}, sk.GetAllBalances(suite.ctx, acc.GetAddress()))

}

func (suite *DrawTestSuite) TestRepayPrincipalOverpay() {
	err := suite.keeper.RepayPrincipal(suite.ctx, suite.addrs[0], "xrp-a", c("jpu", 20000000))
	suite.NoError(err)
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	acc := ak.GetAccount(suite.ctx, suite.addrs[0])
	suite.Equal(i(10000000000), sk.GetBalance(suite.ctx, acc.GetAddress(), "jpu").Amount)
	_, found := suite.keeper.GetCdp(suite.ctx, "xrp-a", 1)
	suite.False(found)
}

/* Note: This test existed only on v0.11.1. On v0.13.0, this test is removed. So, I guess this test is unnecessary.
func (suite *DrawTestSuite) TestAddRepayPrincipalFees() {
	err := suite.keeper.AddCdp(suite.ctx, suite.addrs[2], c("xrp", 1000000000000), c("jpu", 100000000000), "xrp-a")
	suite.NoError(err)
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Minute * 10))
	err = suite.keeper.AccumulateInterest(suite.ctx, "xrp-a")
	suite.NoError(err)
	err = suite.keeper.SynchronizeInterestForRiskyCdps(suite.ctx, sdk.Int{}, sdk.MaxSortableDec, "xrp-a")
	suite.NoError(err)
	err = suite.keeper.AddPrincipal(suite.ctx, suite.addrs[2], "xrp-a", c("jpu", 10000000))
	suite.NoError(err)
	t, _ := suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(2))
	suite.Equal(c("jpu", 92827), t.AccumulatedFees)
	err = suite.keeper.RepayPrincipal(suite.ctx, suite.addrs[2], "xrp-a", c("jpu", 100))
	suite.NoError(err)
	t, _ = suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(2))
	suite.Equal(c("jpu", 92727), t.AccumulatedFees)
	err = suite.keeper.RepayPrincipal(suite.ctx, suite.addrs[2], "xrp-a", c("jpu", 100010092727))
	suite.NoError(err)
	_, f := suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(2))
	suite.False(f)

	err = suite.keeper.AddCdp(suite.ctx, suite.addrs[2], c("xrp", 1000000000000), c("jpu", 100000000), "xrp-a")
	suite.NoError(err)

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * 31536000)) // move forward one year in time
	err = suite.keeper.AccumulateInterest(suite.ctx, "xrp-a")
	suite.NoError(err)
	err = suite.keeper.SynchronizeInterestForRiskyCdps(suite.ctx, sdk.Int{}, sdk.MaxSortableDec, "xrp-a")
	suite.NoError(err)
	err = suite.keeper.AddPrincipal(suite.ctx, suite.addrs[2], "xrp-a", c("jpu", 100000000))
	suite.NoError(err)
	t, _ = suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(3))
	suite.Equal(c("jpu", 5000000), t.AccumulatedFees)
}
*/

func (suite *DrawTestSuite) TestPricefeedFailure() {
	ctx := suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour * 2))
	pfk := suite.app.GetPriceFeedKeeper()
	pfk.SetCurrentPrices(ctx, "xrp:jpy")
	err := suite.keeper.AddPrincipal(ctx, suite.addrs[0], "xrp-a", c("jpu", 10000000))
	suite.Error(err)
	err = suite.keeper.RepayPrincipal(ctx, suite.addrs[0], "xrp-a", c("jpu", 10000000))
	suite.NoError(err)
}

func (suite *DrawTestSuite) TestModuleAccountFailure() {
	suite.Panics(func() {
		ctx := suite.ctx.WithBlockHeader(suite.ctx.BlockHeader())
		ak := suite.app.GetAccountKeeper()
		acc := ak.GetModuleAccount(ctx, types.ModuleName)
		bk := suite.app.GetBankKeeper()
		bk.BurnCoins(ctx, types.ModuleName, bk.GetAllBalances(ctx, acc.GetAddress()))
		suite.keeper.RepayPrincipal(ctx, suite.addrs[0], "xrp-a", c("jpu", 10000000))
	})
}

func TestDrawTestSuite(t *testing.T) {
	suite.Run(t, new(DrawTestSuite))
}
