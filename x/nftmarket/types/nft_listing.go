package types

func (m NftListing) IdBytes() []byte {
	return m.NftId.IdBytes()
}

func (m NftListing) IsActive() bool {
	return m.State == ListingState_SELLING || m.State == ListingState_BIDDING
}

func (ni NftIdentifier) IdBytes() []byte {
	return NftBytes(ni.ClassId, ni.NftId)
}

func (b NftBid) IdBytes() []byte {
	return b.NftId.IdBytes()
}
