package cdp

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/jpyx/x/cdp/keeper"
	"github.com/lcnem/jpyx/x/cdp/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgCreateCDP:
			return handleMsgCreateCdp(ctx, k, msg)

		case *types.MsgDeposit:
			return handleMsgUpdateCdp(ctx, k, msg)

		case *types.MsgWithdraw:
			return handleMsgDeleteCdp(ctx, k, msg)

		case *types.MsgDrawDebt:
			return handleMsgDeleteCdp(ctx, k, msg)

		case *types.MsgRelayDebt:
			return handleMsgDeleteCdp(ctx, k, msg)

		case *types.MsgLiquidate:
			return handleMsgDeleteCdp(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
