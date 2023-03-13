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

func TestExemplaryTraderMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CopytradingKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	sender := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateExemplaryTrader{
			Sender: sender,
		}
		_, err := srv.CreateExemplaryTrader(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetExemplaryTrader(ctx,
			expected.Sender,
		)
		require.True(t, found)
		require.Equal(t, expected.Sender, rst.Address)
	}
}

func TestExemplaryTraderMsgServerUpdate(t *testing.T) {
	sender := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateExemplaryTrader
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateExemplaryTrader{
				Sender: sender,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateExemplaryTrader{
				Sender: "B",
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateExemplaryTrader{
				Sender: sender,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CopytradingKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateExemplaryTrader{
				Sender: sender,
			}
			_, err := srv.CreateExemplaryTrader(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateExemplaryTrader(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetExemplaryTrader(ctx,
					expected.Sender,
				)
				require.True(t, found)
				require.Equal(t, expected.Sender, rst.Address)
			}
		})
	}
}

func TestExemplaryTraderMsgServerDelete(t *testing.T) {
	sender := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteExemplaryTrader
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteExemplaryTrader{
				Sender: sender,
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteExemplaryTrader{
				Sender: "B",
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteExemplaryTrader{
				Sender: sender,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.CopytradingKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateExemplaryTrader(wctx, &types.MsgCreateExemplaryTrader{
				Sender: sender,
			})
			require.NoError(t, err)
			_, err = srv.DeleteExemplaryTrader(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetExemplaryTrader(ctx,
					tc.request.Sender,
				)
				require.False(t, found)
			}
		})
	}
}
