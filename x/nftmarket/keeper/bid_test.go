package keeper_test

import (
	"time"

	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

// test basic functions of bids on nft bids
func (suite *KeeperTestSuite) TestNftBidBasics() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	now := time.Now().UTC()
	bids := []types.NftBid{
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "1",
			},
			Bidder:           owner.String(),
			Amount:           sdk.NewInt64Coin("uguu", 1000000),
			AutomaticPayment: true,
			PaidAmount:       sdk.NewInt(1000000),
			BidTime:          now,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "1",
			},
			Bidder:           owner2.String(),
			Amount:           sdk.NewInt64Coin("uguu", 1000000),
			AutomaticPayment: true,
			PaidAmount:       sdk.NewInt(1000000),
			BidTime:          now,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "2",
			},
			Bidder:           owner.String(),
			Amount:           sdk.NewInt64Coin("uguu", 1000000),
			AutomaticPayment: true,
			PaidAmount:       sdk.NewInt(1000000),
			BidTime:          now,
		},
	}

	for _, bid := range bids {
		suite.app.NftmarketKeeper.SetBid(suite.ctx, bid)
	}

	for _, bid := range bids {
		bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
		suite.Require().NoError(err)
		gotBid, err := suite.app.NftmarketKeeper.GetBid(suite.ctx, bid.NftId.IdBytes(), bidder)
		suite.Require().NoError(err)
		suite.Require().Equal(bid, gotBid)
	}

	// check all bids
	allBids := suite.app.NftmarketKeeper.GetAllBids(suite.ctx)
	suite.Require().Len(allBids, len(bids))

	// check bids by bidder
	bidsByOwner := suite.app.NftmarketKeeper.GetBidsByBidder(suite.ctx, owner)
	suite.Require().Len(bidsByOwner, 2)

	// check bids by nft
	nftBids := suite.app.NftmarketKeeper.GetBidsByNft(suite.ctx, (types.NftIdentifier{
		ClassId: "1",
		NftId:   "1",
	}).IdBytes())
	suite.Require().Len(nftBids, 2)

	// delete all the bids
	for _, bid := range bids {
		suite.app.NftmarketKeeper.DeleteBid(suite.ctx, bid)
	}

	// check all bids
	allBids = suite.app.NftmarketKeeper.GetAllBids(suite.ctx)
	suite.Require().Len(allBids, 0)

	// check bids by bidder
	bidsByOwner = suite.app.NftmarketKeeper.GetBidsByBidder(suite.ctx, owner)
	suite.Require().Len(bidsByOwner, 0)

	// check bids by nft
	nftBids = suite.app.NftmarketKeeper.GetBidsByNft(suite.ctx, (types.NftIdentifier{
		ClassId: "1",
		NftId:   "1",
	}).IdBytes())
	suite.Require().Len(nftBids, 0)
}

// TODO: add test for SetCancelledBid(ctx sdk.Context, bid types.NftBid)
// TODO: add test for GetAllCancelledBids(ctx sdk.Context)
// TODO: add test for GetMaturedCancelledBids(ctx sdk.Context, endTime time.Time)
// TODO: add test for DeleteCancelledBid(ctx sdk.Context, bid types.NftBid)

// TODO: add test for SafeCloseBid(ctx sdk.Context, bid types.NftBid)

// TODO: add test for TotalActiveRankDeposit(ctx sdk.Context, nftIdBytes []byte)

// TODO: add test for PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid)
// TODO: add test for CancelBid(ctx sdk.Context, msg *types.MsgCancelBid)
// TODO: add test for PayFullBid(ctx sdk.Context, msg *types.MsgPayFullBid)
// TODO: add test for HandleMaturedCancelledBids(ctx sdk.Context)
