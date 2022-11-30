package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystem-incentive/keeper"
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

func (suite *KeeperTestSuite) TestAfterNftPaymentWithCommission() {
	tests := []struct {
		testCase        string
		nftId           nftmarkettypes.NftIdentifier
		reward          sdk.Coin
		incentiveUnitId string
		subjectAddrs    []ununifitypes.StringAccAddress
		weights         []sdk.Dec
		recordedBefore  bool
		expectPass      bool
	}{
		{
			testCase: "failure case since incentive unit id was not recorded with nft id",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class1",
				NftId:   "nft1",
			},
			reward:          sdk.Coin{"uguu", sdk.NewInt(100)},
			incentiveUnitId: "incentiveUnitId1",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore: false,
			expectPass:     false,
		},
		{
			testCase: "failure case since there's no fee to distribute",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class2",
				NftId:   "nft2",
			},
			reward:          sdk.Coin{"uguu", sdk.NewInt(0)},
			incentiveUnitId: "incentiveUnitId1",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore: true,
			expectPass:     false,
		},
		{
			testCase: "ordinal case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class3",
				NftId:   "nft3",
			},
			reward:          sdk.Coin{"uguu", sdk.NewInt(100)},
			incentiveUnitId: "incentiveUnitId3",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore: true,
			expectPass:     true,
		},
		{
			testCase: "ordinal case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class4",
				NftId:   "nft4",
			},
			reward:          sdk.Coin{"uguu", sdk.NewInt(100)},
			incentiveUnitId: "incentiveUnitId4",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore: true,
			expectPass:     true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()
		_, err := suite.app.EcosystemincentiveKeeper.Register(suite.ctx, &types.MsgRegister{
			Sender:          tc.subjectAddrs[0],
			IncentiveUnitId: tc.incentiveUnitId,
			SubjectAddrs:    tc.subjectAddrs,
			Weights:         tc.weights,
		})
		suite.Require().NoError(err)
		if tc.recordedBefore {
			suite.app.EcosystemincentiveKeeper.RecordIncentiveUnitIdWithNftId(suite.ctx, tc.nftId, tc.incentiveUnitId)
		}

		suite.app.EcosystemincentiveKeeper.Hooks().AfterNftPaymentWithCommission(suite.ctx, tc.nftId, tc.reward)

		if tc.expectPass {
			_, rewardsForEach := keeper.CalculateRewardsForEachSubject(tc.weights, tc.reward, suite.app.EcosystemincentiveKeeper.GetNftmarketFrontendRewardRate(suite.ctx))
			for i, subject := range tc.subjectAddrs {
				rewardStore, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, sdk.AccAddress(subject))
				suite.Require().True(exists)
				suite.Require().Equal(sdk.NewCoins(sdk.NewCoins(rewardsForEach[i])...), rewardStore.Rewards)
			}
		} else {
			for _, subject := range tc.subjectAddrs {
				_, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, sdk.AccAddress(subject))
				suite.Require().False(exists)
			}
		}

		_, exists := suite.app.EcosystemincentiveKeeper.GetIncentiveUnitIdByNftId(suite.ctx, tc.nftId)
		suite.Require().False(exists)
	}
}

// TODO: test
// func (suite *KeeperTestSuite) TestAfterNftUnlistedWithoutPayment()
