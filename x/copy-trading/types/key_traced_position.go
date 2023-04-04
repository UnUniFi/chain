package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TracingKeyPrefix is the prefix to retrieve all Tracing
	TracedPositionKeyPrefix                = "TracedPosition/value/"
	ExemplaryTraderTracedPositionKeyPrefix = "ExemplaryTrader/TracedPosition/value/"
)

// TracingKey returns the store key to retrieve a Tracing from the index fields
func TracedPositionKey(
	id string,
) []byte {
	var key []byte

	indexBytes := []byte(id)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
