package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) UpdateStrategy(goCtx context.Context, msg *types.MsgUpdateStrategy) (*types.MsgUpdateStrategyResponse, error) {
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	strategy, found := k.GetStrategy(ctx, msg.Denom, msg.Id)
	if !found {
		return nil, types.ErrStrategyNotFound
	}

	strategy.Name = msg.Name
	strategy.Description = msg.Description
	strategy.GitUrl = msg.GitUrl

	return &types.MsgUpdateStrategyResponse{}, nil
}
