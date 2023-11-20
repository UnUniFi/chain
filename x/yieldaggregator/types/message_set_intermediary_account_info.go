package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetIntermediaryAccountInfo{}

func NewMsgSetIntermediaryAccountInfo(sender string, addrs []ChainAddress) *MsgSetIntermediaryAccountInfo {
	return &MsgSetIntermediaryAccountInfo{
		Sender: sender,
		Addrs:  addrs,
	}
}

func (msg MsgSetIntermediaryAccountInfo) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	return nil
}

func (msg MsgSetIntermediaryAccountInfo) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
