package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/UnUniFi/chain/x/ununifidist/simulation"
	"github.com/UnUniFi/chain/x/ununifidist/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"
)

func TestDecodeStore(t *testing.T) {
	prevBlockTime := time.Now().UTC()
	bPrevBlockTime, err := prevBlockTime.MarshalBinary()
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	if err != nil {
		panic(err)
	}
	dec := simulation.NewDecodeStore(cdc)

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   []byte(types.PreviousBlockTimeKey),
				Value: bPrevBlockTime,
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"PreviousBlockTime", fmt.Sprintf("%s\n%s", prevBlockTime, prevBlockTime)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
