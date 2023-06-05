package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/pricefeed/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) PostPrice(c context.Context, msg *types.MsgPostPrice) (*types.MsgPostPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.ValidateAuthorityAndDeposit(ctx, msg.MarketId, msg.From.AccAddress(), msg.Deposit)
	if err != nil {
		return nil, err
	}

	// TODO: slash deposit if the oracle is malicious
	_, err = k.keeper.SetPrice(ctx, msg.From.AccAddress(), msg.MarketId, msg.Price.ToSDKDec(), msg.Expiry)
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

	return &types.MsgPostPriceResponse{}, nil
}
