package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterSymbolInfos{}

func NewMsgRegisterSymbolInfos(sender string, info []SymbolInfo) *MsgRegisterSymbolInfos {
	return &MsgRegisterSymbolInfos{
		Sender: sender,
		Info:   info,
	}
}

func (msg MsgRegisterSymbolInfos) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	return nil
}

func (msg MsgRegisterSymbolInfos) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
