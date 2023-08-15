package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgWithdrawAllRewards{}
	_ sdk.Msg = &MsgWithdrawReward{}
)

func NewMsgWithdrawAllRewards(
	sender string,
) *MsgWithdrawAllRewards {
	return &MsgWithdrawAllRewards{
		Sender: sender,
	}
}

func (msg MsgWithdrawAllRewards) ValidateBasic() error {
	// check if addresses are valid
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		panic(err)
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgWithdrawAllRewards) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)

	return []sdk.AccAddress{addr}
}

func NewMsgWithdrawReward(
	sender string,
	denom string,
) *MsgWithdrawReward {
	return &MsgWithdrawReward{
		Sender: sender,
		Denom:  denom,
	}
}

func (msg MsgWithdrawReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgWithdrawReward) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
