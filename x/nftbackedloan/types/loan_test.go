package types_test

import (
	"testing"
	"time"

	types "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	backedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func TestMinSettlementAmount(t *testing.T) {
	testCases := []struct {
		name      string
		bids      []backedloantypes.NftBid
		expResult types.Coin
	}{
		{
			name:      "empty bids",
			bids:      []backedloantypes.NftBid{},
			expResult: types.Coin{},
		},
		{
			name: "one bid",
			bids: []backedloantypes.NftBid{
				{
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 30),
				},
			},
			expResult: types.NewInt64Coin("uatom", 30),
		},
		{
			name: "two bids, totalDepositAmount < bidAmount",
			bids: []backedloantypes.NftBid{
				{
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 30),
				},
				{
					BidAmount:     types.NewInt64Coin("uatom", 200),
					DepositAmount: types.NewInt64Coin("uatom", 50),
				},
			},
			expResult: types.NewInt64Coin("uatom", 80),
		},
		{
			name: "three bids & bidAmount < totalDepositAmount",
			bids: []backedloantypes.NftBid{
				{
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 30),
				},
				{
					BidAmount:     types.NewInt64Coin("uatom", 200),
					DepositAmount: types.NewInt64Coin("uatom", 20),
				},
				{
					BidAmount:     types.NewInt64Coin("uatom", 300),
					DepositAmount: types.NewInt64Coin("uatom", 15),
				},
			},
			expResult: types.NewInt64Coin("uatom", 100),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := backedloantypes.MinSettlementAmount(tc.bids)
			if !result.IsEqual(tc.expResult) {
				t.Errorf("expected %s, got %s", tc.expResult, result)
			}
		})
	}
}

func TestLiquidationBid(t *testing.T) {
	testCases := []struct {
		name      string
		bids      []backedloantypes.NftBid
		expResult backedloantypes.NftBid
	}{
		{
			name:      "empty bids",
			bids:      []backedloantypes.NftBid{},
			expResult: backedloantypes.NftBid{},
		},
		{
			name: "one bid, paid",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 30),
					PaidAmount:    types.NewInt64Coin("uatom", 70),
				},
			},
			expResult: backedloantypes.NftBid{
				Id: backedloantypes.BidId{
					NftId: &backedloantypes.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
			},
		},
		{
			name: "one bid, paid",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 30),
				},
			},
			expResult: backedloantypes.NftBid{},
		},
		{
			name: "two bids, paid",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 30),
					PaidAmount:    types.NewInt64Coin("uatom", 70),
				},
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     types.NewInt64Coin("uatom", 200),
					DepositAmount: types.NewInt64Coin("uatom", 50),
					PaidAmount:    types.NewInt64Coin("uatom", 150),
				},
			},
			expResult: backedloantypes.NftBid{
				Id: backedloantypes.BidId{
					NftId: &backedloantypes.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
				},
			},
		},
		{
			name: "two bids, unpaid top bid",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 30),
					PaidAmount:    types.NewInt64Coin("uatom", 70),
				},
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     types.NewInt64Coin("uatom", 200),
					DepositAmount: types.NewInt64Coin("uatom", 50),
				},
			},
			expResult: backedloantypes.NftBid{
				Id: backedloantypes.BidId{
					NftId: &backedloantypes.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := backedloantypes.LiquidationBid(tc.bids, time.Now())
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestForForfeitedBidsAndRefundBids(t *testing.T) {
	testCases := []struct {
		name      string
		bids      []backedloantypes.NftBid
		winBid    backedloantypes.NftBid
		expResult []int
	}{
		{
			name: "two bids, paid",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 30),
					PaidAmount:    types.NewInt64Coin("uatom", 70),
				},
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     types.NewInt64Coin("uatom", 200),
					DepositAmount: types.NewInt64Coin("uatom", 50),
					PaidAmount:    types.NewInt64Coin("uatom", 150),
				},
			},
			winBid: backedloantypes.NftBid{
				Id: backedloantypes.BidId{
					NftId: &backedloantypes.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
				},
			},
			expResult: []int{0, 1},
		},
		{
			name: "two bids, unpaid top bid",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 30),
					PaidAmount:    types.NewInt64Coin("uatom", 70),
				},
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     types.NewInt64Coin("uatom", 200),
					DepositAmount: types.NewInt64Coin("uatom", 50),
				},
			},
			winBid: backedloantypes.NftBid{
				Id: backedloantypes.BidId{
					NftId: &backedloantypes.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
			},
			expResult: []int{1, 0},
		},
	}
	for _, tc := range testCases {
		forfeitedBids, refundBids := backedloantypes.ForfeitedBidsAndRefundBids(tc.bids, tc.winBid)
		if tc.expResult[0] != len(forfeitedBids) {
			t.Error("forfeitedBids expected length %d, got %d", tc.expResult[0], len(forfeitedBids))
		}
		if tc.expResult[1] != len(refundBids) {
			t.Error("refundBids expected length %d, got %d", tc.expResult[1], len(refundBids))
		}
	}
}
