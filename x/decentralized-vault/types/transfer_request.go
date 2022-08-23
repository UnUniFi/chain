package types

func (m TransferRequest) IdBytes() []byte {
	return []byte(m.NftId)
}
