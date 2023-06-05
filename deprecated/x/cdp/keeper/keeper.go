package keeper

import (
	"fmt"
	"time"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/UnUniFi/chain/deprecated/cdp/types"
)

type (
	Keeper struct {
		cdc             codec.Codec
		storeKey        storetypes.StoreKey
		memKey          storetypes.StoreKey
		paramSpace      paramtypes.Subspace
		accountKeeper   types.AccountKeeper
		bankKeeper      types.BankKeeper
		auctionKeeper   types.AuctionKeeper
		PricefeedKeeper types.PricefeedKeeper
		hooks           types.CdpHooks
		maccPerms       map[string][]string
	}
)

func NewKeeper(cdc codec.Codec, storeKey, memKey storetypes.StoreKey, paramSpace paramtypes.Subspace, accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	auctionKeeper types.AuctionKeeper, PricefeedKeeper types.PricefeedKeeper, maccPerms map[string][]string) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:             cdc,
		storeKey:        storeKey,
		memKey:          memKey,
		paramSpace:      paramSpace,
		accountKeeper:   accountKeeper,
		bankKeeper:      bankKeeper,
		auctionKeeper:   auctionKeeper,
		PricefeedKeeper: PricefeedKeeper,
		hooks:           nil,
		maccPerms:       maccPerms,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetHooks sets the cdp keeper hooks
func (k *Keeper) SetHooks(hooks types.CdpHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set validator hooks twice")
	}
	k.hooks = hooks
	return k
}

// CdpDenomIndexIterator returns an sdk.Iterator for all cdps with matching collateral denom
func (k Keeper) CdpDenomIndexIterator(ctx sdk.Context, collateralType string) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpKey))
	db, found := k.GetCollateralTypePrefix(ctx, collateralType)
	if !found {
		panic(fmt.Sprintf("denom %s prefix not found", collateralType))
	}
	return sdk.KVStorePrefixIterator(store, types.DenomIterKey(db))
}

// CdpCollateralRatioIndexIterator returns an sdk.Iterator for all cdps that have collateral denom
// matching denom and collateral:debt ratio LESS THAN targetRatio
func (k Keeper) CdpCollateralRatioIndexIterator(ctx sdk.Context, collateralType string, targetRatio sdk.Dec) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpIDCollateralRatioIndex))
	db, found := k.GetCollateralTypePrefix(ctx, collateralType)
	if !found {
		panic(fmt.Sprintf("denom %s prefix not found", collateralType))
	}
	return store.Iterator(types.CollateralRatioIterKey(db, sdk.ZeroDec()), types.CollateralRatioIterKey(db, targetRatio))
}

// IterateAllCdps iterates over all cdps and performs a callback function
func (k Keeper) IterateAllCdps(ctx sdk.Context, cb func(cdp types.Cdp) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var cdp types.Cdp
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &cdp)

		if cb(cdp) {
			break
		}
	}
}

// IterateCdpsByCollateralType iterates over cdps with matching denom and performs a callback function
func (k Keeper) IterateCdpsByCollateralType(ctx sdk.Context, collateralType string, cb func(cdp types.Cdp) (stop bool)) {
	iterator := k.CdpDenomIndexIterator(ctx, collateralType)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var cdp types.Cdp
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &cdp)
		if cb(cdp) {
			break
		}
	}
}

// IterateCdpsByCollateralRatio iterate over cdps with collateral denom equal to denom and
// collateral:debt ratio LESS THAN targetRatio and performs a callback function.
func (k Keeper) IterateCdpsByCollateralRatio(ctx sdk.Context, collateralType string, targetRatio sdk.Dec, cb func(cdp types.Cdp) (stop bool)) {
	iterator := k.CdpCollateralRatioIndexIterator(ctx, collateralType, targetRatio)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		_, id, _ := types.SplitCollateralRatioKey(iterator.Key())
		cdp, found := k.GetCdp(ctx, collateralType, id)
		if !found {
			panic(fmt.Sprintf("cdp %d does not exist", id))
		}
		if cb(cdp) {
			break
		}

	}
}

// GetSliceOfCdpsByRatioAndType returns a slice of cdps of size equal to the input cutoffCount
// sorted by target ratio in ascending order (ie, the lowest collateral:debt ratio cdps are returned first)
func (k Keeper) GetSliceOfCdpsByRatioAndType(ctx sdk.Context, cutoffCount sdk.Int, targetRatio sdk.Dec, collateralType string) (cdps types.Cdps) {
	count := sdk.ZeroInt()
	k.IterateCdpsByCollateralRatio(ctx, collateralType, targetRatio, func(cdp types.Cdp) bool {
		cdps = append(cdps, cdp)
		count = count.Add(sdk.OneInt())
		if count.GTE(cutoffCount) {
			return true
		}
		return false
	})
	return cdps
}

