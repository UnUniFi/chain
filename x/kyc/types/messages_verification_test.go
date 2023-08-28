package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/testutil/sample"
)

func TestMsgCreateVerification_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateVerification
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVerification{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateVerification{
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

func TestMsgUpdateVerification_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateVerification
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateVerification{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateVerification{
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

func TestMsgDeleteVerification_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteVerification
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteVerification{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteVerification{
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
