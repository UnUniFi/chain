package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDepositToVault = "deposit-to-vault"

var _ sdk.Msg = &MsgDepositToVault{}

func NewMsgDepositToVault(sender string, amount sdk.Coin) *MsgDepositToVault {
	return &MsgDepositToVault{
		Sender: sender,
		Amount: amount,
	}
}

func (msg MsgDepositToVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if !msg.Amount.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	if !msg.Amount.IsPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

func (msg *MsgDepositToVault) Route() string {
	return RouterKey
}

func (msg *MsgDepositToVault) Type() string {
	return TypeMsgDepositToVault
}

func (msg MsgDepositToVault) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDepositToVault) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
