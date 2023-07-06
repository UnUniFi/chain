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
			name:      "empty bid",
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
					BidAmount:     types.NewInt64Coin("uatom", 100),
					DepositAmount: types.NewInt64Coin("uatom", 40),
				},
				{
					BidAmount:     types.NewInt64Coin("uatom", 105),
					DepositAmount: types.NewInt64Coin("uatom", 45),
				},
			},
			expResult: types.NewInt64Coin("uatom", 105),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := backedloantypes.MinSettlementAmount(tc.bids)
			if err != nil {
				if tc.name != "empty bid" {
					t.Errorf("unexpected error: %v", err)
				}
			}
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
			name:      "empty bid",
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
					PaidAmount:    types.NewInt64Coin("uatom", 0),
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
					PaidAmount:    types.NewInt64Coin("uatom", 0),
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
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bidsSortedByDeposit := backedloantypes.NftBids(tc.bids).SortHigherDeposit()
			result, err := backedloantypes.LiquidationBid(bidsSortedByDeposit, time.Now())
			if err != nil {
				if tc.name != "empty bid" {
					t.Errorf("unexpected error: %v", err)
				}
			}
			require.Equal(t, tc.expResult.Id.Bidder, result.Id.Bidder)
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
					PaidAmount:    types.NewInt64Coin("uatom", 0),
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
		bidsSortedByDeposit := backedloantypes.NftBids(tc.bids).SortHigherDeposit()
		forfeitedBids, refundBids := backedloantypes.ForfeitedBidsAndRefundBids(bidsSortedByDeposit, tc.winBid)
		if tc.expResult[0] != len(forfeitedBids) {
			t.Errorf("forfeitedBids expected length %d, got %d", tc.expResult[0], len(forfeitedBids))
		}
		if tc.expResult[1] != len(refundBids) {
			t.Errorf("refundBids expected length %d, got %d", tc.expResult[1], len(refundBids))
		}
	}
}

func TestExpectedRepayAmount(t *testing.T) {
	testCases := []struct {
		name       string
		bids       []backedloantypes.NftBid
		borrowBids []backedloantypes.BorrowBid
		expResult  types.Coin
	}{
		{
			name:       "empty bid",
			bids:       []backedloantypes.NftBid{},
			borrowBids: []backedloantypes.BorrowBid{},
			expResult:  types.Coin{},
		},
		{
			name: "one bid",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:          types.NewInt64Coin("uatom", 100000000),
					DepositAmount:      types.NewInt64Coin("uatom", 30000000),
					DepositLendingRate: types.NewDecWithPrec(1, 1),
					// Additional 1 minute for time error correction
					BiddingPeriod: time.Now().Add(time.Hour).Add(time.Minute),
				},
			},
			borrowBids: []backedloantypes.BorrowBid{
				{
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					Amount: types.NewInt64Coin("uatom", 200000000),
				},
			},
			// 30000000 * 0.1 / 365 / 24 = 342.46
			expResult: types.NewInt64Coin("uatom", 30000342),
		},
		{
			name: "2 bid & over borrow",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:          types.NewInt64Coin("uatom", 100000000),
					DepositAmount:      types.NewInt64Coin("uatom", 30000000),
					DepositLendingRate: types.NewDecWithPrec(1, 1),
					BiddingPeriod:      time.Now().Add(time.Hour).Add(time.Minute),
				},
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:          types.NewInt64Coin("uatom", 200000000),
					DepositAmount:      types.NewInt64Coin("uatom", 30000000),
					DepositLendingRate: types.NewDecWithPrec(1, 1),
					BiddingPeriod:      time.Now().Add(time.Hour).Add(time.Minute),
				},
			},
			borrowBids: []backedloantypes.BorrowBid{
				{
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					Amount: types.NewInt64Coin("uatom", 20000000),
				},
				{
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					Amount: types.NewInt64Coin("uatom", 70000000),
				},
			},
			// if over borrow, borrow amount is deposit amount
			// 50000000 * 0.1 / 365 / 24 = 570.7762
			expResult: types.NewInt64Coin("uatom", 50000570),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := backedloantypes.ExpectedRepayAmount(tc.bids, tc.borrowBids, time.Now())
			if err != nil {
				if tc.name != "empty bid" {
					t.Errorf("unexpected error: %v", err)
				}
			}
			if !result.IsEqual(tc.expResult) {
				t.Errorf("expected %s, got %s", tc.expResult, result)
			}
		})
	}
}

