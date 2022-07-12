package keeper_test

import (
	"testing"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/nftmint/keeper"
	"github.com/UnUniFi/chain/x/nftmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

const (
	testAddr = "ununifi1ghaguquuytpdgdhfthmtva0vdjmf4q99k2jhd2"
)

// test basic functions of nftmint
func (suite *KeeperTestSuite) TestNftMintBasics() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, owner)

	classId := keeper.CreateClassId(owner_seq, owner)

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
}

// test of the method to create a class.Id
func TestCreateId(t *testing.T) {
	var seq uint64 = 0
	addr := testAddr
	accAddr, _ := sdk.AccAddressFromBech32(addr)
	classIdSeq0 := keeper.CreateClassId(seq, accAddr)
	err := nfttypes.ValidateClassID(classIdSeq0)
	require.NoError(t, err)

	seq += 1
	classIdSeq1 := keeper.CreateClassId(seq, accAddr)
	require.NoError(t, err)
	require.NotEqual(t, classIdSeq0, classIdSeq1)
}

// TODO: createClass msg
// TODO: sendClass msg
// TODO: updateTokenSupplyCap msg
// TODO: updateBaseTokenUri msg
