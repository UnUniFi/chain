package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (*types.Params, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyParams)
	if bz == nil {
		return nil, sdkerrors.ErrNotFound.Wrap("x/yieldaggregator module params")
	}

	var params types.Params
	if err := k.cdc.Unmarshal(bz, &params); err != nil {
		return nil, types.ErrParsingParams.Wrap(err.Error())
	}

	return &params, nil
}

func (k Keeper) SetParams(ctx sdk.Context, params *types.Params) error {
	store := ctx.KVStore(k.storeKey)

	bz, err := k.cdc.Marshal(params)
	if err != nil {
		return types.ErrParsingParams.Wrap(err.Error())
	}

	store.Set(types.KeyParams, bz)

	return nil
}
