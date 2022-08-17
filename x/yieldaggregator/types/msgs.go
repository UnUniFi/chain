package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgDeposit{}

func NewMsgDeposit(sender sdk.AccAddress) MsgDeposit {
	return MsgDeposit{}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgDeposit) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgWithdraw{}

func NewMsgWithdraw(sender sdk.AccAddress) MsgWithdraw {
	return MsgWithdraw{}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgWithdraw) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgAddFarmingOrder{}

func NewMsgAddFarmingOrder(sender sdk.AccAddress) MsgAddFarmingOrder {
	return MsgAddFarmingOrder{}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgAddFarmingOrder) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgAddFarmingOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgDeleteFarmingOrder{}

func NewMsgDeleteFarmingOrder(sender sdk.AccAddress) MsgDeleteFarmingOrder {
	return MsgDeleteFarmingOrder{}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgDeleteFarmingOrder) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgDeleteFarmingOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgActivateFarmingOrder{}

func NewMsgActivateFarmingOrder(sender sdk.AccAddress) MsgActivateFarmingOrder {
	return MsgActivateFarmingOrder{}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgActivateFarmingOrder) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgActivateFarmingOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgInactivateFarmingOrder{}

func NewMsgInactivateFarmingOrder(sender sdk.AccAddress) MsgInactivateFarmingOrder {
	return MsgInactivateFarmingOrder{}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgInactivateFarmingOrder) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgInactivateFarmingOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgExecuteFarmingOrders{}

func NewMsgExecuteFarmingOrders(sender sdk.AccAddress) MsgExecuteFarmingOrders {
	return MsgExecuteFarmingOrders{}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgExecuteFarmingOrders) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgExecuteFarmingOrders) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress.AccAddress()}
}
