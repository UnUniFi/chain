package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateVault(ctx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	params := k.Keeper.GetParams(sdkCtx)

	if msg.Fee.Denom != params.VaultCreationFee.Denom {
		// return nil, types.ErrInvalidFeeDenom
	}
	if msg.Fee.IsLT(params.VaultCreationFee) {
		// return nil, types.ErrInsufficientFee
	}

	if msg.Deposit.Denom != params.VaultCreationDeposit.Denom {
		// return nil, types.ErrInvalidDepositDenom
	}
	if msg.Deposit.IsLT(params.VaultCreationDeposit) {
		// return nil, types.ErrInsufficientDeposit
	}

	// TODO: transfer fee
	// TODO: transfer deposit

	vault := types.Vault{
		Denom:                  msg.Denom,
		Owner:                  msg.Sender,
		OwnerDeposit:           msg.Deposit,
		WithdrawCommissionRate: msg.CommissionRate,
		StrategyWeights:        msg.StrategyWeights,
	}

	k.Keeper.AppendVault(sdkCtx, vault)
	panic("implement me")
}
