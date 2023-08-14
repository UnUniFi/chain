package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
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

func (k msgServer) CancelListing(c context.Context, msg *types.MsgCancelListing) (*types.MsgCancelListingResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.CancelNftListing(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgCancelListingResponse{}, nil
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

func (k msgServer) SellingDecision(c context.Context, msg *types.MsgSellingDecision) (*types.MsgSellingDecisionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.SetSellingDecision(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgSellingDecisionResponse{}, nil
}

func (k msgServer) PayRemainder(c context.Context, msg *types.MsgPayRemainder) (*types.MsgPayRemainderResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	err := k.keeper.PayRemainder(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgPayRemainderResponse{}, nil
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

func (k msgServer) UpdateParams(c context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	

	return &types.MsgUpdateParamsResponse{}, nil
}
