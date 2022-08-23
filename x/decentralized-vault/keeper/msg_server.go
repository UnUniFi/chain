package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/decentralized-vault/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) NftLocked(c context.Context, msg *types.MsgNftLocked) (*types.MsgNftLockedResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.keeper.MintWrappedNft(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgNftLockedResponse{}, nil
}

func (k msgServer) NftUnlocked(c context.Context, msg *types.MsgNftUnlocked) (*types.MsgNftUnlockedResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.keeper.BurnWrappedNft(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgNftUnlockedResponse{}, nil
}

func (k msgServer) NftTransferRequest(c context.Context, msg *types.MsgNftTransferRequest) (*types.MsgNftTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.keeper.DepositWrappedNft(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgNftTransferRequestResponse{}, nil
}

func (k msgServer) NftRejectTransfer(c context.Context, msg *types.MsgNftRejectTransfer) (*types.MsgNftRejectTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.keeper.WithdrawWrappedNft(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgNftRejectTransferResponse{}, nil
}

func (k msgServer) NftTransferred(c context.Context, msg *types.MsgNftTransferred) (*types.MsgNftTransferredResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.keeper.NftTransferred(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgNftTransferredResponse{}, nil
}
