package types

import (
	// errorsmod "cosmossdk.io/errors"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterInterestRateSwapVault{}

func NewMsgRegisterInterestRateSwapVault(sender string, name, description string) *MsgRegisterInterestRateSwapVault {
	return &MsgRegisterInterestRateSwapVault{
		Sender:      sender,
		Name:        name,
		Description: description,
	}
}

func (msg MsgRegisterInterestRateSwapVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if msg.StrategyContract == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty strategy contract")
	}

	if msg.Name == "" {
		return ErrInvalidVaultName
	}

	if msg.Description == "" {
		return ErrInvalidVaultDescription
	}

	return nil
}

func (msg MsgRegisterInterestRateSwapVault) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
