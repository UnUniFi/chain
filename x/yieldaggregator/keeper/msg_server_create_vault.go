package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) CreateVault(goCtx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.Keeper.GetParams(ctx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
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
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(msg.Fee))
	if err != nil {
		return nil, err
	}

	// transfer deposit to module account (not for investment but for incentivization of deleting unused vault)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(msg.Deposit))
	if err != nil {
		return nil, err
	}

	for _, strategyWeight := range msg.StrategyWeights {
		_, found := k.Keeper.GetStrategy(ctx, msg.Denom, strategyWeight.StrategyId)
		if !found {
			return nil, types.ErrInvalidStrategyInvolved
		}
	}

	vault := types.Vault{
		Denom:                  msg.Denom,
		Owner:                  msg.Sender,
		OwnerDeposit:           msg.Deposit,
		WithdrawCommissionRate: msg.CommissionRate,
		WithdrawReserveRate:    msg.WithdrawReserveRate,
		StrategyWeights:        msg.StrategyWeights,
	}
	id := k.Keeper.AppendVault(ctx, vault)

	return &types.MsgCreateVaultResponse{
		Id: id,
	}, nil
}
