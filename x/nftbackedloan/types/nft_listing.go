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

func (m NftListing) IsBidding() bool {
	return m.State == ListingState_BIDDING
}

func (m NftListing) IsEnded() bool {
	return m.State == ListingState_SELLING_DECISION || m.State == ListingState_LIQUIDATION || m.State == ListingState_SUCCESSFUL_BID
}

func (m NftListing) IsSelling() bool {
	return m.State == ListingState_LISTING || m.State == ListingState_BIDDING
}

func (m NftListing) CanCancelBid() bool {
	return m.CanBid()
}

func (m NftListing) CanBid() bool {
	return m.State == ListingState_LISTING || m.State == ListingState_BIDDING
}
