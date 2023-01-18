package keeper

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func (k Keeper) GetUserPositionsLength(ctx sdk.Context, user sdk.AccAddress) []types.Position {
	store := ctx.KVStore(k.storeKey)

	positions := []types.Position{}
	it := sdk.KVStorePrefixIterator(store, types.AddressPositionKeyPrefix(user))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		position := types.Position{}

		k.cdc.Unmarshal(it.Value(), &position)

		positions = append(positions, position)
	}

	return positions
}

func (k Keeper) CreatePosition(ctx sdk.Context, positionKey []byte, position types.Position) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(*position)
	store.Set(positionKey, bz)
}

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) error {
	sender := msg.Sender.AccAddress()
	positions := k.GetUserPositionsLength(ctx, sender)
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
