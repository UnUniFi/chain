package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func TestMinSettlementAmount(t *testing.T) {
	testCases := []struct {
		name      string
		bids      []types.NftBid
		expResult sdk.Coin
	}{
		{
			name:      "empty bid",
			bids:      []types.NftBid{},
			expResult: sdk.Coin{},
		},
		{
			name: "one bid",
			bids: []types.NftBid{
				{
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 30),
				},
			},
			expResult: sdk.NewInt64Coin("uatom", 30),
		},
		{
			name: "two bids, totalDepositAmount < bidAmount",
			bids: []types.NftBid{
				{
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 30),
				},
				{
					BidAmount:     sdk.NewInt64Coin("uatom", 200),
					DepositAmount: sdk.NewInt64Coin("uatom", 50),
				},
			},
			expResult: sdk.NewInt64Coin("uatom", 80),
		},
		{
			name: "three bids & bidAmount < totalDepositAmount",
			bids: []types.NftBid{
				{
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 30),
				},
				{
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 40),
				},
				{
					BidAmount:     sdk.NewInt64Coin("uatom", 105),
					DepositAmount: sdk.NewInt64Coin("uatom", 45),
				},
			},
			expResult: sdk.NewInt64Coin("uatom", 105),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := types.MinSettlementAmount(tc.bids)
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
	now := time.Now()
	testCases := []struct {
		name      string
		bids      []types.NftBid
		expResult types.NftBid
	}{
		{
			name:      "empty bid",
			bids:      []types.NftBid{},
			expResult: types.NftBid{},
		},
		{
			name: "one bid, paid",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 30),
					PaidAmount:    sdk.NewInt64Coin("uatom", 70),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
					Borrow:        types.Borrowing{},
				},
			},
			expResult: types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
			},
		},
		{
			name: "one bid, paid",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 30),
					PaidAmount:    sdk.NewInt64Coin("uatom", 0),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
				},
			},
			expResult: types.NftBid{},
		},
		{
			name: "two bids, paid",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 30),
					PaidAmount:    sdk.NewInt64Coin("uatom", 70),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
				},
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 200),
					DepositAmount: sdk.NewInt64Coin("uatom", 50),
					PaidAmount:    sdk.NewInt64Coin("uatom", 150),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
				},
			},
			expResult: types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
				},
			},
		},
		{
			name: "two bids, unpaid top bid",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 30),
					PaidAmount:    sdk.NewInt64Coin("uatom", 70),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
				},
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 200),
					DepositAmount: sdk.NewInt64Coin("uatom", 50),
					PaidAmount:    sdk.NewInt64Coin("uatom", 0),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
				},
			},
			expResult: types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
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
			bidsSortedByDeposit := types.NftBids(tc.bids).SortHigherDeposit()
			result, err := types.LiquidationBid(bidsSortedByDeposit, now)
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
		bids      []types.NftBid
		winBid    types.NftBid
		expResult []int
	}{
		{
			name: "two bids, paid",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 30),
					PaidAmount:    sdk.NewInt64Coin("uatom", 70),
				},
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 200),
					DepositAmount: sdk.NewInt64Coin("uatom", 50),
					PaidAmount:    sdk.NewInt64Coin("uatom", 150),
				},
			},
			winBid: types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
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
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100),
					DepositAmount: sdk.NewInt64Coin("uatom", 30),
					PaidAmount:    sdk.NewInt64Coin("uatom", 70),
				},
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 200),
					DepositAmount: sdk.NewInt64Coin("uatom", 50),
					PaidAmount:    sdk.NewInt64Coin("uatom", 0),
				},
			},
			winBid: types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
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
		bidsSortedByDeposit := types.NftBids(tc.bids).SortHigherDeposit()
		forfeitedBids, refundBids := types.ForfeitedBidsAndRefundBids(bidsSortedByDeposit, tc.winBid)
		if tc.expResult[0] != len(forfeitedBids) {
			t.Errorf("forfeitedBids expected length %d, got %d", tc.expResult[0], len(forfeitedBids))
		}
		if tc.expResult[1] != len(refundBids) {
			t.Errorf("refundBids expected length %d, got %d", tc.expResult[1], len(refundBids))
		}
	}
}

