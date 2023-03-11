package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/testutil/nullify"
	"github.com/UnUniFi/chain/x/copy-trading/keeper"
	"github.com/UnUniFi/chain/x/copy-trading/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTracing(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Tracing {
	items := make([]types.Tracing, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetTracing(ctx, items[i])
	}
	return items
}

func TestTracingGet(t *testing.T) {
	keeper, ctx := keepertest.CopytradingKeeper(t)
	items := createNTracing(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTracing(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTracingRemove(t *testing.T) {
	keeper, ctx := keepertest.CopytradingKeeper(t)
	items := createNTracing(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTracing(ctx,
			item.Address,
		)
		_, found := keeper.GetTracing(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestTracingGetAll(t *testing.T) {
	keeper, ctx := keepertest.CopytradingKeeper(t)
	items := createNTracing(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTracing(ctx)),
	)
}
