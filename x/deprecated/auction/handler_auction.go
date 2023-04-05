package auction

import (
	"github.com/UnUniFi/chain/x/auction/keeper"
	"github.com/UnUniFi/chain/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgPlaceBid(ctx sdk.Context, k keeper.Keeper, msg types.MsgPlaceBid) (*sdk.Result, error) {

	err := k.PlaceBid(ctx, msg.AuctionId, msg.Bidder.AccAddress(), msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Bidder.AccAddress().String()),
		),
	)

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}
