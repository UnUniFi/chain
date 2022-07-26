package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/UnUniFi/chain/x/nftmint/keeper"
	"github.com/UnUniFi/chain/x/nftmint/types"
)

const (
	testName              = "test"
	testBaseTokenUri      = "ipfs://testcid-sample/"
	testTokenSupplyCap    = 10000
	testMintingPermission = 0
	testNftID             = "a00"
)

// test basic functions of nftmint
func (suite *KeeperTestSuite) TestNftMintClassBasics() {
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
		Owner:   owner.Bytes(),
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

// test for the CreateClass relating functions
func (suite *KeeperTestSuite) TestCreateClass() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	err := suite.CreateClass(suite.ctx, classId, sender)
	suite.Require().NoError(err)

	// check if Class is set
	class, exists := suite.app.NFTKeeper.GetClass(suite.ctx, classId)
	suite.Require().True(exists)
	expectedClass := nfttypes.Class{
		Id:   classId,
		Name: testName,
	}
	suite.Require().Equal(class, expectedClass)

	// check if ClassAttributes is set
	classAttributes, exists := suite.app.NftmintKeeper.GetClassAttributes(suite.ctx, classId)
	suite.Require().True(exists)
	expectedClassAttributes := types.ClassAttributes{
		ClassId:           classId,
		Owner:             sender.Bytes(),
		BaseTokenUri:      testBaseTokenUri,
		MintingPermission: testMintingPermission,
		TokenSupplyCap:    testTokenSupplyCap,
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
	classNameIdList, exists := suite.app.NftmintKeeper.GetClassNameIdList(suite.ctx, testName)
	suite.Require().True(exists)
	expectedClassNameIdList := types.ClassNameIdList{
		ClassName: testName,
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
		BaseTokenUri:      testBaseTokenUri,
		TokenSupplyCap:    testTokenSupplyCap,
		MintingPermission: testMintingPermission,
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
		Name:              testName,
		BaseTokenUri:      "ipfs", // shorter than the minimum length defined on UnUniFi
		TokenSupplyCap:    testTokenSupplyCap,
		MintingPermission: testMintingPermission,
	}
	err = suite.app.NftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidUri)
	suite.Require().Error(err)

	// in case which contains the invalid token supply cap
	testMsgCreateClassInvalidTokenSupplyCap := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              testName,
		BaseTokenUri:      testBaseTokenUri,
		TokenSupplyCap:    10000000, // bigger than the token supply cap
		MintingPermission: 0,
	}
	err = suite.app.NftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidTokenSupplyCap)
	suite.Require().Error(err)

	// in case which contains the invalid minting permission
	testMsgCreateClassInvalidMintingPermission := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              testName,
		BaseTokenUri:      testBaseTokenUri,
		TokenSupplyCap:    10000,
		MintingPermission: 10, // not allowed minting permission option
	}
	err = suite.app.NftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidMintingPermission)
	suite.Require().Error(err)
}

// test for the SendClass relating functions
func (suite *KeeperTestSuite) TestSendClass() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)

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
	expectedOwner := recipient
	suite.Require().Equal(expectedOwner, classAttributes.Owner.AccAddress())

	// invalid sender of MsgSendclass
	invalidSender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testMsgSendClassInvalidSender := types.MsgSendClass{
		Sender:    invalidSender.Bytes(), // not the owner of class
		ClassId:   classId,
		Recipient: recipient.Bytes(),
	}
	err = suite.app.NftmintKeeper.SendClass(suite.ctx, &testMsgSendClassInvalidSender)
	suite.Require().Error(err)

	// invalid class id specification
	invalidClassId := "nonexistance"
	testMsgCreateClassInvalidClassId := types.MsgSendClass{
		Sender:    sender.Bytes(),
		ClassId:   invalidClassId, // non-existant class
		Recipient: recipient.Bytes(),
	}
	err = suite.app.NftmintKeeper.SendClass(suite.ctx, &testMsgCreateClassInvalidClassId)
	suite.Require().Error(err)
}

