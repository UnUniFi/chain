package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgDeposit{}

func NewMsgDeposit(sender sdk.AccAddress, amounts sdk.Coins, executeOrders bool) MsgDeposit {
	return MsgDeposit{
		FromAddress:   sender.Bytes(),
		Amount:        amounts,
		ExecuteOrders: executeOrders,
	}
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

func NewMsgWithdraw(sender sdk.AccAddress, amounts sdk.Coins) MsgWithdraw {
	return MsgWithdraw{
		FromAddress: sender.Bytes(),
		Amount:      amounts,
	}
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

func NewMsgAddFarmingOrder(sender sdk.AccAddress, order FarmingOrder) MsgAddFarmingOrder {
	return MsgAddFarmingOrder{
		FromAddress: sender.Bytes(),
		Order:       &order,
	}
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

func NewMsgDeleteFarmingOrder(sender sdk.AccAddress, orderId string) MsgDeleteFarmingOrder {
	return MsgDeleteFarmingOrder{
		FromAddress: sender.Bytes(),
		OrderId:     orderId,
	}
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

func NewMsgActivateFarmingOrder(sender sdk.AccAddress, orderId string) MsgActivateFarmingOrder {
	return MsgActivateFarmingOrder{
		FromAddress: sender.Bytes(),
		OrderId:     orderId,
	}
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

func NewMsgInactivateFarmingOrder(sender sdk.AccAddress, orderId string) MsgInactivateFarmingOrder {
	return MsgInactivateFarmingOrder{
		FromAddress: sender.Bytes(),
		OrderId:     orderId,
	}
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

func NewMsgExecuteFarmingOrders(sender sdk.AccAddress, orderIds []string) MsgExecuteFarmingOrders {
	return MsgExecuteFarmingOrders{
		FromAddress: sender.Bytes(),
		OrderIds:    orderIds,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgExecuteFarmingOrders) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgExecuteFarmingOrders) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress.AccAddress()}
}
