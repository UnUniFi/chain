package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetLastPositionId(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(types.KeyPrefixLastPositionId))
	if bz == nil {
		panic("last position id not set in genesis")
	}

	return string(bz)
}

func (k Keeper) IncreaseLastPositionId(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)

	lastPositionId := types.GetPositionIdFromString(k.GetLastPositionId(ctx))
	store.Set([]byte(types.KeyPrefixLastPositionId), types.GetPositionIdBytes(lastPositionId+1))
}

func (k Keeper) GetAllOpenedPositions(ctx sdk.Context) []*types.OpenedPosition {
	store := ctx.KVStore(k.storeKey)

	positions := []*types.OpenedPosition{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixOpenedPosition))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.OpenedPosition{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, &position)
	}

	return positions
}

func (k Keeper) GetAddressOpenedPositions(ctx sdk.Context, user sdk.AccAddress) []*types.OpenedPosition {
	store := ctx.KVStore(k.storeKey)

	positions := []*types.OpenedPosition{}
	it := sdk.KVStorePrefixIterator(store, types.AddressOpenedPositionKeyPrefix(user))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.OpenedPosition{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, &position)
	}

	return positions
}

func (k Keeper) GetAddressClosedPositions(ctx sdk.Context, user sdk.AccAddress) []*types.ClosedPosition {
	store := ctx.KVStore(k.storeKey)

	positions := []*types.ClosedPosition{}
	it := sdk.KVStorePrefixIterator(store, types.AddressClosedPositionKeyPrefix(user))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.ClosedPosition{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, &position)
	}

	return positions
}

func (k Keeper) GetClosedPosition(ctx sdk.Context, positionId string) *types.ClosedPosition {
	store := ctx.KVStore(k.storeKey)

	position := types.ClosedPosition{}
	// TODO: implement this
	return &position
}

func (k Keeper) CreateOpenedPosition(ctx sdk.Context, OpenedPosition types.OpenedPosition) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&OpenedPosition)
	store.Set(types.AddressOpenedPositionWithIdKeyPrefix(OpenedPosition.Address.AccAddress(), OpenedPosition.Id), bz)
}

func (k Keeper) GetIdAddressOpenedPosition(ctx sdk.Context, address sdk.AccAddress, id string) types.OpenedPosition {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.AddressOpenedPositionWithIdKeyPrefix(address, id))
	position := types.OpenedPosition{}
	k.cdc.Unmarshal(bz, &position)

	return position
}

func (k Keeper) DeleteOpenedPosition(ctx sdk.Context, address sdk.AccAddress, id string) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.AddressOpenedPositionWithIdKeyPrefix(address, id))
}

func (k Keeper) CreateClosedPosition(ctx sdk.Context, closedPosition types.ClosedPosition) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&closedPosition)
	store.Set(types.AddressClosedPositionWithIdKeyPrefix(closedPosition.Address.AccAddress(), closedPosition.Id), bz)
}

func (k Keeper) SetRemainingMargin(ctx sdk.Context, positionId string, margin sdk.Coin) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&margin)
	store.Set(types.RemainingMarginKeyPrefix(positionId), bz)
}

func (k Keeper) GetRemainingMargin(ctx sdk.Context, positionId string) sdk.Coin {
	store := ctx.KVStore(k.storeKey)

	margin := sdk.Coin{}
	bz := store.Get(types.RemainingMarginKeyPrefix(positionId))
	k.cdc.MustUnmarshal(bz, &margin)

	return margin
}

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) error {
	sender := msg.Sender.AccAddress()
	lastPositionId := k.GetLastPositionId(ctx)

	positionKey := types.AddressOpenedPositionWithIdKeyPrefix(sender, lastPositionId)
	positionId := string(positionKey)

	// TODO: subtract margin from user's balance
	k.SetRemainingMargin(ctx, positionId, msg.Margin)

	// TODO: need to be refactored
	position, err := types.UnpackPosition(&msg.Position)
	if err != nil {
		return err
	}
	switch position.(type) {
	case *types.PerpetualFuturesPosition:
		return k.OpenPerpetualFuturesPosition(ctx, positionId, msg.Sender.AccAddress(), position.(*types.PerpetualFuturesPosition))
	case *types.PerpetualOptionsPosition:
		return k.OpenPerpetualOptionsPosition(ctx, positionId, msg.Sender.AccAddress(), position.(*types.PerpetualOptionsPosition))
	}

	k.IncreaseLastPositionId(ctx)

	return nil
}

func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) error {
	positionId := msg.PositionId
	openedPosition := k.GetIdAddressOpenedPosition(ctx, msg.Sender.AccAddress(), positionId)

	if msg.Sender.AccAddress().String() != openedPosition.Address.AccAddress().String() {
		return nil // TODO: return error
	}

	k.DeleteOpenedPosition(ctx, msg.Sender.AccAddress(), positionId)

	position, err := types.UnpackOpenedPosition(&openedPosition.Position)
	if err != nil {
		return err
	}
	switch position.(type) {
	case *types.PerpetualFuturesPosition:
		return k.ClosePerpetualFuturesPosition(ctx, openedPosition, position.(*types.PerpetualFuturesOpenedPosition))
	case *types.PerpetualOptionsPosition:
		return k.ClosePerpetualOptionsPosition(ctx, openedPosition, position.(*types.PerpetualOptionsOpenedPosition))
	}

	return nil
}
