package kyc_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/testutil/nullify"
	"github.com/UnUniFi/chain/x/kyc"
	"github.com/UnUniFi/chain/x/kyc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KycKeeper(t)
	kyc.InitGenesis(ctx, *k, genesisState)
	got := kyc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProviderList, got.ProviderList)
	require.Equal(t, genesisState.ProviderCount, got.ProviderCount)
	require.ElementsMatch(t, genesisState.VerificationList, got.VerificationList)
	// this line is used by starport scaffolding # genesis/test/assert
}
