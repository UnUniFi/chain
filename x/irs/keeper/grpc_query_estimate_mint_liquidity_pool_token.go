package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateMintLiquidityPoolToken(c context.Context, req *types.QueryEstimateMintLiquidityPoolTokenRequest) (*types.QueryEstimateMintLiquidityPoolTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	// initial deposit
	if tranche.TotalShares.IsZero() {
		return &types.QueryEstimateMintLiquidityPoolTokenResponse{
			RequiredAmount: sdk.Coins{},
		}, nil
	}
	desiredAmount, ok := sdk.NewIntFromString(req.DesiredAmount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	// we do an abstract calculation on the lp liquidity coins needed to have
	// the designated amount of given shares of the pool without performing swap
	neededLpLiquidity, err := GetMaximalNoSwapLPAmount(ctx, tranche, desiredAmount)
	if err != nil {
		return nil, err
	}
	return &types.QueryEstimateMintLiquidityPoolTokenResponse{
		RequiredAmount: neededLpLiquidity,
	}, nil
}
