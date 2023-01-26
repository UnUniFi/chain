package derivatives

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

func levyImaginaryFundingRateAndLiquidateInsufficientMarginPositions(ctx sdk.Context, k keeper.Keeper) {
	// TODO: Iterate open position by for loop and levy imaginary funding rate from margin (principal, collateral).
	// imaginary_funding_rate increases in proportion to the net position of traders.
	// If traders' net position is long, imaginary funding rate is positive. If traders' net position is short, imaginary funding rate is negative.
	// If FR is positive, opened long positions are levied the imaginary funding rate and opened short positions get the funding rate (without imaginary funding rate commission rate).
	// If FR is negative, opened short positions are levied the imaginary funding rate and opened long positions get the funding rate (without imaginary funding rate commission rate).
	// imaginary funding rate commission rate is needed to be defined in Params
	// TODO: If position_profit - principal is close to zero (e.g. under 50% of principal), forcibly close the position.
	// This "50%" should also be parameterized in Params. The name should be "margin_maintenance_rate": sdk.Dec
	params := k.GetParams(ctx)
	wrappedPositions := k.GetAllPositions(ctx)
	assets := k.GetPoolAssets(ctx)
	fundingRateProportionalCoefficient := params.FundingRateProportionalCoefficient
	commissionRate := params.CommissionRate

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
			imaginaryFundingRate := imaginaryFundingRates[futuresPosition.Denom]
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

			if sdk.NewDecFromInt(depositedMargin.Amount).Mul(sdk.NewDecWithPrec(1, 0)).LT(principal.Mul(params.MarginMaintenanceRate)) {
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
