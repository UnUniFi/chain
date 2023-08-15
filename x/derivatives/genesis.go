package derivatives

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	if err := k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), genState.PoolMarketCap); err != nil {
		panic(err)
	}

	if len(genState.PerpetualFuturesGrossPositionOfMarket) > 0 {
		for _, perpetualFuturesGrossPositionOfMarket := range genState.PerpetualFuturesGrossPositionOfMarket {
			k.SetPerpetualFuturesGrossPositionOfMarket(ctx, perpetualFuturesGrossPositionOfMarket)
		}
	}
	initialPerpetualFuturesGrossPositionOfMarkets := types.GetMarketsOutOfPerpetualFuturesGrossPositionOfMarket(genState.PerpetualFuturesGrossPositionOfMarket)
	for _, market := range genState.Params.PerpetualFutures.Markets {
		// set initial net position
		if !market.InMarketSet(initialPerpetualFuturesGrossPositionOfMarkets) {
			// Position reference for Long
			perpetualFuturesGrossPositionOfMarketLong := types.NewPerpetualFuturesGrossPositionOfMarket(*market, types.PositionType_LONG, sdk.ZeroInt())
			k.SetPerpetualFuturesGrossPositionOfMarket(ctx, perpetualFuturesGrossPositionOfMarketLong)

			// Position reference for Short
			perpetualFuturesGrossPositionOfMarketShort := types.NewPerpetualFuturesGrossPositionOfMarket(*market, types.PositionType_SHORT, sdk.ZeroInt())
			k.SetPerpetualFuturesGrossPositionOfMarket(ctx, perpetualFuturesGrossPositionOfMarketShort)
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
	genesis.PerpetualFuturesGrossPositionOfMarket = k.GetAllPerpetualFuturesGrossPositionOfMarket(ctx)

	return genesis
}
