package types_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
)

func TestBidEqual(t *testing.T) {
	testCases := []struct {
		name      string
		bida      types.NftBid
		bidb      types.NftBid
		expResult bool
	}{
		{
			"equal pattern 1",
			types.NftBid{
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
			types.NftBid{
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
			true,
		},
		{
			"equal pattern 2 difference deposit amount",
			types.NftBid{
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
			types.NftBid{
				NftId: types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder:             "ununifi155u042u8wk3al32h3vzxu989jj76k4zcu44v6w",
				BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
				DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(52)),
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
			true,
		},
		{
			"difference bidder",
			types.NftBid{
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
			types.NftBid{
				NftId: types.NftIdentifier{
					ClassId: "a10",
					NftId:   "a10",
				},
				Bidder:             "ununifi1v0h8j7x7kfys29kj4uwdudcc9y0nx6twwxahla",
				BidAmount:          sdk.NewCoin("uguu", sdk.NewInt(100)),
				DepositAmount:      sdk.NewCoin("uguu", sdk.NewInt(52)),
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
			false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bida.Equal(tc.bidb)
			if tc.expResult == result {
			} else {
				t.Error(tc.name, "not expect result")
			}
		})
	}
}

func TestCalcInterest(t *testing.T) {
	testCases := []struct {
		name        string
		bida        types.NftBid
		lendCoin    sdk.Coin
		lendingRate sdk.Dec
		start       time.Time
		end         time.Time
		expResult   sdk.Coin
	}{
		{
			"one year interest",
			types.NftBid{
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
			sdk.NewCoin("uguu", sdk.NewInt(100)),
			sdk.MustNewDecFromStr("0.1"),
			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * 365),
			sdk.NewCoin("uguu", sdk.NewInt(10)),
		},
		{
			"one day interest",
			types.NftBid{
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
			sdk.NewCoin("uguu", sdk.NewInt(100000)),
			sdk.MustNewDecFromStr("0.1"),
			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Hour * 24),
			sdk.NewCoin("uguu", sdk.NewInt(27)),
		},
		{
			"check round off",
			types.NftBid{
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
			sdk.NewCoin("uguu", sdk.NewInt(100032)),
			sdk.MustNewDecFromStr("0.1"),
			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Hour * 24),
			sdk.NewCoin("uguu", sdk.NewInt(27)),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bida.CalcInterest(tc.lendCoin, tc.lendingRate, tc.start, tc.end)
			// fmt.Println(result)
			// fmt.Println(tc.expResult)
			if tc.expResult.Equal(result) {
			} else {
				t.Error(tc.name, "not expect result")
			}
		})
	}
}

func TestRepayThenGetReceipt(t *testing.T) {
	testCases := []struct {
		name            string
		borrowing       types.Borrowing
		payAmount       sdk.Coin
		payTime         time.Time
		interest        sdk.Coin
		expectReceipt   types.RepayReceipt
		expectBorrowing types.Borrowing
	}{
		{
			"pay all principal and interest",
			types.Borrowing{
				Amount:             sdk.NewCoin("uguu", sdk.NewInt(100)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(1)),
				StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			sdk.NewCoin("uguu", sdk.NewInt(201)),
			time.Now(),
			sdk.NewCoin("uguu", sdk.NewInt(2)),
			types.RepayReceipt{
				Charge:             sdk.NewCoin("uguu", sdk.NewInt(100)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(1)),
			},
			types.Borrowing{
				Amount:             sdk.NewCoin("uguu", sdk.NewInt(0)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(2)),
				StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			"pay all interest and pay part of principal",
			types.Borrowing{
				Amount:             sdk.NewCoin("uguu", sdk.NewInt(100)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(1)),
				StartAt:            time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			sdk.NewCoin("uguu", sdk.NewInt(20)),
			time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			sdk.NewCoin("uguu", sdk.NewInt(12)),
			types.RepayReceipt{
				Charge:             sdk.NewCoin("uguu", sdk.NewInt(0)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(11)),
			},
			types.Borrowing{
				Amount:             sdk.NewCoin("uguu", sdk.NewInt(91)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
				StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			"pay all interest",
			types.Borrowing{
				Amount:             sdk.NewCoin("uguu", sdk.NewInt(100)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(1)),
				StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			sdk.NewCoin("uguu", sdk.NewInt(10)),
			time.Now(),
			sdk.NewCoin("uguu", sdk.NewInt(11)),
			types.RepayReceipt{
				Charge:             sdk.NewCoin("uguu", sdk.NewInt(0)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(10)),
			},
			types.Borrowing{
				Amount:             sdk.NewCoin("uguu", sdk.NewInt(100)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(11)),
				StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			"pay part of the interest",
			types.Borrowing{
				Amount:             sdk.NewCoin("uguu", sdk.NewInt(100)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(0)),
				StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			sdk.NewCoin("uguu", sdk.NewInt(5)),
			time.Now(),
			sdk.NewCoin("uguu", sdk.NewInt(10)),
			types.RepayReceipt{
				Charge:             sdk.NewCoin("uguu", sdk.NewInt(0)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(5)),
			},
			types.Borrowing{
				Amount:             sdk.NewCoin("uguu", sdk.NewInt(100)),
				PaidInterestAmount: sdk.NewCoin("uguu", sdk.NewInt(5)),
				StartAt:            time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			f := func(result sdk.Coin) func(lendCoin sdk.Coin, start, end time.Time) sdk.Coin {
				return func(lendCoin sdk.Coin, start, end time.Time) sdk.Coin {
					return result
				}
			}
			result := tc.borrowing.RepayThenGetReceipt(tc.payAmount, tc.payTime, f(tc.interest))
			// fmt.Println(result)
			// fmt.Println(tc.expectReceipt)
			if tc.expectReceipt.Charge.Equal(result.Charge) && tc.expectReceipt.PaidInterestAmount.Equal(result.PaidInterestAmount) {
			} else {
				t.Error(tc.name, "not expect result")
			}
			if tc.expectBorrowing.Equal(tc.borrowing) {
			} else {
				fmt.Println("tc.borrowing")
				fmt.Println(tc.borrowing)
				fmt.Println("tc.expectBorrowing")
				fmt.Println(tc.expectBorrowing)
				t.Error(tc.name, "not expect borrow")
			}
		})
	}
}

func TestSortLiquidation(t *testing.T) {

	// test SortLiquidation in NftBids
	testCases := []struct {
		name   string
		bids   types.NftBids
		expect types.NftBids
	}{
		{
			"higher bidder",
			types.NftBids{
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(100)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(50)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(20)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(102)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(2)),
				},
			},
			types.NftBids{
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(20)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(102)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(2)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(100)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(50)),
				},
			},
		},
		{
			"higher deposit",
			types.NftBids{
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(50)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(20)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(2)),
				},
			},
			types.NftBids{
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(50)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(20)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(2)),
				},
			},
		},
		{
			"higher deposit and greater than qDash",
			types.NftBids{
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(100)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(50)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(102)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(20)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(2)),
				},
			},
			types.NftBids{
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(102)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(20)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(2)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(100)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(50)),
				},
			},
		},
		{
			"higher deposit and greater than qDash 2",
			types.NftBids{
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(2)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(102)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(20)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(100)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(50)),
				},
			},
			types.NftBids{
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(102)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(20)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(103)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(2)),
				},
				types.NftBid{
					BidAmount:     sdk.NewCoin("uguu", sdk.NewInt(100)),
					DepositAmount: sdk.NewCoin("uguu", sdk.NewInt(50)),
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bids.SortLiquidation()
			for i := 0; i < len(result); i++ {
				if result[i].BidAmount.Equal(tc.expect[i].BidAmount) && result[i].DepositAmount.Equal(tc.expect[i].DepositAmount) {
				} else {
					t.Error(tc.name, "not expect result")
					t.Log(i)
					t.Log(result[i].BidAmount, tc.expect[i].BidAmount)
					t.Log(result[i].DepositAmount, tc.expect[i].DepositAmount)
				}
			}
		})
	}

}

