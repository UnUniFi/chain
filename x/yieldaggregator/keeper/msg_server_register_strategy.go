package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) RegisterStrategy(goCtx context.Context, msg *types.MsgRegisterStrategy) (*types.MsgRegisterStrategyResponse, error) {
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	k.AppendStrategy(ctx, msg.Denom, types.Strategy{
		Denom:           msg.Denom,
		ContractAddress: msg.ContractAddress,
		Name:            msg.Name,
		GitUrl:          msg.GitUrl,
	})

	return &types.MsgRegisterStrategyResponse{}, nil
}
