package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCdp{}

func NewMsgCreateCdp(creator string) *MsgCreateCdp {
	return &MsgCreateCdp{
		Creator: creator,
	}
}

func (msg *MsgCreateCdp) Route() string {
	return RouterKey
}

func (msg *MsgCreateCdp) Type() string {
	return "CreateCdp"
}

func (msg *MsgCreateCdp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCdp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCdp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCdp{}

func NewMsgUpdateCdp(creator string, id string) *MsgUpdateCdp {
	return &MsgUpdateCdp{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgUpdateCdp) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCdp) Type() string {
	return "UpdateCdp"
}

func (msg *MsgUpdateCdp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCdp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCdp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateCdp{}

func NewMsgDeleteCdp(creator string, id string) *MsgDeleteCdp {
	return &MsgDeleteCdp{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteCdp) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCdp) Type() string {
	return "DeleteCdp"
}

func (msg *MsgDeleteCdp) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCdp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCdp) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
