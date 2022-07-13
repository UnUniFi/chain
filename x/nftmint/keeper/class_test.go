package keeper_test

import (
	"fmt"
	"testing"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/nftmint/keeper"
	"github.com/UnUniFi/chain/x/nftmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

// test basic functions of nftmint
func (suite *KeeperTestSuite) TestNftMintBasics() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, owner)

	classId := keeper.CreateClassId(owner_seq, owner)

	// test relating ClassAttributes
	classAttributes := types.ClassAttributes{
		ClassId: classId,
	}
	// check setting ClassAttributes function
	err := suite.app.NftmintKeeper.SetClassAttributes(suite.ctx, classAttributes)
	suite.Require().NoError(err)
	// check getting ClassAttributes function
	gotClassAttributes, exists := suite.app.NftmintKeeper.GetClassAttributes(suite.ctx, classAttributes.ClassId)
	suite.Require().True(exists)
	suite.Require().Equal(classAttributes, gotClassAttributes)

	// test relating OwningClassIdList
	var classIdList []string
	classIdList = append(classIdList, classId)
	owningClassIdList := types.OwningClassIdList{
		Owner:   ununifitypes.StringAccAddress(owner),
		ClassId: classIdList,
	}
	// check setting OwningClassIdList function
	err = suite.app.NftmintKeeper.SetOwningClassIdList(suite.ctx, owningClassIdList)
	suite.Require().NoError(err)
	// check getting OwningClassIdList function
	gotOwningClassIdList, exists := suite.app.NftmintKeeper.GetOwningClassIdList(suite.ctx, owner)
	suite.Require().True(exists)
	suite.Require().Equal(owningClassIdList, gotOwningClassIdList)

	// test relating ClassNameIdList
	testClassName := "test"
	classNameIdList := types.ClassNameIdList{
		ClassName: testClassName,
		ClassId:   classIdList,
	}
	// check setting ClassNameIdList function
	err = suite.app.NftmintKeeper.SetClassNameIdList(suite.ctx, classNameIdList)
	suite.Require().NoError(err)
	// check getting ClassNameIdList function
	gotClassNameIdList, exists := suite.app.NftmintKeeper.GetClassNameIdList(suite.ctx, testClassName)
	suite.Require().True(exists)
	suite.Require().Equal(classNameIdList, gotClassNameIdList)
}

// test of the method to create a class.Id
func TestCreateId(t *testing.T) {
	var seq uint64 = 0
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	fmt.Println(addr.String())
	accAddr, _ := sdk.AccAddressFromBech32(addr.String())
	classIdSeq0 := keeper.CreateClassId(seq, accAddr)
	err := nfttypes.ValidateClassID(classIdSeq0)
	require.NoError(t, err)

	// add one to imitate actual account sequence transition of the expected situation
	seq += 1
	classIdSeq1 := keeper.CreateClassId(seq, accAddr)
	require.NoError(t, err)
	require.NotEqual(t, classIdSeq0, classIdSeq1)
}

