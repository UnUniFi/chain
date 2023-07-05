package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
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

func (k msgServer) DepositToPool(c context.Context, msg *types.MsgDepositToPool) (*types.MsgDepositToPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.MintLiquidityProviderToken(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgDepositToPoolResponse{}, nil
}

func (k msgServer) WithdrawFromPool(c context.Context, msg *types.MsgWithdrawFromPool) (*types.MsgWithdrawFromPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.BurnLiquidityProviderToken(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgWithdrawFromPoolResponse{}, nil
}

func (k msgServer) OpenPosition(c context.Context, msg *types.MsgOpenPosition) (*types.MsgOpenPositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.OpenPosition(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgOpenPositionResponse{}, nil
}

func (k msgServer) ClosePosition(c context.Context, msg *types.MsgClosePosition) (*types.MsgClosePositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.ClosePosition(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgClosePositionResponse{}, nil
}

func (k msgServer) ReportLiquidation(c context.Context, msg *types.MsgReportLiquidation) (*types.MsgReportLiquidationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.ReportLiquidation(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgReportLiquidationResponse{}, nil
}

func (k msgServer) ReportLevyPeriod(c context.Context, msg *types.MsgReportLevyPeriod) (*types.MsgReportLevyPeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.ReportLevyPeriod(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgReportLevyPeriodResponse{}, nil
}

func (k msgServer) AddMargin(c context.Context, msg *types.MsgAddMargin) (*types.MsgAddMarginResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.AddMargin(ctx, sdk.AccAddress(msg.Sender), msg.PositionId, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgAddMarginResponse{}, nil
}

func (k msgServer) RemoveMargin(c context.Context, msg *types.MsgRemoveMargin) (*types.MsgRemoveMarginResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.RemoveMargin(ctx, sdk.AccAddress(msg.Sender), msg.PositionId, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgRemoveMarginResponse{}, nil
}
