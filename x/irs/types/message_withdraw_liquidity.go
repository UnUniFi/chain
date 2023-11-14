package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWithdrawLiquidity{}

func NewMsgWithdrawLiquidity(sender string, strategyContract string, shareAmount sdk.Int) *MsgWithdrawLiquidity {
	return &MsgWithdrawLiquidity{
		Sender:           sender,
		StrategyContract: strategyContract,
		ShareAmount:      shareAmount,
	}
}

func (msg MsgWithdrawLiquidity) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if msg.StrategyContract == "" {
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
