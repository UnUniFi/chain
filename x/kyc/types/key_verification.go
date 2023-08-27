package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VerificationKeyPrefix is the prefix to retrieve all Verification
	VerificationKeyPrefix = "Verification/value/"
)

// VerificationKey returns the store key to retrieve a Verification from the index fields
func VerificationKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
