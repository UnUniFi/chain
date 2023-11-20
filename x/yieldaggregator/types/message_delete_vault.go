package types

import (
	// errorsmod "cosmossdk.io/errors"
	// "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteVault{}

func NewMsgDeleteVault(sender string, vaultId uint64) *MsgDeleteVault {
	return &MsgDeleteVault{
		Sender:  sender,
		VaultId: vaultId,
	}
}

func (msg MsgDeleteVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	return nil
}

func (msg MsgDeleteVault) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
