package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "nftmint"

	// ChainName defines the chain name to put StoreKey for avoiding the KVStore collision to sdk's nft module KVStore key
	ChainName = "ununifi"

	// StoreKey defines the primary module store key
	StoreKey = ChainName + ModuleName

	// RouterKey is the message route for nftmint
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines thee in-memory store key
	MemStoreKey = "mem_nftmint"
)

var (
	// KeyPrefixClassAttributes defines prefix key for ClassAttributes
	KeyPrefixClassAttributes = []byte{0x01}

	// KeyPrefixNFTAttributes defines prefix key for NFTAttributes
	KeyPrefixNFTMinter = []byte{0x02}

	// KeyPrefixOwningClassList defines prefix key for OwningClassList
	KeyPrefixOwningClassIdList = []byte{0x03}

	// KeyPrefixClassNameIdList defines prefix key for ClassNameIdList
	KeyPrefixClassNameIdList = []byte{0x04}
)

func NFTMinterKey(classID, nftID string) []byte {
	nftIdentifier := classID + nftID
	return []byte(nftIdentifier)
}

func OwningClassIdListKey(owner sdk.AccAddress) []byte {
	ownerAddr, _ := sdk.AccAddressFromBech32(owner.String())
	return ownerAddr.Bytes()
}
