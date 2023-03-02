package keeper_test

import (
	"testing"

	keepertest "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/testutil/nullify"
	"github.com/UnUniFi/chain/x/yield-aggregator/keeper"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNStrategy(keeper *keeper.Keeper, ctx sdk.Context, vaultDenom string, n int) []types.Strategy {
	items := make([]types.Strategy, n)
	for i := range items {
		items[i].Id = keeper.AppendStrategy(ctx, vaultDenom, items[i])
	}
	return items
}

func TestStrategyGet(t *testing.T) {
	keeper, ctx := keepertest.YieldAggregatorKeeper(t)
	vaultDenom := "uatom"
	items := createNStrategy(keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		got, found := keeper.GetStrategy(ctx, vaultDenom, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestStrategyRemove(t *testing.T) {
	keeper, ctx := keepertest.YieldAggregatorKeeper(t)
	vaultDenom := "uatom"
	items := createNStrategy(keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		keeper.RemoveStrategy(ctx, vaultDenom, item.Id)
		_, found := keeper.GetStrategy(ctx, vaultDenom, item.Id)
		require.False(t, found)
	}
}

func TestStrategyGetAll(t *testing.T) {
	keeper, ctx := keepertest.YieldAggregatorKeeper(t)
	vaultDenom := "uatom"
	items := createNStrategy(keeper, ctx, vaultDenom, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStrategy(ctx, vaultDenom)),
	)
}

func TestStrategyCount(t *testing.T) {
	keeper, ctx := keepertest.YieldAggregatorKeeper(t)
	vaultDenom := "uatom"
	items := createNStrategy(keeper, ctx, vaultDenom, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetStrategyCount(ctx, vaultDenom))
}
