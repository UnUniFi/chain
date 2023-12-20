package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateSwapMaturedYtToUt(c context.Context, req *types.QueryEstimateSwapMaturedYtToUtRequest) (*types.QueryEstimateSwapMaturedYtToUtResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.PoolId)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	if uint64(ctx.BlockTime().Unix()) <= tranche.StartTime+tranche.Maturity {
		return nil, types.ErrTrancheNotMatured
	}
	ytDenom := types.YtDenom(tranche)
	ytAmount, ok := sdk.NewIntFromString(req.YtAmount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	redeemAmount, err := k.CalculateRedeemYtAmount(ctx, tranche, sdk.NewCoin(ytDenom, ytAmount))
	if err != nil {
		return nil, err
	}
	depositInfo := k.GetStrategyDepositInfo(ctx, tranche.StrategyContract)
	return &types.QueryEstimateSwapMaturedYtToUtResponse{
		UtAmount: sdk.NewCoin(depositInfo.Denom, redeemAmount),
	}, nil
}
