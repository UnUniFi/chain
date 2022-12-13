package types

func (m NftBid) Equal(b NftBid) bool {
	return m.Bidder == b.Bidder && m.NftId == b.NftId && m.BidAmount.Equal(b)
}
