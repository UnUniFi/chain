package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) TranchePtAPYs(c context.Context, req *types.QueryTranchePtAPYsRequest) (*types.QueryTranchePtAPYsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	depositInfo := k.GetStrategyDepositInfo(ctx, tranche.StrategyContract)
	swapCoin := sdk.NewCoin(depositInfo.Denom, sdk.NewInt(1_000_000))
	pt, err := k.SimulateSwapPoolTokens(ctx, tranche, swapCoin)
	if err != nil {
		return nil, err
	}
	restMaturity := tranche.StartTime + tranche.Maturity - uint64(ctx.BlockTime().Unix())
	maturityPerYear := sdk.NewDecFromInt(sdk.NewIntFromUint64(restMaturity)).QuoInt(sdk.NewInt(365 * 24 * 60 * 60))
	ptAPY := sdk.NewDecFromInt(pt.Amount.Sub(swapCoin.Amount)).QuoInt(swapCoin.Amount).Mul(maturityPerYear)
	ptRate := sdk.NewDecFromInt(pt.Amount).QuoInt(swapCoin.Amount)

	return &types.QueryTranchePtAPYsResponse{
		PtApy:       ptAPY,
		PtRatePerUt: ptRate,
	}, nil
}

func (k Keeper) TrancheYtAPYs(c context.Context, req *types.QueryTrancheYtAPYsRequest) (*types.QueryTrancheYtAPYsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	restMaturity := tranche.StartTime + tranche.Maturity - uint64(ctx.BlockTime().Unix())
	maturityPerYear := sdk.NewDecFromInt(sdk.NewIntFromUint64(restMaturity)).QuoInt(sdk.NewInt(365 * 24 * 60 * 60))

	ytDenom := types.YtDenom(tranche)
	yt := sdk.NewCoin(ytDenom, sdk.NewInt(1_000_000))
	requiredUt, err := k.CalculateRequiredUtSwapToYt(ctx, tranche, yt.Amount)
	if err != nil {
		return nil, err
	}
	redeemUt, err := k.CalculateRedeemYtAmount(ctx, tranche, yt)
	if err != nil {
		return nil, err
	}
	ytAPY := sdk.NewDecFromInt(redeemUt.Sub(requiredUt.Amount)).QuoInt(requiredUt.Amount).Mul(maturityPerYear)
	ytRate := sdk.NewDecFromInt(yt.Amount).QuoInt(requiredUt.Amount)

	return &types.QueryTrancheYtAPYsResponse{
		YtApy:       ytAPY,
		YtRatePerUt: ytRate,
	}, nil
}

func (k Keeper) TranchePoolAPYs(c context.Context, req *types.QueryTranchePoolAPYsRequest) (*types.QueryTranchePoolAPYsResponse, error) {
	// TODO: implement
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	depositInfo := k.GetStrategyDepositInfo(ctx, tranche.StrategyContract)
	swapCoin := sdk.NewCoin(depositInfo.Denom, sdk.NewInt(1_000_000))
	pt, err := k.SimulateSwapPoolTokens(ctx, tranche, swapCoin)
	if err != nil {
		return nil, err
	}
	restMaturity := tranche.StartTime + tranche.Maturity - uint64(ctx.BlockTime().Unix())
	maturityPerYear := sdk.NewDecFromInt(sdk.NewIntFromUint64(restMaturity)).QuoInt(sdk.NewInt(365 * 24 * 60 * 60))
	ptAPY := sdk.NewDecFromInt(pt.Amount.Sub(swapCoin.Amount)).QuoInt(swapCoin.Amount).Mul(maturityPerYear)
	ptRate := sdk.NewDecFromInt(pt.Amount).QuoInt(swapCoin.Amount)

	ptDenom := types.PtDenom(tranche)
	var utPercentage sdk.Dec
	var ptPercentage sdk.Dec
	if ptDenom == tranche.PoolAssets[0].Denom {
		total := sdk.NewDecFromInt(tranche.PoolAssets[1].Amount).Mul(ptRate).Add(sdk.NewDecFromInt(tranche.PoolAssets[0].Amount))
		utPercentage = sdk.NewDecFromInt(tranche.PoolAssets[1].Amount).Quo(total)
		ptPercentage = sdk.NewDecFromInt(tranche.PoolAssets[0].Amount).Quo(total)
	} else {
		total := sdk.NewDecFromInt(tranche.PoolAssets[0].Amount).Mul(ptRate).Add(sdk.NewDecFromInt(tranche.PoolAssets[1].Amount))
		utPercentage = sdk.NewDecFromInt(tranche.PoolAssets[0].Amount).Quo(total)
		ptPercentage = sdk.NewDecFromInt(tranche.PoolAssets[1].Amount).Quo(total)
	}
	discountPtAPY := ptAPY.Mul(ptPercentage)

	return &types.QueryTranchePoolAPYsResponse{
		LiquidityApy:       sdk.ZeroDec(),
		DiscountPtApy:      discountPtAPY,
		LiquidityRatePerUt: sdk.ZeroDec(),
		UtPercentageInPool: utPercentage,
		PtPercentageInPool: ptPercentage,
	}, nil
}
