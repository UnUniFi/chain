package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	ununifitypes "github.com/UnUniFi/chain/deprecated/types"
	ecoincentivetypes "github.com/UnUniFi/chain/x/ecosystemincentive/types"
	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

// test basic functions of nft listing
func (suite *KeeperTestSuite) TestNftListingBasics() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	now := time.Now().UTC()
	future := now.Add(time.Second)
	listings := []types.NftListing{
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "1",
			},
			Owner:              owner.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_LISTING,
			BidToken:           "uguu",
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			StartedAt:          now,
			EndAt:              now,
			FullPaymentEndAt:   time.Time{},
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
			CollectedAmount: sdk.Coin{
				Denom:  "uguu",
				Amount: sdk.ZeroInt(),
			},
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "2",
			},
			Owner:              owner.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_BIDDING,
			BidToken:           "uguu",
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			StartedAt:          now,
			EndAt:              now,
			FullPaymentEndAt:   time.Time{},
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
			CollectedAmount: sdk.Coin{
				Denom:  "uguu",
				Amount: sdk.ZeroInt(),
			},
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "3",
			},
			Owner:              owner.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_END_LISTING,
			BidToken:           "uguu",
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			StartedAt:          now,
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
			CollectedAmount: sdk.Coin{
				Denom:  "uguu",
				Amount: sdk.ZeroInt(),
			},
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "4",
			},
			Owner:              owner.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_SELLING_DECISION,
			BidToken:           "uguu",
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			StartedAt:          time.Time{},
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
			CollectedAmount: sdk.Coin{
				Denom:  "uguu",
				Amount: sdk.ZeroInt(),
			},
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "2",
				NftId:   "1",
			},
			Owner:              owner2.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_SUCCESSFUL_BID,
			BidToken:           "uguu",
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			StartedAt:          time.Time{},
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: now,
			AutoRelistedCount:  0,
			CollectedAmount: sdk.Coin{
				Denom:  "uguu",
				Amount: sdk.ZeroInt(),
			},
		},
		// {
		// 	NftId: types.NftIdentifier{
		// 		ClassId: "2",
		// 		NftId:   "2",
		// 	},
		// 	Owner:              owner2.String(),
		// 	ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
		// 	State:              types.ListingState_LIQUIDATION,
		// 	BidToken:           "uguu",
		// 	MinimumDepositRate: "0.1",
		// 	StartedAt:          time.Time{},
		// 	EndAt:              now,
		// 	FullPaymentEndAt:   now,
		// 	SuccessfulBidEndAt: now,
		// 	AutoRelistedCount:  0,
		// },
	}

	for _, listing := range listings {
		suite.app.NftmarketKeeper.SetNftListing(suite.ctx, listing)
	}

	for _, listing := range listings {
		gotListing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, listing.IdBytes())
		suite.Require().NoError(err)
		suite.Require().Equal(listing, gotListing)
	}

	// check all listings
	allListings := suite.app.NftmarketKeeper.GetAllNftListings(suite.ctx)
	suite.Require().Len(allListings, len(listings))

	// check listing by owner
	listingsByOwner := suite.app.NftmarketKeeper.GetListingsByOwner(suite.ctx, owner)
	suite.Require().Len(listingsByOwner, 4)

	// check active listings (bidding or listing status ending now)
	activeNftListings := suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, now)
	suite.Require().Len(activeNftListings, 0)
	// check active listings (bidding or listing status ending future time)
	activeNftListings = suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(activeNftListings, 2)

	// full payment listings (sell decision or end listing status ending now)
	fullPaymentNftListingsEnding := suite.app.NftmarketKeeper.GetFullPaymentNftListingsEndingAt(suite.ctx, now)
	suite.Require().Len(fullPaymentNftListingsEnding, 0)
	// full payment listings (sell decision or end listing status ending future)
	fullPaymentNftListingsEnding = suite.app.NftmarketKeeper.GetFullPaymentNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(fullPaymentNftListingsEnding, 2)

	// successful listing endings (ending now)
	successfulNftListingsEnding := suite.app.NftmarketKeeper.GetSuccessfulBidNftListingsEndingAt(suite.ctx, now)
	suite.Require().Len(successfulNftListingsEnding, 0)
	// successful listing endings (ending future)
	successfulNftListingsEnding = suite.app.NftmarketKeeper.GetSuccessfulBidNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(successfulNftListingsEnding, 1)

	// delete all the listings
	for _, listing := range listings {
		suite.app.NftmarketKeeper.DeleteNftListing(suite.ctx, listing)
	}

	// check queries after deleting all
	allListings = suite.app.NftmarketKeeper.GetAllNftListings(suite.ctx)
	suite.Require().Len(allListings, 0)

	// listings by owner
	listingsByOwner = suite.app.NftmarketKeeper.GetListingsByOwner(suite.ctx, owner)
	suite.Require().Len(listingsByOwner, 0)

	// queries for active listings
	activeNftListings = suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(activeNftListings, 0)
	// queries for full payment queue
	fullPaymentNftListingsEnding = suite.app.NftmarketKeeper.GetFullPaymentNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(fullPaymentNftListingsEnding, 0)
	// queries for successful ending queue
	successfulNftListingsEnding = suite.app.NftmarketKeeper.GetSuccessfulBidNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(successfulNftListingsEnding, 0)
}

func (suite *KeeperTestSuite) TestListNft() {
	acc1 := suite.addrs[0]
	acc2 := suite.addrs[1]
	keeper := suite.keeper
	nftKeeper := suite.nftKeeper

	tests := []struct {
		testCase         string
		classId          string
		nftId            string
		nftOwner         sdk.AccAddress
		lister           sdk.AccAddress
		bidToken         string
		mintBefore       bool
		listBefore       bool
		expectPass       bool
		statusListedHook bool
	}{
		{
			testCase:         "not existing nft",
			classId:          "class1",
			nftId:            "nft1",
			nftOwner:         acc1,
			lister:           acc1,
			bidToken:         "uguu",
			mintBefore:       false,
			listBefore:       false,
			expectPass:       false,
			statusListedHook: false,
		},
		{
			testCase:         "already listed",
			classId:          "class2",
			nftId:            "nft2",
			nftOwner:         acc1,
			lister:           acc1,
			bidToken:         "uguu",
			mintBefore:       true,
			listBefore:       true,
			expectPass:       false,
			statusListedHook: false,
		},
		{
			testCase:         "not owned nft",
			classId:          "class3",
			nftId:            "nft3",
			nftOwner:         acc1,
			lister:           acc2,
			bidToken:         "uguu",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       false,
			statusListedHook: false,
		},
		{
			testCase:         "unsupported bid token",
			classId:          "class4",
			nftId:            "nft4",
			nftOwner:         acc1,
			lister:           acc1,
			bidToken:         "xxxx",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       false,
			statusListedHook: false,
		},
		{
			testCase:         "successful listing with default active rank",
			classId:          "class5",
			nftId:            "nft5",
			nftOwner:         acc1,
			lister:           acc1,
			bidToken:         "uguu",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       true,
			statusListedHook: false,
		},
		{
			testCase:         "successful listing with non-default active rank",
			classId:          "class6",
			nftId:            "nft6",
			nftOwner:         acc1,
			lister:           acc1,
			bidToken:         "uguu",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       true,
			statusListedHook: false,
		},
		{
			testCase:         "successful anther owner",
			classId:          "class7",
			nftId:            "nft7",
			nftOwner:         acc2,
			lister:           acc2,
			bidToken:         "uguu",
			mintBefore:       true,
			listBefore:       false,
			expectPass:       true,
			statusListedHook: false,
		},
	}

	for _, tc := range tests {
		statusAfterNftListed = false
		if tc.mintBefore {
			nftKeeper.SaveClass(suite.ctx, nfttypes.Class{
				Id:          tc.classId,
				Name:        tc.classId,
				Symbol:      tc.classId,
				Description: tc.classId,
				Uri:         tc.classId,
			})
			err := nftKeeper.Mint(suite.ctx, nfttypes.NFT{
				ClassId: tc.classId,
				Id:      tc.nftId,
				Uri:     tc.nftId,
				UriHash: tc.nftId,
			}, tc.nftOwner)
			suite.Require().NoError(err)
		}
		if tc.listBefore {
			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:             ununifitypes.StringAccAddress(tc.lister),
				NftId:              types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId},
				ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:           tc.bidToken,
				MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
			})
			suite.Require().NoError(err)
		}
		err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             ununifitypes.StringAccAddress(tc.lister),
			NftId:              types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId},
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			BidToken:           tc.bidToken,
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			// params := keeper.GetParamSet(suite.ctx)
			// get listing
			listing, err := keeper.GetNftListingByIdBytes(suite.ctx, (types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}).IdBytes())
			suite.Require().NoError(err)

			// check ownership is transferred
			moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
			nftOwner := nftKeeper.GetOwner(suite.ctx, tc.classId, tc.nftId)
			suite.Require().Equal(nftOwner.String(), moduleAddr.String())

			// check startedAt is set as current time
			suite.Require().Equal(suite.ctx.BlockTime(), listing.StartedAt)

			// check endAt is set from initial listing duration
			// suite.Require().Equal(suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingPeriodInitial)), listing.EndAt, tc.testCase)
		} else {
			suite.Require().Error(err)
		}

		suite.Require().Equal(tc.statusListedHook, statusAfterNftListed, tc.testCase)
	}
}

