package types

func (m NftListing) IdBytes() []byte {
	return NftBytes(m.NftId.ClassId, m.NftId.NftId)
}

func (m NftListing) IsActive() bool {
	return m.State == ListingState_BIDDING
}
