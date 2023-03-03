package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransferVaultOwnership(ctx context.Context, msg *types.MsgTransferVaultOwnership) (*types.MsgTransferVaultOwnershipResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.Context()
	panic("implement me")
}
