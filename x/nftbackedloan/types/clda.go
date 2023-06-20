package types

import "time"

type Bid struct {
	Id                 BidId
	Price              uint64
	Deposit            uint64
	AnnualInterestRate uint64
	Start              time.Time
	End                time.Time
}

func SortBidsByPrice(bids []Bid) []Bid {
	return bids
}

func SortBidsByDeposit(bids []Bid) []Bid {
	return bids
}

func NextSettlementRightBidId(bidsSortedByDeposit []Bid, skip map[BidId]bool, forfeited uint64, needToCollectAmount uint64) BidId {
	var ret BidId
	for _, bid := range bidsSortedByDeposit {
		if skip[bid.Id] {
			continue
		}
		if bid.Price+forfeited < needToCollectAmount {
			continue
		}
		ret = bid.Id
		break
	}

	return ret
}

func AdjustAnnualRateWithDuration(annualInterestRate uint64, duration time.Duration) uint64 {
	return annualInterestRate * uint64(duration.Seconds()) / 31536000
}

func MinimumSettlementAmount(bidsSortedByPrice []Bid) uint64 {
	var minimumSettlementAmount uint64
	var forfeitedDeposit uint64

	for _, bid := range bidsSortedByPrice {
		if minimumSettlementAmount == 0 || minimumSettlementAmount > bid.Price+forfeitedDeposit {
			minimumSettlementAmount = bid.Price + forfeitedDeposit
		}
		forfeitedDeposit += bid.Deposit
	}

	return minimumSettlementAmount
}

func IsBidAbleToCancel(cancellingBidId BidId, bidsSortedByPrice []Bid, repayAmount uint64) bool {
	return MinimumSettlementAmount(
		FilterBidsById(bidsSortedByPrice, cancellingBidId),
	) >= repayAmount
}

func AdditionallyAcceptableRepayAmount(bidsSortedByPrice []Bid, bidsSortedByInterestRate []Bid, borrowed []struct {
	BorrowedAmount     uint64
	AnnualInterestRate uint64
	InterestAmount     uint64
	additionalDuration time.Duration
}) uint64 {
	var repayAmount uint64

	for i, _ := range borrowed {
		repayAmount += (1+AdjustAnnualRateWithDuration(borrowed[i].AnnualInterestRate, borrowed[i].additionalDuration))*borrowed[i].BorrowedAmount + borrowed[i].InterestAmount
	}

	return MinimumSettlementAmount(bidsSortedByPrice) - repayAmount
}

func AdditionallyAcceptableBorrowAmount(additionallyAcceptableRepayAmount uint64, bidsSortedByInterestRate []struct {
	Bid
	NotBorrowedAmount  uint64
	additionalDuration time.Duration
}) uint64 {
	var repayAmount uint64
	var borrowableAmount uint64

	for _, bid := range bidsSortedByInterestRate {
		repayAmount += (1 + AdjustAnnualRateWithDuration(bid.AnnualInterestRate, bid.additionalDuration)) * bid.NotBorrowedAmount

		if repayAmount > additionallyAcceptableRepayAmount {
			break
		}
		borrowableAmount += bid.NotBorrowedAmount
	}

	return borrowableAmount
}

// Do not consider refinancing
func BorrowableAmountInMinimumInterestCombination(bidsSortedByPrice []Bid, bidsSortedByInterestRate []struct {
	Bid
	NotBorrowedAmount  uint64
	additionalDuration time.Duration
}) uint64 {
	minimumSettlementAmount := MinimumSettlementAmount(bidsSortedByPrice)

	return AdditionallyAcceptableBorrowAmount(minimumSettlementAmount, bidsSortedByInterestRate)
}

func FilterBidsById(bids []Bid, excludeId BidId) []Bid {
	var ret []Bid
	for _, bid := range bids {
		if bid.Id == excludeId {
			continue
		}
		ret = append(ret, bid)
	}
	return ret
}
