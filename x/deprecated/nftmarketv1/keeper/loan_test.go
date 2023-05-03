package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/nftmarketv1/types"
)

func (suite *KeeperTestSuite) TestDebtBasics() {
	debts := []types.Loan{
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "1",
			},
			Loan: sdk.NewInt64Coin("uguu", 1000000),
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "2",
			},
			Loan: sdk.NewInt64Coin("uguu", 1000000),
		},
	}

	for _, debt := range debts {
		suite.app.NftmarketKeeper.SetDebt(suite.ctx, debt)
	}

	for _, debt := range debts {
		loan := suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, debt.NftId.IdBytes())
		suite.Require().Equal(loan, debt)
	}

	// check all debts
	allDebts := suite.app.NftmarketKeeper.GetAllDebts(suite.ctx)
	suite.Require().Len(allDebts, len(debts))

	// delete all the debts
	for _, debt := range debts {
		suite.app.NftmarketKeeper.DeleteDebt(suite.ctx, debt.NftId.IdBytes())
	}

	// check all debts
	allDebts = suite.app.NftmarketKeeper.GetAllDebts(suite.ctx)
	suite.Require().Len(allDebts, 0)
}

func (suite *KeeperTestSuite) TestIncreaseDecreaseDebt() {
	nftIdentifier := types.NftIdentifier{
		ClassId: "1",
		NftId:   "1",
	}

	loan := suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.Coin{})

	suite.app.NftmarketKeeper.IncreaseDebt(suite.ctx, nftIdentifier, sdk.NewInt64Coin("uguu", 1000000))
	loan = suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.NewInt64Coin("uguu", 1000000))

	suite.app.NftmarketKeeper.DecreaseDebt(suite.ctx, nftIdentifier, sdk.NewInt64Coin("uguu", 500000))
	loan = suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.NewInt64Coin("uguu", 500000))

	suite.app.NftmarketKeeper.DecreaseDebt(suite.ctx, nftIdentifier, sdk.NewInt64Coin("uguu", 500000))
	loan = suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.NewInt64Coin("uguu", 0))
}

