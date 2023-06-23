package keeper_test

import (
	"testing"

	testkeeper "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// TODO: add test for followings after full implementation
// LiquidityProviderTokenRealAPY
// LiquidityProviderTokenNominalAPY
// PerpetualFutures
// PerpetualFuturesMarket
// PerpetualOptions
// Pool

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.DerivativesKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