func (suite *KeeperTestSuite) TestCancelNftListing() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	keeper := suite.keeper
	nftKeeper := suite.nftKeeper

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase           string
		classId            string
		nftId              string
		nftOwner           sdk.AccAddress
		canceller          sdk.AccAddress
		cancelAfter        time.Duration
		numBids            int
		listBefore         bool
		endedListing       bool
		expectPass         bool
		statusUnlistedHook bool
	}{
		{
			testCase:           "not existing listing",
			classId:            "class1",
			nftId:              "nft1",
			nftOwner:           acc1,
			canceller:          acc1,
			cancelAfter:        time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:            0,
			listBefore:         false,
			endedListing:       false,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "not owned nft listing",
			classId:            "class2",
			nftId:              "nft2",
			nftOwner:           acc1,
			canceller:          acc2,
			cancelAfter:        time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:            0,
			listBefore:         true,
			endedListing:       false,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "cancel time not pass",
			classId:            "class3",
			nftId:              "nft3",
			nftOwner:           acc1,
			canceller:          acc1,
			cancelAfter:        0,
			numBids:            0,
			listBefore:         true,
			endedListing:       false,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "already ended listing",
			classId:            "class4",
			nftId:              "nft4",
			nftOwner:           acc1,
			canceller:          acc1,
			cancelAfter:        time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:            0,
			listBefore:         true,
			endedListing:       true,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "successful cancel without cancel fee",
			classId:            "class5",
			nftId:              "nft5",
			nftOwner:           acc1,
			canceller:          acc1,
			cancelAfter:        time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:            0,
			listBefore:         true,
			endedListing:       false,
			expectPass:         true,
			statusUnlistedHook: false,
		},
		{
			testCase:           "successful cancel with cancel fee",
			classId:            "class6",
			nftId:              "nft6",
			nftOwner:           acc1,
			canceller:          acc1,
			cancelAfter:        time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:            0,
			listBefore:         true,
			endedListing:       false,
			expectPass:         true,
			statusUnlistedHook: false,
		},
	}

	for _, tc := range tests {
		statusAfterNftUnlistedWithoutPayment = false

		nftKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err := nftKeeper.Mint(suite.ctx, nfttypes.NFT{
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

		for i := 0; i < tc.numBids; i++ {
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
			suite.Require().NoError(err)
		}

		if tc.endedListing {
			err := suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
				Sender: ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:  nftIdentifier,
			})
			suite.Require().NoError(err)
		}

		oldCancellerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.canceller, "uguu")
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(tc.cancelAfter))
		err = keeper.CancelNftListing(suite.ctx, &types.MsgCancelNftListing{
			Sender: ununifitypes.StringAccAddress(tc.canceller),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			// check all bids are closed and returned
			nftBids := keeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().Len(nftBids, 0)

			// check cancel fee is reduced from listing owner
			if tc.numBids > 0 {
				cancellerNewBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.canceller, "uguu")
				suite.Require().True(cancellerNewBalance.Amount.LT(oldCancellerBalance.Amount))
			}

			// check nft ownership is returned back to owner
			owner := nftKeeper.GetOwner(suite.ctx, tc.classId, tc.nftId)
			suite.Require().Equal(owner, tc.nftOwner)

			// check nft listing is deleted
			_, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().Error(err)
		} else {
			suite.Require().Error(err)
		}

		suite.Require().Equal(tc.statusUnlistedHook, statusAfterNftUnlistedWithoutPayment, tc.testCase)
	}
}

// func (suite *KeeperTestSuite) TestExpandListingPeriod() {
// 	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

// 	// params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

// 	tests := []struct {
// 		testCase     string
// 		classId      string
// 		nftId        string
// 		nftOwner     sdk.AccAddress
// 		executor     sdk.AccAddress
// 		numBids      int
// 		listBefore   bool
// 		endedListing bool
// 		expectPass   bool
// 	}{
// 		{
// 			testCase:     "not existing listing",
// 			classId:      "class1",
// 			nftId:        "nft1",
// 			nftOwner:     acc1,
// 			executor:     acc1,
// 			numBids:      0,
// 			listBefore:   false,
// 			endedListing: false,
// 			expectPass:   false,
// 		},
// 		{
// 			testCase:     "not owned nft listing",
// 			classId:      "class2",
// 			nftId:        "nft2",
// 			nftOwner:     acc1,
// 			executor:     acc2,
// 			numBids:      0,
// 			listBefore:   true,
// 			endedListing: false,
// 			expectPass:   false,
// 		},
// 		{
// 			testCase:     "already ended nft listing",
// 			classId:      "class3",
// 			nftId:        "nft3",
// 			nftOwner:     acc1,
// 			executor:     acc2,
// 			numBids:      1,
// 			listBefore:   true,
// 			endedListing: true,
// 			expectPass:   false,
// 		},
// 		{
// 			testCase:     "successful nft listing period extend",
// 			classId:      "class4",
// 			nftId:        "nft4",
// 			nftOwner:     acc1,
// 			executor:     acc1,
// 			numBids:      1,
// 			listBefore:   true,
// 			endedListing: false,
// 			expectPass:   true,
// 		},
// 	}

// 	for _, tc := range tests {

// 		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
// 		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
// 		suite.NoError(err)
// 		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.executor, sdk.Coins{coin})
// 		suite.NoError(err)

// 		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
// 			Id:          tc.classId,
// 			Name:        tc.classId,
// 			Symbol:      tc.classId,
// 			Description: tc.classId,
// 			Uri:         tc.classId,
// 		})
// 		err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
// 			ClassId: tc.classId,
// 			Id:      tc.nftId,
// 			Uri:     tc.nftId,
// 			UriHash: tc.nftId,
// 		}, tc.nftOwner)
// 		suite.Require().NoError(err)

