package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VerificationKeyPrefix is the prefix to retrieve all Verification
	VerificationKeyPrefix = "Verification/value/"
)

// VerificationKey returns the store key to retrieve a Verification from the index fields
func VerificationKey(
	customer string,
	providerId uint64,
) []byte {
	var key []byte

	indexBytes := []byte(customer)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)
	key = append(key, GetProviderIDBytes(providerId)...)

	return key
}

// GetProviderIDBytes returns the byte representation of the ID
func GetProviderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
