package types

import (
	"testing"

	"github.com/UnUniFi/chain/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgWithdrawFromVault_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgWithdrawFromVault
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgWithdrawFromVault{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgWithdrawFromVault{
				Sender:        sample.AccAddress(),
				VaultId:       1,
				LpTokenAmount: sdk.NewInt(1000),
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
