package auction

import (
	"fmt"

	"github.com/UnUniFi/chain/x/auction/keeper"
	"github.com/UnUniFi/chain/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper, genState types.GenesisState) {
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

	k.SetParams(ctx, genState.Params)

	totalAuctionCoins := sdk.NewCoins()
	auctions, err := types.UnpackGenesisAuctions(genState.Auctions)
	if err != nil {
		panic(err)
	}
	for _, a := range auctions {
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
	balances := bankKeeper.GetAllBalances(ctx, moduleAcc.GetAddress())
	if !balances.IsEqual(totalAuctionCoins) {
		panic(fmt.Sprintf("total auction coins (%s) do not equal (%s) module account (%s) ", balances, types.ModuleName, totalAuctionCoins))
	}
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
	packed, err := types.PackGenesisAuctions(genAuctions)
	if err != nil {
		panic(err)
	}
	ret := types.NewGenesisState(nextAuctionID, params, packed)
	return &ret
}
