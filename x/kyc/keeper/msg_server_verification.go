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

	// TODO: check provider Id

	k.SetVerification(
		ctx,
		verification,
	)
	return &types.MsgCreateVerificationResponse{}, nil
}

func (k msgServer) DeleteVerification(goCtx context.Context, msg *types.MsgDeleteVerification) (*types.MsgDeleteVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	verification, isFound := k.GetVerification(
		ctx,
		msg.Customer,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	verification.Address = msg.Customer
	verification.ProviderIds = []uint64{}

	// TODO: check provider Id

	for _, providerId := range verification.ProviderIds {
		if providerId != msg.ProviderId {
			verification.ProviderIds = append(verification.ProviderIds, providerId)
		}
	}

	if len(verification.ProviderIds) == 0 {
		k.RemoveVerification(
			ctx,
			msg.Customer,
		)
	} else {
		k.SetVerification(
			ctx,
			verification,
		)
	}

	return &types.MsgDeleteVerificationResponse{}, nil
}
