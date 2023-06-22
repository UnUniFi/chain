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
	// If the position is under the liquidation condition, the position should be liquidated
	err := k.ReportLiquidation(
		ctx,
		&types.MsgReportLiquidation{
			Sender:          sender.String(),
			PositionId:      positionId,
			RewardRecipient: sender.String(),
		})
	if err != nil && err != types.ErrLiquidationNotNeeded {
		return err
	} else if err == nil {
		// If the position is liquidated, the position should've been deleted already
		return nil
	} else {
		// If the position is not liquidated, the position should be updated
		k.SetPosition(ctx, *position)
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
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, amount)
}
