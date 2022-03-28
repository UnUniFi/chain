package simulation

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/UnUniFi/chain/x/auction/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding auction type
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal([]byte(kvA.Key), []byte(types.AuctionKey)):
			var auctionAny codectypes.Any
			auctionAny.Unmarshal(kvA.Value)
			auction := types.MustUnpackAuction(&auctionAny)
			return fmt.Sprintf("%v\n%v", auction, auction)

		case bytes.Equal([]byte(kvA.Key), []byte(types.AuctionByTimeKey)),
			bytes.Equal([]byte(kvA.Key), []byte(types.NextAuctionIDKey)):
			auctionIDA := binary.BigEndian.Uint64(kvA.Value)
			auctionIDB := binary.BigEndian.Uint64(kvB.Value)
			return fmt.Sprintf("%d\n%d", auctionIDA, auctionIDB)

		default:
			panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
		}
	}
}
