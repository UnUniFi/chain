package keeper_test

import (
	gocontext "context"
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/UnUniFi/chain/x/nftmarketv1/types"
)

func TestGRPCQuery(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) TestListedNfts() {
	var req *types.QueryListedNftsRequest
	testCases := []struct {
		msg        string
		malleate   func(index int, require *require.Assertions)
		expError   string
		listingNft []types.NftListing
		postTest   func(index int, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListing)
	}{
		{
			"success empty",
			func(index int, require *require.Assertions) {
				req = &types.QueryListedNftsRequest{}
			},
			"",
			[]types.NftListing(nil),
			func(index int, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListing) {
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
			[]types.NftListing{},
			func(index int, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListing) {
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
			[]types.NftListing{
				{
					NftId:              types.NftIdentifier{ClassId: "class2", NftId: "nft2"},
					Owner:              s.addrs[0].String(),
					ListingType:        0,
					State:              0,
					BidToken:           "uguu",
					MinBid:             sdk.NewInt(0),
					BidActiveRank:      0x1,
					StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:              time.Date(1, time.January, 1, 0, 1, 0, 0, time.UTC),
					FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					AutoRelistedCount:  0x0,
				},
				{
					NftId:       types.NftIdentifier{ClassId: "class5", NftId: "nft5"},
					Owner:       s.addrs[0].String(),
					ListingType: 0, State: 0, BidToken: "uguu",
					MinBid:             sdk.NewInt(0),
					BidActiveRank:      1,
					StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:              time.Date(1, time.January, 1, 0, 1, 0, 0, time.UTC),
					FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					AutoRelistedCount:  0,
				},
				{
					NftId:       types.NftIdentifier{ClassId: "class6", NftId: "nft6"},
					Owner:       s.addrs[0].String(),
					ListingType: 0, State: 0, BidToken: "uguu",
					MinBid:             sdk.NewInt(0),
					BidActiveRank:      100,
					StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:              time.Date(1, time.January, 1, 0, 1, 0, 0, time.UTC),
					FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					AutoRelistedCount:  0,
				},
			},
			func(index int, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListing) {
				require.Equal(res.Listings, expListingNft, "the error occurred on:%d", index)
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
			[]types.NftListing{
				{
					NftId:       types.NftIdentifier{ClassId: "class7", NftId: "nft7"},
					Owner:       s.addrs[1].String(),
					ListingType: 0, State: 0, BidToken: "uguu",
					MinBid:             sdk.NewInt(0),
					BidActiveRank:      1,
					StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:              time.Date(1, time.January, 1, 0, 1, 0, 0, time.UTC),
					FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					AutoRelistedCount:  0,
				},
			},
			func(index int, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListing) {
				require.Equal(res.Listings, expListingNft, "the error occurred on:%d", index)
			},
		},
		{
			"Success all req",
			func(index int, require *require.Assertions) {
				req = &types.QueryListedNftsRequest{}
			},
			"",
			[]types.NftListing{
				{
					NftId:              types.NftIdentifier{ClassId: "class2", NftId: "nft2"},
					Owner:              s.addrs[0].String(),
					ListingType:        0,
					State:              0,
					BidToken:           "uguu",
					MinBid:             sdk.NewInt(0),
					BidActiveRank:      0x1,
					StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:              time.Date(1, time.January, 1, 0, 1, 0, 0, time.UTC),
					FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					AutoRelistedCount:  0x0,
				},
				{
					NftId:       types.NftIdentifier{ClassId: "class5", NftId: "nft5"},
					Owner:       s.addrs[0].String(),
					ListingType: 0, State: 0, BidToken: "uguu",
					MinBid:             sdk.NewInt(0),
					BidActiveRank:      1,
					StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:              time.Date(1, time.January, 1, 0, 1, 0, 0, time.UTC),
					FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					AutoRelistedCount:  0,
				},
				{
					NftId:       types.NftIdentifier{ClassId: "class6", NftId: "nft6"},
					Owner:       s.addrs[0].String(),
					ListingType: 0, State: 0, BidToken: "uguu",
					MinBid:             sdk.NewInt(0),
					BidActiveRank:      100,
					StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:              time.Date(1, time.January, 1, 0, 1, 0, 0, time.UTC),
					FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					AutoRelistedCount:  0,
				},
				{
					NftId:       types.NftIdentifier{ClassId: "class7", NftId: "nft7"},
					Owner:       s.addrs[1].String(),
					ListingType: 0, State: 0, BidToken: "uguu",
					MinBid:             sdk.NewInt(0),
					BidActiveRank:      1,
					StartedAt:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					EndAt:              time.Date(1, time.January, 1, 0, 1, 0, 0, time.UTC),
					FullPaymentEndAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					SuccessfulBidEndAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					AutoRelistedCount:  0,
				},
			},
			func(index int, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListing) {
				require.Equal(res.Listings, expListingNft, "the error occurred on:%d", index)
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
			[]types.NftListing(nil),
			func(index int, require *require.Assertions, res *types.QueryListedNftsResponse, expListingNft []types.NftListing) {
				require.Equal(res.Listings, expListingNft, "the error occurred on:%d", index)
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
			tc.postTest(index, require, result, tc.listingNft)
		})
	}
}

func (s *KeeperTestSuite) TestLoan() {
	testCases := []struct {
		msg       string
		malleate  func(index int, require *require.Assertions)
		req       *types.QueryLoanRequest
		expError  string
		expResult types.QueryLoanResponse
	}{
		{
			"success empty",
			func(index int, require *require.Assertions) {
			},
			&types.QueryLoanRequest{},
			"",
			types.QueryLoanResponse{
				Loan: types.Loan{
					NftId: types.NftIdentifier{},
					Loan: sdk.Coin{
						Amount: sdk.NewInt(0),
					},
				},
				BorrowingLimit: sdk.ZeroInt(),
			},
		},
		{
			"fail invalid class id",
			func(index int, require *require.Assertions) {
			},
			&types.QueryLoanRequest{
				ClassId: "ddfdifd",
				NftId:   "a10",
			},
			"",
			types.QueryLoanResponse{
				Loan: types.Loan{
					NftId: types.NftIdentifier{},
					Loan: sdk.Coin{
						Amount: sdk.NewInt(0),
					},
				},
				BorrowingLimit: sdk.ZeroInt(),
			},
		},
		{
			"Success",
			func(index int, require *require.Assertions) {
				s.TestBorrow()
			},
			&types.QueryLoanRequest{
				ClassId: "class5",
				NftId:   "nft5",
			},
			"",
			types.QueryLoanResponse{
				Loan: types.Loan{
					NftId: types.NftIdentifier{
						ClassId: "class5",
						NftId:   "nft5",
					},
					Loan: sdk.Coin{
						Denom:  "uguu",
						Amount: sdk.NewInt(2000000),
					},
				},
				BorrowingLimit: sdk.NewInt(3500000),
			},
		},
	}
	for index, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			require := s.Require()
			tc.malleate(index, require)
			result, err := s.queryClient.Loan(gocontext.Background(), tc.req)
			if tc.expError == "" {
				require.NoError(err)
				require.Equal(result, &tc.expResult, "the error occurred on:%d", index)
			} else {
				require.Error(err)
				require.Contains(err.Error(), tc.expError)
			}
		})
	}
}

func (s *KeeperTestSuite) TestClasses() {
	testCases := []struct {
		msg       string
		malleate  func(index int, require *require.Assertions)
		req       *types.QueryListedClassesRequest
		expError  string
		expResult types.QueryListedClassesResponse
	}{
		{
			"success empty",
			func(index int, require *require.Assertions) {
			},
			&types.QueryListedClassesRequest{},
			"",
			types.QueryListedClassesResponse{
				Classes: []*types.QueryListedClassResponse(nil),
			},
		},
		{
			"Success list nft",
			func(index int, require *require.Assertions) {
				s.TestListNft()
			},
			&types.QueryListedClassesRequest{},
			"",
			types.QueryListedClassesResponse{
				Classes: []*types.QueryListedClassResponse{
					{
						ClassId:     "class2",
						Name:        "class2",
						Description: "class2",
						Symbol:      "class2",
						Uri:         "class2",
						Urihash:     "",
						Nfts: []*types.ListedNft{
							{
								Id:      "nft2",
								Uri:     "nft2",
								UriHash: "nft2",
							},
						},
						NftCount: 1,
					},
					{
						ClassId:     "class5",
						Name:        "class5",
						Description: "class5",
						Symbol:      "class5",
						Uri:         "class5",
						Urihash:     "",
						Nfts: []*types.ListedNft{
							{
								Id:      "nft5",
								Uri:     "nft5",
								UriHash: "nft5",
							},
						},
						NftCount: 1,
					},
					{
						ClassId:     "class6",
						Name:        "class6",
						Description: "class6",
						Symbol:      "class6",
						Uri:         "class6",
						Urihash:     "",
						Nfts: []*types.ListedNft{
							{
								Id:      "nft6",
								Uri:     "nft6",
								UriHash: "nft6",
							},
						},
						NftCount: 1,
					},
					{
						ClassId:     "class7",
						Name:        "class7",
						Description: "class7",
						Symbol:      "class7",
						Uri:         "class7",
						Urihash:     "",
						Nfts: []*types.ListedNft{
							{
								Id:      "nft7",
								Uri:     "nft7",
								UriHash: "nft7",
							},
						},
						NftCount: 1,
					},
				},
			},
		},
		{
			"Success plaace bid",
			func(index int, require *require.Assertions) {
				s.TestPlaceBid()
			},
			&types.QueryListedClassesRequest{},
			"",
			types.QueryListedClassesResponse{
				Classes: []*types.QueryListedClassResponse{
					{
						ClassId:     "class5",
						Name:        "class5",
						Description: "class5",
						Symbol:      "class5",
						Uri:         "class5",
						Urihash:     "",
						Nfts: []*types.ListedNft{
							{
								Id:      "nft5",
								Uri:     "nft5",
								UriHash: "nft5",
							},
						},
						NftCount: 1,
					},
				},
			},
		},
	}
	for index, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			require := s.Require()
			tc.malleate(index, require)
			result, err := s.queryClient.ListedClasses(gocontext.Background(), tc.req)
			if tc.expError == "" {
				require.NoError(err)
				require.Equal(result, &tc.expResult, "the error occurred on:%d", index)
			} else {
				require.Error(err)
				require.Contains(err.Error(), tc.expError)
			}
		})
	}
}
