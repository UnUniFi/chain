package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgMintLiquidityProviderToken{}

func NewMsgMintLiquidityProviderToken(sender sdk.AccAddress, amount sdk.Coin) MsgMintLiquidityProviderToken {
	return MsgMintLiquidityProviderToken{
		Sender: sender.Bytes(),
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
