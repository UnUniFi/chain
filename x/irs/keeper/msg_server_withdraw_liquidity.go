package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) WithdrawLiquidity(goCtx context.Context, msg *types.MsgWithdrawLiquidity) (*types.MsgWithdrawLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	return &types.MsgWithdrawLiquidityResponse{}, nil
}
