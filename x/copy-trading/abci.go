package copy_trading

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/copy-trading/keeper"
	"github.com/UnUniFi/chain/x/copy-trading/types"

	derivativesTypes "github.com/UnUniFi/chain/x/derivatives/types"

	proto "github.com/gogo/protobuf/proto"
)

func openPositionHandler(ctx sdk.Context, k keeper.Keeper, event sdk.Event) {
	var sender, positionId string
	for _, attribute := range event.Attributes {
		if string(attribute.Key) == "sender" {
			sender = string(attribute.Value)
		}
		if string(attribute.Key) == "position_id" {
			positionId = string(attribute.Value)
		}
	}

	trader, found := k.GetExemplaryTrader(ctx, sender)
	if !found {
		return
	}

	tracings := k.GetExemplaryTraderTracing(ctx, trader.Address)
	position := k.DerivativesKeeper.GetPosition(ctx, positionId)
	for _, tracing := range tracings {
		k.TracePosition(ctx, tracing, position)
	}
}

func closePositionHandler(ctx sdk.Context, k keeper.Keeper, event sdk.Event) {
	var sender, positionId string
	for _, attribute := range event.Attributes {
		if string(attribute.Key) == "sender" {
			sender = string(attribute.Value)
		}
		if string(attribute.Key) == "position_id" {
			positionId = string(attribute.Value)
		}
	}

	// TODO: get position ids list of traced position and if position id is in list, then close it
	// k.DerivativesKeeper.ClosePosition(ctx, &derivativesTypes.MsgClosePosition{Sender: sender, PositionId: positionId})
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	for _, event := range ctx.EventManager().Events() {
		if event.Type == proto.MessageName(&derivativesTypes.EventPerpetualFuturesPositionOpened{}) {
			openPositionHandler(ctx, k, event)
		}
		if event.Type == proto.MessageName(&derivativesTypes.EventPerpetualFuturesPositionClosed{}) {
			closePositionHandler(ctx, k, event)
		}
	}
}
