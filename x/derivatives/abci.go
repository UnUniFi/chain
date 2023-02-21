package derivatives

import (
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
	CheckPotision(ctx, k)
}

// todo: fixme this function is temporary treatment.
func CheckPotision(ctx sdk.Context, k keeper.Keeper) {
	potisions := k.GetAllPositions(ctx)
	params := k.GetParams(ctx)
	for _, potision := range potisions {
		if potision.NeedLiqudation(params.PerpetualFutures.MarginMaintenanceRate) {
			msg := types.MsgReportLiquidation{
				Sender:          potision.Address,
				PositionId:      potision.Id,
				RewardRecipient: potision.Address,
			}
			k.ReportLiquidation(ctx, &msg)
		}
		msg := types.MsgReportLevyPeriod{
			Sender:          potision.Address,
			PositionId:      potision.Id,
			RewardRecipient: potision.Address,
		}
		k.ReportLevyPeriod(ctx, &msg)
	}
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	setPoolMarketCapSnapshot(ctx, k)
	saveBlockTime(ctx, k)
}
