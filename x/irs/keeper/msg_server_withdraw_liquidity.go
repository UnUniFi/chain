package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) WithdrawLiquidity(goCtx context.Context, msg *types.MsgWithdrawLiquidity) (*types.MsgWithdrawLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Burn lp tokens and get tokens from tranche pool for PT + Deposit-Token
	_, err = k.WithdrawFromLiquidityPool(ctx, sender, msg.TrancheId, msg.ShareAmount, msg.TokenOutMins)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawLiquidityResponse{}, nil
}
