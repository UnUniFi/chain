package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/types"
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
	sender sdk.AccAddress,
	incentiveId string,
	subjectAccAddrs []sdk.AccAddress,
	weights []sdk.Dec,
) *MsgRegister {
	var subjectAddrs []types.StringAccAddress
	for _, accAddr := range subjectAccAddrs {
		subjectAddrs = append(subjectAddrs, accAddr.Bytes())
	}

	return &MsgRegister{
		IncentiveId:  incentiveId,
		SubjectAddrs: subjectAddrs,
		Weights:      weights,
	}
}

func (msg MsgRegister) Route() string { return RouterKey }

func (msg MsgRegister) Type() string { return TypeMsgRegister }

func (msg MsgRegister) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgRegister) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgRegister) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

func NewMsgWithdrawAllRewards(
	sender sdk.AccAddress,
) *MsgWithdrawAllRewards {
	return &MsgWithdrawAllRewards{
		Sender: sender.Bytes(),
	}
}

func (msg MsgWithdrawAllRewards) Route() string { return RouterKey }

func (msg MsgWithdrawAllRewards) Type() string { return TypeMsgWithdrawAllRewards }

func (msg MsgWithdrawAllRewards) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgWithdrawAllRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgWithdrawAllRewards) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

func NewMsgWithdrawReward(
	sender sdk.AccAddress,
	denom string,
) *MsgWithdrawReward {
	return &MsgWithdrawReward{
		Sender: sender.Bytes(),
		Denom:  denom,
	}
}

func (msg MsgWithdrawReward) Route() string { return RouterKey }

func (msg MsgWithdrawReward) Type() string { return TypeMsgWithdrawReward }

func (msg MsgWithdrawReward) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgWithdrawReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgWithdrawReward) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}
