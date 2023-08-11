package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
)

func MinSettlementAmount(bids []Bid, listing Listing) (types.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrBidDoesNotExists
	}
	bidsSortedByPrice := NftBids(bids).SortHigherPrice()
	return FindMinSettlementAmount(bidsSortedByPrice, listing)
}

func FindMinSettlementAmount(bidsSortedByPrice []Bid, listing Listing) (types.Coin, error) {
	if len(bidsSortedByPrice) == 0 {
		return types.Coin{}, ErrBidDoesNotExists
	}
	minimumSettlementAmount := types.NewCoin(listing.BidDenom, sdk.NewInt(0))
	forfeitedDeposit := types.NewCoin(listing.BidDenom, sdk.NewInt(0))

	for _, bid := range bidsSortedByPrice {
		if minimumSettlementAmount.IsZero() || bid.Price.Add(forfeitedDeposit).IsLT(minimumSettlementAmount) {
			minimumSettlementAmount = bid.Price.Add(forfeitedDeposit)
		}
		forfeitedDeposit = forfeitedDeposit.Add(bid.Deposit)
	}

	// If higher than the total deposits, the minimum settlement amount is the total deposits.
	if forfeitedDeposit.IsLT(minimumSettlementAmount) {
		minimumSettlementAmount = forfeitedDeposit
	}

	return minimumSettlementAmount, nil
}

// arg: bids sorted by higher deposit & liquidation time
// return: win bid, deposit collect bids, refund bids, error
func LiquidationBid(bidsSortedByDeposit []Bid, listing Listing, time time.Time) (Bid, []Bid, []Bid, error) {

	if len(bidsSortedByDeposit) == 0 {
		return Bid{}, nil, nil, ErrBidDoesNotExists
	}
	settlementAmount, _ := ExistRepayAmountAtLiquidation(bidsSortedByDeposit, listing, time)
	forfeitedDeposit := types.NewCoin(listing.BidDenom, sdk.NewInt(0))
	var winnerBid Bid
	notInspectedBidsSortedByDeposit := NftBids(bidsSortedByDeposit)
	forfeitedBids := []Bid{}
	refundBids := []Bid{}

	// prevention of infinite loops
	highestBid, err := notInspectedBidsSortedByDeposit.GetHighestBid()
	if err != nil {
		return Bid{}, nil, nil, err
	}
	if highestBid.Price.IsLT(settlementAmount) {
		return Bid{}, nil, nil, ErrCannotLiquidation
	}

	// loop until all bids are handled
	for len(notInspectedBidsSortedByDeposit) > 0 {
		for _, bid := range notInspectedBidsSortedByDeposit {
			// if the win bid is found, other bids are refunded.
			if !winnerBid.IsNil() {
				refundBids = append(refundBids, bid)
				notInspectedBidsSortedByDeposit = notInspectedBidsSortedByDeposit.RemoveBid(bid)
				continue
			}
			// skip, if the bid is paid & the bid amount + forfeited deposit < settlement amount
			if bid.Price.Add(forfeitedDeposit).IsLT(settlementAmount) {
				continue
			}
			// if liquidation is available, the bid is win or collect
			if bid.IsPaidSalePrice() {
				winnerBid = bid
			} else {
				forfeitedDeposit = forfeitedDeposit.Add(bid.Deposit)
				forfeitedBids = append(forfeitedBids, bid)
			}
			notInspectedBidsSortedByDeposit = notInspectedBidsSortedByDeposit.RemoveBid(bid)
		}
	}

	// if the win bid is not found (no one paid case)
	if winnerBid.IsNil() {
		// No error, if liquidation is available
		if settlementAmount.IsLTE(forfeitedDeposit) {
			return Bid{}, forfeitedBids, nil, nil
		}
		// With Error, if liquidation is not available
		return Bid{}, forfeitedBids, refundBids, ErrCannotLiquidation
	}
	return winnerBid, forfeitedBids, refundBids, nil
}

func ForfeitedBidsAndRefundBids(bidsSortedByDeposit []Bid, winnerBid Bid) ([]Bid, []Bid) {
	isDecidedWinningBid := false
	forfeitedBids := []Bid{}
	refundBids := []Bid{}
	for _, bid := range bidsSortedByDeposit {
		if bid.Id.Bidder == winnerBid.Id.Bidder {
			isDecidedWinningBid = true
			continue
		}
		if isDecidedWinningBid {
			refundBids = append(refundBids, bid)
			continue
		}
		if bid.IsPaidSalePrice() {
			refundBids = append(refundBids, bid)
		} else {
			forfeitedBids = append(forfeitedBids, bid)
		}
	}
	return forfeitedBids, refundBids
}

func ExpectedRepayAmount(bids []Bid, borrowBids []BorrowBid, listing Listing, time time.Time) (sdk.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrBidDoesNotExists
	}
	expectedRepayAmount := types.NewCoin(listing.BidDenom, sdk.NewInt(0))
	for _, borrowBid := range borrowBids {
		for _, nftBid := range bids {
			if borrowBid.Bidder == nftBid.Id.Bidder {
				if nftBid.Deposit.Denom != borrowBid.Amount.Denom {
					return types.Coin{}, ErrInvalidBorrowDenom
				}

				if borrowBid.Amount.IsLTE(nftBid.Deposit) {
					expectedInterest := nftBid.CalcCompoundInterest(borrowBid.Amount, time, nftBid.Expiry)
					expectedRepayAmount = expectedRepayAmount.Add(borrowBid.Amount).Add(expectedInterest)
				} else {
					// if borrow larger than deposit, then borrow all deposit
					expectedInterest := nftBid.CalcCompoundInterest(nftBid.Deposit, time, nftBid.Expiry)
					expectedRepayAmount = expectedRepayAmount.Add(nftBid.Deposit).Add(expectedInterest)
				}
			}
		}
	}
	if expectedRepayAmount.IsZero() {
		return types.Coin{}, ErrInvalidRepayAmount
	}
	return expectedRepayAmount, nil
}

