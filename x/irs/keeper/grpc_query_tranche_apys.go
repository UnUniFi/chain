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
	swapCoin := sdk.NewCoin(tranche.DepositDenom, sdk.NewInt(1_000_000))
	pt, err := k.SimulateSwapPoolTokens(ctx, tranche, swapCoin)
	if err != nil {
		return nil, err
	}
	restMaturity := tranche.StartTime + tranche.Maturity - uint64(ctx.BlockTime().Unix())
	maturityPerYear := sdk.NewDecFromInt(sdk.NewIntFromUint64(restMaturity)).QuoInt(sdk.NewInt(365 * 24 * 60 * 60))
	ptAPY := sdk.NewDecFromInt(pt.Amount.Sub(swapCoin.Amount)).QuoInt(swapCoin.Amount).Mul(maturityPerYear)
	ptRate := sdk.NewDecFromInt(pt.Amount).QuoInt(swapCoin.Amount)

	return &types.QueryTranchePtAPYsResponse{
		PtApy:            ptAPY,
		PtRatePerDeposit: ptRate,
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

	ytDenom := types.YtDenom(tranche)
	yt := sdk.NewCoin(ytDenom, sdk.NewInt(1_000_000))
	requiredDeposit, err := k.CalculateRequiredDepositSwapToYt(ctx, tranche, yt.Amount)
	if err != nil {
		return nil, err
	}
	if requiredDeposit.IsZero() {
		return &types.QueryTrancheYtAPYsResponse{
			YtApy:            sdk.ZeroDec(),
			YtRatePerDeposit: sdk.ZeroDec(),
		}, nil
	}

	// YT APY = stATOM APY * SwapRate (stATOM => YT) - 1
	info := k.GetStrategyDepositInfo(ctx, tranche.StrategyContract)
	lsDenomApy, err := sdk.NewDecFromStr(info.DepositDenomApy)
	if err != nil {
		lsDenomApy = sdk.ZeroDec()
	}
	ytRate := sdk.NewDecFromInt(yt.Amount).QuoInt(requiredDeposit.Amount)
	ytAPY := lsDenomApy.Mul(ytRate).Sub(sdk.OneDec())

	return &types.QueryTrancheYtAPYsResponse{
		YtApy:            ytAPY,
		YtRatePerDeposit: ytRate,
	}, nil
}

func (k Keeper) TranchePoolAPYs(c context.Context, req *types.QueryTranchePoolAPYsRequest) (*types.QueryTranchePoolAPYsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	swapCoin := sdk.NewCoin(tranche.DepositDenom, sdk.NewInt(1_000_000))
	pt, err := k.SimulateSwapPoolTokens(ctx, tranche, swapCoin)
	if err != nil {
		return nil, err
	}
	restMaturity := tranche.StartTime + tranche.Maturity - uint64(ctx.BlockTime().Unix())
	maturityPerYear := sdk.NewDecFromInt(sdk.NewIntFromUint64(restMaturity)).QuoInt(sdk.NewInt(365 * 24 * 60 * 60))
	ptAPY := sdk.NewDecFromInt(pt.Amount.Sub(swapCoin.Amount)).QuoInt(swapCoin.Amount).Mul(maturityPerYear)
	ptRate := sdk.NewDecFromInt(pt.Amount).QuoInt(swapCoin.Amount)

	ptDenom := types.PtDenom(tranche)
	var depositTokenPercentage sdk.Dec
	var ptPercentage sdk.Dec
	if len(tranche.PoolAssets) != 2 {
		return nil, types.ErrInvalidPoolAssets
	}
	if ptDenom == tranche.PoolAssets[0].Denom {
		total := sdk.NewDecFromInt(tranche.PoolAssets[1].Amount).Mul(ptRate).Add(sdk.NewDecFromInt(tranche.PoolAssets[0].Amount))
		depositTokenPercentage = sdk.NewDecFromInt(tranche.PoolAssets[1].Amount).Quo(total)
		ptPercentage = sdk.NewDecFromInt(tranche.PoolAssets[0].Amount).Quo(total)
	} else {
		total := sdk.NewDecFromInt(tranche.PoolAssets[0].Amount).Mul(ptRate).Add(sdk.NewDecFromInt(tranche.PoolAssets[1].Amount))
		depositTokenPercentage = sdk.NewDecFromInt(tranche.PoolAssets[0].Amount).Quo(total)
		ptPercentage = sdk.NewDecFromInt(tranche.PoolAssets[1].Amount).Quo(total)
	}
	discountPtAPY := ptAPY.Mul(ptPercentage)

	lpAmount := sdk.NewInt(1_000_000)
	requiredCoins, err := GetMaximalNoSwapLPAmount(ctx, tranche, lpAmount)
	if err != nil {
		return nil, err
	}
	var depositTokenAmount sdk.Dec
	if ptDenom == requiredCoins[0].Denom {
		depositTokenAmount = sdk.NewDecFromInt(requiredCoins[0].Amount).Quo(ptRate).Add(sdk.NewDecFromInt(requiredCoins[1].Amount))
	} else {
		depositTokenAmount = sdk.NewDecFromInt(requiredCoins[1].Amount).Quo(ptRate).Add(sdk.NewDecFromInt(requiredCoins[0].Amount))
	}
	lpRate := sdk.NewDecFromInt(lpAmount).Quo(depositTokenAmount)

	return &types.QueryTranchePoolAPYsResponse{
		LiquidityApy:                 sdk.ZeroDec(),
		LiquidityRatePerDeposit:      lpRate,
		DiscountPtApy:                discountPtAPY,
		DepositTokenPercentageInPool: depositTokenPercentage,
		PtPercentageInPool:           ptPercentage,
	}, nil
}
