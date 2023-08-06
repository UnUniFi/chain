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

	position.RemainingMargin = position.RemainingMargin.Add(amount)

	// Send additional margin by sender to the margin manager module account
	if err := k.SendMarginToMarginManager(ctx, sender, sdk.NewCoins(amount)); err != nil {
		return err
	}

	err := k.SetPosition(ctx, *position)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) RemoveMargin(ctx sdk.Context, withdrawer sdk.AccAddress, positionId string, amount sdk.Coin) error {
	position := k.GetPositionWithId(ctx, positionId)
	if position == nil {
		return types.ErrPositionDoesNotExist
	}

	// // Check withdrawer (sender) matches the owner of the position
	// if position.Address != withdrawer.String() {
	// 	return types.ErrUnauthorized
	// }

	// check withdrawer has owner nft
	owner := k.GetPositionNFTOwner(ctx, positionId)
	if owner.String() != withdrawer.String() {
		return types.ErrUnauthorized
	}

	// Update RemainingMargin
	updatedRemainingMargin, err := position.RemainingMargin.SafeSub(amount)
	// Check if the updated margin is positive
	if err != nil {
		return err
	}
	position.RemainingMargin = updatedRemainingMargin

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

	err = k.SetPosition(ctx, *position)
	if err != nil {
		return err
	}

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

func (k Keeper) SendPendingMargin(ctx sdk.Context, amount sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.MarginManager, types.PendingPaymentManager, amount)
}
