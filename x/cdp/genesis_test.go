package cdp_test

import (
	"fmt"
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
		suite.Run(tc.name, func() {
			params := tc.args.params // suite.keeper.GetParams(suite.ctx)

			var prevDistTimes cdptypes.GenesisAccumulationTimes
			var totalPrincipals cdptypes.GenesisTotalPrincipals

			for _, cp := range params.CollateralParams {
				interestFactor, found := suite.keeper.GetInterestFactor(suite.ctx, cp.Type)
				if !found {
					interestFactor = sdk.OneDec()
				}
				// Governance param changes happen in the end blocker. If a new collateral type is added and then the chain
				// is exported before the BeginBlocker can run, previous accrual time won't be found. We can't set it to
				// current block time because it is not available in the export ctx. We should panic instead of exporting
				// bad state.
				previDisTime, f := suite.keeper.GetPreviousAccrualTime(suite.ctx, cp.Type)
				if !f {
					panic(fmt.Sprintf("expected previous accrual time to be set in state for %s", cp.Type))
				}
				prevDistTimes = append(prevDistTimes, cdptypes.NewGenesisAccumulationTime(cp.Type, previDisTime, interestFactor))
				tp := suite.keeper.GetTotalPrincipal(suite.ctx, cp.Type, cdptypes.DefaultStableDenom)
				genTotalPrincipal := cdptypes.NewGenesisTotalPrincipal(cp.Type, tp)
				totalPrincipals = append(totalPrincipals, genTotalPrincipal)
			}
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
			NewPricefeedGenStateMulti(tApp),
			NewCDPGenStateMulti(tApp),
		)
	})

	cdpGS := NewCDPGenStateMulti(tApp)
	gs := cdptypes.GenesisState{}
	cdptypes.ModuleCdc.UnmarshalJSON(cdpGS["cdp"], &gs)
	gs.Cdps = cdps()
	gs.StartingCdpId = uint64(5)
	appGS := app.GenesisState{"cdp": cdptypes.ModuleCdc.MustMarshalJSON(&gs)}
	suite.NotPanics(func() {
		tApp.InitializeFromGenesisStates(
			NewPricefeedGenStateMulti(tApp),
			appGS,
		)
	})

}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
