package keeper_test

import (
	"github.com/UnUniFi/chain/x/nftmint/keeper"
	"github.com/UnUniFi/chain/x/nftmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

const (
	testNFTId = "a00"
)

// test for the MintNFT relating functions
func (suite *KeeperTestSuite) TestMintNFT() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)

	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)

	recipient := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testMsgMintNFT := types.MsgMintNFT{
		Sender:    sender.Bytes(),
		ClassId:   classId,
		NftId:     testNFTId,
		Recipient: recipient.Bytes(),
	}
	err := suite.app.NftmintKeeper.MintNFT(suite.ctx, &testMsgMintNFT)
	suite.Require().NoError(err)
	// check owner
	owner := suite.app.NFTKeeper.GetOwner(suite.ctx, classId, testNFTId)
	expectedOwner := recipient
	suite.Require().Equal(expectedOwner, owner)
	// check minter
	minter, exists := suite.app.NftmintKeeper.GetNFTMinter(suite.ctx, classId, testNFTId)
	suite.Require().True(exists)
	expectedMinter := sender
	suite.Require().Equal(expectedMinter, minter)

	// invalid nft id case to give invalid nft uri length on UnUniFi
	invalidNFTIdUri := "test"
	for i := 0; i < types.DefaultMaxUriLen; i++ {
		invalidNFTIdUri += "a"
	}
	testMsgMintNFTInvalidNftIdUri := types.MsgMintNFT{
		Sender:    sender.Bytes(),
		ClassId:   classId,
		NftId:     invalidNFTIdUri,
		Recipient: recipient.Bytes(),
	}
	err = suite.app.NftmintKeeper.MintNFT(suite.ctx, &testMsgMintNFTInvalidNftIdUri)
	suite.Require().Error(err)
	exists = suite.app.NFTKeeper.HasNFT(suite.ctx, classId, invalidNFTIdUri)
	suite.Require().False(exists)

	// invalid nft id by sdk's x/nft validation in case for not being called through message
	invalidNFTId := "a" // shorter than the defined minimum length by sdk's x/nft module
	testMsgMintNFTInvalidNftId := types.MsgMintNFT{
		Sender:    sender.Bytes(),
		ClassId:   classId,
		NftId:     invalidNFTId,
		Recipient: recipient.Bytes(),
	}
	err = suite.app.NftmintKeeper.MintNFT(suite.ctx, &testMsgMintNFTInvalidNftId)
	suite.Require().Error(err)
	exists = suite.app.NFTKeeper.HasNFT(suite.ctx, classId, invalidNFTId)
	suite.Require().False(exists)

	// invalid case by the minting permission limitation
	// estimate MintingPermission = 0 (OnlyOwner)
	notOwnerOfClass := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	testNFTId2 := testNFTId + "2"
	testMsgMintNFTInvalidMinter := types.MsgMintNFT{
		Sender:    notOwnerOfClass.Bytes(),
		ClassId:   classId,
		NftId:     testNFTId2,
		Recipient: recipient.Bytes(),
	}
	err = suite.app.NftmintKeeper.MintNFT(suite.ctx, &testMsgMintNFTInvalidMinter)
	suite.Require().Error(err)
	exists = suite.app.NFTKeeper.HasNFT(suite.ctx, classId, testNFTId2)
	suite.Require().False(exists)

	// invalid case which is over the defined token supply cap
	classId = "test"
	testMsgCreateClass := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              "test",
		BaseTokenUri:      "ipfs://testuri/",
		TokenSupplyCap:    1,
		MintingPermission: 0,
	}
	_ = suite.app.NftmintKeeper.CreateClass(suite.ctx, classId, &testMsgCreateClass)
	_ = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{ClassId: classId, Id: testNFTId}, sender)
	testMsgMintNFTOverCap := types.MsgMintNFT{
		Sender:    sender.Bytes(),
		ClassId:   classId,
		NftId:     testNFTId2,
		Recipient: sender.Bytes(),
	}
	err = suite.app.NftmintKeeper.MintNFT(suite.ctx, &testMsgMintNFTOverCap)
	suite.Require().Error(err)
	exists = suite.app.NFTKeeper.HasNFT(suite.ctx, classId, testNFTId2)
	suite.Require().False(exists)
}

// tests for the BurnNFT relating functions
func (suite *KeeperTestSuite) TestBurnNFT() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)

	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.MintNFT(suite.ctx, classId, testNFTId, sender)

	testMsgBurnNFT := types.MsgBurnNFT{
		Sender:  sender.Bytes(),
		ClassId: classId,
		NftId:   testNFTId,
	}
	err := suite.app.NftmintKeeper.BurnNFT(suite.ctx, &testMsgBurnNFT)
	suite.Require().NoError(err)
	// check if burned successfully
	exists := suite.app.NFTKeeper.HasNFT(suite.ctx, classId, testNFTId)
	suite.Require().False(exists)

	// invalid case which sender is not the owner of the nft
	invalidSender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	_ = suite.MintNFT(suite.ctx, classId, testNFTId, sender)
	testMsgBurnNFTInvalidSender := types.MsgBurnNFT{
		Sender:  invalidSender.Bytes(),
		ClassId: classId,
		NftId:   testNFTId,
	}
	err = suite.app.NftmintKeeper.BurnNFT(suite.ctx, &testMsgBurnNFTInvalidSender)
	suite.Require().Error(err)
	// check if not burned as intended
	exists = suite.app.NFTKeeper.HasNFT(suite.ctx, classId, testNFTId)
	suite.Require().True(exists)
}

// tests for UpdateNFTUri function
func (suite *KeeperTestSuite) TestUpdateNFTUri() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)

	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.MintNFT(suite.ctx, classId, testNFTId, sender)

	updatingBaseTokenUri := "ipfs://test-latest/"
	err := suite.app.NftmintKeeper.UpdateNFTUri(suite.ctx, classId, updatingBaseTokenUri)
	suite.Require().NoError(err)
	nft, _ := suite.app.NFTKeeper.GetNFT(suite.ctx, classId, testNFTId)
	expectedNFTUri := updatingBaseTokenUri + testNFTId
	suite.Require().Equal(expectedNFTUri, nft.Uri)

	// invalid BaseTokenUri length defined on UnUniFi
	invalidBaseTokenUri := "invalid"
	for i := 0; i < types.DefaultMaxUriLen; i++ {
		invalidBaseTokenUri += "a"
	}
	err = suite.app.NftmintKeeper.UpdateNFTUri(suite.ctx, classId, invalidBaseTokenUri)
	suite.Require().Error(err)
	// check if nft uri doesn't change after updating
	nft, _ = suite.app.NFTKeeper.GetNFT(suite.ctx, classId, testNFTId)
	expectedNFTUri = updatingBaseTokenUri + testNFTId
	suite.Require().Equal(expectedNFTUri, nft.Uri)
}

// mint nft method for keeper test
func (suite *KeeperTestSuite) MintNFT(ctx sdk.Context, classID, nftID string, sender sdk.AccAddress) error {
	_ = suite.CreateClass(suite.ctx, classID, sender)

	testMsgMintNFT := types.MsgMintNFT{
		Sender:    sender.Bytes(),
		ClassId:   classID,
		NftId:     nftID,
		Recipient: sender.Bytes(),
	}
	err := suite.app.NftmintKeeper.MintNFT(suite.ctx, &testMsgMintNFT)
	return err
}
