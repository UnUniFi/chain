package types

import (
	"sort"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	osmomath "github.com/UnUniFi/chain/osmomath"
)

func (m NftBid) Equal(b NftBid) bool {
	if m.IsNil() || b.IsNil() {
		if m.IsNil() && b.IsNil() {
			return true
		} else {
			return false
		}
	}
	return m.Id.Bidder == b.Id.Bidder && m.Id.NftId.ClassId == b.Id.NftId.ClassId && m.Id.NftId.NftId == b.Id.NftId.NftId && m.Price.Equal(b.Price)
}
func (m NftBid) IsLT(b NftBid) bool {
	if b.Price.IsLTE(m.Price) {
		return false
	}
	if b.Deposit.IsLTE(m.Deposit) {
		return false
	}
	if b.InterestRate.GTE(m.InterestRate) {
		return false
	}

	return true
}

func (m NftBid) GetIdToByte() []byte {
	return NftBidBytes(m.Id.NftId.ClassId, m.Id.NftId.NftId, m.Id.Bidder)
}

func (m NftBid) IsBorrowed() bool {
	return m.Borrow.Amount.IsPositive()
}

func (m NftBid) LiquidationAmount(time time.Time) sdk.Coin {
	interestAmount := m.CalcCompoundInterest(m.Borrow.Amount, m.Borrow.LastRepaidAt, time)
	return m.Borrow.Amount.Add(interestAmount)
}

func (m NftBid) CompoundInterest(end time.Time) sdk.Coin {
	return m.CalcCompoundInterest(m.Borrow.Amount, m.Borrow.LastRepaidAt, end)
}

func (m NftBid) CalcCompoundInterest(lendCoin sdk.Coin, startTime time.Time, endTime time.Time) sdk.Coin {
	lendingDuration := endTime.Sub(startTime)
	if lendingDuration <= 0 {
		return sdk.NewCoin(lendCoin.Denom, sdk.ZeroInt())
	}
	oneYearDays := sdk.NewInt(365)
	oneDayHours := sdk.NewInt(24)
	durationUnitsYear := sdk.NewDecFromInt(sdk.NewInt(int64(lendingDuration.Hours()))).QuoInt(oneDayHours).QuoInt(oneYearDays)

	e := osmomath.NewDecWithPrec(2718281, 6) // 2.718281
	interestRateBidDec := osmomath.BigDecFromSDKDec(m.InterestRate)
	durationBigDec := osmomath.BigDecFromSDKDec(durationUnitsYear)
	// compoundInterestRate = exp ^ (interestRate * duration) - 1
	compoundRate := e.Power(interestRateBidDec.Mul(durationBigDec)).Sub(osmomath.OneDec()).SDKDec()
	result := sdk.NewDecFromInt(lendCoin.Amount).Mul(compoundRate)
	return sdk.NewCoin(lendCoin.Denom, result.RoundInt())
}

func (m NftBid) Repay(repayAmount sdk.Coin, payTime time.Time) RepayInfo {
	interest := m.CalcCompoundInterest(m.Borrow.Amount, m.Borrow.LastRepaidAt, payTime)
	total := m.Borrow.Amount.Add(interest)

	if repayAmount.IsGTE(total) {
		remainingAmount := sdk.NewCoin(m.Borrow.Amount.Denom, sdk.ZeroInt())
		return RepayInfo{
			RepaidAmount:         total,
			RepaidInterestAmount: interest,
			RemainingAmount:      remainingAmount,
			LastRepaidAt:         payTime,
		}
	} else {
		remainingAmount := total.Sub(repayAmount)
		return RepayInfo{
			RepaidAmount:         repayAmount,
			RepaidInterestAmount: interest,
			RemainingAmount:      remainingAmount,
			LastRepaidAt:         payTime,
		}
	}
}

func (m NftBid) RepayInFull(payTime time.Time) RepayInfo {
	interest := m.CalcCompoundInterest(m.Borrow.Amount, m.Borrow.LastRepaidAt, payTime)
	total := m.Borrow.Amount.Add(interest)

	remainingAmount := sdk.NewCoin(m.Borrow.Amount.Denom, sdk.ZeroInt())
	return RepayInfo{
		RepaidAmount:         total,
		RepaidInterestAmount: interest,
		RemainingAmount:      remainingAmount,
		LastRepaidAt:         payTime,
	}
}

func (m NftBid) BidderPaidAmount() sdk.Coin {
	return m.PaidAmount.Add(m.Deposit)
}

func (m NftBid) IsPaidSalePrice() bool {
	fullPaidAmount := m.BidderPaidAmount()
	return fullPaidAmount.Equal(m.Price)
}

func (m NftBid) CanCancel() bool {
	return !m.IsBorrowed()
}

func (m NftBid) CanReBid() bool {
	return !m.IsBorrowed()
}

func (m NftBid) IsNil() bool {
	return m.Id.Bidder == ""
}

type NftBids []NftBid

func (m NftBids) SortLowerInterestRate() NftBids {
	dest := append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].InterestRate.LT(dest[j].InterestRate)
	})
	return dest
}

func (m NftBids) SortHigherInterestRate() NftBids {
	dest := append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].InterestRate.GT(dest[j].InterestRate)
	})
	return dest
}

func (m NftBids) SortLowerExpiryDate() NftBids {
	dest := append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].Expiry.Before(dest[j].Expiry)
	})
	return dest
}

func (m NftBids) SortHigherDeposit() NftBids {
	dest := append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].Deposit.IsGTE(dest[j].Deposit)
	})
	return dest
}

func (m NftBids) SortHigherPrice() NftBids {
	dest := append(NftBids{}, m...)
	sort.SliceStable(dest, func(i, j int) bool {
		return dest[i].Price.IsGTE(dest[j].Deposit)
	})
	return dest
}

func (m NftBids) GetHighestBid() (NftBid, error) {
	if len(m) == 0 {
		return NftBid{}, ErrBidDoesNotExists
	}
	highestBid := m[0]
	for _, bid := range m {
		if highestBid.Price.IsLT(bid.Price) {
			highestBid = bid
		}
	}

	return highestBid, nil
}

func (m NftBids) GetBidByBidder(bidder string) NftBid {
	for _, bid := range m {
		if bid.Id.Bidder == bidder {
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
		excludeList[s.Id.Bidder] = true
	}
	var newArr NftBids
	for _, s := range m {
		if !excludeList[s.Id.Bidder] {
			newArr = append(newArr, s)
		}
	}
	return newArr
}

func (m NftBids) TotalBorrowedAmount() sdk.Coin {
	if len(m) == 0 {
		return sdk.Coin{}
	}
	coin := sdk.NewCoin(m[0].Borrow.Amount.Denom, sdk.ZeroInt())
	for _, bid := range m {
		coin = coin.Add(bid.Borrow.Amount)
	}
	return coin
}

func (m NftBids) TotalCompoundInterest(end time.Time) sdk.Coin {
	if len(m) == 0 {
		return sdk.Coin{}
	}
	coin := sdk.NewCoin(m[0].Borrow.Amount.Denom, sdk.ZeroInt())
	for _, bid := range m {
		coin = coin.Add(bid.CalcCompoundInterest(bid.Borrow.Amount, bid.Borrow.LastRepaidAt, end))
	}
	return coin
}
