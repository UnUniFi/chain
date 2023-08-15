package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func TestIsBorrowed(t *testing.T) {
	testCases := []struct {
		name      string
		bid       types.Bid
		expResult bool
	}{
		{
			"Exist Borrowing",
			types.Bid{
				Id: types.BidId{
					NftId: &types.NftId{
						ClassId: "a10",
						TokenId: "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				Loan: types.Loan{
					Amount:       sdk.NewCoin("uguu", sdk.NewInt(100)),
					LastRepaidAt: time.Now(),
				},
			},
			true,
		},
		{
			"No borrowing",
			types.Bid{
				Id: types.BidId{
					NftId: &types.NftId{
						ClassId: "a10",
						TokenId: "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				Loan: types.Loan{
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
			result := tc.bid.IsBorrowed()
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
		bid       types.Bid
		lendCoin  sdk.Coin
		startTime time.Time
		endTime   time.Time
		expResult sdk.Coin
	}{
		{
			"Interest test 1 month",
			types.Bid{
				Id: types.BidId{
					NftId: &types.NftId{
						ClassId: "a10",
						TokenId: "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				InterestRate: sdk.NewDecWithPrec(1, 1),
				Loan: types.Loan{
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
			types.Bid{
				Id: types.BidId{
					NftId: &types.NftId{
						ClassId: "a10",
						TokenId: "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				InterestRate: sdk.NewDecWithPrec(1, 1),
				Loan: types.Loan{
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

func TestRepayInfo(t *testing.T) {
	now := time.Now()
	nextYear := time.Now().Add(time.Hour * 24 * 365)
	testCases := []struct {
		name          string
		bid           types.Bid
		repayAmount   sdk.Coin
		repaymentTime time.Time
		expResult     types.RepayInfo
	}{
		{
			"Repay partial",
			types.Bid{
				Id: types.BidId{
					NftId: &types.NftId{
						ClassId: "a10",
						TokenId: "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				InterestRate: sdk.NewDecWithPrec(1, 1),
				Loan: types.Loan{
					Amount:       sdk.NewCoin("uguu", sdk.NewInt(1000000)),
					LastRepaidAt: now,
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(200000)),
			nextYear,
			types.RepayInfo{
				RepaidAmount:         sdk.NewCoin("uguu", sdk.NewInt(200000)),
				RepaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(105171)),
				// 1105171 - 200000 = 905171
				RemainingAmount: sdk.NewCoin("uguu", sdk.NewInt(905171)),
				LastRepaidAt:    nextYear,
			},
		},
		{
			"Repay over amount",
			types.Bid{
				Id: types.BidId{
					NftId: &types.NftId{
						ClassId: "a10",
						TokenId: "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				InterestRate: sdk.NewDecWithPrec(1, 1),
				Loan: types.Loan{
					Amount:       sdk.NewCoin("uguu", sdk.NewInt(1000000)),
					LastRepaidAt: now,
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(1200000)),
			nextYear,
			types.RepayInfo{
				RepaidAmount:         sdk.NewCoin("uguu", sdk.NewInt(1105171)),
				RepaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(105171)),
				RemainingAmount:      sdk.NewCoin("uguu", sdk.NewInt(0)),
				LastRepaidAt:         nextYear,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bid.RepayInfo(tc.repayAmount, tc.repaymentTime)
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestIsPaidSalePrice(t *testing.T) {
	testCases := []struct {
		name      string
		bid       types.Bid
		expResult bool
	}{
		{
			"Paid",
			types.Bid{
				Id: types.BidId{
					NftId: &types.NftId{
						ClassId: "a10",
						TokenId: "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				Price:      sdk.NewCoin("uguu", sdk.NewInt(1000000)),
				Deposit:    sdk.NewCoin("uguu", sdk.NewInt(200000)),
				PaidAmount: sdk.NewCoin("uguu", sdk.NewInt(800000)),
			},
			true,
		},
		{
			"Not paid",
			types.Bid{
				Id: types.BidId{
					NftId: &types.NftId{
						ClassId: "a10",
						TokenId: "a10",
					},
					Bidder: "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				},
				Price:      sdk.NewCoin("uguu", sdk.NewInt(1000000)),
				Deposit:    sdk.NewCoin("uguu", sdk.NewInt(200000)),
				PaidAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
			},
			false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bid.IsPaidSalePrice()
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestTotalBorrowedAmount(t *testing.T) {
	testCases := []struct {
		name      string
		bids      types.NftBids
		expResult sdk.Coin
	}{
		{
			"Total borrow amount",
			types.NftBids{
				types.Bid{
					Loan: types.Loan{
						Amount: sdk.NewCoin("uguu", sdk.NewInt(1000000)),
					},
				},
				types.Bid{
					Loan: types.Loan{
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
			result := tc.bids.TotalBorrowedAmount()
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
				types.Bid{
					InterestRate: sdk.NewDecWithPrec(1, 1),
					Loan: types.Loan{
						Amount:       sdk.NewCoin("uguu", sdk.NewInt(1000000)),
						LastRepaidAt: now,
					},
				},
				types.Bid{
					InterestRate: sdk.NewDecWithPrec(2, 1),
					Loan: types.Loan{
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
