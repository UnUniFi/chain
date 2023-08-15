package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/eventhook/types"
)

func (k msgServer) RegisterHook(goCtx context.Context, msg *types.MsgRegisterHook) (*types.MsgRegisterHookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Sender != k.authority {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "unauthorized")
	}

	hook := types.Hook{
		Name:            msg.Name,
		ContractAddress: msg.ContractAddress,
		GitUrl:          msg.GitUrl,
		EventType:       msg.EventType,
		EventAttributes: msg.EventAttributes,
	}
	id := k.AppendHook(ctx, msg.EventType, hook)

	return &types.MsgRegisterHookResponse{
		Id: id,
	}, nil
}
