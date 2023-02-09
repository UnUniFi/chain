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

func (k msgServer) MintLiquidityProviderToken(c context.Context, msg *types.MsgMintLiquidityProviderToken) (*types.MsgMintLiquidityProviderTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.MintLiquidityProviderToken(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgMintLiquidityProviderTokenResponse{}, nil
}

func (k msgServer) BurnLiquidityProviderToken(c context.Context, msg *types.MsgBurnLiquidityProviderToken) (*types.MsgBurnLiquidityProviderTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.BurnLiquidityProviderToken(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgBurnLiquidityProviderTokenResponse{}, nil
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
