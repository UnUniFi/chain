package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
)

func MinSettlementAmount(bids []NftBid) (types.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrNotExistsBid
	}
	bidsSortedByPrice := NftBids(bids).SortHigherPrice()
	return FindMinSettlementAmount(bidsSortedByPrice)
}

func FindMinSettlementAmount(bidsSortedByPrice []NftBid) (types.Coin, error) {
	if len(bidsSortedByPrice) == 0 {
		return types.Coin{}, ErrNotExistsBid
	}
	minimumSettlementAmount := types.NewCoin(bidsSortedByPrice[0].BidAmount.Denom, sdk.NewInt(0))
	forfeitedDeposit := types.NewCoin(bidsSortedByPrice[0].DepositAmount.Denom, sdk.NewInt(0))

	for _, bid := range bidsSortedByPrice {
		if minimumSettlementAmount.IsZero() || bid.BidAmount.Add(forfeitedDeposit).IsLT(minimumSettlementAmount) {
			minimumSettlementAmount = bid.BidAmount.Add(forfeitedDeposit)
		}
		forfeitedDeposit = forfeitedDeposit.Add(bid.DepositAmount)
	}

	// If higher than the total deposits, the minimum settlement amount is the total deposits.
	if forfeitedDeposit.IsLT(minimumSettlementAmount) {
		minimumSettlementAmount = forfeitedDeposit
	}

	return minimumSettlementAmount, nil
}

func LiquidationBid(bidsSortedByDeposit []NftBid, time time.Time) (NftBid, error) {
	if len(bidsSortedByDeposit) == 0 {
		return NftBid{}, ErrNotExistsBid
	}
	settlementAmount, _ := ExistRepayAmountAtTime(bidsSortedByDeposit, time)
	forfeitedDeposit := types.NewCoin(bidsSortedByDeposit[0].DepositAmount.Denom, sdk.NewInt(0))
	var ret NftBid

	for _, bid := range bidsSortedByDeposit {
		if !bid.IsPaidBidAmount() {
			forfeitedDeposit = forfeitedDeposit.Add(bid.DepositAmount)
			continue
		}
		if bid.BidAmount.Add(forfeitedDeposit).IsLT(settlementAmount) {
			continue
		}
		ret = bid
		break
	}

	if ret.IsNil() {
		return NftBid{}, nil
	}

	return ret, nil
}

func ForfeitedBidsAndRefundBids(bidsSortedByDeposit []NftBid, winBid NftBid) ([]NftBid, []NftBid) {
	isDecidedWinningBid := false
	forfeitedBids := []NftBid{}
	refundBids := []NftBid{}
	for _, bid := range bidsSortedByDeposit {
		if bid.Id.Bidder == winBid.Id.Bidder {
			isDecidedWinningBid = true
			continue
		}
		if isDecidedWinningBid {
			refundBids = append(refundBids, bid)
			continue
		}
		if bid.IsPaidBidAmount() {
			refundBids = append(refundBids, bid)
		} else {
			forfeitedBids = append(forfeitedBids, bid)
		}
	}
	return forfeitedBids, refundBids
}

func ExpectedRepayAmount(bids []NftBid, borrowBids []BorrowBid, time time.Time) (sdk.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrNotExistsBid
	}
	expectedRepayAmount := types.NewCoin(bids[0].BidAmount.Denom, sdk.NewInt(0))
	for _, borrowBid := range borrowBids {
		for _, nftBid := range bids {
			if nftBid.DepositAmount.Denom != borrowBid.Amount.Denom {
				return types.Coin{}, ErrInvalidBidDenom
			}
			if borrowBid.Bidder == nftBid.Id.Bidder {
				if borrowBid.Amount.IsLTE(nftBid.DepositAmount) {
					expectedInterest := nftBid.CalcCompoundInterest(borrowBid.Amount, time, nftBid.ExpiryAt)
					expectedRepayAmount = expectedRepayAmount.Add(borrowBid.Amount).Add(expectedInterest)
				} else {
					// if borrow larger than deposit, then borrow all deposit
					expectedInterest := nftBid.CalcCompoundInterest(nftBid.DepositAmount, time, nftBid.ExpiryAt)
					expectedRepayAmount = expectedRepayAmount.Add(nftBid.DepositAmount).Add(expectedInterest)
				}
			}
		}
	}
	if expectedRepayAmount.IsZero() {
		return types.Coin{}, ErrInvalidRepayAmount
	}
	return expectedRepayAmount, nil
}