// test for the UpdateTokenSupplyCap relating functions
func (suite *KeeperTestSuite) TestUpdateTokenSupplyCap() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)

	var updatingTokenSupplyCap uint64 = 100
	testMsgUpdateTokenSupplyCap := types.MsgUpdateTokenSupplyCap{
		Sender:         sender.Bytes(),
		ClassId:        classId,
		TokenSupplyCap: updatingTokenSupplyCap,
	}
	err := suite.app.NftmintKeeper.UpdateTokenSupplyCap(suite.ctx, &testMsgUpdateTokenSupplyCap)
	suite.Require().NoError(err)
	classAttributes, exists := suite.app.NftmintKeeper.GetClassAttributes(suite.ctx, classId)
	suite.Require().True(exists)
	expectedTokenSupplyCap := updatingTokenSupplyCap
	suite.Require().Equal(expectedTokenSupplyCap, classAttributes.TokenSupplyCap)

	// invalid sender of MsgUpdateTokenSupplyCap
	invalidSender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testMsgUpdateTokenSupplyCapInvalidSender := types.MsgUpdateTokenSupplyCap{
		Sender:         invalidSender.Bytes(), // not the owner of class
		ClassId:        classId,
		TokenSupplyCap: 100,
	}
	err = suite.app.NftmintKeeper.UpdateTokenSupplyCap(suite.ctx, &testMsgUpdateTokenSupplyCapInvalidSender)
	suite.Require().Error(err)

	// invalid token supply cap specification
	testMsgUpdateTokenSupplyCapInvalidCap := types.MsgUpdateTokenSupplyCap{
		Sender:         sender.Bytes(),
		ClassId:        classId,
		TokenSupplyCap: 1000000, // bigger than the maximum token supply cap on UnUniFi
	}
	err = suite.app.NftmintKeeper.UpdateTokenSupplyCap(suite.ctx, &testMsgUpdateTokenSupplyCapInvalidCap)
	suite.Require().Error(err)

	// invalid case which current token supply is bigger than the updating supply cap
	_ = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{ClassId: classId, Id: "a00"}, sender)
	_ = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{ClassId: classId, Id: "a01"}, sender)
	testMsgUpdateTokenSupplyCapSmaller := types.MsgUpdateTokenSupplyCap{
		Sender:         sender.Bytes(),
		ClassId:        classId,
		TokenSupplyCap: 1, // smaller than the current token supply 2 of the specified class
	}
	err = suite.app.NftmintKeeper.UpdateTokenSupplyCap(suite.ctx, &testMsgUpdateTokenSupplyCapSmaller)
	suite.Require().Error(err)
}

// test for the UpdateBaseTokenUri relating functions
func (suite *KeeperTestSuite) TestUpdateBaseTokenUri() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)
	_ = suite.MintNFT(suite.ctx, classId, testNftID, sender)

	invalidSender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	var baseTokenUriInvalidLonger string
	for i := 0; i <= types.DefaultMaxUriLen; i++ {
		baseTokenUriInvalidLonger += "a"
	}

	tests := []struct {
		testCase          string
		msg               types.MsgUpdateBaseTokenUri
		validSender       bool
		validBaseTokenUir bool
	}{
		{
			testCase: "invalid sender",
			msg: types.MsgUpdateBaseTokenUri{
				Sender:       invalidSender.Bytes(), // not the owner of class
				ClassId:      classId,
				BaseTokenUri: "ipfs://testcid-sample-latest/",
			},
			validSender:       false,
			validBaseTokenUir: true,
		},
		{
			testCase: "updating BaseTokenUri is longer than the maximum length on UnUniFi",
			msg: types.MsgUpdateBaseTokenUri{
				Sender:       sender.Bytes(),
				ClassId:      classId,
				BaseTokenUri: baseTokenUriInvalidLonger,
			},
			validSender:       true,
			validBaseTokenUir: false,
		},
		{
			testCase: "updating BaseTokenUri is longer than the maximum length on UnUniFi",
			msg: types.MsgUpdateBaseTokenUri{
				Sender:       sender.Bytes(),
				ClassId:      classId,
				BaseTokenUri: "t",
			},
			validSender:       true,
			validBaseTokenUir: false,
		},
		{
			testCase: "successful case",
			msg: types.MsgUpdateBaseTokenUri{
				Sender:       sender.Bytes(),
				ClassId:      classId,
				BaseTokenUri: "ipfs://testcid-sample-latest/",
			},
			validSender:       true,
			validBaseTokenUir: true,
		},
	}

	for _, tc := range tests {
		err := suite.app.NftmintKeeper.UpdateBaseTokenUri(suite.ctx, &tc.msg)

		// invalid cases
		if !tc.validSender || !tc.validBaseTokenUir {
			suite.Require().Error(err)
			gotClassAttributes, _ := suite.app.NftmintKeeper.GetClassAttributes(suite.ctx, classId)
			suite.Require().Equal(testBaseTokenUri, gotClassAttributes.BaseTokenUri)
			nft, _ := suite.app.NFTKeeper.GetNFT(suite.ctx, classId, testNftID)
			expectedNFTUri := testBaseTokenUri + testNFTId
			suite.Require().Equal(expectedNFTUri, nft.Uri)
		}

		// valid case
		if tc.validSender && tc.validBaseTokenUir {
			suite.Require().NoError(err)
			gotClassAttributes, _ := suite.app.NftmintKeeper.GetClassAttributes(suite.ctx, classId)
			suite.Require().Equal(tc.msg.BaseTokenUri, gotClassAttributes.BaseTokenUri)
			nft, _ := suite.app.NFTKeeper.GetNFT(suite.ctx, classId, testNftID)
			expectedNFTUri := tc.msg.BaseTokenUri + testNftID
			suite.Require().Equal(expectedNFTUri, nft.Uri)
		}
	}
}

// execute CreateClass as the common function
func (suite *KeeperTestSuite) CreateClass(ctx sdk.Context, classID string, sender sdk.AccAddress) error {
	testMsgCreateClass := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              testName,
		BaseTokenUri:      testBaseTokenUri,
		TokenSupplyCap:    testTokenSupplyCap,
		MintingPermission: testMintingPermission,
	}

	err := suite.app.NftmintKeeper.CreateClass(ctx, classID, &testMsgCreateClass)
	return err
}
