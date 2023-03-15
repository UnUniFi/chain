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
	KeyPrefixRewardStore = []byte{0x02}

	// KeyPrefixIncentiveUnitIdByNftId defines prefix key for nft_id with incentive_id
	KeyPrefixIncentiveUnitIdByNftId = []byte{0x03}

	// KeyPrefixIncentiveUnitIdByAddr defines prefix key for incentiveUnitIdsByAddr with address
	KeyPrefixIncentiveUnitIdsByAddr = []byte{0x04}
)
