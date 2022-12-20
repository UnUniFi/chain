package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	return m.State == ListingState_SELLING_DECISION || m.State == ListingState_END_LISTING
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
	return b.NftId.IdBytes()
}

func (m NftListing) CanRefinancing(allBids, expiredBids []NftBid) bool {
	fmt.Println("CanRefinancing-val")
	fmt.Println(allBids)
	fmt.Println(expiredBids)
	fmt.Println("CanRefinancing")
	fmt.Println(m.AutomaticRefinancing)
	if !m.AutomaticRefinancing {
		return false
	}
	fmt.Println("usableAmount")
	usableAmount := m.MaxPossibleBorrowAmount(allBids, expiredBids)
	fmt.Print(usableAmount)
	fmt.Println()
	fmt.Println("liquidationAmount")
	liquidationAmount := m.CalcAmount(expiredBids)
	fmt.Println(liquidationAmount)
	if liquidationAmount.GT(usableAmount) {
		return false
	}
	return true
}

func (m NftListing) CalcAmount(bids []NftBid) sdk.Int {
	return m.CalcAmountF(bids, func(NftBid) bool { return false })
}

func (m NftListing) CalcAmountF(bids []NftBid, conditionF func(bid NftBid) bool) sdk.Int {
	maxPossibleBorrowAmount := sdk.ZeroInt()
	for _, bid := range bids {
		if conditionF(bid) {
			continue
		}
		maxPossibleBorrowAmount = maxPossibleBorrowAmount.Add(bid.DepositAmount.Amount)
	}
	return maxPossibleBorrowAmount
}

func (m NftListing) MaxPossibleBorrowAmount(bids, expiredBids []NftBid) sdk.Int {
	f := func() func(NftBid) bool {
		expiredBidsMap := map[string]NftBid{}
		for _, bid := range expiredBids {
			expiredBidsMap[bid.GetIdToString()] = bid
		}
		return func(bid NftBid) bool {
			if _, ok := expiredBidsMap[bid.GetIdToString()]; ok {
				return true
			} else {
				return false
			}
		}
	}
	return m.CalcAmountF(bids, f())
}

func (m NftListing) IsSelling() bool {
	return m.State == ListingState_LISTING || m.State == ListingState_BIDDING
}
