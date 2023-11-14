package types

import (
	// errorsmod "cosmossdk.io/errors"
	// "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDepositToTranche{}

func NewMsgDepositToTranche(sender string, strategyContract string, trancheType TrancheType, trancheMaturity uint64) *MsgDepositToTranche {
	return &MsgDepositToTranche{
		Sender:           sender,
		StrategyContract: strategyContract,
		TrancheType:      trancheType,
		TrancheMaturity:  trancheMaturity,
	}
}

func (msg MsgDepositToTranche) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if msg.StrategyContract == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty strategy contract")
	}

	return nil
}

func (msg MsgDepositToTranche) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
