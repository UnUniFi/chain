package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWithdrawFromTranche{}

func NewMsgWithdrawFromTranche(sender string, strategyContract string, principalToken sdk.Coin) *MsgWithdrawFromTranche {
	return &MsgWithdrawFromTranche{
		Sender:           sender,
		StrategyContract: strategyContract,
		PrincipalToken:   principalToken,
	}
}

func (msg MsgWithdrawFromTranche) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if msg.StrategyContract == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "empty strategy contract")
	}

	if !msg.PrincipalToken.IsPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.PrincipalToken.String())
	}

	return nil
}

func (msg MsgWithdrawFromTranche) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
