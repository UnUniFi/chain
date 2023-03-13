package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTracing = "create_tracing"
	TypeMsgDeleteTracing = "delete_tracing"
)

var _ sdk.Msg = &MsgCreateTracing{}

func NewMsgCreateTracing(
	sender string,
) *MsgCreateTracing {
	return &MsgCreateTracing{
		Sender: sender,
	}
}

func (msg *MsgCreateTracing) Route() string {
	return RouterKey
}

func (msg *MsgCreateTracing) Type() string {
	return TypeMsgCreateTracing
}

func (msg *MsgCreateTracing) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgCreateTracing) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTracing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTracing{}

func NewMsgDeleteTracing(
	sender string,

) *MsgDeleteTracing {
	return &MsgDeleteTracing{
		Sender: sender,
	}
}
func (msg *MsgDeleteTracing) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTracing) Type() string {
	return TypeMsgDeleteTracing
}

func (msg *MsgDeleteTracing) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgDeleteTracing) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTracing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
