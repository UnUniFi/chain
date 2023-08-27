package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateVerification(goCtx context.Context, msg *types.MsgCreateVerification) (*types.MsgCreateVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetVerification(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var verification = types.Verification{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetVerification(
		ctx,
		verification,
	)
	return &types.MsgCreateVerificationResponse{}, nil
}

func (k msgServer) UpdateVerification(goCtx context.Context, msg *types.MsgUpdateVerification) (*types.MsgUpdateVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetVerification(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var verification = types.Verification{
		Creator: msg.Creator,
		Index:   msg.Index,
	}

	k.SetVerification(ctx, verification)

	return &types.MsgUpdateVerificationResponse{}, nil
}

func (k msgServer) DeleteVerification(goCtx context.Context, msg *types.MsgDeleteVerification) (*types.MsgDeleteVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetVerification(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveVerification(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteVerificationResponse{}, nil
}
