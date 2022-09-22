package types

const (
	// ModuleName defines the module name
	ModuleName = "ecosystemincentive"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// RouterKey is the message route for nftmint
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ecosystemincentive"
)

var (
	// KeyPrefixIncentiveUnit defines prefix key for IncentiveUnit
	KeyPrefixIncentiveUnit = []byte{0x01}

	// KeyPrefixReward defines prefix key for Reward
	KeyPrefixReward = []byte{0x02}
)
