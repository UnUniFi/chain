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
	_, found := k.GetVerification(
		ctx,
		msg.Customer,
		msg.ProviderId,
	)

	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// TODO: check provider Id

	k.SetVerification(
		ctx,
		types.Verification{
			Address:    msg.Customer,
			ProviderId: msg.ProviderId,
		},
	)
	return &types.MsgCreateVerificationResponse{}, nil
}

func (k msgServer) DeleteVerification(goCtx context.Context, msg *types.MsgDeleteVerification) (*types.MsgDeleteVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetVerification(
		ctx,
		msg.Customer,
		msg.ProviderId,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// TODO: check provider Id

	k.RemoveVerification(
		ctx,
		msg.Customer,
		msg.ProviderId,
	)

	return &types.MsgDeleteVerificationResponse{}, nil
}
