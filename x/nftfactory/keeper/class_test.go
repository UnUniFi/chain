package keeper_test

import (
	"testing"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/stretchr/testify/require"

	ununifitypes "github.com/UnUniFi/chain/deprecated/types"
	"github.com/UnUniFi/chain/x/nftfactory/keeper"
	"github.com/UnUniFi/chain/x/nftfactory/types"
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
	owner_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, owner)

	classId := keeper.CreateClassId(owner_seq, owner)

	// test relating ClassAttributes
	classAttributes := types.ClassAttributes{
		ClassId: classId,
	}
	// check setting ClassAttributes function
	err := suite.nftmintKeeper.SetClassAttributes(suite.ctx, classAttributes)
	suite.Require().NoError(err)
	// check getting ClassAttributes function
	gotClassAttributes, exists := suite.nftmintKeeper.GetClassAttributes(suite.ctx, classAttributes.ClassId)
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
	err = suite.nftmintKeeper.SetOwningClassIdList(suite.ctx, owningClassIdList)
	suite.Require().NoError(err)
	// check getting OwningClassIdList function
	gotOwningClassIdList, exists := suite.nftmintKeeper.GetOwningClassIdList(suite.ctx, owner)
	suite.Require().True(exists)
	suite.Require().Equal(owningClassIdList, gotOwningClassIdList)

	// test relating ClassNameIdList
	testClassName := "test"
	classNameIdList := types.ClassNameIdList{
		ClassName: testClassName,
		ClassId:   classIdList,
	}
	// check setting ClassNameIdList function
	err = suite.nftmintKeeper.SetClassNameIdList(suite.ctx, classNameIdList)
	suite.Require().NoError(err)
	// check getting ClassNameIdList function
	gotClassNameIdList, exists := suite.nftmintKeeper.GetClassNameIdList(suite.ctx, testClassName)
	suite.Require().True(exists)
	suite.Require().Equal(classNameIdList, gotClassNameIdList)
}

// test of the method to create a class.Id
func TestCreateId(t *testing.T) {
	var seq uint64 = 0
	addr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	accAddr, _ := sdk.AccAddressFromBech32(addr.String())
	classIdSeq0 := keeper.CreateClassId(seq, accAddr)
	err := types.ValidateClassID(classIdSeq0)
	require.NoError(t, err)

	// add one to imitate actual account sequence transition of the expected situation
	seq += 1
	classIdSeq1 := keeper.CreateClassId(seq, accAddr)
	require.NoError(t, err)
	require.NotEqual(t, classIdSeq0, classIdSeq1)
}

// test for the CreateClass relating functions
func (suite *KeeperTestSuite) TestCreateClass() {
	sender := suite.addrs[0]
	sender_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	err := suite.CreateClass(suite.ctx, classId, sender)
	suite.Require().NoError(err)

	// check if Class is set
	class, exists := suite.nftKeeper.GetClass(suite.ctx, classId)
	suite.Require().True(exists)
	expectedClass := nfttypes.Class{
		Id:   classId,
		Name: testName,
	}
	suite.Require().Equal(class, expectedClass)

	// check if ClassAttributes is set
	classAttributes, exists := suite.nftmintKeeper.GetClassAttributes(suite.ctx, classId)
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
	owningClassIdList, exists := suite.nftmintKeeper.GetOwningClassIdList(suite.ctx, sender.Bytes())
	suite.Require().True(exists)
	var classIdList []string
	classIdList = append(classIdList, classId)
	expectedOwningClassIdList := types.OwningClassIdList{
		Owner:   sender.Bytes(),
		ClassId: classIdList,
	}
	suite.Require().Equal(owningClassIdList, expectedOwningClassIdList)

	// check if ClassNameIdList is set
	classNameIdList, exists := suite.nftmintKeeper.GetClassNameIdList(suite.ctx, testName)
	suite.Require().True(exists)
	expectedClassNameIdList := types.ClassNameIdList{
		ClassName: testName,
		ClassId:   classIdList,
	}
	suite.Require().Equal(classNameIdList, expectedClassNameIdList)

	senderInInvalidCase := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	senderInInvalidCase_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, senderInInvalidCase)
	classIdInInvalidCase := keeper.CreateClassId(senderInInvalidCase_seq, senderInInvalidCase)
	// in case which contains the invalid name
	testMsgCreateClassInvalidName := types.MsgCreateClass{
		Sender:            senderInInvalidCase.Bytes(),
		Name:              "t", // shorter than the minimum length defined on UnUniFi
		BaseTokenUri:      testBaseTokenUri,
		TokenSupplyCap:    testTokenSupplyCap,
		MintingPermission: testMintingPermission,
	}
	err = suite.nftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidName)
	suite.Require().Error(err)

	// check if data objects aren't set
	exists = suite.nftKeeper.HasClass(suite.ctx, classIdInInvalidCase)
	suite.Require().False(exists)
	_, exists = suite.nftmintKeeper.GetClassAttributes(suite.ctx, classIdInInvalidCase)
	suite.Require().False(exists)
	_, exists = suite.nftmintKeeper.GetOwningClassIdList(suite.ctx, senderInInvalidCase.Bytes())
	suite.Require().False(exists)
	_, exists = suite.nftmintKeeper.GetClassNameIdList(suite.ctx, testMsgCreateClassInvalidName.Name)
	suite.Require().False(exists)

	// in case which contains the invalid uri
	testMsgCreateClassInvalidUri := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              testName,
		BaseTokenUri:      "ipfs", // shorter than the minimum length defined on UnUniFi
		TokenSupplyCap:    testTokenSupplyCap,
		MintingPermission: testMintingPermission,
	}
	err = suite.nftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidUri)
	suite.Require().Error(err)

	// in case which contains the invalid token supply cap
	testMsgCreateClassInvalidTokenSupplyCap := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              testName,
		BaseTokenUri:      testBaseTokenUri,
		TokenSupplyCap:    10000000, // bigger than the token supply cap
		MintingPermission: 0,
	}
	err = suite.nftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidTokenSupplyCap)
	suite.Require().Error(err)

	// in case which contains the invalid minting permission
	testMsgCreateClassInvalidMintingPermission := types.MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              testName,
		BaseTokenUri:      testBaseTokenUri,
		TokenSupplyCap:    10000,
		MintingPermission: 10, // not allowed minting permission option
	}
	err = suite.nftmintKeeper.CreateClass(suite.ctx, classIdInInvalidCase, &testMsgCreateClassInvalidMintingPermission)
	suite.Require().Error(err)
}

