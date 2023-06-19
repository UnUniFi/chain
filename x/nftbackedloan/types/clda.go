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

// Do not consider refinancing
func BorrowableAmount(bidsSortedByPrice []Bid, bidsSortedByInterestRate []Bid, duration time.Duration) uint64 {
	var minimumSettlementAmount uint64
	var forfeitedDeposit uint64

	for _, bid := range bidsSortedByPrice {
		if minimumSettlementAmount == 0 || minimumSettlementAmount > bid.Price+forfeitedDeposit {
			minimumSettlementAmount = bid.Price + forfeitedDeposit
		}
		forfeitedDeposit += bid.Deposit
	}

	var repayAmount uint64
	var borrowableAmount uint64

	for _, bid := range bidsSortedByInterestRate {
		repayAmount += (1 + AdjustAnnualRateWithDuration(bid.AnnualInterestRate, duration)) * bid.Deposit

		if repayAmount > minimumSettlementAmount {
			break
		}
		borrowableAmount += bid.Deposit
	}

	return borrowableAmount
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

func IsBidAbleToCancel(cancellingBidId BidId, bidsSortedByPrice []Bid, bidsSortedByInterestRate []Bid, durationToClosestExpiry time.Duration, borrowedAmount uint64) bool {
	return BorrowableAmount(
		FilterBidsById(bidsSortedByPrice, cancellingBidId),
		FilterBidsById(bidsSortedByInterestRate, cancellingBidId),
		durationToClosestExpiry,
	) >= borrowedAmount
}
