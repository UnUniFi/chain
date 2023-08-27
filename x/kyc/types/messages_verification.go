package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateVerification = "create_verification"
	TypeMsgUpdateVerification = "update_verification"
	TypeMsgDeleteVerification = "delete_verification"
)

var _ sdk.Msg = &MsgCreateVerification{}

func NewMsgCreateVerification(
	creator string,
	index string,

) *MsgCreateVerification {
	return &MsgCreateVerification{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreateVerification) Route() string {
	return RouterKey
}

func (msg *MsgCreateVerification) Type() string {
	return TypeMsgCreateVerification
}

func (msg *MsgCreateVerification) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVerification) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVerification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateVerification{}

func NewMsgUpdateVerification(
	creator string,
	index string,

) *MsgUpdateVerification {
	return &MsgUpdateVerification{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdateVerification) Route() string {
	return RouterKey
}

func (msg *MsgUpdateVerification) Type() string {
	return TypeMsgUpdateVerification
}

func (msg *MsgUpdateVerification) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateVerification) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateVerification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteVerification{}

func NewMsgDeleteVerification(
	creator string,
	index string,

) *MsgDeleteVerification {
	return &MsgDeleteVerification{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteVerification) Route() string {
	return RouterKey
}

func (msg *MsgDeleteVerification) Type() string {
	return TypeMsgDeleteVerification
}

func (msg *MsgDeleteVerification) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteVerification) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteVerification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
