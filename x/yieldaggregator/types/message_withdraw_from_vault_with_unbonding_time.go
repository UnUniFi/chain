package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWithdrawFromVaultWithUnbondingTime{}

func NewMsgWithdrawFromVaultWithUnbondingTime(sender string, vaultId uint64, lpTokenAmount sdk.Int) *MsgWithdrawFromVaultWithUnbondingTime {
	return &MsgWithdrawFromVaultWithUnbondingTime{
		Sender:        sender,
		VaultId:       vaultId,
		LpTokenAmount: lpTokenAmount,
	}
}

func (msg MsgWithdrawFromVaultWithUnbondingTime) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if !msg.LpTokenAmount.IsPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.LpTokenAmount.String())
	}

	return nil
}

func (msg MsgWithdrawFromVaultWithUnbondingTime) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
