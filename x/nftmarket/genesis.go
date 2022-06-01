package nftmarket

import (
	"github.com/UnUniFi/chain/x/nftmarket/keeper"
	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, gs types.GenesisState) {
}

// ExportGenesis export genesis state for nftmarket module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	return types.GenesisState{}
}
