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

func NextSettlementRightBidId(bidsSortedByDeposit []Bid, skip map[BidId]bool, borrowed uint64, forfeited uint64) *BidId {
	for _, bid := range bidsSortedByDeposit {
		if skip[bid.Id] {
			continue
		}
		if forfeited+bid.Price > borrowed {
			return &bid.Id
		}
	}

	return nil
}

// Simulate repay amount with given borrow amount
// Utilizing lowest interest rate first
// No refinancing
func SimulateRepayAmount(bidsSortedByInterestRate []Bid, borrowAmount uint64, now time.Time, maxDuration time.Duration) uint64 {
	return 0
}

// Simulate borrow amount with given repay amount
// Utilizing lowest interest rate first
// No refinancing
func SimulateBorrowAmount(bidsSortedByInterestRate []Bid, repayAmount uint64, now time.Time, maxDuration time.Duration) uint64 {

	return 0
}

// Do not consider refinancing
func BorrowableAmount(bidsSortedByPrice []Bid, bidsSortedByInterestRate []Bid, now time.Time) uint64 {
	var sumDeposit uint64
	var closestExpire time.Time
	for _, bid := range bidsSortedByPrice {
		sumDeposit += bid.Deposit
		if closestExpire.IsZero() || closestExpire.After(bid.End) {
			closestExpire = bid.End
		}
		maxDuration := closestExpire.Sub(now)

		simulatedRepayAmount := SimulateRepayAmount(bidsSortedByInterestRate, sumDeposit, now, maxDuration)

		if simulatedRepayAmount > bid.Price {
			return SimulateBorrowAmount(bidsSortedByInterestRate, sumDeposit-bid.Deposit+bid.Price, now, maxDuration)
		}
	}

	return sumDeposit
}
