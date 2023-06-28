package types

import (
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgDepositToPool{}

func NewMsgDepositToPool(sender string, amount sdk.Coin) MsgDepositToPool {
	return MsgDepositToPool{
		Sender: sender,
		Amount: amount,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgDepositToPool) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgDepositToPool) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgWithdrawFromPool{}

func NewMsgWithdrawFromPool(sender string, lptAmount sdk.Int, denom string) MsgWithdrawFromPool {
	return MsgWithdrawFromPool{
		Sender:      sender,
		LptAmount:   lptAmount,
		RedeemDenom: denom,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgWithdrawFromPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgWithdrawFromPool) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgOpenPosition{}

func NewMsgOpenPosition(sender string, margin sdk.Coin, market Market, positionInstance codecTypes.Any) MsgOpenPosition {
	return MsgOpenPosition{
		Sender:           sender,
		Margin:           margin,
		Market:           market,
		PositionInstance: positionInstance,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgOpenPosition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgOpenPosition) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgClosePosition{}

func NewMsgClosePosition(sender string, positionId string) MsgClosePosition {
	return MsgClosePosition{
		Sender:     sender,
		PositionId: positionId,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClosePosition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClosePosition) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgReportLiquidation{}

func NewMsgReportLiquidation(sender string, positionId string, rewardRecipient string) MsgReportLiquidation {
	return MsgReportLiquidation{
		Sender:          sender,
		PositionId:      positionId,
		RewardRecipient: rewardRecipient,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgReportLiquidation) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgReportLiquidation) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgReportLevyPeriod{}

func NewMsgReportLevyPeriod() MsgReportLevyPeriod {
	return MsgReportLevyPeriod{}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgReportLevyPeriod) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgReportLevyPeriod) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

var _ sdk.Msg = &MsgAddMargin{}

func NewMsgAddMargin(sender string, positionId string, amount sdk.Coin) MsgAddMargin {
	return MsgAddMargin{
		Sender:     sender,
		PositionId: positionId,
		Amount:     amount,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgAddMargin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgAddMargin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgRemoveMargin{}

func NewMsgRemoveMargin(sender string, positionId string, amount sdk.Coin) MsgRemoveMargin {
	return MsgRemoveMargin{
		Sender:     sender,
		PositionId: positionId,
		Amount:     amount,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgRemoveMargin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgRemoveMargin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
