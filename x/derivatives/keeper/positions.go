package keeper

import (
	"errors"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetLastPositionId(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(types.KeyPrefixLastPositionId))

	return string(bz)
}

func (k Keeper) GetLastPosition(ctx sdk.Context) types.Position {
	store := ctx.KVStore(k.storeKey)

	position := types.Position{}

	it := sdk.KVStoreReversePrefixIterator(store, []byte(types.KeyPrefixPosition))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.Position{}
		k.cdc.Unmarshal(it.Value(), &position)
		return position
	}

	return position
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

func (k Keeper) GetAllPositions(ctx sdk.Context) []types.Position {
	store := ctx.KVStore(k.storeKey)

	positions := []types.Position{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixPosition))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.Position{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, position)
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

func (k Keeper) SetPosition(ctx sdk.Context, position types.Position) {
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
	// todo check sender amount for margin

	lastPosition := k.GetLastPosition(ctx)

	var positionId string
	if lastPosition.Id == "" {
		positionId = "0"
	} else {
		// increment position id
		lastPositionId, _ := strconv.Atoi(lastPosition.Id)
		positionId = strconv.Itoa(lastPositionId + 1)
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.Sender.AccAddress(), types.ModuleName, sdk.NewCoins(msg.Margin)); err != nil {
		return err
	}

	// fixme check first bank.send last
	positionInstance, err := types.UnpackPositionInstance(msg.PositionInstance)
	if err != nil {
		return err
	}

	var position *types.Position
	switch positionInstance := positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		position, err = k.OpenPerpetualFuturesPosition(ctx, positionId, msg.Sender, msg.Margin, msg.Market, *positionInstance)
	case *types.PerpetualOptionsPositionInstance:
		position, err = k.OpenPerpetualOptionsPosition(ctx, positionId, msg.Sender, msg.Margin, msg.Market, *positionInstance)
	default:
		panic("")
	}

	if err != nil {
		return err
	}

	k.SetPosition(ctx, *position)
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

	switch positionInstance := positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := types.NewPerpetualFuturesPosition(*position, *positionInstance)
		err = k.ClosePerpetualFuturesPosition(ctx, perpetualFuturesPosition)
		break
	case *types.PerpetualOptionsPositionInstance:
		err = k.ClosePerpetualOptionsPosition(ctx, *position, *positionInstance)
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

	switch positionInstance := positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := types.NewPerpetualFuturesPosition(*position, *positionInstance)
		err = k.ReportLiquidationNeededPerpetualFuturesPosition(ctx, msg.RewardRecipient, perpetualFuturesPosition)
		break
	case *types.PerpetualOptionsPositionInstance:
		err = k.ReportLiquidationNeededPerpetualOptionsPosition(ctx, msg.RewardRecipient, *position, *positionInstance)
		break
	default:
		panic("")
	}

	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) ReportLevyPeriod(ctx sdk.Context, msg *types.MsgReportLevyPeriod) error {
	position := k.GetPositionWithId(ctx, msg.PositionId)

	if position == nil {
		return errors.New("position not found")
	}

	if ctx.BlockTime().Sub(position.LastLeviedAt) < time.Duration(8)*time.Hour {
		return errors.New("It hasn't passed 8 hours since last levy")
	}

	positionInstance, err := types.UnpackPositionInstance(position.PositionInstance)
	if err != nil {
		return err
	}

	switch positionInstance := positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		err = k.ReportLevyPeriodPerpetualFuturesPosition(ctx, msg.RewardRecipient, *position, *positionInstance)
		break
	case *types.PerpetualOptionsPositionInstance:
		// err = k.ReportLevyPeriodPerpetualOptionsPosition(ctx, msg.RewardRecipient, *position, *positionInstance)
		break
	default:
		panic("")
	}

	if err != nil {
		return err
	}

	return nil
}
