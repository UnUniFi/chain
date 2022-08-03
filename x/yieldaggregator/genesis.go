package yieldaggregator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, acc := range genState.AssetManagementAccounts {
		k.SetAssetManagementAccount(ctx, acc)
	}

	for _, target := range genState.AssetManagementTargets {
		k.SetAssetManagementTarget(ctx, target)
	}

	for _, order := range genState.FarmingOrders {
		k.SetFarmingOrder(ctx, order)
	}

	// TODO: set FarmingUnits
	// TODO: set UserInfos
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.AssetManagementAccounts = k.GetAllAssetManagementAccounts(ctx)
	genesis.AssetManagementTargets = k.GetAllAssetManagementTargets(ctx)
	genesis.FarmingOrders = k.GetAllFarmingOrders(ctx)

	// TODO: set FarmingUnits
	// TODO: set UserInfos
	return genesis
}
