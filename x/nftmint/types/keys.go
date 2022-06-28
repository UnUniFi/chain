package types

const (
	// ModuleName defines the module name
	ModuleName = "nftmint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for nftmint
	RouterKey = ModuleName

	// MemStoreKey defines thee in-memory store key
	MemStoreKey = "mem_nftmint"
)

var (
	// KeyPrefixClassAttributes defines prefix key for ClassAttributes
	KeyPrefixClassAttributes = []byte{0x01}
)
