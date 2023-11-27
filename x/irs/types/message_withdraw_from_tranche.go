package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWithdrawFromTranche{}

func NewMsgWithdrawFromTranche(sender string, trancheId uint64, trancheType TrancheType, token sdk.Coin) *MsgWithdrawFromTranche {
	return &MsgWithdrawFromTranche{
		Sender:      sender,
		TrancheId:   trancheId,
		TrancheType: trancheType,
		Token:       token,
	}
}

func (msg MsgWithdrawFromTranche) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if msg.TrancheId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid tranche id")
	}

	if !msg.Token.IsPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Token.String())
	}

	return nil
}

func (msg MsgWithdrawFromTranche) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
