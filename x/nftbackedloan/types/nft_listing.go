package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (m Listing) IdBytes() []byte {
	return m.NftId.IdBytes()
}

func (m Listing) ClassIdBytes() []byte {
	return m.NftId.ClassIdBytes()
}

func (m Listing) IsActive() bool {
	return m.State == ListingState_LISTING || m.State == ListingState_BIDDING
}

func (m Listing) IsFullPayment() bool {
	return m.State == ListingState_SELLING_DECISION || m.State == ListingState_LIQUIDATION
}

func (m Listing) IsSuccessfulBid() bool {
	return m.State == ListingState_SUCCESSFUL_BID
}

func (ni NftId) IdBytes() []byte {
	return NftBytes(ni.ClassId, ni.TokenId)
}

func (ni NftId) ClassIdBytes() []byte {
	return []byte(ni.ClassId)
}

func (b Bid) IdBytes() []byte {
	return b.Id.NftId.IdBytes()
}

func (m Listing) IsBidding() bool {
	return m.State == ListingState_BIDDING
}

func (m Listing) IsEnded() bool {
	return m.State == ListingState_SELLING_DECISION || m.State == ListingState_LIQUIDATION || m.State == ListingState_SUCCESSFUL_BID
}

func (m Listing) CanCancelBid() bool {
	return m.CanBid()
}

func (m Listing) CanBid() bool {
	return m.State == ListingState_LISTING || m.State == ListingState_BIDDING
}

func (m Listing) IsNegativeCollectedAmount() bool {
	return m.CollectedAmountNegative
}

func (m Listing) AddCollectedAmount(amount sdk.Coin) Listing {
	if m.CollectedAmountNegative {
		if m.CollectedAmount.IsLTE(amount) {
			m.CollectedAmount = amount.Sub(m.CollectedAmount)
			m.CollectedAmountNegative = false
		} else {
			m.CollectedAmount = m.CollectedAmount.Sub(amount)
		}
	} else {
		m.CollectedAmount = m.CollectedAmount.Add(amount)
	}
	return m
}

func (m Listing) SubCollectedAmount(amount sdk.Coin) Listing {
	if m.CollectedAmountNegative {
		m.CollectedAmount = m.CollectedAmount.Add(amount)
	} else {
		if m.CollectedAmount.IsLTE(amount) {
			m.CollectedAmount = amount.Sub(m.CollectedAmount)
			m.CollectedAmountNegative = true
		} else {
			m.CollectedAmount = m.CollectedAmount.Sub(amount)
		}
	}
	return m
}
