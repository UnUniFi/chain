package nftmarket

import (
	"github.com/UnUniFi/chain/x/nftmarket/keeper"
	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, gs types.GenesisState) {
	k.SetParamSet(ctx, gs.Params)
	for _, listing := range gs.Listings {
		k.SetNftListing(ctx, listing)
	}
	for _, bid := range gs.Bids {
		k.SetBid(ctx, bid)
	}
	for _, bid := range gs.CancelledBids {
		k.SetCancelledBid(ctx, bid)
	}
	for _, loan := range gs.Loans {
		k.SetDebt(ctx, loan)
	}
	k.SetParamSet(ctx, gs.Params)
}

// ExportGenesis export genesis state for nftmarket module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Params:        k.GetParamSet(ctx),
		Listings:      k.GetAllNftListings(ctx),
		Bids:          k.GetAllBids(ctx),
		CancelledBids: k.GetAllCancelledBids(ctx),
		Loans:         k.GetAllDebts(ctx),
	}
}
