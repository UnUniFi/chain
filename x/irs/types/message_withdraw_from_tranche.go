package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWithdrawFromTranche{}

func NewMsgWithdrawFromTranche(sender string, trancheId uint64, trancheType TrancheType, tokens sdk.Coins, requiredUt sdk.Int) *MsgWithdrawFromTranche {
	return &MsgWithdrawFromTranche{
		Sender:      sender,
		TrancheId:   trancheId,
		TrancheType: trancheType,
		Tokens:      tokens,
		RequiredUt:  requiredUt,
	}
}

func (msg MsgWithdrawFromTranche) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if msg.TrancheId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid tranche id")
	}

	if sdk.Coins(msg.Tokens).Empty() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, sdk.Coins(msg.Tokens).String())
	}

	return nil
}

func (msg MsgWithdrawFromTranche) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
