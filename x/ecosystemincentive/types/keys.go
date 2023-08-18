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
	// KeyPrefixRecipient defines prefix key for Incentive-Recipient
	KeyPrefixRecipient = []byte{0x01}

	// KeyPrefixReward defines prefix key for Reward
	KeyPrefixRewardStore = []byte{0x02}

	// KeyPrefixRecipientByNftId defines prefix key for nft_id with recipient
	KeyPrefixRecipientByNftId = []byte{0x03}
)
