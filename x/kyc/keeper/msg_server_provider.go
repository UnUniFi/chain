package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/UnUniFi/chain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateProvider(goCtx context.Context, msg *types.MsgCreateProvider) (*types.MsgCreateProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Sender != k.authority {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "sender is not the authority address")
	}

	var provider = types.Provider{
		Address:         msg.Sender,
		Name:            msg.Name,
		Identity:        msg.Identity,
		Website:         msg.Website,
		SecurityContact: msg.SecurityContact,
		Details:         msg.Details,
	}

	id := k.AppendProvider(
		ctx,
		provider,
	)

	return &types.MsgCreateProviderResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateProvider(goCtx context.Context, msg *types.MsgUpdateProvider) (*types.MsgUpdateProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	provider, found := k.GetProvider(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Sender != provider.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	provider.Address = msg.Address
	provider.Name = msg.Name
	provider.Identity = msg.Identity
	provider.Website = msg.Website
	provider.SecurityContact = msg.SecurityContact
	provider.Details = msg.Details

	k.SetProvider(ctx, provider)

	return &types.MsgUpdateProviderResponse{}, nil
}
