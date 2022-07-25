package keeper_test

import (
	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/tendermint/tendermint/crypto/ed25519"
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
	// not existing listing
	// invalid repay denom
	// check debt increased
	// check tokens are sent from end user to module
}

func (suite *KeeperTestSuite) Liquidate() {
	// TODO: implement it once logic is implemented!
}
