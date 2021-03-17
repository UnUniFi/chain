package cdp

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/jpyx/x/cdp/keeper"
	"github.com/lcnem/jpyx/x/cdp/types"
)

func handleMsgCreateCdp(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateCdp) (*sdk.Result, error) {
	k.CreateCdp(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateCdp(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateCdp) (*sdk.Result, error) {
	var cdp = types.CDP{
		Creator: msg.Creator,
		Id:      msg.Id,
	}

	// Checks that the element exists
	if !k.HasCdp(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetCdpOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCdp(ctx, cdp)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteCdp(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteCdp) (*sdk.Result, error) {
	if !k.HasCdp(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetCdpOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.DeleteCdp(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
