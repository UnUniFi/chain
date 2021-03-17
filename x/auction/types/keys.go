package types

import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "auction"

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
	ParamsKey        = "Params-value-"
	AuctionKey       = "Auction-value-"
	AuctionCountKey  = "Auction-count-"
	AuctionByTimeKey = "Auction-by-time-"
	NextAuctionIDKey = "NextAuctionID-value-"
)

// GetAuctionKey returns the bytes of an auction key
func GetAuctionKey(auctionID uint64) []byte {
	return Uint64ToBytes(auctionID)
}

// GetAuctionByTimeKey returns the key for iterating auctions by time
func GetAuctionByTimeKey(endTime time.Time, auctionID uint64) []byte {
	return append(sdk.FormatTimeBytes(endTime), Uint64ToBytes(auctionID)...)
}

// Uint64ToBytes converts a uint64 into fixed length bytes for use in store keys.
func Uint64ToBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(id))
	return bz
}

// Uint64FromBytes converts some fixed length bytes back into a uint64.
func Uint64FromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
