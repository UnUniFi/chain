package derivatives

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

func levyImaginaryFundingRate(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)
	positions := k.GetAllPositions(ctx)
	assets := k.GetPoolAssets(ctx)
	fundingRateProportionalCoefficient := params.PerpetualFutures.ImaginaryFundingRateProportionalCoefficient
	commissionRate := params.PerpetualFutures.CommissionRate

	imaginaryFundingRates := make(map[types.Market]sdk.Dec)

	for _, asset := range assets {
		netPosition := k.GetPerpetualFuturesNetPositionOfMarket(ctx, asset.Denom)
		imaginaryFundingRate := netPosition.Mul(fundingRateProportionalCoefficient)
		imaginaryFundingRates[asset.Denom] = imaginaryFundingRate
	}

	for _, position := range positions {
		positionInstance, err := types.UnpackPositionInstance(position.PositionInstance)
		if err != nil {
			panic("unable to unpack open position")
		}

		switch positionInstance.(type) {
		case *types.PerpetualFuturesPositionInstance:
			futuresPosition := positionInstance.(*types.PerpetualFuturesPositionInstance)
			remainingMargin := *k.GetRemainingMargin(ctx, position.Id)
			imaginaryFundingRate := imaginaryFundingRates[position.Market]
			imaginaryFundingFee := sdk.NewDecFromInt(remainingMargin.Amount).Mul(imaginaryFundingRate).RoundInt()
			commissionFee := sdk.NewDecFromInt(remainingMargin.Amount).Mul(commissionRate).RoundInt()

			if imaginaryFundingRate.IsNegative() {
				if futuresPosition.PositionType == types.PositionType_SHORT {
					remainingMargin.Amount = remainingMargin.Amount.Sub(imaginaryFundingFee)
				} else {
					remainingMargin.Amount = remainingMargin.Amount.Add(imaginaryFundingFee.Sub(commissionFee))
				}
			} else {
				if futuresPosition.PositionType == types.PositionType_LONG {
					remainingMargin.Amount = remainingMargin.Amount.Sub(imaginaryFundingFee)
				} else {
					remainingMargin.Amount = remainingMargin.Amount.Add(imaginaryFundingFee.Sub(commissionFee))
				}
			}

			k.SetRemainingMargin(ctx, position.Id, remainingMargin)

			return
		case *types.PerpetualOptionsPositionInstance:
			return
		}
	}
}

func setPoolMarketCapSnapshot(ctx sdk.Context, k keeper.Keeper) {
	k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), k.GetPoolMarketCap(ctx))
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	levyImaginaryFundingRate(ctx, k)
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	setPoolMarketCapSnapshot(ctx, k)
}