// test for the SendClassOwnership relating functions
func (suite *KeeperTestSuite) TestSendClassOwnership() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)

	recipient := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testMsgSendClassOwnership := types.MsgSendClassOwnership{
		Sender:    sender.Bytes(),
		ClassId:   classId,
		Recipient: recipient.Bytes(),
	}
	err := suite.nftmintKeeper.SendClassOwnership(suite.ctx, &testMsgSendClassOwnership)
	suite.Require().NoError(err)
	// check if recipient address becomes new owner of class
	classAttributes, exists := suite.nftmintKeeper.GetClassAttributes(suite.ctx, classId)
	suite.Require().True(exists)
	expectedOwner := recipient
	suite.Require().Equal(expectedOwner, classAttributes.Owner.AccAddress())

	// invalid sender of MsgSendClassOwnership
	invalidSender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testMsgSendClassOwnershipInvalidSender := types.MsgSendClassOwnership{
		Sender:    invalidSender.Bytes(), // not the owner of class
		ClassId:   classId,
		Recipient: recipient.Bytes(),
	}
	err = suite.nftmintKeeper.SendClassOwnership(suite.ctx, &testMsgSendClassOwnershipInvalidSender)
	suite.Require().Error(err)

	// invalid class id specification
	invalidClassId := "nonexistance"
	testMsgCreateClassInvalidClassId := types.MsgSendClassOwnership{
		Sender:    sender.Bytes(),
		ClassId:   invalidClassId, // non-existant class
		Recipient: recipient.Bytes(),
	}
	err = suite.nftmintKeeper.SendClassOwnership(suite.ctx, &testMsgCreateClassInvalidClassId)
	suite.Require().Error(err)
}

// test for the UpdateTokenSupplyCap relating functions
func (suite *KeeperTestSuite) TestUpdateTokenSupplyCap() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)

	var updatingTokenSupplyCap uint64 = 100
	testMsgUpdateTokenSupplyCap := types.MsgUpdateTokenSupplyCap{
		Sender:         sender.Bytes(),
		ClassId:        classId,
		TokenSupplyCap: updatingTokenSupplyCap,
	}
	err := suite.nftmintKeeper.UpdateTokenSupplyCap(suite.ctx, &testMsgUpdateTokenSupplyCap)
	suite.Require().NoError(err)
	classAttributes, exists := suite.nftmintKeeper.GetClassAttributes(suite.ctx, classId)
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
	err = suite.nftmintKeeper.UpdateTokenSupplyCap(suite.ctx, &testMsgUpdateTokenSupplyCapInvalidSender)
	suite.Require().Error(err)

	// invalid token supply cap specification
	testMsgUpdateTokenSupplyCapInvalidCap := types.MsgUpdateTokenSupplyCap{
		Sender:         sender.Bytes(),
		ClassId:        classId,
		TokenSupplyCap: 1000000, // bigger than the maximum token supply cap on UnUniFi
	}
	err = suite.nftmintKeeper.UpdateTokenSupplyCap(suite.ctx, &testMsgUpdateTokenSupplyCapInvalidCap)
	suite.Require().Error(err)

	// invalid case which current token supply is bigger than the updating supply cap
	_ = suite.nftKeeper.Mint(suite.ctx, nfttypes.NFT{ClassId: classId, Id: "a00"}, sender)
	_ = suite.nftKeeper.Mint(suite.ctx, nfttypes.NFT{ClassId: classId, Id: "a01"}, sender)
	testMsgUpdateTokenSupplyCapSmaller := types.MsgUpdateTokenSupplyCap{
		Sender:         sender.Bytes(),
		ClassId:        classId,
		TokenSupplyCap: 1, // smaller than the current token supply 2 of the specified class
	}
	err = suite.nftmintKeeper.UpdateTokenSupplyCap(suite.ctx, &testMsgUpdateTokenSupplyCapSmaller)
	suite.Require().Error(err)
}

