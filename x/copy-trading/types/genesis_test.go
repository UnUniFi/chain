package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/x/copy-trading/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ExemplaryTraderList: []types.ExemplaryTrader{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				TracingList: []types.Tracing{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated exemplaryTrader",
			genState: &types.GenesisState{
				ExemplaryTraderList: []types.ExemplaryTrader{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated tracing",
			genState: &types.GenesisState{
				TracingList: []types.Tracing{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
