package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

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
	ctx := sdk.UnwrapSDKContext(c)

	subjectInfoList, err := k.keeper.Register(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&types.EventRegister{
		IncentiveUnitId:  msg.IncentiveUnitId,
		SubjectInfoLists: *subjectInfoList,
	}); err != nil {
		return nil, err
	}

	return &types.MsgRegisterResponse{}, nil
}

func (k msgServer) WithdrawAllRewards(c context.Context, msg *types.MsgWithdrawAllRewards) (*types.MsgWithdrawAllRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	rewards, err := k.keeper.WithdrawAllRewards(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&types.EventWithdrawAllRewards{
		Sender:              msg.Sender,
		AllWithdrawnRewards: rewards,
	}); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawAllRewardsResponse{}, nil
}

func (k msgServer) WithdrawReward(c context.Context, msg *types.MsgWithdrawReward) (*types.MsgWithdrawRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	reward, err := k.keeper.WithdrawReward(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&types.EventWithdrawReward{
		Sender:          msg.Sender,
		WithdrawnReward: reward,
	}); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawRewardResponse{}, nil
}
