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

// GetJPYXMintingClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetJPYXMintingClaim(ctx sdk.Context, addr sdk.AccAddress) (types.JPYXMintingClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JPYXMintingClaimKeyPrefix)
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JPYXMintingClaimKeyPrefix)
	bz := k.cdc.MustMarshalBinaryBare(&c)
	store.Set(c.Owner, bz)

}

// DeleteJPYXMintingClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteJPYXMintingClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JPYXMintingClaimKeyPrefix)
	store.Delete(owner)
}

// IterateJPYXMintingClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateJPYXMintingClaims(ctx sdk.Context, cb func(c types.JPYXMintingClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JPYXMintingClaimKeyPrefix)
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousJPYXMintingRewardAccrualTimeKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return time.Time{}, false
	}
	blockTime.UnmarshalBinary(bz)

	return blockTime, true
}

// SetPreviousJPYXMintingAccrualTime sets the last time a collateral type accrued JPYX minting rewards
func (k Keeper) SetPreviousJPYXMintingAccrualTime(ctx sdk.Context, ctype string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousJPYXMintingRewardAccrualTimeKeyPrefix)
	bz, _ := blockTime.MarshalBinary()
	store.Set([]byte(ctype), bz)
}

// IterateJPYXMintingAccrualTimes iterates over all previous JPYX minting accrual times and preforms a callback function
func (k Keeper) IterateJPYXMintingAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousJPYXMintingRewardAccrualTimeKeyPrefix)
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

// GetJPYXMintingRewardFactor returns the current reward factor for an individual collateral type
func (k Keeper) GetJPYXMintingRewardFactor(ctx sdk.Context, ctype string) (factor sdk.Dec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JPYXMintingRewardFactorKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return sdk.ZeroDec(), false
	}
	factor.Unmarshal(bz)

	return factor, true
}

// SetJPYXMintingRewardFactor sets the current reward factor for an individual collateral type
func (k Keeper) SetJPYXMintingRewardFactor(ctx sdk.Context, ctype string, factor sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.JPYXMintingRewardFactorKeyPrefix)
	bz, _ := factor.Marshal()
	store.Set([]byte(ctype), bz)
}
