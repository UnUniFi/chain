package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) DepositToTranche(goCtx context.Context, msg *types.MsgDepositToTranche) (*types.MsgDepositToTrancheResponse, error) {
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	return &types.MsgDepositToTrancheResponse{}, nil
}