// GetPreviousAccrualTime returns the last time an individual market accrued interest
func (k Keeper) GetPreviousAccrualTime(ctx sdk.Context, ctype string) (previousAccrualTime time.Time, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousAccrualTime))
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return time.Time{}, false
	}
	previousAccrualTime.UnmarshalBinary(bz)

	return previousAccrualTime, true
}

// SetPreviousAccrualTime sets the most recent accrual time for a particular market
func (k Keeper) SetPreviousAccrualTime(ctx sdk.Context, ctype string, previousAccrualTime time.Time) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PreviousAccrualTime))
	bz, _ := previousAccrualTime.MarshalBinary()
	store.Set([]byte(ctype), bz)
}

// GetInterestFactor returns the current interest factor for an individual collateral type
func (k Keeper) GetInterestFactor(ctx sdk.Context, ctype string) (interestFactor sdk.Dec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterestFactor))
	bz := store.Get([]byte(ctype))
	if bz == nil {
		return sdk.ZeroDec(), false
	}
	interestFactor.Unmarshal(bz)

	return interestFactor, true
}

// SetInterestFactor sets the current interest factor for an individual collateral type
func (k Keeper) SetInterestFactor(ctx sdk.Context, ctype string, interestFactor sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterestFactor))
	bz, _ := interestFactor.Marshal()
	store.Set([]byte(ctype), bz)
}

// IncrementTotalPrincipal increments the total amount of debt that has been drawn with that collateral type
func (k Keeper) IncrementTotalPrincipal(ctx sdk.Context, collateralType string, principal sdk.Coin) {
	total := k.GetTotalPrincipal(ctx, collateralType, principal.Denom)
	total = total.Add(principal.Amount)
	k.SetTotalPrincipal(ctx, collateralType, principal.Denom, total)
}

// DecrementTotalPrincipal decrements the total amount of debt that has been drawn for a particular collateral type
func (k Keeper) DecrementTotalPrincipal(ctx sdk.Context, collateralType string, principal sdk.Coin) {
	total := k.GetTotalPrincipal(ctx, collateralType, principal.Denom)
	// NOTE: negative total principal can happen in tests due to rounding errors
	// in fee calculation
	total = sdk.MaxInt(total.Sub(principal.Amount), sdk.ZeroInt())
	k.SetTotalPrincipal(ctx, collateralType, principal.Denom, total)
}

// GetTotalPrincipal returns the total amount of principal that has been drawn for a particular collateral
func (k Keeper) GetTotalPrincipal(ctx sdk.Context, collateralType, principalDenom string) (total sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrincipalKey))
	bz := store.Get([]byte(collateralType + principalDenom))
	if bz == nil {
		k.SetTotalPrincipal(ctx, collateralType, principalDenom, sdk.ZeroInt())
		return sdk.ZeroInt()
	}
	total.Unmarshal(bz)

	return total
}

// SetTotalPrincipal sets the total amount of principal that has been drawn for the input collateral
func (k Keeper) SetTotalPrincipal(ctx sdk.Context, collateralType, principalDenom string, total sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PrincipalKey))
	_, found := k.GetCollateralTypePrefix(ctx, collateralType)
	if !found {
		panic(fmt.Sprintf("collateral not found: %s", collateralType))
	}
	bz, _ := total.Marshal()
	store.Set([]byte(collateralType+principalDenom), bz)
}

// getModuleAccountCoins gets the total coin balance of this coin currently held by module accounts
func (k Keeper) getModuleAccountCoins(ctx sdk.Context, denom string) sdk.Coins {
	totalModCoinBalance := sdk.NewCoins(sdk.NewCoin(denom, sdk.ZeroInt()))
	for macc := range k.maccPerms {
		_macc := k.accountKeeper.GetModuleAccount(ctx, macc)
		modCoinBalance := k.bankKeeper.GetAllBalances(ctx, _macc.GetAddress()).AmountOf(denom)
		if modCoinBalance.IsPositive() {
			totalModCoinBalance = totalModCoinBalance.Add(sdk.NewCoin(denom, modCoinBalance))
		}
	}
	return totalModCoinBalance
}
