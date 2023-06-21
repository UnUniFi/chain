package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
)

func MinSettlementAmount(bids []NftBid) types.Coin {
	bidsSortedByPrice := NftBids(bids).SortHigherPrice()
	minimumSettlementAmount := types.NewCoin(bidsSortedByPrice[0].BidAmount.Denom, sdk.NewInt(0))
	forfeitedDeposit := types.NewCoin(bidsSortedByPrice[0].DepositAmount.Denom, sdk.NewInt(0))

	for _, bid := range bidsSortedByPrice {
		if minimumSettlementAmount.Size() == 0 || bid.BidAmount.Add(forfeitedDeposit).IsLT(minimumSettlementAmount) {
			minimumSettlementAmount = bid.BidAmount.Add(forfeitedDeposit)
		}
		forfeitedDeposit = forfeitedDeposit.Add(bid.DepositAmount)
	}

	return minimumSettlementAmount
}

func ExpectedRepayAmount(bids []NftBid, borrowBids []BorrowBid, time time.Time) sdk.Coin {
	expectedRepayAmount := types.NewCoin(bids[0].BidAmount.Denom, sdk.NewInt(0))
	for _, borrowBid := range borrowBids {
		for _, nftBid := range bids {
			if borrowBid.Bidder == nftBid.Id.Bidder {
				if borrowBid.Amount.IsLTE(nftBid.DepositAmount) {
					expectedInterest := nftBid.CalcInterest(borrowBid.Amount, nftBid.DepositLendingRate, time, nftBid.BiddingPeriod)
					expectedRepayAmount = expectedRepayAmount.Add(borrowBid.Amount).Add(expectedInterest)
				} else {
					// if borrow larger than deposit, then borrow all deposit
					expectedInterest := nftBid.CalcInterest(nftBid.DepositAmount, nftBid.DepositLendingRate, time, nftBid.BiddingPeriod)
					expectedRepayAmount = expectedRepayAmount.Add(nftBid.DepositAmount).Add(expectedInterest)
				}
			}
		}
	}
	return expectedRepayAmount
}

func ExistRepayAmount(bids []NftBid) sdk.Coin {
	existRepayAmount := types.NewCoin(bids[0].BidAmount.Denom, sdk.NewInt(0))
	for _, nftBid := range bids {
		for _, borrowing := range nftBid.Borrowings {
			existInterest := nftBid.CalcInterest(borrowing.Amount, nftBid.DepositLendingRate, borrowing.StartAt, nftBid.BiddingPeriod)
			existRepayAmount = existRepayAmount.Add(borrowing.Amount).Add(existInterest)
		}
	}
	return existRepayAmount
}

func IsAbleToBorrow(bids []NftBid, borrowBids []BorrowBid, time time.Time) bool {
	minimumSettlementAmount := MinSettlementAmount(bids)
	expectedRepayAmount := ExpectedRepayAmount(bids, borrowBids, time)
	// todo : include existRepayAmount
	return expectedRepayAmount.IsLTE(minimumSettlementAmount)

}

func IsAbleToCancelBid(cancellingBidId BidId, bids []NftBid) bool {
	var afterCanceledBids []NftBid
	for _, bid := range bids {
		if bid.Id != cancellingBidId {
			afterCanceledBids = append(afterCanceledBids, bid)
		}
	}
	minimumSettlementAmount := MinSettlementAmount(afterCanceledBids)
	existRepayAmount := ExistRepayAmount(bids)
	return existRepayAmount.IsLTE(minimumSettlementAmount)
}
