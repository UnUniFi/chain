package keeper_test

import (
	"errors"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/cdp/keeper"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
)

type CdpTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	app    app.TestApp
	ctx    sdk.Context
}

func (suite *CdpTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	tApp.InitializeFromGenesisStates(
		NewPricefeedGenStateMulti(tApp),
		NewCDPGenStateMulti(tApp),
	)
	keeper := tApp.GetCDPKeeper()
	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = keeper
}

func (suite *CdpTestSuite) TestAddCdp() {
	_, addrs := app.GeneratePrivKeyAddressPairs(2)
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	acc := ak.NewAccountWithAddress(suite.ctx, addrs[0])
	sk.GetAllBalances(suite.ctx, acc.GetAddress())
	testutil.FundAccount(suite.app.BankKeeper, suite.ctx, acc.GetAddress(), cs(c("xrp", 200000000), c("btc", 500000000)))
	ak.SetAccount(suite.ctx, acc)
	err := suite.keeper.AddCdp(suite.ctx, addrs[0], c("xrp", 200000000), c("jpu", 10000000), "btc-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrInvalidCollateral))
	err = suite.keeper.AddCdp(suite.ctx, addrs[0], c("xrp", 200000000), c("jpu", 26000000), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrInvalidCollateralRatio))
	err = suite.keeper.AddCdp(suite.ctx, addrs[0], c("xrp", 500000000), c("jpu", 26000000), "xrp-a")
	suite.Error(err) // insufficient balance
	err = suite.keeper.AddCdp(suite.ctx, addrs[0], c("xrp", 200000000), c("xjpy", 10000000), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrDebtNotSupported))

	acc2 := ak.NewAccountWithAddress(suite.ctx, addrs[1])
	testutil.FundAccount(suite.app.BankKeeper, suite.ctx, acc2.GetAddress(), cs(c("btc", 500000000000)))
	ak.SetAccount(suite.ctx, acc2)
	err = suite.keeper.AddCdp(suite.ctx, addrs[1], c("btc", 500000000000), c("jpu", 500000000001), "btc-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrExceedsDebtLimit))

	ctx := suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Hour * 2))
	pk := suite.app.GetPriceFeedKeeper()
	err = pk.SetCurrentPrices(ctx, "xrp:jpy")
	suite.Error(err)
	ok := suite.keeper.UpdatePricefeedStatus(ctx, "xrp:jpy")
	suite.False(ok)
	err = suite.keeper.AddCdp(ctx, addrs[0], c("xrp", 100000000), c("jpu", 10000000), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrPricefeedDown))

	err = pk.SetCurrentPrices(suite.ctx, "xrp:jpy")
	ok = suite.keeper.UpdatePricefeedStatus(suite.ctx, "xrp:jpy")
	suite.True(ok)
	suite.NoError(err)
	err = suite.keeper.AddCdp(suite.ctx, addrs[0], c("xrp", 100000000), c("jpu", 10000000), "xrp-a")
	suite.NoError(err)
	id := suite.keeper.GetNextCdpID(suite.ctx)
	suite.Equal(uint64(2), id)
	tp := suite.keeper.GetTotalPrincipal(suite.ctx, "xrp-a", "jpu")
	suite.Equal(i(10000000), tp)

	macc := ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	suite.Equal(cs(c("debtjpu", 10000000), c("xrp", 100000000)), sk.GetAllBalances(suite.ctx, macc.GetAddress()))
	acc = ak.GetAccount(suite.ctx, addrs[0])
	suite.Equal(cs(c("jpu", 10000000), c("xrp", 100000000), c("btc", 500000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))

	err = suite.keeper.AddCdp(suite.ctx, addrs[0], c("btc", 500000000), c("jpu", 26667000000), "btc-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrInvalidCollateralRatio))

	err = suite.keeper.AddCdp(suite.ctx, addrs[0], c("btc", 500000000), c("jpu", 100000000), "btc-a")
	suite.NoError(err)
	id = suite.keeper.GetNextCdpID(suite.ctx)
	suite.Equal(uint64(3), id)
	tp = suite.keeper.GetTotalPrincipal(suite.ctx, "btc-a", "jpu")
	suite.Equal(i(100000000), tp)
	macc = ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	suite.Equal(cs(c("debtjpu", 110000000), c("xrp", 100000000), c("btc", 500000000)), sk.GetAllBalances(suite.ctx, macc.GetAddress()))
	acc = ak.GetAccount(suite.ctx, addrs[0])
	suite.Equal(cs(c("jpu", 110000000), c("xrp", 100000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))

	err = suite.keeper.AddCdp(suite.ctx, addrs[0], c("lol", 100), c("jpu", 10), "lol-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrCollateralNotSupported))
	err = suite.keeper.AddCdp(suite.ctx, addrs[0], c("xrp", 100), c("jpu", 10), "xrp-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrCdpAlreadyExists))
}

func (suite *CdpTestSuite) TestAddGetCdp() {
	_, addrs := app.GeneratePrivKeyAddressPairs(2)
	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	acc := ak.NewAccountWithAddress(suite.ctx, addrs[0])
	sk.GetAllBalances(suite.ctx, acc.GetAddress())
	testutil.FundAccount(suite.app.BankKeeper, suite.ctx, acc.GetAddress(), cs(c("xrp", 200000000), c("btc", 500000000)))
	ak.SetAccount(suite.ctx, acc)
	err := suite.keeper.AddCdp(suite.ctx, addrs[0], c("btc", 100000000), c("jpu", 10000000), "btc-a")
	suite.NoError(err)
	id := suite.keeper.GetNextCdpID(suite.ctx)
	suite.Equal(uint64(2), id)
	_, found := suite.keeper.GetCdp(suite.ctx, "btc-a", 1)
	suite.True(found)
	_, found2 := suite.keeper.GetCdpByOwnerAndCollateralType(suite.ctx, addrs[0], "btc-a")
	suite.True(found2)
}

func (suite *CdpTestSuite) TestGetSetCollateralTypeByte() {
	_, found := suite.keeper.GetCollateralTypePrefix(suite.ctx, "lol-a")
	suite.False(found)
	db, found := suite.keeper.GetCollateralTypePrefix(suite.ctx, "xrp-a")
	suite.True(found)
	suite.Equal(byte(0x20), db)
}

func (suite *CdpTestSuite) TestGetDebtDenomMap() {
	denomMap := cdptypes.NewDebtDenomMap(cdptypes.DefaultDebtParams)
	suite.keeper.SetDebtDenomMap(suite.ctx, denomMap)
	t := suite.keeper.GetDebtDenomMap(suite.ctx)
	suite.Equal(denomMap, t)

	empty := cdptypes.DebtDenomMap{}
	suite.keeper.SetDebtDenomMap(suite.ctx, empty)
	t = suite.keeper.GetDebtDenomMap(suite.ctx)
	suite.Equal(empty, t)
}

func (suite *CdpTestSuite) TestGetNextCdpID() {
	id := suite.keeper.GetNextCdpID(suite.ctx)
	suite.Equal(cdptypes.DefaultCdpStartingID, id)
}

func (suite *CdpTestSuite) TestGetSetCdp() {
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	cdp := cdptypes.NewCdp(cdptypes.DefaultCdpStartingID, addrs[0], c("xrp", 1), "xrp-a", c("jpu", 1), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	err := suite.keeper.SetCdp(suite.ctx, cdp)
	suite.NoError(err)

	t, found := suite.keeper.GetCdp(suite.ctx, "xrp-a", cdptypes.DefaultCdpStartingID)
	suite.True(found)
	suite.Equal(cdp, t)
	_, found = suite.keeper.GetCdp(suite.ctx, "xrp-a", uint64(2))
	suite.False(found)
	suite.keeper.DeleteCdp(suite.ctx, cdp)
	_, found = suite.keeper.GetCdp(suite.ctx, "btc-a", cdptypes.DefaultCdpStartingID)
	suite.False(found)
}

func (suite *CdpTestSuite) TestGetSetCdpId() {
	_, addrs := app.GeneratePrivKeyAddressPairs(2)
	cdp := cdptypes.NewCdp(cdptypes.DefaultCdpStartingID, addrs[0], c("xrp", 1), "xrp-a", c("jpu", 1), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	err := suite.keeper.SetCdp(suite.ctx, cdp)
	suite.NoError(err)
	suite.keeper.IndexCdpByOwner(suite.ctx, cdp)
	id, found := suite.keeper.GetCdpID(suite.ctx, addrs[0], "xrp-a")
	suite.True(found)
	suite.Equal(cdptypes.DefaultCdpStartingID, id)
	_, found = suite.keeper.GetCdpID(suite.ctx, addrs[0], "lol-a")
	suite.False(found)
	_, found = suite.keeper.GetCdpID(suite.ctx, addrs[1], "xrp-a")
	suite.False(found)
}

func (suite *CdpTestSuite) TestGetSetCdpByOwnerAndCollateralType() {
	_, addrs := app.GeneratePrivKeyAddressPairs(2)
	cdp := cdptypes.NewCdp(cdptypes.DefaultCdpStartingID, addrs[0], c("xrp", 1), "xrp-a", c("jpu", 1), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	err := suite.keeper.SetCdp(suite.ctx, cdp)
	suite.NoError(err)
	suite.keeper.IndexCdpByOwner(suite.ctx, cdp)
	t, found := suite.keeper.GetCdpByOwnerAndCollateralType(suite.ctx, addrs[0], "xrp-a")
	suite.True(found)
	suite.Equal(cdp, t)
	_, found = suite.keeper.GetCdpByOwnerAndCollateralType(suite.ctx, addrs[0], "lol-a")
	suite.False(found)
	_, found = suite.keeper.GetCdpByOwnerAndCollateralType(suite.ctx, addrs[1], "xrp-a")
	suite.False(found)
	suite.NotPanics(func() { suite.keeper.IndexCdpByOwner(suite.ctx, cdp) })
}

func (suite *CdpTestSuite) TestCalculateCollateralToDebtRatio() {
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	cdp := cdptypes.NewCdp(cdptypes.DefaultCdpStartingID, addrs[0], c("xrp", 3), "xrp-a", c("jpu", 1), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	cr := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, cdp.Collateral, cdp.Type, cdp.Principal)
	suite.Equal(sdk.MustNewDecFromStr("3.0"), cr)
	cdp = cdptypes.NewCdp(cdptypes.DefaultCdpStartingID, addrs[0], c("xrp", 1), "xrp-a", c("jpu", 2), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	cr = suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, cdp.Collateral, cdp.Type, cdp.Principal)
	suite.Equal(sdk.MustNewDecFromStr("0.5"), cr)
}

func (suite *CdpTestSuite) TestSetCdpByCollateralRatio() {
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	cdp := cdptypes.NewCdp(cdptypes.DefaultCdpStartingID, addrs[0], c("xrp", 3), "xrp-a", c("jpu", 1), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	cr := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, cdp.Collateral, cdp.Type, cdp.Principal)
	suite.NotPanics(func() { suite.keeper.IndexCdpByCollateralRatio(suite.ctx, cdp.Type, cdp.Id, cr) })
}

func (suite *CdpTestSuite) TestIterateCdps() {
	cdps := cdps()
	for _, c := range cdps {
		err := suite.keeper.SetCdp(suite.ctx, c)
		suite.NoError(err)
		suite.keeper.IndexCdpByOwner(suite.ctx, c)
		cr := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, c.Collateral, c.Type, c.Principal)
		suite.keeper.IndexCdpByCollateralRatio(suite.ctx, c.Type, c.Id, cr)
	}
	t := suite.keeper.GetAllCdps(suite.ctx)
	suite.Equal(8, len(t))
}

func (suite *CdpTestSuite) TestIterateCdpsByCollateralType() {
	cdps := cdps()
	for _, c := range cdps {
		err := suite.keeper.SetCdp(suite.ctx, c)
		suite.NoError(err)
		suite.keeper.IndexCdpByOwner(suite.ctx, c)
		cr := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, c.Collateral, c.Type, c.Principal)
		suite.keeper.IndexCdpByCollateralRatio(suite.ctx, c.Type, c.Id, cr)
	}
	xrpCdps := suite.keeper.GetAllCdpsByCollateralType(suite.ctx, "xrp-a")
	suite.Equal(3, len(xrpCdps))
	xrpCdps = suite.keeper.GetAllCdpsByCollateralType(suite.ctx, "xrp-b")
	suite.Equal(3, len(xrpCdps))
	btcCdps := suite.keeper.GetAllCdpsByCollateralType(suite.ctx, "btc-a")
	suite.Equal(1, len(btcCdps))
	btcCdps = suite.keeper.GetAllCdpsByCollateralType(suite.ctx, "btc-b")
	suite.Equal(1, len(btcCdps))
	suite.keeper.DeleteCdp(suite.ctx, cdps[0])
	suite.keeper.RemoveCdpOwnerIndex(suite.ctx, cdps[0])
	xrpCdps = suite.keeper.GetAllCdpsByCollateralType(suite.ctx, "xrp-a")
	suite.Equal(2, len(xrpCdps))
	suite.keeper.DeleteCdp(suite.ctx, cdps[1])
	suite.keeper.RemoveCdpOwnerIndex(suite.ctx, cdps[1])
	ids, found := suite.keeper.GetCdpIdsByOwner(suite.ctx, cdps[1].Owner.AccAddress())
	suite.True(found)
	suite.Equal(1, len(ids))
	suite.Equal(uint64(3), ids[0])
}

func (suite *CdpTestSuite) TestIterateCdpsByCollateralRatio() {
	cdps := cdps()
	for _, c := range cdps {
		err := suite.keeper.SetCdp(suite.ctx, c)
		suite.NoError(err)
		suite.keeper.IndexCdpByOwner(suite.ctx, c)
		cr := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, c.Collateral, c.Type, c.Principal)
		suite.keeper.IndexCdpByCollateralRatio(suite.ctx, c.Type, c.Id, cr)
	}
	xrpCdps := suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", d("1.25"))
	suite.Equal(0, len(xrpCdps))
	xrpCdps = suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", d("1.25").Add(sdk.SmallestDec()))
	suite.Equal(1, len(xrpCdps))
	xrpCdps = suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", d("2.0").Add(sdk.SmallestDec()))
	suite.Equal(2, len(xrpCdps))
	xrpCdps = suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", d("100.0").Add(sdk.SmallestDec()))
	suite.Equal(3, len(xrpCdps))
	suite.keeper.DeleteCdp(suite.ctx, cdps[0])
	suite.keeper.RemoveCdpOwnerIndex(suite.ctx, cdps[0])
	cr := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, cdps[0].Collateral, cdps[0].Type, cdps[0].Principal)
	suite.keeper.RemoveCdpCollateralRatioIndex(suite.ctx, cdps[0].Type, cdps[0].Id, cr)
	xrpCdps = suite.keeper.GetAllCdpsByCollateralTypeAndRatio(suite.ctx, "xrp-a", d("2.0").Add(sdk.SmallestDec()))
	suite.Equal(1, len(xrpCdps))
}

func (suite *CdpTestSuite) TestValidateCollateral() {
	c := sdk.NewCoin("xrp", sdk.NewInt(1))
	err := suite.keeper.ValidateCollateral(suite.ctx, c, "xrp-a")
	suite.NoError(err)
	c = sdk.NewCoin("lol", sdk.NewInt(1))
	err = suite.keeper.ValidateCollateral(suite.ctx, c, "lol-a")
	suite.Require().True(errors.Is(err, cdptypes.ErrCollateralNotSupported))
}

func (suite *CdpTestSuite) TestValidatePrincipal() {
	d := sdk.NewCoin("jpu", sdk.NewInt(10000000))
	err := suite.keeper.ValidatePrincipalAdd(suite.ctx, d)
	suite.NoError(err)
	d = sdk.NewCoin("xjpy", sdk.NewInt(1))
	err = suite.keeper.ValidatePrincipalAdd(suite.ctx, d)
	suite.Require().True(errors.Is(err, cdptypes.ErrDebtNotSupported))
	d = sdk.NewCoin("jpu", sdk.NewInt(1000000000001))
	err = suite.keeper.ValidateDebtLimit(suite.ctx, "xrp-a", d)
	suite.Require().True(errors.Is(err, cdptypes.ErrExceedsDebtLimit))
	d = sdk.NewCoin("jpu", sdk.NewInt(100000000))
	err = suite.keeper.ValidateDebtLimit(suite.ctx, "xrp-a", d)
	suite.NoError(err)
}

func (suite *CdpTestSuite) TestCalculateCollateralizationRatio() {
	c := cdps()[1]
	err := suite.keeper.SetCdp(suite.ctx, c)
	suite.NoError(err)
	suite.keeper.IndexCdpByOwner(suite.ctx, c)
	cr := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, c.Collateral, c.Type, c.Principal)
	suite.keeper.IndexCdpByCollateralRatio(suite.ctx, c.Type, c.Id, cr)
	cr, err = suite.keeper.CalculateCollateralizationRatio(suite.ctx, c.Collateral, c.Type, c.Principal, c.AccumulatedFees, "spot")
	suite.NoError(err)
	suite.Equal(d("2.5"), cr)
	c.AccumulatedFees = sdk.NewCoin("jpu", i(10000000))
	cr, err = suite.keeper.CalculateCollateralizationRatio(suite.ctx, c.Collateral, c.Type, c.Principal, c.AccumulatedFees, "spot")
	suite.NoError(err)
	suite.Equal(d("1.25"), cr)
}

func (suite *CdpTestSuite) TestMintBurnDebtCoins() {
	cd := cdps()[1]
	denomMap := suite.keeper.GetDebtDenomMap(suite.ctx)
	err := suite.keeper.MintDebtCoins(suite.ctx, cdptypes.ModuleName, denomMap[cd.Principal.Denom], cd.Principal)
	suite.NoError(err)
	suite.Require().Panics(func() {
		_ = suite.keeper.MintDebtCoins(suite.ctx, "notamodule", denomMap[cd.Principal.Denom], cd.Principal)
	})

	ak := suite.app.GetAccountKeeper()
	sk := suite.app.GetBankKeeper()
	acc := ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	suite.Equal(cs(c("debtjpu", 10000000)), sk.GetAllBalances(suite.ctx, acc.GetAddress()))

	err = suite.keeper.BurnDebtCoins(suite.ctx, cdptypes.ModuleName, denomMap[cd.Principal.Denom], cd.Principal)
	suite.NoError(err)
	suite.Require().Panics(func() {
		_ = suite.keeper.BurnDebtCoins(suite.ctx, "notamodule", denomMap[cd.Principal.Denom], cd.Principal)
	})
	acc = ak.GetModuleAccount(suite.ctx, cdptypes.ModuleName)
	suite.Equal(sdk.Coins{}, sk.GetAllBalances(suite.ctx, acc.GetAddress()))
}

// func TestCdpTestSuite(t *testing.T) {
// 	suite.Run(t, new(CdpTestSuite))
// }
