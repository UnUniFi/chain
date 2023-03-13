package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/copy-trading/types"
)

func (k msgServer) CreateTracing(goCtx context.Context, msg *types.MsgCreateTracing) (*types.MsgCreateTracingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTracing(
		ctx,
		msg.Sender,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var tracing = types.Tracing{
		Address: msg.Sender,
	}

	k.SetTracing(
		ctx,
		tracing,
	)
	return &types.MsgCreateTracingResponse{}, nil
}

func (k msgServer) DeleteTracing(goCtx context.Context, msg *types.MsgDeleteTracing) (*types.MsgDeleteTracingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTracing(
		ctx,
		msg.Sender,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Sender != valFound.Address {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTracing(
		ctx,
		msg.Sender,
	)

	return &types.MsgDeleteTracingResponse{}, nil
}
