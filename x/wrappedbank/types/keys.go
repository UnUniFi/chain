package types

const (
	// ModuleName defines the module name
	ModuleName = "wrappedbank"

	// Module account for nft trading fee collection
	// use ecosystem-incentive module account for now
	// [unused] NftTradingFee = "nfttradingfee"

	// StoreKey defines the primary module store key
	StoreKey = "wrappedbank"

	// RouterKey is the message route for nftmarket
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)
