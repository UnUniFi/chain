package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExemplaryTraderKeyPrefix is the prefix to retrieve all ExemplaryTrader
	ExemplaryTraderKeyPrefix = "ExemplaryTrader/value/"
)

// ExemplaryTraderKey returns the store key to retrieve a ExemplaryTrader from the index fields
func ExemplaryTraderKey(
	address string,
) []byte {
	var key []byte

	indexBytes := []byte(address)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
