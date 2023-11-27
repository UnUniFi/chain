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

	// Put tokens on tranche pool and get lp token
	_, _, err = k.DepositToLiquidityPool(ctx, sender, msg.TrancheId, sdk.ZeroInt(), msg.TokenInMaxs)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositLiquidityResponse{}, nil
}