// 		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
// 		if tc.listBefore {
// 			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
// 				Sender:             ununifitypes.StringAccAddress(tc.nftOwner),
// 				NftId:              nftIdentifier,
// 				ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
// 				BidToken:           "uguu",
// 				MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
// 			})
// 			suite.Require().NoError(err)
// 		}

// 		lastBidder := sdk.AccAddress{}
// 		for i := 0; i < tc.numBids; i++ {
// 			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 			lastBidder = bidder

// 			// init tokens to addr
// 			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
// 			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
// 			suite.NoError(err)
// 			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{coin})
// 			suite.NoError(err)

// 			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
// 				Sender:           ununifitypes.StringAccAddress(bidder),
// 				NftId:            nftIdentifier,
// 				BidAmount:        coin,
// 				AutomaticPayment: false,
// 			})
// 			suite.Require().NoError(err)
// 		}

// 		if tc.endedListing {
// 			err := suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
// 				Sender: ununifitypes.StringAccAddress(tc.nftOwner),
// 				NftId:  nftIdentifier,
// 			})
// 			suite.Require().NoError(err)
// 		}

// 		oldListing, _ := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())

// 		oldExecutorBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.executor, "uguu")
// 		oldLastBidderBalance := suite.app.BankKeeper.GetBalance(suite.ctx, lastBidder, "uguu")
// 		err = suite.app.NftmarketKeeper.ExpandListingPeriod(suite.ctx, &types.MsgExpandListingPeriod{
// 			Sender: ununifitypes.StringAccAddress(tc.executor),
// 			NftId:  nftIdentifier,
// 		})

// 		if tc.expectPass {
// 			suite.Require().NoError(err)

// 			// check fee is paid
// 			newExecutorBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.executor, "uguu")
// 			suite.Require().True(newExecutorBalance.Amount.LT(oldExecutorBalance.Amount))

// 			// check winner bidders get extend fee
// 			newLastBidderBalance := suite.app.BankKeeper.GetBalance(suite.ctx, lastBidder, "uguu")
// 			suite.Require().True(newLastBidderBalance.Amount.GT(oldLastBidderBalance.Amount))

// 			// check listing endAt is extended
// 			listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
// 			suite.Require().NoError(err)
// 			suite.Require().True(oldListing.EndAt.Before(listing.EndAt))
// 		} else {
// 			suite.Require().Error(err)
// 		}
// 	}
// }

func (suite *KeeperTestSuite) TestSellingDecision() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase      string
		classId       string
		nftId         string
		nftOwner      sdk.AccAddress
		executor      sdk.AccAddress
		numBids       int
		enoughAutoPay bool
		autoPayment   bool
		listBefore    bool
		endedListing  bool
		expectPass    bool
	}{
		{
			testCase:      "not existing listing",
			classId:       "class1",
			nftId:         "nft1",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       0,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    false,
			endedListing:  false,
			expectPass:    false,
		},
		{
			testCase:      "not owned nft listing",
			classId:       "class2",
			nftId:         "nft2",
			nftOwner:      acc1,
			executor:      acc2,
			numBids:       0,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  false,
			expectPass:    false,
		},
		{
			testCase:      "already ended nft listing",
			classId:       "class3",
			nftId:         "nft3",
			nftOwner:      acc1,
			executor:      acc2,
			numBids:       1,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  true,
			expectPass:    false,
		},
		{
			testCase:      "successful nft selling decision with automatic payment",
			classId:       "class4",
			nftId:         "nft4",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: true,
			autoPayment:   true,
			listBefore:    true,
			endedListing:  false,
			expectPass:    true,
		},
		{
			testCase:      "successful nft selling decision with automatic payment enabled with not enough balance",
			classId:       "class5",
			nftId:         "nft5",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: false,
			autoPayment:   true,
			listBefore:    true,
			endedListing:  false,
			expectPass:    true,
		},
		{
			testCase:      "successful nft selling decision without automatic payment",
			classId:       "class6",
			nftId:         "nft6",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  false,
			expectPass:    true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.executor, sdk.Coins{coin})
		suite.NoError(err)

		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
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

		lastBidder := sdk.AccAddress{}
		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
			lastBidder = bidder

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			halfCoin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)/2))
			mintCoin := coin
			if !tc.enoughAutoPay {
				mintCoin = sdk.NewInt64Coin("uguu", int64(1000000*(i+1)/2))
			}
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{mintCoin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{mintCoin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:             ununifitypes.StringAccAddress(bidder),
				NftId:              nftIdentifier,
				BidAmount:          coin,
				DepositAmount:      halfCoin,
				AutomaticPayment:   tc.autoPayment,
				DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
				BiddingPeriod:      time.Now().Add(time.Hour * 24),
			})
			suite.Require().NoError(err)
		}

		if tc.endedListing {
			err := suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
				Sender: ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:  nftIdentifier,
			})
			suite.Require().NoError(err)
		}

		err = suite.app.NftmarketKeeper.SellingDecision(suite.ctx, &types.MsgSellingDecision{
			Sender: ununifitypes.StringAccAddress(tc.executor),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)
			if tc.autoPayment {
				bid, err := suite.app.NftmarketKeeper.GetBid(suite.ctx, nftIdentifier.IdBytes(), lastBidder)
				suite.Require().NoError(err)
				if tc.enoughAutoPay {
					// check automatic payment execution when user has enough balance
					suite.Require().Equal(bid.PaidAmount.Amount.Add(bid.DepositAmount.Amount), bid.BidAmount.Amount, tc.testCase)
				} else {
					// check automatic payment when the user does not have enough balance
					suite.Require().NotEqual(bid.PaidAmount, bid.BidAmount.Amount)
				}
			}

			// check full payment end time update
			listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().NoError(err)
			suite.Require().Equal(listing.State, types.ListingState_SELLING_DECISION)
			suite.Require().Equal(suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingFullPaymentPeriod)), listing.FullPaymentEndAt)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestEndNftListing() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase           string
		classId            string
		nftId              string
		nftOwner           sdk.AccAddress
		executor           sdk.AccAddress
		numBids            int
		enoughAutoPay      bool
		autoPayment        bool
		listBefore         bool
		endedListing       bool
		expectPass         bool
		statusUnlistedHook bool
	}{
		{
			testCase:           "not existing listing",
			classId:            "class1",
			nftId:              "nft1",
			nftOwner:           acc1,
			executor:           acc1,
			numBids:            0,
			enoughAutoPay:      true,
			autoPayment:        false,
			listBefore:         false,
			endedListing:       false,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "not owned nft listing",
			classId:            "class2",
			nftId:              "nft2",
			nftOwner:           acc1,
			executor:           acc2,
			numBids:            0,
			enoughAutoPay:      true,
			autoPayment:        false,
			listBefore:         true,
			endedListing:       false,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "already ended nft listing",
			classId:            "class3",
			nftId:              "nft3",
			nftOwner:           acc1,
			executor:           acc2,
			numBids:            1,
			enoughAutoPay:      true,
			autoPayment:        false,
			listBefore:         true,
			endedListing:       true,
			expectPass:         false,
			statusUnlistedHook: false,
		},
		{
			testCase:           "successful nft listing ending when no bid",
			classId:            "class4",
			nftId:              "nft4",
			nftOwner:           acc1,
			executor:           acc1,
			numBids:            0,
			enoughAutoPay:      true,
			autoPayment:        false,
			listBefore:         true,
			endedListing:       false,
			expectPass:         true,
			statusUnlistedHook: true,
		},
		{
			testCase:           "successful nft listing ending with automatic payment enabled with not enough balance",
			classId:            "class5",
			nftId:              "nft5",
			nftOwner:           acc1,
			executor:           acc1,
			numBids:            1,
			enoughAutoPay:      false,
			autoPayment:        true,
			listBefore:         true,
			endedListing:       false,
			expectPass:         true,
			statusUnlistedHook: false,
		},
		{
			testCase:           "successful nft listing ending without automatic payment",
			classId:            "class6",
			nftId:              "nft6",
			nftOwner:           acc1,
			executor:           acc1,
			numBids:            4,
			enoughAutoPay:      true,
			autoPayment:        false,
			listBefore:         true,
			endedListing:       false,
			expectPass:         true,
			statusUnlistedHook: false,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()
		keeper := suite.keeper
		nftKeeper := suite.nftKeeper
		statusAfterNftUnlistedWithoutPayment = false

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.executor, sdk.Coins{coin})
		suite.NoError(err)

		nftKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = nftKeeper.Mint(suite.ctx, nfttypes.NFT{
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

		lastBidder := sdk.AccAddress{}
		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
			lastBidder = bidder

			// init tokens to addr
			bidAmount := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			depositAmount := sdk.NewInt64Coin("uguu", int64(100000*(i+1)))
			mintCoin := bidAmount
			if !tc.enoughAutoPay {
				mintCoin = sdk.NewInt64Coin("uguu", int64(1000000*(i+1)/2))
			}
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{mintCoin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{mintCoin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:             ununifitypes.StringAccAddress(bidder),
				NftId:              nftIdentifier,
				BidAmount:          bidAmount,
				BiddingPeriod:      time.Now().Add(time.Hour * 24),
				DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
				AutomaticPayment:   tc.autoPayment,
				DepositAmount:      depositAmount,
			})
			suite.Require().NoError(err)
		}

		if tc.endedListing {
			err := keeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
				Sender: ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:  nftIdentifier,
			})
			suite.Require().NoError(err)
		}

		err = keeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
			Sender: ununifitypes.StringAccAddress(tc.executor),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)
			if tc.autoPayment {
				bid, err := keeper.GetBid(suite.ctx, nftIdentifier.IdBytes(), lastBidder)
				suite.Require().NoError(err)
				if tc.enoughAutoPay {
					// check automatic payment execution when user has enough balance
					suite.Require().Equal(bid.PaidAmount, bid.BidAmount.Amount)
				} else {
					// check automatic payment when the user does not have enough balance
					suite.Require().NotEqual(bid.PaidAmount, bid.BidAmount.Amount)
				}
			}

			if tc.numBids == 0 {
				// successful end when there's no bid - delete listing directly and transfer nft back to owner
				_, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().Error(err)
			} else {
				// check full payment end time update
				listing, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_END_LISTING)
				suite.Require().Equal(suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingFullPaymentPeriod)), listing.FullPaymentEndAt)

				// // check non-active bids are cancelled automatically
				// bids := suite.app.NftmarketKeeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
				// suite.Require().True(len(bids) <= int(listing.BidActiveRank))
			}

		} else {
			suite.Require().Error(err)
		}
		suite.Require().Equal(tc.statusUnlistedHook, statusAfterNftUnlistedWithoutPayment)
	}
}

