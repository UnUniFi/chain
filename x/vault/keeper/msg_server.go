package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/vault/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) ParamsUpdate(c context.Context, msg *types.MsgParamsUpdate) (*types.MsgParamsUpdateResponse, error) {
	// to do
	return &types.MsgParamsUpdateResponse{}, nil
}
