package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/copy-trading/types"

	derivativesTypes "github.com/UnUniFi/chain/x/derivatives/types"
)

// SetTracing set a specific tracing in the store from its index
func (k Keeper) SetTracing(ctx sdk.Context, tracing types.Tracing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TracingKeyPrefix))
	b := k.cdc.MustMarshal(&tracing)
	store.Set(types.TracingKey(
		tracing.Address,
	), b)
}

// GetTracing returns a tracing from its index
func (k Keeper) GetTracing(
	ctx sdk.Context,
	index string,

) (val types.Tracing, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TracingKeyPrefix))

	b := store.Get(types.TracingKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTracing removes a tracing from the store
func (k Keeper) RemoveTracing(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TracingKeyPrefix))
	store.Delete(types.TracingKey(
		index,
	))
}

// GetAllTracing returns all tracing
func (k Keeper) GetAllTracing(ctx sdk.Context) (list []types.Tracing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TracingKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tracing
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetExemplaryTraderTracing(ctx sdk.Context, exemplaryTrader string) (list []types.Tracing) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExemplaryTraderTracingKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tracing
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GenerateTracedPositionOpenMsg(ctx sdk.Context, tracing types.Tracing, position derivativesTypes.Position) derivativesTypes.MsgOpenPosition {
	tracerAddress := sdk.MustAccAddressFromBech32(tracing.Address)
	balance := k.bankKeeper.SpendableCoins(ctx, tracerAddress)

	unpackedPositionInstance, _ := derivativesTypes.UnpackPositionInstance(position.PositionInstance)

	switch positionInstance := unpackedPositionInstance.(type) {
	case *derivativesTypes.PerpetualFuturesPositionInstance:
		{
			tracedPositionSize := tracing.SizeCoefficient.Mul(positionInstance.Size_)
			tracedPositionLeverage := tracing.LeverageCoefficient.Mul(sdk.NewDec(int64(positionInstance.Leverage))).TruncateInt().Int64()
			tracedMargin := position.RemainingMargin // TODO

			tracedPositionInstance := derivativesTypes.PerpetualFuturesPositionInstance{
				PositionType: positionInstance.PositionType,
				Size_:        tracedPositionSize,
				Leverage:     uint32(tracedPositionLeverage),
			}

			tracedMsg := derivativesTypes.NewMsgOpenPosition(sdk.AccAddress(tracing.Address), tracedMargin, position.Market, tracedPositionInstance)

			return tracedMsg
		}
	case *derivativesTypes.PerpetualOptionsPositionInstance:
		{
			panic("not implemented")
		}
	}

	// TODO:
	// traced_position_size = tracing.SizeCoefficient * position_size
	// traced_position_leverage = tracing.LeverageCoefficient * position_leverage
	// if traced_position_margin > balance { adjust_traced_position_size to traced_position_margin = balance }
}

func (k Keeper) TracePosition(ctx sdk.Context, tracing types.Tracing, position derivativesTypes.Position) {
	tracedPositionMsg := k.GenerateTracedPositionOpenMsg(ctx, tracing, position)

	k.DerivativesKeeper.OpenPosition(ctx, &tracedPositionMsg)
}
