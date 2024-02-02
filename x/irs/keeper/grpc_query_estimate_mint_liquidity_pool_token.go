package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateMintLiquidityPoolToken(c context.Context, req *types.QueryEstimateMintLiquidityPoolTokenRequest) (*types.QueryEstimateMintLiquidityPoolTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	pool, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	// initial deposit
	if pool.TotalShares.IsZero() {
		return &types.QueryEstimateMintLiquidityPoolTokenResponse{
			MintAmount:               sdk.NewCoin(types.LsDenom(pool), types.OneShare),
			AdditionalRequiredAmount: sdk.Coin{},
		}, nil
	}
	tokenInAmount, ok := sdk.NewIntFromString(req.Amount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	mint, require, err := k.CalculateMintLpAmount(ctx, pool, sdk.NewCoin(req.Denom, tokenInAmount))
	if err != nil {
		return nil, err
	}
	return &types.QueryEstimateMintLiquidityPoolTokenResponse{
		MintAmount:               mint,
		AdditionalRequiredAmount: require,
	}, nil
}
