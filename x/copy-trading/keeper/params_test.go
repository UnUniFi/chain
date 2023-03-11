package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/x/copy-trading/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CopytradingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
