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

	for _, unit := range genState.FarmingUnits {
		k.SetFarmingUnit(ctx, unit)
	}

	for _, deposit := range genState.UserDeposits {
		addr := sdk.MustAccAddressFromBech32(deposit.User)
		k.SetUserDeposit(ctx, addr, deposit.Amount)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.AssetManagementAccounts = k.GetAllAssetManagementAccounts(ctx)
	genesis.AssetManagementTargets = k.GetAllAssetManagementTargets(ctx)
	genesis.FarmingOrders = k.GetAllFarmingOrders(ctx)
	genesis.FarmingUnits = k.GetAllFarmingUnits(ctx)
	genesis.UserDeposits = k.GetAllUserDeposits(ctx)

	return genesis
}
