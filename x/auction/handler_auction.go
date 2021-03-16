package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/auction/keeper"
	"github.com/lcnem/jpyx/x/auction/types"
)

func handleMsgPlaceBid(ctx sdk.Context, k keeper.Keeper, msg *types.MsgPlaceBid) (*sdk.Result, error) {
	k.PlaceBid(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
