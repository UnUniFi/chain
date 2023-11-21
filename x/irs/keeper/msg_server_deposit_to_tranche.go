package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) DepositToTranche(goCtx context.Context, msg *types.MsgDepositToTranche) (*types.MsgDepositToTrancheResponse, error) {
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	if msg.TrancheType == types.TrancheType_NORMAL_YIELD { // Both PT and YT
		err := k.MintPtYtPair(ctx, sender, msg.StrategyContract, msg.TrancheMaturity, msg.Token)
		if err != nil {
			return nil, err
		}
	} else if msg.TrancheType == types.TrancheType_FIXED_YIELD {
		// Buy PT from AMM with msg.TrancheMaturity for msg.SpendAmount
		k.SwapUtToPt()
	} else if msg.TrancheType == types.TrancheType_LEVERAGED_VARIABLE_YIELD {
		// Borrow msg.AmountToBuy from AMM pool
		// MintPtYtPair
		// Sell msg.AmountToBuy worth of PT
		// Return borrowed amount
		k.SwapUtToYt()
	}

	return &types.MsgDepositToTrancheResponse{}, nil
}
