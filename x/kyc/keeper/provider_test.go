package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "testchain/testutil/keeper"
	"testchain/testutil/nullify"
	"testchain/x/kyc/keeper"
	"testchain/x/kyc/types"
)

func createNProvider(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Provider {
	items := make([]types.Provider, n)
	for i := range items {
		items[i].Id = keeper.AppendProvider(ctx, items[i])
	}
	return items
}

func TestProviderGet(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNProvider(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetProvider(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestProviderRemove(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNProvider(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProvider(ctx, item.Id)
		_, found := keeper.GetProvider(ctx, item.Id)
		require.False(t, found)
	}
}

func TestProviderGetAll(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNProvider(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProvider(ctx)),
	)
}

func TestProviderCount(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNProvider(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetProviderCount(ctx))
}
