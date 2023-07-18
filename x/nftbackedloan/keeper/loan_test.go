package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
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
		suite.app.NftbackedloanKeeper.SetDebt(suite.ctx, debt)
	}

	for _, debt := range debts {
		loan := suite.app.NftbackedloanKeeper.GetDebtByNft(suite.ctx, debt.NftId.IdBytes())
		suite.Require().Equal(loan, debt)
	}

	// check all debts
	allDebts := suite.app.NftbackedloanKeeper.GetAllDebts(suite.ctx)
	suite.Require().Len(allDebts, len(debts))

	// delete all the debts
	for _, debt := range debts {
		suite.app.NftbackedloanKeeper.DeleteDebt(suite.ctx, debt.NftId.IdBytes())
	}

	// check all debts
	allDebts = suite.app.NftbackedloanKeeper.GetAllDebts(suite.ctx)
	suite.Require().Len(allDebts, 0)
}

func (suite *KeeperTestSuite) TestIncreaseDecreaseDebt() {
	nftIdentifier := types.NftIdentifier{
		ClassId: "1",
		NftId:   "1",
	}

	loan := suite.app.NftbackedloanKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.Coin{})

	suite.app.NftbackedloanKeeper.IncreaseDebt(suite.ctx, nftIdentifier, sdk.NewInt64Coin("uguu", 1000000))
	loan = suite.app.NftbackedloanKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.NewInt64Coin("uguu", 1000000))

	suite.app.NftbackedloanKeeper.DecreaseDebt(suite.ctx, nftIdentifier, sdk.NewInt64Coin("uguu", 500000))
	loan = suite.app.NftbackedloanKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.NewInt64Coin("uguu", 500000))

	suite.app.NftbackedloanKeeper.DecreaseDebt(suite.ctx, nftIdentifier, sdk.NewInt64Coin("uguu", 500000))
	loan = suite.app.NftbackedloanKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.NewInt64Coin("uguu", 0))
}

func (suite *KeeperTestSuite) TestBorrow() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

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
			amount:       sdk.NewInt64Coin("uguu", 1000),
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
			originAmount: sdk.NewInt64Coin("uguu", 1000),
			amount:       sdk.NewInt64Coin("uguu", 2000),
			listBefore:   true,
			expectPass:   true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		_ = suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
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
			err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:               tc.nftOwner.String(),
				NftId:                nftIdentifier,
				BidToken:             "uguu",
				MinimumDepositRate:   sdk.MustNewDecFromStr("0.01"),
				AutomaticRefinancing: false,
			})
			suite.Require().NoError(err)
		}

		for i := 0; i < tc.prevBids; i++ {
			// init tokens to addr
			bidAmount := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			depositAmount := sdk.NewInt64Coin("uguu", int64(10000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{bidAmount})
			suite.NoError(err)

			err := suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:             bidder.String(),
				NftId:              nftIdentifier,
				BidAmount:          bidAmount,
				BiddingPeriod:      time.Now().Add(time.Hour * 24),
				DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
				AutomaticPayment:   false,
				DepositAmount:      depositAmount,
			})
			suite.Require().NoError(err)
		}

		if tc.originAmount.IsPositive() {
			err := suite.app.NftbackedloanKeeper.Borrow(suite.ctx, &types.MsgBorrow{
				Sender: tc.borrower.String(),
				NftId:  nftIdentifier,
				BorrowBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: tc.originAmount,
					},
				},
			})
			suite.Require().NoError(err, tc.testCase)
		}

		oldBorrowerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.borrower, "uguu")
		err = suite.app.NftbackedloanKeeper.Borrow(suite.ctx, &types.MsgBorrow{
			Sender: tc.borrower.String(),
			NftId:  nftIdentifier,
			BorrowBids: []types.BorrowBid{
				{
					Bidder: bidder.String(),
					Amount: tc.amount,
				},
			},
		})

		if tc.expectPass {
			suite.Require().NoError(err, tc.testCase)

			// check borrow balance increase
			borrowerNewBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.borrower, "uguu")
			suite.Require().True(borrowerNewBalance.Amount.GT(oldBorrowerBalance.Amount))

			// check debt increase
			loan := suite.app.NftbackedloanKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().True(loan.Loan.Amount.IsPositive())
		} else {
			suite.Require().Error(err, tc.testCase)
		}
	}
}

func (suite *KeeperTestSuite) TestRepay() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

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
			borrowAmount: sdk.NewInt64Coin("uguu", 1000),
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
			borrowAmount: sdk.NewInt64Coin("uguu", 10000),
			amount:       sdk.NewInt64Coin("uguu", 10000),
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
			borrowAmount: sdk.NewInt64Coin("uguu", 10000),
			amount:       sdk.NewInt64Coin("uguu", 1000),
			listBefore:   true,
			expectPass:   true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		_ = suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
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
			err := suite.app.NftbackedloanKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:             tc.nftOwner.String(),
				NftId:              nftIdentifier,
				BidToken:           "uguu",
				MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			})
			suite.Require().NoError(err)
		}

		for i := 0; i < tc.prevBids; i++ {
			// init tokens to addr
			bidAmount := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			depositAmount := sdk.NewInt64Coin("uguu", int64(100000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{bidAmount})
			suite.NoError(err)

			err := suite.app.NftbackedloanKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:             bidder.String(),
				NftId:              nftIdentifier,
				BidAmount:          bidAmount,
				BiddingPeriod:      time.Now().Add(time.Hour * 24),
				DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
				AutomaticPayment:   false,
				DepositAmount:      depositAmount,
			})
			suite.Require().NoError(err)
		}

		if tc.borrowAmount.IsPositive() {
			err := suite.app.NftbackedloanKeeper.Borrow(suite.ctx, &types.MsgBorrow{
				Sender: tc.borrower.String(),
				NftId:  nftIdentifier,
				BorrowBids: []types.BorrowBid{
					{
						Bidder: bidder.String(),
						Amount: tc.borrowAmount,
					},
				},
			})
			suite.Require().NoError(err, tc.testCase)
		}

		oldRepayerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.borrower, "uguu")
		err = suite.app.NftbackedloanKeeper.Repay(suite.ctx, &types.MsgRepay{
			Sender: tc.borrower.String(),
			NftId:  nftIdentifier,
			RepayBids: []types.BorrowBid{
				{
					Bidder: bidder.String(),
					Amount: tc.amount,
				},
			},
		})

		if tc.expectPass {
			suite.Require().NoError(err, tc.testCase)

			repayerNewBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.borrower, "uguu")
			suite.Require().True(repayerNewBalance.Amount.LT(oldRepayerBalance.Amount))

			// check debt decrease
			loan := suite.app.NftbackedloanKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().True(loan.Loan.Amount.Equal(tc.borrowAmount.Amount.Sub(tc.amount.Amount)))
		} else {
			suite.Require().Error(err, tc.testCase)
		}
	}
}
