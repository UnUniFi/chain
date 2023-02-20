package types

import (
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifiTypes "github.com/UnUniFi/chain/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgMintLiquidityProviderToken{}

func NewMsgMintLiquidityProviderToken(sender ununifiTypes.StringAccAddress, amount sdk.Coin) MsgMintLiquidityProviderToken {
	return MsgMintLiquidityProviderToken{
		Sender: sender,
		Amount: amount,
	}
}

// Route return the message type used for routing the message.
func (msg MsgMintLiquidityProviderToken) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgMintLiquidityProviderToken) Type() string { return "mint_lpt" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgMintLiquidityProviderToken) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgMintLiquidityProviderToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgMintLiquidityProviderToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgBurnLiquidityProviderToken{}

func NewMsgBurnLiquidityProviderToken(sender sdk.AccAddress, amount sdk.Int, denom string) MsgBurnLiquidityProviderToken {
	return MsgBurnLiquidityProviderToken{
		Sender:      sender.Bytes(),
		Amount:      amount,
		RedeemDenom: denom,
	}
}

// Route return the message type used for routing the message.
func (msg MsgBurnLiquidityProviderToken) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgBurnLiquidityProviderToken) Type() string { return "burn_lpt" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgBurnLiquidityProviderToken) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgBurnLiquidityProviderToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgBurnLiquidityProviderToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgOpenPosition{}

func NewMsgOpenPosition(sender sdk.AccAddress, margin sdk.Coin, market Market, positionInstance codecTypes.Any) MsgOpenPosition {
	return MsgOpenPosition{
		Sender:           sender.Bytes(),
		Margin:           margin,
		Market:           market,
		PositionInstance: positionInstance,
	}
}

// Route return the message type used for routing the message.
func (msg MsgOpenPosition) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgOpenPosition) Type() string { return "open_position" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgOpenPosition) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgOpenPosition) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgOpenPosition) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgClosePosition{}

func NewMsgClosePosition(sender ununifiTypes.StringAccAddress, positionId string) MsgClosePosition {
	return MsgClosePosition{
		Sender:     sender,
		PositionId: positionId,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClosePosition) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClosePosition) Type() string { return "close_position" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClosePosition) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClosePosition) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClosePosition) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgReportLiquidation{}

func NewMsgReportLiquidation(sender ununifiTypes.StringAccAddress, positionId string, rewardRecipient ununifiTypes.StringAccAddress) MsgReportLiquidation {
	return MsgReportLiquidation{
		Sender:          sender,
		PositionId:      positionId,
		RewardRecipient: rewardRecipient,
	}
}

// Route return the message type used for routing the message.
func (msg MsgReportLiquidation) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgReportLiquidation) Type() string {
	return "report_liquidation"
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgReportLiquidation) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgReportLiquidation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgReportLiquidation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

var _ sdk.Msg = &MsgReportLevyPeriod{}

func NewMsgReportLevyPeriod() MsgReportLevyPeriod {
	return MsgReportLevyPeriod{}
}

// Route return the message type used for routing the message.
func (msg MsgReportLevyPeriod) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgReportLevyPeriod) Type() string {
	return "report_levy_period"
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgReportLevyPeriod) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgReportLevyPeriod) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgReportLevyPeriod) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
