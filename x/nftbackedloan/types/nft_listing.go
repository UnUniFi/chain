package types

func (m NftListing) IdBytes() []byte {
	return m.NftId.IdBytes()
}

func (m NftListing) ClassIdBytes() []byte {
	return m.NftId.ClassIdBytes()
}

func (m NftListing) IsActive() bool {
	return m.State == ListingState_LISTING || m.State == ListingState_BIDDING
}

func (m NftListing) IsFullPayment() bool {
	return m.State == ListingState_SELLING_DECISION || m.State == ListingState_LIQUIDATION
}

func (m NftListing) IsSuccessfulBid() bool {
	return m.State == ListingState_SUCCESSFUL_BID
}

func (ni NftIdentifier) IdBytes() []byte {
	return NftBytes(ni.ClassId, ni.NftId)
}

func (ni NftIdentifier) ClassIdBytes() []byte {
	return []byte(ni.ClassId)
}

func (b NftBid) IdBytes() []byte {
	return b.Id.NftId.IdBytes()
}

// func (m NftListing) CanRefinancing(allBids, expiredBids []NftBid, now time.Time) bool {
// 	if !m.AutomaticRefinancing {
// 		return false
// 	}
// 	usableAmount := m.MaxPossibleBorrowAmount(allBids, expiredBids)
// 	liquidationAmount := NftBids(expiredBids).LiquidationAmount(m.BidDenom, now)
// 	if liquidationAmount.Amount.GT(usableAmount) {
// 		return false
// 	}
// 	return true
// }

// func (m NftListing) CalcAmount(bids []NftBid) sdk.Int {
// 	return m.CalcAmountF(bids, func(NftBid) bool { return false })
// }

// func (m NftListing) CalcAmountF(bids []NftBid, conditionF func(bid NftBid) bool) sdk.Int {
// 	DepositAmount := sdk.ZeroInt()
// 	for _, bid := range bids {
// 		if conditionF(bid) {
// 			continue
// 		}
// 		DepositAmount = DepositAmount.Add(bid.DepositAmount.Amount)
// 	}
// 	return DepositAmount
// }

// func (m NftListing) MaxPossibleBorrowAmount(bids, expiredBids []NftBid) sdk.Int {
// 	newBids := NftBids(bids).MakeExcludeExpiredBids(expiredBids)
// 	borrowableAmount := newBids.BorrowableAmount(m.BidDenom)
// 	return borrowableAmount.Amount
// }

func (m NftListing) IsSelling() bool {
	return m.State == ListingState_LISTING || m.State == ListingState_BIDDING
}

func (m NftListing) CanCancelBid() bool {
	return m.CanBid()
}

func (m NftListing) CanBid() bool {
	return m.State == ListingState_LISTING || m.State == ListingState_BIDDING
}
