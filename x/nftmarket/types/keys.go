package types

const (
	// ModuleName defines the module name
	ModuleName = "nftmarket"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for nftmarket
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

const (
	// nft listing info by nft_id
	KeyPrefixNftListing = "nft_listing"
	// nft listing by owner
	KeyPrefixAddressNftListing = "address_nft_listing"
	// nft bid by nft_id
	KeyPrefixNftBid = "nft_bid"
	// nft bid by owner
	KeyPrefixAddressBid = "address_bid"
	// nft loan by nft_id
	KeyPrefixNftLoan = "nft_loan"
	// nft loan by owner
	KeyPrefixAddressLoan = "address_loan"
	// rewards by address
	KeyPrefixAddressRewards = "rewards"
)
