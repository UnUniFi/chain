package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

func (suite *KeeperTestSuite) TestRegister() {

	tests := []struct {
		testCase             string
		recipientContainerId string
		subjectAddrs         []string
		weights              []sdk.Dec
		success              bool
	}{
		{
			testCase:             "ordinal success case",
			recipientContainerId: "test1",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights: []sdk.Dec{sdk.MustNewDecFromStr("1")},
			success: true,
		},
		{
			testCase:             "multiple subjects success case",
			recipientContainerId: "test2",
			subjectAddrs: []string{
				suite.addrs[0].String(),
				suite.addrs[1].String(),
				suite.addrs[2].String(),
			},
			weights: []sdk.Dec{
				sdk.MustNewDecFromStr("0.33"),
				sdk.MustNewDecFromStr("0.33"),
				sdk.MustNewDecFromStr("0.34"),
			},
			success: true,
		},
		{
			testCase:             "failure due to the duplicated inentiveUnitId",
			recipientContainerId: "test1",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights: []sdk.Dec{sdk.MustNewDecFromStr("1")},
			success: false,
		},
		{
			testCase:             "failure due to invalid recipientContainerId",
			recipientContainerId: "",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights: []sdk.Dec{
				sdk.MustNewDecFromStr("1"),
			},
			success: false,
		},
	}

	for _, tc := range tests {
		subjectInfo, err := suite.app.EcosystemincentiveKeeper.Register(suite.ctx, &types.MsgRegister{
			Sender:               suite.addrs[0].String(),
			RecipientContainerId: tc.recipientContainerId,
			Addresses:            tc.subjectAddrs,
			Weights:              tc.weights,
		})
		if tc.success {
			suite.Require().NoError(err)
			for i, subject := range *subjectInfo {
				suite.Require().Equal(subject.SubjectAddr, tc.subjectAddrs[i])
				suite.Require().Equal(subject.Weight, tc.weights[i])

				recipientContainerIdsByAddr := suite.app.EcosystemincentiveKeeper.GetRecipientContainerIdsByAddr(suite.ctx, subject.SubjectAddr.AccAddress())
				suite.Require().Contains(recipientContainerIdsByAddr.RecipientContainerIds, tc.recipientContainerId)
			}
		} else {
			suite.Require().Error(err)
		}
	}
}
