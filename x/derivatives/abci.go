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

	k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), k.GetPoolMarketCap(ctx))
}

func saveBlockTime(ctx sdk.Context, k keeper.Keeper) {
	k.SaveBlockTimestamp(ctx, ctx.BlockHeight(), ctx.BlockTime())
}

// BeginBlocker
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	// TODO: make this function calling every 8 hours.
	// saving `last_levy_ifr_block_time` in store is one of ways to do so.
	// levyImaginaryFundingRate(ctx, k)
	CheckPosition(ctx, k)
}

// todo: fixme this function is temporary treatment.
func CheckPosition(ctx sdk.Context, k keeper.Keeper) {
	positions := k.GetAllPositions(ctx)
	params := k.GetParams(ctx)
	for _, position := range positions {
		if position.Validate() != nil {
			// this is temporary treatment.
			fmt.Println("this is temporary treatment.")
			posForLog, _ := types.NewPerpetualFuturesPositionFromPosition(position)
			fmt.Println("deleted position:")
			fmt.Println(posForLog.String())

			k.DeletePosition(ctx, sdk.AccAddress(position.Address), position.Id)
			continue
		}
		// this is temporary treatment.
		// if position id is 0,3,5,7,9 or 10 delete position
		if position.Id == "0" || position.Id == "3" || position.Id == "5" || position.Id == "7" || position.Id == "9" || position.Id == "10" {
			posForLog, _ := types.NewPerpetualFuturesPositionFromPosition(position)
			fmt.Println("deleted position for users:")
			fmt.Println(posForLog.String())
			k.DeletePosition(ctx, sdk.AccAddress(position.Address), position.Id)
			continue
		}

		currentBaseUsdRate, currentQuoteUsdRate, err := k.GetPairUsdPriceFromMarket(ctx, position.Market)
		if err != nil {
			// todo: user logger
			fmt.Println("failed to get pair usd price from market")
			fmt.Println(err)
			continue
		}
		if position.NeedLiquidation(params.PerpetualFutures.MarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate) {
			msg := types.MsgReportLiquidation{
				Sender:          position.Address,
				PositionId:      position.Id,
				RewardRecipient: position.Address,
			}
			k.ReportLiquidation(ctx, &msg)
		}
		msg := types.MsgReportLevyPeriod{
			Sender:          position.Address,
			PositionId:      position.Id,
			RewardRecipient: position.Address,
		}
		k.ReportLevyPeriod(ctx, &msg)
	}
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	setPoolMarketCapSnapshot(ctx, k)
	saveBlockTime(ctx, k)
}
