package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) TransferRequestedNfts(c context.Context, req *types.QueryTransferRequestedNftsRequest) (*types.QueryTransferRequestedNftsResponse, error) {
	var limit int
	if int(req.NftLimit) > 0 {
		limit = int(req.NftLimit)
	} else {
		limit = 10
	}
	ctx := sdk.UnwrapSDKContext(c)
	requsts, err := k.GetTransferRequests(ctx, limit)
	if err != nil {
		return nil, err
	}

	var transferRequestNfts []*types.QueryTransferRequestedNftResponse
	for _, v := range requsts {
		transferRequestNft, err := k.makeTransferRequestedNftResponse(ctx, v)
		if err != nil {
			return nil, status.Error(codes.NotFound, "not found nft")
		}
		transferRequestNfts = append(transferRequestNfts, transferRequestNft)
	}

	return &types.QueryTransferRequestedNftsResponse{
		Nfts: transferRequestNfts,
	}, nil
}

func (k Keeper) TransferRequestedNft(c context.Context, req *types.QueryTransferRequestedNftRequest) (*types.QueryTransferRequestedNftResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	transferReq, err := k.GetTransferRequestByIdBytes(ctx, []byte(req.NftId))
	if (transferReq == types.TransferRequest{} || err != nil) {
		return nil, status.Error(codes.NotFound, "not exist transfer request")
	}
	return k.makeTransferRequestedNftResponse(ctx, transferReq)
}

func (k Keeper) makeTransferRequestedNftResponse(ctx sdk.Context, tr types.TransferRequest) (*types.QueryTransferRequestedNftResponse, error) {
	nft, exists := k.nftKeeper.GetNFT(ctx, types.WrappedClassId, tr.NftId)
	if !exists {
		return nil, status.Error(codes.NotFound, "not exist transfer request")
	}
	addr, _ := sdk.AccAddressFromBech32(tr.Owner)
	ownerAddress := ununifitypes.StringAccAddress(addr)
	return &types.QueryTransferRequestedNftResponse{
		Nft: &types.TransferRequestNft{
			Id:         nft.Id,
			Uri:        nft.Uri,
			UriHash:    nft.UriHash,
			Owner:      ownerAddress,
			EthAddress: tr.EthAddress,
		},
	}, nil
}
