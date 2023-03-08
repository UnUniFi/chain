package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateVault(ctx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	params := k.Keeper.GetParams(sdkCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		// TODO
		return nil, err
	}

	if msg.Fee.Denom != params.VaultCreationFee.Denom {
		// return nil, types.ErrInvalidFeeDenom
		return nil, nil
	}
	if msg.Fee.IsLT(params.VaultCreationFee) {
		// return nil, types.ErrInsufficientFee
		return nil, nil
	}

	if msg.Deposit.Denom != params.VaultCreationDeposit.Denom {
		// return nil, types.ErrInvalidDepositDenom
		return nil, nil
	}
	if msg.Deposit.IsLT(params.VaultCreationDeposit) {
		// return nil, types.ErrInsufficientDeposit
		return nil, nil
	}

	// transfer fee
	err = k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, sender, types.ModuleName, sdk.NewCoins(msg.Fee))
	if err != nil {
		return nil, err
	}
	// transfer deposit
	err = k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, sender, types.ModuleName, sdk.NewCoins(msg.Deposit))
	if err != nil {
		return nil, err
	}

	vault := types.Vault{
		Denom:                  msg.Denom,
		Owner:                  msg.Sender,
		OwnerDeposit:           msg.Deposit,
		WithdrawCommissionRate: msg.CommissionRate,
		StrategyWeights:        msg.StrategyWeights,
	}

	k.Keeper.AppendVault(sdkCtx, vault)

	return &types.MsgCreateVaultResponse{}, nil
}
