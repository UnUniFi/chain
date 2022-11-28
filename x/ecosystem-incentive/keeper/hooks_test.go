package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

func (suite *KeeperTestSuite) TestAfterNftListed() {
	tests := []struct {
		testCase        string
		nftId           nftmarkettypes.NftIdentifier
		incentiveUnitId string
		subjectAddrs    []ununifitypes.StringAccAddress
		weights         []sdk.Dec
		txMemo          string
		registerBefore  bool
		expectPass      bool
		expectPanic     bool
		memoFormat      bool
	}{
		{
			testCase: "ordinal success case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class1",
				NftId:   "nft1",
			},
			incentiveUnitId: "incentiveUnitId1",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			txMemo:         `{"version":"v1","incentive-unit-id":"incentiveUnitId1"}`,
			registerBefore: true,
			expectPass:     true,
			expectPanic:    false,
			memoFormat:     true,
		},
		{
			testCase: "incentive unit id is not registered",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class2",
				NftId:   "nft2",
			},
			txMemo:         `{"version":"v1","incentive-unit-id":"incentiveUnitId2"}`,
			registerBefore: false,
			expectPass:     false,
			expectPanic:    false,
			memoFormat:     true,
		},
		{
			testCase: "panic since incentive unit id is already recorded with nft id",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class1",
				NftId:   "nft1",
			},
			txMemo:         `{"version":"v1","incentive-unit-id":"incentiveUnitId2"}`,
			registerBefore: false,
			expectPass:     false,
			expectPanic:    true,
			memoFormat:     true,
		},
		{
			testCase: "invalid memo format",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class3",
				NftId:   "nft3",
			},
			incentiveUnitId: "incentiveUnitId3",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			txMemo:         `{"error":true,"version":"v1","incentive-unit-id":"incentiveUnitId3"}`,
			registerBefore: true,
			expectPass:     false,
			expectPanic:    false,
			memoFormat:     false,
		},
		{
			testCase: "invalid memo format version",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class4",
				NftId:   "nft4",
			},
			incentiveUnitId: "incentiveUnitId4",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			txMemo:         `{"version":"v0","incentive-unit-id":"incentiveUnitId4"}`,
			registerBefore: true,
			expectPass:     false,
			expectPanic:    false,
			memoFormat:     false,
		},
	}

	for _, tc := range tests {
		if tc.registerBefore {
			_, err := suite.app.EcosystemincentiveKeeper.Register(suite.ctx, &types.MsgRegister{
				Sender:          tc.subjectAddrs[0],
				IncentiveUnitId: tc.incentiveUnitId,
				SubjectAddrs:    tc.subjectAddrs,
				Weights:         tc.weights,
			})
			suite.Require().NoError(err)
		}

		if tc.expectPanic {
			suite.Panics(func() {
				suite.app.EcosystemincentiveKeeper.Hooks().AfterNftListed(suite.ctx, tc.nftId, tc.txMemo)
			})
		} else {
			suite.NotPanics(func() {
				suite.app.EcosystemincentiveKeeper.Hooks().AfterNftListed(suite.ctx, tc.nftId, tc.txMemo)
			})
		}

		incentiveUnitId, exists := suite.app.EcosystemincentiveKeeper.GetIncentiveUnitIdByNftId(suite.ctx, tc.nftId)
		if tc.expectPass {
			suite.Require().True(exists)
			suite.Require().Equal(tc.incentiveUnitId, incentiveUnitId)
		} else if tc.expectPanic {
			suite.Require().True(exists)
			suite.Require().NotEqual(tc.incentiveUnitId, incentiveUnitId)
		} else {
			suite.Require().False(exists)
		}
	}
}

// TODO: test
// func (suite *KeeperTestSuite) TestAfterNftPaymentWithCommission()
// TODO: test
// func (suite *KeeperTestSuite) TestAfterNftUnlistedWithoutPayment()