func (suite *KeeperTestSuite) TestProcessEndingNftListings() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase            string
		classId             string
		nftId               string
		nftOwner            sdk.AccAddress
		numBids             int
		relistedCount       uint64
		expectedToEnd       bool
		expectedToBeRemoved bool
		statusUnlistedHook  bool
	}{
		{
			testCase:            "no bid nft listing extend when relisted count not reached the limit",
			classId:             "class1",
			nftId:               "nft1",
			nftOwner:            acc1,
			numBids:             0,
			relistedCount:       0,
			expectedToEnd:       false,
			expectedToBeRemoved: false,
			statusUnlistedHook:  false,
		},
		// {
		// 	testCase:            "no bid nft listing end when relisted count reached",
		// 	classId:             "class2",
		// 	nftId:               "nft2",
		// 	nftOwner:            acc1,
		// 	numBids:             0,
		// 	relistedCount:       params.AutoRelistingCountIfNoBid,
		// 	expectedToEnd:       true,
		// 	expectedToBeRemoved: true,
		// 	statusUnlistedHook:  true,
		// },
		// {
		// 	testCase:            "bids existing nft listing end when relisted count not reached",
		// 	classId:             "class3",
		// 	nftId:               "nft3",
		// 	nftOwner:            acc1,
		// 	numBids:             1,
		// 	relistedCount:       0,
		// 	expectedToEnd:       true,
		// 	expectedToBeRemoved: false,
		// 	statusUnlistedHook:  false,
		// },
		// {
		// 	testCase:            "bids existing nft listing end when relisted count reached",
		// 	classId:             "class4",
		// 	nftId:               "nft4",
		// 	nftOwner:            acc1,
		// 	numBids:             1,
		// 	relistedCount:       params.AutoRelistingCountIfNoBid,
		// 	expectedToEnd:       true,
		// 	expectedToBeRemoved: false,
		// 	statusUnlistedHook:  false,
		// },
	}

	for _, tc := range tests {
		suite.SetupTest()
		keeper := suite.keeper
		nftKeeper := suite.nftKeeper
		statusAfterNftUnlistedWithoutPayment = false

		now := time.Now().UTC()
		suite.ctx = suite.ctx.WithBlockTime(now)

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.nftOwner, sdk.Coins{coin})
		suite.NoError(err)

		nftKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = nftKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		err = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             ununifitypes.StringAccAddress(tc.nftOwner),
			NftId:              nftIdentifier,
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			BidToken:           "uguu",
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
		})
		suite.Require().NoError(err)
		listing, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
		suite.Require().NoError(err)
		listing.AutoRelistedCount = tc.relistedCount
		keeper.SetNftListing(suite.ctx, listing)

		for i := 0; i < tc.numBids; i++ {
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
				AutomaticPayment:   true,
				DepositAmount:      depositAmount,
			})
			suite.Require().NoError(err)
		}

		suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second * time.Duration(params.NftListingPeriodInitial+1)))
		keeper.ProcessEndingNftListings(suite.ctx)

		if tc.expectedToBeRemoved {
			_, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().Error(err, tc.testCase)
		} else {
			listing, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().NoError(err)

			if tc.expectedToEnd {
				suite.Require().Equal(listing.State, types.ListingState_END_LISTING, tc.testCase)
			} else {
				suite.Require().NotEqual(listing.State, types.ListingState_END_LISTING)
			}
		}
		suite.Require().Equal(tc.statusUnlistedHook, statusAfterNftUnlistedWithoutPayment)
	}
}

