package auction

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/auction/keeper"
	"github.com/lcnem/jpyx/x/auction/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the auction
	// for _, elem := range genState.AuctionList {
	// 	k.SetAuction(ctx, *elem)
	// }

	// // Set auction count
	// k.SetAuctionCount(ctx, int64(len(genState.AuctionList)))
	if err := genState.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	k.SetNextAuctionID(ctx, genState.NextAuctionId)

	k.SetParams(ctx, *genState.Params)

	totalAuctionCoins := sdk.NewCoins()
	for _, a := range genState.Auctions {
		k.SetAuction(ctx, a)
		// find the total coins that should be present in the module account
		totalAuctionCoins = totalAuctionCoins.Add(a.GetModuleAccountCoins()...)
	}

	// check if the module account exists
	moduleAcc := accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
	// check module coins match auction coins
	// Note: Other sdk modules do not check this, instead just using the existing module account coins, or if zero, setting them.
	// if !moduleAcc.GetCoins().IsEqual(totalAuctionCoins) {
	// 	panic(fmt.Sprintf("total auction coins (%s) do not equal (%s) module account (%s) ", moduleAcc.GetCoins(), ModuleName, totalAuctionCoins))
	// }
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all auction
	// auctionList := k.GetAllAuction(ctx)
	// for _, elem := range auctionList {
	// 	elem := elem
	// 	genesis.AuctionList = append(genesis.AuctionList, &elem)
	// }

	// return genesis
	nextAuctionID, err := k.GetNextAuctionID(ctx)
	if err != nil {
		panic(err)
	}

	params := k.GetParams(ctx)

	genAuctions := types.GenesisAuctions{} // return empty list instead of nil if no auctions
	k.IterateAuctions(ctx, func(a types.Auction) bool {
		ga, ok := a.(types.GenesisAuction)
		if !ok {
			panic("could not convert stored auction to GenesisAuction type")
		}
		genAuctions = append(genAuctions, ga)
		return false
	})
	ret := types.NewGenesisState(nextAuctionID, params, genAuctions)
	return &ret
}
