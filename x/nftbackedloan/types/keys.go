package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "nftbackedloan"

	// Module account for nft trading fee collection
	// use ecosystem-incentive module account for now
	// [unused] NftTradingFee = "nfttradingfee"

	// StoreKey defines the primary module store key
	StoreKey = "_" + ModuleName

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
	// active nft listing by end time
	KeyPrefixEndTimeNftListing = "end_time_nft_listing"
	// full payment period listing by end time
	KeyPrefixFullPaymentPeriodListing = "full_payment_period_listing"
	// successful bid listing by nft send time
	KeyPrefixSuccessfulBidListing = "successful_bid_listing"
	// nft bid by nft_id
	KeyPrefixNftBid = "nft_bid_original"
	// nft bid cancelled
	KeyPrefixNftBidCancelled = "nft_bid_cancelled"
	// nft bid by owner
	KeyPrefixAddressBid = "address_bid"
	// nft bid by end time
	KeyPrefixEndTimeNftBid = "end_time_nft_bid"
	// nft loan by nft_id
	KeyPrefixNftLoan = "nft_loan"
	// nft loan by owner
	KeyPrefixAddressNftLoan = "address_nft_loan"
	// rewards by address
	KeyPrefixAddressRewards = "rewards"
	// class by address
	KeyPrefixClass = "class"
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

func ClassKey(addr sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixClass), address.MustLengthPrefix(addr)...)
}

func ClassIdKey(classId string) []byte {
	return ClassKey([]byte(classId))
}

func NftBidBytes(classId, nftId, bidder string) []byte {
	return append(append(NftBytes(classId, nftId), byte(0xFF)), []byte(bidder)...)
}

func NftBidBytesToBidId(NftBidBytes []byte) BidId {
	stringList := []string{}
	chunk := []byte{}
	separate := byte(0xFF)
	for _, idChar := range NftBidBytes {
		if idChar == separate {
			stringList = append(stringList, string(chunk))
			chunk = []byte{}
			continue
		}
		chunk = append(chunk, idChar)
	}
	stringList = append(stringList, string(chunk))
	if len(stringList) != 3 {
		panic("nft bid byte wrong format")
	}

	return BidId{
		NftId: &NftIdentifier{
			ClassId: stringList[0],
			NftId:   stringList[1],
		},
		Bidder: stringList[2],
	}
}
