package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TracingKeyPrefix is the prefix to retrieve all Tracing
	TracingKeyPrefix = "Tracing/value/"

	ExemplaryTraderTracingKeyPrefix = "ExemplaryTrader/Tracing/value/"
)

// TracingKey returns the store key to retrieve a Tracing from the index fields
func TracingKey(
	address string,
) []byte {
	var key []byte

	indexBytes := []byte(address)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func ExemplaryTraderTracingKey(
	exemplaryTraderAddress string,
	tracerAddress string,
) []byte {
	var key []byte

	exemplaryTraderAddressBytes := []byte(exemplaryTraderAddress)
	tracerAddressBytes := []byte(tracerAddress)
	key = append(key, exemplaryTraderAddressBytes...)
	key = append(key, []byte("/")...)
	key = append(key, tracerAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
