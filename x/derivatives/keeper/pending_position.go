package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) SetPendingPaymentPosition(ctx sdk.Context, positionId string, amount sdk.Coin) {
	pendingPaymentPosition := types.PendingPaymentPosition{
		Id:               positionId,
		RefundableAmount: amount,
		CreatedAt:        ctx.BlockTime(),
		CreatedHeight:    uint64(ctx.BlockHeight()),
	}
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&pendingPaymentPosition)
	store.Set(types.PendingPaymentPositionWithIdKeyPrefix(positionId), bz)
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

func (k Keeper) ClosePendingPaymentPosition(ctx sdk.Context, pendingPosition types.PendingPaymentPosition, address sdk.AccAddress) error {
	err := k.ClosePendingPaymentFuturePosition(ctx, pendingPosition, address)
	if err == nil {
		return nil
	}
	err = k.ClosePendingPaymentOptionPosition(ctx, pendingPosition, address)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) ClosePendingPaymentFuturePosition(ctx sdk.Context, pendingPosition types.PendingPaymentPosition, address sdk.AccAddress) error {
	owner := k.GetFuturePositionNFTOwner(ctx, pendingPosition.Id)
	if owner.String() != address.String() {
		return types.ErrUnauthorized
	}
	sendDisabled, err := k.GetFuturePositionNFTSendDisabled(ctx, pendingPosition.Id)
	if err != nil {
		return err
	}
	if sendDisabled {
		return types.ErrPositionNFTSendDisabled
	}
	if pendingPosition.RefundableAmount.IsPositive() {
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.PendingPaymentManager, address, sdk.NewCoins(pendingPosition.RefundableAmount))
		if err != nil {
			return err
		}
	}
	err = k.CloseFuturePositionNFT(ctx, pendingPosition.Id)
	if err != nil {
		return err
	}
	k.DeletePendingPaymentPosition(ctx, pendingPosition.Id)
	return nil
}

func (k Keeper) ClosePendingPaymentOptionPosition(ctx sdk.Context, pendingPosition types.PendingPaymentPosition, address sdk.AccAddress) error {
	// todo: impl
	return types.ErrNotImplemented
}
