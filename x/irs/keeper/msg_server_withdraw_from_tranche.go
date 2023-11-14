package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) WithdrawFromTranche(goCtx context.Context, msg *types.MsgWithdrawFromTranche) (*types.MsgWithdrawFromTrancheResponse, error) {
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	return &types.MsgWithdrawFromTrancheResponse{}, nil
}
