package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/testutil/sample"
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
				Sender:  sample.AccAddress(),
				VaultId: 1,
				Amount:  sdk.NewCoin("uatom", sdk.NewInt(1000)),
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
