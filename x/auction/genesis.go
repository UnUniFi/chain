package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/auction/keeper"
	"github.com/lcnem/jpyx/x/auction/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the auction
	for _, elem := range genState.AuctionList {
		k.SetAuction(ctx, *elem)
	}

	// Set auction count
	k.SetAuctionCount(ctx, int64(len(genState.AuctionList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all auction
	auctionList := k.GetAllAuction(ctx)
	for _, elem := range auctionList {
		elem := elem
		genesis.AuctionList = append(genesis.AuctionList, &elem)
	}

	return genesis
}
