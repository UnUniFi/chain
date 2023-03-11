package copytrading

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/copy-trading/keeper"
	"github.com/UnUniFi/chain/x/copy-trading/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the exemplaryTrader
	for _, elem := range genState.ExemplaryTraderList {
		k.SetExemplaryTrader(ctx, elem)
	}
	// Set all the tracing
	for _, elem := range genState.TracingList {
		k.SetTracing(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ExemplaryTraderList = k.GetAllExemplaryTrader(ctx)
	genesis.TracingList = k.GetAllTracing(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
