package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

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
	KeyPrefixAddressNftLoan = "address_nft_loan"
	// rewards by address
	KeyPrefixAddressRewards = "rewards"
)

func NftBytes(classId, nftId uint64) []byte {
	return append(sdk.Uint64ToBigEndian(classId), sdk.Uint64ToBigEndian(nftId)...)
}

func NftListingKey(classId, nftId uint64) []byte {
	return append([]byte(KeyPrefixNftListing), NftBytes(classId, nftId)...)
}

func NftAddressNftListingKey(addr sdk.AccAddress, classId, nftId uint64) []byte {
	return append(append([]byte(KeyPrefixAddressNftListing), address.MustLengthPrefix(addr)...), NftBytes(classId, nftId)...)
}

func NftBidKey(classId, nftId uint64, bidder sdk.AccAddress) []byte {
	return append(append([]byte(KeyPrefixNftBid), NftBytes(classId, nftId)...), address.MustLengthPrefix(bidder)...)
}

func AddressBidKey(classId, nftId uint64, bidder sdk.AccAddress) []byte {
	return append(append([]byte(KeyPrefixAddressBid), address.MustLengthPrefix(bidder)...), NftBytes(classId, nftId)...)
}

func NftLoanKey(classId, nftId uint64) []byte {
	return append([]byte(KeyPrefixNftLoan), NftBytes(classId, nftId)...)
}

func AddressNftLoanKey(addr sdk.AccAddress, classId, nftId uint64) []byte {
	return append(append([]byte(KeyPrefixAddressNftLoan), address.MustLengthPrefix(addr)...), NftBytes(classId, nftId)...)
}

func AddressRewardsKey(addr sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixAddressRewards), address.MustLengthPrefix(addr)...)
}
