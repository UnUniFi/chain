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

	for _, reward := range genState.DailyPercents {
		k.SetDailyRewardPercent(ctx, reward)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := &types.GenesisState{
		Params:                  k.GetParams(ctx),
		AssetManagementAccounts: k.GetAllAssetManagementAccounts(ctx),
		AssetManagementTargets:  k.GetAllAssetManagementTargets(ctx),
		FarmingOrders:           k.GetAllFarmingOrders(ctx),
		FarmingUnits:            k.GetAllFarmingUnits(ctx),
		UserDeposits:            k.GetAllUserDeposits(ctx),
		DailyPercents:           k.GetAllDailyRewardPercents(ctx),
	}

	return genesis
}