func TestExistRepayAmount(t *testing.T) {
	testCases := []struct {
		name      string
		bids      []backedloantypes.NftBid
		expResult types.Coin
	}{
		{
			name:      "empty bid",
			bids:      []backedloantypes.NftBid{},
			expResult: types.Coin{},
		},
		{
			name: "one bid",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:          types.NewInt64Coin("uatom", 100000000),
					DepositAmount:      types.NewInt64Coin("uatom", 30000000),
					DepositLendingRate: types.NewDecWithPrec(1, 1),
					// Additional 1 minute for time error correction
					BiddingPeriod: time.Now().Add(time.Hour).Add(time.Minute),
					Borrowings: []backedloantypes.Borrowing{
						{
							Amount:  types.NewInt64Coin("uatom", 20000000),
							StartAt: time.Now(),
						},
					},
				},
			},
			// 20000000 * 0.1 / 365 / 24 = 228.31
			expResult: types.NewInt64Coin("uatom", 20000228),
		},
		{
			name: "2 bid",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:          types.NewInt64Coin("uatom", 100000000),
					DepositAmount:      types.NewInt64Coin("uatom", 30000000),
					DepositLendingRate: types.NewDecWithPrec(1, 1),
					// Additional 1 minute for time error correction
					BiddingPeriod: time.Now().Add(time.Hour).Add(time.Minute),
					Borrowings: []backedloantypes.Borrowing{
						{
							Amount:  types.NewInt64Coin("uatom", 20000000),
							StartAt: time.Now(),
						},
					},
				},
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:          types.NewInt64Coin("uatom", 200000000),
					DepositAmount:      types.NewInt64Coin("uatom", 40000000),
					DepositLendingRate: types.NewDecWithPrec(1, 1),
					// Additional 1 minute for time error correction
					BiddingPeriod: time.Now().Add(time.Hour).Add(time.Minute),
					Borrowings: []backedloantypes.Borrowing{
						{
							Amount:  types.NewInt64Coin("uatom", 40000000),
							StartAt: time.Now(),
						},
					},
				},
			},
			// 60000000 * 0.1 / 365 / 24 = 685.89
			expResult: types.NewInt64Coin("uatom", 60000685),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := backedloantypes.ExistRepayAmount(tc.bids)
			if err != nil {
				if tc.name != "empty bid" {
					t.Errorf("unexpected error: %v", err)
				}
			}
			if !result.IsEqual(tc.expResult) {
				t.Errorf("expected %s, got %s", tc.expResult, result)
			}
		})
	}
}

func TestIsAbleToBorrow(t *testing.T) {
	bids := []backedloantypes.NftBid{
		{
			Id: backedloantypes.BidId{
				NftId: &backedloantypes.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			BidAmount:          types.NewInt64Coin("uatom", 100000000),
			DepositAmount:      types.NewInt64Coin("uatom", 30000000),
			DepositLendingRate: types.NewDecWithPrec(5, 2), // 5%
			BiddingPeriod:      time.Now().Add(time.Hour * 72),
		},
		{
			Id: backedloantypes.BidId{
				NftId: &backedloantypes.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
			},
			BidAmount:          types.NewInt64Coin("uatom", 200000000),
			DepositAmount:      types.NewInt64Coin("uatom", 80000000),
			DepositLendingRate: types.NewDecWithPrec(1, 1), // 10%
			BiddingPeriod:      time.Now().Add(time.Hour * 24),
		},
	}
	testCases := []struct {
		name       string
		borrowBids []backedloantypes.BorrowBid
		expResult  bool
	}{
		{
			name: "able to borrow",
			borrowBids: []backedloantypes.BorrowBid{
				{
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					Amount: types.NewInt64Coin("uatom", 30000000),
				},
				{
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					Amount: types.NewInt64Coin("uatom", 20000000),
				},
			},
			expResult: true,
		},
		{
			name: "unable to borrow",
			borrowBids: []backedloantypes.BorrowBid{
				{
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					Amount: types.NewInt64Coin("uatom", 30000000),
				},
				{
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					Amount: types.NewInt64Coin("uatom", 80000000),
				},
			},
			expResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := backedloantypes.IsAbleToBorrow(bids, tc.borrowBids, time.Now())
			if result != tc.expResult {
				t.Errorf("expected %v, got %v", tc.expResult, result)
			}
		})
	}
}

