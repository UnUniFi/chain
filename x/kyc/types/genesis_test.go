package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/x/kyc/types"
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

				ProviderList: []types.Provider{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ProviderCount: 2,
				VerificationList: []types.Verification{
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
			desc: "duplicated provider",
			genState: &types.GenesisState{
				ProviderList: []types.Provider{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid provider count",
			genState: &types.GenesisState{
				ProviderList: []types.Provider{
					{
						Id: 1,
					},
				},
				ProviderCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated verification",
			genState: &types.GenesisState{
				VerificationList: []types.Verification{
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
