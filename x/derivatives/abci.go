package derivatives

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
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
}

func setPoolMarketCapSnapshot(ctx sdk.Context, k keeper.Keeper) {
	k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), k.GetPoolMarketCap(ctx))
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	levyImaginaryFundingRateAndLiquidateInsufficientMarginPositions(ctx, k)
	setPoolMarketCapSnapshot(ctx, k)
}
