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
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	tranche, found := k.GetTranchePool(ctx, msg.TrancheId)
	if !found {
		return nil, types.ErrTrancheNotFound
	}

	if msg.TrancheType == types.TrancheType_FIXED_YIELD {
		// If matured, send required amount from unbonded from the share
		if tranche.StartTime+tranche.Maturity <= uint64(ctx.BlockTime().Unix()) {
			err := k.RedeemPtAtMaturity(ctx, sender, tranche, msg.Token)
			if err != nil {
				return nil, err
			}
		}

		// Else, sell PT from AMM with msg.TrancheMaturity for msg.PTAmount
		err := k.SwapPtToUt(ctx, sender, tranche, msg.Token)
		if err != nil {
			return nil, err
		}
	} else if msg.TrancheType == types.TrancheType_LEVERAGED_VARIABLE_YIELD {
		// If matured, send required amount from unbonded from the share
		if tranche.StartTime+tranche.Maturity <= uint64(ctx.BlockTime().Unix()) {
			err := k.RedeemYtAtMaturity(ctx, sender, tranche, msg.Token)
			if err != nil {
				return nil, err
			}
		}
		// Else
		panic("not allowed to withdraw yt before being matured")
		// Put required amount of msg.PT from user wallet
		// Close position
		// Start redemption for strategy share
		// k.SwapYtToUt()
	} else { // All yield
		panic("Not possible to withdraw both tokens")
	}

	return &types.MsgWithdrawFromTrancheResponse{}, nil
}