func (suite *KeeperTestSuite) TestCreateClass() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)

	testMsgCreateClass := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              "test",
		BaseTokenUri:      "ipfs://testcid-sample/",
		TokenSupplyCap:    10000,
		MintingPermission: 0,
	}
	err := suite.app.NftmintKeeper.CreateClass(suite.ctx, classId, &testMsgCreateClass)
	suite.Require().NoError(err)

	// check if Class is set
	class, exists := suite.app.NFTKeeper.GetClass(suite.ctx, classId)
	suite.Require().True(exists)
	expectedClass := nfttypes.Class{
		Id:   classId,
		Name: testMsgCreateClass.Name,
	}
	suite.Require().Equal(class, expectedClass)

	// check if ClassAttributes is set
	classAttributes, exists := suite.app.NftmintKeeper.GetClassAttributes(suite.ctx, classId)
	suite.Require().True(exists)
	expectedClassAttributes := types.ClassAttributes{
		ClassId:           classId,
		Owner:             ununifitypes.StringAccAddress(sender),
		BaseTokenUri:      testMsgCreateClass.BaseTokenUri,
		MintingPermission: testMsgCreateClass.MintingPermission,
		TokenSupplyCap:    testMsgCreateClass.TokenSupplyCap,
	}
	suite.Require().Equal(classAttributes, expectedClassAttributes)

	// check if OwningClassIdList is set
	owningClassIdList, exists := suite.app.NftmintKeeper.GetOwningClassIdList(suite.ctx, sender.Bytes())
	suite.Require().True(exists)
	var classIdList []string
	classIdList = append(classIdList, classId)
	expectedOwningClassIdList := types.OwningClassIdList{
		Owner:   sender.Bytes(),
		ClassId: classIdList,
	}
	suite.Require().Equal(owningClassIdList, expectedOwningClassIdList)

	// check if ClassNameIdList is set
	classNameIdList, exists := suite.app.NftmintKeeper.GetClassNameIdList(suite.ctx, testMsgCreateClass.Name)
	suite.Require().True(exists)
	expectedClassNameIdList := types.ClassNameIdList{
		ClassName: testMsgCreateClass.Name,
		ClassId:   classIdList,
	}
	suite.Require().Equal(classNameIdList, expectedClassNameIdList)

	senderInInvalidCase := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	senderInInvalidCase_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, senderInInvalidCase)
	classIdInInvalidCase := keeper.CreateClassId(senderInInvalidCase_seq, senderInInvalidCase)
	// in case which contains the invalid name
	testMsgCreateClassInvalidName := types.MsgCreateClass{
		Sender:            senderInInvalidCase.Bytes(),
		Name:              "t", // shorter than the minimum length defined on UnUniFi
		BaseTokenUri:      "ipfs://testcid-sample/",
		TokenSupplyCap:    10000,
		MintingPermission: 0,
	}
	err = suite.app.NftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidName)
	suite.Require().Error(err)

	// check if data objects aren't set
	exists = suite.app.NFTKeeper.HasClass(suite.ctx, classIdInInvalidCase)
	suite.Require().False(exists)
	_, exists = suite.app.NftmintKeeper.GetClassAttributes(suite.ctx, classIdInInvalidCase)
	suite.Require().False(exists)
	_, exists = suite.app.NftmintKeeper.GetOwningClassIdList(suite.ctx, senderInInvalidCase.Bytes())
	suite.Require().False(exists)
	_, exists = suite.app.NftmintKeeper.GetClassNameIdList(suite.ctx, testMsgCreateClassInvalidName.Name)
	suite.Require().False(exists)

	// in case which contains the invalid uri
	testMsgCreateClassInvalidUri := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              "test",
		BaseTokenUri:      "ipfs", // shorter than the minimum length defined on UnUniFi
		TokenSupplyCap:    10000,
		MintingPermission: 0,
	}
	err = suite.app.NftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidUri)
	suite.Require().Error(err)

	// in case which contains the invalid token supply cap
	testMsgCreateClassInvalidTokenSupplyCap := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              "test",
		BaseTokenUri:      "ipfs",
		TokenSupplyCap:    10000000, // bigger than the token supply cap
		MintingPermission: 0,
	}
	err = suite.app.NftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidTokenSupplyCap)
	suite.Require().Error(err)

	// in case which contains the invalid minting permission
	testMsgCreateClassInvalidMintingPermission := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              "test",
		BaseTokenUri:      "ipfs",
		TokenSupplyCap:    10000,
		MintingPermission: 10,
	}
	err = suite.app.NftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidMintingPermission)
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) SendClass() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)

	testMsgCreateClass := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              "test",
		BaseTokenUri:      "ipfs://testcid-sample/",
		TokenSupplyCap:    10000,
		MintingPermission: 0,
	}
	_ = suite.app.NftmintKeeper.CreateClass(suite.ctx, classId, &testMsgCreateClass)

	recipient := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testMsgSendClass := types.MsgSendClass{
		Sender:    sender.Bytes(),
		ClassId:   classId,
		Recipient: recipient.Bytes(),
	}
	err := suite.app.NftmintKeeper.SendClass(suite.ctx, &testMsgSendClass)
	suite.Require().NoError(err)
	// check if recipient address becomes new owner of class
	classAttributes, exists := suite.app.NftmintKeeper.GetClassAttributes(suite.ctx, classId)
	suite.Require().True(exists)
	suite.Require().Equal(recipient, classAttributes.Owner.AccAddress())

	// invalid sender of MsgSendclass
	invalidSender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testMsgCreateClassInvalidSender := types.MsgSendClass{
		Sender:    invalidSender.Bytes(),
		ClassId:   classId,
		Recipient: recipient.Bytes(),
	}
	err = suite.app.NftmintKeeper.SendClass(suite.ctx, &testMsgCreateClassInvalidSender)
	suite.Require().Error(err)

	// invalid class id specification
	invalidClassId := "nonexistance"
	testMsgCreateClassInvalidClassId := types.MsgSendClass{
		Sender:    sender.Bytes(),
		ClassId:   invalidClassId,
		Recipient: recipient.Bytes(),
	}
	err = suite.app.NftmintKeeper.SendClass(suite.ctx, &testMsgCreateClassInvalidClassId)
	suite.Require().Error(err)
}

// TODO: sendClass msg

// TODO: updateTokenSupplyCap msg
// TODO: updateBaseTokenUri msg
