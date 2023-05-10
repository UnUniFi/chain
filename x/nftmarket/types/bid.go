package types

import (
	"sort"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m NftBid) Equal(b NftBid) bool {
	return m.Bidder == b.Bidder && m.NftId == b.NftId && m.BidAmount.Equal(b.BidAmount)
}
func (m NftBid) IsLT(b NftBid) bool {
	if b.BidAmount.IsLTE(m.BidAmount) {
		return false
	}
	if b.DepositAmount.IsLTE(m.DepositAmount) {
		return false
	}
	if b.DepositLendingRate.GTE(m.DepositLendingRate) {
		return false
	}

	return true
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

func (m NftBid) TotalInterestAmount(endTime time.Time) sdk.Coin {
	totalInterestAmount := sdk.NewCoin(m.InterestAmount.Denom, m.InterestAmount.Amount)
	for _, v := range m.Borrowings {
		totalInterestAmount = totalInterestAmount.Add(m.CalcInterest(v.Amount, m.DepositLendingRate, v.StartAt, endTime))
	}
	return totalInterestAmount
}

func (m NftBid) TotalInterestAmountDec(endTime time.Time) sdk.DecCoin {
	totalInterestAmount := sdk.NewDecCoin(m.InterestAmount.Denom, m.InterestAmount.Amount)
	for _, v := range m.Borrowings {
		interest := m.CalcInterest(v.Amount, m.DepositLendingRate, v.StartAt, endTime)
		totalInterestAmount = totalInterestAmount.Add(sdk.NewDecCoin(interest.Denom, interest.Amount))
	}
	return totalInterestAmount
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

func (m NftBids) SortRepay() NftBids {
	return m.SortHigherLendingRate()
}

func (m NftBids) SortLiquidation() NftBids {
	return m.SortDepositAboveAvgBid()
}

func (m NftBids) SortLowerLendingRate() NftBids {
	dest := NftBids{}
	dest = append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].DepositLendingRate.LT(dest[j].DepositLendingRate)
	})
	return dest
}

func (m NftBids) SortHigherLendingRate() NftBids {
	dest := NftBids{}
	dest = append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].DepositLendingRate.GT(dest[j].DepositLendingRate)
	})
	return dest
}

func (m NftBids) SortDepositAboveAvgBid() NftBids {
	dest := NftBids{}
	if len(m) == 0 {
		return dest
	}
	qDash := m.GetAverageBidAmount()
	dest = append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		if dest[i].BidAmount.IsLT(qDash) {
			return false
		}
		if dest[j].BidAmount.IsLT(qDash) {
			return true
		}
		return dest[i].DepositAmount.IsGTE(dest[j].DepositAmount)
	})
	return dest
}

func (m NftBids) SortLowerBiddingPeriod() NftBids {
	dest := NftBids{}
	dest = append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].BiddingPeriod.Before(dest[j].BiddingPeriod)
	})
	return dest
}

func (m NftBids) SortHigherDeposit() NftBids {
	dest := NftBids{}
	dest = append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].DepositAmount.IsGTE(dest[j].DepositAmount)
	})
	return dest
}

