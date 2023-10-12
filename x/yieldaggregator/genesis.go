package yieldaggregator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, &genState.Params)

	for _, strategy := range genState.Strategies {
		k.AppendStrategy(ctx, strategy.Denom, strategy)
	}
	for _, vault := range genState.Vaults {
		k.AppendVault(ctx, vault)
	}
	for _, deposit := range genState.PendingDeposits {
		k.SetPendingDeposit(ctx, deposit.VaultId, deposit.Amount)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	params, _ := k.GetParams(ctx)
	if params != nil {
		genesis.Params = *params
	}
	genesis.Strategies = k.GetAllStrategy(ctx, "")
	genesis.Vaults = k.GetAllVault(ctx)
	genesis.PendingDeposits = k.GetAllPendingDeposits(ctx)

	return genesis
}
