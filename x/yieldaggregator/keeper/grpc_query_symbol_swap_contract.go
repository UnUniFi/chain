package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) SymbolSwapContractMap(c context.Context, req *types.QuerySymbolSwapContractMapRequest) (*types.QuerySymbolSwapContractMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QuerySymbolSwapContractMapResponse{
		Mappings: k.GetAllSymbolSwapContractMap(ctx),
	}, nil
}
