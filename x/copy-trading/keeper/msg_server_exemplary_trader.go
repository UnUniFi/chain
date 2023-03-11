package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/copy-trading/types"
)

func (k msgServer) CreateExemplaryTrader(goCtx context.Context, msg *types.MsgCreateExemplaryTrader) (*types.MsgCreateExemplaryTraderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetExemplaryTrader(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var exemplaryTrader = types.ExemplaryTrader{
		Address: msg.Creator,
	}

	k.SetExemplaryTrader(
		ctx,
		exemplaryTrader,
	)
	return &types.MsgCreateExemplaryTraderResponse{}, nil
}

func (k msgServer) UpdateExemplaryTrader(goCtx context.Context, msg *types.MsgUpdateExemplaryTrader) (*types.MsgUpdateExemplaryTraderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetExemplaryTrader(
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

	var exemplaryTrader = types.ExemplaryTrader{
		Address: msg.Creator,
	}

	k.SetExemplaryTrader(ctx, exemplaryTrader)

	return &types.MsgUpdateExemplaryTraderResponse{}, nil
}

func (k msgServer) DeleteExemplaryTrader(goCtx context.Context, msg *types.MsgDeleteExemplaryTrader) (*types.MsgDeleteExemplaryTraderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetExemplaryTrader(
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

	k.RemoveExemplaryTrader(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteExemplaryTraderResponse{}, nil
}
