package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "nftmarket"

	// StoreKey defines the primary module store key
	StoreKey = "ununifinftmarket"

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
	// nft listing by end time
	KeyPrefixEndTimeNftListing = "end_time_nft_listing"
	// nft bid by nft_id
	KeyPrefixNftBid = "nft_bid"
	// nft bid cancelled
	KeyPrefixNftBidCancelled = "nft_bid_cancelled"
	// nft bid by owner
	KeyPrefixAddressBid = "address_bid"
	// nft loan by nft_id
	KeyPrefixNftLoan = "nft_loan"
	// nft loan by owner
	KeyPrefixAddressNftLoan = "address_nft_loan"
	// rewards by address
	KeyPrefixAddressRewards = "rewards"
)

func NftBytes(classId, nftId string) []byte {
	return append(append([]byte(classId), byte(0xFF)), []byte(nftId)...)
}

func NftListingKey(idBytes []byte) []byte {
	return append([]byte(KeyPrefixNftListing), idBytes...)
}

func NftAddressNftListingPrefixKey(addr sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixAddressNftListing), address.MustLengthPrefix(addr)...)
}

func NftAddressNftListingKey(addr sdk.AccAddress, nftIdBytes []byte) []byte {
	return append(append([]byte(KeyPrefixAddressNftListing), address.MustLengthPrefix(addr)...), nftIdBytes...)
}

func NftBidKey(nftIdBytes []byte, bidder sdk.AccAddress) []byte {
	return append(append([]byte(KeyPrefixNftBid), nftIdBytes...), address.MustLengthPrefix(bidder)...)
}

func AddressBidKeyPrefix(bidder sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixAddressBid), address.MustLengthPrefix(bidder)...)
}

func AddressBidKey(nftIdBytes []byte, bidder sdk.AccAddress) []byte {
	return append(append([]byte(KeyPrefixAddressBid), address.MustLengthPrefix(bidder)...), nftIdBytes...)
}

func NftLoanKey(nftIdBytes []byte) []byte {
	return append([]byte(KeyPrefixNftLoan), nftIdBytes...)
}

func AddressNftLoanKey(addr sdk.AccAddress, nftIdBytes []byte) []byte {
	return append(append([]byte(KeyPrefixAddressNftLoan), address.MustLengthPrefix(addr)...), nftIdBytes...)
}

func AddressRewardsKey(addr sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixAddressRewards), address.MustLengthPrefix(addr)...)
}
