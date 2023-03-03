package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/testutil/nullify"
	"github.com/UnUniFi/chain/x/yield-aggregator/keeper"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

func createNVault(keeper *keeper.Keeper, ctx sdk.Context, denom string, n int) []types.Vault {
	items := make([]types.Vault, n)
	for i := range items {
		items[i] = types.Vault{
			Denom:          denom,
			CommissionRate: sdk.MustNewDecFromStr("0.001"),
		}
		items[i].Id = keeper.AppendVault(ctx, items[i])
	}
	return items
}

func TestVaultGet(t *testing.T) {
	keeper, ctx := keepertest.YieldAggregatorKeeper(t)
	vaultDenom := "uatom"
	items := createNVault(keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		got, found := keeper.GetVault(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVaultRemove(t *testing.T) {
	keeper, ctx := keepertest.YieldAggregatorKeeper(t)
	vaultDenom := "uatom"
	items := createNVault(keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		keeper.RemoveVault(ctx, item.Id)
		_, found := keeper.GetVault(ctx, item.Id)
		require.False(t, found)
	}
}

func TestVaultGetAll(t *testing.T) {
	keeper, ctx := keepertest.YieldAggregatorKeeper(t)
	vaultDenom := "uatom"
	items := createNVault(keeper, ctx, vaultDenom, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVault(ctx)),
	)
}

func TestVaultCount(t *testing.T) {
	keeper, ctx := keepertest.YieldAggregatorKeeper(t)
	vaultDenom := "uatom"
	items := createNVault(keeper, ctx, vaultDenom, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVaultCount(ctx))
}
