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

	queriedPositions, err := k.MakeQueriedPositions(ctx, positions)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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
	positions := k.GetAddressPositionsVal(ctx, address)

	queriedPositions, err := k.MakeQueriedPositions(ctx, positions)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAddressPositionsResponse{
		Positions: queriedPositions,
	}, nil
}

func (k Keeper) MakeQueriedPositions(ctx sdk.Context, positions types.Positions) ([]types.QueriedPosition, error) {
	queriedPositions := make([]types.QueriedPosition, 0)
	usdMap := map[string]sdk.Dec{}
	for _, position := range positions {

		if _, ok := usdMap[position.Market.BaseDenom]; !ok {
			price, err := k.GetCurrentPrice(ctx, position.Market.BaseDenom)
			if err != nil {
				return nil, err
			}
			usdMap[position.Market.BaseDenom] = price
		}

		if _, ok := usdMap[position.Market.QuoteDenom]; !ok {
			price, err := k.GetCurrentPrice(ctx, position.Market.QuoteDenom)
			if err != nil {
				return nil, err
			}
			usdMap[position.Market.QuoteDenom] = price
		}

		if _, ok := usdMap[position.RemainingMargin.Denom]; !ok {
			price, err := k.GetCurrentPrice(ctx, position.RemainingMargin.Denom)
			if err != nil {
				return nil, err
			}
			usdMap[position.RemainingMargin.Denom] = price
		}
		currentBaseUsdRate := usdMap[position.Market.BaseDenom]
		currentQuoteUsdRate := usdMap[position.Market.QuoteDenom]
		currentMarginUsdRate := usdMap[position.RemainingMargin.Denom]

		perpetualFuturesPosition, err := types.NewPerpetualFuturesPositionFromPosition(position)
		if err != nil {
			// not implemented perpetual options
			return nil, status.Error(codes.Internal, err.Error())
		}

		closedPairRate := currentBaseUsdRate.Quo(currentQuoteUsdRate)
		profit := perpetualFuturesPosition.CalcProfitAndLoss(closedPairRate)
		queriedPosition := types.QueriedPosition{
			Position:              position,
			ValuationProfit:       sdk.NewCoin("uusd", profit),
			MarginMaintenanceRate: perpetualFuturesPosition.MarginMaintenanceRate(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate),
			EffectiveMargin:       sdk.NewCoin("uusd", types.NormalToMicroDenom(perpetualFuturesPosition.EffectiveMargin(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate))),
		}
		queriedPositions = append(queriedPositions, queriedPosition)
	}
	return queriedPositions, nil
}

func (k Keeper) Position(c context.Context, req *types.QueryPositionRequest) (*types.QueryPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	position := k.GetPositionWithId(ctx, req.PositionId)
	if position == nil {
		return &types.QueryPositionResponse{}, nil
	}

	perpetualFuturesPosition, err := types.NewPerpetualFuturesPositionFromPosition(*position)
	if err != nil {
		panic(err)
	}
	currentBaseUsdRate, err := k.GetCurrentPrice(ctx, position.Market.BaseDenom)
	if err != nil {
		panic(err)
	}

	currentQuoteUsdRate, err := k.GetCurrentPrice(ctx, position.Market.QuoteDenom)
	if err != nil {
		panic(err)
	}
	currentMarginUsdRate, err := k.GetCurrentPrice(ctx, position.RemainingMargin.Denom)
	if err != nil {
		panic(err)
	}
	closedPairRate := currentBaseUsdRate.Quo(currentQuoteUsdRate)
	profit := perpetualFuturesPosition.CalcProfitAndLoss(closedPairRate)
	return &types.QueryPositionResponse{
		Position:              position,
		ValuationProfit:       sdk.NewCoin("uusd", profit),
		MarginMaintenanceRate: perpetualFuturesPosition.MarginMaintenanceRate(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate),
		EffectiveMargin:       sdk.NewCoin("uusd", types.NormalToMicroDenom(perpetualFuturesPosition.EffectiveMargin(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate))),
	}, nil
}

func (k Keeper) PerpetualFuturesPositionSize(c context.Context, req *types.QueryPerpetualFuturesPositionSizeRequest) (*types.QueryPerpetualFuturesPositionSizeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	positions := types.Positions(k.GetAddressPositionsVal(ctx, address))
	getPriceFunc := func(ctx sdk.Context) func(denom string) (sdk.Dec, error) {
		return func(denom string) (sdk.Dec, error) {
			return k.GetCurrentPrice(ctx, denom)
		}
	}
	var result sdk.Dec
	if req.PositionType == types.PositionType_LONG {
		result = positions.EvaluateLongPositions(getPriceFunc(ctx))
	} else if req.PositionType == types.PositionType_SHORT {
		result = positions.EvaluateShortPositions(getPriceFunc(ctx))
	} else {
		return nil, status.Error(codes.InvalidArgument, "invalid position type")
	}
	return &types.QueryPerpetualFuturesPositionSizeResponse{
		TotalPositionSizeUsd: sdk.NewCoin("usd", result.RoundInt()),
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
