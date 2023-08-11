package types

import (
	"sort"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	osmomath "github.com/UnUniFi/chain/osmomath"
)

func (m Bid) Equal(b Bid) bool {
	if m.IsNil() || b.IsNil() {
		if m.IsNil() && b.IsNil() {
			return true
		} else {
			return false
		}
	}
	return m.Id.Bidder == b.Id.Bidder && m.Id.NftId.ClassId == b.Id.NftId.ClassId && m.Id.NftId.TokenId == b.Id.NftId.TokenId && m.Price.Equal(b.Price)
}
func (m Bid) IsLT(b Bid) bool {
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

func (m Bid) GetIdToByte() []byte {
	return NftBidBytes(m.Id.NftId.ClassId, m.Id.NftId.TokenId, m.Id.Bidder)
}

func (m Bid) IsBorrowed() bool {
	return m.Borrow.Amount.IsPositive()
}

func (m Bid) LiquidationAmount(time time.Time) sdk.Coin {
	interestAmount := m.CalcCompoundInterest(m.Borrow.Amount, m.Borrow.LastRepaidAt, time)
	return m.Borrow.Amount.Add(interestAmount)
}

func (m Bid) CompoundInterest(end time.Time) sdk.Coin {
	return m.CalcCompoundInterest(m.Borrow.Amount, m.Borrow.LastRepaidAt, end)
}

func (m Bid) CalcCompoundInterest(lendCoin sdk.Coin, startTime time.Time, endTime time.Time) sdk.Coin {
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

func (m Bid) RepayInfo(repayAmount sdk.Coin, repaymentTime time.Time) RepayInfo {
	interest := m.CalcCompoundInterest(m.Borrow.Amount, m.Borrow.LastRepaidAt, repaymentTime)
	total := m.Borrow.Amount.Add(interest)

	if repayAmount.IsGTE(total) {
		remainingAmount := sdk.NewCoin(m.Borrow.Amount.Denom, sdk.ZeroInt())
		return RepayInfo{
			RepaidAmount:         total,
			RepaidInterestAmount: interest,
			RemainingAmount:      remainingAmount,
			LastRepaidAt:         repaymentTime,
		}
	} else {
		remainingAmount := total.Sub(repayAmount)
		return RepayInfo{
			RepaidAmount:         repayAmount,
			RepaidInterestAmount: interest,
			RemainingAmount:      remainingAmount,
			LastRepaidAt:         repaymentTime,
		}
	}
}

func (m Bid) RepayInfoInFull(repaymentTime time.Time) RepayInfo {
	m.RepayInfo(m.Borrow.Amount, repaymentTime)
	interest := m.CalcCompoundInterest(m.Borrow.Amount, m.Borrow.LastRepaidAt, repaymentTime)
	total := m.Borrow.Amount.Add(interest)

	remainingAmount := sdk.NewCoin(m.Borrow.Amount.Denom, sdk.ZeroInt())
	return RepayInfo{
		RepaidAmount:         total,
		RepaidInterestAmount: interest,
		RemainingAmount:      remainingAmount,
		LastRepaidAt:         repaymentTime,
	}
}

func (m Bid) BidderPaidAmount() sdk.Coin {
	return m.PaidAmount.Add(m.Deposit)
}

func (m Bid) IsPaidSalePrice() bool {
	fullPaidAmount := m.BidderPaidAmount()
	return fullPaidAmount.Equal(m.Price)
}

func (m Bid) CanCancel() bool {
	return !m.IsBorrowed()
}

func (m Bid) CanReBid() bool {
	return !m.IsBorrowed()
}

func (m Bid) IsNil() bool {
	return m.Id.Bidder == ""
}

type NftBids []Bid

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

func (m NftBids) GetHighestBid() (Bid, error) {
	if len(m) == 0 {
		return Bid{}, ErrBidDoesNotExists
	}
	highestBid := m[0]
	for _, bid := range m {
		if highestBid.Price.IsLT(bid.Price) {
			highestBid = bid
		}
	}

	return highestBid, nil
}

func (m NftBids) GetBidByBidder(bidder string) Bid {
	for _, bid := range m {
		if bid.Id.Bidder == bidder {
			return bid
		}
	}
	return Bid{}
}

func (m NftBids) RemoveBid(targetBid Bid) NftBids {
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
