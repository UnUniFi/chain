package derivatives

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

func setPoolMarketCapSnapshot(ctx sdk.Context, k keeper.Keeper) {
	// move on only if price is ready
	if !k.IsPriceReady(ctx) {
		return
	}
	poolMarketCap, err := k.GetPoolMarketCap(ctx)
	if err != nil {
		return
	}

	k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), poolMarketCap)
}

func saveBlockTime(ctx sdk.Context, k keeper.Keeper) {
	k.SaveBlockTimestamp(ctx, ctx.BlockHeight(), ctx.BlockTime())
}

// BeginBlocker
// In mainnet, BeginBlocker will have no function.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	CheckPosition(ctx, k)
}

// todo: fixme this function is temporary treatment.
// In mainnet, this function will be executed by off chain bots.
func CheckPosition(ctx sdk.Context, k keeper.Keeper) {
	positions := k.GetAllPositions(ctx)
	params := k.GetParams(ctx)
	quoteTicker := k.GetPoolQuoteTicker(ctx)
	for _, position := range positions {
		currentBaseUsdRate, currentQuoteUsdRate, err := k.GetPairUsdPriceFromMarket(ctx, position.Market)
		if err != nil {
			// todo: user logger
			fmt.Println("failed to get pair usd price from market")
			fmt.Println(err)
			continue
		}
		baseMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.BaseDenom, currentBaseUsdRate)
		quoteMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.QuoteDenom, currentQuoteUsdRate)
		moduleAddr := k.GetModuleAddress()
		if position.NeedLiquidation(params.PerpetualFutures.MarginMaintenanceRate, baseMetricsRate, quoteMetricsRate) {
			msg := types.MsgReportLiquidation{
				Sender:          moduleAddr.String(),
				PositionId:      position.Id,
				RewardRecipient: moduleAddr.String(),
			}
			k.ReportLiquidation(ctx, &msg)
		}
		msg := types.MsgReportLevyPeriod{
			Sender:          moduleAddr.String(),
			PositionId:      position.Id,
			RewardRecipient: moduleAddr.String(),
		}
		k.ReportLevyPeriod(ctx, &msg)
	}
}

// EndBlocker
// In EndBlocker, the snapshot of pool market cap and block time are saved.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	setPoolMarketCapSnapshot(ctx, k)
	saveBlockTime(ctx, k)
}
