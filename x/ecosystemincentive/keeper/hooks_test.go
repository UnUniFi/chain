package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
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
			txMemo:         `{"version":"v1","incentive_unit_id":"incentiveUnitId1"}`,
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
			txMemo:         `{"version":"v1","incentive_unit_id":"incentiveUnitId2"}`,
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
			txMemo:         `{"version":"v1","incentive_unit_id":"incentiveUnitId2"}`,
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
			txMemo:         `{"error":true,"version":"v1","incentive_unit_id":"incentiveUnitId3"}`,
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
			txMemo:         `{"version":"v0","incentive_unit_id":"incentiveUnitId4"}`,
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
		// calculate the reward for incentiveUnit using the default rate
		expRewardForIncentiveUnit sdk.Int
		// calculate the reward for stakers using the default rate
		expRewardForStakers sdk.Coin
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
			weights:                   []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore:            true,
			expectPass:                true,
			expRewardForIncentiveUnit: sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt(),
			expRewardForStakers:       sdk.NewCoin("uguu", sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt()),
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
			weights:                   []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore:            true,
			expectPass:                true,
			expRewardForIncentiveUnit: sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt(),
			expRewardForStakers:       sdk.NewCoin("uguu", sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt()),
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

		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.reward})
		_ = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{tc.reward})

		suite.app.EcosystemincentiveKeeper.Hooks().AfterNftPaymentWithCommission(suite.ctx, tc.nftId, tc.reward)

		if tc.expectPass {
			totalRewardForIncentiveUnit := sdk.ZeroInt()
			for _, subject := range tc.subjectAddrs {
				rewardStore, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, sdk.AccAddress(subject))
				totalRewardForIncentiveUnit = totalRewardForIncentiveUnit.Add(rewardStore.Rewards.AmountOf(tc.reward.Denom))
				suite.Require().True(exists)
			}
			suite.Require().Equal(tc.expRewardForIncentiveUnit, totalRewardForIncentiveUnit)

			// check the reward distribution for stakers by checking the balance of the fee_collector module account
			feeCollectorAcc := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
			feeCollectorAccBalance := suite.app.BankKeeper.GetBalance(suite.ctx, feeCollectorAcc, tc.reward.Denom)
			suite.Require().Equal(tc.expRewardForStakers, feeCollectorAccBalance)
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

func (suite *KeeperTestSuite) TestAfterNftUnlistedWithoutPayment() {
	tests := []struct {
		testCase        string
		nftId           nftmarkettypes.NftIdentifier
		incentiveUnitId string
		subjectAddrs    []ununifitypes.StringAccAddress
		weights         []sdk.Dec
		registerBefore  bool
		expectPass      bool
	}{
		{
			testCase: "ordinal case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class1",
				NftId:   "nft1",
			},
			incentiveUnitId: "incentiveUnitId1",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			registerBefore: true,
		},
		{
			testCase: "not recorded case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class2",
				NftId:   "nft2",
			},
			incentiveUnitId: "incentiveUnitId2",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			registerBefore: true,
		},
	}

	for _, tc := range tests {
		_, err := suite.app.EcosystemincentiveKeeper.Register(suite.ctx, &types.MsgRegister{
			Sender:          tc.subjectAddrs[0],
			IncentiveUnitId: tc.incentiveUnitId,
			SubjectAddrs:    tc.subjectAddrs,
			Weights:         tc.weights,
		})
		suite.Require().NoError(err)
		if tc.registerBefore {
			suite.app.EcosystemincentiveKeeper.RecordIncentiveUnitIdWithNftId(suite.ctx, tc.nftId, tc.incentiveUnitId)
		}

		suite.NotPanics(func() {
			suite.app.EcosystemincentiveKeeper.Hooks().AfterNftUnlistedWithoutPayment(suite.ctx, tc.nftId)
		})

		_, exists := suite.app.EcosystemincentiveKeeper.GetIncentiveUnitIdByNftId(suite.ctx, tc.nftId)
		suite.Require().False(exists)
	}
}
