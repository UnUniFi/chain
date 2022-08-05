package decentralizedvault_test

// import (
// 	"testing"

// 	keepertest "github.com/UnUniFi/chain/testutil/keeper"
// 	"github.com/UnUniFi/chain/testutil/nullify"
// 	"github.com/UnUniFi/chain/x/decentralized-vault"
// 	"github.com/UnUniFi/chain/x/decentralized-vault/types"
// 	"github.com/stretchr/testify/require"
// )

// func TestGenesis(t *testing.T) {
// 	genesisState := types.GenesisState{
// 		Params:	types.DefaultParams(),

// 		// this line is used by starport scaffolding # genesis/test/state
// 	}

// 	k, ctx := keepertest.DecentralizedvaultKeeper(t)
// 	decentralizedvault.InitGenesis(ctx, *k, genesisState)
// 	got := decentralizedvault.ExportGenesis(ctx, *k)
// 	require.NotNil(t, got)

// 	nullify.Fill(&genesisState)
// 	nullify.Fill(got)

// 	// this line is used by starport scaffolding # genesis/test/assert
// }