func (suite *KeeperTestSuite) TestActiveNftListingsEndingAtQueueRemovalOnNftListingEnd() {
	suite.SetupTest()

	classId := "class1"
	nftId := "nf1"
	nftOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()

	suite.ctx = suite.ctx.WithBlockTime(now)
	coin := sdk.NewInt64Coin("uguu", int64(1000000000))
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, nftOwner, sdk.Coins{coin})
	suite.NoError(err)

	suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
		Id:          classId,
		Name:        classId,
		Symbol:      classId,
		Description: classId,
		Uri:         classId,
	})
	err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
		ClassId: classId,
		Id:      nftId,
		Uri:     nftId,
		UriHash: nftId,
	}, nftOwner)
	suite.Require().NoError(err)

	nftIdentifier := types.NftIdentifier{ClassId: classId, NftId: nftId}
	err = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
		Sender:             ununifitypes.StringAccAddress(nftOwner),
		NftId:              nftIdentifier,
		ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
		BidToken:           "uguu",
		MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
	})
	suite.Require().NoError(err)

	listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().NoError(err)
	suite.Require().True(listing.IsActive())

	// check number before end listing
	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)
	activeNftListings := suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, now.Add(time.Second*time.Duration(params.NftListingPeriodInitial+1)))
	suite.Require().Len(activeNftListings, 1)

	err = suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
		Sender: ununifitypes.StringAccAddress(nftOwner),
		NftId:  nftIdentifier,
	})
	suite.Require().NoError(err)

	// check number after end listing
	activeNftListings = suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, now.Add(time.Second*time.Duration(params.NftListingPeriodInitial+1)))
	suite.Require().Len(activeNftListings, 0)
}

func (suite *KeeperTestSuite) TestHandleFullPaymentPeriodEndings() {

	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase           string
		classId            string
		nftId              string
		nftOwner           sdk.AccAddress
		numBids            int
		listingState       types.ListingState
		fullPay            bool
		statusUnlistedHook bool
	}{
		{
			testCase:           "selling decision listing when highest bid is paid",
			classId:            "class1",
			nftId:              "nft1",
			nftOwner:           acc1,
			numBids:            2,
			listingState:       types.ListingState_SELLING_DECISION,
			fullPay:            true,
			statusUnlistedHook: false,
		}, // add successful listing state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status
		{
			testCase:           "selling decision listing when highest bid is not paid and no more bids",
			classId:            "class2",
			nftId:              "nft2",
			nftOwner:           acc1,
			numBids:            1,
			listingState:       types.ListingState_SELLING_DECISION,
			fullPay:            false,
			statusUnlistedHook: false,
		}, // status => ListingState_LISTING
		{
			testCase:           "selling decision listing when highest bid is not paid, and more bids",
			classId:            "class2",
			nftId:              "nft2",
			nftOwner:           acc1,
			numBids:            2,
			listingState:       types.ListingState_SELLING_DECISION,
			fullPay:            false,
			statusUnlistedHook: false,
		}, // status => ListingState_BIDDING
		{
			testCase:           "ended listing, when fully paid bid exists",
			classId:            "class2",
			nftId:              "nft2",
			nftOwner:           acc1,
			numBids:            2,
			listingState:       types.ListingState_END_LISTING,
			fullPay:            true,
			statusUnlistedHook: false,
		}, // add successful bid state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status, close all the other bids
		{
			testCase:           "ended listing, when fully paid bid does not exist",
			classId:            "class2",
			nftId:              "nft2",
			nftOwner:           acc1,
			numBids:            2,
			listingState:       types.ListingState_END_LISTING,
			fullPay:            true,
			statusUnlistedHook: false,
		}, // all the bids closed, pay depositCollected, nft listing delete, transfer nft to fully paid bidder
	}

	for _, tc := range tests {
		suite.SetupTest()
		keeper := suite.keeper
		nftKeeper := suite.nftKeeper
		statusAfterNftUnlistedWithoutPayment = false

		now := time.Now().UTC()
		suite.ctx = suite.ctx.WithBlockTime(now)

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.nftOwner, sdk.Coins{coin})
		suite.NoError(err)

		nftKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = nftKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		err = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             ununifitypes.StringAccAddress(tc.nftOwner),
			NftId:              nftIdentifier,
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			BidToken:           "uguu",
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
		})
		suite.Require().NoError(err)
		listing, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
		suite.Require().NoError(err)

		for i := 0; i < tc.numBids; i++ {
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
				AutomaticPayment:   true,
				DepositAmount:      depositAmount,
			})
			suite.Require().NoError(err)

			if tc.fullPay {
				err := keeper.PayFullBid(suite.ctx, &types.MsgPayFullBid{
					Sender: ununifitypes.StringAccAddress(bidder),
					NftId:  nftIdentifier,
				})
				suite.Require().NoError(err)
			}
		}

		listing.State = tc.listingState
		keeper.SetNftListing(suite.ctx, listing)

		oldNftOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.nftOwner, "uguu")
		suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second * time.Duration(params.NftListingPeriodInitial+1)))
		keeper.HandleFullPaymentsPeriodEndings(suite.ctx)

		switch tc.listingState {
		case types.ListingState_SELLING_DECISION:
			if tc.fullPay {
				// add successful listing state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status
				listing, err = keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_SUCCESSFUL_BID)
				suite.Require().Equal(listing.SuccessfulBidEndAt, suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingNftDeliveryPeriod)))
			} else if tc.numBids > 1 {
				// status => ListingState_BIDDING
				listing, err = keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_BIDDING)
			} else {
				// status => ListingState_LISTING
				listing, err = keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_LISTING)
			}
		case types.ListingState_END_LISTING:
			if tc.fullPay {
				// add successful bid state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status, close all the other bids
				listing, err = keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_SUCCESSFUL_BID)
				suite.Require().Equal(listing.SuccessfulBidEndAt, suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingNftDeliveryPeriod)))
			} else {
				// all the bids closed, pay depositCollected, nft listing delete, transfer nft to fully paid bidder
				_, err = keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().Error(err)

				bids := keeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().Len(bids, 0)

				newOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.nftOwner, "uguu")
				suite.Require().True(newOwnerBalance.Amount.GT(oldNftOwnerBalance.Amount))

				nft, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(nft.Owner, tc.nftOwner)
			}
		}
		suite.Require().Equal(tc.statusUnlistedHook, statusAfterNftUnlistedWithoutPayment)
	}
}

func (suite *KeeperTestSuite) TestDelieverSuccessfulBids() {
	suite.SetupTest()
	keeper := suite.keeper

	classId := "class1"
	nftId := "nf1"
	nftOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()

	suite.ctx = suite.ctx.WithBlockTime(now)
	bidAmount := sdk.NewInt64Coin("uguu", int64(1000000000))
	depositAmount := sdk.NewInt64Coin("uguu", int64(100000000))
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, nftOwner, sdk.Coins{bidAmount})
	suite.NoError(err)

	suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
		Id:          classId,
		Name:        classId,
		Symbol:      classId,
		Description: classId,
		Uri:         classId,
	})
	err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
		ClassId: classId,
		Id:      nftId,
		Uri:     nftId,
		UriHash: nftId,
	}, nftOwner)
	suite.Require().NoError(err)

	nftIdentifier := types.NftIdentifier{ClassId: classId, NftId: nftId}
	err = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
		Sender:             ununifitypes.StringAccAddress(nftOwner),
		NftId:              nftIdentifier,
		ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
		BidToken:           "uguu",
		MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
	})
	suite.Require().NoError(err)

	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// init tokens to addr
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{bidAmount})
	suite.NoError(err)

	err = suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
		Sender:             ununifitypes.StringAccAddress(bidder),
		NftId:              nftIdentifier,
		BidAmount:          bidAmount,
		BiddingPeriod:      time.Now().Add(time.Hour * 24),
		DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
		AutomaticPayment:   true,
		DepositAmount:      depositAmount,
	})
	suite.Require().NoError(err)
	err = keeper.PayFullBid(suite.ctx, &types.MsgPayFullBid{
		Sender: ununifitypes.StringAccAddress(bidder),
		NftId:  nftIdentifier,
	})
	suite.Require().NoError(err)

	listing, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().NoError(err)
	listing.SuccessfulBidEndAt = now
	listing.State = types.ListingState_SUCCESSFUL_BID
	keeper.SetNftListing(suite.ctx, listing)

	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second))
	oldNftOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, nftOwner, "uguu")

	keeper.DelieverSuccessfulBids(suite.ctx)

	// check nft transfer
	newNftOwner := suite.app.NFTKeeper.GetOwner(suite.ctx, classId, nftId)
	suite.Require().NoError(err)
	suite.Require().Equal(newNftOwner.String(), bidder.String())

	// check fee paid
	newOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, nftOwner, "uguu")
	suite.Require().True(newOwnerBalance.Amount.GT(oldNftOwnerBalance.Amount))

	// check bid deleted
	bids := keeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Len(bids, 0)

	// check nft listing deleted
	_, err = keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Error(err)

	// check if AfterNftPaymentWithCommission is called
	suite.Require().True(statusAfterNftPaymentWithCommission)
}

