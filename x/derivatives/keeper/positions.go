package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/UnUniFi/chain/x/derivatives/types"
	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
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

	lastPositionId := k.GetLastPositionId(ctx)
	store.Set([]byte(types.KeyPrefixLastPositionId), []byte(lastPositionId+1))
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

func (k Keeper) SatRemainingMargin(ctx sdk.Context, positionId string, margin sdk.Coin) {
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
	k.SatRemainingMargin(ctx, positionId, msg.Margin)

	// TODO: need to be refactored
	wrappedPosition := types.OpenedPosition{
		Id:       string(positionKey),
		Address:  msg.Sender,
		OpenedAt: *timestamppb.New(ctx.BlockTime()),
		Position: nil,
	}

	k.CreateOpenedPosition(ctx, wrappedPosition)

	// TODO: need to be refactored
	position, err := types.UnpackOpenedPosition(&msg.Position)
	if err != nil {
		return err
	}
	switch position.(type) {
	case *types.PerpetualFuturesPosition:
		return k.OpenPerpetualFuturesPosition(ctx, msg.Sender.AccAddress(), lastPositionId, position.(*types.PerpetualFuturesPosition))
	case *types.PerpetualOptionsPosition:
		return k.OpenPerpetualOptionsPosition(ctx, msg.Sender.AccAddress(), position.(*types.PerpetualOptionsPosition))
	}

	k.IncreaseLastPositionId(ctx)

	return nil
}

func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) error {
	positionId := msg.PositionId
	OpenedPosition := k.GetIdAddressOpenedPosition(ctx, msg.Sender.AccAddress(), positionId)

	if msg.Sender.AccAddress().String() != OpenedPosition.Address.AccAddress().String() {
		return nil // TODO: return error
	}

	k.DeleteOpenedPosition(ctx, msg.Sender.AccAddress(), positionId)

	closedPosition := types.ClosedPosition{
		Id: positionId,
		// TODO:
	}

	k.CreateClosedPosition(ctx, closedPosition)

	// TODO: need to be refactored
	position, err := types.UnpackOpenedPosition(&OpenedPosition.Position)
	if err != nil {
		return err
	}
	switch position.(type) {
	case *types.PerpetualFuturesPosition:
		return k.ClosePerpetualFuturesPosition(ctx, closedPosition, position.(*types.PerpetualFuturesClosedPosition))
	case *types.PerpetualOptionsPosition:
		return k.ClosePerpetualOptionsPosition(ctx, closedPosition, position.(*types.PerpetualOptionsClosedPosition))
	}

	return nil
}
