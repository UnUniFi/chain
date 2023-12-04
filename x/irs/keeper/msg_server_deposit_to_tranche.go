package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) DepositToTranche(goCtx context.Context, msg *types.MsgDepositToTranche) (*types.MsgDepositToTrancheResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.DepositToTranchePool(ctx, sender, msg.TrancheId, msg.TrancheType, msg.Token, msg.RequiredYt)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositToTrancheResponse{}, nil
}
