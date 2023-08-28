package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/x/kyc/types"
)

func TestProviderMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateProvider(ctx, &types.MsgCreateProvider{Sender: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestProviderMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateProvider
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateProvider{Sender: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateProvider{Sender: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateProvider{Sender: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateProvider(ctx, &types.MsgCreateProvider{Sender: creator})
			require.NoError(t, err)

			_, err = srv.UpdateProvider(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
