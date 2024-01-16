package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWithdrawLiquidity{}

func NewMsgWithdrawLiquidity(sender string, trancheId uint64, shareAmount math.Int, tokenOutMins sdk.Coins) *MsgWithdrawLiquidity {
	return &MsgWithdrawLiquidity{
		Sender:       sender,
		TrancheId:    trancheId,
		ShareAmount:  shareAmount,
		TokenOutMins: tokenOutMins,
	}
}

func (msg MsgWithdrawLiquidity) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if msg.TrancheId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty strategy contract")
	}

	if !msg.ShareAmount.IsPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.ShareAmount.String())
	}

	return nil
}

func (msg MsgWithdrawLiquidity) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
