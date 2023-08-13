package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) addDenomFromCreator(ctx sdk.Context, creator, classId string) {
	store := k.GetCreatorPrefixStore(ctx, creator)
	store.Set([]byte(classId), []byte(classId))
}

func (k Keeper) getDenomsFromCreator(ctx sdk.Context, creator string) []string {
	store := k.GetCreatorPrefixStore(ctx, creator)

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	classIds := []string{}
	for ; iterator.Valid(); iterator.Next() {
		classIds = append(classIds, string(iterator.Key()))
	}
	return classIds
}

func (k Keeper) GetAllDenomsIterator(ctx sdk.Context) sdk.Iterator {
	return k.GetCreatorsPrefixStore(ctx).Iterator(nil, nil)
}
