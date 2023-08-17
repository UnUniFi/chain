package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	osmomath "github.com/UnUniFi/chain/osmomath"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) EstimateRedeemAmount(c context.Context, req *types.QueryEstimateRedeemAmountRequest) (*types.QueryEstimateRedeemAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	params, err := k.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	vault, found := k.GetVault(ctx, req.Id)
	if !found {
		return nil, types.ErrInvalidVaultId
	}
	burnAmount, ok := sdk.NewIntFromString(req.BurnAmount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	principal := k.EstimateRedeemAmountInternal(ctx, vault.Denom, vault.Id, burnAmount)

	// Unstake funds from Strategy
	amountToUnbond := principal.Amount

	// implement fees on withdrawal
	amountInVault := k.VaultWithdrawalAmount(ctx, vault)
	amountUnbonding := k.VaultUnbondingAmountInStrategies(ctx, vault)

	reserveMaintenanceRate := sdk.ZeroDec()
	if amountInVault.GT(amountToUnbond) {
		reserveMaintenanceRate = sdk.NewDecFromInt(amountInVault.Sub(amountToUnbond)).Quo(sdk.NewDecFromInt(amountInVault.Add(amountUnbonding)))
	}

	e := osmomath.NewDecWithPrec(2718281, 6) // 2.718281
	withdrawFeeRate := osmomath.OneDec().
		Quo(e.Power(osmomath.BigDecFromSDKDec(reserveMaintenanceRate).MulInt64(10))).
		SDKDec()

	withdrawFee := sdk.NewDecFromInt(principal.Amount).Mul(withdrawFeeRate).RoundInt()
	withdrawAmount := principal.Amount.Sub(withdrawFee)

	withdrawModuleCommissionFee := sdk.NewDecFromInt(withdrawAmount).Mul(params.CommissionRate).RoundInt()
	withdrawVaultCommissionFee := sdk.NewDecFromInt(withdrawAmount).Mul(vault.WithdrawCommissionRate).RoundInt()
	withdrawAmountWithoutCommission := withdrawAmount.Sub(withdrawModuleCommissionFee).Sub(withdrawVaultCommissionFee)
	fee := withdrawModuleCommissionFee.Add(withdrawVaultCommissionFee)

	return &types.QueryEstimateRedeemAmountResponse{
		ShareAmount:  sdk.NewCoin(types.GetLPTokenDenom(vault.Id), burnAmount),
		Fee:          sdk.NewCoin(principal.Denom, fee),
		RedeemAmount: sdk.NewCoin(principal.Denom, withdrawAmountWithoutCommission),
	}, nil
}
