package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/testutil/nullify"
	"github.com/UnUniFi/chain/x/kyc/keeper"
	"github.com/UnUniFi/chain/x/kyc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVerification(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Verification {
	items := make([]types.Verification, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetVerification(ctx, items[i])
	}
	return items
}

func TestVerificationGet(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNVerification(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVerification(ctx,
			item.Address,
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
			item.Address,
		)
		_, found := keeper.GetVerification(ctx,
			item.Address,
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
