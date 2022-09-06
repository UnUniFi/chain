package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/UnUniFi/chain/x/nftmint/keeper"
	"github.com/UnUniFi/chain/x/nftmint/types"
)

func (suite *KeeperTestSuite) TestQueryClassAttributes() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)

	req := types.QueryClassAttributesRequest{
		ClassId: classId,
	}
	res, err := suite.app.NftmintKeeper.ClassAttributes(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().Equal(classId, res.ClassAttributes.ClassId)
	suite.Require().Equal(testBaseTokenUri, res.ClassAttributes.BaseTokenUri)

	invalidReq := types.QueryClassAttributesRequest{
		ClassId: "invalidClassId",
	}
	_, err = suite.app.NftmintKeeper.ClassAttributes(suite.ctx, &invalidReq)
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestQueryNftMinter() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	sender_seq, _ := suite.app.AccountKeeper.GetSequence(suite.ctx, sender)
	classId := keeper.CreateClassId(sender_seq, sender)
	_ = suite.CreateClass(suite.ctx, classId, sender)
	_ = suite.MintNFT(suite.ctx, classId, testNFTId, sender)

	req := types.QueryNFTMinterRequest{
		ClassId: classId,
		NftId:   testNFTId,
	}
	res, err := suite.app.NftmintKeeper.NFTMinter(suite.ctx, &req)
	suite.Require().NoError(err)
	suite.Require().Equal(sender.String(), res.Minter)
}
