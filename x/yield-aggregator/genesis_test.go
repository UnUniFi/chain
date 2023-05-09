package yield_aggregator_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	keepertest "github.com/UnUniFi/chain/testutil/keeper"
// 	"github.com/UnUniFi/chain/testutil/nullify"
// 	"github.com/UnUniFi/chain/x/yield-aggregator"
// 	"github.com/UnUniFi/chain/x/yield-aggregator/types"
// )

// func TestGenesis(t *testing.T) {
// 	genesisState := types.GenesisState{
// 		Params: types.DefaultParams(),

// 		// this line is used by starport scaffolding # genesis/test/state
// 	}

// 	k, ctx := keepertest.YieldAggregatorKeeper(t)
// 	yield_aggregator.InitGenesis(ctx, *k, genesisState)
// 	got := yield_aggregator.ExportGenesis(ctx, *k)
// 	require.NotNil(t, got)

// 	nullify.Fill(&genesisState)
// 	nullify.Fill(got)

// 	// this line is used by starport scaffolding # genesis/test/assert
// }
