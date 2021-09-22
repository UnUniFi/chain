package types

const (
	// ModuleName defines the module name
	ModuleName = "botanydist"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// BotanydistMacc module account for botanydist
	BotanydistMacc = ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ParamsKey            = "Params-value-"
	RewardKey            = "Reward-value-"
	RewardCountKey       = "Reward-count-"
	PreviousBlockTimeKey = "PreviousBlockTime"
	GovDenomKey          = "GovDenom"
)
