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

// TODO: rename MsgClaim to MsgClaimLiquidityProviderRewards
func (k msgServer) Claim(c context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.Keeper.Claim(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgClaimResponse{}, nil
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