func (suite *KeeperTestSuite) TestProcessPaymentWithCommissionFee() {
	tests := []struct {
		testCase   string
		loanAmount sdk.Int
	}{
		{
			testCase:   "zero loan",
			loanAmount: sdk.ZeroInt(),
		},
		{
			testCase:   "positive loan",
			loanAmount: sdk.NewInt(10),
		},
	}

	for _, tc := range tests {
		suite.SetupTest()
		keeper := suite.keeper
		statusAfterNftPaymentWithCommission = false

		amount := sdk.NewInt(1000000)
		owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewCoin("uguu", amount)})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{sdk.NewCoin("uguu", amount)})
		suite.NoError(err)

		var nftId types.NftIdentifier
		keeper.ProcessPaymentWithCommissionFee(suite.ctx, owner, "uguu", amount, tc.loanAmount, nftId)

		params := keeper.GetParamSet(suite.ctx)
		fee := amount.Mul(sdk.NewInt(int64(params.NftListingCommissionFee))).Quo(sdk.NewInt(100))
		listingPayment := amount.Sub(fee).Sub(tc.loanAmount)

		// check fee paid to NftTradingFee
		tradingFeeModuleAcc := suite.app.AccountKeeper.GetModuleAddress(ecoincentivetypes.ModuleName)
		tradingFeeBal := suite.app.BankKeeper.GetBalance(suite.ctx, tradingFeeModuleAcc, "uguu")
		suite.Require().Equal(tradingFeeBal, sdk.NewCoin("uguu", fee))

		// check fee to lister
		ownerBal := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "uguu")
		suite.Require().Equal(ownerBal, sdk.NewCoin("uguu", listingPayment))

		// check if AfterNftPaymentWithCommission is called
		suite.Require().True(statusAfterNftPaymentWithCommission)
	}
}

func (suite *KeeperTestSuite) TestDelieverSuccessfulBidForPositiveLoan() {
	suite.SetupTest()
	keeper := suite.keeper

	classId := "class1"
	nftId := "nf1"
	nftOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()

	suite.ctx = suite.ctx.WithBlockTime(now)
	bidAmount := sdk.NewInt64Coin("uguu", int64(1000000000))
	depositAmount := sdk.NewInt64Coin("uguu", int64(100000000))
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, nftOwner, sdk.Coins{bidAmount})
	suite.NoError(err)

	suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
		Id:          classId,
		Name:        classId,
		Symbol:      classId,
		Description: classId,
		Uri:         classId,
	})
	err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
		ClassId: classId,
		Id:      nftId,
		Uri:     nftId,
		UriHash: nftId,
	}, nftOwner)
	suite.Require().NoError(err)

	nftIdentifier := types.NftIdentifier{ClassId: classId, NftId: nftId}
	err = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
		Sender:             ununifitypes.StringAccAddress(nftOwner),
		NftId:              nftIdentifier,
		ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
		BidToken:           "uguu",
		MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
	})
	suite.Require().NoError(err)

	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// init tokens to addr
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{bidAmount})
	suite.NoError(err)

	err = suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
		Sender:             ununifitypes.StringAccAddress(bidder),
		NftId:              nftIdentifier,
		BidAmount:          bidAmount,
		BiddingPeriod:      time.Now().Add(time.Hour * 24),
		DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
		AutomaticPayment:   true,
		DepositAmount:      depositAmount,
	})
	suite.Require().NoError(err)
	err = keeper.PayFullBid(suite.ctx, &types.MsgPayFullBid{
		Sender: ununifitypes.StringAccAddress(bidder),
		NftId:  nftIdentifier,
	})
	suite.Require().NoError(err)

	listing, err := keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().NoError(err)
	listing.SuccessfulBidEndAt = now
	listing.State = types.ListingState_SUCCESSFUL_BID
	keeper.SetNftListing(suite.ctx, listing)

	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second))

	loanAmount := sdk.NewInt64Coin("uguu", 10)
	params := keeper.GetParamSet(suite.ctx)
	commissionFee := params.NftListingCommissionFee
	feeAmount := bidAmount.Amount.Mul(sdk.NewInt(int64(commissionFee))).Quo(sdk.NewInt(100))
	err = suite.app.NftmarketKeeper.Borrow(suite.ctx, &types.MsgBorrow{
		Sender: nftOwner.Bytes(),
		NftId:  nftIdentifier,
		Amount: loanAmount,
	})
	suite.Require().NoError(err)

	oldNftOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, nftOwner, "uguu")
	suite.NotPanics(func() {
		keeper.DelieverSuccessfulBids(suite.ctx)
	})

	// check nft transfer
	newNftOwner := suite.app.NFTKeeper.GetOwner(suite.ctx, classId, nftId)
	suite.Require().NoError(err)
	suite.Require().Equal(newNftOwner.String(), bidder.String())

	// check fee paid
	newOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, nftOwner, "uguu")
	suite.Require().True(newOwnerBalance.Amount.Equal(oldNftOwnerBalance.Amount.Add(bidAmount.Amount).Sub(feeAmount).Sub(loanAmount.Amount)))

	// check bid deleted
	bids := keeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Len(bids, 0)

	// check nft listing deleted
	_, err = keeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Error(err)

	suite.Require().True(statusAfterNftPaymentWithCommission)
}

