package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgRegisterStrategy{}

func NewMsgRegisterStrategy(sender string, denom, contractAddr, name, gitUrl string) MsgRegisterStrategy {
	return MsgRegisterStrategy{
		Sender:          sender,
		Denom:           denom,
		ContractAddress: contractAddr,
		Name:            name,
		GitUrl:          gitUrl,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgRegisterStrategy) ValidateBasic() error {
	if msg.Denom == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty strategy denom")
	}
	if msg.Name == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty strategy name")
	}
	if msg.ContractAddress == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty strategy contract address")
	}
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgRegisterStrategy) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
