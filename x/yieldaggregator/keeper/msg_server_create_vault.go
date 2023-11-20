package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) CreateVault(goCtx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params, err := k.Keeper.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	feeCollector, err := sdk.AccAddressFromBech32(params.FeeCollectorAddress)
	if err != nil {
		return nil, err
	}

	if msg.Fee.Denom != params.VaultCreationFee.Denom {
		return nil, types.ErrInvalidFeeDenom
	}
	if msg.Fee.IsLT(params.VaultCreationFee) {
		return nil, types.ErrInsufficientFee
	}

	if msg.Deposit.Denom != params.VaultCreationDeposit.Denom {
		return nil, types.ErrInvalidDepositDenom
	}
	if msg.Deposit.IsLT(params.VaultCreationDeposit) {
		return nil, types.ErrInsufficientDeposit
	}

	// transfer fee
	err = k.bankKeeper.SendCoins(ctx, sender, feeCollector, sdk.NewCoins(msg.Fee))
	if err != nil {
		return nil, err
	}

	// transfer deposit to module account (not for investment but for incentivization of deleting unused vault)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(msg.Deposit))
	if err != nil {
		return nil, err
	}

	for _, strategyWeight := range msg.StrategyWeights {
		val, found := k.Keeper.GetStrategy(ctx, strategyWeight.Denom, strategyWeight.StrategyId)
		if !found {
			return nil, types.ErrInvalidStrategyInvolved
		}
		denomInfo := k.GetDenomInfo(ctx, val.Denom)
		if denomInfo.Symbol != msg.Symbol {
			return nil, types.ErrDenomDoesNotMatchVaultSymbol
		}
	}

	vault := types.Vault{
		Symbol:                 msg.Symbol,
		Name:                   msg.Name,
		Description:            msg.Description,
		Owner:                  msg.Sender,
		OwnerDeposit:           msg.Deposit,
		WithdrawCommissionRate: msg.CommissionRate,
		WithdrawReserveRate:    msg.WithdrawReserveRate,
		StrategyWeights:        msg.StrategyWeights,
		FeeCollectorAddress:    msg.FeeCollectorAddress,
	}
	id := k.Keeper.AppendVault(ctx, vault)

	return &types.MsgCreateVaultResponse{
		Id: id,
	}, nil
}
