package keeper_test

import (
	"time"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/tendermint/tendermint/crypto/ed25519"
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
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          now,
			EndAt:              now,
			FullPaymentEndAt:   time.Time{},
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
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
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          now,
			EndAt:              now,
			FullPaymentEndAt:   time.Time{},
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
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
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          now,
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
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
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          time.Time{},
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
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
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          time.Time{},
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: now,
			AutoRelistedCount:  0,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "2",
				NftId:   "2",
			},
			Owner:              owner2.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_LIQUIDATION,
			BidToken:           "uguu",
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          time.Time{},
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: now,
			AutoRelistedCount:  0,
		},
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
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase   string
		classId    string
		nftId      string
		nftOwner   sdk.AccAddress
		lister     sdk.AccAddress
		bidToken   string
		activeRank uint64
		mintBefore bool
		listBefore bool
		expectPass bool
	}{
		{
			testCase:   "not existing nft",
			classId:    "class1",
			nftId:      "nft1",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "uguu",
			activeRank: 1,
			mintBefore: false,
			listBefore: false,
			expectPass: false,
		},
		{
			testCase:   "already listed",
			classId:    "class2",
			nftId:      "nft2",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "uguu",
			activeRank: 1,
			mintBefore: true,
			listBefore: true,
			expectPass: false,
		},
		{
			testCase:   "not owned nft",
			classId:    "class3",
			nftId:      "nft3",
			nftOwner:   acc1,
			lister:     acc2,
			bidToken:   "uguu",
			activeRank: 1,
			mintBefore: true,
			listBefore: false,
			expectPass: false,
		},
		{
			testCase:   "unsupported bid token",
			classId:    "class4",
			nftId:      "nft4",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "xxxx",
			activeRank: 1,
			mintBefore: true,
			listBefore: false,
			expectPass: false,
		},
		{
			testCase:   "successful listing with default active rank",
			classId:    "class5",
			nftId:      "nft5",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "uguu",
			activeRank: 0,
			mintBefore: true,
			listBefore: false,
			expectPass: true,
		},
		{
			testCase:   "successful listing with non-default active rank",
			classId:    "class6",
			nftId:      "nft6",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "uguu",
			activeRank: 100,
			mintBefore: true,
			listBefore: false,
			expectPass: true,
		},
	}

	for _, tc := range tests {
		if tc.mintBefore {
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
		}
		if tc.listBefore {
			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:        ununifitypes.StringAccAddress(tc.lister),
				NftId:         types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId},
				ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:      tc.bidToken,
				MinBid:        sdk.ZeroInt(),
				BidActiveRank: tc.activeRank,
			})
			suite.Require().NoError(err)
		}
		err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:        ununifitypes.StringAccAddress(tc.lister),
			NftId:         types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId},
			ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
			BidToken:      tc.bidToken,
			MinBid:        sdk.ZeroInt(),
			BidActiveRank: tc.activeRank,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)
			// get listing
			listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, (types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}).IdBytes())
			suite.Require().NoError(err)

			// check ownership is transferred
			moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
			nftOwner := suite.app.NFTKeeper.GetOwner(suite.ctx, tc.classId, tc.nftId)
			suite.Require().Equal(nftOwner.String(), moduleAddr.String())

			// check bid active rank is set to default if zero
			if tc.activeRank == 0 {
				suite.Require().Equal(params.DefaultBidActiveRank, listing.BidActiveRank)
			}

			// check startedAt is set as current time
			suite.Require().Equal(suite.ctx.BlockTime(), listing.StartedAt)

			// check endAt is set from initial listing duration
			suite.Require().Equal(suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingPeriodInitial)), listing.EndAt)
		} else {
			suite.Require().Error(err)
		}
	}
}

// TODO:Add test for CancelNftListing(ctx sdk.Context, msg *types.MsgCancelNftListing) error
// TODO:Add test for ExpandListingPeriod(ctx sdk.Context, msg *types.MsgExpandListingPeriod) error
// TODO:Add test for SellingDecision(ctx sdk.Context, msg *types.MsgSellingDecision) error
// TODO:Add test for EndNftListing(ctx sdk.Context, msg *types.MsgEndNftListing) error
// TODO:Add test for ProcessEndingNftListings(ctx sdk.Context)
// TODO:Add test for HandleFullPaymentsPeriodEndings(ctx sdk.Context)
// TODO:Add test for DelieverSuccessfulBids(ctx sdk.Context)
// TODO:Add test for ProcessPaymentWithCommissionFee(ctx sdk.Context, listingOwner sdk.AccAddress, denom string, amount sdk.Int)

// TODO: add test for NftListing following scenario
// Create Nft
// List Nft
// End NftListing
// Check GetActiveNftListingsEndingAt is not showing any value
