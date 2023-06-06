package keeper_test

import (
	"fmt"
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	ununifitypes "github.com/UnUniFi/chain/deprecated/types"
	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) TestNftBidBasics() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	normalListing := types.NftListing{
		NftId: types.NftIdentifier{
			ClassId: "class1",
			NftId:   "nft1",
		},
		Owner:                owner.String(),
		ListingType:          types.ListingType_DIRECT_ASSET_BORROW,
		BidToken:             "uguu",
		MinimumDepositRate:   sdk.MustNewDecFromStr("0.1"),
		AutomaticRefinancing: true,
		MinimumBiddingPeriod: time.Hour * 1,
		State:                types.ListingState_LISTING,
	}
	endListing := types.NftListing{
		NftId: types.NftIdentifier{
			ClassId: "class2",
			NftId:   "nft2",
		},
		Owner:                owner.String(),
		ListingType:          types.ListingType_DIRECT_ASSET_BORROW,
		BidToken:             "uguu",
		MinimumDepositRate:   sdk.MustNewDecFromStr("0.1"),
		AutomaticRefinancing: true,
		MinimumBiddingPeriod: time.Hour * 1,
		State:                types.ListingState_END_LISTING,
	}
	biddingListing := types.NftListing{
		NftId: types.NftIdentifier{
			ClassId: "class3",
			NftId:   "nft3",
		},
		Owner:                owner.String(),
		ListingType:          types.ListingType_DIRECT_ASSET_BORROW,
		BidToken:             "uguu",
		MinimumDepositRate:   sdk.MustNewDecFromStr("0.1"),
		AutomaticRefinancing: true,
		MinimumBiddingPeriod: time.Hour * 1,
		State:                types.ListingState_BIDDING,
	}
	testCases := []struct {
		testCase string
		bid      types.MsgPlaceBid
		listing  *types.NftListing
		expErr   error
	}{
		{
			testCase: "test basic functions of bids on nft bids1",
			bid: types.MsgPlaceBid{
				Sender:           bidder.Bytes(),
				NftId:            normalListing.NftId,
				BidAmount:        sdk.NewInt64Coin("uguu", 1000000),
				DepositAmount:    sdk.NewInt64Coin("uguu", 1000000/2),
				AutomaticPayment: true,
				BiddingPeriod:    time.Now().Add(time.Hour * 2),
			},
			listing: &normalListing,
			expErr:  nil,
		},
		{
			testCase: "test basic functions of bids on nft bids2",
			bid: types.MsgPlaceBid{
				Sender:           bidder.Bytes(),
				NftId:            biddingListing.NftId,
				BidAmount:        sdk.NewInt64Coin("uguu", 1000000),
				DepositAmount:    sdk.NewInt64Coin("uguu", 1000000/2),
				AutomaticPayment: true,
				BiddingPeriod:    time.Now().Add(time.Hour * 2),
			},
			listing: &biddingListing,
			expErr:  nil,
		},
		{
			testCase: "not exist listing",
			bid: types.MsgPlaceBid{
				Sender:           bidder.Bytes(),
				NftId:            normalListing.NftId,
				BidAmount:        sdk.NewInt64Coin("uguu", 1000000),
				DepositAmount:    sdk.NewInt64Coin("uguu", 1000000/2),
				AutomaticPayment: true,
				BiddingPeriod:    time.Now().Add(time.Hour * 2),
			},
			listing: nil,
			expErr:  fmt.Errorf("nft listing does not exist"),
		},
		{
			testCase: "can not bid",
			bid: types.MsgPlaceBid{
				Sender:           bidder.Bytes(),
				NftId:            endListing.NftId,
				BidAmount:        sdk.NewInt64Coin("uguu", 1000000),
				DepositAmount:    sdk.NewInt64Coin("uguu", 1000000/2),
				AutomaticPayment: true,
				BiddingPeriod:    time.Now().Add(time.Hour * 2),
			},
			listing: &endListing,
			expErr:  types.ErrNftListingNotInBidState,
		},
		{
			testCase: "invalid bid token",
			bid: types.MsgPlaceBid{
				Sender:           bidder.Bytes(),
				NftId:            normalListing.NftId,
				BidAmount:        sdk.NewInt64Coin("xxx", 1000000),
				DepositAmount:    sdk.NewInt64Coin("xxx", 1000000/2),
				AutomaticPayment: true,
				BiddingPeriod:    time.Now().Add(time.Hour * 2),
			},
			listing: &normalListing,
			expErr:  types.ErrInvalidBidDenom,
		},
		{
			testCase: "not enough bidding period",
			bid: types.MsgPlaceBid{
				Sender:           bidder.Bytes(),
				NftId:            normalListing.NftId,
				BidAmount:        sdk.NewInt64Coin("uguu", 1000000),
				DepositAmount:    sdk.NewInt64Coin("uguu", 1000000/2),
				AutomaticPayment: true,
				BiddingPeriod:    time.Now().Add(time.Minute * 1),
			},
			listing: &normalListing,
			expErr:  types.ErrSmallBiddingPeriod,
		},
	}
	for _, tc := range testCases {
		var err error
		initUguuAmount := sdk.NewInt64Coin("uguu", 100000000)
		initAmount := sdk.NewCoins(initUguuAmount)
		err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, initAmount)
		suite.Require().NoError(err, tc.testCase)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName,
			tc.bid.Sender.AccAddress(), initAmount)
		suite.Require().NoError(err, tc.testCase)
		if tc.listing != nil {
			suite.keeper.SetNftListing(suite.ctx, *tc.listing)
		}

		err = suite.keeper.PlaceBid(suite.ctx, &tc.bid)
		if tc.expErr != nil {
			suite.Require().Error(err, tc.testCase)
			suite.Require().Equal(tc.expErr.Error(), err.Error(), tc.testCase)
			continue
		}

		suite.Require().NoError(err, tc.testCase)
		afterAmount := suite.app.BankKeeper.GetBalance(suite.ctx, tc.bid.Sender.AccAddress(), "uguu")
		suite.Equal(afterAmount.Add(tc.bid.DepositAmount), initUguuAmount, tc.testCase)

		// cleanup
		suite.keeper.DeleteNftListing(suite.ctx, *tc.listing)
		err = suite.app.BankKeeper.SendCoinsFromAccountToModule(suite.ctx, tc.bid.Sender.AccAddress(), types.ModuleName, sdk.NewCoins(afterAmount))
		suite.Require().NoError(err, tc.testCase)
		cleanupAmount := suite.app.BankKeeper.GetBalance(suite.ctx, tc.bid.Sender.AccAddress(), "uguu")
		suite.Equal(cleanupAmount, sdk.NewCoin("uguu", sdk.ZeroInt()), tc.testCase)
	}
}

