package types

import (
	"sort"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m NftBid) Equal(b NftBid) bool {
	return m.Bidder == b.Bidder && m.NftId == b.NftId && m.BidAmount.Equal(b.BidAmount)
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
		liquidationAmount = liquidationAmount.Add(m.CalcInterest(v.Amount, m.DepositLendingRate, v.StartAt, endTime))
		liquidationAmount = liquidationAmount.Sub(v.PaidInterestAmount)
	}
	return liquidationAmount
}

func (m NftBid) BorrowingAmount() sdk.Coin {
	BorrowingAmount := sdk.NewCoin(m.DepositAmount.Denom, sdk.ZeroInt())
	for _, v := range m.Borrowings {
		BorrowingAmount = BorrowingAmount.Add(v.Amount)
		BorrowingAmount = BorrowingAmount.Sub(v.PaidInterestAmount)
	}
	return BorrowingAmount
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

func (m NftBid) CalcInterest(lendCoin sdk.Coin, lendingRate sdk.Dec, start, end time.Time) sdk.Coin {
	lendingDuration := end.Sub(start)
	oneYearDays := sdk.NewInt(365)
	oneDayHours := sdk.NewInt(24)

	yearInterest := lendingRate.Mul(sdk.NewDecFromInt(lendCoin.Amount))
	durationUnitsYear := sdk.NewDecFromInt(sdk.NewInt(int64(lendingDuration.Hours()))).QuoInt(oneDayHours).QuoInt(oneYearDays)

	result := durationUnitsYear.Mul(yearInterest)
	return sdk.NewCoin(lendCoin.Denom, result.RoundInt())
}

func (m NftBid) CalcInterestF() func(lendCoin sdk.Coin, start, end time.Time) sdk.Coin {
	f := func(rate sdk.Dec) func(lendCoin sdk.Coin, start, end time.Time) sdk.Coin {
		return func(lendCoin sdk.Coin, start, end time.Time) sdk.Coin {
			return m.CalcInterest(lendCoin, rate, start, end)
		}
	}
	return f(m.DepositLendingRate)
}

func (m NftBid) FullPaidAmount() sdk.Coin {
	return m.PaidAmount.Add(m.DepositAmount)
}

func (m NftBid) IsPaidBidAmount() bool {
	fullPaidAmount := m.FullPaidAmount()
	return fullPaidAmount.Equal(m.BidAmount)
}

func (m NftBid) CanCancel() bool {
	return !m.IsBorrowing()
}

func (m NftBid) CanReBid() bool {
	return !m.IsBorrowing()
}

func (m NftBid) IsNil() bool {
	return m.Bidder == ""
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

func (m NftBids) GetBidByBidder(bidder string) NftBid {
	for _, bid := range m {
		if bid.Bidder == bidder {
			return bid
		}
	}
	return NftBid{}
}

// todo: add proto then use it
type RepayReceipt struct {
	Charge             sdk.Coin
	PaidInterestAmount sdk.Coin
}

func (m *Borrowing) RepayThenGetReceipt(payAmount sdk.Coin, payTime time.Time, calcInterestF func(lendCoin sdk.Coin, start, end time.Time) sdk.Coin) RepayReceipt {
	principal := m.Amount
	interest := calcInterestF(principal, m.StartAt, payTime)
	interest = interest.Sub(m.PaidInterestAmount)
	paidInterestAmount := sdk.NewCoin(principal.Denom, sdk.ZeroInt())
	total := sdk.NewCoin(principal.Denom, sdk.ZeroInt())
	total = total.Add(principal)
	total = total.Add(interest)
	// bigger msg Amount
	if payAmount.IsGTE(total) {
		// bid.InterestAmount = bid.InterestAmount.Add(interest)
		payAmount = payAmount.Sub(total)
		paidInterestAmount = paidInterestAmount.Add(interest)
		m.Amount.Amount = sdk.ZeroInt()
		m.PaidInterestAmount = m.PaidInterestAmount.Add(paidInterestAmount)
	} else {
		// bigger total Amount
		if payAmount.IsGTE(interest) {
			// can paid interest
			if payAmount.Amount.GT(interest.Amount) {
				// all paid interest and part paid principal

				// bid.InterestAmount = bid.InterestAmount.Add(interest)
				payAmount = payAmount.Sub(interest)
				m.Amount = principal.Sub(payAmount)
				m.PaidInterestAmount.Amount = sdk.ZeroInt()
				m.StartAt = payTime

				payAmount.Amount = sdk.ZeroInt()
				paidInterestAmount = interest
			} else {
				// all paid interest
				// bid.InterestAmount = bid.InterestAmount.Add(interest)
				// m.PaidInterestAmount = m.PaidInterestAmount.Add(interest)
				m.PaidInterestAmount = m.PaidInterestAmount.Add(interest)
				paidInterestAmount = paidInterestAmount.Add(interest)
				payAmount = payAmount.Sub(interest)
			}
		} else {
			// can not paid interest
			// bid.InterestAmount.Add(payAmount)
			paidInterestAmount = payAmount
			m.PaidInterestAmount = m.PaidInterestAmount.Add(payAmount)
			payAmount.Amount = sdk.ZeroInt()
		}
		payAmount.Amount = sdk.ZeroInt()
	}
	return RepayReceipt{
		PaidInterestAmount: paidInterestAmount,
		Charge:             payAmount,
	}
}

func (m Borrowing) IsAllRepaid() bool {
	return m.Amount.IsZero()
}

func (a Borrowing) Equal(b Borrowing) bool {
	return a.Amount.Equal(b.Amount) &&
		a.PaidInterestAmount.Equal(b.PaidInterestAmount) &&
		a.StartAt.Location() == b.StartAt.Location()
}
