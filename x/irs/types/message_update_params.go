package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgUpdateParams{}

func NewMsgUpdateParams(sender string, params Params) *MsgUpdateParams {
	return &MsgUpdateParams{
		Sender: sender,
		Params: params,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgUpdateParams) ValidateBasic() error {
	if err := msg.Params.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
