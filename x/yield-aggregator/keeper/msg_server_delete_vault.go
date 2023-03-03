package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteVault(ctx context.Context, msg *types.MsgDeleteVault) (*types.MsgDeleteVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	k.Keeper.RemoveVault(sdkCtx, msg.VaultId)
	panic("implement me")
}
