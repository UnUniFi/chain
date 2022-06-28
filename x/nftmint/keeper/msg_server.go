package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/nftmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateClass(c context.Context, msg *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// TODO: Make some validation additionaly against:
	// - name
	// - symbol
	// - baseTokenUri
	// - description
	// - classUri

	classID, err := k.keeper.CreateClassId(ctx, msg.Sender.AccAddress())
	if err != nil {
		return nil, err
	}

	err = k.keeper.CreateClass(ctx, classID, msg.Name, msg.Symbol, msg.Description, msg.ClassUri)
	if err != nil {
		return nil, err
	}

	k.keeper.CreateClassAttributes(ctx, classID, msg.Sender.AccAddress(), msg.BaseTokenUri, msg.MintingPermission, msg.TokenSupplyCap)

	return &types.MsgCreateClassResponse{}, nil
}

func (k msgServer) SendClass(c context.Context, msg *types.MsgSendClass) (*types.MsgSendClassResponse, error) {
	return &types.MsgSendClassResponse{}, nil
}
func (k msgServer) UpdateBaseTokenUri(c context.Context, msg *types.MsgUpdateBaseTokenUri) (*types.MsgUpdateBaseTokenUriResponse, error) {
	return &types.MsgUpdateBaseTokenUriResponse{}, nil
}
func (k msgServer) UpdateTokenSupplyCap(c context.Context, msg *types.MsgUpdateTokenSupplyCap) (*types.MsgUpdateTokenSupplyCapResponse, error) {
	return &types.MsgUpdateTokenSupplyCapResponse{}, nil
}
func (k msgServer) MintNFT(c context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	return &types.MsgMintNFTResponse{}, nil
}
func (k msgServer) BurnNFT(c context.Context, msg *types.MsgBurnNFT) (*types.MsgBurnNFTResponse, error) {
	return &types.MsgBurnNFTResponse{}, nil
}
