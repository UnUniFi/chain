package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) AddMargin(ctx sdk.Context, sender sdk.AccAddress, positionId string, amount sdk.Coin) error {
	position := k.GetPositionWithId(ctx, positionId)
	if position == nil {
		return types.ErrPositionDoesNotExist
	}

	// Send additional margin by sender to the margin manager module account
	if err := k.SendMarginToMarginManager(ctx, sender, sdk.NewCoins(amount)); err != nil {
		return err
	}

	// Add margin to the position
	position.RemainingMargin = position.RemainingMargin.Add(amount)

	// Make sure if the position is not under the liquidation condition
	if position.RemainingMargin.IsNegative() {
		return types.ErrNegativeMargin
	}

	// Check if the updated margin is not under liquidation
	params := k.GetParams(ctx)
	currentBaseUsdRate, currentQuoteUsdRate, err := k.GetPairUsdPriceFromMarket(ctx, position.Market)
	if err != nil {
		return err
	}
	quoteTicker := k.GetPoolQuoteTicker(ctx)
	baseMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.BaseDenom, currentBaseUsdRate)
	quoteMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.QuoteDenom, currentQuoteUsdRate)

	if position.NeedLiquidation(params.PerpetualFutures.MarginMaintenanceRate, baseMetricsRate, quoteMetricsRate) {
		// Emit event that the position is needed to be liquidated
		_ = ctx.EventManager().EmitTypedEvent(&types.EventLiquidationNeeded{
			PositionId: positionId,
		})

		// Return err as the result of this tx
		return types.ErrLiquidationNeeded
	}

	k.SetPosition(ctx, *position)

	return nil
}

func (k Keeper) WithdrawMargin(ctx sdk.Context, withdrawer sdk.AccAddress, positionId string, amount sdk.Coin) error {
	position := k.GetPositionWithId(ctx, positionId)
	if position == nil {
		return types.ErrPositionDoesNotExist
	}

	// Check withdrawer (sender) matches the owner of the position
	if position.Address != withdrawer.String() {
		return types.ErrUnauthorized
	}

	// Update RemainingMargin
	position.RemainingMargin = position.RemainingMargin.Sub(amount)

	// Check if the updated margin is positive
	if position.RemainingMargin.IsNegative() {
		return types.ErrNegativeMargin
	}

	// Check if the updated margin is not under liquidation
	params := k.GetParams(ctx)
	currentBaseUsdRate, currentQuoteUsdRate, err := k.GetPairUsdPriceFromMarket(ctx, position.Market)
	if err != nil {
		return err
	}
	quoteTicker := k.GetPoolQuoteTicker(ctx)
	baseMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.BaseDenom, currentBaseUsdRate)
	quoteMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.QuoteDenom, currentQuoteUsdRate)

	if position.NeedLiquidation(params.PerpetualFutures.MarginMaintenanceRate, baseMetricsRate, quoteMetricsRate) {
		// Return err as the result of this tx
		return types.ErrTooMuchMarginToWithdraw
	}

	// Send margin from the margin manager module account to the withdrawer
	if err := k.SendBackMargin(ctx, withdrawer, sdk.NewCoins(amount)); err != nil {
		return err
	}

	k.SetPosition(ctx, *position)

	return nil
}

func (k Keeper) SendMarginToMarginManager(ctx sdk.Context, sender sdk.AccAddress, margin sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.MarginManager, margin)
}

func (k Keeper) SendCoinFromMarginManagerToPool(ctx sdk.Context, amount sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.MarginManager, types.ModuleName, amount)
}

func (k Keeper) SendCoinFromPoolToMarginManager(ctx sdk.Context, amount sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.MarginManager, amount)
}

func (k Keeper) SendBackMargin(ctx sdk.Context, recipient sdk.AccAddress, amount sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.MarginManager, recipient, amount)
}
