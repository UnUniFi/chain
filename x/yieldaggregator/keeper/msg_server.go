package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
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

func (k msgServer) Deposit(c context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.Deposit(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgDepositResponse{}, nil
}

func (k msgServer) Withdraw(c context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.Withdraw(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgWithdrawResponse{}, nil
}

func (k msgServer) AddFarmingOrder(c context.Context, msg *types.MsgAddFarmingOrder) (*types.MsgAddFarmingOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.AddFarmingOrder(ctx, *msg.Order)
	if err != nil {
		return nil, err
	}
	return &types.MsgAddFarmingOrderResponse{}, nil
}

func (k msgServer) DeleteFarmingOrder(c context.Context, msg *types.MsgDeleteFarmingOrder) (*types.MsgDeleteFarmingOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	k.Keeper.DeleteFarmingOrder(ctx, msg.FromAddress.AccAddress(), msg.OrderId)
	return &types.MsgDeleteFarmingOrderResponse{}, nil
}

func (k msgServer) ActivateFarmingOrder(c context.Context, msg *types.MsgActivateFarmingOrder) (*types.MsgActivateFarmingOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	k.Keeper.ActivateFarmingOrder(ctx, msg.FromAddress.AccAddress(), msg.OrderId)
	return &types.MsgActivateFarmingOrderResponse{}, nil
}

func (k msgServer) InactivateFarmingOrder(c context.Context, msg *types.MsgInactivateFarmingOrder) (*types.MsgInactivateFarmingOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	k.Keeper.InactivateFarmingOrder(ctx, msg.FromAddress.AccAddress(), msg.OrderId)
	return &types.MsgInactivateFarmingOrderResponse{}, nil
}

func (k msgServer) ExecuteFarmingOrders(c context.Context, msg *types.MsgExecuteFarmingOrders) (*types.MsgExecuteFarmingOrdersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	k.Keeper.ExecuteFarmingOrders(ctx, msg.FromAddress.AccAddress())
	return &types.MsgExecuteFarmingOrdersResponse{}, nil
}
