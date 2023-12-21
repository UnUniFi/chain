package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateRedeemLiquidityPoolToken(c context.Context, req *types.QueryEstimateRedeemLiquidityPoolTokenRequest) (*types.QueryEstimateRedeemLiquidityPoolTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	totalSharesAmount := tranche.GetTotalShares()
	shareInAmount, ok := sdk.NewIntFromString(req.Amount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	if shareInAmount.GT(totalSharesAmount.Amount) {
		return nil, types.ErrInvalidTotalShares
	} else if shareInAmount.LTE(sdk.ZeroInt()) {
		return nil, types.ErrInvalidTotalShares
	}
	exitFee := tranche.ExitFee
	exitCoins, err := tranche.ExitPool(ctx, shareInAmount, exitFee)
	if err != nil {
		return nil, err
	}
	return &types.QueryEstimateRedeemLiquidityPoolTokenResponse{
		RedeemAmount: exitCoins,
	}, nil
}
