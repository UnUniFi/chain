package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/testutil/sample"
)

func TestMsgCreateProvider_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateProvider
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateProvider{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateProvider{
				Sender: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateProvider_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateProvider
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateProvider{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateProvider{
				Sender: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteProvider_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteProvider
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteProvider{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteProvider{
				Sender: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
