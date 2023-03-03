package keeper

import (
	"context"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawFromVault(ctx context.Context, msg *types.MsgWithdrawFromVault) (*types.MsgWithdrawFromVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.Context()
	panic("implement me")
}
