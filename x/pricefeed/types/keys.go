package types

const (
	// ModuleName defines the module name
	ModuleName = "pricefeed"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ParamsKey       = "Params-value-"
	MarketKey       = "Market-value-"
	MarketCountKey  = "Market-count-"
	OracleKey       = "Oracle-value-"
	PriceKey        = "Price-value-"
	PriceCountKey   = "Price-count-"
	RawPriceKey     = "RawPrice-value-"
	CurrentPriceKey = "CurrentPrice-value-"
	RawPriceFeedKey = "RawPriceFeed-value-"
)

// CurrentPriceKeySuffix returns the prefix for the current price
func CurrentPriceKeySuffix(marketID string) []byte {
	return append(KeyPrefix(CurrentPriceKey), []byte(marketID)...)
}

// RawPriceKeySuffix returns the prefix for the raw price
func RawPriceKeySuffix(marketID string) []byte {
	return append(KeyPrefix(RawPriceFeedKey), []byte(marketID)...)
}
