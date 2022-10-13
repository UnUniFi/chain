package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/tendermint/tendermint/crypto/ed25519"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

func (suite *KeeperTestSuite) TestRecordNftIdWithIncentiveUnitId() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase        string
		classId         string
		nftId           string
		incentiveUnitId string
		subjectAddrs    []ununifitypes.StringAccAddress
		weights         []sdk.Dec
		registerBefore  bool
		expectPass      bool
	}{
		{
			testCase:        "not registered",
			classId:         "class1",
			nftId:           "nft1",
			incentiveUnitId: "id1",
			subjectAddrs:    []ununifitypes.StringAccAddress{sender.Bytes()},
			weights:         []sdk.Dec{sdk.MustNewDecFromStr("1.0")},
			registerBefore:  false,
			expectPass:      false,
		},
		{
			testCase:        "registed",
			classId:         "class2",
			nftId:           "nft2",
			incentiveUnitId: "id2",
			subjectAddrs:    []ununifitypes.StringAccAddress{sender.Bytes()},
			weights:         []sdk.Dec{sdk.MustNewDecFromStr("1.0")},
			registerBefore:  true,
			expectPass:      true,
		},
		{
			testCase:        "already recorded",
			classId:         "class2",
			nftId:           "nft2",
			incentiveUnitId: "id3",
			subjectAddrs:    []ununifitypes.StringAccAddress{sender.Bytes()},
			weights:         []sdk.Dec{sdk.MustNewDecFromStr("1.0")},
			registerBefore:  true,
			expectPass:      false,
		},
	}
	for _, test := range tests {
		nftId := nftmarkettypes.NftIdentifier{
			ClassId: test.classId,
			NftId:   test.nftId,
		}

		if test.registerBefore {
			_, _ = suite.app.EcosystemincentiveKeeper.Register(
				suite.ctx,
				&types.MsgRegister{
					Sender:          sender.Bytes(),
					IncentiveUnitId: test.incentiveUnitId,
					SubjectAddrs:    test.subjectAddrs,
					Weights:         test.weights,
				},
			)

			err := suite.RecordNftIdWithIncentiveUnitIdTest(suite.ctx, nftId, test.incentiveUnitId)

			if test.expectPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		} else {
			err := suite.RecordNftIdWithIncentiveUnitIdTest(suite.ctx, nftId, test.incentiveUnitId)
			suite.Require().Error(err)
		}

	}
}

// RecordNftIdWithIncentiveUnitIdTest is a mehtod to have the exact same logic
// for being used in test cases to return error as return value
// since the normal RecordNftIdWithIncentiveUnitId doesn't return any value by intention
func (suite *KeeperTestSuite) RecordNftIdWithIncentiveUnitIdTest(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, incentiveUnitId string) error {
	// panic if the nftId is already recorded in the store.
	if _, exists := suite.app.EcosystemincentiveKeeper.GetNftIdForFrontend(ctx, nftId); exists {
		return types.ErrRecordedNftId
	}

	// check incentiveUnitId is already registered
	if _, exists := suite.app.EcosystemincentiveKeeper.GetIncentiveUnit(ctx, incentiveUnitId); !exists {
		return types.ErrNotRegisteredIncentiveUnitId
	}

	if err := suite.app.EcosystemincentiveKeeper.SetNftIdForFrontend(ctx, nftId, incentiveUnitId); err != nil {
		return err
	}

	return nil
}