func TestIsAbleToCancelBid(t *testing.T) {
	testCases := []struct {
		name      string
		bids      []backedloantypes.NftBid
		cancelBid backedloantypes.BidId
		expResult bool
	}{
		{
			name: "able to cancel",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:          types.NewInt64Coin("uatom", 100000000),
					DepositAmount:      types.NewInt64Coin("uatom", 30000000),
					DepositLendingRate: types.NewDecWithPrec(5, 2), // 5%
					BiddingPeriod:      time.Now().Add(time.Hour * 72),
					Borrowings:         []backedloantypes.Borrowing{},
				},
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:          types.NewInt64Coin("uatom", 200000000),
					DepositAmount:      types.NewInt64Coin("uatom", 80000000),
					DepositLendingRate: types.NewDecWithPrec(1, 1), // 10%
					BiddingPeriod:      time.Now().Add(time.Hour * 24),
					Borrowings:         []backedloantypes.Borrowing{},
				},
			},
			cancelBid: backedloantypes.BidId{
				NftId: &backedloantypes.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			expResult: true,
		},
		{
			name: "unable to cancel",
			bids: []backedloantypes.NftBid{
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:          types.NewInt64Coin("uatom", 100000000),
					DepositAmount:      types.NewInt64Coin("uatom", 30000000),
					DepositLendingRate: types.NewDecWithPrec(5, 2), // 5%
					BiddingPeriod:      time.Now().Add(time.Hour * 72),
					Borrowings:         []backedloantypes.Borrowing{},
				},
				{
					Id: backedloantypes.BidId{
						NftId: &backedloantypes.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:          types.NewInt64Coin("uatom", 200000000),
					DepositAmount:      types.NewInt64Coin("uatom", 80000000),
					DepositLendingRate: types.NewDecWithPrec(1, 1), // 10%
					BiddingPeriod:      time.Now().Add(time.Hour * 24),
					Borrowings: []backedloantypes.Borrowing{
						{
							// borrow 100% of deposit
							Amount:  types.NewInt64Coin("uatom", 80000000),
							StartAt: time.Now(),
						},
					},
				},
			},
			cancelBid: backedloantypes.BidId{
				NftId: &backedloantypes.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			expResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := backedloantypes.IsAbleToCancelBid(tc.cancelBid, tc.bids)
			if result != tc.expResult {
				t.Errorf("expected %v, got %v", tc.expResult, result)
			}
		})
	}
}

func TestIsAbleToReBid(t *testing.T) {
	bids := []backedloantypes.NftBid{
		{
			Id: backedloantypes.BidId{
				NftId: &backedloantypes.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			BidAmount:          types.NewInt64Coin("uatom", 100000000),
			DepositAmount:      types.NewInt64Coin("uatom", 30000000),
			DepositLendingRate: types.NewDecWithPrec(5, 2), // 5%
			BiddingPeriod:      time.Now().Add(time.Hour * 72),
		},
		{
			Id: backedloantypes.BidId{
				NftId: &backedloantypes.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
			},
			BidAmount:          types.NewInt64Coin("uatom", 200000000),
			DepositAmount:      types.NewInt64Coin("uatom", 70000000),
			DepositLendingRate: types.NewDecWithPrec(1, 1), // 10%
			BiddingPeriod:      time.Now().Add(time.Hour * 24),
			Borrowings: []backedloantypes.Borrowing{
				{
					// borrow 100% of deposit
					Amount:  types.NewInt64Coin("uatom", 70000000),
					StartAt: time.Now(),
				},
			},
		},
		{
			Id: backedloantypes.BidId{
				NftId: &backedloantypes.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi1y3t7sp0nfe2nfda7r9gf628g6ym6e7d44evfv6",
			},
			BidAmount:          types.NewInt64Coin("uatom", 200000000),
			DepositAmount:      types.NewInt64Coin("uatom", 80000000),
			DepositLendingRate: types.NewDecWithPrec(1, 1), // 10%
			BiddingPeriod:      time.Now().Add(time.Hour * 24),
			Borrowings: []backedloantypes.Borrowing{
				{
					// borrow 100% of deposit
					Amount:  types.NewInt64Coin("uatom", 80000000),
					StartAt: time.Now(),
				},
			},
		},
	}

	testCases := []struct {
		name      string
		oldBidId  backedloantypes.BidId
		newBid    backedloantypes.NftBid
		expResult bool
	}{
		{
			name: "able to re-bid",
			oldBidId: backedloantypes.BidId{
				NftId: &backedloantypes.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			newBid: backedloantypes.NftBid{
				Id: backedloantypes.BidId{
					NftId: &backedloantypes.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				BidAmount:          types.NewInt64Coin("uatom", 200000000),
				DepositAmount:      types.NewInt64Coin("uatom", 50000000),
				DepositLendingRate: types.NewDecWithPrec(5, 2), // 5%
				BiddingPeriod:      time.Now().Add(time.Hour * 72),
			},
			expResult: true,
		},
		{
			name: "unable to re-bid",
			oldBidId: backedloantypes.BidId{
				NftId: &backedloantypes.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			newBid: backedloantypes.NftBid{
				Id: backedloantypes.BidId{
					NftId: &backedloantypes.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				BidAmount:          types.NewInt64Coin("uatom", 200),
				DepositAmount:      types.NewInt64Coin("uatom", 50),
				DepositLendingRate: types.NewDecWithPrec(5, 3), // 0.5%
				BiddingPeriod:      time.Now().Add(time.Hour * 72),
			},
			expResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := backedloantypes.IsAbleToReBid(bids, tc.oldBidId, tc.newBid)
			if result != tc.expResult {
				t.Errorf("expected %v, got %v", tc.expResult, result)
			}
		})
	}
}
