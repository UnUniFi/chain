package irs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/keeper"
	"github.com/UnUniFi/chain/x/irs/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, &genState.Params)

	for _, vault := range genState.Vaults {
		k.SetVault(ctx, vault)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	params, _ := k.GetParams(ctx)
	if params != nil {
		genesis.Params = *params
	}
	genesis.Vaults = k.GetAllVault(ctx)

	return genesis
}
