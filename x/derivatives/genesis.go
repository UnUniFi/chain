package derivatives

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	// todo load genesis state when restart
	for _, asset := range genState.Params.PoolParams.AcceptedAssets {
		k.AddPoolAsset(ctx, *asset)
	}

	if err := k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), genState.PoolMarketCap); err != nil {
		panic(err)
	}

	for _, market := range genState.Params.PerpetualFutures.Markets {
		// set initial net position
		k.SetPerpetualFuturesNetPositionOfMarket(ctx, *market, sdk.NewDec(0))
	}

	for _, position := range genState.Positions {
		k.SetPosition(ctx, position)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
