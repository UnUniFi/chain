package derivatives

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

func levyImaginaryFundingRate(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)
	positions := k.GetAllPositions(ctx)

	perpetualFuturesMarkets := params.PerpetualFutures.Markets
	perpetualFuturesImaginaryFundingRates := make(map[types.Market]sdk.Dec)

	for _, perpetualFuturesMarket := range perpetualFuturesMarkets {
		netPosition := k.GetPerpetualFuturesNetPositionOfMarket(ctx, *perpetualFuturesMarket)
		imaginaryFundingRate := netPosition.Mul(params.PerpetualFutures.ImaginaryFundingRateProportionalCoefficient)
		perpetualFuturesImaginaryFundingRates[*perpetualFuturesMarket] = imaginaryFundingRate
	}

	for _, position := range positions {
		positionInstance, err := types.UnpackPositionInstance(position.PositionInstance)
		if err != nil {
			panic("unable to unpack open position")
		}

		switch positionInstance.(type) {
		case *types.PerpetualFuturesPositionInstance:
			futuresPosition := positionInstance.(*types.PerpetualFuturesPositionInstance)
			remainingMargin := position.RemainingMargin
			imaginaryFundingRate := perpetualFuturesImaginaryFundingRates[position.Market]
			imaginaryFundingFee := sdk.NewDecFromInt(remainingMargin.Amount).Mul(imaginaryFundingRate).RoundInt()
			commissionFee := sdk.NewDecFromInt(remainingMargin.Amount).Mul(params.PerpetualFutures.CommissionRate).RoundInt()

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

			return
		case *types.PerpetualOptionsPositionInstance:
			return
		}
	}
}

func setPoolMarketCapSnapshot(ctx sdk.Context, k keeper.Keeper) {
	k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), k.GetPoolMarketCap(ctx))
}

func saveBlockTime(ctx sdk.Context, k keeper.Keeper) {
	k.SaveBlockTimestamp(ctx, ctx.BlockHeight(), ctx.BlockTime())
}

// BeginBlocker
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	// TODO: make this function calling every 8 hours.
	// saving `last_levy_ifr_block_time` in store is one of ways to do so.
	levyImaginaryFundingRate(ctx, k)
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	setPoolMarketCapSnapshot(ctx, k)
	saveBlockTime(ctx, k)
}
