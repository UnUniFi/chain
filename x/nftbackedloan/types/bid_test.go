package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func TestIsBorrowing(t *testing.T) {
	testCases := []struct {
		name      string
		bid       types.NftBid
		expResult bool
	}{
		{
			"Exist Borrowing",
			types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				Borrow: types.Borrowing{
					Amount:       sdk.NewCoin("uguu", sdk.NewInt(100)),
					LastRepaidAt: time.Now(),
				},
			},
			true,
		},
		{
			"No borrowing",
			types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				Borrow: types.Borrowing{
					Amount:       sdk.NewCoin("uguu", sdk.NewInt(0)),
					LastRepaidAt: time.Now(),
				},
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bid.IsBorrowing()
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestCalcCompoundInterest(t *testing.T) {
	now := time.Now()
	nextMonth := time.Now().Add(time.Hour * 24 * 30)
	nextYear := time.Now().Add(time.Hour * 24 * 365)
	testCases := []struct {
		name      string
		bid       types.NftBid
		lendCoin  sdk.Coin
		startTime time.Time
		endTime   time.Time
		expResult sdk.Coin
	}{
		{
			"Interest test 1 month",
			types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				InterestRate: sdk.NewDecWithPrec(1, 1),
				Borrow: types.Borrowing{
					Amount:       sdk.NewCoin("uguu", sdk.NewInt(10000000)),
					LastRepaidAt: now,
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(1000000)),
			now,
			nextMonth,
			// 10000000 * e^(30/365 * 0.1) - 10000000 = 8253
			sdk.NewCoin("uguu", sdk.NewInt(8253)),
		},
		{
			"Interest test 1 year",
			types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				InterestRate: sdk.NewDecWithPrec(1, 1),
				Borrow: types.Borrowing{
					Amount:       sdk.NewCoin("uguu", sdk.NewInt(10000000)),
					LastRepaidAt: now,
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(1000000)),
			now,
			nextYear,
			// 10000000 * e^(0.1) - 10000000 = 105171
			sdk.NewCoin("uguu", sdk.NewInt(105171)),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bid.CalcCompoundInterest(tc.lendCoin, tc.startTime, tc.endTime)
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestRepaidResult(t *testing.T) {
	now := time.Now()
	nextYear := time.Now().Add(time.Hour * 24 * 365)
	testCases := []struct {
		name        string
		bid         types.NftBid
		repayAmount sdk.Coin
		payTime     time.Time
		expResult   types.RepayResult
	}{
		{
			"Repay partial",
			types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				InterestRate: sdk.NewDecWithPrec(1, 1),
				Borrow: types.Borrowing{
					Amount:       sdk.NewCoin("uguu", sdk.NewInt(1000000)),
					LastRepaidAt: now,
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(200000)),
			nextYear,
			types.RepayResult{
				RepaidAmount:         sdk.NewCoin("uguu", sdk.NewInt(200000)),
				RepaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(105171)),
				// 1105171 - 200000 = 905171
				RemainingBorrowAmount: sdk.NewCoin("uguu", sdk.NewInt(905171)),
				LastRepaidAt:          nextYear,
			},
		},
		{
			"Repay over amount",
			types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				InterestRate: sdk.NewDecWithPrec(1, 1),
				Borrow: types.Borrowing{
					Amount:       sdk.NewCoin("uguu", sdk.NewInt(1000000)),
					LastRepaidAt: now,
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(1200000)),
			nextYear,
			types.RepayResult{
				RepaidAmount:          sdk.NewCoin("uguu", sdk.NewInt(1105171)),
				RepaidInterestAmount:  sdk.NewCoin("uguu", sdk.NewInt(105171)),
				RemainingBorrowAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
				LastRepaidAt:          nextYear,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bid.RepaidResult(tc.repayAmount, tc.payTime)
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestIsPaidBidAmount(t *testing.T) {
	testCases := []struct {
		name      string
		bid       types.NftBid
		expResult bool
	}{
		{
			"Paid",
			types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(1000000)),
				DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(200000)),
				PaidAmount:    sdk.NewCoin("uguu", sdk.NewInt(800000)),
			},
			true,
		},
		{
			"Not paid",
			types.NftBid{
				Id: types.BidId{
					NftId: &types.NftIdentifier{
						ClassId: "a10",
						NftId:   "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(1000000)),
				DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(200000)),
				PaidAmount:    sdk.NewCoin("uguu", sdk.NewInt(0)),
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bid.IsPaidBidAmount()
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestTotalBorrowAmount(t *testing.T) {
	testCases := []struct {
		name      string
		bids      types.NftBids
		expResult sdk.Coin
	}{
		{
			"Total borrow amount",
			types.NftBids{
				types.NftBid{
					Borrow: types.Borrowing{
						Amount: sdk.NewCoin("uguu", sdk.NewInt(1000000)),
					},
				},
				types.NftBid{
					Borrow: types.Borrowing{
						Amount: sdk.NewCoin("uguu", sdk.NewInt(2000000)),
					},
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(3000000)),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bids.TotalBorrowAmount()
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestTotalCompoundInterest(t *testing.T) {
	now := time.Now()
	nextYear := time.Now().Add(time.Hour * 24 * 365)
	testCases := []struct {
		name      string
		bids      types.NftBids
		expResult sdk.Coin
	}{
		{
			"Total borrow amount",
			types.NftBids{
				types.NftBid{
					InterestRate: sdk.NewDecWithPrec(1, 1),
					Borrow: types.Borrowing{
						Amount:       sdk.NewCoin("uguu", sdk.NewInt(1000000)),
						LastRepaidAt: now,
					},
				},
				types.NftBid{
					InterestRate: sdk.NewDecWithPrec(2, 1),
					Borrow: types.Borrowing{
						Amount:       sdk.NewCoin("uguu", sdk.NewInt(1000000)),
						LastRepaidAt: now,
					},
				},
			},
			// 105171 + 221403 = 326574
			sdk.NewCoin("uguu", sdk.NewInt(326574)),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bids.TotalCompoundInterest(nextYear)
			require.Equal(t, tc.expResult, result)
		})
	}
}
