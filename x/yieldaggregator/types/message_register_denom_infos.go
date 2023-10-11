package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterDenomInfos{}

func NewMsgRegisterDenomInfos(sender string, info []DenomInfo) *MsgRegisterDenomInfos {
	return &MsgRegisterDenomInfos{
		Sender: sender,
		Info:   info,
	}
}

func (msg MsgRegisterDenomInfos) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	return nil
}

func (msg MsgRegisterDenomInfos) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
