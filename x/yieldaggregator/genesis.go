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

	strategyCountMap := make(map[string]uint64)
	vaultCount := uint64(0)

	for _, strategy := range genState.Strategies {
		k.SetStrategy(ctx, strategy.Denom, strategy)

		if strategy.Id+1 > strategyCountMap[strategy.Denom] {
			strategyCountMap[strategy.Denom] = strategy.Id + 1
		}
	}
	for _, vault := range genState.Vaults {
		k.SetVault(ctx, vault)

		if vault.Id+1 > vaultCount {
			vaultCount = vault.Id + 1
		}
	}
	for denom, count := range strategyCountMap {
		k.SetStrategyCount(ctx, denom, count)
	}
	k.SetVaultCount(ctx, vaultCount)
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

	return genesis
}
