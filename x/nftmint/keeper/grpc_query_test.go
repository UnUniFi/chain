package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmint/keeper"
	"github.com/UnUniFi/chain/x/nftmint/types"
)

func (suite *KeeperTestSuite) TestQueryClassAttributes() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)

	req := types.QueryClassAttributesRequest{
		ClassId: classId,
	}
	res, err := suite.nftmintKeeper.ClassAttributes(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().Equal(classId, res.ClassAttributes.ClassId)
	suite.Require().Equal(testBaseTokenUri, res.ClassAttributes.BaseTokenUri)

	invalidReq := types.QueryClassAttributesRequest{
		ClassId: "invalidClassId",
	}
	_, err = suite.nftmintKeeper.ClassAttributes(suite.ctx, &invalidReq)
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestQueryNftMinter() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)
	_ = suite.MintNFT(suite.ctx, classId, testNFTId, sender)

	req := types.QueryNFTMinterRequest{
		ClassId: classId,
		NftId:   testNFTId,
	}
	res, err := suite.nftmintKeeper.NFTMinter(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().Equal(sender.String(), res.Minter)
}

func (suite *KeeperTestSuite) TestQueryClassIdsByOwner() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)
	_ = suite.MintNFT(suite.ctx, classId, testNFTId, sender)

	req := types.QueryClassIdsByOwnerRequest{
		Owner: sender.String(),
	}
	res, err := suite.nftmintKeeper.ClassIdsByOwner(suite.ctx, &req)
	suite.Require().NoError(err)
	var classIds []string
	classIds = append(classIds, classId)
	expectedRes := types.OwningClassIdList{
		Owner:   sender.Bytes(),
		ClassId: classIds[:],
	}
	suite.Require().Equal(&expectedRes, res.OwningClassIdList)
}

func (suite *KeeperTestSuite) TestQueryIdsByName() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)
	_ = suite.MintNFT(suite.ctx, classId, testNFTId, sender)

	req := types.QueryClassIdsByNameRequest{
		ClassName: testName,
	}
	res, err := suite.nftmintKeeper.ClassIdsByName(suite.ctx, &req)
	suite.Require().NoError(err)
	var classIds []string
	classIds = append(classIds, classId)
	expectedRes := types.ClassNameIdList{
		ClassName: testName,
		ClassId:   classIds,
	}
	suite.Require().Equal(&expectedRes, res.ClassNameIdList)
}
