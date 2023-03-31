package types

import (
	// errorsmod "cosmossdk.io/errors"
	// "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferVaultOwnership = "transfer-vault-ownership"

var _ sdk.Msg = &MsgCreateVault{}

func NewMsgTransferVaultOwnership(sender string, vaultId uint64, recipient string) *MsgTransferVaultOwnership {
	return &MsgTransferVaultOwnership{
		Sender:    sender,
		VaultId:   vaultId,
		Recipient: recipient,
	}
}

func (msg MsgTransferVaultOwnership) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid recipient address: %s", err)
	}

	return nil
}

func (msg *MsgTransferVaultOwnership) Route() string {
	return RouterKey
}

func (msg *MsgTransferVaultOwnership) Type() string {
	return TypeMsgDepositToVault
}

func (msg MsgTransferVaultOwnership) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgTransferVaultOwnership) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
