package keeper_test

import (
	"errors"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/cdp/keeper"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
)

type DepositTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	app    app.TestApp
	ctx    sdk.Context
	addrs  []sdk.AccAddress
}

func (suite *DepositTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	_, addrs := app.GeneratePrivKeyAddressPairs(10)
	authGS := app.NewAuthGenState(
		tApp,
		addrs[0:2],
		[]sdk.Coins{
			cs(c("xrp", 500000000), c("btc", 500000000)),
			cs(c("xrp", 200000000))})
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

func (suite *DepositTestSuite) TestGetSetDeposit() {
	d, found := suite.keeper.GetDeposit(suite.ctx, uint64(1), suite.addrs[0])
	suite.True(found)
	td := cdptypes.NewDeposit(uint64(1), suite.addrs[0], c("xrp", 400000000))
	suite.True(d.Equals(td))
	ds := suite.keeper.GetDeposits(suite.ctx, uint64(1))
	suite.Equal(1, len(ds))
	suite.True(ds[0].Equals(td))
	suite.keeper.DeleteDeposit(suite.ctx, uint64(1), suite.addrs[0])
	_, found = suite.keeper.GetDeposit(suite.ctx, uint64(1), suite.addrs[0])
	suite.False(found)
	ds = suite.keeper.GetDeposits(suite.ctx, uint64(1))
	suite.Equal(0, len(ds))
}

func (suite *DepositTestSuite) TestDepositCollateral() {
	err := suite.keeper.DepositCollateral(suite.ctx, suite.addrs[0], suite.addrs[0], c("xrp", 10000000), "xrp-a")
	suite.NoError(err)
	d, found := suite.keeper.GetDeposit(suite.ctx, uint64(1), suite.addrs[0])
	suite.True(found)
	td := cdptypes.NewDeposit(uint64(1), suite.addrs[0], c("xrp", 410000000))
	suite.True(d.Equals(td))
	ds := suite.keeper.GetDeposits(suite.ctx, uint64(1))
	suite.Equal(1, len(ds))
	suite.True(ds[0].Equals(td))
	cd, _ := suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(1))
	suite.Equal(c("xrp", 410000000), cd.Collateral)
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	acc := ak.GetAccount(suite.ctx, suite.addrs[0])
	suite.Equal(i(90000000), sk.GetBalance(suite.ctx, acc.GetAddress(), "xrp").Amount)

	err = suite.keeper.DepositCollateral(suite.ctx, suite.addrs[0], suite.addrs[0], c("btc", 1), "btc-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrCdpNotFound))

	err = suite.keeper.DepositCollateral(suite.ctx, suite.addrs[1], suite.addrs[0], c("xrp", 1), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrCdpNotFound))

	err = suite.keeper.DepositCollateral(suite.ctx, suite.addrs[0], suite.addrs[1], c("xrp", 10000000), "xrp-a")
	suite.NoError(err)
	d, found = suite.keeper.GetDeposit(suite.ctx, uint64(1), suite.addrs[1])
	suite.True(found)
	td = cdptypes.NewDeposit(uint64(1), suite.addrs[1], c("xrp", 10000000))
	suite.True(d.Equals(td))
	ds = suite.keeper.GetDeposits(suite.ctx, uint64(1))
	suite.Equal(2, len(ds))
	suite.True(ds[1].Equals(td))
}

func (suite *DepositTestSuite) TestWithdrawCollateral() {
	err := suite.keeper.WithdrawCollateral(suite.ctx, suite.addrs[0], suite.addrs[0], c("xrp", 400000000), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrInvalidCollateralRatio))
	err = suite.keeper.WithdrawCollateral(suite.ctx, suite.addrs[0], suite.addrs[0], c("xrp", 321000000), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrInvalidCollateralRatio))
	err = suite.keeper.WithdrawCollateral(suite.ctx, suite.addrs[1], suite.addrs[0], c("xrp", 10000000), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrCdpNotFound))

	cd, _ := suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(1))
	cd.AccumulatedFees = c("jpu", 1)
	err = suite.keeper.SetCdp(suite.ctx, cd)
	suite.NoError(err)
	err = suite.keeper.WithdrawCollateral(suite.ctx, suite.addrs[0], suite.addrs[0], c("xrp", 320000000), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrInvalidCollateralRatio))

	err = suite.keeper.WithdrawCollateral(suite.ctx, suite.addrs[0], suite.addrs[0], c("xrp", 10000000), "xrp-a")
	suite.NoError(err)
	dep, _ := suite.keeper.GetDeposit(suite.ctx, uint64(1), suite.addrs[0])
	td := cdptypes.NewDeposit(uint64(1), suite.addrs[0], c("xrp", 390000000))
	suite.True(dep.Equals(td))
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	acc := ak.GetAccount(suite.ctx, suite.addrs[0])
	suite.Equal(i(110000000), sk.GetBalance(suite.ctx, acc.GetAddress(), "xrp").Amount)

	err = suite.keeper.WithdrawCollateral(suite.ctx, suite.addrs[0], suite.addrs[1], c("xrp", 10000000), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrDepositNotFound))
}

// func TestDepositTestSuite(t *testing.T) {
// 	suite.Run(t, new(DepositTestSuite))
// }
