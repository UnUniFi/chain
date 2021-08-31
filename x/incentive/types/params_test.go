package types_test

import (
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/suite"

	incentivetypes "github.com/lcnem/jpyx/x/incentive/types"
)

type ParamTestSuite struct {
	suite.Suite
}

func (suite *ParamTestSuite) SetupTest() {}

func (suite *ParamTestSuite) TestParamValidation() {
	type args struct {
		jpyxMintingRewardPeriods   incentivetypes.RewardPeriods
		hardSupplyRewardPeriods    incentivetypes.MultiRewardPeriods
		hardBorrowRewardPeriods    incentivetypes.MultiRewardPeriods
		hardDelegatorRewardPeriods incentivetypes.RewardPeriods
		multipliers                incentivetypes.Multipliers
		end                        time.Time
	}

	type errArgs struct {
		expectPass bool
		contains   string
	}
	type test struct {
		name    string
		args    args
		errArgs errArgs
	}

	testCases := []test{
		{
			"default",
			args{
				jpyxMintingRewardPeriods:   incentivetypes.DefaultRewardPeriods,
				hardSupplyRewardPeriods:    incentivetypes.DefaultMultiRewardPeriods,
				hardBorrowRewardPeriods:    incentivetypes.DefaultMultiRewardPeriods,
				hardDelegatorRewardPeriods: incentivetypes.DefaultRewardPeriods,
				multipliers:                incentivetypes.DefaultMultipliers,
				end:                        incentivetypes.DefaultClaimEnd,
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"valid",
			args{
				jpyxMintingRewardPeriods: incentivetypes.RewardPeriods{incentivetypes.NewRewardPeriod(
					true, "bnb-a", time.Date(2020, 10, 15, 14, 0, 0, 0, time.UTC), time.Date(2024, 10, 15, 14, 0, 0, 0, time.UTC),
					sdk.NewCoin(incentivetypes.JpyxMintingRewardDenom, sdk.NewInt(122354)))},
				multipliers: incentivetypes.Multipliers{
					incentivetypes.NewMultiplier(
						incentivetypes.Small, 1, sdk.MustNewDecFromStr("0.25"),
					),
					incentivetypes.NewMultiplier(
						incentivetypes.Large, 1, sdk.MustNewDecFromStr("1.0"),
					),
				},
				hardSupplyRewardPeriods:    incentivetypes.DefaultMultiRewardPeriods,
				hardBorrowRewardPeriods:    incentivetypes.DefaultMultiRewardPeriods,
				hardDelegatorRewardPeriods: incentivetypes.DefaultRewardPeriods,
				end:                        time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			params := incentivetypes.NewParams(tc.args.jpyxMintingRewardPeriods, tc.args.hardSupplyRewardPeriods,
				tc.args.hardBorrowRewardPeriods, tc.args.hardDelegatorRewardPeriods, tc.args.multipliers, tc.args.end,
			)
			err := params.Validate()
			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func TestParamTestSuite(t *testing.T) {
	suite.Run(t, new(ParamTestSuite))
}
