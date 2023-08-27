package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "testchain/testutil/keeper"
	"testchain/testutil/nullify"
	"testchain/x/kyc/keeper"
	"testchain/x/kyc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVerification(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Verification {
	items := make([]types.Verification, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetVerification(ctx, items[i])
	}
	return items
}

func TestVerificationGet(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNVerification(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVerification(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestVerificationRemove(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNVerification(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVerification(ctx,
			item.Index,
		)
		_, found := keeper.GetVerification(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestVerificationGetAll(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNVerification(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVerification(ctx)),
	)
}