func TestExpectedRepayAmount(t *testing.T) {
	now := time.Now()
	nextMonth := now.Add(time.Hour * 24 * 30)
	testCases := []struct {
		name       string
		bids       []types.NftBid
		borrowBids []types.BorrowBid
		expResult  sdk.Coin
	}{
		{
			name:       "empty bid",
			bids:       []types.NftBid{},
			borrowBids: []types.BorrowBid{},
			expResult:  sdk.Coin{},
		},
		{
			name: "one bid",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 30000000),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
					// Additional 1 minute for time error correction
					ExpiryAt: nextMonth,
				},
			},
			borrowBids: []types.BorrowBid{
				{
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					Amount: sdk.NewInt64Coin("uatom", 200000000),
				},
			},
			// 30000000 * 0.1 / 365 / 24 = 342.46
			expResult: sdk.NewInt64Coin("uatom", 30000342),
		},
		{
			name: "2 bid & over borrow",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 30000000),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
					ExpiryAt:      nextMonth,
				},
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 200000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 30000000),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
					ExpiryAt:      nextMonth,
				},
			},
			borrowBids: []types.BorrowBid{
				{
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					Amount: sdk.NewInt64Coin("uatom", 20000000),
				},
				{
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					Amount: sdk.NewInt64Coin("uatom", 70000000),
				},
			},
			// if over borrow, borrow amount is deposit amount
			// 50000000 * 0.1 / 365 / 24 = 570.7762
			expResult: sdk.NewInt64Coin("uatom", 50000570),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := types.ExpectedRepayAmount(tc.bids, tc.borrowBids, now)
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
	now := time.Now()
	nextMonth := now.Add(time.Hour * 24 * 30)
	testCases := []struct {
		name      string
		bids      []types.NftBid
		expResult sdk.Coin
	}{
		{
			name:      "empty bid",
			bids:      []types.NftBid{},
			expResult: sdk.Coin{},
		},
		{
			name: "one bid",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 30000000),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
					// Additional 1 minute for time error correction
					ExpiryAt: nextMonth,
					Borrow: types.Borrowing{
						Amount:       sdk.NewInt64Coin("uatom", 20000000),
						LastRepaidAt: now,
					},
				},
			},
			// 20000000 * 0.1 / 365 / 24 = 228.31
			expResult: sdk.NewInt64Coin("uatom", 20000228),
		},
		{
			name: "2 bid",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 30000000),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
					// Additional 1 minute for time error correction
					ExpiryAt: nextMonth,
					Borrow: types.Borrowing{
						Amount:       sdk.NewInt64Coin("uatom", 20000000),
						LastRepaidAt: now,
					},
				},
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 200000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 40000000),
					InterestRate:  sdk.NewDecWithPrec(1, 1),
					// Additional 1 minute for time error correction
					ExpiryAt: nextMonth,
					Borrow: types.Borrowing{
						Amount:       sdk.NewInt64Coin("uatom", 40000000),
						LastRepaidAt: now,
					},
				},
			},
			// 60000000 * 0.1 / 365 / 24 = 685.89
			expResult: sdk.NewInt64Coin("uatom", 60000685),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := types.ExistRepayAmount(tc.bids)
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
	now := time.Now()
	bids := []types.NftBid{
		{
			Id: types.BidId{
				NftId: &types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			BidAmount:     sdk.NewInt64Coin("uatom", 100000000),
			DepositAmount: sdk.NewInt64Coin("uatom", 30000000),
			InterestRate:  sdk.NewDecWithPrec(5, 2), // 5%
			ExpiryAt:      now.Add(time.Hour * 72),
		},
		{
			Id: types.BidId{
				NftId: &types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
			},
			BidAmount:     sdk.NewInt64Coin("uatom", 200000000),
			DepositAmount: sdk.NewInt64Coin("uatom", 80000000),
			InterestRate:  sdk.NewDecWithPrec(1, 1), // 10%
			ExpiryAt:      now.Add(time.Hour * 24),
		},
	}
	testCases := []struct {
		name       string
		borrowBids []types.BorrowBid
		expResult  bool
	}{
		{
			name: "able to borrow",
			borrowBids: []types.BorrowBid{
				{
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					Amount: sdk.NewInt64Coin("uatom", 30000000),
				},
				{
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					Amount: sdk.NewInt64Coin("uatom", 20000000),
				},
			},
			expResult: true,
		},
		{
			name: "unable to borrow",
			borrowBids: []types.BorrowBid{
				{
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					Amount: sdk.NewInt64Coin("uatom", 30000000),
				},
				{
					Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					Amount: sdk.NewInt64Coin("uatom", 80000000),
				},
			},
			expResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := types.IsAbleToBorrow(bids, tc.borrowBids, now)
			if result != tc.expResult {
				t.Errorf("expected %v, got %v", tc.expResult, result)
			}
		})
	}
}

