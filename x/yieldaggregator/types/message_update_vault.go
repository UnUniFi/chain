package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateVault{}

func NewMsgUpdateVault(sender string, id uint64, denom, name, description, feeCollector string) *MsgUpdateVault {
	return &MsgUpdateVault{
		Sender:              sender,
		Id:                  id,
		Denom:               denom,
		Name:                name,
		Description:         description,
		FeeCollectorAddress: feeCollector,
	}
}

func (msg MsgUpdateVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.FeeCollectorAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid fee collector address: %s", err)
	}

	if msg.Denom == "" {
		return sdkerrors.ErrInvalidRequest.Wrapf("empty denom")
	}

	if msg.Name == "" {
		return sdkerrors.ErrInvalidRequest.Wrapf("empty name")
	}

	if msg.Description == "" {
		return sdkerrors.ErrInvalidRequest.Wrapf("empty description")
	}

	return nil
}

func (msg MsgUpdateVault) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
