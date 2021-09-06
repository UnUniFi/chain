package incentive

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/jpyx/x/incentive/keeper"
	"github.com/lcnem/jpyx/x/incentive/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgClaimCdpMintingReward:
			return handleMsgClaimCdpMintingReward(ctx, k, msg)
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgClaimCdpMintingReward(ctx sdk.Context, k keeper.Keeper, msg *types.MsgClaimCdpMintingReward) (*sdk.Result, error) {

	err := k.ClaimCdpMintingReward(ctx, msg.Sender.AccAddress(), msg.MultiplierName)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}
