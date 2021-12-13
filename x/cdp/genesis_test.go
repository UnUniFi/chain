package cdp_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/app"
	cdpkeeper "github.com/UnUniFi/chain/x/cdp/keeper"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx    sdk.Context
	keeper cdpkeeper.Keeper
}

func (suite *GenesisTestSuite) TestInvalidGenState() {
	type args struct {
		params             cdptypes.Params
		cdps               cdptypes.Cdps
		deposits           cdptypes.Deposits
		startingID         uint64
		debtDenom          string
		govDenom           string
		genAccumTimes      cdptypes.GenesisAccumulationTimes
		genTotalPrincipals cdptypes.GenesisTotalPrincipals
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
				params:             cdptypes.DefaultParams(),
				cdps:               cdptypes.Cdps{},
				deposits:           cdptypes.Deposits{},
				debtDenom:          "",
				govDenom:           cdptypes.DefaultGovDenom,
				genAccumTimes:      cdptypes.DefaultGenesis().PreviousAccumulationTimes,
				genTotalPrincipals: cdptypes.DefaultGenesis().TotalPrincipals,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "debt denom invalid",
			},
		},
		{
			name: "empty gov denom",
			args: args{
				params:             cdptypes.DefaultParams(),
				cdps:               cdptypes.Cdps{},
				deposits:           cdptypes.Deposits{},
				debtDenom:          cdptypes.DefaultDebtDenom,
				govDenom:           "",
				genAccumTimes:      cdptypes.DefaultGenesis().PreviousAccumulationTimes,
				genTotalPrincipals: cdptypes.DefaultGenesis().TotalPrincipals,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "gov denom invalid",
			},
		},
		{
			name: "interest factor below one",
			args: args{
				params:             cdptypes.DefaultParams(),
				cdps:               cdptypes.Cdps{},
				deposits:           cdptypes.Deposits{},
				debtDenom:          cdptypes.DefaultDebtDenom,
				govDenom:           cdptypes.DefaultGovDenom,
				genAccumTimes:      cdptypes.GenesisAccumulationTimes{cdptypes.NewGenesisAccumulationTime("bnb-a", time.Time{}, sdk.OneDec().Sub(sdk.SmallestDec()))},
				genTotalPrincipals: cdptypes.DefaultGenesis().TotalPrincipals,
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "interest factor should be ≥ 1.0",
			},
		},
		{
			name: "negative total principal",
			args: args{
				params:             cdptypes.DefaultParams(),
				cdps:               cdptypes.Cdps{},
				deposits:           cdptypes.Deposits{},
				debtDenom:          cdptypes.DefaultDebtDenom,
				govDenom:           cdptypes.DefaultGovDenom,
				genAccumTimes:      cdptypes.DefaultGenesis().PreviousAccumulationTimes,
				genTotalPrincipals: cdptypes.GenesisTotalPrincipals{cdptypes.NewGenesisTotalPrincipal("bnb-a", sdk.NewInt(-1))},
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "total principal should be positive",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			gs := cdptypes.NewGenesisState(tc.args.params, tc.args.cdps, tc.args.deposits, tc.args.startingID,
				tc.args.debtDenom, tc.args.govDenom, tc.args.genAccumTimes, tc.args.genTotalPrincipals)
			err := gs.Validate()
			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
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
	tApp.AppCodec().MustUnmarshalJSON(cdpGS[cdptypes.ModuleName], &gs)
	gs.Cdps = cdps()
	gs.StartingCdpId = uint64(5)
	appGS := app.GenesisState{cdptypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&gs)}

	tApp = app.NewTestApp()
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