// TestReListingDataManagement is a test to check if the all data related listing like bid and loan
// works as intended when the list is re-listed by system or even by user's hand
func (suite *KeeperTestSuite) TestReListingDataManagement() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)
	nftIdentifier := types.NftIdentifier{ClassId: "class1", NftId: "nft1"}
	nftOwner := acc1

	tests := []struct {
		testCase     string
		numBids      int
		listingState types.ListingState
		fullPay      bool
		borrow       bool
	}{
		{
			testCase:     "selling decision listing when highest bid is  paid and no more bids",
			numBids:      1,
			listingState: types.ListingState_SELLING_DECISION,
			fullPay:      true,
			borrow:       true,
		}, // status => ListingState_LISTING
		{
			testCase:     "selling decision listing when highest bid is not paid and no more bids",
			numBids:      1,
			listingState: types.ListingState_SELLING_DECISION,
			borrow:       true,
		}, // status => ListingState_LISTING
		{
			testCase:     "selling decision listing when highest bid is not paid, and more bids",
			numBids:      2,
			listingState: types.ListingState_SELLING_DECISION,
			borrow:       true,
		}, // status => ListingState_BIDDING
		{
			testCase:     "end listing when no bids is not paid",
			numBids:      1,
			listingState: types.ListingState_END_LISTING,
			borrow:       true,
		}, // status => delist
		{
			testCase:     "end listing when no bids is not paid",
			numBids:      1,
			listingState: types.ListingState_END_LISTING,
			fullPay:      true,
			borrow:       true,
		}, // status => delist
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now().UTC()
		suite.ctx = suite.ctx.WithBlockTime(now)

		bidAmount := sdk.NewInt64Coin("uguu", int64(1000000000))
		depositAmount := sdk.NewInt64Coin("uguu", int64(100000000))
		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
		_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, nftOwner, sdk.Coins{bidAmount})

		_ = suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:     nftIdentifier.ClassId,
			Name:   nftIdentifier.ClassId,
			Symbol: nftIdentifier.ClassId,
		})
		_ = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: nftIdentifier.ClassId,
			Id:      nftIdentifier.NftId,
		}, nftOwner)

		_ = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:             ununifitypes.StringAccAddress(nftOwner),
			NftId:              nftIdentifier,
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			BidToken:           "uguu",
			MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
		})
		listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
		suite.Require().NoError(err)

		for i := 0; i < tc.numBids; i++ {
			if tc.borrow {
				suite.PlaceAndBorrow(bidAmount, depositAmount, nftIdentifier, nftOwner, tc.fullPay, 2)
			} else {
				bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
				// init tokens to addr
				bidAmount := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
				depositAmount := sdk.NewInt64Coin("uguu", int64(100000*(i+1)))
				_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{bidAmount})
				_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{bidAmount})

				_ = suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
					Sender:             ununifitypes.StringAccAddress(bidder),
					NftId:              nftIdentifier,
					BidAmount:          bidAmount,
					BiddingPeriod:      time.Now().Add(time.Hour * 24),
					DepositLendingRate: sdk.MustNewDecFromStr("0.05"),
					AutomaticPayment:   false,
					DepositAmount:      depositAmount,
				})
			}
		}

		listing.State = tc.listingState
		suite.app.NftmarketKeeper.SetNftListing(suite.ctx, listing)

		suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second * time.Duration(params.NftListingPeriodInitial+1)))
		suite.app.NftmarketKeeper.HandleFullPaymentsPeriodEndings(suite.ctx)
		bids := suite.app.NftmarketKeeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
		loan := suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())

		// check bids and loan data specifically
		switch tc.listingState {
		case types.ListingState_SELLING_DECISION:
			if tc.numBids > 1 {
				// status => ListingState_BIDDING
				listing, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_BIDDING)
				suite.Require().NotEmpty(bids)
				suite.Require().NotEmpty(loan)
			} else if tc.fullPay {
				// status => ListingState_SUCCESSFUL_BID
				// this simulates the state of re-lisitng from the situation the list is successfully sent to
				// a buyer and buyer lists it
				listing, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_SUCCESSFUL_BID)
				suite.Require().Len(bids, 1)
				suite.Require().Empty(loan)
			} else {
				// status => ListingState_LISTING
				listing, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_LISTING)
				suite.Require().Empty(bids)
				suite.Require().Empty(loan)
			}
		case types.ListingState_END_LISTING:
			if tc.fullPay {
				// pay depositCollected, nft listing delete, transfer nft to fully paid bidder
				// this simulates the state of re-listing from the situation the list was handed to buyer
				// and that buyer lists
				listing, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_SUCCESSFUL_BID)
				suite.Require().Len(bids, 1)
				suite.Require().Empty(loan)
			} else {
				// all the bids closed, pay depositCollected, nft listing delete
				// this simulates the state of re-listing from the situation the list was handed to buyer
				_, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().Error(err)
				suite.Require().Empty(bids)
				suite.Require().Empty(loan)
			}
		}
	}
}

