package keeper

// import (
// 	"context"
// 	"strings"

// 	errorsmod "cosmossdk.io/errors"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

// 	"github.com/bianjieai/nft-transfer/types"
// 	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
// )

// var _ types.MsgServer = Keeper{}

// // Transfer defines a rpc handler method for MsgTransfer.
// func (k Keeper) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	sender, err := sdk.AccAddressFromBech32(msg.Sender)
// 	if err != nil {
// 		return nil, err
// 	}

// 	nftData, found := k.nftKeeper.GetNftData(ctx, msg.ClassId, msg.TokenId)
// 	if !found {
// 		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown NFT %s", msg.Id)
// 	}

// 	if nftData.SendDisabled {
// 		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Sending NFT %s of class %s is disabled", msg.Id, msg.ClassId)
// 	}

// 	sequence, err := k.SendTransfer(
// 		ctx, msg.SourcePort, msg.SourceChannel, msg.ClassId, msg.TokenIds,
// 		sender, msg.Receiver, msg.TimeoutHeight, msg.TimeoutTimestamp, msg.Memo,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	k.Logger(ctx).Info("IBC non-fungible token transfer",
// 		"classID", msg.ClassId,
// 		"tokenIDs", strings.Join(msg.TokenIds, ","),
// 		"sender", msg.Sender,
// 		"receiver", msg.Receiver,
// 	)

// 	ctx.EventManager().EmitEvents(sdk.Events{
// 		sdk.NewEvent(
// 			types.EventTypeTransfer,
// 			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
// 			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
// 		),
// 		sdk.NewEvent(
// 			sdk.EventTypeMessage,
// 			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
// 		),
// 	})

// 	return &types.MsgTransferResponse{Sequence: sequence}, nil
// }
