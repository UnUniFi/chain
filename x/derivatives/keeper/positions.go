package keeper

import (
	"fmt"
	"math/big"
	"time"

	cdcTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetUserPositions(ctx sdk.Context, user sdk.AccAddress) []types.Position {
	store := ctx.KVStore(k.storeKey)

	positions := []types.Position{}
	it := sdk.KVStorePrefixIterator(store, types.AddressPositionKeyPrefix(user))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		positionAny := cdcTypes.Any{}
		k.cdc.Unmarshal(it.Value(), &positionAny)

		position, _ := types.UnpackPosition(&positionAny)

		positions = append(positions, position)
	}

	return positions
}

func (k Keeper) CreatePosition(ctx sdk.Context, positionKey []byte, positionAny cdcTypes.Any) {
	store := ctx.KVStore(k.storeKey)

	wrappedPosition := types.WrappedPosition{
		Id:       "",
		Address:  nil,
		StartAt:  *timestamppb.New(time.Now()), // TODO
		Position: positionAny,
	}

	bz := k.cdc.MustMarshal(&wrappedPosition)
	store.Set(positionKey, bz)
}

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) error {
	sender := msg.Sender.AccAddress()
	positions := k.GetUserPositions(ctx, sender)
	positionCount := len(positions)

	positionKey := types.AddressPositionWithIdKeyPrefix(sender, positionCount+1)

	// Not sure how to convert any type to position type
	k.CreatePosition(ctx, positionKey, msg.Position)

	return nil
}

func (k Keeper) Claim(ctx sdk.Context, msg *types.MsgClaim) error {

	return nil
}

func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) error {
	return nil
}
