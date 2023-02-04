package keeper

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/UnUniFi/chain/x/pricefeed/types"
)

type (
	Keeper struct {
		cdc        codec.Codec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramSpace paramtypes.Subspace
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(cdc codec.Codec, storeKey, memKey storetypes.StoreKey, paramSpace paramtypes.Subspace, bankKeeper types.BankKeeper) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramSpace: paramSpace,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetPrice updates the posted price for a specific oracle
func (k Keeper) SetPrice(
	ctx sdk.Context,
	oracle sdk.AccAddress,
	marketID string,
	price sdk.Dec,
	expiry time.Time) (types.PostedPrice, error) {
	// If the expiry is less than or equal to the current blockheight, we consider the price valid
	if !expiry.After(ctx.BlockTime()) {
		return types.PostedPrice{}, types.ErrExpired
	}

	store := ctx.KVStore(k.storeKey)
	prices, err := k.GetRawPrices(ctx, marketID)
	if err != nil {
		return types.PostedPrice{}, err
	}
	var index int
	found := false
	for i := range prices {
		if prices[i].OracleAddress.AccAddress().Equals(oracle) {
			index = i
			found = true
			break
		}
	}

	// set the price for that particular oracle
	if found {
		prices[index] = types.NewPostedPrice(marketID, oracle, price, expiry)
	} else {
		prices = append(prices, types.NewPostedPrice(marketID, oracle, price, expiry))
		index = len(prices) - 1
	}

	// Emit an event containing the oracle's new price
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeOracleUpdatedPrice,
			sdk.NewAttribute(types.AttributeMarketID, marketID),
			sdk.NewAttribute(types.AttributeOracle, oracle.String()),
			sdk.NewAttribute(types.AttributeMarketPrice, price.String()),
			sdk.NewAttribute(types.AttributeExpiry, expiry.UTC().String()),
		),
	)

	bz, _ := json.Marshal(prices)
	store.Set(types.RawPriceKeySuffix(marketID), bz)
	return prices[index], nil
}

// SetCurrentPrices updates the price of an asset to the median of all valid oracle inputs
func (k Keeper) SetCurrentPrices(ctx sdk.Context, marketID string) error {
	_, ok := k.GetMarket(ctx, marketID)
	if !ok {
		return sdkerrors.Wrap(types.ErrInvalidMarket, marketID)
	}
	// store current price
	validPrevPrice := true
	prevPrice, err := k.GetCurrentPrice(ctx, marketID)
	if err != nil {
		validPrevPrice = false
	}

	prices, err := k.GetRawPrices(ctx, marketID)
	if err != nil {
		return err
	}
	var notExpiredPrices types.CurrentPrices
	// filter out expired prices
	for _, v := range prices {
		if v.Expiry.After(ctx.BlockTime()) {
			notExpiredPrices = append(notExpiredPrices, types.NewCurrentPrice(v.MarketId, v.Price))
		}
	}

	if len(notExpiredPrices) == 0 {
		// NOTE: The current price stored will continue storing the most recent (expired)
		// price if this is not set.
		// This zero's out the current price stored value for that market and ensures
		// that Cdp methods that GetCurrentPrice will return error.
		k.setCurrentPrice(ctx, marketID, types.CurrentPrice{})
		return types.ErrNoValidPrice
	}

	medianPrice := k.CalculateMedianPrice(ctx, notExpiredPrices)

	// check case that market price was not set in genesis
	if validPrevPrice && !medianPrice.Equal(prevPrice.Price) {
		// only emit event if price has changed
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMarketPriceUpdated,
				sdk.NewAttribute(types.AttributeMarketID, marketID),
				sdk.NewAttribute(types.AttributeMarketPrice, medianPrice.String()),
			),
		)
	}

	currentPrice := types.NewCurrentPrice(marketID, medianPrice)
	k.setCurrentPrice(ctx, marketID, currentPrice)

	return nil
}

func (k Keeper) setCurrentPrice(ctx sdk.Context, marketID string, currentPrice types.CurrentPrice) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.CurrentPriceKeySuffix(marketID), k.cdc.MustMarshal(&currentPrice))
}

// CalculateMedianPrice calculates the median prices for the input prices.
func (k Keeper) CalculateMedianPrice(ctx sdk.Context, prices types.CurrentPrices) sdk.Dec {
	l := len(prices)

	if l == 1 {
		// Return immediately if there's only one price
		return prices[0].Price
	}
	// sort the prices
	sort.Slice(prices, func(i, j int) bool {
		return prices[i].Price.LT(prices[j].Price)
	})
	// for even numbers of prices, the median is calculated as the mean of the two middle prices
	if l%2 == 0 {
		median := k.calculateMeanPrice(ctx, prices[l/2-1:l/2+1])
		return median
	}
	// for odd numbers of prices, return the middle element
	return prices[l/2].Price

}

func (k Keeper) calculateMeanPrice(ctx sdk.Context, prices types.CurrentPrices) sdk.Dec {
	sum := prices[0].Price.Add(prices[1].Price)
	mean := sum.Quo(sdk.NewDec(2))
	return mean
}

// GetCurrentPrice fetches the current median price of all oracles for a specific market
func (k Keeper) GetCurrentPrice(ctx sdk.Context, marketID string) (types.CurrentPrice, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.CurrentPriceKeySuffix(marketID))

	if bz == nil {
		return types.CurrentPrice{}, types.ErrNoValidPrice
	}
	var price types.CurrentPrice
	err := k.cdc.Unmarshal(bz, &price)
	if err != nil {
		return types.CurrentPrice{}, err
	}
	if price.Price.Equal(sdk.ZeroDec()) {
		return types.CurrentPrice{}, types.ErrNoValidPrice
	}
	return price, nil
}

// IterateCurrentPrices iterates over all current price objects in the store and performs a callback function
func (k Keeper) IterateCurrentPrices(ctx sdk.Context, cb func(cp types.CurrentPrice) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CurrentPriceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var cp types.CurrentPrice
		k.cdc.MustUnmarshal(iterator.Value(), &cp)
		if cb(cp) {
			break
		}
	}
}

// GetCurrentPrices returns all current price objects from the store
func (k Keeper) GetCurrentPrices(ctx sdk.Context) types.CurrentPrices {
	cps := types.CurrentPrices{}
	k.IterateCurrentPrices(ctx, func(cp types.CurrentPrice) (stop bool) {
		cps = append(cps, cp)
		return false
	})
	return cps
}

// GetRawPrices fetches the set of all prices posted by oracles for an asset
func (k Keeper) GetRawPrices(ctx sdk.Context, marketID string) (types.PostedPrices, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.RawPriceKeySuffix(marketID))
	if bz == nil {
		return types.PostedPrices{}, nil
	}
	var prices types.PostedPrices
	err := json.Unmarshal(bz, &prices)
	if err != nil {
		return types.PostedPrices{}, err
	}
	return prices, nil
}
