package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgClaimJPYXMintingReward{}
var _ sdk.Msg = &MsgClaimHardLiquidityProviderReward{}

// MsgClaimJPYXMintingReward message type used to claim JPYX minting rewards
type MsgClaimJPYXMintingReward struct {
	Sender         sdk.AccAddress `json:"sender" yaml:"sender"`
	MultiplierName string         `json:"multiplier_name" yaml:"multiplier_name"`
}

// NewMsgClaimJPYXMintingReward returns a new MsgClaimJPYXMintingReward.
func NewMsgClaimJPYXMintingReward(sender sdk.AccAddress, multiplierName string) MsgClaimJPYXMintingReward {
	return MsgClaimJPYXMintingReward{
		Sender:         sender,
		MultiplierName: multiplierName,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimJPYXMintingReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimJPYXMintingReward) Type() string { return "claim_jpyx_minting_reward" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimJPYXMintingReward) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	return MultiplierName(strings.ToLower(msg.MultiplierName)).IsValid()
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimJPYXMintingReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimJPYXMintingReward) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgClaimHardLiquidityProviderReward message type used to claim Hard liquidity provider rewards
type MsgClaimHardLiquidityProviderReward struct {
	Sender         sdk.AccAddress `json:"sender" yaml:"sender"`
	MultiplierName string         `json:"multiplier_name" yaml:"multiplier_name"`
}

// NewMsgClaimHardLiquidityProviderReward returns a new MsgClaimHardLiquidityProviderReward.
func NewMsgClaimHardLiquidityProviderReward(sender sdk.AccAddress, multiplierName string) MsgClaimHardLiquidityProviderReward {
	return MsgClaimHardLiquidityProviderReward{
		Sender:         sender,
		MultiplierName: multiplierName,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimHardLiquidityProviderReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimHardLiquidityProviderReward) Type() string {
	return "claim_hard_liquidity_provider_reward"
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimHardLiquidityProviderReward) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	return MultiplierName(strings.ToLower(msg.MultiplierName)).IsValid()
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimHardLiquidityProviderReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimHardLiquidityProviderReward) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
