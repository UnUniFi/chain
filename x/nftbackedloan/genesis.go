package nftbackedloan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/keeper"
	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, gs types.GenesisState) {
	err := k.SetParams(ctx, &gs.Params)
	if err != nil {
		panic(err)
	}
	for _, listing := range gs.Listings {
		k.SaveNftListing(ctx, listing)
	}
	for _, bid := range gs.Bids {
		k.SetBid(ctx, bid)
	}
}

// ExportGenesis export genesis state for nftbackedloan module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	params, _ := k.GetParams(ctx)
	if params != nil {
		genesis.Params = *params
	}
	genesis.Listings = k.GetAllNftListings(ctx)
	genesis.Bids = k.GetAllBids(ctx)

	return genesis
}
