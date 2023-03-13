package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/x/copy-trading/keeper"
	"github.com/UnUniFi/chain/x/copy-trading/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTracingMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CopytradingKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	sender := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateTracing{
			Sender: sender,
		}
		_, err := srv.CreateTracing(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetTracing(ctx,
			expected.Sender,
		)
		require.True(t, found)
		require.Equal(t, expected.Sender, rst.Address)
	}
}

func TestTracingMsgServerDelete(t *testing.T) {
	sender := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteTracing
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteTracing{
				Sender: sender,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteTracing{
				Sender: "B",
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteTracing{
				Sender: sender,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CopytradingKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateTracing(wctx, &types.MsgCreateTracing{
				Sender: sender,
			})
			require.NoError(t, err)
			_, err = srv.DeleteTracing(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetTracing(ctx,
					tc.request.Sender,
				)
				require.False(t, found)
			}
		})
	}
}
