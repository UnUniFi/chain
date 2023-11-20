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

	// TODO:
	if msg.TrancheType == types.TrancheType_FIXED_YIELD {
		// If matured, send required amount from unbonded from the share
		k.RedeemPtAtMaturity()
		// Else, well PT from AMM with msg.TrancheMaturity for msg.PTAmount
		k.SwapPtToUt()
	} else if msg.TrancheType == types.TrancheType_LEVERAGED_VARIABLE_YIELD {
		// If matured, send required amount from unbonded from the share
		k.RedeemYtAtMaturity()
		// Else
		// Put required amount of msg.PT from user wallet
		// Close position
		// Start redemption for strategy share
		k.SwapYtToUt()
	} else { // All yield
		panic("Not possible to withdraw both tokens")
	}

	return &types.MsgWithdrawFromTrancheResponse{}, nil
}
