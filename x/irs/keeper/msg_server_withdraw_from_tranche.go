package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) WithdrawFromTranche(goCtx context.Context, msg *types.MsgWithdrawFromTranche) (*types.MsgWithdrawFromTrancheResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.WithdrawFromTranchePool(ctx, sender, msg.TrancheId, msg.TrancheType, msg.Tokens, msg.RequiredRedeemAmount)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawFromTrancheResponse{}, nil
}
