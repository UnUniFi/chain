package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/derivatives/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) OpenPosition(c context.Context, msg *types.MsgOpenPosition) (*types.MsgOpenPositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.OpenPosition(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgOpenPositionResponse{}, nil
}
