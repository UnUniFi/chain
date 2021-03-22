package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/lcnem/jpyx/x/incentive/types"
)

type (
	Keeper struct {
		cdc           codec.Marshaler
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		paramSpace    paramtypes.Subspace
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		cdpKeeper     types.CdpKeeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey,
	paramSpace paramtypes.Subspace, accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	cdpKeeper types.CdpKeeper) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramSpace:    paramSpace,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		cdpKeeper:     cdpKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetJpyxMintingClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetJpyxMintingClaim(ctx sdk.Context, addr sdk.AccAddress) (types.JpyxMintingClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JpyxMintingClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types.JpyxMintingClaim{}, false
	}
	var c types.JpyxMintingClaim
	k.cdc.MustUnmarshalBinaryBare(bz, &c)
	return c, true
}

// SetJpyxMintingClaim sets the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) SetJpyxMintingClaim(ctx sdk.Context, c types.JpyxMintingClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JpyxMintingClaimKeyPrefix)
	bz := k.cdc.MustMarshalBinaryBare(&c)
	store.Set(c.Owner, bz)

}

// DeleteJpyxMintingClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteJpyxMintingClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JpyxMintingClaimKeyPrefix)
	store.Delete(owner)
}

// IterateJpyxMintingClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateJpyxMintingClaims(ctx sdk.Context, cb func(c types.JpyxMintingClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JpyxMintingClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types.JpyxMintingClaim
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllJpyxMintingClaims returns all Claim objects in the store
func (k Keeper) GetAllJpyxMintingClaims(ctx sdk.Context) types.JpyxMintingClaims {
	cs := types.JpyxMintingClaims{}
	k.IterateJpyxMintingClaims(ctx, func(c types.JpyxMintingClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// GetPreviousJpyxMintingAccrualTime returns the last time a collateral type accrued Jpyx minting rewards
func (k Keeper) GetPreviousJpyxMintingAccrualTime(ctx sdk.Context, ctype string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousJpyxMintingRewardAccrualTimeKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return time.Time{}, false
	}
	blockTime.UnmarshalBinary(bz)

	return blockTime, true
}

// SetPreviousJpyxMintingAccrualTime sets the last time a collateral type accrued Jpyx minting rewards
func (k Keeper) SetPreviousJpyxMintingAccrualTime(ctx sdk.Context, ctype string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousJpyxMintingRewardAccrualTimeKeyPrefix)
	bz, _ := blockTime.MarshalBinary()
	store.Set([]byte(ctype), bz)
}

// IterateJpyxMintingAccrualTimes iterates over all previous Jpyx minting accrual times and preforms a callback function
func (k Keeper) IterateJpyxMintingAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousJpyxMintingRewardAccrualTimeKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var accrualTime time.Time
		var collateralType string
		collateralType = string(iterator.Value())
		accrualTime.UnmarshalBinary(iterator.Value())
		if cb(collateralType, accrualTime) {
			break
		}
	}
}

// GetJpyxMintingRewardFactor returns the current reward factor for an individual collateral type
func (k Keeper) GetJpyxMintingRewardFactor(ctx sdk.Context, ctype string) (factor sdk.Dec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JpyxMintingRewardFactorKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return sdk.ZeroDec(), false
	}
	factor.Unmarshal(bz)

	return factor, true
}

// SetJpyxMintingRewardFactor sets the current reward factor for an individual collateral type
func (k Keeper) SetJpyxMintingRewardFactor(ctx sdk.Context, ctype string, factor sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JpyxMintingRewardFactorKeyPrefix)
	bz, _ := factor.Marshal()
	store.Set([]byte(ctype), bz)
}
