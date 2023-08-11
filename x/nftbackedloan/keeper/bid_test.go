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

	listing := types.Listing{
		NftId:              types.NftId{ClassId: "class1", TokenId: "nft1"},
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
				NftId:            types.NftId{ClassId: "class999", TokenId: "nft99"},
				Price:            sdk.NewInt64Coin("uatom", 10000000),
				Expiry:           time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.NewDecWithPrec(1, 1),
				AutomaticPayment: true,
				Deposit:          sdk.NewInt64Coin("uatom", 1000000),
			},
			expectError: types.ErrNftListingDoesNotExist,
			expectPass:  false,
		},
		{
			testCase: "Bid denom is not the same as listing denom",
			msgBid: types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            types.NftId{ClassId: "class1", TokenId: "nft1"},
				Price:            sdk.NewInt64Coin("uatom", 10000000),
				Expiry:           time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.NewDecWithPrec(1, 1),
				AutomaticPayment: true,
				Deposit:          sdk.NewInt64Coin("uatom", 1000000),
			},
			expectError: types.ErrInvalidPriceDenom,
			expectPass:  false,
		},
		{
			testCase: "pass first bid",
			msgBid: types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            types.NftId{ClassId: "class1", TokenId: "nft1"},
				Price:            sdk.NewInt64Coin("uguu", 10000000),
				Expiry:           time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.NewDecWithPrec(1, 1),
				AutomaticPayment: true,
				Deposit:          sdk.NewInt64Coin("uguu", 1000000),
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

		_ = suite.app.UnUniFiNFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:     listing.NftId.ClassId,
			Name:   listing.NftId.ClassId,
			Symbol: listing.NftId.ClassId,
		})
		_ = suite.app.UnUniFiNFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: listing.NftId.ClassId,
			Id:      listing.NftId.TokenId,
		}, owner)

		err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             listing.Owner,
			NftId:              listing.NftId,
			BidDenom:           listing.BidDenom,
			MinimumDepositRate: listing.MinimumDepositRate,
		})
		suite.Require().NoError(err)

		err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.msgBid.Price})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{tc.msgBid.Price})
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &tc.msgBid)
		if tc.expectPass {
			suite.NoError(err)
		} else {
			suite.Require().Equal(tc.expectError, err)
		}
	}
}

// CloseBid Test
func (suite *KeeperTestSuite) TestSafeCloseBid() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	listing := types.Listing{
		NftId:              types.NftId{ClassId: "class1", TokenId: "nft1"},
		Owner:              owner.String(),
		State:              types.ListingState_LISTING,
		BidDenom:           "uguu",
		MinimumDepositRate: sdk.NewDecWithPrec(1, 1),
		StartedAt:          time.Now(),
	}

	tests := []struct {
		testCase           string
		msgBid             types.MsgPlaceBid
		expectPrice        sdk.Coin
		expectClosedAmount sdk.Coin
	}{
		{
			testCase: "pass first bid",
			msgBid: types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            types.NftId{ClassId: "class1", TokenId: "nft1"},
				Price:            sdk.NewInt64Coin("uguu", 10000000),
				Expiry:           time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.NewDecWithPrec(1, 1),
				AutomaticPayment: true,
				Deposit:          sdk.NewInt64Coin("uguu", 1000000),
			},
			expectPrice:        sdk.NewInt64Coin("uguu", 9000000),
			expectClosedAmount: sdk.NewInt64Coin("uguu", 10000000),
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now()
		suite.ctx = suite.ctx.WithBlockTime(now)

		_ = suite.app.UnUniFiNFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          listing.NftId.ClassId,
			Name:        listing.NftId.ClassId,
			Symbol:      listing.NftId.ClassId,
			Description: listing.NftId.ClassId,
			Uri:         listing.NftId.ClassId,
		})
		_ = suite.app.UnUniFiNFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: listing.NftId.ClassId,
			Id:      listing.NftId.TokenId,
			Uri:     listing.NftId.TokenId,
			UriHash: listing.NftId.TokenId,
		}, owner)

		err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             listing.Owner,
			NftId:              listing.NftId,
			BidDenom:           listing.BidDenom,
			MinimumDepositRate: listing.MinimumDepositRate,
		})
		suite.Require().NoError(err)

		err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.msgBid.Price})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{tc.msgBid.Price})
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &tc.msgBid)
		suite.NoError(err)
		bid, err := suite.app.NftbackedloanKeeper.GetBid(suite.ctx, tc.msgBid.NftId.IdBytes(), bidder)
		suite.NoError(err)
		balance := suite.app.BankKeeper.GetBalance(suite.ctx, bidder, "uguu")
		suite.Equal(tc.expectPrice, balance)

		err = suite.app.NftbackedloanKeeper.SafeCloseBid(suite.ctx, bid)
		suite.NoError(err)
		balance = suite.app.BankKeeper.GetBalance(suite.ctx, bidder, "uguu")
		suite.Equal(tc.expectClosedAmount, balance)
	}
}

