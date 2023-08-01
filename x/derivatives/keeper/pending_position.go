package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) SetPendingPaymentPosition(ctx sdk.Context, position types.Position, amount sdk.Coin) {
	pendingPaymentPosition := types.PendingPaymentPosition{
		Id:               position.Id,
		RefundableAmount: amount,
		CreatedAt:        ctx.BlockTime(),
		CreatedHeight:    uint64(ctx.BlockHeight()),
	}
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&pendingPaymentPosition)
	store.Set(types.PendingPaymentPositionWithIdKeyPrefix(position.Id), bz)
}

func (k Keeper) DeletePendingPaymentPosition(ctx sdk.Context, Id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PendingPaymentPositionWithIdKeyPrefix(Id))
}

func (k Keeper) GetPendingPaymentPosition(ctx sdk.Context, Id string) *types.PendingPaymentPosition {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PendingPaymentPositionWithIdKeyPrefix(Id))
	if bz == nil {
		return nil
	}
	pendingPaymentPosition := types.PendingPaymentPosition{}
	k.cdc.MustUnmarshal(bz, &pendingPaymentPosition)
	return &pendingPaymentPosition
}

func (k Keeper) ClosePendingPaymentPosition(ctx sdk.Context, pendingPosition types.PendingPaymentPosition, address string) error {
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}
	if pendingPosition.RefundableAmount.IsLT(pendingPosition.RemainingMargin) {
		loss := pendingPosition.RemainingMargin.Sub(pendingPosition.RefundableAmount)
		// Send margin-loss from MarginManager
		if err := k.SendBackMargin(ctx, addr, sdk.NewCoins(pendingPosition.RefundableAmount)); err != nil {
			return err
		}
		// Send loss to the pool
		if err := k.SendCoinFromMarginManagerToPool(ctx, sdk.NewCoins(loss)); err != nil {
			return err
		}
	} else {
		profit := pendingPosition.RefundableAmount.Sub(pendingPosition.RemainingMargin)
		// Send margin from MarginManager
		if err := k.SendBackMargin(ctx, addr, sdk.NewCoins(pendingPosition.RemainingMargin)); err != nil {
			return err
		}
		// Send profit from the pool
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(profit)); err != nil {
			return err
		}
	}
	k.DeletePendingPaymentPosition(ctx, pendingPosition.Id)
	return nil
}