// test for the UpdateBaseTokenUri relating functions
func (suite *KeeperTestSuite) TestUpdateBaseTokenUri() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, sender)
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
		err := suite.nftmintKeeper.UpdateBaseTokenUri(suite.ctx, &tc.msg)

		// invalid cases
		if !tc.validSender || !tc.validBaseTokenUir {
			suite.Require().Error(err)
			gotClassAttributes, _ := suite.nftmintKeeper.GetClassAttributes(suite.ctx, classId)
			suite.Require().Equal(testBaseTokenUri, gotClassAttributes.BaseTokenUri)
			nft, _ := suite.nftKeeper.GetNFT(suite.ctx, classId, testNftID)
			expectedNFTUri := testBaseTokenUri + testNFTId
			suite.Require().Equal(expectedNFTUri, nft.Uri)
		}

		// valid case
		if tc.validSender && tc.validBaseTokenUir {
			suite.Require().NoError(err)
			gotClassAttributes, _ := suite.nftmintKeeper.GetClassAttributes(suite.ctx, classId)
			suite.Require().Equal(tc.msg.BaseTokenUri, gotClassAttributes.BaseTokenUri)
			nft, _ := suite.nftKeeper.GetNFT(suite.ctx, classId, testNftID)
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

	err := suite.nftmintKeeper.CreateClass(ctx, classID, &testMsgCreateClass)
	return err
}

// test for the GetClassAttributes relating functions
func (suite *KeeperTestSuite) TestGetClassAttributes() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, owner)
	classId := keeper.CreateClassId(owner_seq, owner)
	_ = suite.CreateClass(suite.ctx, classId, owner)

	tests := []struct {
		testCase   string
		classId    string
		validClass bool
	}{
		{
			testCase:   "invalid class name",
			classId:    "invalid_class_id",
			validClass: false,
		},
		{
			testCase:   "successful case",
			classId:    classId,
			validClass: true,
		},
	}
	for _, tc := range tests {
		res, valid := suite.nftmintKeeper.GetClassAttributes(suite.ctx, tc.classId)

		// invalid cases
		if !tc.validClass {
			suite.Require().False(valid)
			suite.Require().Equal(res.ClassId, "")
		}

		// valid case
		if tc.validClass {
			suite.Require().True(valid)
			suite.Require().Equal(res.ClassId, classId)
			suite.Require().Equal(res.Owner.AccAddress(), owner)
		}
	}
}

// test for the GetOwningClassIdList relating functions
func (suite *KeeperTestSuite) TestGetOwningClassIdList() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, owner)
	classId := keeper.CreateClassId(owner_seq, owner)
	_ = suite.CreateClass(suite.ctx, classId, owner)

	invalidOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase     string
		ownerAddress sdk.AccAddress
		validOwner   bool
	}{
		{
			testCase:     "invalid sender",
			ownerAddress: invalidOwner,
			validOwner:   false,
		},
		{
			testCase:     "successful case",
			ownerAddress: owner,
			validOwner:   true,
		},
	}
	for _, tc := range tests {
		res, valid := suite.nftmintKeeper.GetOwningClassIdList(suite.ctx, tc.ownerAddress)

		// invalid cases
		if !tc.validOwner {
			suite.Require().False(valid)
			suite.Require().Equal(len(res.ClassId), 0)
		}

		// valid case
		if tc.validOwner {
			suite.Require().True(valid)
			suite.Require().Equal(res.Owner, ununifitypes.StringAccAddress(tc.ownerAddress))
			suite.Require().Equal(res.ClassId[0], classId)
		}
	}
}

// test for the GetClassNameIdList relating functions
func (suite *KeeperTestSuite) TestGetClassNameIdList() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner_seq, _ := suite.accountKeeper.GetSequence(suite.ctx, owner)
	classId := keeper.CreateClassId(owner_seq, owner)
	_ = suite.CreateClass(suite.ctx, classId, owner)

	tests := []struct {
		testCase   string
		className  string
		validClass bool
	}{
		{
			testCase:   "invalid class name",
			className:  "invalid_class",
			validClass: false,
		},
		{
			testCase:   "successful case",
			className:  testName,
			validClass: true,
		},
	}
	for _, tc := range tests {
		res, valid := suite.nftmintKeeper.GetClassNameIdList(suite.ctx, tc.className)

		// invalid cases
		if !tc.validClass {
			suite.Require().False(valid)
			suite.Require().Equal(len(res.ClassId), 0)
		}

		// valid case
		if tc.validClass {
			suite.Require().True(valid)
			suite.Require().Equal(res.ClassId[0], classId)
		}
	}
}