func (m NftBids) GetAverageBidAmount() sdk.Coin {
	if len(m) == 0 {
		return sdk.Coin{}
	}
	denom := m[0].BidAmount.Denom
	totalAmount := sdk.NewCoin(denom, sdk.ZeroInt())
	for _, bid := range m {
		totalAmount = totalAmount.Add(bid.BidAmount)
	}
	return sdk.NewCoin(denom, totalAmount.Amount.Quo(sdk.NewInt(int64(len(m)))))
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

func (m NftBids) RemoveBid(targetBid NftBid) NftBids {
	return m.RemoveBids(NftBids{targetBid})
}

func (m NftBids) RemoveBids(excludeBids NftBids) NftBids {
	excludeList := make(map[string]bool)
	for _, s := range excludeBids {
		excludeList[s.Bidder] = true
	}
	var newArr NftBids
	for _, s := range m {
		if !excludeList[s.Bidder] {
			newArr = append(newArr, s)
		}
	}
	return newArr
}

func (m NftBids) MakeExcludeExpiredBids(expiredBids NftBids) NftBids {
	return m.RemoveBids(expiredBids)
}

func (m NftBids) MakeBorrowedBidExcludeExpiredBids(borrowAmount sdk.Coin, start time.Time, expiredBids NftBids) NftBids {
	newBids := m.MakeExcludeExpiredBids(expiredBids)
	newBids.BorrowFromBids(borrowAmount, start)
	return newBids
}
func (m NftBids) MakeCollectBidsAndRefundBids() (NftBids, NftBids) {
	collectedBids := NftBids{}
	refundBids := NftBids{}
	existWinner := false
	for _, bid := range m {
		if existWinner {
			if bid.IsPaidBidAmount() {
				existWinner = true
				continue
			} else {
				collectedBids = append(collectedBids, bid)
			}
		}
		refundBids = append(refundBids, bid)
	}
	return collectedBids, refundBids
}

// get winner bid
func (m NftBids) GetWinnerBid() NftBid {
	for _, bid := range m {
		if bid.IsPaidBidAmount() {
			return bid
		}
	}
	return NftBid{}
}

func (m *NftBids) BorrowFromBids(borrowAmount sdk.Coin, start time.Time) {
	bids := []NftBid(*m)
	for i := 0; i < len(bids); i++ {
		bid := &bids[i]
		if borrowAmount.IsZero() {
			break
		}

		usableAmount := bid.BorrowableAmount()
		if usableAmount.Amount.IsZero() {
			continue
		}

		// bigger msg Amount
		if borrowAmount.IsGTE(usableAmount) {
			borrow := Borrowing{
				Amount:             sdk.NewCoin(usableAmount.Denom, usableAmount.Amount),
				StartAt:            start,
				PaidInterestAmount: sdk.NewCoin(usableAmount.Denom, sdk.ZeroInt()),
			}
			bid.Borrowings = append(bid.Borrowings, borrow)
			borrowAmount = borrowAmount.Sub(borrow.Amount)
		} else {
			borrow := Borrowing{
				Amount:             sdk.NewCoin(borrowAmount.Denom, borrowAmount.Amount),
				StartAt:            start,
				PaidInterestAmount: sdk.NewCoin(borrowAmount.Denom, sdk.ZeroInt()),
			}
			bid.Borrowings = append(bid.Borrowings, borrow)
			borrowAmount.Amount = sdk.ZeroInt()
		}
		// todo: execute func
		// k.SetBid(ctx, bid)

	}
}

func (m NftBids) BorrowableAmount(denom string) sdk.Coin {
	coin := sdk.NewCoin(denom, sdk.ZeroInt())
	for _, s := range m {
		coin = coin.Add(s.BorrowableAmount())
	}
	return coin
}

func (m NftBids) TotalDeposit() sdk.Coin {
	if len(m) == 0 {
		return sdk.Coin{}
	}
	coin := sdk.NewCoin(m[0].DepositAmount.Denom, sdk.ZeroInt())
	for _, s := range m {
		coin = coin.Add(s.DepositAmount)
	}
	return coin
}

func (m NftBids) TotalInterestAmount(end time.Time) sdk.Coin {
	if len(m) == 0 {
		return sdk.Coin{}
	}
	coin := sdk.NewCoin(m[0].DepositAmount.Denom, sdk.ZeroInt())
	for _, bid := range m {
		coin = coin.Add(bid.TotalInterestAmount(end))
	}
	return coin
}

func (m NftBids) LiquidationAmount(denom string, end time.Time) sdk.Coin {
	coin := sdk.NewCoin(denom, sdk.ZeroInt())
	for _, s := range m {
		coin = coin.Add(s.LiquidationAmount(end))
	}
	return coin
}

func (m NftBids) FindKickOutBid(newBid NftBid, end time.Time) NftBid {
	HigherDepositBids := m.SortHigherDeposit()
	kickOutBid := NftBid{}
	for _, b := range HigherDepositBids {
		if b.IsLT(newBid) {
			refundAmount := b.TotalInterestAmount(end)
			refundAmount = refundAmount.Add(b.DepositAmount)
			if refundAmount.IsLT(newBid.DepositAmount) {
				kickOutBid = b
				break
			}
		}
	}
	return kickOutBid
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
	total = total.Add(principal).Add(interest)
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

func CalcPartInterest(total, surplus sdk.Int, interest sdk.DecCoin) sdk.Int {
	if total.IsZero() {
		return sdk.ZeroInt()
	}
	decTotalInterest := sdk.NewDecFromInt(total)
	decSurplusAmount := sdk.NewDecFromInt(surplus)
	rate := interest.Amount.Quo(decTotalInterest)
	return decSurplusAmount.Mul(rate).TruncateInt()
}
