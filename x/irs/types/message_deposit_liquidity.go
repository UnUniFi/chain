package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDepositLiquidity{}

func NewMsgDepositLiquidity(sender string, trancheId uint64, shareOutAmount math.Int, tokenInMaxs sdk.Coins) *MsgDepositLiquidity {
	return &MsgDepositLiquidity{
		Sender:         sender,
		TrancheId:      trancheId,
		ShareOutAmount: shareOutAmount,
		TokenInMaxs:    tokenInMaxs,
	}
}

func (msg MsgDepositLiquidity) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if msg.TrancheId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid tranche id")
	}

	if !msg.ShareOutAmount.IsPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.ShareOutAmount.String())
	}

	return nil
}

func (msg MsgDepositLiquidity) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
