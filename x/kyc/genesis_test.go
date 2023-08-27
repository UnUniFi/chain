package kyc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "testchain/testutil/keeper"
	"testchain/testutil/nullify"
	"testchain/x/kyc"
	"testchain/x/kyc/types"
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
				Index: "0",
			},
			{
				Index: "1",
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