func ExistRepayAmount(bids []NftBid) (sdk.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrBidDoesNotExists
	}
	existRepayAmount := types.NewCoin(bids[0].BidAmount.Denom, sdk.NewInt(0))
	for _, nftBid := range bids {
		existInterest := nftBid.CalcCompoundInterest(nftBid.Borrow.Amount, nftBid.Borrow.LastRepaidAt, nftBid.ExpiryAt)
		existRepayAmount = existRepayAmount.Add(nftBid.Borrow.Amount).Add(existInterest)
	}
	return existRepayAmount, nil
}

func ExistRepayAmountAtTime(bids []NftBid, time time.Time) (sdk.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrBidDoesNotExists
	}
	existRepayAmount := types.NewCoin(bids[0].BidAmount.Denom, sdk.NewInt(0))
	for _, nftBid := range bids {
		existInterest := nftBid.CalcCompoundInterest(nftBid.Borrow.Amount, nftBid.Borrow.LastRepaidAt, time)
		existRepayAmount = existRepayAmount.Add(nftBid.Borrow.Amount).Add(existInterest)
	}
	return existRepayAmount, nil
}

func MaxBorrowAmount(bids []NftBid, time time.Time) (types.Coin, error) {
	if len(bids) == 0 {
		return types.Coin{}, ErrNotExistsBid
	}
	minSettlement, err := MinSettlementAmount(bids)
	if err != nil {
		return types.Coin{}, err
	}
	bidsSortedByRate := NftBids(bids).SortLowerInterestRate()
	borrowAmount := types.NewCoin(bids[0].BidAmount.Denom, sdk.NewInt(0))

	for _, bid := range bidsSortedByRate {
		expectedInterest := bid.CalcCompoundInterest(bid.DepositAmount, time, bid.ExpiryAt)
		repayAmount := bid.DepositAmount.Add(expectedInterest)
		if repayAmount.IsLT(minSettlement) {
			minSettlement = minSettlement.Sub(repayAmount)
			borrowAmount = borrowAmount.Add(bid.DepositAmount)
		} else {
			discountAmount := minSettlement.Amount.Mul(bid.DepositAmount.Amount).Quo(repayAmount.Amount)
			borrowAmount = borrowAmount.Add(types.NewCoin(borrowAmount.Denom, discountAmount))
			break
		}
	}
	return borrowAmount, nil
}

func IsAbleToBorrow(bids []NftBid, borrowBids []BorrowBid, time time.Time) bool {
	minimumSettlementAmount, err := MinSettlementAmount(bids)
	if err != nil {
		return false
	}
	expectedRepayAmount, err := ExpectedRepayAmount(bids, borrowBids, time)
	if err != nil {
		return false
	}
	return expectedRepayAmount.IsLTE(minimumSettlementAmount)
}

func IsAbleToCancelBid(cancellingBidId BidId, bids []NftBid) bool {
	// check if not borrowed before this func
	var afterCanceledBids []NftBid
	for _, bid := range bids {
		if bid.Id.Bidder != cancellingBidId.Bidder {
			afterCanceledBids = append(afterCanceledBids, bid)
		}
	}
	minimumSettlementAmount, err := MinSettlementAmount(afterCanceledBids)
	if err != nil {
		return false
	}
	existRepayAmount, err := ExistRepayAmount(bids)
	if err != nil {
		return false
	}
	return existRepayAmount.IsLTE(minimumSettlementAmount)
}

func IsAbleToReBid(bids []NftBid, oldBidId BidId, newBid NftBid) bool {
	// check if not borrowed before this func
	if oldBidId.Bidder != newBid.Id.Bidder {
		return false
	}

	var afterUpdatedBids []NftBid
	for _, bid := range bids {
		if bid.Id.Bidder != oldBidId.Bidder {
			afterUpdatedBids = append(afterUpdatedBids, bid)
		}
	}
	afterUpdatedBids = append(afterUpdatedBids, newBid)
	minimumSettlementAmount, err := MinSettlementAmount(afterUpdatedBids)
	if err != nil {
		return false
	}
	existRepayAmount, err := ExistRepayAmount(bids)
	if err != nil {
		return false
	}
	return existRepayAmount.IsLTE(minimumSettlementAmount)
}
