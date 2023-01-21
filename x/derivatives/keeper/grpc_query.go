package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

var _ types.QueryServer = Keeper{}

// TODO: remove
func (k Keeper) LiquidityProviderRewardsSinceLastRedemption(c context.Context, req *types.QueryLiquidityProviderRewardsSinceLastRedemptionRequest) (*types.QueryLiquidityProviderRewardsSinceLastRedemptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	totalSupply := k.bankKeeper.GetSupply(ctx, "DLP")
	user := sdk.AccAddress(req.Address)
	userBalance := k.bankKeeper.GetBalance(ctx, user, "DLP")
	// TODO: not sure about the fee payout token - is it DLP or real asset?
	accumulatedFee := k.GetAccumulatedFee(ctx)

	tempAmount := accumulatedFee.Amount.Mul(userBalance.Amount)
	feeAmount := tempAmount.BigInt().Div(tempAmount.BigInt(), totalSupply.Amount.BigInt())

	return &types.QueryLiquidityProviderRewardsSinceLastRedemptionResponse{
		Amount: sdk.Coins{sdk.NewCoin(accumulatedFee.Denom, sdk.NewInt(feeAmount.Int64()))},
	}, nil
}

func (k Keeper) LiquidityProviderTokenRealAPY(c context.Context, req *types.QueryLiquidityProviderTokenRealAPYRequest) (*types.QueryLiquidityProviderTokenRealAPYResponse, error) {
}

func (k Keeper) LiquidityProviderTokenNominalAPY(c context.Context, req *types.QueryLiquidityProviderTokenNominalAPYRequest) (*types.QueryLiquidityProviderTokenNominalAPYResponse, error) {
}

func (k Keeper) Positions(c context.Context, req *types.QueryPositionsRequest) (*types.QueryPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	positions := []*types.WrappedPosition{}

	ctx := sdk.UnwrapSDKContext(c)
	if req.Address == "" {
		positions = k.GetAllPositions(ctx)
	} else {
		positions = k.GetUserPositions(ctx, sdk.AccAddress(req.Address))
	}

	return &types.QueryPositionsResponse{
		Positions: positions,
	}, nil
}