//todo make Rebid tests
//todo make FirstBid tests
//todo make ManualBid tests

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
			BidAmount:        sdk.NewInt64Coin("uguu", 1000000),
			AutomaticPayment: true,
			DepositAmount:    sdk.NewInt64Coin("uguu", 1000000),
			BidTime:          now,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "1",
			},
			Bidder:           owner2.String(),
			BidAmount:        sdk.NewInt64Coin("uguu", 1000000),
			AutomaticPayment: true,
			DepositAmount:    sdk.NewInt64Coin("uguu", 1000000),
			BidTime:          now,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "2",
			},
			Bidder:           owner.String(),
			BidAmount:        sdk.NewInt64Coin("uguu", 1000000),
			AutomaticPayment: true,
			DepositAmount:    sdk.NewInt64Coin("uguu", 1000000),
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
			BidAmount:        sdk.NewInt64Coin("uguu", 1000000),
			AutomaticPayment: true,
			DepositAmount:    sdk.NewInt64Coin("uguu", 1000000),
			BidTime:          now,
			PaidAmount:       sdk.NewInt64Coin("uguu", 0),
		},
	}

	for _, bid := range bids {
		suite.app.NftmarketKeeper.SetBid(suite.ctx, bid)
	}

	// try safe close of bids when module account does not have enough balance
	for i, bid := range bids {
		cacheCtx, _ := suite.ctx.CacheContext()
		err := suite.app.NftmarketKeeper.SafeCloseBid(cacheCtx, bid)
		suite.Require().Error(err, i)
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

func (suite *KeeperTestSuite) TestCancelBid() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase        string
		classId         string
		nftId           string
		nftOwner        sdk.AccAddress
		bidder          sdk.AccAddress
		prevBids        int
		bidAmount       sdk.Coin
		depositAmount   sdk.Coin
		listBefore      bool
		cancelAfter     time.Duration
		loanAmount      sdk.Coin
		expectPass      bool
		expectCancelFee bool
	}{
		{
			testCase:        "bid on not listed nft",
			classId:         "class1",
			nftId:           "nft1",
			nftOwner:        acc1,
			bidder:          bidder,
			prevBids:        0,
			bidAmount:       sdk.NewInt64Coin("uguu", 0),
			depositAmount:   sdk.NewInt64Coin("uguu", 0),
			listBefore:      false,
			loanAmount:      sdk.NewInt64Coin("uguu", 0),
			cancelAfter:     time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			expectPass:      false,
			expectCancelFee: false,
		},
		{
			testCase:        "did not bid previously",
			classId:         "class4",
			nftId:           "nft4",
			nftOwner:        acc1,
			bidder:          bidder,
			prevBids:        1,
			bidAmount:       sdk.NewInt64Coin("uguu", 0),
			depositAmount:   sdk.NewInt64Coin("uguu", 0),
			listBefore:      true,
			loanAmount:      sdk.NewInt64Coin("uguu", 0),
			cancelAfter:     time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			expectPass:      false,
			expectCancelFee: false,
		},
		{
			testCase:        "cancelling just after bid",
			classId:         "class2",
			nftId:           "nft2",
			nftOwner:        acc1,
			bidder:          bidder,
			prevBids:        1,
			bidAmount:       sdk.NewInt64Coin("uguu", 10000000),
			depositAmount:   sdk.NewInt64Coin("uguu", 1000000),
			listBefore:      true,
			loanAmount:      sdk.NewInt64Coin("uguu", 0),
			cancelAfter:     0,
			expectPass:      false,
			expectCancelFee: false,
		},
		{
			testCase:        "cancel single bid case",
			classId:         "class3",
			nftId:           "nft3",
			nftOwner:        acc1,
			bidder:          bidder,
			prevBids:        0,
			bidAmount:       sdk.NewInt64Coin("uguu", 10000000),
			depositAmount:   sdk.NewInt64Coin("uguu", 1000000),
			listBefore:      true,
			loanAmount:      sdk.NewInt64Coin("uguu", 0),
			cancelAfter:     time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			expectPass:      false,
			expectCancelFee: false,
		},
		// {
		// 	testCase:        "successful bid cancel on active rank with loan with cancel fee",
		// 	classId:         "class5",
		// 	nftId:           "nft5",
		// 	nftOwner:        acc1,
		// 	bidder:          bidder,
		// 	prevBids:        2,
		// 	bidAmount:       sdk.NewInt64Coin("uguu", 100000000),
		// 	depositAmount:   sdk.NewInt64Coin("uguu", 10000000),
		// 	listBefore:      true,
		// 	loanAmount:      sdk.NewInt64Coin("uguu", 10000000),
		// 	cancelAfter:     time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
		// 	expectPass:      true,
		// 	expectCancelFee: true,
		// },
		// {
		// 	testCase:        "successful bid cancel on active rank without loan without cancel fee",
		// 	classId:         "class5",
		// 	nftId:           "nft5",
		// 	nftOwner:        acc1,
		// 	bidder:          bidder,
		// 	prevBids:        2,
		// 	bidAmount:       sdk.NewInt64Coin("uguu", 100000000),
		// 	depositAmount:   sdk.NewInt64Coin("uguu", 10000000),
		// 	listBefore:      true,
		// 	loanAmount:      sdk.NewInt64Coin("uguu", 0),
		// 	cancelAfter:     time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
		// 	expectPass:      true,
		// 	expectCancelFee: false,
		// },
		// {
		// 	testCase:        "successful bid cancel on not active rank",
		// 	classId:         "class5",
		// 	nftId:           "nft5",
		// 	nftOwner:        acc1,
		// 	bidder:          bidder,
		// 	prevBids:        2,
		// 	bidAmount:       sdk.NewInt64Coin("uguu", 1000),
		// 	depositAmount:   sdk.NewInt64Coin("uguu", 100),
		// 	listBefore:      true,
		// 	loanAmount:      sdk.NewInt64Coin("uguu", 0),
		// 	cancelAfter:     time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
		// 	expectPass:      true,
		// 	expectCancelFee: false,
		// },
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now().UTC()
		suite.ctx = suite.ctx.WithBlockTime(now)

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
				Sender:             ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:              nftIdentifier,
				ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:           "uguu",
				MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			})
			suite.Require().NoError(err)
		}

		for i := 0; i < tc.prevBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

			// init tokens to addr
			bidAmount := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			depositAmount := sdk.NewInt64Coin("uguu", int64(100000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{bidAmount})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:             ununifitypes.StringAccAddress(bidder),
				NftId:              nftIdentifier,
				BidAmount:          bidAmount,
				BiddingPeriod:      time.Now().Add(time.Hour * 24),
				DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
				AutomaticPayment:   false,
				DepositAmount:      depositAmount,
			})
			fmt.Println(bidAmount)
			suite.Require().NoError(err)
		}

		if tc.bidAmount.IsPositive() {
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.bidAmount})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.bidder, sdk.Coins{tc.bidAmount})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:             ununifitypes.StringAccAddress(bidder),
				NftId:              nftIdentifier,
				BidAmount:          tc.bidAmount,
				BiddingPeriod:      time.Now().Add(time.Hour * 24),
				DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
				AutomaticPayment:   false,
				DepositAmount:      tc.depositAmount,
			})
			suite.Require().NoError(err)
		}

		// originBid, _ := suite.app.NftmarketKeeper.GetBid(suite.ctx, nftIdentifier.IdBytes(), tc.bidder)

		if tc.loanAmount.IsPositive() {
			suite.app.NftmarketKeeper.SetDebt(suite.ctx, types.Loan{
				NftId: nftIdentifier,
				Loan:  tc.loanAmount,
			})
		}
		suite.ctx = suite.ctx.WithBlockTime(now.Add(tc.cancelAfter))
		err = suite.app.NftmarketKeeper.CancelBid(suite.ctx, &types.MsgCancelBid{
			Sender: ununifitypes.StringAccAddress(tc.bidder),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			// bid removal check
			_, err := suite.app.NftmarketKeeper.GetBid(suite.ctx, nftIdentifier.IdBytes(), tc.bidder)
			suite.Require().Error(err)

			// cancelled bid creation check
			cancelledBids := suite.app.NftmarketKeeper.GetAllCancelledBids(suite.ctx)
			suite.Require().Len(cancelledBids, 1)

			// cancelled bid delievery time check
			suite.Require().Equal(cancelledBids[0].BidTime, suite.ctx.BlockTime().Add(time.Duration(params.BidTokenDisburseSecondsAfterCancel)*time.Second))

			// cancel fee check if in active rank
			// if tc.expectCancelFee {
			// 	suite.Require().True(cancelledBids[0].PaidAmount.LT(originBid.PaidAmount))
			// } else {
			// 	suite.Require().True(cancelledBids[0].PaidAmount.Equal(originBid.PaidAmount))
			// }
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestPayFullBid() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase        string
		classId         string
		nftId           string
		nftOwner        sdk.AccAddress
		bidder          sdk.AccAddress
		bidAmount       sdk.Coin
		depositAmount   sdk.Coin
		listBefore      bool
		loanAmount      sdk.Coin
		expectPass      bool
		expectCancelFee bool
	}{
		{
			testCase:        "bid on not listed nft",
			classId:         "class1",
			nftId:           "nft1",
			nftOwner:        acc1,
			bidder:          bidder,
			bidAmount:       sdk.NewInt64Coin("uguu", 0),
			depositAmount:   sdk.NewInt64Coin("uguu", 0),
			listBefore:      false,
			loanAmount:      sdk.NewInt64Coin("uguu", 0),
			expectPass:      false,
			expectCancelFee: false,
		},
		{
			testCase:        "did not bid previously",
			classId:         "class4",
			nftId:           "nft4",
			nftOwner:        acc1,
			bidder:          bidder,
			bidAmount:       sdk.NewInt64Coin("uguu", 0),
			depositAmount:   sdk.NewInt64Coin("uguu", 0),
			listBefore:      true,
			loanAmount:      sdk.NewInt64Coin("uguu", 0),
			expectPass:      false,
			expectCancelFee: false,
		},
		{
			testCase:        "successful full pay",
			classId:         "class5",
			nftId:           "nft5",
			nftOwner:        acc1,
			bidder:          bidder,
			bidAmount:       sdk.NewInt64Coin("uguu", 100000000),
			depositAmount:   sdk.NewInt64Coin("uguu", 10000000),
			listBefore:      true,
			loanAmount:      sdk.NewInt64Coin("uguu", 10000000),
			expectPass:      true,
			expectCancelFee: true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now().UTC()
		suite.ctx = suite.ctx.WithBlockTime(now)

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
				Sender:             ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:              nftIdentifier,
				ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:           "uguu",
				MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			})
			suite.Require().NoError(err)
		}

		if tc.bidAmount.IsPositive() {
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.bidAmount})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.bidder, sdk.Coins{tc.bidAmount})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:             ununifitypes.StringAccAddress(bidder),
				NftId:              nftIdentifier,
				BidAmount:          tc.bidAmount,
				BiddingPeriod:      time.Now().Add(time.Hour * 24),
				DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
				AutomaticPayment:   false,
				DepositAmount:      tc.depositAmount,
			})
			suite.Require().NoError(err)
		}

		oldBidderBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.bidder, "uguu")

		err = suite.app.NftmarketKeeper.PayFullBid(suite.ctx, &types.MsgPayFullBid{
			Sender: ununifitypes.StringAccAddress(tc.bidder),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			// check balance changes after execution
			newBidderBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.bidder, "uguu")
			suite.Require().True(newBidderBalance.Amount.LT(oldBidderBalance.Amount))

			// check paid amount changes after execution
			bid, err := suite.app.NftmarketKeeper.GetBid(suite.ctx, nftIdentifier.IdBytes(), tc.bidder)
			suite.Require().NoError(err)
			suite.Require().Equal(bid.BidAmount.Amount, bid.PaidAmount.Amount.Add(bid.DepositAmount.Amount), tc.testCase)

			// re-execute full pay
			err = suite.app.NftmarketKeeper.PayFullBid(suite.ctx, &types.MsgPayFullBid{
				Sender: ununifitypes.StringAccAddress(tc.bidder),
				NftId:  nftIdentifier,
			})
			suite.Require().NoError(err)

			// check balance after reexecution
			new2BidderBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.bidder, "uguu")
			suite.Require().True(newBidderBalance.Amount.Equal(new2BidderBalance.Amount))
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestHandleMaturedCancelledBids() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	now := time.Now().UTC()
	cancelledBids := []types.NftBid{
		// TODO: check again
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "1",
			},
			Bidder:             owner.String(),
			BidAmount:          sdk.NewInt64Coin("uguu", 1000000),
			DepositAmount:      sdk.NewInt64Coin("uguu", 100000),
			PaidAmount:         sdk.NewInt64Coin("uguu", 0),
			BiddingPeriod:      time.Now().Add(time.Hour * 24),
			DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
			AutomaticPayment:   true,
			BidTime:            now,
			InterestAmount:     sdk.NewInt64Coin("uguu", 0),
			Borrowings:         []types.Borrowing{},
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "1",
			},
			Bidder:             owner2.String(),
			BidAmount:          sdk.NewInt64Coin("uguu", 1000000),
			DepositAmount:      sdk.NewInt64Coin("uguu", 100000),
			PaidAmount:         sdk.NewInt64Coin("uguu", 0),
			BiddingPeriod:      time.Now().Add(time.Hour * 24),
			DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
			AutomaticPayment:   true,
			BidTime:            now,
			InterestAmount:     sdk.NewInt64Coin("uguu", 0),
			Borrowings:         []types.Borrowing{},
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "2",
			},
			Bidder:             owner.String(),
			BidAmount:          sdk.NewInt64Coin("uguu", 1000000),
			DepositAmount:      sdk.NewInt64Coin("uguu", 100000),
			PaidAmount:         sdk.NewInt64Coin("uguu", 0),
			BiddingPeriod:      time.Now().Add(time.Hour * 24),
			DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
			AutomaticPayment:   true,
			BidTime:            now.Add(time.Second),
			InterestAmount:     sdk.NewInt64Coin("uguu", 0),
			Borrowings:         []types.Borrowing{},
		},
	}

	for _, bid := range cancelledBids {
		suite.app.NftmarketKeeper.SetCancelledBid(suite.ctx, bid)
	}

	// check matured cancelled bids
	maturedCancelledBids := suite.app.NftmarketKeeper.GetMaturedCancelledBids(suite.ctx, now.Add(time.Second))
	suite.Require().Len(maturedCancelledBids, 2)

	// allocate tokens to the module
	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second))
	coin := sdk.NewInt64Coin("uguu", int64(1000000000))
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{coin})
	suite.NoError(err)

	// execute matured cancelled bids
	err = suite.app.NftmarketKeeper.HandleMaturedCancelledBids(suite.ctx)
	suite.Require().NoError(err)

	// check matured cancelled bids after handle
	maturedCancelledBids = suite.app.NftmarketKeeper.GetMaturedCancelledBids(suite.ctx, now.Add(time.Second))
	suite.Require().Len(maturedCancelledBids, 0)
}
