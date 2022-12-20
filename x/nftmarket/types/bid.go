package types

import (
	"sort"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m NftBid) Equal(b NftBid) bool {
	return m.Bidder == b.Bidder && m.NftId == b.NftId && m.BidAmount.Equal(b)
}
func (m NftBid) GetIdToByte() []byte {
	return NftBidBytes(m.NftId.ClassId, m.NftId.NftId, m.Bidder)
}

// todo check test
func (m NftBid) GetIdToString() string {
	return string(NftBidBytes(m.NftId.ClassId, m.NftId.NftId, m.Bidder))
}

func (m NftBid) IsBorrowing() bool {
	return len(m.Borrowings) != 0
}

func (m NftBid) LiquidationAmount(endTime time.Time) sdk.Coin {
	liquidationAmount := sdk.NewCoin(m.DepositAmount.Denom, sdk.ZeroInt())
	for _, v := range m.Borrowings {
		liquidationAmount = liquidationAmount.Add(v.Amount)
		liquidationAmount = liquidationAmount.Add(CalcInterest(v.Amount, m.DepositLendingRate, v.StartAt, endTime))
		liquidationAmount = liquidationAmount.Sub(v.PaidInterestAmount)
	}
	return liquidationAmount
}

func (m NftBid) BorrowingAmount() sdk.Coin {
	liquidationAmount := sdk.NewCoin(m.DepositAmount.Denom, sdk.ZeroInt())
	for _, v := range m.Borrowings {
		liquidationAmount = liquidationAmount.Add(v.Amount)
		liquidationAmount = liquidationAmount.Sub(v.PaidInterestAmount)
	}
	return liquidationAmount
}

func (m NftBid) BorrowableAmount() sdk.Coin {
	borrowableAmount := m.DepositAmount
	borrowabingAmount := sdk.NewCoin(m.DepositAmount.Denom, sdk.ZeroInt())
	for _, v := range m.Borrowings {
		borrowabingAmount = borrowabingAmount.Add(v.Amount)
		borrowabingAmount = borrowabingAmount.Sub(v.PaidInterestAmount)
	}
	return borrowableAmount.Sub(borrowabingAmount)
}

func CalcInterest(lendCoin sdk.Coin, lendingRate sdk.Dec, start, end time.Time) sdk.Coin {
	interest := sdk.ZeroInt()
	return sdk.NewCoin(lendCoin.Denom, interest)
}

func (m NftBid) GetFullPaidAmount() sdk.Coin {
	return m.PaidAmount.Add(m.DepositAmount)
}

func (m NftBid) IsPaidBidAmount() bool {
	fullPaidAmount := m.GetFullPaidAmount()
	return fullPaidAmount.Equal(m.BidAmount)
}

type NftBids []NftBid

func (m NftBids) SortBorrowing() NftBids {
	return m.SortLowerLendingRate()
}

func (m NftBids) SortLiquidation() NftBids {
	return m.SortLowerDepositAmount()
}

func (m NftBids) SortLowerLendingRate() NftBids {
	dest := NftBids{}
	dest = append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].DepositLendingRate.LT(dest[j].DepositLendingRate)
	})
	return dest
}

func (m NftBids) SortLowerDepositAmount() NftBids {
	dest := NftBids{}
	dest = append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].DepositAmount.IsLT(dest[j].DepositAmount)
	})
	return dest
}

func (m NftBids) GetHighestBid() NftBid {

	highestBidder := NftBid{
		BidAmount: sdk.NewCoin(m[0].BidAmount.Denom, sdk.ZeroInt()),
	}
	for _, bid := range m {
		if highestBidder.BidAmount.IsLT(bid.BidAmount) {
			highestBidder = bid
		}
	}

	return highestBidder
}
