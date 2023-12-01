package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) DepositToTranche(goCtx context.Context, msg *types.MsgDepositToTranche) (*types.MsgDepositToTrancheResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	tranche, found := k.GetTranchePool(ctx, msg.TrancheId)
	if !found {
		return nil, types.ErrTrancheNotFound
	}

	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	if msg.TrancheType == types.TrancheType_NORMAL_YIELD { // Both PT and YT
		_, err := k.MintPtYtPair(ctx, sender, tranche, msg.Token)
		if err != nil {
			return nil, err
		}
	} else if msg.TrancheType == types.TrancheType_FIXED_YIELD {
		// Buy PT from AMM with msg.TrancheMaturity for msg.SpendAmount
		err := k.SwapUtToPt(ctx, sender, tranche, msg.Token)
		if err != nil {
			return nil, err
		}
	} else if msg.TrancheType == types.TrancheType_LEVERAGED_VARIABLE_YIELD {
		// Borrow msg.AmountToBuy from AMM pool
		// MintPtYtPair
		// Sell msg.AmountToBuy worth of PT
		// Return borrowed amount
		k.SwapUtToYt(ctx, sender, tranche, msg.RequiredYt)
	}

	return &types.MsgDepositToTrancheResponse{}, nil
}
