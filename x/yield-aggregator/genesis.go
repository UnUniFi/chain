package yield_aggregator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/keeper"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.SetStrategy(ctx, "ibc/DDDDDWWWWW", types.Strategy{
		Denom:           "ibc/DDDDDWWWWW",
		Id:              1,
		ContractAddress: "x/stake-ibc",
		Name:            "testStaking",
		GitUrl:          "",
	})

	for _, Strategies := range genState.Strategies {
		k.SetStrategy(ctx, Strategies.Denom, Strategies)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
