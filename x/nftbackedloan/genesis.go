package nftbackedloan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/keeper"
	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, gs types.GenesisState) {
	k.SetParamSet(ctx, gs.Params)
	for _, listing := range gs.Listings {
		k.SaveListedNft(ctx, listing)
	}
	for _, bid := range gs.Bids {
		k.SetBid(ctx, bid)
	}
	// for _, loan := range gs.Loans {
	// 	k.SetDebt(ctx, loan)
	// }
}

// ExportGenesis export genesis state for nftbackedloan module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{
		Params:   k.GetParamSet(ctx),
		Listings: k.GetAllListedNfts(ctx),
		Bids:     k.GetAllBids(ctx),
		// Loans:    k.GetAllDebts(ctx),
	}
}
