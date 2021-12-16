package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/ununifidist/keeper"
	ununifidisttypes "github.com/UnUniFi/chain/x/ununifidist/types"
)

type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	// supplyKeeper ununifidisttypes.SupplyKeeper
	accountKeeper ununifidisttypes.AccountKeeper
	bankKeeper    ununifidisttypes.BankKeeper
	app           app.TestApp
	ctx           sdk.Context
}

var (
	testPeriods = ununifidisttypes.Periods{
		ununifidisttypes.Period{
			Start:     time.Date(2020, time.March, 1, 1, 0, 0, 0, time.UTC),
			End:       time.Date(2021, time.March, 1, 1, 0, 0, 0, time.UTC),
			Inflation: sdk.MustNewDecFromStr("1.000000003022265980"),
		},
	}
)

func (suite *KeeperTestSuite) SetupTest() {
	// config := sdk.GetConfig()
	// app.SetBech32AddressPrefixes(config)
	tApp := app.NewTestApp()
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	coins := []sdk.Coins{sdk.NewCoins(sdk.NewCoin("uguu", sdk.NewInt(1000000000000)))}
	authGS := app.NewAuthGenState(
		tApp, addrs, coins)

	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

	params := ununifidisttypes.NewParams(true, testPeriods)
	guuGs := ununifidisttypes.NewGenesisState(params, ununifidisttypes.DefaultPreviousBlockTime, ununifidisttypes.DefaultGovDenom)
	// gs := app.GenesisState{ununifidisttypes.ModuleName: ununifidisttypes.ModuleCdc.MustMarshalJSON(ununifidisttypes.NewGenesisState(params, ununifidisttypes.DefaultPreviousBlockTime))}
	gs := app.GenesisState{ununifidisttypes.ModuleName: ununifidisttypes.ModuleCdc.MustMarshalJSON(&guuGs)}
	tApp.InitializeFromGenesisStates(
		authGS,
		gs,
	)
	// keeper := tApp.GetKavadistKeeper()
	// sk := tApp.GetSupplyKeeper()
	keeper := tApp.GetUnunifidistKeeper()
	sk := tApp.GetBankKeeper()
	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = keeper
	// suite.supplyKeeper = sk
	suite.accountKeeper = tApp.AccountKeeper
	suite.bankKeeper = sk
}

func (suite *KeeperTestSuite) TestMintExpiredPeriod() {
	govDenom, _ := suite.keeper.GetGovDenom(suite.ctx)
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	initialSupply := suite.bankKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(govDenom)
	suite.NotPanics(func() { suite.keeper.SetPreviousBlockTime(suite.ctx, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)) })
	ctx := suite.ctx.WithBlockTime(time.Date(2022, 1, 1, 0, 7, 0, 0, time.UTC))
	err := suite.keeper.MintPeriodInflation(ctx)
	suite.NoError(err)
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.Equal(initialSupply, finalSupply)
}

func (suite *KeeperTestSuite) TestMintPeriodNotStarted() {
	govDenom, _ := suite.keeper.GetGovDenom(suite.ctx)
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	initialSupply := suite.bankKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(govDenom)
	suite.NotPanics(func() { suite.keeper.SetPreviousBlockTime(suite.ctx, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)) })
	ctx := suite.ctx.WithBlockTime(time.Date(2019, 1, 1, 0, 7, 0, 0, time.UTC))
	err := suite.keeper.MintPeriodInflation(ctx)
	suite.NoError(err)
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.Equal(initialSupply, finalSupply)
}

func (suite *KeeperTestSuite) TestMintOngoingPeriod() {
	govDenom, _ := suite.keeper.GetGovDenom(suite.ctx)
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	initialSupply := suite.bankKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(govDenom)
	suite.NotPanics(func() {
		suite.keeper.SetPreviousBlockTime(suite.ctx, time.Date(2020, time.March, 1, 1, 0, 1, 0, time.UTC))
	})
	ctx := suite.ctx.WithBlockTime(time.Date(2021, 2, 28, 23, 59, 59, 0, time.UTC))
	err := suite.keeper.MintPeriodInflation(ctx)
	suite.NoError(err)
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.True(finalSupply.GT(initialSupply))
	// mAcc := suite.supplyKeeper.GetModuleAccount(ctx, ununifidisttypes.ModuleName)
	// mAccSupply := mAcc.GetCoins().AmountOf(ununifidisttypes.GovDenom)
	// suite.True(mAccSupply.Equal(finalSupply.Sub(initialSupply)))
	mAddr := suite.accountKeeper.GetModuleAddress(ununifidisttypes.ModuleName)
	mAddrSupply := suite.bankKeeper.GetAllBalances(ctx, mAddr).AmountOf(govDenom)
	suite.True(mAddrSupply.Equal(finalSupply.Sub(initialSupply)))
	// expect that inflation is ~10%
	expectedSupply := sdk.NewDecFromInt(initialSupply).Mul(sdk.MustNewDecFromStr("1.1"))
	supplyError := sdk.OneDec().Sub((sdk.NewDecFromInt(finalSupply).Quo(expectedSupply))).Abs()
	suite.True(supplyError.LTE(sdk.MustNewDecFromStr("0.001")))
}

func (suite *KeeperTestSuite) TestMintPeriodTransition() {
	govDenom, _ := suite.keeper.GetGovDenom(suite.ctx)
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	initialSupply := suite.bankKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(govDenom)
	params := suite.keeper.GetParams(suite.ctx)
	periods := ununifidisttypes.Periods{
		testPeriods[0],
		ununifidisttypes.Period{
			Start:     time.Date(2021, time.March, 1, 1, 0, 0, 0, time.UTC),
			End:       time.Date(2022, time.March, 1, 1, 0, 0, 0, time.UTC),
			Inflation: sdk.MustNewDecFromStr("1.000000003022265980"),
		},
	}
	params.Periods = periods
	suite.NotPanics(func() {
		suite.keeper.SetParams(suite.ctx, params)
	})
	suite.NotPanics(func() {
		suite.keeper.SetPreviousBlockTime(suite.ctx, time.Date(2020, time.March, 1, 1, 0, 1, 0, time.UTC))
	})
	ctx := suite.ctx.WithBlockTime(time.Date(2021, 3, 10, 0, 0, 0, 0, time.UTC))
	err := suite.keeper.MintPeriodInflation(ctx)
	suite.NoError(err)
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.True(finalSupply.GT(initialSupply))
}

func (suite *KeeperTestSuite) TestMintNotActive() {
	govDenom, _ := suite.keeper.GetGovDenom(suite.ctx)
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	initialSupply := suite.bankKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(govDenom)
	params := suite.keeper.GetParams(suite.ctx)
	params.Active = false
	suite.NotPanics(func() {
		suite.keeper.SetParams(suite.ctx, params)
	})
	suite.NotPanics(func() {
		suite.keeper.SetPreviousBlockTime(suite.ctx, time.Date(2020, time.March, 1, 1, 0, 1, 0, time.UTC))
	})
	ctx := suite.ctx.WithBlockTime(time.Date(2021, 2, 28, 23, 59, 59, 0, time.UTC))
	err := suite.keeper.MintPeriodInflation(ctx)
	suite.NoError(err)
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(ununifidisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.Equal(initialSupply, finalSupply)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
