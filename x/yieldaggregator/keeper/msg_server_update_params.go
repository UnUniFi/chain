package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	k.SetParams(ctx, &msg.Params)

	return &types.MsgUpdateParamsResponse{}, nil
}
