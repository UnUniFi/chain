package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DepositToVault(ctx context.Context, msg *types.MsgDepositToVault) (*types.MsgDepositToVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.Context()
	panic("implement me")
}
