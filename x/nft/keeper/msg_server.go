package keeper

import (
	"context"
	// "encoding/json"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
	// authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var _ nft.MsgServer = Keeper{}

// Send implements Send method of the types.MsgServer.
func (k Keeper) Send(goCtx context.Context, msg *nft.MsgSend) (*nft.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	owner := k.GetOwner(ctx, msg.ClassId, msg.Id)
	if !owner.Equals(sender) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is not the owner of nft %s", sender, msg.Id)
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	// Check NftData
	nftData, found := k.GetNftData(ctx, msg.ClassId, msg.Id)
	if !found {
		return nil, nft.ErrNFTNotExists
	}

	if nftData.SendDisabled {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Sending NFT %s of class %s is disabled", msg.Id, msg.ClassId)
	}

	// Check ClassData
	// Must be done after checking NftData
	// class, found := k.GetClassData(ctx, msg.ClassId)
	// if !found {
	// 	return nil, nft.ErrClassNotExists
	// }
	// if class.SendPrehookContract != "" {
	// 	contractAddr, err := sdk.AccAddressFromBech32(class.SendPrehookContract)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	address := authtypes.NewModuleAddress(nft.ModuleName)
	// 	wasmMsg, err := json.Marshal(struct{}{})
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	_, err = k.wasmKeeper.Execute(ctx, contractAddr, address, []byte(wasmMsg), sdk.Coins{})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	if err := k.Transfer(ctx, msg.ClassId, msg.Id, receiver); err != nil {
		return nil, err
	}

	_ = ctx.EventManager().EmitTypedEvent(&nft.EventSend{
		ClassId:  msg.ClassId,
		Id:       msg.Id,
		Sender:   msg.Sender,
		Receiver: msg.Receiver,
	})
	return &nft.MsgSendResponse{}, nil
}
