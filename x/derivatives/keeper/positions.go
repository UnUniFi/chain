package keeper

import (
	"fmt"
	"math/big"
	"time"

	cdcTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetUserPositions(ctx sdk.Context, user sdk.AccAddress) []types.WrappedPosition {
	store := ctx.KVStore(k.storeKey)

	positions := []types.WrappedPosition{}
	it := sdk.KVStorePrefixIterator(store, types.AddressPositionKeyPrefix(user))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.WrappedPosition{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, position)
	}

	return positions
}

func (k Keeper) CreatePosition(ctx sdk.Context, wrappedPosition types.WrappedPosition) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&wrappedPosition)
	store.Set(types.AddressPositionWithIdKeyPrefix(wrappedPosition.Address.AccAddress(), 0), bz) // TODO: id
}

func (k Keeper) GetPosition(ctx sdk.Context, address sdk.AccAddress, id int) types.WrappedPosition {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.AddressPositionWithIdKeyPrefix(address, id))
	position := types.WrappedPosition{}
	k.cdc.Unmarshal(bz, &position)

	return position
}

func (k Keeper) DeletePosition(ctx sdk.Context, address sdk.AccAddress, id int) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.AddressPositionWithIdKeyPrefix(address, id))
}

func (k Keeper) CreateClosedPosition(ctx sdk.Context, wrappedPosition types.WrappedPosition) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&wrappedPosition)
	store.Set(types.AddressClosedPositionWithIdKeyPrefix(wrappedPosition.Address.AccAddress(), 0), bz) // TODO: id
}

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) error {
	sender := msg.Sender.AccAddress()
	positions := k.GetUserPositions(ctx, sender)
	positionCount := len(positions)

	positionKey := types.AddressPositionWithIdKeyPrefix(sender, positionCount+1)

	wrappedPosition := types.WrappedPosition{
		Id:       string(positionKey), // TODO
		Address:  msg.Sender,
		StartAt:  *timestamppb.New(time.Now()), // TODO
		Position: msg.Position,
	}

	// Not sure how to convert any type to position type
	k.CreatePosition(ctx, wrappedPosition)

	position, err := types.UnpackPosition(&msg.Position)
	if err != nil {
		return err
	}
	switch position.(type) {
	case *types.PerpetualFuturesPosition:
		return k.OpenPerpetualFuturesPosition(ctx, msg.Sender.AccAddress(), position.(*types.PerpetualFuturesPosition))
	case *types.PerpetualOptionsPosition:
		return k.OpenPerpetualOptionsPosition(ctx, msg.Sender.AccAddress(), position.(*types.PerpetualOptionsPosition))
	}

	return nil
}

func (k Keeper) Claim(ctx sdk.Context, msg *types.MsgClaim) error {

	return nil
}

func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) error {
	wrappedPosition := k.GetPosition(ctx, msg.Sender.AccAddress(), 0) // TODO: id
	k.DeletePosition(ctx, msg.Sender.AccAddress(), 0)                 // TODO: id

	k.CreateClosedPosition(ctx, wrappedPosition)

	position, err := types.UnpackPosition(&wrappedPosition.Position)
	if err != nil {
		return err
	}
	switch position.(type) {
	case *types.PerpetualFuturesPosition:
		return k.ClosePerpetualFuturesPosition(ctx, msg.Sender.AccAddress(), position.(*types.PerpetualFuturesPosition))
	case *types.PerpetualOptionsPosition:
		return k.ClosePerpetualOptionsPosition(ctx, msg.Sender.AccAddress(), position.(*types.PerpetualOptionsPosition))
	}

	return nil
}
