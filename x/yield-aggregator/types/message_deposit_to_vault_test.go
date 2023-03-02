package types

import (
	"testing"

	"github.com/UnUniFi/chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgDepositToVault_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDepositToVault
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDepositToVault{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDepositToVault{
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