func TestIsAbleToCancelBid(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name      string
		bids      []types.NftBid
		cancelBid types.BidId
		expResult bool
	}{
		{
			name: "able to cancel",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 30000000),
					InterestRate:  sdk.NewDecWithPrec(5, 2), // 5%
					ExpiryAt:      now.Add(time.Hour * 72),
					Borrow: types.Borrowing{
						Amount:       sdk.NewInt64Coin("uatom", 0),
						LastRepaidAt: now,
					},
				},
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 200000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 80000000),
					InterestRate:  sdk.NewDecWithPrec(1, 1), // 10%
					ExpiryAt:      now.Add(time.Hour * 24),
					Borrow: types.Borrowing{
						Amount:       sdk.NewInt64Coin("uatom", 0),
						LastRepaidAt: now,
					},
				},
			},
			cancelBid: types.BidId{
				NftId: &types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			expResult: true,
		},
		{
			name: "unable to cancel",
			bids: []types.NftBid{
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 100000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 30000000),
					InterestRate:  sdk.NewDecWithPrec(5, 2), // 5%
					ExpiryAt:      now.Add(time.Hour * 72),
					Borrow: types.Borrowing{
						Amount:       sdk.NewInt64Coin("uatom", 0),
						LastRepaidAt: now,
					},
				},
				{
					Id: types.BidId{
						NftId: &types.NftIdentifier{
							ClassId: "a10",
							NftId:   "a10",
						},
						Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
					},
					BidAmount:     sdk.NewInt64Coin("uatom", 200000000),
					DepositAmount: sdk.NewInt64Coin("uatom", 80000000),
					InterestRate:  sdk.NewDecWithPrec(1, 1), // 10%
					ExpiryAt:      now.Add(time.Hour * 24),
					Borrow: types.Borrowing{
						// borrow 100% of deposit
						Amount:       sdk.NewInt64Coin("uatom", 80000000),
						LastRepaidAt: now,
					},
				},
			},
			cancelBid: types.BidId{
				NftId: &types.NftIdentifier{
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
			result := types.IsAbleToCancelBid(tc.cancelBid, tc.bids)
			if result != tc.expResult {
				t.Errorf("expected %v, got %v", tc.expResult, result)
			}
		})
	}
}

func TestIsAbleToReBid(t *testing.T) {
	now := time.Now()
	bids := []types.NftBid{
		{
			Id: types.BidId{
				NftId: &types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			BidAmount:     sdk.NewInt64Coin("uatom", 100000000),
			DepositAmount: sdk.NewInt64Coin("uatom", 30000000),
			InterestRate:  sdk.NewDecWithPrec(5, 2), // 5%
			ExpiryAt:      now.Add(time.Hour * 72),
		},
		{
			Id: types.BidId{
				NftId: &types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
			},
			BidAmount:     sdk.NewInt64Coin("uatom", 200000000),
			DepositAmount: sdk.NewInt64Coin("uatom", 70000000),
			InterestRate:  sdk.NewDecWithPrec(1, 1), // 10%
			ExpiryAt:      now.Add(time.Hour * 24),
			Borrow: types.Borrowing{
				Amount:       sdk.NewInt64Coin("uatom", 70000000),
				LastRepaidAt: now,
			},
		},
		{
			Id: types.BidId{
				NftId: &types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi1y3t7sp0nfe2nfda7r9gf628g6ym6e7d44evfv6",
			},
			BidAmount:     sdk.NewInt64Coin("uatom", 200000000),
			DepositAmount: sdk.NewInt64Coin("uatom", 80000000),
			InterestRate:  sdk.NewDecWithPrec(1, 1), // 10%
			ExpiryAt:      now.Add(time.Hour * 24),
			Borrow: types.Borrowing{
				// borrow 100% of deposit
				Amount:       sdk.NewInt64Coin("uatom", 80000000),
				LastRepaidAt: now,
			},
		},
	}

	testCases := []struct {
		name      string
		oldBidId  types.BidId
		newBid    types.NftBid
		expResult bool
	}{
		{
			name: "able to re-bid",
			oldBidId: types.BidId{
				NftId: &types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			newBid: types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				BidAmount:     sdk.NewInt64Coin("uatom", 200000000),
				DepositAmount: sdk.NewInt64Coin("uatom", 50000000),
				InterestRate:  sdk.NewDecWithPrec(5, 2), // 5%
				ExpiryAt:      now.Add(time.Hour * 72),
			},
			expResult: true,
		},
		{
			name: "unable to re-bid",
			oldBidId: types.BidId{
				NftId: &types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
			},
			newBid: types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				BidAmount:     sdk.NewInt64Coin("uatom", 200),
				DepositAmount: sdk.NewInt64Coin("uatom", 50),
				InterestRate:  sdk.NewDecWithPrec(5, 3), // 0.5%
				ExpiryAt:      now.Add(time.Hour * 72),
			},
			expResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := types.IsAbleToReBid(bids, tc.oldBidId, tc.newBid)
			if result != tc.expResult {
				t.Errorf("expected %v, got %v", tc.expResult, result)
			}
		})
	}
}
