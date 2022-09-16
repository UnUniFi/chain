package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
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

func (k msgServer) Register(c context.Context, msg *types.MsgRegister) (*types.MsgRegisterResponse, error) {
	return &types.MsgRegisterResponse{}, nil
}

func (k msgServer) WithdrawAllRewards(c context.Context, msg *types.MsgWithdrawAllRewards) (*types.MsgWithdrawAllRewardsResponse, error) {
	return &types.MsgWithdrawAllRewardsResponse{}, nil
}

func (k msgServer) WithdrawReward(c context.Context, msg *types.MsgWithdrawReward) (*types.MsgWithdrawRewardResponse, error) {
	return &types.MsgWithdrawRewardResponse{}, nil
}