func ExistRepayAmount(bids []Bid, listing Listing) (sdk.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrBidDoesNotExists
	}
	existRepayAmount := types.NewCoin(listing.BidDenom, sdk.NewInt(0))
	for _, nftBid := range bids {
		existInterest := nftBid.CalcCompoundInterest(nftBid.Borrow.Amount, nftBid.Borrow.LastRepaidAt, nftBid.Expiry)
		existRepayAmount = existRepayAmount.Add(nftBid.Borrow.Amount).Add(existInterest)
	}
	return existRepayAmount, nil
}

func ExistRepayAmountAtTime(bids []Bid, listing Listing, time time.Time) (sdk.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrBidDoesNotExists
	}
	existRepayAmount := types.NewCoin(listing.BidDenom, sdk.NewInt(0))
	for _, nftBid := range bids {
		existInterest := nftBid.CalcCompoundInterest(nftBid.Borrow.Amount, nftBid.Borrow.LastRepaidAt, time)
		existRepayAmount = existRepayAmount.Add(nftBid.Borrow.Amount).Add(existInterest)
	}
	return existRepayAmount, nil
}

func ExistRepayAmountAtLiquidation(bids []Bid, listing Listing, liquidationTime time.Time) (sdk.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrBidDoesNotExists
	}
	existRepayAmount := types.NewCoin(listing.BidDenom, sdk.NewInt(0))
	for _, nftBid := range bids {
		if nftBid.IsPaidSalePrice() {
			existInterest := nftBid.CalcCompoundInterest(nftBid.Borrow.Amount, nftBid.Borrow.LastRepaidAt, liquidationTime)
			existRepayAmount = existRepayAmount.Add(nftBid.Borrow.Amount).Add(existInterest)
		}
		existRepayAmount = existRepayAmount.Add(nftBid.Borrow.Amount)
	}
	return existRepayAmount, nil
}

func MaxBorrowAmount(bids []Bid, listing Listing, time time.Time) (types.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrBidDoesNotExists
	}
	minSettlement, err := MinSettlementAmount(bids, listing)
	if err != nil {
		return types.Coin{}, err
	}
	bidsSortedByRate := NftBids(bids).SortLowerInterestRate()
	borrowAmount := types.NewCoin(listing.BidDenom, sdk.NewInt(0))

	for _, bid := range bidsSortedByRate {
		expectedInterest := bid.CalcCompoundInterest(bid.Deposit, time, bid.Expiry)
		repayAmount := bid.Deposit.Add(expectedInterest)
		if repayAmount.IsLT(minSettlement) {
			minSettlement = minSettlement.Sub(repayAmount)
			borrowAmount = borrowAmount.Add(bid.Deposit)
		} else {
			discountAmount := minSettlement.Amount.Mul(bid.Deposit.Amount).Quo(repayAmount.Amount)
			borrowAmount = borrowAmount.Add(types.NewCoin(borrowAmount.Denom, discountAmount))
			break
		}
	}
	return borrowAmount, nil
}

func IsAbleToBorrow(bids []Bid, borrowBids []BorrowBid, listing Listing, time time.Time) bool {
	minimumSettlementAmount, err := MinSettlementAmount(bids, listing)
	if err != nil {
		return false
	}
	expectedRepayAmount, err := ExpectedRepayAmount(bids, borrowBids, listing, time)
	if err != nil {
		return false
	}
	return expectedRepayAmount.IsLTE(minimumSettlementAmount)
}

func IsAbleToCancelBid(cancellingBidId BidId, bids []Bid, listing Listing) bool {
	// check if not borrowed before this func
	var afterCanceledBids []Bid
	for _, bid := range bids {
		if bid.Id.Bidder != cancellingBidId.Bidder {
			afterCanceledBids = append(afterCanceledBids, bid)
		}
	}
	minimumSettlementAmount, err := MinSettlementAmount(afterCanceledBids, listing)
	if err != nil {
		return false
	}
	existRepayAmount, err := ExistRepayAmount(bids, listing)
	if err != nil {
		return false
	}
	return existRepayAmount.IsLTE(minimumSettlementAmount)
}

func IsAbleToReBid(bids []Bid, oldBidId BidId, newBid Bid, listing Listing) bool {
	// check if not borrowed before this func
	if oldBidId.Bidder != newBid.Id.Bidder {
		return false
	}

	var afterUpdatedBids []Bid
	for _, bid := range bids {
		if bid.Id.Bidder != oldBidId.Bidder {
			afterUpdatedBids = append(afterUpdatedBids, bid)
		}
	}
	afterUpdatedBids = append(afterUpdatedBids, newBid)
	minimumSettlementAmount, err := MinSettlementAmount(afterUpdatedBids, listing)
	if err != nil {
		return false
	}
	existRepayAmount, err := ExistRepayAmount(bids, listing)
	if err != nil {
		return false
	}
	return existRepayAmount.IsLTE(minimumSettlementAmount)
}