// test GetAverageBidAmount
func TestGetAverageBidAmount(t *testing.T) {
	testCases := []struct {
		name   string
		bids   types.NftBids
		expect sdk.Coin
	}{
		{
			"empty",
			types.NftBids{},
			sdk.Coin{},
		},
		{
			"one",
			types.NftBids{
				types.NftBid{
					BidAmount: sdk.NewCoin("uguu", sdk.NewInt(100)),
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(100)),
		},
		{
			"two",
			types.NftBids{
				types.NftBid{
					BidAmount: sdk.NewCoin("uguu", sdk.NewInt(100)),
				},
				types.NftBid{
					BidAmount: sdk.NewCoin("uguu", sdk.NewInt(200)),
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(150)),
		},
		{
			"three",
			types.NftBids{
				types.NftBid{
					BidAmount: sdk.NewCoin("uguu", sdk.NewInt(100)),
				},
				types.NftBid{
					BidAmount: sdk.NewCoin("uguu", sdk.NewInt(200)),
				},
				types.NftBid{
					BidAmount: sdk.NewCoin("uguu", sdk.NewInt(300)),
				},
			},
			sdk.NewCoin("uguu", sdk.NewInt(200)),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.bids.GetAverageBidAmount()
			if result.Equal(tc.expect) {
			} else {
				t.Error(tc.name, "not expect result")
				t.Log(result, tc.expect)
			}
		})
	}
}
