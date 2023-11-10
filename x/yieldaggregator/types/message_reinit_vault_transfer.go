package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgReinitVaultTransfer{}

func NewMsgReinitVaultTransfer(sender string, vaultId uint64, strategyDenom string, strategyId uint64, amount sdk.Coin) *MsgReinitVaultTransfer {
	return &MsgReinitVaultTransfer{
		Sender:        sender,
		VaultId:       vaultId,
		StrategyDenom: strategyDenom,
		StrategyId:    strategyId,
		Amount:        amount,
	}
}

func (msg MsgReinitVaultTransfer) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	return nil
}

func (msg MsgReinitVaultTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
