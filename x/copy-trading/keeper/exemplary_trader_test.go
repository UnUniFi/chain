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

func createNExemplaryTrader(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ExemplaryTrader {
	items := make([]types.ExemplaryTrader, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetExemplaryTrader(ctx, items[i])
	}
	return items
}

func TestExemplaryTraderGet(t *testing.T) {
	keeper, ctx := keepertest.CopytradingKeeper(t)
	items := createNExemplaryTrader(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExemplaryTrader(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestExemplaryTraderRemove(t *testing.T) {
	keeper, ctx := keepertest.CopytradingKeeper(t)
	items := createNExemplaryTrader(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExemplaryTrader(ctx,
			item.Address,
		)
		_, found := keeper.GetExemplaryTrader(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestExemplaryTraderGetAll(t *testing.T) {
	keeper, ctx := keepertest.CopytradingKeeper(t)
	items := createNExemplaryTrader(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExemplaryTrader(ctx)),
	)
}
