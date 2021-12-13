package cdp

import (
	"fmt"

	"github.com/UnUniFi/chain/x/cdp/keeper"
	"github.com/UnUniFi/chain/x/cdp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgCreateCdp:
			return handleMsgCreateCdp(ctx, k, msg)

		case *types.MsgDeposit:
			return handleMsgDeposit(ctx, k, msg)

		case *types.MsgWithdraw:
			return handleMsgWithdraw(ctx, k, msg)

		case *types.MsgDrawDebt:
			return handleMsgDrawDebt(ctx, k, msg)

		case *types.MsgRepayDebt:
			return handleMsgRepayDebt(ctx, k, msg)

		case *types.MsgLiquidate:
			return handleMsgLiquidate(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
