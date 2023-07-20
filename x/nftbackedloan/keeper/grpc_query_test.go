package keeper_test

import (
	gocontext "context"
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func TestGRPCQuery(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) TestListedNfts() {
	var req *types.QueryListedNftsRequest
	type postTest func(index int, msg string, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListingDetail)
	testCases := []struct {
		msg        string
		malleate   func(index int, require *require.Assertions)
		expError   string
		listingNft []types.NftListingDetail
		postTest   postTest
	}{
		{
			"success empty",
			func(index int, require *require.Assertions) {
				req = &types.QueryListedNftsRequest{}
			},
			"",
			[]types.NftListingDetail(nil),
			func(index int, msg string, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListingDetail) {
				require.Equal(res.Listings, expListingNft, "the error occurred on:%d", index)
			},
		},
		{
			"fail invalid Owner addr",
			func(index int, require *require.Assertions) {
				req = &types.QueryListedNftsRequest{
					Owner: "owner",
				}
			},
			"invalid request. address wrong",
			[]types.NftListingDetail{},
			func(index int, msg string, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListingDetail) {
			},
		},
		{
			"Success owner1",
			func(index int, require *require.Assertions) {
				s.TestListNft()
				req = &types.QueryListedNftsRequest{
					Owner: s.addrs[0].String(),
				}
			},
			"",
			[]types.NftListingDetail{
				{
					Listing: types.NftListing{
						NftId:              types.NftIdentifier{ClassId: "class2", NftId: "nft2"},
						Owner:              s.addrs[0].String(),
						State:              1,
						BidDenom:           "uguu",
						MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
						StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						LiquidatedAt:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						CollectedAmount: sdk.Coin{
							Denom:  "uguu",
							Amount: sdk.ZeroInt(),
						},
						CollectedAmountNegative: false,
					},
					NftInfo: types.NftInfo{
						Id:      "nft2",
						Uri:     "nft2",
						UriHash: "nft2",
					},
				},
				{
					Listing: types.NftListing{
						NftId: types.NftIdentifier{ClassId: "class5", NftId: "nft5"},
						Owner: s.addrs[0].String(),
						State: 1, BidDenom: "uguu",
						MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
						StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						LiquidatedAt:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						CollectedAmount: sdk.Coin{
							Denom:  "uguu",
							Amount: sdk.ZeroInt(),
						},
						CollectedAmountNegative: false,
					},
					NftInfo: types.NftInfo{
						Id:      "nft5",
						Uri:     "nft5",
						UriHash: "nft5",
					},
				},
				{
					Listing: types.NftListing{
						NftId: types.NftIdentifier{ClassId: "class6", NftId: "nft6"},
						Owner: s.addrs[0].String(),
						State: 1, BidDenom: "uguu",
						MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
						StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						LiquidatedAt:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						CollectedAmount: sdk.Coin{
							Denom:  "uguu",
							Amount: sdk.ZeroInt(),
						},
						CollectedAmountNegative: false,
					},
					NftInfo: types.NftInfo{
						Id:      "nft6",
						Uri:     "nft6",
						UriHash: "nft6",
					},
				},
			},
			func(index int, msg string, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListingDetail) {
				require.Equal(expListingNft, res.Listings, "the error occurred on:%d", msg, res.Listings[index].Listing.NftId)
			},
		},
		{
			"Success owner2",
			func(index int, require *require.Assertions) {
				req = &types.QueryListedNftsRequest{
					Owner: s.addrs[1].String(),
				}
			},
			"",
			[]types.NftListingDetail{
				{
					Listing: types.NftListing{
						NftId: types.NftIdentifier{ClassId: "class7", NftId: "nft7"},
						Owner: s.addrs[1].String(),
						State: 1, BidDenom: "uguu",
						MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
						StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						LiquidatedAt:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						CollectedAmount: sdk.Coin{
							Denom:  "uguu",
							Amount: sdk.ZeroInt(),
						},
						CollectedAmountNegative: false,
					},
					NftInfo: types.NftInfo{
						Id:      "nft7",
						Uri:     "nft7",
						UriHash: "nft7",
					},
				},
			},
			func(index int, msg string, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListingDetail) {
				require.Equal(res.Listings, expListingNft, "the error occurred on:%d", msg)
			},
		},
		{
			"Success all req",
			func(index int, require *require.Assertions) {
				req = &types.QueryListedNftsRequest{}
			},
			"",
			[]types.NftListingDetail{
				{
					Listing: types.NftListing{
						NftId:              types.NftIdentifier{ClassId: "class2", NftId: "nft2"},
						Owner:              s.addrs[0].String(),
						State:              1,
						BidDenom:           "uguu",
						MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
						StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						LiquidatedAt:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						CollectedAmount: sdk.Coin{
							Denom:  "uguu",
							Amount: sdk.ZeroInt(),
						},
						CollectedAmountNegative: false,
					},
					NftInfo: types.NftInfo{
						Id:      "nft2",
						Uri:     "nft2",
						UriHash: "nft2",
					},
				},
				{
					Listing: types.NftListing{
						NftId: types.NftIdentifier{ClassId: "class5", NftId: "nft5"},
						Owner: s.addrs[0].String(),
						State: 1, BidDenom: "uguu",
						MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
						StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						LiquidatedAt:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						CollectedAmount: sdk.Coin{
							Denom:  "uguu",
							Amount: sdk.ZeroInt(),
						},
						CollectedAmountNegative: false,
					},
					NftInfo: types.NftInfo{
						Id:      "nft5",
						Uri:     "nft5",
						UriHash: "nft5",
					},
				},
				{
					Listing: types.NftListing{
						NftId: types.NftIdentifier{ClassId: "class6", NftId: "nft6"},
						Owner: s.addrs[0].String(),
						State: 1, BidDenom: "uguu",
						MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
						StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						LiquidatedAt:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						CollectedAmount: sdk.Coin{
							Denom:  "uguu",
							Amount: sdk.ZeroInt(),
						},
						CollectedAmountNegative: false,
					},
					NftInfo: types.NftInfo{
						Id:      "nft6",
						Uri:     "nft6",
						UriHash: "nft6",
					},
				},
				{
					Listing: types.NftListing{
						NftId: types.NftIdentifier{ClassId: "class7", NftId: "nft7"},
						Owner: s.addrs[1].String(),
						State: 1, BidDenom: "uguu",
						MinimumDepositRate: sdk.MustNewDecFromStr("0.1"),
						StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						LiquidatedAt:       time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
						CollectedAmount: sdk.Coin{
							Denom:  "uguu",
							Amount: sdk.ZeroInt(),
						},
						CollectedAmountNegative: false,
					},
					NftInfo: types.NftInfo{
						Id:      "nft7",
						Uri:     "nft7",
						UriHash: "nft7",
					},
				},
			},
			func(index int, msg string, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListingDetail) {
				require.Equal(res.Listings, expListingNft, "the error occurred on:%d", msg)
			},
		},
		{
			"success empty owner",
			func(index int, require *require.Assertions) {
				req = &types.QueryListedNftsRequest{
					Owner: s.addrs[2].String(),
				}
			},
			"",
			[]types.NftListingDetail(nil),
			func(index int, msg string, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListingDetail) {
				require.Equal(res.Listings, expListingNft, "the error occurred on:%d", msg, index)
			},
		},
	}
	for index, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			require := s.Require()
			tc.malleate(index, require)
			result, err := s.queryClient.ListedNfts(gocontext.Background(), req)
			if tc.expError == "" {
				require.NoError(err)
			} else {
				require.Error(err)
				require.Contains(err.Error(), tc.expError)
			}
			tc.postTest(index, tc.msg, require, result, tc.listingNft)
		})
	}
}
