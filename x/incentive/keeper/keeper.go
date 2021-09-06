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

// GetCdpMintingClaim returns the claim in the store corresponding the the input address collateral type and id and a boolean for if the claim was found
func (k Keeper) GetCdpMintingClaim(ctx sdk.Context, addr sdk.AccAddress) (types.CdpMintingClaim, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CdpMintingClaimKeyPrefix)
	bz := store.Get(addr)
	if bz == nil {
		return types.CdpMintingClaim{}, false
	}
	var c types.CdpMintingClaim
	k.cdc.MustUnmarshalBinaryBare(bz, &c)
	return c, true
}

// SetCdpMintingClaim sets the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) SetCdpMintingClaim(ctx sdk.Context, c types.CdpMintingClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CdpMintingClaimKeyPrefix)
	bz := k.cdc.MustMarshalBinaryBare(&c)
	store.Set(c.Owner, bz)

}

// DeleteCdpMintingClaim deletes the claim in the store corresponding to the input address, collateral type, and id
func (k Keeper) DeleteCdpMintingClaim(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CdpMintingClaimKeyPrefix)
	store.Delete(owner)
}

// IterateCdpMintingClaims iterates over all claim  objects in the store and preforms a callback function
func (k Keeper) IterateCdpMintingClaims(ctx sdk.Context, cb func(c types.CdpMintingClaim) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CdpMintingClaimKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var c types.CdpMintingClaim
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &c)
		if cb(c) {
			break
		}
	}
}

// GetAllCdpMintingClaims returns all Claim objects in the store
func (k Keeper) GetAllCdpMintingClaims(ctx sdk.Context) types.CdpMintingClaims {
	cs := types.CdpMintingClaims{}
	k.IterateCdpMintingClaims(ctx, func(c types.CdpMintingClaim) (stop bool) {
		cs = append(cs, c)
		return false
	})
	return cs
}

// GetPreviousCdpMintingAccrualTime returns the last time a collateral type accrued Cdp minting rewards
func (k Keeper) GetPreviousCdpMintingAccrualTime(ctx sdk.Context, ctype string) (blockTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousCdpMintingRewardAccrualTimeKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return time.Time{}, false
	}
	blockTime.UnmarshalBinary(bz)

	return blockTime, true
}

// SetPreviousCdpMintingAccrualTime sets the last time a collateral type accrued Cdp minting rewards
func (k Keeper) SetPreviousCdpMintingAccrualTime(ctx sdk.Context, ctype string, blockTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousCdpMintingRewardAccrualTimeKeyPrefix)
	bz, _ := blockTime.MarshalBinary()
	store.Set([]byte(ctype), bz)
}

// IterateCdpMintingAccrualTimes iterates over all previous Cdp minting accrual times and preforms a callback function
func (k Keeper) IterateCdpMintingAccrualTimes(ctx sdk.Context, cb func(string, time.Time) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PreviousCdpMintingRewardAccrualTimeKeyPrefix)
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

// GetCdpMintingRewardFactor returns the current reward factor for an individual collateral type
func (k Keeper) GetCdpMintingRewardFactor(ctx sdk.Context, ctype string) (factor sdk.Dec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CdpMintingRewardFactorKeyPrefix)
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return sdk.ZeroDec(), false
	}
	factor.Unmarshal(bz)

	return factor, true
}

// SetCdpMintingRewardFactor sets the current reward factor for an individual collateral type
func (k Keeper) SetCdpMintingRewardFactor(ctx sdk.Context, ctype string, factor sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CdpMintingRewardFactorKeyPrefix)
	bz, _ := factor.Marshal()
	store.Set([]byte(ctype), bz)
}

func (k Keeper) GetGenesisDenoms(ctx sdk.Context) (denoms *types.GenesisDenoms, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GenesisDenomsKeyPrefix)
	if bz == nil {
		return types.DefaultGenesisDenoms(), false
	}
	denoms.Unmarshal(bz)

	return denoms, true
}

func (k Keeper) SetGenesisDenoms(ctx sdk.Context, denoms *types.GenesisDenoms) {
	store := ctx.KVStore(k.storeKey)
	bz, _ := denoms.Marshal()
	store.Set(types.GenesisDenomsKeyPrefix, bz)
}
