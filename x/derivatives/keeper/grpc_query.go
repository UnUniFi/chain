package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) LiquidityProviderTokenRealAPY(c context.Context, req *types.QueryLiquidityProviderTokenRealAPYRequest) (*types.QueryLiquidityProviderTokenRealAPYResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	rate := k.GetLPNominalYieldRate(ctx, req.BeforeHeight, req.AfterHeight)
	annualized := k.AnnualizeYieldRate(ctx, rate, req.BeforeHeight, req.AfterHeight)

	return &types.QueryLiquidityProviderTokenRealAPYResponse{Apy: &annualized}, nil
}

func (k Keeper) LiquidityProviderTokenNominalAPY(c context.Context, req *types.QueryLiquidityProviderTokenNominalAPYRequest) (*types.QueryLiquidityProviderTokenNominalAPYResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	rate := k.GetLPNominalYieldRate(ctx, req.BeforeHeight, req.AfterHeight)
	annualized := k.AnnualizeYieldRate(ctx, rate, req.BeforeHeight, req.AfterHeight)

	return &types.QueryLiquidityProviderTokenNominalAPYResponse{Apy: &annualized}, nil
}

func (k Keeper) PerpetualFutures(c context.Context, req *types.QueryPerpetualFuturesRequest) (*types.QueryPerpetualFuturesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	ctx.BlockHeight()
	metricsQuoteTicker := ""
	volume24Hours := sdk.NewDec(0)
	fees24Hours := sdk.NewDec(0)
	longPositions := sdk.NewDec(0)
	shortPositions := sdk.NewDec(0)

	return &types.QueryPerpetualFuturesResponse{
		MetricsQuoteTicker: metricsQuoteTicker,
		Volume_24Hours:     &volume24Hours,
		Fees_24Hours:       &fees24Hours,
		LongPositions:      &longPositions,
		ShortPositions:     &shortPositions,
	}, nil
}

func (k Keeper) PerpetualFuturesMarket(c context.Context, req *types.QueryPerpetualFuturesMarketRequest) (*types.QueryPerpetualFuturesMarketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	ctx.BlockHeight()
	price := sdk.NewDec(0)
	metricsQuoteTicker := ""
	volume24Hours := sdk.NewDec(0)
	fees24Hours := sdk.NewDec(0)
	longPositions := sdk.NewDec(0)
	shortPositions := sdk.NewDec(0)

	return &types.QueryPerpetualFuturesMarketResponse{
		Price:              &price,
		MetricsQuoteTicker: metricsQuoteTicker,
		Volume_24Hours:     &volume24Hours,
		Fees_24Hours:       &fees24Hours,
		LongPositions:      &longPositions,
		ShortPositions:     &shortPositions,
	}, nil
}

func (k Keeper) PerpetualOptions(c context.Context, req *types.QueryPerpetualOptionsRequest) (*types.QueryPerpetualOptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	ctx.BlockHeight()

	return &types.QueryPerpetualOptionsResponse{}, nil
}

func (k Keeper) PerpetualOptionsMarket(c context.Context, req *types.QueryPerpetualOptionsMarketRequest) (*types.QueryPerpetualOptionsMarketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	ctx.BlockHeight()

	return &types.QueryPerpetualOptionsMarketResponse{}, nil
}

func (k Keeper) Pool(c context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	metricsQuoteTicker := ""
	poolMarketCap := k.GetPoolMarketCap(ctx)
	volume24Hours := sdk.NewDec(0)
	fees24Hours := sdk.NewDec(0)

	return &types.QueryPoolResponse{
		MetricsQuoteTicker: metricsQuoteTicker,
		PoolMarketCap:      &poolMarketCap,
		Volume_24Hours:     &volume24Hours,
		Fees_24Hours:       &fees24Hours,
	}, nil
}

func (k Keeper) AllPositions(c context.Context, req *types.QueryAllPositionsRequest) (*types.QueryAllPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: pagination

	ctx := sdk.UnwrapSDKContext(c)
	positions := k.GetAllPositions(ctx)

	queriedPositions := make([]types.QueriedPosition, 0)
	for _, position := range positions {
		queriedPosition := types.QueriedPosition{
			Position:    position,
			QuoteTicker: "USD",
		}
		ins, err := types.UnpackPositionInstance(position.PositionInstance)
		if err != nil {
			return nil, status.Error(codes.Internal, "internal error")
		}

		switch positionInstance := ins.(type) {
		case *types.PerpetualFuturesPositionInstance:
			break
		case *types.PerpetualOptionsPositionInstance:
			panic(positionInstance)
			break
		}

		queriedPositions = append(queriedPositions, queriedPosition)
	}

	return &types.QueryAllPositionsResponse{
		Positions: queriedPositions,
	}, nil
}

func (k Keeper) AddressPositions(c context.Context, req *types.QueryAddressPositionsRequest) (*types.QueryAddressPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: pagination

	ctx := sdk.UnwrapSDKContext(c)
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	positions := k.GetAddressPositions(ctx, address)

	queriedPoisitions := make([]types.QueriedPosition, 0)
	for _, position := range positions {
		queriedPosition := types.QueriedPosition{
			Position:    *position,
			QuoteTicker: "USD",
		}
		ins, err := types.UnpackPositionInstance(position.PositionInstance)
		if err != nil {
			return nil, status.Error(codes.Internal, "internal error")
		}

		switch positionInstance := ins.(type) {
		case *types.PerpetualFuturesPositionInstance:
			break
		case *types.PerpetualOptionsPositionInstance:
			panic(positionInstance)
			break
		}

		queriedPoisitions = append(queriedPoisitions, queriedPosition)
	}

	return &types.QueryAddressPositionsResponse{
		Positions: queriedPoisitions,
	}, nil
}

func (k Keeper) EstimateDLPTokenAmount(c context.Context, req *types.QueryEstimateDLPTokenAmountRequest) (*types.QueryEstimateDLPTokenAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	mintAmount, mintFee, err := k.DetermineMintingLPTokenAmount(ctx, sdk.NewCoin(req.MintDenom, *req.Amount))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryEstimateDLPTokenAmountResponse{
		Amount: &mintAmount.Amount,
		Fee:    &mintFee.Amount,
	}, nil
}

func (k Keeper) EstimateRedeemAmount(c context.Context, req *types.QueryEstimateRedeemAmountRequest) (*types.QueryEstimateRedeemAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	redeemAmount, redeemFee, err := k.GetRedeemDenomAmount(ctx, *req.LptAmount, req.RedeemDenom)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryEstimateRedeemAmountResponse{
		Amount: &redeemAmount.Amount,
		Fee:    &redeemFee.Amount,
	}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
