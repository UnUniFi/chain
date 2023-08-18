package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateStrategy{}

func NewMsgUpdateStrategy(sender string, denom string, id uint64, name, description, gitUrl string) *MsgUpdateStrategy {
	return &MsgUpdateStrategy{
		Sender:      sender,
		Denom:       denom,
		Id:          id,
		Name:        name,
		Description: description,
		GitUrl:      gitUrl,
	}
}

func (msg MsgUpdateStrategy) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
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

	if msg.GitUrl == "" {
		return sdkerrors.ErrInvalidRequest.Wrapf("empty git url")
	}

	return nil
}

func (msg MsgUpdateStrategy) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
