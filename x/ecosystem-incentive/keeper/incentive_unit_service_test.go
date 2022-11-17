package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

func (suite *KeeperTestSuite) TestRegister() {

	tests := []struct {
		testCase        string
		incentiveUnitId string
		subjectAddrs    []ununifitypes.StringAccAddress
		weights         []sdk.Dec
		success         bool
	}{
		{
			testCase:        "ordinal success case",
			incentiveUnitId: "test1",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights: []sdk.Dec{sdk.MustNewDecFromStr("1")},
			success: true,
		},
		{
			testCase:        "multiple subjects success case",
			incentiveUnitId: "test2",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
				ununifitypes.StringAccAddress(suite.addrs[1]),
				ununifitypes.StringAccAddress(suite.addrs[2]),
			},
			weights: []sdk.Dec{
				sdk.MustNewDecFromStr("0.33"),
				sdk.MustNewDecFromStr("0.33"),
				sdk.MustNewDecFromStr("0.34"),
			},
			success: true,
		},
		{
			testCase:        "failure due to the duplicated inentiveUnitId",
			incentiveUnitId: "test1",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights: []sdk.Dec{sdk.MustNewDecFromStr("1")},
			success: false,
		},
		{
			testCase:        "failure due to invalid incentiveUnitId",
			incentiveUnitId: "",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights: []sdk.Dec{
				sdk.MustNewDecFromStr("1"),
			},
			success: false,
		},
	}

	for _, tc := range tests {
		subjectInfo, err := suite.app.EcosystemincentiveKeeper.Register(suite.ctx, &types.MsgRegister{
			Sender:          ununifitypes.StringAccAddress(suite.addrs[0]),
			IncentiveUnitId: tc.incentiveUnitId,
			SubjectAddrs:    tc.subjectAddrs,
			Weights:         tc.weights,
		})
		if tc.success {
			suite.Require().NoError(err)
			for i, subject := range *subjectInfo {
				suite.Require().Equal(subject.Address, tc.subjectAddrs[i])
				suite.Require().Equal(subject.Weight, tc.weights[i])
			}
		} else {
			suite.Require().Error(err)
		}
	}
}
