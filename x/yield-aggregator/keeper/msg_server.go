package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
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

func (k msgServer) DepositToVault(ctx context.Context, msg *types.MsgDepositToVault) (*types.MsgDepositToVaultResponse, error) {
	k.Keeper.DepositToVault()
	panic("implement me")
}

func (k msgServer) WithdrawFromVault(ctx context.Context, msg *types.MsgWithdrawFromVault) (*types.MsgWithdrawFromVaultResponse, error) {
	k.Keeper.WithdrawFromVault()
	panic("implement me")
}
