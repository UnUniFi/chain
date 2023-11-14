package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) DepositLiquidity(goCtx context.Context, msg *types.MsgDepositLiquidity) (*types.MsgDepositLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	_ = sender
	_ = ctx

	return &types.MsgDepositLiquidityResponse{}, nil
}
