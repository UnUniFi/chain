package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRegister           = "register"
	TypeMsgWithdrawAllRewards = "withdraw-all-rewards"
	TypeMsgWithdrawReward     = "withdraw-reward"
)

var (
	_ sdk.Msg = &MsgRegister{}
	_ sdk.Msg = &MsgWithdrawAllRewards{}
	_ sdk.Msg = &MsgWithdrawReward{}
)

func NewMsgRegister(
	sender string,
	recipientContainerId string,
	subjectAddrs []string,
	weights []sdk.Dec,
) *MsgRegister {
	return &MsgRegister{
		Sender:               sender,
		RecipientContainerId: recipientContainerId,
		Addresses:            subjectAddrs,
		Weights:              weights,
	}
}

func (msg MsgRegister) Route() string { return RouterKey }

func (msg MsgRegister) Type() string { return TypeMsgRegister }

func (msg MsgRegister) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	for _, addr := range msg.Addresses {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return err
		}
	}

	// return err if the number of elements in subjects and weights aren't same
	if len(msg.Addresses) != len(msg.Weights) {
		return sdkerrors.Wrapf(ErrSubjectsWeightsNumUnmatched, "subjects element num: %d, weights element num: %d", len(msg.Addresses), len(msg.Weights))
	}

	// the summed number of all weights must be 1
	totalWeight := sdk.ZeroDec()
	for _, weight := range msg.Weights {
		totalWeight = totalWeight.Add(weight)
	}
	if !(totalWeight.Equal(sdk.OneDec())) {
		return sdkerrors.Wrap(ErrInvalidTotalWeight, totalWeight.String())
	}

	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgRegister) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgRegister) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgWithdrawAllRewards(
	sender string,
) *MsgWithdrawAllRewards {
	return &MsgWithdrawAllRewards{
		Sender: sender,
	}
}

func (msg MsgWithdrawAllRewards) Route() string { return RouterKey }

func (msg MsgWithdrawAllRewards) Type() string { return TypeMsgWithdrawAllRewards }

func (msg MsgWithdrawAllRewards) ValidateBasic() error {
	// check if addresses are valid
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		panic(err)
	}

	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgWithdrawAllRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
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

func (msg MsgWithdrawReward) Route() string { return RouterKey }

func (msg MsgWithdrawReward) Type() string { return TypeMsgWithdrawReward }

func (msg MsgWithdrawReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgWithdrawReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgWithdrawReward) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