// test LiquidationProcessExitsWinner
func (suite *KeeperTestSuite) TestLiquidationProcessExitsWinner() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	bidder1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now()
	refundFCheckerGen := func(testCase string, expRefundBids types.NftBids, expTotalInterest, expSurplusAmount sdk.Coin, expListing types.NftListing) func(ctx sdk.Context, refundBids types.NftBids, totalInterest, surplusAmount sdk.Coin, listing types.NftListing) error {
		return func(ctx sdk.Context, refundBids types.NftBids, totalInterest, surplusAmount sdk.Coin, listing types.NftListing) error {
			suite.Equal(expRefundBids, refundBids, testCase)
			suite.Equal(expTotalInterest, totalInterest, testCase)
			suite.Equal(expSurplusAmount, surplusAmount, testCase)
			suite.Equal(expListing, listing, testCase)
			return nil
		}
	}
	type funcArg struct {
		collectBid types.NftBids
		refundBid  types.NftBids
		listing    types.NftListing
		winnerBid  types.NftBid
		blockTime  time.Time
	}
	type funcFExp struct {
		expRefundBid                       types.NftBids
		expTotalInterest, expSurplusAmount sdk.Coin
		expListing                         types.NftListing
	}
	tcs := []struct {
		testCase  string
		funcArg   funcArg
		funcFExp  funcFExp
		expResult error
	}{
		{
			"no refund bid and no collect bid",
			funcArg{
				collectBid: types.NftBids{},
				refundBid:  types.NftBids{},
				listing: types.NftListing{
					NftId: types.NftIdentifier{
						ClassId: "1",
						NftId:   "1",
					},
					Owner:              owner.String(),
					ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
					State:              types.ListingState_LISTING,
					BidToken:           "uguu",
					MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
					StartedAt:          now,
					EndAt:              now,
					FullPaymentEndAt:   time.Time{},
					SuccessfulBidEndAt: time.Time{},
					AutoRelistedCount:  0,
					CollectedAmount: sdk.Coin{
						Denom:  "uguu",
						Amount: sdk.ZeroInt(),
					},
				},
				winnerBid: types.NftBid{
					NftId: types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder:             "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
					DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(50)),
					PaidAmount:         sdk.NewCoin("uguu", sdk.NewInt(0)),
					BiddingPeriod:      time.Now(),
					DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
					AutomaticPayment:   true,
					BidTime:            time.Now(),
					InterestAmount:     sdk.NewCoin("uguu", sdk.NewInt(0)),
					Borrowings:         []types.Borrowing{},
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: bidder1.String(),
					},
				},
				blockTime: time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			funcFExp{
				expRefundBid:     types.NftBids{},
				expTotalInterest: sdk.Coin{"uguu", sdk.NewInt(0)},
				expSurplusAmount: sdk.Coin{"uguu", sdk.NewInt(0)},
				expListing: types.NftListing{
					NftId: types.NftIdentifier{
						ClassId: "1",
						NftId:   "1",
					},
					Owner:              owner.String(),
					ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
					State:              types.ListingState_LISTING,
					BidToken:           "uguu",
					MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
					StartedAt:          now,
					EndAt:              now,
					FullPaymentEndAt:   time.Time{},
					SuccessfulBidEndAt: time.Time{},
					AutoRelistedCount:  0,
					CollectedAmount: sdk.Coin{
						Denom:  "uguu",
						Amount: sdk.ZeroInt(),
					},
				},
			},
			nil,
		},
		{
			"refund bid and no collect bid",
			funcArg{
				collectBid: types.NftBids{},
				refundBid: types.NftBids{
					{
						BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
						DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(50)),
						DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
						Borrowings: []types.Borrowing{
							{
								Amount:             sdk.NewCoin("uguu", sdk.NewInt(10)),
								PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
								StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						},
						InterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
					},
					{
						BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(99)),
						DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(45)),
						DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
						InterestAmount:     sdk.NewCoin("uguu", sdk.NewInt(10)),
						Borrowings: []types.Borrowing{
							{
								Amount:             sdk.NewCoin("uguu", sdk.NewInt(10)),
								PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
								StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						},
						PaidAmount: sdk.NewCoin("uguu", sdk.NewInt(55)),
					},
				},
				listing: types.NftListing{
					NftId: types.NftIdentifier{
						ClassId: "1",
						NftId:   "1",
					},
					Owner:              owner.String(),
					ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
					State:              types.ListingState_LISTING,
					BidToken:           "uguu",
					MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
					StartedAt:          now,
					EndAt:              now,
					FullPaymentEndAt:   time.Time{},
					SuccessfulBidEndAt: time.Time{},
					AutoRelistedCount:  0,
					CollectedAmount: sdk.Coin{
						Denom:  "uguu",
						Amount: sdk.ZeroInt(),
					},
				},
				winnerBid: types.NftBid{
					NftId: types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder:             "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
					DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(50)),
					PaidAmount:         sdk.NewCoin("uguu", sdk.NewInt(0)),
					BiddingPeriod:      time.Now(),
					DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
					AutomaticPayment:   true,
					BidTime:            time.Now(),
					InterestAmount:     sdk.NewCoin("uguu", sdk.NewInt(0)),
					Borrowings:         []types.Borrowing{},
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
				},
				blockTime: time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			funcFExp{
				expRefundBid: types.NftBids{
					{
						BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
						DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(50)),
						DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
						Borrowings: []types.Borrowing{
							{
								Amount:             sdk.NewCoin("uguu", sdk.NewInt(10)),
								PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
								StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						},
						InterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
					},
					{
						BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(99)),
						DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(45)),
						DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
						InterestAmount:     sdk.NewCoin("uguu", sdk.NewInt(10)),
						Borrowings: []types.Borrowing{
							{
								Amount:             sdk.NewCoin("uguu", sdk.NewInt(10)),
								PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
								StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						},
						PaidAmount: sdk.NewCoin("uguu", sdk.NewInt(55)),
					},
				},
				expTotalInterest: sdk.Coin{"uguu", sdk.NewInt(12)},
				expSurplusAmount: sdk.Coin{"uguu", sdk.NewInt(5)},
				expListing: types.NftListing{
					NftId: types.NftIdentifier{
						ClassId: "1",
						NftId:   "1",
					},
					Owner:              owner.String(),
					ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
					State:              types.ListingState_LISTING,
					BidToken:           "uguu",
					MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
					StartedAt:          now,
					EndAt:              now,
					FullPaymentEndAt:   time.Time{},
					SuccessfulBidEndAt: time.Time{},
					AutoRelistedCount:  0,
					CollectedAmount: sdk.Coin{
						Denom:  "uguu",
						Amount: sdk.ZeroInt(),
					},
				},
			},
			nil,
		},
		{
			"refund bid and collect bid",
			funcArg{
				collectBid: types.NftBids{
					{
						BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
						DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(50)),
						DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
						Borrowings: []types.Borrowing{
							{
								Amount:             sdk.NewCoin("uguu", sdk.NewInt(10)),
								PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
								StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						},
						InterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
						Bidder:         bidder1.String(),
						NftId: types.NftIdentifier{
							ClassId: "1",
							NftId:   "1",
						},
					},
				},
				refundBid: types.NftBids{
					{
						BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
						DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(50)),
						DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
						Borrowings: []types.Borrowing{
							{
								Amount:             sdk.NewCoin("uguu", sdk.NewInt(10)),
								PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
								StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						},
						InterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
					},
					{
						BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(99)),
						DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(45)),
						DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
						InterestAmount:     sdk.NewCoin("uguu", sdk.NewInt(10)),
						Borrowings: []types.Borrowing{
							{
								Amount:             sdk.NewCoin("uguu", sdk.NewInt(10)),
								PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
								StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						},
						PaidAmount: sdk.NewCoin("uguu", sdk.NewInt(55)),
					},
				},
				listing: types.NftListing{
					NftId: types.NftIdentifier{
						ClassId: "1",
						NftId:   "1",
					},
					Owner:              owner.String(),
					ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
					State:              types.ListingState_LISTING,
					BidToken:           "uguu",
					MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
					StartedAt:          now,
					EndAt:              now,
					FullPaymentEndAt:   time.Time{},
					SuccessfulBidEndAt: time.Time{},
					AutoRelistedCount:  0,
					CollectedAmount: sdk.Coin{
						Denom:  "uguu",
						Amount: sdk.ZeroInt(),
					},
				},
				winnerBid: types.NftBid{
					NftId: types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder:             "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
					DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(50)),
					PaidAmount:         sdk.NewCoin("uguu", sdk.NewInt(0)),
					BiddingPeriod:      time.Now(),
					DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
					AutomaticPayment:   true,
					BidTime:            time.Now(),
					InterestAmount:     sdk.NewCoin("uguu", sdk.NewInt(0)),
					Borrowings:         []types.Borrowing{},
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
				},
				blockTime: time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			funcFExp{
				expRefundBid: types.NftBids{
					{
						BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
						DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(50)),
						DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
						Borrowings: []types.Borrowing{
							{
								Amount:             sdk.NewCoin("uguu", sdk.NewInt(10)),
								PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
								StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						},
						InterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
					},
					{
						BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(99)),
						DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(45)),
						DepositLendingRate: sdk.MustNewDecFromStr("0.1"),
						InterestAmount:     sdk.NewCoin("uguu", sdk.NewInt(10)),
						Borrowings: []types.Borrowing{
							{
								Amount:             sdk.NewCoin("uguu", sdk.NewInt(10)),
								PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
								StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
							},
						},
						PaidAmount: sdk.NewCoin("uguu", sdk.NewInt(55)),
					},
				},
				expTotalInterest: sdk.Coin{"uguu", sdk.NewInt(12)},
				expSurplusAmount: sdk.Coin{"uguu", sdk.NewInt(55)},
				expListing: types.NftListing{
					NftId: types.NftIdentifier{
						ClassId: "1",
						NftId:   "1",
					},
					Owner:              owner.String(),
					ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
					State:              types.ListingState_LISTING,
					BidToken:           "uguu",
					MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
					StartedAt:          now,
					EndAt:              now,
					FullPaymentEndAt:   time.Time{},
					SuccessfulBidEndAt: time.Time{},
					AutoRelistedCount:  0,
					CollectedAmount: sdk.Coin{
						Denom:  "uguu",
						Amount: sdk.NewInt(50),
					},
				},
			},
			nil,
		},
	}
	suite.SetupTest()
	for _, tc := range tcs {
		err := suite.keeper.LiquidationProcessExitsWinner(suite.ctx,
			tc.funcArg.collectBid, tc.funcArg.refundBid,
			tc.funcArg.listing, tc.funcArg.winnerBid,
			tc.funcArg.blockTime,
			refundFCheckerGen(tc.testCase, tc.funcFExp.expRefundBid, tc.funcFExp.expTotalInterest, tc.funcFExp.expSurplusAmount, tc.funcFExp.expListing),
		)
		if tc.expResult != nil {
			suite.Equal(tc.expResult, err)
		} else {
			suite.NoError(err)
		}
	}
}
