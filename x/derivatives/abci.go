package derivatives

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

func levyImaginaryFundingRateAndLiquidateInsufficientMarginPositions(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)
	wrappedPositions := k.GetAllPositions(ctx)
	assets := k.GetPoolAssets(ctx)
	fundingRateProportionalCoefficient := params.PerpetualFutures.ImaginaryFundingRateProportionalCoefficient
	commissionRate := params.PerpetualFutures.CommissionRate

	imaginaryFundingRates := make(map[string]sdk.Dec)

	for _, asset := range assets {
		netPosition := k.GetPerpetualFuturesNetPositionOfDenom(ctx, asset.Denom)
		imaginaryFundingRate := netPosition.Mul(fundingRateProportionalCoefficient)
		imaginaryFundingRates[asset.Denom] = imaginaryFundingRate
	}

	for _, wrappedPosition := range wrappedPositions {
		positionId := types.GetPositionIdFromString(wrappedPosition.Id)
		position, err := types.UnpackPosition(&wrappedPosition.Position)
		if err != nil {
			panic("unable to unpack open position")
		}

		switch position.(type) {
		case *types.PerpetualFuturesPosition:
			futuresPosition := position.(*types.PerpetualFuturesPosition)
			depositedMargin := k.GetDepositedMargin(ctx, positionId)
			imaginaryFundingRate := imaginaryFundingRates[futuresPosition.Pair.Denom]
			imaginaryFundingFee := sdk.NewDecFromInt(depositedMargin.Amount).Mul(imaginaryFundingRate).RoundInt()
			commissionFee := sdk.NewDecFromInt(depositedMargin.Amount).Mul(commissionRate).RoundInt()

			principal := types.CalculatePrincipal(*futuresPosition)

			if imaginaryFundingRate.IsNegative() {
				if futuresPosition.PositionType == types.PositionType_SHORT {
					depositedMargin.Amount = depositedMargin.Amount.Sub(imaginaryFundingFee)
				} else {
					depositedMargin.Amount = depositedMargin.Amount.Add(imaginaryFundingFee.Sub(commissionFee))
				}
			} else {
				if futuresPosition.PositionType == types.PositionType_LONG {
					depositedMargin.Amount = depositedMargin.Amount.Sub(imaginaryFundingFee)
				} else {
					depositedMargin.Amount = depositedMargin.Amount.Add(imaginaryFundingFee.Sub(commissionFee))
				}
			}

			k.SaveDepositedMargin(ctx, positionId, depositedMargin)

			if sdk.NewDecFromInt(depositedMargin.Amount).Mul(sdk.NewDecWithPrec(1, 0)).LT(principal.Mul(params.PerpetualFutures.MarginMaintenanceRate)) {
				k.ClosePerpetualFuturesPosition(ctx, wrappedPosition.Address.AccAddress(), positionId, futuresPosition)
			}
			return
		case *types.PerpetualOptionsPosition:
			return
		}
	}
}

func setPoolMarketCapSnapshot(ctx sdk.Context, k keeper.Keeper) {
	k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), k.GetPoolMarketCap(ctx))
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	levyImaginaryFundingRateAndLiquidateInsufficientMarginPositions(ctx, k)
	setPoolMarketCapSnapshot(ctx, k)
}