func (suite *KeeperTestSuite) TestBorrow() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase     string
		classId      string
		nftId        string
		nftOwner     sdk.AccAddress
		borrower     sdk.AccAddress
		prevBids     int
		originAmount sdk.Coin
		amount       sdk.Coin
		listBefore   bool
		expectPass   bool
	}{
		{
			testCase:     "borrow on not listed nft",
			classId:      "class1",
			nftId:        "nft1",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     0,
			originAmount: sdk.NewInt64Coin("uguu", 0),
			amount:       sdk.NewInt64Coin("uguu", 10000000),
			listBefore:   false,
			expectPass:   false,
		},
		{
			testCase:     "borrow request by non owner",
			classId:      "class3",
			nftId:        "nft3",
			nftOwner:     acc1,
			borrower:     acc2,
			prevBids:     2,
			originAmount: sdk.NewInt64Coin("uguu", 0),
			amount:       sdk.NewInt64Coin("uguu", 1),
			listBefore:   true,
			expectPass:   false,
		},
		{
			testCase:     "invalid borrow denom",
			classId:      "class2",
			nftId:        "nft2",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     2,
			originAmount: sdk.NewInt64Coin("uguu", 0),
			amount:       sdk.NewInt64Coin("xxxx", 10000000),
			listBefore:   true,
			expectPass:   false,
		},
		{
			testCase:     "more than max debt",
			classId:      "class3",
			nftId:        "nft3",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     1,
			originAmount: sdk.NewInt64Coin("uguu", 0),
			amount:       sdk.NewInt64Coin("uguu", 1000000000),
			listBefore:   true,
			expectPass:   false,
		},
		{
			testCase:     "successful 1st time borrow",
			classId:      "class5",
			nftId:        "nft5",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     2,
			originAmount: sdk.NewInt64Coin("uguu", 0),
			amount:       sdk.NewInt64Coin("uguu", 1000000),
			listBefore:   true,
			expectPass:   true,
		},
		{
			testCase:     "successful 2nd time borrow",
			classId:      "class5",
			nftId:        "nft5",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     4,
			originAmount: sdk.NewInt64Coin("uguu", 1000000),
			amount:       sdk.NewInt64Coin("uguu", 1000000),
			listBefore:   true,
			expectPass:   true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err := suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		if tc.listBefore {
			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:        ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:         nftIdentifier,
				ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:      "uguu",
				MinBid:        sdk.NewInt(10),
				BidActiveRank: 2,
			})
			suite.Require().NoError(err)
		}

		for i := 0; i < tc.prevBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{coin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           ununifitypes.StringAccAddress(bidder),
				NftId:            nftIdentifier,
				Amount:           coin,
				AutomaticPayment: false,
			})
			suite.Require().NoError(err)
		}

		if tc.originAmount.IsPositive() {
			err := suite.app.NftmarketKeeper.Borrow(suite.ctx, &types.MsgBorrow{
				Sender: ununifitypes.StringAccAddress(tc.borrower),
				NftId:  nftIdentifier,
				Amount: tc.originAmount,
			})
			suite.Require().NoError(err)
		}

		oldBorrowerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.borrower, "uguu")
		err = suite.app.NftmarketKeeper.Borrow(suite.ctx, &types.MsgBorrow{
			Sender: ununifitypes.StringAccAddress(tc.borrower),
			NftId:  nftIdentifier,
			Amount: tc.amount,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			// check borrow balance increase
			borrowerNewBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.borrower, "uguu")
			suite.Require().True(borrowerNewBalance.Amount.GT(oldBorrowerBalance.Amount))

			// check debt increase
			loan := suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().True(loan.Loan.Amount.IsPositive())
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestRepay() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase     string
		classId      string
		nftId        string
		nftOwner     sdk.AccAddress
		borrower     sdk.AccAddress
		prevBids     int
		borrowAmount sdk.Coin
		amount       sdk.Coin
		listBefore   bool
		expectPass   bool
	}{
		{
			testCase:     "repay on not listed nft",
			classId:      "class1",
			nftId:        "nft1",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     0,
			borrowAmount: sdk.NewInt64Coin("uguu", 0),
			amount:       sdk.NewInt64Coin("uguu", 10000000),
			listBefore:   false,
			expectPass:   false,
		},
		{
			testCase:     "repay request by non owner",
			classId:      "class3",
			nftId:        "nft3",
			nftOwner:     acc1,
			borrower:     acc2,
			prevBids:     2,
			borrowAmount: sdk.NewInt64Coin("uguu", 0),
			amount:       sdk.NewInt64Coin("uguu", 1),
			listBefore:   true,
			expectPass:   false,
		},
		{
			testCase:     "invalid repay denom",
			classId:      "class2",
			nftId:        "nft2",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     2,
			borrowAmount: sdk.NewInt64Coin("uguu", 0),
			amount:       sdk.NewInt64Coin("xxxx", 10000000),
			listBefore:   true,
			expectPass:   false,
		},
		{
			testCase:     "repay more than debt",
			classId:      "class3",
			nftId:        "nft3",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     1,
			borrowAmount: sdk.NewInt64Coin("uguu", 100),
			amount:       sdk.NewInt64Coin("uguu", 10000),
			listBefore:   true,
			expectPass:   false,
		},
		{
			testCase:     "successful full repay",
			classId:      "class5",
			nftId:        "nft5",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     2,
			borrowAmount: sdk.NewInt64Coin("uguu", 1000000),
			amount:       sdk.NewInt64Coin("uguu", 1000000),
			listBefore:   true,
			expectPass:   true,
		},
		{
			testCase:     "successful partial repay",
			classId:      "class5",
			nftId:        "nft5",
			nftOwner:     acc1,
			borrower:     acc1,
			prevBids:     2,
			borrowAmount: sdk.NewInt64Coin("uguu", 1000000),
			amount:       sdk.NewInt64Coin("uguu", 100000),
			listBefore:   true,
			expectPass:   true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err := suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		if tc.listBefore {
			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:        ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:         nftIdentifier,
				ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:      "uguu",
				MinBid:        sdk.NewInt(10),
				BidActiveRank: 2,
			})
			suite.Require().NoError(err)
		}

		for i := 0; i < tc.prevBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{coin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           ununifitypes.StringAccAddress(bidder),
				NftId:            nftIdentifier,
				Amount:           coin,
				AutomaticPayment: false,
			})
			suite.Require().NoError(err)
		}

		if tc.borrowAmount.IsPositive() {
			err := suite.app.NftmarketKeeper.Borrow(suite.ctx, &types.MsgBorrow{
				Sender: ununifitypes.StringAccAddress(tc.borrower),
				NftId:  nftIdentifier,
				Amount: tc.borrowAmount,
			})
			suite.Require().NoError(err)
		}

		oldRepayerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.borrower, "uguu")
		err = suite.app.NftmarketKeeper.Repay(suite.ctx, &types.MsgRepay{
			Sender: ununifitypes.StringAccAddress(tc.borrower),
			NftId:  nftIdentifier,
			Amount: tc.amount,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			repayerNewBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.borrower, "uguu")
			suite.Require().True(repayerNewBalance.Amount.LT(oldRepayerBalance.Amount))

			// check debt decrease
			loan := suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().True(loan.Loan.Amount.Equal(tc.borrowAmount.Amount.Sub(tc.amount.Amount)))
		} else {
			suite.Require().Error(err)
		}
	}
}

