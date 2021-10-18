package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/lcnem/jpyx/app"
	"github.com/lcnem/jpyx/x/botanydist/keeper"
	botanydisttypes "github.com/lcnem/jpyx/x/botanydist/types"
)

type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	// supplyKeeper botanydisttypes.SupplyKeeper
	accountKeeper botanydisttypes.AccountKeeper
	bankKeeper    botanydisttypes.BankKeeper
	app           app.TestApp
	ctx           sdk.Context
}

var (
	testPeriods = botanydisttypes.Periods{
		botanydisttypes.Period{
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
	coins := []sdk.Coins{sdk.NewCoins(sdk.NewCoin("ujcbn", sdk.NewInt(1000000000000)))}
	authGS := app.NewAuthGenState(
		tApp, addrs, coins)

	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

	params := botanydisttypes.NewParams(true, testPeriods)
	jcbnGs := botanydisttypes.NewGenesisState(params, botanydisttypes.DefaultPreviousBlockTime, botanydisttypes.DefaultGovDenom)
	// gs := app.GenesisState{botanydisttypes.ModuleName: botanydisttypes.ModuleCdc.MustMarshalJSON(botanydisttypes.NewGenesisState(params, botanydisttypes.DefaultPreviousBlockTime))}
	gs := app.GenesisState{botanydisttypes.ModuleName: botanydisttypes.ModuleCdc.MustMarshalJSON(&jcbnGs)}
	tApp.InitializeFromGenesisStates(
		authGS,
		gs,
	)
	// keeper := tApp.GetKavadistKeeper()
	// sk := tApp.GetSupplyKeeper()
	keeper := tApp.GetBotanydistKeeper()
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
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
	initialSupply := suite.bankKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(govDenom)
	suite.NotPanics(func() { suite.keeper.SetPreviousBlockTime(suite.ctx, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)) })
	ctx := suite.ctx.WithBlockTime(time.Date(2022, 1, 1, 0, 7, 0, 0, time.UTC))
	err := suite.keeper.MintPeriodInflation(ctx)
	suite.NoError(err)
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.Equal(initialSupply, finalSupply)
}

func (suite *KeeperTestSuite) TestMintPeriodNotStarted() {
	govDenom, _ := suite.keeper.GetGovDenom(suite.ctx)
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
	initialSupply := suite.bankKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(govDenom)
	suite.NotPanics(func() { suite.keeper.SetPreviousBlockTime(suite.ctx, time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)) })
	ctx := suite.ctx.WithBlockTime(time.Date(2019, 1, 1, 0, 7, 0, 0, time.UTC))
	err := suite.keeper.MintPeriodInflation(ctx)
	suite.NoError(err)
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.Equal(initialSupply, finalSupply)
}

func (suite *KeeperTestSuite) TestMintOngoingPeriod() {
	govDenom, _ := suite.keeper.GetGovDenom(suite.ctx)
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
	initialSupply := suite.bankKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(govDenom)
	suite.NotPanics(func() {
		suite.keeper.SetPreviousBlockTime(suite.ctx, time.Date(2020, time.March, 1, 1, 0, 1, 0, time.UTC))
	})
	ctx := suite.ctx.WithBlockTime(time.Date(2021, 2, 28, 23, 59, 59, 0, time.UTC))
	err := suite.keeper.MintPeriodInflation(ctx)
	suite.NoError(err)
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.True(finalSupply.GT(initialSupply))
	// mAcc := suite.supplyKeeper.GetModuleAccount(ctx, botanydisttypes.ModuleName)
	// mAccSupply := mAcc.GetCoins().AmountOf(botanydisttypes.GovDenom)
	// suite.True(mAccSupply.Equal(finalSupply.Sub(initialSupply)))
	mAddr := suite.accountKeeper.GetModuleAddress(botanydisttypes.ModuleName)
	mAddrSupply := suite.bankKeeper.GetAllBalances(ctx, mAddr).AmountOf(govDenom)
	suite.True(mAddrSupply.Equal(finalSupply.Sub(initialSupply)))
	// expect that inflation is ~10%
	expectedSupply := sdk.NewDecFromInt(initialSupply).Mul(sdk.MustNewDecFromStr("1.1"))
	supplyError := sdk.OneDec().Sub((sdk.NewDecFromInt(finalSupply).Quo(expectedSupply))).Abs()
	suite.True(supplyError.LTE(sdk.MustNewDecFromStr("0.001")))
}

func (suite *KeeperTestSuite) TestMintPeriodTransition() {
	govDenom, _ := suite.keeper.GetGovDenom(suite.ctx)
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
	initialSupply := suite.bankKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(govDenom)
	params := suite.keeper.GetParams(suite.ctx)
	periods := botanydisttypes.Periods{
		testPeriods[0],
		botanydisttypes.Period{
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
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.True(finalSupply.GT(initialSupply))
}

func (suite *KeeperTestSuite) TestMintNotActive() {
	govDenom, _ := suite.keeper.GetGovDenom(suite.ctx)
	// initialSupply := suite.supplyKeeper.GetSupply(suite.ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
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
	// finalSupply := suite.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(botanydisttypes.GovDenom)
	finalSupply := suite.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(govDenom)
	suite.Equal(initialSupply, finalSupply)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
