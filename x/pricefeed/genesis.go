package pricefeed

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/pricefeed/keeper"
	"github.com/UnUniFi/chain/x/pricefeed/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set the markets and oracles from params
	k.SetParams(ctx, genState.Params)

	// Iterate through the posted prices and set them in the store if they are not expired
	for _, pp := range genState.PostedPrices {
		if pp.Expiry.After(ctx.BlockTime()) {
			oracleAddress, err := sdk.AccAddressFromBech32(pp.OracleAddress)
			if err != nil {
				panic(err)
			}
			_, err = k.SetPrice(ctx, oracleAddress, pp.MarketId, pp.Price, pp.Expiry)
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
		rps := k.GetRawPrices(ctx, market.MarketId)

		if len(rps) == 0 {
			continue
		}
		err := k.SetCurrentPrices(ctx, market.MarketId)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	// Get the params for markets and oracles
	params := k.GetParams(ctx)

	var postedPrices []types.PostedPrice
	for _, market := range k.GetMarkets(ctx) {
		pp := k.GetRawPrices(ctx, market.MarketId)
		postedPrices = append(postedPrices, pp...)
	}
	return types.NewGenesisState(params, postedPrices)
}
