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

	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	_, err = k.keeper.GetOracle(ctx, msg.MarketId, from)
	if err != nil {
		return nil, err
	}

	_, err = k.keeper.SetPrice(ctx, from, msg.MarketId, msg.Price, msg.Expiry)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
		),
	)

	// Gas-less
	ctx.GasMeter().RefundGas(ctx.GasMeter().GasConsumed(), "pricefeed: PostPrice")

	return &types.MsgPostPriceResponse{}, nil
}
