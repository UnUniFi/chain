package pricefeed

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/pricefeed/keeper"
	"github.com/lcnem/jpyx/x/pricefeed/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set the markets and oracles from params
	k.SetParams(ctx, genState.Params)

	// Iterate through the posted prices and set them in the store if they are not expired
	for _, pp := range genState.PostedPrices {
		if pp.Expiry.After(ctx.BlockTime()) {
			_, err := k.SetPrice(ctx, pp.OracleAddress.AccAddress(), pp.MarketId, pp.Price, pp.Expiry)
			if err != nil {
				panic(err)
			}
		}
	}
	params := k.GetParams(ctx)

	// Set the current price (if any) based on what's now in the store
	for _, market := range params.Markets {
		if !market.Active {
			continue
		}
		rps, err := k.GetRawPrices(ctx, market.MarketId)
		if err != nil {
			panic(err)
		}
		if len(rps) == 0 {
			continue
		}
		err = k.SetCurrentPrices(ctx, market.MarketId)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	params := k.GetParams(ctx)

	var postedPrices []types.PostedPrice
	for _, market := range k.GetMarkets(ctx) {
		pp, err := k.GetRawPrices(ctx, market.MarketId)
		if err != nil {
			panic(err)
		}
		postedPrices = append(postedPrices, pp...)
	}
	return types.NewGenesisState(params, postedPrices)
}
