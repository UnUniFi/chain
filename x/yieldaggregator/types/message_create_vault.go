package types

import (
	// errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateVault{}

func NewMsgCreateVault(sender string, denom string, commissionRate sdk.Dec, withdrawReserveRate sdk.Dec, strategyWeights []StrategyWeight, fee types.Coin, deposit types.Coin) *MsgCreateVault {
	return &MsgCreateVault{
		Sender:              sender,
		Denom:               denom,
		CommissionRate:      commissionRate,
		StrategyWeights:     strategyWeights,
		Fee:                 fee,
		Deposit:             deposit,
		WithdrawReserveRate: withdrawReserveRate,
	}
}

func (msg MsgCreateVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address: %s", err)
	}

	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return err
	}

	if msg.CommissionRate.IsNegative() || msg.CommissionRate.GTE(sdk.OneDec()) {
		return ErrInvalidCommissionRate
	}

	if msg.WithdrawReserveRate.LTE(sdk.ZeroDec()) || msg.WithdrawReserveRate.GTE(sdk.OneDec()) {
		return ErrInvalidWithdrawReserveRate
	}

	usedStrategy := make(map[uint64]bool)
	weightSum := sdk.ZeroDec()
	for _, strategyWeight := range msg.StrategyWeights {
		weightSum = weightSum.Add(strategyWeight.Weight)
		if usedStrategy[strategyWeight.StrategyId] {
			return ErrDuplicatedStrategy
		}
		usedStrategy[strategyWeight.StrategyId] = true
	}

	if !weightSum.Equal(sdk.OneDec()) {
		return ErrInvalidStrategyWeightSum
	}

	return nil
}

func (msg MsgCreateVault) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
