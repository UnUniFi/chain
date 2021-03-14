package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"

	"github.com/lcnem/jpyx/x/incentive/types"
)

// Keeper keeper for the incentive module
type Keeper struct {
	accountKeeper types.AccountKeeper
	cdc           *codec.Codec
	cdpKeeper     types.CdpKeeper
	key           sdk.StoreKey
	paramSubspace subspace.Subspace
	supplyKeeper  types.SupplyKeeper
	stakingKeeper types.StakingKeeper
}

// NewKeeper creates a new keeper
func NewKeeper(
	cdc *codec.Codec, key sdk.StoreKey, paramstore subspace.Subspace, sk types.SupplyKeeper,
	cdpk types.CdpKeeper, ak types.AccountKeeper, stk types.StakingKeeper,
) Keeper {

	return Keeper{
		accountKeeper: ak,
		cdc:           cdc,
		cdpKeeper:     cdpk,
		key:           key,
		paramSubspace: paramstore.WithKeyTable(types.ParamKeyTable()),
		supplyKeeper:  sk,
		stakingKeeper: stk,
	}
}

// GetJPYXMintingClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetJPYXMintingClaim(ctx sdk.Context, addr sdk.AccAddress) (types.JPYXMintingClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JPYXMintingClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types.JPYXMintingClaim{}, false
	}
	var c types.JPYXMintingClaim
	k.cdc.MustUnmarshalBinaryBare(bz, &c)
	return c, true
}

// SetJPYXMintingClaim sets the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) SetJPYXMintingClaim(ctx sdk.Context, c types.JPYXMintingClaim) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JPYXMintingClaimKeyPrefix)
	bz := k.cdc.MustMarshalBinaryBare(c)
	store.Set(c.Owner, bz)

}

// DeleteJPYXMintingClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteJPYXMintingClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JPYXMintingClaimKeyPrefix)
	store.Delete(owner)
}

// IterateJPYXMintingClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateJPYXMintingClaims(ctx sdk.Context, cb func(c types.JPYXMintingClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JPYXMintingClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types.JPYXMintingClaim
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllJPYXMintingClaims returns all Claim objects in the store
func (k Keeper) GetAllJPYXMintingClaims(ctx sdk.Context) types.JPYXMintingClaims {
	cs := types.JPYXMintingClaims{}
	k.IterateJPYXMintingClaims(ctx, func(c types.JPYXMintingClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// GetPreviousJPYXMintingAccrualTime returns the last time a collateral type accrued JPYX minting rewards
func (k Keeper) GetPreviousJPYXMintingAccrualTime(ctx sdk.Context, ctype string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJPYXMintingRewardAccrualTimeKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return time.Time{}, false
	}
	k.cdc.MustUnmarshalBinaryBare(bz, &blockTime)
	return blockTime, true
}

// SetPreviousJPYXMintingAccrualTime sets the last time a collateral type accrued JPYX minting rewards
func (k Keeper) SetPreviousJPYXMintingAccrualTime(ctx sdk.Context, ctype string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJPYXMintingRewardAccrualTimeKeyPrefix)
	store.Set([]byte(ctype), k.cdc.MustMarshalBinaryBare(blockTime))
}

// IterateJPYXMintingAccrualTimes iterates over all previous JPYX minting accrual times and preforms a callback function
func (k Keeper) IterateJPYXMintingAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.PreviousJPYXMintingRewardAccrualTimeKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var accrualTime time.Time
		var collateralType string
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &collateralType)
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &accrualTime)
		if cb(collateralType, accrualTime) {
			break
		}
	}
}

// GetJPYXMintingRewardFactor returns the current reward factor for an individual collateral type
func (k Keeper) GetJPYXMintingRewardFactor(ctx sdk.Context, ctype string) (factor sdk.Dec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JPYXMintingRewardFactorKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return sdk.ZeroDec(), false
	}
	k.cdc.MustUnmarshalBinaryBare(bz, &factor)
	return factor, true
}

// SetJPYXMintingRewardFactor sets the current reward factor for an individual collateral type
func (k Keeper) SetJPYXMintingRewardFactor(ctx sdk.Context, ctype string, factor sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.JPYXMintingRewardFactorKeyPrefix)
	store.Set([]byte(ctype), k.cdc.MustMarshalBinaryBare(factor))
}
