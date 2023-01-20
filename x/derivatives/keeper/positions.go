package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetLastPositionId(ctx sdk.Context) (id uint64) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(types.KeyPrefixLastPositionId))
	if bz == nil {
		panic("last position id not set in genesis")
	}

	id = types.GetPositionIdFromBytes(bz)
	return
}

func (k Keeper) GetAllPositions(ctx sdk.Context) []*types.WrappedPosition {
	store := ctx.KVStore(k.storeKey)

	positions := []*types.WrappedPosition{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixPosition))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.WrappedPosition{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, &position)
	}

	return positions
}

func (k Keeper) GetUserPositions(ctx sdk.Context, user sdk.AccAddress) []*types.WrappedPosition {
	store := ctx.KVStore(k.storeKey)

	positions := []*types.WrappedPosition{}
	it := sdk.KVStorePrefixIterator(store, types.AddressPositionKeyPrefix(user))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.WrappedPosition{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, &position)
	}

	return positions
}

func (k Keeper) CreatePosition(ctx sdk.Context, wrappedPosition types.WrappedPosition) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&wrappedPosition)
	positionId := types.GetPositionIdFromString(wrappedPosition.Id)
	store.Set(types.AddressPositionWithIdKeyPrefix(wrappedPosition.Address.AccAddress(), positionId), bz)
}

func (k Keeper) GetPosition(ctx sdk.Context, address sdk.AccAddress, id uint64) types.WrappedPosition {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.AddressPositionWithIdKeyPrefix(address, id))
	position := types.WrappedPosition{}
	k.cdc.Unmarshal(bz, &position)

	return position
}

func (k Keeper) DeletePosition(ctx sdk.Context, address sdk.AccAddress, id uint64) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.AddressPositionWithIdKeyPrefix(address, id))
}

func (k Keeper) CreateClosedPosition(ctx sdk.Context, wrappedPosition types.WrappedPosition) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&wrappedPosition)
	positionId := types.GetPositionIdFromString(wrappedPosition.Id)
	store.Set(types.AddressClosedPositionWithIdKeyPrefix(wrappedPosition.Address.AccAddress(), positionId), bz)
}

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) error {
	sender := msg.Sender.AccAddress()
	lastPositionId := k.GetLastPositionId(ctx)

	positionKey := types.AddressPositionWithIdKeyPrefix(sender, lastPositionId+1)

	wrappedPosition := types.WrappedPosition{
		Id:       string(positionKey),
		Address:  msg.Sender,
		StartAt:  *timestamppb.New(time.Now()), // TODO
		Position: msg.Position,
	}

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

func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) error {
	positionId := types.GetPositionIdFromBytes([]byte(msg.PositionId))
	wrappedPosition := k.GetPosition(ctx, msg.Sender.AccAddress(), positionId)
	k.DeletePosition(ctx, msg.Sender.AccAddress(), positionId)

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
