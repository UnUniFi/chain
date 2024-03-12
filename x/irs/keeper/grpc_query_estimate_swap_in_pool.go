package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateSwapInPool(c context.Context, req *types.QueryEstimateSwapInPoolRequest) (*types.QueryEstimateSwapInPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	depositAmount, ok := sdk.NewIntFromString(req.Amount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	ptDenom := types.PtDenom(tranche)
	if req.Denom == ptDenom {
		if tranche.StartTime+tranche.Maturity <= uint64(ctx.BlockTime().Unix()) {
			info := k.GetStrategyDepositInfo(ctx, tranche.StrategyContract)
			rate := sdk.MustNewDecFromStr(info.DepositDenomRate)
			if rate.IsZero() {
				return &types.QueryEstimateSwapInPoolResponse{
					Amount: sdk.NewCoin(tranche.DepositDenom, sdk.ZeroInt()),
				}, nil
			}
			amount := sdk.NewDecFromInt(depositAmount).Quo(rate).TruncateInt()
			return &types.QueryEstimateSwapInPoolResponse{
				Amount: sdk.NewCoin(tranche.DepositDenom, amount),
			}, nil
		}
	}

	amount, err := k.SimulateSwapPoolTokens(ctx, tranche, sdk.NewCoin(req.Denom, depositAmount))
	if err != nil {
		return nil, err
	}

	return &types.QueryEstimateSwapInPoolResponse{
		Amount: amount,
	}, nil
}
