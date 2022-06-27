package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/nftmint/types"
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

func (k msgServer) CreateClass(c context.Context, msg *types.MsgCreateClass) (*types.MsgCreateClassRespone, error)
func (k msgServer) SendClass(c context.Context, msg *types.MsgSendClass) (*types.MsgSendClassRespone, error)
func (k msgServer) UpdateBaseTokenUri(c context.Context, msg *types.MsgUpdateBaseTokenUri) (*types.MsgUpdateBaseTokenUriResponse, error)
func (k msgServer) UpdateTokenSupplyCap(c context.Context, msg *types.MsgUpdateTokenSupplyCap) (*types.MsgUpdateTokenSupplyCapResponse, error)
func (k msgServer) MintNFT(c context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTReponse, error)
func (k msgServer) BurnNFT(c context.Context, msg *types.MsgBurnNFT) (*types.MsgBurnNFTResponse, error)
