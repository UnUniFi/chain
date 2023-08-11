package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) TestManualBorrow() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	listing := types.Listing{
		NftId:          types.NftId{ClassId: "class1", TokenId: "nft1"},
		Owner:          owner.String(),
		State:          types.ListingState_LISTING,
		BidDenom:       "uguu",
		MinDepositRate: sdk.NewDecWithPrec(1, 1),
		StartedAt:      time.Now(),
	}
	msgBid := types.MsgPlaceBid{
		Sender:           bidder.String(),
		NftId:            types.NftId{ClassId: "class1", TokenId: "nft1"},
		Price:            sdk.NewInt64Coin("uguu", 10000000),
		Expiry:           time.Now().Add(time.Hour * 24),
		InterestRate:     sdk.NewDecWithPrec(1, 1),
		AutomaticPayment: true,
		Deposit:          sdk.NewInt64Coin("uguu", 1000000),
	}

	tests := []struct {
		testCase     string
		msgBorrow    types.MsgBorrow
		expectPass   bool
		expectAmount sdk.Coin
	}{
		{
			testCase: "fail with not listing",
			msgBorrow: types.MsgBorrow{
				Sender: owner.String(),
				// invalid nft id
				NftId: types.NftId{ClassId: "class99", TokenId: "nft99"},
				BorrowBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: sdk.NewInt64Coin("uguu", 500000),
					},
				},
			},
			expectPass:   false,
			expectAmount: sdk.NewInt64Coin("uguu", 0)},
		{
			testCase: "fail with not owner",
			msgBorrow: types.MsgBorrow{
				// invalid sender
				Sender: bidder.String(),
				NftId:  types.NftId{ClassId: "class1", TokenId: "nft1"},
				BorrowBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: sdk.NewInt64Coin("uguu", 500000),
					},
				},
			},
			expectPass:   false,
			expectAmount: sdk.NewInt64Coin("uguu", 0),
		},
		{
			testCase: "fail with invalid denom",
			msgBorrow: types.MsgBorrow{
				Sender: owner.String(),
				NftId:  types.NftId{ClassId: "class1", TokenId: "nft1"},
				BorrowBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						// invalid denom
						Amount: sdk.NewInt64Coin("uatom", 500000),
					},
				},
			},
			expectPass:   false,
			expectAmount: sdk.NewInt64Coin("uguu", 0),
		},
		{
			testCase: "pass with partial borrow",
			msgBorrow: types.MsgBorrow{
				Sender: owner.String(),
				NftId:  types.NftId{ClassId: "class1", TokenId: "nft1"},
				BorrowBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: sdk.NewInt64Coin("uguu", 500000),
					},
				},
			},
			expectPass:   true,
			expectAmount: sdk.NewInt64Coin("uguu", 500000),
		},
		{
			testCase: "pass with over borrow",
			msgBorrow: types.MsgBorrow{
				Sender: owner.String(),
				NftId:  types.NftId{ClassId: "class1", TokenId: "nft1"},
				BorrowBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: sdk.NewInt64Coin("uguu", 2000000),
					},
				},
			},
			expectPass:   true,
			expectAmount: sdk.NewInt64Coin("uguu", 1000000),
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
			Sender:         listing.Owner,
			NftId:          listing.NftId,
			BidDenom:       listing.BidDenom,
			MinDepositRate: listing.MinDepositRate,
		})
		suite.Require().NoError(err)

		err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{msgBid.Price})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{msgBid.Price})
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &msgBid)
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.ManualBorrow(suite.ctx, tc.msgBorrow.NftId, tc.msgBorrow.BorrowBids, tc.msgBorrow.Sender)
		if tc.expectPass {
			suite.NoError(err)
			balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "uguu")
			suite.Equal(tc.expectAmount, balance)
		} else {
			suite.Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestManualRepay() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	listing := types.Listing{
		NftId:          types.NftId{ClassId: "class1", TokenId: "nft1"},
		Owner:          owner.String(),
		State:          types.ListingState_LISTING,
		BidDenom:       "uguu",
		MinDepositRate: sdk.NewDecWithPrec(1, 1),
		StartedAt:      time.Now(),
	}
	msgBid := types.MsgPlaceBid{
		Sender:           bidder.String(),
		NftId:            types.NftId{ClassId: "class1", TokenId: "nft1"},
		Price:            sdk.NewInt64Coin("uguu", 10000000),
		Expiry:           time.Now().Add(time.Hour * 24),
		InterestRate:     sdk.NewDecWithPrec(1, 1),
		AutomaticPayment: true,
		Deposit:          sdk.NewInt64Coin("uguu", 1000000),
	}
	msgBorrow := types.MsgBorrow{
		Sender: owner.String(),
		NftId:  types.NftId{ClassId: "class1", TokenId: "nft1"},
		BorrowBids: []types.BorrowBid{
			{
				Bidder: bidder.String(),
				Amount: sdk.NewInt64Coin("uguu", 1000000),
			},
		},
	}

	tests := []struct {
		testCase     string
		msgRepay     types.MsgRepay
		expectPass   bool
		expectAmount sdk.Coin
	}{
		{
			testCase: "fail with not listing",
			msgRepay: types.MsgRepay{
				Sender: owner.String(),
				// invalid nft id
				NftId: types.NftId{ClassId: "class99", TokenId: "nft99"},
				RepayBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: sdk.NewInt64Coin("uguu", 500000),
					},
				},
			},
			expectPass:   false,
			expectAmount: sdk.NewInt64Coin("uguu", 0)},
		{
			testCase: "fail with not owner",
			msgRepay: types.MsgRepay{
				// invalid sender
				Sender: bidder.String(),
				NftId:  types.NftId{ClassId: "class1", TokenId: "nft1"},
				RepayBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: sdk.NewInt64Coin("uguu", 500000),
					},
				},
			},
			expectPass:   false,
			expectAmount: sdk.NewInt64Coin("uguu", 0),
		},
		{
			testCase: "fail with invalid denom",
			msgRepay: types.MsgRepay{
				Sender: owner.String(),
				NftId:  types.NftId{ClassId: "class1", TokenId: "nft1"},
				RepayBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						// invalid denom
						Amount: sdk.NewInt64Coin("uatom", 500000),
					},
				},
			},
			expectPass:   false,
			expectAmount: sdk.NewInt64Coin("uguu", 0),
		},
		{
			testCase: "pass with partial repay",
			msgRepay: types.MsgRepay{
				Sender: owner.String(),
				NftId:  types.NftId{ClassId: "class1", TokenId: "nft1"},
				RepayBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: sdk.NewInt64Coin("uguu", 500000),
					},
				},
			},
			expectPass:   true,
			expectAmount: sdk.NewInt64Coin("uguu", 1500000),
		},
		{
			testCase: "pass with over repay",
			msgRepay: types.MsgRepay{
				Sender: owner.String(),
				NftId:  types.NftId{ClassId: "class1", TokenId: "nft1"},
				RepayBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: sdk.NewInt64Coin("uguu", 2000000),
					},
				},
			},
			expectPass:   true,
			expectAmount: sdk.NewInt64Coin("uguu", 1000000),
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
			Sender:         listing.Owner,
			NftId:          listing.NftId,
			BidDenom:       listing.BidDenom,
			MinDepositRate: listing.MinDepositRate,
		})
		suite.Require().NoError(err)

		err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{msgBid.Price})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{msgBid.Deposit})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, owner, sdk.Coins{msgBid.Deposit})
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &msgBid)
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.ManualBorrow(suite.ctx, msgBorrow.NftId, msgBorrow.BorrowBids, msgBorrow.Sender)
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.ManualRepay(suite.ctx, tc.msgRepay.NftId, tc.msgRepay.RepayBids, tc.msgRepay.Sender)
		if tc.expectPass {
			suite.NoError(err)
			balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "uguu")
			suite.Equal(tc.expectAmount, balance)
		} else {
			suite.Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestSendInterestToBidder() {
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	initAmount := sdk.NewInt64Coin("uguu", 1000000)
	tests := []struct {
		testCase     string
		bid          types.Bid
		interest     sdk.Coin
		expectPass   bool
		expectAmount sdk.Coin
	}{
		{
			testCase: "pass with valid interest",
			bid: types.Bid{
				Id: types.BidId{
					Bidder: bidder.String(),
				},
			},
			interest:     sdk.NewInt64Coin("uguu", 1000000),
			expectPass:   true,
			expectAmount: sdk.NewInt64Coin("uguu", 1000000),
		},
		{
			testCase: "fail with 0 interest",
			bid: types.Bid{
				Id: types.BidId{
					Bidder: bidder.String(),
				},
			},
			interest:     sdk.NewInt64Coin("uguu", 0),
			expectPass:   true,
			expectAmount: sdk.NewInt64Coin("uguu", 0),
		},
		{
			testCase: "fail with nil interest",
			bid: types.Bid{
				Id: types.BidId{
					Bidder: bidder.String(),
				},
			},
			interest:     sdk.Coin{},
			expectPass:   false,
			expectAmount: sdk.NewInt64Coin("uguu", 0),
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now()
		suite.ctx = suite.ctx.WithBlockTime(now)

		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{initAmount})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{initAmount})
		suite.NoError(err)

		err = suite.app.NftbackedloanKeeper.SendInterestToBidder(suite.ctx, tc.bid, tc.interest)
		if tc.expectPass {
			suite.NoError(err)
			balance := suite.app.BankKeeper.GetBalance(suite.ctx, bidder, "uguu")
			suite.Equal(tc.expectAmount, balance)
		} else {
			suite.Error(err)
		}
	}
}
