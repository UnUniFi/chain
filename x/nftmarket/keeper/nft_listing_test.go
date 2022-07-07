package keeper_test

import (
	"time"

	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

// test basic functions of nft listing
func (suite *KeeperTestSuite) TestNftListingBasics() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

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
			StartedAt:          time.Time{},
			EndAt:              time.Time{},
			FullPaymentEndAt:   time.Time{},
			SuccessfulBidEndAt: time.Time{},
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

	allListings := suite.app.NftmarketKeeper.GetAllNftListings(suite.ctx)
	suite.Require().Len(allListings, 1)
	suite.Require().Equal(listings, allListings)

	// TODO: add tests for following functions as well
	// GetListingsByOwner(ctx sdk.Context, owner sdk.AccAddress) []types.NftListing {
	// SetNftListing(ctx sdk.Context, listing types.NftListing) {
	// DeleteNftListing(ctx sdk.Context, listing types.NftListing) {
	// GetActiveNftListingsEndingAt(ctx sdk.Context, endTime time.Time) []types.NftListing {
	// GetFullPaymentNftListingsEndingAt(ctx sdk.Context, endTime time.Time) []types.NftListing {
	// GetSuccessfulBidNftListingsEndingAt(ctx sdk.Context, endTime time.Time) []types.NftListing {

}

// TODO: add test for NftListing following scenario
// Create Nft
// List Nft
// End NftListing
// Check GetActiveNftListingsEndingAt is not showing any value
