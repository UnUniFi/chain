package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/UnUniFi/chain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateVerification(goCtx context.Context, msg *types.MsgCreateVerification) (*types.MsgCreateVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	verification, _ := k.GetVerification(
		ctx,
		msg.Customer,
	)

	verification.Address = msg.Customer
	verification.ProviderIds = append(verification.ProviderIds, msg.ProviderId)

	k.SetVerification(
		ctx,
		verification,
	)
	return &types.MsgCreateVerificationResponse{}, nil
}

func (k msgServer) UpdateVerification(goCtx context.Context, msg *types.MsgUpdateVerification) (*types.MsgUpdateVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetVerification(
		ctx,
		msg.Customer,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	panic("TODO: implement check")
	// if msg.Creator != valFound.Creator {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	var verification = types.Verification{
		Address: msg.Customer,
	}

	k.SetVerification(ctx, verification)

	return &types.MsgUpdateVerificationResponse{}, nil
}

func (k msgServer) DeleteVerification(goCtx context.Context, msg *types.MsgDeleteVerification) (*types.MsgDeleteVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetVerification(
		ctx,
		msg.Customer,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	panic("TODO: implement check")
	// if msg.Creator != valFound.Creator {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	k.RemoveVerification(
		ctx,
		msg.Customer,
	)

	return &types.MsgDeleteVerificationResponse{}, nil
}
