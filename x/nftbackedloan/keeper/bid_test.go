package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) TestPlaceBid() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	listing := types.NftListing{
		NftId:              types.NftIdentifier{ClassId: "class1", NftId: "nft1"},
		Owner:              owner.String(),
		State:              types.ListingState_LISTING,
		BidDenom:           "uguu",
		MinimumDepositRate: sdk.NewDecWithPrec(1, 1),
		StartedAt:          time.Now(),
	}

	tests := []struct {
		testCase    string
		msgBid      types.MsgPlaceBid
		expectError error
		expectPass  bool
	}{
		{
			testCase: "No listing",
			msgBid: types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            types.NftIdentifier{ClassId: "class999", NftId: "nft99"},
				BidAmount:        sdk.NewInt64Coin("uatom", 10000000),
				ExpiryAt:         time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.NewDecWithPrec(1, 1),
				AutomaticPayment: true,
				DepositAmount:    sdk.NewInt64Coin("uatom", 1000000),
			},
			expectError: types.ErrNftListingDoesNotExist,
			expectPass:  false,
		},
		{
			testCase: "Bid denom is not the same as listing denom",
			msgBid: types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            types.NftIdentifier{ClassId: "class1", NftId: "nft1"},
				BidAmount:        sdk.NewInt64Coin("uatom", 10000000),
				ExpiryAt:         time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.NewDecWithPrec(1, 1),
				AutomaticPayment: true,
				DepositAmount:    sdk.NewInt64Coin("uatom", 1000000),
			},
			expectError: types.ErrInvalidBidDenom,
			expectPass:  false,
		},
		{
			testCase: "pass first bid",
			msgBid: types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            types.NftIdentifier{ClassId: "class1", NftId: "nft1"},
				BidAmount:        sdk.NewInt64Coin("uguu", 10000000),
				ExpiryAt:         time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.NewDecWithPrec(1, 1),
				AutomaticPayment: true,
				DepositAmount:    sdk.NewInt64Coin("uguu", 1000000),
			},
			expectError: nil,
			expectPass:  true,
		},
		// Test Rebid
		// Now check from the frontend
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now()
		suite.ctx = suite.ctx.WithBlockTime(now)

		_ = suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          listing.NftId.ClassId,
			Name:        listing.NftId.ClassId,
			Symbol:      listing.NftId.ClassId,
			Description: listing.NftId.ClassId,
			Uri:         listing.NftId.ClassId,
		})
		_ = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: listing.NftId.ClassId,
			Id:      listing.NftId.NftId,
			Uri:     listing.NftId.NftId,
			UriHash: listing.NftId.NftId,
		}, owner)

		err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             listing.Owner,
			NftId:              listing.NftId,
			BidDenom:           listing.BidDenom,
			MinimumDepositRate: listing.MinimumDepositRate,
		})
		suite.Require().NoError(err)

		err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.msgBid.BidAmount})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{tc.msgBid.BidAmount})
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &tc.msgBid)
		if tc.expectPass {
			suite.NoError(err)
		} else {
			suite.Require().Equal(tc.expectError, err)
		}
	}
}

func (suite *KeeperTestSuite) TestSafeCloseBid() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	listing := types.NftListing{
		NftId:              types.NftIdentifier{ClassId: "class1", NftId: "nft1"},
		Owner:              owner.String(),
		State:              types.ListingState_LISTING,
		BidDenom:           "uguu",
		MinimumDepositRate: sdk.NewDecWithPrec(1, 1),
		StartedAt:          time.Now(),
	}

	tests := []struct {
		testCase           string
		msgBid             types.MsgPlaceBid
		expectBidAmount    sdk.Coin
		expectClosedAmount sdk.Coin
	}{
		{
			testCase: "pass first bid",
			msgBid: types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            types.NftIdentifier{ClassId: "class1", NftId: "nft1"},
				BidAmount:        sdk.NewInt64Coin("uguu", 10000000),
				ExpiryAt:         time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.NewDecWithPrec(1, 1),
				AutomaticPayment: true,
				DepositAmount:    sdk.NewInt64Coin("uguu", 1000000),
			},
			expectBidAmount:    sdk.NewInt64Coin("uguu", 9000000),
			expectClosedAmount: sdk.NewInt64Coin("uguu", 10000000),
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now()
		suite.ctx = suite.ctx.WithBlockTime(now)

		_ = suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          listing.NftId.ClassId,
			Name:        listing.NftId.ClassId,
			Symbol:      listing.NftId.ClassId,
			Description: listing.NftId.ClassId,
			Uri:         listing.NftId.ClassId,
		})
		_ = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: listing.NftId.ClassId,
			Id:      listing.NftId.NftId,
			Uri:     listing.NftId.NftId,
			UriHash: listing.NftId.NftId,
		}, owner)

		err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             listing.Owner,
			NftId:              listing.NftId,
			BidDenom:           listing.BidDenom,
			MinimumDepositRate: listing.MinimumDepositRate,
		})
		suite.Require().NoError(err)

		err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.msgBid.BidAmount})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{tc.msgBid.BidAmount})
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &tc.msgBid)
		suite.NoError(err)
		bids := types.NftBids(suite.app.NftbackedloanKeeper.GetBidsByNft(suite.ctx, listing.IdBytes()))
		bid := bids.GetBidByBidder(bidder.String())
		balance := suite.app.BankKeeper.GetBalance(suite.ctx, bidder, "uguu")
		suite.Equal(tc.expectBidAmount, balance)

		err = suite.app.NftbackedloanKeeper.SafeCloseBid(suite.ctx, bid)
		suite.NoError(err)
		balance = suite.app.BankKeeper.GetBalance(suite.ctx, bidder, "uguu")
		suite.Equal(tc.expectClosedAmount, balance)
	}
}
