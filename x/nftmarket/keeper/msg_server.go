package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) ListNft(c context.Context, msg *types.MsgListNft) (*types.MsgListNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.keeper.ListNft(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgListNftResponse{}, nil
}

func (k msgServer) CancelNftListing(c context.Context, msg *types.MsgCancelNftListing) (*types.MsgCancelNftListingResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.CancelNftListing(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgCancelNftListingResponse{}, nil
}

func (k msgServer) ExpandListingPeriod(c context.Context, msg *types.MsgExpandListingPeriod) (*types.MsgExpandListingPeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.ExpandListingPeriod(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgExpandListingPeriodResponse{}, nil
}

func (k msgServer) PlaceBid(c context.Context, msg *types.MsgPlaceBid) (*types.MsgPlaceBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.PlaceBid(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgPlaceBidResponse{}, nil
}

func (k msgServer) CancelBid(c context.Context, msg *types.MsgCancelBid) (*types.MsgCancelBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.CancelBid(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgCancelBidResponse{}, nil
}

func (k msgServer) EndNftListing(c context.Context, msg *types.MsgEndNftListing) (*types.MsgEndNftListingResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.EndNftListing(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgEndNftListingResponse{}, nil
}

func (k msgServer) PayFullBid(c context.Context, msg *types.MsgPayFullBid) (*types.MsgPayFullBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.PayFullBid(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgPayFullBidResponse{}, nil
}

func (k msgServer) Borrow(c context.Context, msg *types.MsgBorrow) (*types.MsgBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.Borrow(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgBorrowResponse{}, nil
}

func (k msgServer) Repay(c context.Context, msg *types.MsgRepay) (*types.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.Repay(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgRepayResponse{}, nil
}

func (k msgServer) MintStableCoin(c context.Context, msg *types.MsgMintStableCoin) (*types.MsgMintStableCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.MsgMintStableCoinResponse{}, nil
}

func (k msgServer) BurnStableCoin(c context.Context, msg *types.MsgBurnStableCoin) (*types.MsgBurnStableCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.MsgBurnStableCoinResponse{}, nil
}

func (k msgServer) Liquidate(c context.Context, msg *types.MsgLiquidate) (*types.MsgLiquidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.Liquidate(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgLiquidateResponse{}, nil
}
