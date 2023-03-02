package keeper

import (
	"context"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawFromVault(ctx context.Context, msg *types.MsgWithdrawFromVault) (*types.MsgWithdrawFromVaultResponse, error) {
	k.Keeper.WithdrawFromVault(sdk.UnwrapSDKContext(ctx), msg.Sender, msg.VaultDenom, msg.LpTokenAmount)
	panic("implement me")
}
