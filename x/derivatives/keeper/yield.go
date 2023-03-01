package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) SaveBlockTimestamp(ctx sdk.Context, height int64, blockTime time.Time) {
	store := ctx.KVStore(k.storeKey)

	store.Set(types.BlockTimestampWithHeight(height), types.GetBlockTimestampBytes(blockTime.Unix()))
}

func (k Keeper) GetBlockTimestamp(ctx sdk.Context, height int64) time.Time {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.BlockTimestampWithHeight(height))
	if bz == nil {
		return time.Time{}
	}

	return time.Unix(types.GetBlockTimestampFromBytes(bz), 0)
}

func (k Keeper) GetLPTPriceFromSnapshot(ctx sdk.Context, height int64) sdk.Dec {
	poolMarketCap := k.GetPoolMarketCapSnapshot(ctx, height)
	lptSupply := k.GetLPTokenSupplySnapshot(ctx, height)

	lptPrice := poolMarketCap.CalculateLPTokenPrice(lptSupply)
	return lptPrice
}

func (k Keeper) GetLPNominalYieldRate(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
	lptPriceBefore := k.GetLPTPriceFromSnapshot(ctx, beforeHeight)
	if lptPriceBefore.IsZero() {
		return sdk.ZeroDec()
	}

	lptPriceAfter := sdk.ZeroDec()
	if afterHeight > ctx.BlockHeight() {
		return sdk.ZeroDec()
	} else if afterHeight == ctx.BlockHeight() {
		lptPriceAfter = k.GetLPTokenPrice(ctx)
	} else {
		lptPriceAfter = k.GetLPTPriceFromSnapshot(ctx, afterHeight)
	}

	diff := lptPriceAfter.Sub(lptPriceBefore)

	return diff.Quo(lptPriceBefore)
}

func (k Keeper) GetInflationRateOfAssetsInPool(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
	poolMarketCapBefore := k.GetPoolMarketCapSnapshot(ctx, beforeHeight)
	poolMarketCapAfter := k.GetPoolMarketCapSnapshot(ctx, afterHeight)

	poolMarketCapAfterWithBeforeAmount := sdk.NewDec(0)

	// TODO: consider an overflow of poolMarketCapAfter[i]
	// It might be better to use map type with string key (denom is used for key)
	for i := range poolMarketCapBefore.Breakdown {
		amountBefore := poolMarketCapBefore.Breakdown[i].Amount
		priceAfter := poolMarketCapAfter.Breakdown[i].Price

		poolMarketCapAfterWithBeforeAmount.Add(sdk.NewDecFromInt(amountBefore).Mul(priceAfter))
	}

	diff := poolMarketCapAfterWithBeforeAmount.Sub(poolMarketCapBefore.Total)

	return diff.Quo(poolMarketCapBefore.Total)
}

func (k Keeper) GetLPRealYieldRate(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
	// This is known as Fisher equation in Economics
	nominalInterestRate := k.GetLPNominalYieldRate(ctx, beforeHeight, afterHeight)
	inflationRate := k.GetInflationRateOfAssetsInPool(ctx, beforeHeight, afterHeight)

	nominalInterestRatePlus1 := nominalInterestRate.Add(sdk.NewDec(1))
	inflationRatePlus1 := inflationRate.Add(sdk.NewDec(1))

	quo := nominalInterestRatePlus1.Quo(inflationRatePlus1)

	realInterestRate := quo.Sub(sdk.NewDec(1))

	return realInterestRate
}

func (k Keeper) AnnualizeYieldRate(ctx sdk.Context, yieldRate sdk.Dec, beforeHeight int64, afterHeight int64) sdk.Dec {
	beforeBlockTimestamp := k.GetBlockTimestamp(ctx, beforeHeight).Unix()
	afterBlockTimestamp := k.GetBlockTimestamp(ctx, afterHeight).Unix()

	if beforeBlockTimestamp == afterBlockTimestamp {
		return sdk.ZeroDec()
	}

	annualizedYieldRate := yieldRate.Mul(sdk.NewDec(86400 * 365)).Quo(sdk.NewDec(afterBlockTimestamp - beforeBlockTimestamp))
	return annualizedYieldRate
}
