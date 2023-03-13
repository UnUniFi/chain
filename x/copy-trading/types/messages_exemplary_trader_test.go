package types

import (
	"testing"

	"github.com/UnUniFi/chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateExemplaryTrader_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateExemplaryTrader
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateExemplaryTrader{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateExemplaryTrader{
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

func TestMsgUpdateExemplaryTrader_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateExemplaryTrader
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateExemplaryTrader{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateExemplaryTrader{
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

func TestMsgDeleteExemplaryTrader_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteExemplaryTrader
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteExemplaryTrader{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteExemplaryTrader{
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