// TestLoanManagement is a test to see if the management of loan data is properly working.
// Specifically, this tests loan situation in HandleFullPaymentPeriodEnding method.
func (suite *KeeperTestSuite) TestLoanManagement() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)
	bidAmount := sdk.NewCoin("uguu", sdk.NewInt(100))
	nftOwner := acc1
	nftIdentifier := types.NftIdentifier{ClassId: "class1", NftId: "nft1"}

	tests := []struct {
		testCase     string
		listingState types.ListingState
		fullPay      bool
		multiBid     bool
		overBorrow   bool
	}{
		{
			testCase:     "unit borrow in selling decision listing when highest bid is paid",
			listingState: types.ListingState_SELLING_DECISION,
			fullPay:      true,
			multiBid:     false,
		}, // add successful listing state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status
		{
			testCase:     "unit borrow in selling decision listing when highest bid is not paid and no more bids",
			listingState: types.ListingState_SELLING_DECISION,
			fullPay:      false,
			multiBid:     false,
		}, // status => ListingState_LISTING
		//
		{
			testCase:     "multi borrow in selling decision listing when highest bid is not paid, and more bids",
			listingState: types.ListingState_SELLING_DECISION,
			fullPay:      false,
			multiBid:     true,
			overBorrow:   true,
		}, // status => ListingState_BIDDING
		// loan data is removed since only one bid exists.
		{
			testCase:     "multi borrow in selling decision listing when highest bid is not paid, and more bids",
			listingState: types.ListingState_SELLING_DECISION,
			fullPay:      false,
			multiBid:     true,
			overBorrow:   false,
		}, // status => ListingState_BIDDING
		// loan data is removed since only one bid exists.
		{
			testCase:     "borrow in ended listing, when fully paid bid exists",
			listingState: types.ListingState_END_LISTING,
			fullPay:      true,
			multiBid:     false,
		}, // add successful bid state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status, close all the other bids
		// and loan data is just removed.
		{
			testCase:     "borrow in ended listing, when fully paid bid does not exist",
			listingState: types.ListingState_END_LISTING,
			fullPay:      false,
			multiBid:     false,
		}, // all the bids closed, pay depositCollected, nft listing delete, transfer nft to fully paid bidder
		// and loan data is just removed.
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now().UTC()
		suite.ctx = suite.ctx.WithBlockTime(now)

		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
		_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, nftOwner, sdk.Coins{bidAmount})

		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:     nftIdentifier.ClassId,
			Name:   nftIdentifier.ClassId,
			Symbol: nftIdentifier.ClassId,
		})
		_ = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: nftIdentifier.ClassId,
			Id:      nftIdentifier.NftId,
		}, nftOwner)

		_ = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:        ununifitypes.StringAccAddress(nftOwner),
			NftId:         nftIdentifier,
			ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
			BidToken:      "uguu",
			MinBid:        sdk.ZeroInt(),
			BidActiveRank: 10,
		})
		listing, _ := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())

		if !tc.multiBid {
			suite.PlaceAndBorrow(bidAmount, nftIdentifier, nftOwner, tc.fullPay, 10)
		} else if tc.overBorrow {
			for i := 0; i < 2; i++ {
				suite.PlaceAndBorrow(bidAmount, nftIdentifier, nftOwner, tc.fullPay, 10)
			}
		} else {
			suite.PlaceAndBorrow(bidAmount, nftIdentifier, nftOwner, tc.fullPay, 10)
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
			_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
			_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{bidAmount})

			_ = suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           ununifitypes.StringAccAddress(bidder),
				NftId:            nftIdentifier,
				Amount:           bidAmount,
				AutomaticPayment: true,
			})
		}

		listing.State = tc.listingState
		suite.app.NftmarketKeeper.SetNftListing(suite.ctx, listing)

		suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second * time.Duration(params.NftListingPeriodInitial+1)))
		suite.app.NftmarketKeeper.HandleFullPaymentsPeriodEndings(suite.ctx)
		loan := suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())

		switch tc.listingState {
		case types.ListingState_SELLING_DECISION:
			if tc.fullPay {
				// afte the fullpay, every loan must be removed
				suite.Require().Empty(loan.Loan)
			} else {
				if tc.multiBid && tc.overBorrow {
					suite.Require().Equal(bidAmount.Amount.QuoRaw(10), loan.Loan.Amount)
				} else {
					suite.Require().Empty(loan.Loan)
				}
			}
		case types.ListingState_END_LISTING:
			// under END_LISTING condition, loan must be removed after HandleFullPaymentsPeriodEndings
			suite.Require().Empty(loan.Loan)
		}
	}
}

// this method is for TestLoanManagement
func (suite *KeeperTestSuite) PlaceAndBorrow(coin sdk.Coin, nftId types.NftIdentifier, nftOwner sdk.AccAddress, fullPay bool, bidActiveRank uint64) {
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
	_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{coin})

	err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
		Sender:           ununifitypes.StringAccAddress(bidder),
		NftId:            nftId,
		Amount:           coin,
		AutomaticPayment: true,
	})
	suite.Require().NoError(err)
	err = suite.app.NftmarketKeeper.Borrow(suite.ctx, &types.MsgBorrow{
		Sender: ununifitypes.StringAccAddress(nftOwner),
		NftId:  nftId,
		Amount: sdk.NewCoin("uguu", coin.Amount.Quo(sdk.NewInt(int64(bidActiveRank)))),
	})
	suite.Require().NoError(err)

	if fullPay {
		err := suite.app.NftmarketKeeper.PayFullBid(suite.ctx, &types.MsgPayFullBid{
			Sender: ununifitypes.StringAccAddress(bidder),
			NftId:  nftId,
		})
		suite.Require().NoError(err)
	}
}

func (suite *KeeperTestSuite) Liquidate() {
	// TODO: implement it once logic is implemented!
}
