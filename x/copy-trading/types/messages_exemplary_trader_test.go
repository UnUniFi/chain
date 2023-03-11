package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/UnUniFi/chain/testutil/sample"
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
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateExemplaryTrader{
				Creator: sample.AccAddress(),
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
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateExemplaryTrader{
				Creator: sample.AccAddress(),
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
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteExemplaryTrader{
				Creator: sample.AccAddress(),
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
