package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/auction/keeper"
	"github.com/lcnem/jpyx/x/auction/types"
)

func handleMsgPlaceBid(ctx sdk.Context, k keeper.Keeper, msg types.MsgPlaceBid) (*sdk.Result, error) {

	err := k.PlaceBid(ctx, msg.AuctionId, msg.Bidder, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Bidder.String()),
		),
	)

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}
