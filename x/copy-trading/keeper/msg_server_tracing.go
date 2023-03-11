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
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var tracing = types.Tracing{
		Address: msg.Creator,
	}

	k.SetTracing(
		ctx,
		tracing,
	)
	return &types.MsgCreateTracingResponse{}, nil
}

func (k msgServer) UpdateTracing(goCtx context.Context, msg *types.MsgUpdateTracing) (*types.MsgUpdateTracingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTracing(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Address {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var tracing = types.Tracing{
		Address: msg.Creator,
	}

	k.SetTracing(ctx, tracing)

	return &types.MsgUpdateTracingResponse{}, nil
}

func (k msgServer) DeleteTracing(goCtx context.Context, msg *types.MsgDeleteTracing) (*types.MsgDeleteTracingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTracing(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Address {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTracing(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteTracingResponse{}, nil
}
