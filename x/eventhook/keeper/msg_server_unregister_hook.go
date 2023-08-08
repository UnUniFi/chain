package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/eventhook/types"
)

func (k msgServer) UnregisterHook(goCtx context.Context, msg *types.MsgUnregisterHook) (*types.MsgUnregisterHookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Sender != k.authority {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "unauthorized")
	}

	_, found := k.GetHook(ctx, msg.EventType, msg.Id)

	if !found {
		return nil, sdkerrors.Wrap(types.ErrHookNotFound, "hook not found")
	}

	k.RemoveHook(ctx, msg.EventType, msg.Id)

	return &types.MsgUnregisterHookResponse{}, nil
}
