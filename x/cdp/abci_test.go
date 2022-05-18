package cdp_test

import (
	"math/rand"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simulation "github.com/cosmos/cosmos-sdk/types/simulation"

	tmabcitypes "github.com/tendermint/tendermint/abci/types"
	tmtime "github.com/tendermint/tendermint/libs/time"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/UnUniFi/chain/app"
	auctiontypes "github.com/UnUniFi/chain/x/auction/types"
	cdp "github.com/UnUniFi/chain/x/cdp"
	cdpkeeper "github.com/UnUniFi/chain/x/cdp/keeper"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
)

type ModuleTestSuite struct {
	suite.Suite

	keeper       cdpkeeper.Keeper
	addrs        []sdk.AccAddress
	app          app.TestApp
	cdps         cdptypes.Cdps
	ctx          sdk.Context
	liquidations liquidationTracker
}

type liquidationTracker struct {
	xrp  []uint64
	btc  []uint64
	debt int64
}

func (suite *ModuleTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	coins := []sdk.Coins{}
	tracker := liquidationTracker{}

	for j := 0; j < 100; j++ {
		coins = append(coins, cs(c("btc", 100000000), c("xrp", 10000000000)))
	}
	_, addrs := app.GeneratePrivKeyAddressPairs(100)

	authGS := app.NewAuthGenState(
		tApp,
		addrs,
		coins,
	)
	tApp.InitializeFromGenesisStates(
		authGS,
		NewPricefeedGenStateMulti(tApp),
		NewCDPGenStateMulti(tApp),
	)
	suite.ctx = ctx
	suite.app = tApp
	suite.keeper = tApp.GetCDPKeeper()
	suite.cdps = cdptypes.Cdps{}
	suite.addrs = addrs
	suite.liquidations = tracker
}

func (suite *ModuleTestSuite) createCdps() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	cdps := make(cdptypes.Cdps, 100)
	_, addrs := app.GeneratePrivKeyAddressPairs(100)
	coins := []sdk.Coins{}
	tracker := liquidationTracker{}

	for j := 0; j < 100; j++ {
		coins = append(coins, cs(c("btc", 100000000), c("xrp", 10000000000)))
	}

	authGS := app.NewAuthGenState(
		tApp,
		addrs,
		coins,
	)
	tApp.InitializeFromGenesisStates(
		authGS,
		NewPricefeedGenStateMulti(tApp),
		NewCDPGenStateMulti(tApp),
	)

	suite.ctx = ctx
	suite.app = tApp
	suite.keeper = tApp.GetCDPKeeper()

	for j := 0; j < 100; j++ {
		collateral := "xrp"
		amount := 10000000000
		debt := simulation.RandIntBetween(rand.New(rand.NewSource(int64(j))), 750000000, 1249000000)
		if j%2 == 0 {
			collateral = "btc"
			amount = 100000000
			debt = simulation.RandIntBetween(rand.New(rand.NewSource(int64(j))), 2700000000, 5332000000)
			if debt >= 4000000000 {
				tracker.btc = append(tracker.btc, uint64(j+1))
				tracker.debt += int64(debt)
			}
		} else {
			if debt >= 1000000000 {
				tracker.xrp = append(tracker.xrp, uint64(j+1))
				tracker.debt += int64(debt)
			}
		}
		suite.Nil(suite.keeper.AddCdp(suite.ctx, addrs[j], c(collateral, int64(amount)), c("jpu", int64(debt)), collateral+"-a"))
		c, f := suite.keeper.GetCdp(suite.ctx, collateral+"-a", uint64(j+1))
		suite.True(f)
		cdps[j] = c
	}

	suite.cdps = cdps
	suite.addrs = addrs
	suite.liquidations = tracker
}

func (suite *ModuleTestSuite) setPrice(price sdk.Dec, market string) {
	pfKeeper := suite.app.GetPriceFeedKeeper()

	pfKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, market, price, suite.ctx.BlockTime().Add(time.Hour*3))
	err := pfKeeper.SetCurrentPrices(suite.ctx, market)
	suite.NoError(err)
	pp, err := pfKeeper.GetCurrentPrice(suite.ctx, market)
	suite.NoError(err)
	suite.Equal(price, pp.Price)
}
func (suite *ModuleTestSuite) TestBeginBlock() {
	suite.createCdps()
	ak := suite.app.GetAccountKeeper()
	bk := suite.app.GetBankKeeper()
	acc := ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	originalXrpCollateral := bk.GetAllBalances(suite.ctx, acc.GetAddress()).AmountOf("xrp")
	suite.setPrice(d("0.2"), "xrp:jpy")
	cdp.BeginBlocker(suite.ctx, tmabcitypes.RequestBeginBlock{Header: suite.ctx.BlockHeader()}, suite.keeper)
	acc = ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	finalXrpCollateral := bk.GetAllBalances(suite.ctx, acc.GetAddress()).AmountOf("xrp")
	seizedXrpCollateral := originalXrpCollateral.Sub(finalXrpCollateral)
	xrpLiquidations := int(seizedXrpCollateral.Quo(i(10000000000)).Int64())
	suite.Equal(10, xrpLiquidations)

	acc = ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	originalBtcCollateral := bk.GetAllBalances(suite.ctx, acc.GetAddress()).AmountOf("btc")
	suite.setPrice(d("6000"), "btc:jpy")
	cdp.BeginBlocker(suite.ctx, tmabcitypes.RequestBeginBlock{Header: suite.ctx.BlockHeader()}, suite.keeper)
	acc = ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	finalBtcCollateral := bk.GetAllBalances(suite.ctx, acc.GetAddress()).AmountOf("btc")
	seizedBtcCollateral := originalBtcCollateral.Sub(finalBtcCollateral)
	btcLiquidations := int(seizedBtcCollateral.Quo(i(100000000)).Int64())
	suite.Equal(10, btcLiquidations)

	acc = ak.GetModuleAccount(suite.ctx, auctiontypes.ModuleName)
	suite.Equal(int64(71955653865), bk.GetAllBalances(suite.ctx, acc.GetAddress()).AmountOf("debtjpu").Int64())

}

func (suite *ModuleTestSuite) TestSeizeSingleCdpWithFees() {
	err := suite.keeper.AddCdp(suite.ctx, suite.addrs[0], c("xrp", 10000000000), c("jpu", 1000000000), "xrp-a")
	suite.NoError(err)
	suite.Equal(i(1000000000), suite.keeper.GetTotalPrincipal(suite.ctx, "xrp-a", "jpu"))
	ak := suite.app.GetAccountKeeper()
	bk := suite.app.GetBankKeeper()
	cdpMacc := ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	suite.Equal(i(1000000000), bk.GetAllBalances(suite.ctx, cdpMacc.GetAddress()).AmountOf("debtjpu"))
	for i := 0; i < 100; i++ {
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * 6))
		cdp.BeginBlocker(suite.ctx, tmabcitypes.RequestBeginBlock{Header: suite.ctx.BlockHeader()}, suite.keeper)
	}

	cdpMacc = ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	suite.Equal(i(1000000891), bk.GetAllBalances(suite.ctx, cdpMacc.GetAddress()).AmountOf("debtjpu"))
	cdp, _ := suite.keeper.GetCdp(suite.ctx, "xrp-a", 1)

	err = suite.keeper.SeizeCollateral(suite.ctx, cdp)
	suite.NoError(err)
	_, found := suite.keeper.GetCdp(suite.ctx, "xrp-a", 1)
	suite.False(found)
}

// func TestModuleTestSuite(t *testing.T) {
// 	suite.Run(t, new(ModuleTestSuite))
// }
