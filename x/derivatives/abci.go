package derivatives

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

func setPoolMarketCapSnapshot(ctx sdk.Context, k keeper.Keeper) {
	// move on only if price is ready
	if !IsPriceReady(ctx, k) {
		return
	}

	k.SetPoolMarketCapSnapshot(ctx, ctx.BlockHeight(), k.GetPoolMarketCap(ctx))
}

func IsPriceReady(ctx sdk.Context, k keeper.Keeper) bool {
	assets := k.GetPoolAssets(ctx)

	for _, asset := range assets {
		_, err := k.GetAssetPrice(ctx, asset.Denom)
		if err != nil {
			ctx.EventManager().EmitTypedEvent(&types.EventPriceIsNotFeeded{
				Asset: asset.String(),
			})

			return false
		}
	}

	return true
}

func saveBlockTime(ctx sdk.Context, k keeper.Keeper) {
	k.SaveBlockTimestamp(ctx, ctx.BlockHeight(), ctx.BlockTime())
}

// BeginBlocker
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	// TODO: make this function calling every 8 hours.
	// saving `last_levy_ifr_block_time` in store is one of ways to do so.
	// levyImaginaryFundingRate(ctx, k)
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	setPoolMarketCapSnapshot(ctx, k)
	saveBlockTime(ctx, k)
}
