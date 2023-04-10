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

	if len(genState.PerpetualFuturesNetPositionOfMarket) > 0 {
		for _, perpetualFuturesNetPositionOfMarket := range genState.PerpetualFuturesNetPositionOfMarket {
			k.SetPerpetualFuturesNetPositionOfMarket(ctx, perpetualFuturesNetPositionOfMarket)
		}
	}
	initialPerpetualFuturesNetPositionOfMarkets := types.GetMarketsOutOfPerpetualFuturesNetPositionOfMarket(genState.PerpetualFuturesNetPositionOfMarket)
	for _, market := range genState.Params.PerpetualFutures.Markets {
		// set initial net position
		if !market.InMarketSet(initialPerpetualFuturesNetPositionOfMarkets) {
			// Position reference for Long
			perpetualFuturesNetPositionOfMarketLong := types.NewPerpetualFuturesNetPositionOfMarket(*market, types.PositionType_LONG, sdk.ZeroInt())
			k.SetPerpetualFuturesNetPositionOfMarket(ctx, perpetualFuturesNetPositionOfMarketLong)

			// Position reference for Short
			perpetualFuturesNetPositionOfMarketShort := types.NewPerpetualFuturesNetPositionOfMarket(*market, types.PositionType_SHORT, sdk.ZeroInt())
			k.SetPerpetualFuturesNetPositionOfMarket(ctx, perpetualFuturesNetPositionOfMarketShort)
		}
	}

	for _, position := range genState.Positions {
		k.SetPosition(ctx, position)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Positions = k.GetAllPositions(ctx)
	genesis.PoolMarketCap = k.GetPoolMarketCapSnapshot(ctx, ctx.BlockHeight())
	genesis.PerpetualFuturesNetPositionOfMarket = k.GetAllPerpetualFuturesNetPositionOfMarket(ctx)

	return genesis
}
