package simulation

import (
	"bytes"
	"fmt"
	"time"

	"github.com/UnUniFi/chain/x/ununifidist/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal([]byte(kvA.Key[:1]), []byte(types.PreviousBlockTimeKey)):
			var timeA, timeB time.Time
			timeA.UnmarshalBinary(kvA.Value)
			timeB.UnmarshalBinary(kvB.Value)
			return fmt.Sprintf("%s\n%s", timeA, timeB)

		default:
			panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
		}
	}
}
