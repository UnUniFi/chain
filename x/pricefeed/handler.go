package pricefeed

import (
	"fmt"

	"github.com/UnUniFi/chain/x/pricefeed/keeper"
	"github.com/UnUniFi/chain/x/pricefeed/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgPostPrice:
			return HandleMsgPostPrice(ctx, k, msg)
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// HandleMsgPostPrice handles prices posted by oracles
func HandleMsgPostPrice(
	ctx sdk.Context,
	k keeper.Keeper,
	msg *types.MsgPostPrice) (*sdk.Result, error) {
	err := k.ValidateAuthorityAndDeposit(ctx, msg.MarketId, msg.From.AccAddress(), msg.Deposit)
	if err != nil {
		return nil, err
	}
	_, err = k.SetPrice(ctx, msg.From.AccAddress(), msg.MarketId, msg.Price.ToSDKDec(), msg.Expiry)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From.AccAddress().String()),
		),
	)

	// Gas-less
	ctx.GasMeter().RefundGas(ctx.GasMeter().GasConsumed(), "pricefeed: PostPrice")

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
