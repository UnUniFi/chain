package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

// Register method record subjects info in RecipientContainer type
func (k Keeper) Register(ctx sdk.Context, msg *types.MsgRegister) (*[]types.WeightedAddress, error) {
	// check if the RecipientContainerId is already registered
	if _, exists := k.GetRecipientContainer(ctx, msg.RecipientContainerId); exists {
		return nil, sdkerrors.Wrap(types.ErrRegisteredIncentiveId, msg.RecipientContainerId)
	}

	// check the length of the RecipientContainerId by referring MaxInentiveUnitIdLen in the Params
	if err := types.ValidateRecipientContainerId(msg.RecipientContainerId); err != nil {
		return nil, err
	}

	var subjectInfoList []types.WeightedAddress
	for i := 0; i < len(msg.Addresses); i++ {
		subjectInfo := types.NewSubjectInfo(msg.Addresses[i], msg.Weights[i])
		subjectInfoList = append(subjectInfoList, subjectInfo)
	}

	recipientContainer := types.NewRecipientContainer(msg.RecipientContainerId, subjectInfoList)

	if err := k.SetRecipientContainer(ctx, recipientContainer); err != nil {
		return nil, err
	}

	// operation related to RecipientContainerIdsByAddr
	// if exists already, add incentuve unit id in msg into data
	// if not, newly create and set
	for _, addr := range msg.Addresses {
		recipientContainerIdsByAddr := k.GetRecipientContainerIdsByAddr(ctx, sdk.MustAccAddressFromBech32(addr))
		recipientContainerIdsByAddr = recipientContainerIdsByAddr.CreateOrUpdate(addr, msg.RecipientContainerId)

		if err := k.SetRecipientContainerIdsByAddr(ctx, recipientContainerIdsByAddr); err != nil {
			return nil, err
		}
	}

	return &subjectInfoList, nil
}

func (k Keeper) SetRecipientContainer(ctx sdk.Context, recipientContainer types.RecipientContainer) error {
	bz, err := k.cdc.Marshal(&recipientContainer)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainer))
	prefixStore.Set([]byte(recipientContainer.Id), bz)

	return nil
}

func (k Keeper) SetRecipientContainerIdsByAddr(ctx sdk.Context, recipientContainerIdsByAddr types.BelongingRecipientContainers) error {
	bz, err := k.cdc.Marshal(&recipientContainerIdsByAddr)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainerIdsByAddr))
	// Use byte array of accAddress as key
	addressKeyBytes := sdk.MustAccAddressFromBech32(recipientContainerIdsByAddr.Address).Bytes()
	prefixStore.Set(addressKeyBytes, bz)

	return nil
}

func (k Keeper) GetRecipientContainer(ctx sdk.Context, id string) (types.RecipientContainer, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainer))

	bz := prefixStore.Get([]byte(id))
	if bz == nil {
		return types.RecipientContainer{}, false
	}

	var recipientContainer types.RecipientContainer
	k.cdc.MustUnmarshal(bz, &recipientContainer)
	return recipientContainer, true
}

func (k Keeper) GetRecipientContainerIdsByAddr(ctx sdk.Context, address sdk.AccAddress) types.BelongingRecipientContainers {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainerIdsByAddr))

	bz := prefixStore.Get(address)
	if bz == nil {
		return types.BelongingRecipientContainers{}
	}

	var recipientContainerIdsByAddr types.BelongingRecipientContainers
	k.cdc.MustUnmarshal(bz, &recipientContainerIdsByAddr)
	return recipientContainerIdsByAddr
}

func (k Keeper) GetAllRecipientContainers(ctx sdk.Context) []types.RecipientContainer {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixRecipientContainer))
	defer it.Close()

	allRecipientContainers := []types.RecipientContainer{}
	for ; it.Valid(); it.Next() {
		var recipientContainer types.RecipientContainer
		k.cdc.MustUnmarshal(it.Value(), &recipientContainer)

		allRecipientContainers = append(allRecipientContainers, recipientContainer)
	}

	return allRecipientContainers
}

func (k Keeper) GetAllRecipientContainerIdsByAddrs(ctx sdk.Context) []types.BelongingRecipientContainers {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixRecipientContainerIdsByAddr))
	defer it.Close()

	allRecipientContainerIdsByAddrs := []types.BelongingRecipientContainers{}
	for ; it.Valid(); it.Next() {
		var recipientContainerIdsByAddr types.BelongingRecipientContainers
		k.cdc.MustUnmarshal(it.Value(), &recipientContainerIdsByAddr)

		allRecipientContainerIdsByAddrs = append(allRecipientContainerIdsByAddrs, recipientContainerIdsByAddr)
	}

	return allRecipientContainerIdsByAddrs
}

func (k Keeper) DeleteRecipientContainer(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainer))

	prefixStore.Delete([]byte(id))
}

func (k Keeper) DeleteRecipientContainerIdsByAddr(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainerIdsByAddr))

	prefixStore.Delete(address)
}
