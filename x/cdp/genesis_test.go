package cdp_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	tmtime "github.com/tendermint/tendermint/types/time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/app"
	cdpkeeper "github.com/lcnem/jpyx/x/cdp/keeper"
	cdptypes "github.com/lcnem/jpyx/x/cdp/types"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx    sdk.Context
	keeper cdpkeeper.Keeper
}

func (suite *GenesisTestSuite) TestInvalidGenState() {
	type args struct {
		params          cdptypes.Params
		cdps            cdptypes.Cdps
		deposits        cdptypes.Deposits
		startingID      uint64
		debtDenom       string
		govDenom        string
		prevDistTime    time.Time
		savingsRateDist sdk.Int
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}
	type genesisTest struct {
		name    string
		args    args
		errArgs errArgs
	}
	testCases := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			name: "empty debt denom",
			args: args{
				params:          cdptypes.DefaultParams(),
				cdps:            cdptypes.Cdps{},
				deposits:        cdptypes.Deposits{},
				debtDenom:       "",
				govDenom:        cdptypes.DefaultGovDenom,
				prevDistTime:    tmtime.Canonical(time.Unix(0, 0)), // cdptypes.DefaultPreviousDistributionTime,
				savingsRateDist: sdk.NewInt(0),                     // cdptypes.DefaultSavingsRateDistributed,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "debt denom invalid",
			},
		},
		{
			name: "empty gov denom",
			args: args{
				params:          cdptypes.DefaultParams(),
				cdps:            cdptypes.Cdps{},
				deposits:        cdptypes.Deposits{},
				debtDenom:       cdptypes.DefaultDebtDenom,
				govDenom:        "",
				prevDistTime:    tmtime.Canonical(time.Unix(0, 0)), // cdptypes.DefaultPreviousDistributionTime,
				savingsRateDist: sdk.NewInt(0),                     // cdptypes.DefaultSavingsRateDistributed,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "gov denom invalid",
			},
		},
		{
			name: "empty distribution time",
			args: args{
				params:          cdptypes.DefaultParams(),
				cdps:            cdptypes.Cdps{},
				deposits:        cdptypes.Deposits{},
				debtDenom:       cdptypes.DefaultDebtDenom,
				govDenom:        cdptypes.DefaultGovDenom,
				prevDistTime:    time.Time{},
				savingsRateDist: sdk.NewInt(0), // cdptypes.DefaultSavingsRateDistributed,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "previous distribution time not set",
			},
		},
		{
			name: "negative savings rate distributed",
			args: args{
				params:          cdptypes.DefaultParams(),
				cdps:            cdptypes.Cdps{},
				deposits:        cdptypes.Deposits{},
				debtDenom:       cdptypes.DefaultDebtDenom,
				govDenom:        cdptypes.DefaultGovDenom,
				prevDistTime:    tmtime.Canonical(time.Unix(0, 0)), // cdptypes.DefaultPreviousDistributionTime,
				savingsRateDist: sdk.NewInt(-100),
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "savings rate distributed should not be negative",
			},
		},
	}
	for _, tc := range testCases {
		prevDistTimes := cdptypes.GenesisAccumulationTimes{}
		totalPrincipals := cdptypes.GenesisTotalPrincipals{}
		suite.Run(tc.name, func() {
			prevDistTime := cdptypes.GenesisAccumulationTime{
				CollateralType:           tc.args.cdps.String(),
				PreviousAccumulationTime: tc.args.prevDistTime,
				InterestFactor:           sdk.NewDec(tc.args.savingsRateDist.Int64()),
			}
			prevDistTimes = append(prevDistTimes, prevDistTime)
			totalPrincipal := cdptypes.GenesisTotalPrincipal{
				CollateralType: tc.args.cdps.String(),
				TotalPrincipal: sdk.NewInt(tc.args.savingsRateDist.Int64()),
			}
			totalPrincipals := append(totalPrincipals, totalPrincipal)
			gs := cdptypes.NewGenesisState(tc.args.params, tc.args.cdps, tc.args.deposits, tc.args.startingID,
				tc.args.debtDenom, tc.args.govDenom, prevDistTimes, totalPrincipals)
			err := gs.Validate()
			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.T().Log(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *GenesisTestSuite) TestValidGenState() {
	tApp := app.NewTestApp()

	suite.NotPanics(func() {
		tApp.InitializeFromGenesisStates(
			NewPricefeedGenStateMulti(),
			NewCDPGenStateMulti(),
		)
	})

	cdpGS := NewCDPGenStateMulti()
	gs := cdptypes.GenesisState{}
	cdptypes.ModuleCdc.UnmarshalJSON(cdpGS["cdp"], &gs)
	gs.Cdps = cdps()
	gs.StartingCdpId = uint64(5)
	appGS := app.GenesisState{"cdp": cdptypes.ModuleCdc.MustMarshalJSON(&gs)}
	suite.NotPanics(func() {
		tApp.InitializeFromGenesisStates(
			NewPricefeedGenStateMulti(),
			appGS,
		)
	})

}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