func (suite *KeeperTestSuite) TestPayRemainder() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	listing := types.Listing{
		NftId:              types.NftId{ClassId: "class1", TokenId: "nft1"},
		Owner:              owner.String(),
		State:              types.ListingState_LISTING,
		BidDenom:           "uguu",
		MinimumDepositRate: sdk.NewDecWithPrec(1, 1),
		StartedAt:          time.Now(),
	}

	tests := []struct {
		testCase           string
		msgBid             types.MsgPlaceBid
		initAmount         sdk.Coin
		expectAfterDeposit sdk.Coin
		expectAfterPayment sdk.Coin
	}{
		{
			testCase: "pass first bid",
			msgBid: types.MsgPlaceBid{
				Sender:           bidder.String(),
				NftId:            types.NftId{ClassId: "class1", TokenId: "nft1"},
				Price:            sdk.NewInt64Coin("uguu", 10000000),
				Expiry:           time.Now().Add(time.Hour * 24),
				InterestRate:     sdk.NewDecWithPrec(1, 1),
				AutomaticPayment: true,
				Deposit:          sdk.NewInt64Coin("uguu", 1000000),
			},
			initAmount:         sdk.NewInt64Coin("uguu", 20000000),
			expectAfterDeposit: sdk.NewInt64Coin("uguu", 19000000),
			expectAfterPayment: sdk.NewInt64Coin("uguu", 10000000),
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now()
		suite.ctx = suite.ctx.WithBlockTime(now)

		_ = suite.app.UnUniFiNFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          listing.NftId.ClassId,
			Name:        listing.NftId.ClassId,
			Symbol:      listing.NftId.ClassId,
			Description: listing.NftId.ClassId,
			Uri:         listing.NftId.ClassId,
		})
		_ = suite.app.UnUniFiNFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: listing.NftId.ClassId,
			Id:      listing.NftId.TokenId,
			Uri:     listing.NftId.TokenId,
			UriHash: listing.NftId.TokenId,
		}, owner)

		err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             listing.Owner,
			NftId:              listing.NftId,
			BidDenom:           listing.BidDenom,
			MinimumDepositRate: listing.MinimumDepositRate,
		})
		suite.Require().NoError(err)

		err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.initAmount})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{tc.initAmount})
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &tc.msgBid)
		suite.NoError(err)
		balance := suite.app.BankKeeper.GetBalance(suite.ctx, bidder, "uguu")
		suite.Equal(tc.expectAfterDeposit, balance)

		err = suite.app.NftbackedloanKeeper.PayRemainder(suite.ctx, &types.MsgPayRemainder{
			Sender: bidder.String(),
			NftId:  tc.msgBid.NftId,
		})
		suite.NoError(err)
		balance = suite.app.BankKeeper.GetBalance(suite.ctx, bidder, "uguu")
		suite.Equal(tc.expectAfterPayment, balance)
	}
}
