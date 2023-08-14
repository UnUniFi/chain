package types

import (
	"strings"
)

const (
	// ModuleName defines the module name
	ModuleName = "nftfactory"

	// StoreKey defines the primary module store key
	StoreKey = "_" + ModuleName

	// RouterKey is the message route for nftfactory
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines thee in-memory store key
	MemStoreKey = "mem_nftfactory"
)

// KeySeparator is used to combine parts of the keys in the store
const KeySeparator = "|"

var (
	ClassAuthorityMetadataKey = "authoritymetadata"
	ClassIdsPrefixKey         = "class_ids"
	CreatorPrefixKey          = "creator"
	AdminPrefixKey            = "admin"
)

// GetDenomPrefixStore returns the store prefix where all the data associated with a specific denom
// is stored
func GetDenomPrefixStore(classId string) []byte {
	return []byte(strings.Join([]string{ClassIdsPrefixKey, classId, ""}, KeySeparator))
}

// GetCreatorsPrefix returns the store prefix where the list of the denoms created by a specific
// creator are stored
func GetCreatorPrefix(creator string) []byte {
	return []byte(strings.Join([]string{CreatorPrefixKey, creator, ""}, KeySeparator))
}

// GetCreatorsPrefix returns the store prefix where a list of all creator addresses are stored
func GetCreatorsPrefix() []byte {
	return []byte(strings.Join([]string{CreatorPrefixKey, ""}, KeySeparator))
}
