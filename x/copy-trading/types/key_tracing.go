package types

import "encoding/binary"

var _ binary.ByteOrder

const (
    // TracingKeyPrefix is the prefix to retrieve all Tracing
	TracingKeyPrefix = "Tracing/value/"
)

// TracingKey returns the store key to retrieve a Tracing from the index fields
func TracingKey(
index string,
) []byte {
	var key []byte
    
    indexBytes := []byte(index)
    key = append(key, indexBytes...)
    key = append(key, []byte("/")...)
    
	return key
}