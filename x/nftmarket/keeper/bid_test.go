package keeper_test

import (
	"time"

	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
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

func (suite *KeeperTestSuite) TestCancelledBid() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	now := time.Now().UTC()
	cancelledBids := []types.NftBid{
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
			BidTime:          now.Add(time.Second),
		},
	}

	for _, bid := range cancelledBids {
		suite.app.NftmarketKeeper.SetCancelledBid(suite.ctx, bid)
	}

	// check all cancelled bids
	allCancelledBids := suite.app.NftmarketKeeper.GetAllCancelledBids(suite.ctx)
	suite.Require().Len(allCancelledBids, len(cancelledBids))

	// check matured cancelled bids
	maturedCancelledBids := suite.app.NftmarketKeeper.GetMaturedCancelledBids(suite.ctx, now.Add(time.Second))
	suite.Require().Len(maturedCancelledBids, 2)

	// check normal bids
	allBids := suite.app.NftmarketKeeper.GetAllBids(suite.ctx)
	suite.Require().Len(allBids, 0)

	// delete all cancelled bids
	for _, bid := range cancelledBids {
		suite.app.NftmarketKeeper.DeleteCancelledBid(suite.ctx, bid)
	}

	// check all cancelled bids
	allCancelledBids = suite.app.NftmarketKeeper.GetAllCancelledBids(suite.ctx)
	suite.Require().Len(allCancelledBids, 0)

	// check matured cancelled bids
	maturedCancelledBids = suite.app.NftmarketKeeper.GetMaturedCancelledBids(suite.ctx, now)
	suite.Require().Len(maturedCancelledBids, 0)
}

func (suite *KeeperTestSuite) TestSafeCloseBid() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

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
	}

	for _, bid := range bids {
		suite.app.NftmarketKeeper.SetBid(suite.ctx, bid)
	}

	// try safe close of bids when module account does not have enough balance
	for _, bid := range bids {
		cacheCtx, _ := suite.ctx.CacheContext()
		err := suite.app.NftmarketKeeper.SafeCloseBid(cacheCtx, bid)
		suite.Require().Error(err)
	}

	// allocate tokens to the module
	coin := sdk.NewInt64Coin("uguu", int64(1000000000))
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{coin})
	suite.NoError(err)

	// try safe close of bids when module account has enough balance
	for _, bid := range bids {
		cacheCtx, _ := suite.ctx.CacheContext()
		err := suite.app.NftmarketKeeper.SafeCloseBid(cacheCtx, bid)
		suite.Require().NoError(err)

		// check tokens are received
		balance := suite.app.BankKeeper.GetBalance(cacheCtx, owner, "uguu")
		suite.Require().True(balance.IsPositive())
	}
}

// TODO: add test for TotalActiveRankDeposit(ctx sdk.Context, nftIdBytes []byte)

// TODO: add test for PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid)
// TODO: add test for CancelBid(ctx sdk.Context, msg *types.MsgCancelBid)
// TODO: add test for PayFullBid(ctx sdk.Context, msg *types.MsgPayFullBid)
// TODO: add test for HandleMaturedCancelledBids(ctx sdk.Context)
