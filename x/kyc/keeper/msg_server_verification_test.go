package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/x/kyc/keeper"
	"github.com/UnUniFi/chain/x/kyc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestVerificationMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.KycKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateVerification{Sender: creator,
			Customer: strconv.Itoa(i),
		}
		_, err := srv.CreateVerification(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetVerification(ctx,
			expected.Customer,
		)
		require.True(t, found)
		require.Equal(t, expected.Customer, rst.Address)
	}
}

func TestVerificationMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteVerification
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteVerification{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteVerification{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteVerification{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.KycKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateVerification(wctx, &types.MsgCreateVerification{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteVerification(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetVerification(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
