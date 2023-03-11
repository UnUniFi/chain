package copytrading_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/testutil/nullify"
	copytrading "github.com/UnUniFi/chain/x/copy-trading"
	"github.com/UnUniFi/chain/x/copy-trading/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CopytradingKeeper(t)
	copytrading.InitGenesis(ctx, *k, genesisState)
	got := copytrading.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ExemplaryTraderList, got.ExemplaryTraderList)
	require.ElementsMatch(t, genesisState.TracingList, got.TracingList)
	// this line is used by starport scaffolding # genesis/test/assert
}
