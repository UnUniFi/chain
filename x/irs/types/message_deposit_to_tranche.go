package types

import (
	// errorsmod "cosmossdk.io/errors"
	// "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDepositToTranche{}

func NewMsgDepositToTranche(sender string, trancheId uint64, trancheType TrancheType, token sdk.Coin, requiredYt sdk.Coin) *MsgDepositToTranche {
	return &MsgDepositToTranche{
		Sender:      sender,
		TrancheId:   trancheId,
		TrancheType: trancheType,
		Token:       token,
		RequiredYt:  requiredYt,
	}
}

func (msg MsgDepositToTranche) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if msg.TrancheId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty tranche id")
	}

	return nil
}

func (msg MsgDepositToTranche) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
