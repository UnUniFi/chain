package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawFromVault = "withdraw-from-vault"

var _ sdk.Msg = &MsgWithdrawFromVault{}

func NewMsgWithdrawFromVault(sender string, vaultDenom string, lpTokenAmount sdk.Int) *MsgWithdrawFromVault {
	return &MsgWithdrawFromVault{
		Sender:        sender,
		VaultDenom:    vaultDenom,
		LpTokenAmount: lpTokenAmount,
	}
}

func (msg MsgWithdrawFromVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if err := sdk.ValidateDenom(msg.VaultDenom); err != nil {
		return sdkerrors.ErrInvalidCoins.Wrapf("invalid vault denom: %s", err)
	}

	if !msg.LpTokenAmount.IsPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.LpTokenAmount.String())
	}

	return nil
}

func (msg MsgWithdrawFromVault) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgWithdrawFromVault) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
