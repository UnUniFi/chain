package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetLastPositionId(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(types.KeyPrefixLastPositionId))

	return string(bz)
}

func (k Keeper) IncreaseLastPositionId(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(types.KeyPrefixLastPositionId))
	if bz == nil {
		store.Set([]byte(types.KeyPrefixLastPositionId), types.GetPositionIdBytes(0))
	}

	lastPositionId := types.GetPositionIdFromString(k.GetLastPositionId(ctx))
	store.Set([]byte(types.KeyPrefixLastPositionId), types.GetPositionIdBytes(lastPositionId+1))
}

func (k Keeper) GetAllPositions(ctx sdk.Context) []*types.Position {
	store := ctx.KVStore(k.storeKey)

	positions := []*types.Position{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixPosition))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.Position{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, &position)
	}

	return positions
}

func (k Keeper) GetPositionWithId(ctx sdk.Context, id string) *types.Position {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.PositionWithIdKeyPrefix(id))
	if bz == nil {
		return nil
	}
	position := types.Position{}
	k.cdc.MustUnmarshal(bz, &position)

	return &position
}

func (k Keeper) GetAddressPositions(ctx sdk.Context, user sdk.AccAddress) []*types.Position {
	store := ctx.KVStore(k.storeKey)

	positions := []*types.Position{}
	it := sdk.KVStorePrefixIterator(store, types.AddressPositionKeyPrefix(user))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.Position{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, &position)
	}

	return positions
}

func (k Keeper) GetAddressPositionWithId(ctx sdk.Context, address sdk.AccAddress, id string) *types.Position {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.AddressPositionWithIdKeyPrefix(address, id))
	if bz == nil {
		return nil
	}
	position := types.Position{}
	k.cdc.MustUnmarshal(bz, &position)

	return &position
}

func (k Keeper) CreatePosition(ctx sdk.Context, position types.Position) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&position)
	store.Set(types.PositionWithIdKeyPrefix(position.Id), bz)
	store.Set(types.AddressPositionWithIdKeyPrefix(position.Address.AccAddress(), position.Id), bz)
}

func (k Keeper) DeletePosition(ctx sdk.Context, address sdk.AccAddress, id string) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.PositionWithIdKeyPrefix(id))
	store.Delete(types.AddressPositionWithIdKeyPrefix(address, id))
}

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) error {
	sender := msg.Sender.AccAddress()
	lastPositionId := k.GetLastPositionId(ctx)

	positionKey := types.AddressPositionWithIdKeyPrefix(sender, lastPositionId)
	positionId := string(positionKey)

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.Sender.AccAddress(), types.ModuleName, sdk.NewCoins(msg.Margin))
	k.SetRemainingMargin(ctx, positionId, msg.Margin)

	positionInstance, err := types.UnpackPositionInstance(msg.PositionInstance)
	if err != nil {
		return err
	}

	var position *types.Position
	switch positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		position, err = k.OpenPerpetualFuturesPosition(ctx, positionId, msg.Sender, msg.Margin, msg.Market, *positionInstance.(*types.PerpetualFuturesPositionInstance))
	case *types.PerpetualOptionsPositionInstance:
		position, err = k.OpenPerpetualOptionsPosition(ctx, positionId, msg.Sender, msg.Margin, msg.Market, *positionInstance.(*types.PerpetualOptionsPositionInstance))
	default:
		panic("")
	}

	if err != nil {
		return err
	}

	k.CreatePosition(ctx, *position)
	k.IncreaseLastPositionId(ctx)

	return nil
}

func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) error {
	positionId := msg.PositionId
	position := k.GetAddressPositionWithId(ctx, msg.Sender.AccAddress(), positionId)

	if position == nil {
		return errors.New("position not found")
	}

	if msg.Sender.AccAddress().String() != position.Address.AccAddress().String() {
		return errors.New("not owner")
	}

	positionInstance, err := types.UnpackPositionInstance(position.PositionInstance)
	if err != nil {
		return err
	}

	switch positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		err = k.ClosePerpetualFuturesPosition(ctx, *position, *positionInstance.(*types.PerpetualFuturesPositionInstance))
		break
	case *types.PerpetualOptionsPositionInstance:
		err = k.ClosePerpetualOptionsPosition(ctx, *position, *positionInstance.(*types.PerpetualOptionsPositionInstance))
		break
	default:
		panic("")
	}

	if err != nil {
		return err
	}

	k.DeletePosition(ctx, msg.Sender.AccAddress(), positionId)

	return nil
}

func (k Keeper) ReportLiquidation(ctx sdk.Context, msg *types.MsgReportLiquidation) error {
	position := k.GetPositionWithId(ctx, msg.PositionId)

	if position == nil {
		return errors.New("position not found")
	}

	positionInstance, err := types.UnpackPositionInstance(position.PositionInstance)
	if err != nil {
		return err
	}

	remainingMargin := *k.GetRemainingMargin(ctx, msg.PositionId)

	switch positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		err = k.ReportLiquidationNeededPerpetualFuturesPosition(ctx, msg.RewardRecipient, remainingMargin, *position, *positionInstance.(*types.PerpetualFuturesPositionInstance))
		break
	case *types.PerpetualOptionsPositionInstance:
		err = k.ReportLiquidationNeededPerpetualOptionsPosition(ctx, msg.RewardRecipient, remainingMargin, *position, *positionInstance.(*types.PerpetualOptionsPositionInstance))
		break
	default:
		panic("")
	}

	if err != nil {
		return err
	}

	return nil
}
