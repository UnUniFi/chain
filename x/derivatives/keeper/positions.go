package keeper

import (
	"encoding/binary"
	"errors"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetLastPositionId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(types.KeyPrefixLastPositionId))
	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
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

	lastPositionId := k.GetLastPositionId(ctx)
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

func (k Keeper) GetAddressPositionsVal(ctx sdk.Context, user sdk.AccAddress) []types.Position {
	store := ctx.KVStore(k.storeKey)

	positions := []types.Position{}
	it := sdk.KVStorePrefixIterator(store, types.AddressPositionKeyPrefix(user))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.Position{}
		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, position)
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

func (k Keeper) SetPosition(ctx sdk.Context, position types.Position) error {
	addr, err := sdk.AccAddressFromBech32(position.Address)
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&position)
	store.Set(types.PositionWithIdKeyPrefix(position.Id), bz)
	store.Set(types.AddressPositionWithIdKeyPrefix(addr, position.Id), bz)

	return nil
}

func (k Keeper) DeletePosition(ctx sdk.Context, address sdk.AccAddress, id string) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.PositionWithIdKeyPrefix(id))
	store.Delete(types.AddressPositionWithIdKeyPrefix(address, id))
}

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) error {
	// check sender amount for margin
	if !k.IsAssetAcceptable(ctx, msg.Margin.Denom) {
		return errors.New("margin denom is not acceptable")
	}

	newPositionId := strconv.FormatUint(k.GetLastPositionId(ctx)+1, 10)

	// fixme check first bank.send last
	positionInstance, err := types.UnpackPositionInstance(msg.PositionInstance)
	if err != nil {
		return err
	}

	var position *types.Position
	switch positionInstance := positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		position, err = k.OpenPerpetualFuturesPosition(ctx, newPositionId, msg.Sender, msg.Margin, msg.Market, *positionInstance)
	case *types.PerpetualOptionsPositionInstance:
		position, err = k.OpenPerpetualOptionsPosition(ctx, newPositionId, msg.Sender, msg.Margin, msg.Market, *positionInstance)
	default:
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "position instance: %s", positionInstance)
	}

	if err != nil {
		return err
	}

	// mint position nft
	err = k.MintPositionNFT(ctx, *position)
	if err != nil {
		return err
	}

	err = k.SetPosition(ctx, *position)
	if err != nil {
		return err
	}
	k.IncreaseLastPositionId(ctx)

	return nil
}

func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) error {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	positionId := msg.PositionId
	position := k.GetPositionWithId(ctx, positionId)

	if position == nil {
		return types.ErrPositionDoesNotExist
	}

	// todo:  add pending position

	// if msg.Sender != position.Address {
	// 	return errors.New("not owner")
	// }

	// check withdrawer has owner nft
	owner := k.GetPositionNFTOwner(ctx, positionId)
	if owner.String() != msg.Sender {
		return types.ErrNotPositionNFTOwner
	}
	sendDisabled, err := k.GetPositionNFTSendDisabled(ctx, positionId)
	if err != nil {
		return err
	}
	if sendDisabled {
		return types.ErrPositionNFTSendDisabled
	}

	positionInstance, err := types.UnpackPositionInstance(position.PositionInstance)
	if err != nil {
		return err
	}

	switch positionInstance := positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := types.NewPerpetualFuturesPosition(*position, *positionInstance)
		err = k.ClosePerpetualFuturesPosition(ctx, perpetualFuturesPosition)
	case *types.PerpetualOptionsPositionInstance:
		err = k.ClosePerpetualOptionsPosition(ctx, *position, *positionInstance)
	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "position instance: %s", positionInstance)
	}

	if err != nil {
		return err
	}

	k.DeletePosition(ctx, sender, positionId)

	// delete position nft
	err = k.ClosePositionNFT(ctx, *position)
	if err != nil {
		return err
	}

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

	params := k.GetParams(ctx)

	currentBaseUsdRate, currentQuoteUsdRate, err := k.GetPairUsdPriceFromMarket(ctx, position.Market)
	if err != nil {
		return err
	}
	quoteTicker := k.GetPoolQuoteTicker(ctx)
	baseMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.BaseDenom, currentBaseUsdRate)
	quoteMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.QuoteDenom, currentQuoteUsdRate)

	switch positionInstance := positionInstance.(type) {
	case *types.PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := types.NewPerpetualFuturesPosition(*position, *positionInstance)
		if position.NeedLiquidation(params.PerpetualFutures.MarginMaintenanceRate, baseMetricsRate, quoteMetricsRate) {
			err = k.LiquidateFuturesPosition(ctx, msg.RewardRecipient, perpetualFuturesPosition, params.PerpetualFutures.CommissionRate, params.PoolParams.ReportLiquidationRewardRate)
		} else {
			return types.ErrLiquidationNotNeeded
		}
		break
	case *types.PerpetualOptionsPositionInstance:
		err = k.ReportLiquidationNeededPerpetualOptionsPosition(ctx, msg.RewardRecipient, *position, *positionInstance)
		break
	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "position instance: %s", positionInstance)
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

	params := k.GetParams(ctx)
	if position.LastLeviedAt.Add(time.Duration(params.PoolParams.LevyPeriodRequiredSeconds) * time.Second).After(ctx.BlockTime()) {
		return errors.New("levy period is allowed after the time set by params")
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "position instance: %s", positionInstance)
	}

	if err != nil {
		return err
	}

	return nil
}

// UnmarshalPosition unmarshals a position from a store value
func (k Keeper) UnmarshalPosition(bz []byte) (position types.Position, err error) {
	err = k.cdc.UnmarshalInterface(bz, &position)
	return position, err
}
