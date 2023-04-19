package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetPoolAcceptedAssetsConf(ctx sdk.Context) []types.PoolAssetConf {
	params := k.GetParams(ctx)
	return params.PoolParams.AcceptedAssetsConf
}

func (k Keeper) GetPoolAcceptedAssetConfByDenom(ctx sdk.Context, denom string) types.PoolAssetConf {
	params := k.GetParams(ctx)

	for _, assetConf := range params.PoolParams.AcceptedAssetsConf {
		if assetConf.Denom == denom {
			return assetConf
		}
	}

	panic(fmt.Sprintf("asset %s is not accepted", denom))
}

func (k Keeper) IsAssetAcceptable(ctx sdk.Context, denom string) bool {
	assets := k.GetPoolAcceptedAssetsConf(ctx)

	for _, asset := range assets {
		if asset.Denom == denom {
			return true
		}
	}

	return false
}

// GetAssetBalanceInPoolByDenom is used to get token balance of "derivatives" module account which is the liquidity pool.
func (k Keeper) GetAssetBalanceInPoolByDenom(ctx sdk.Context, denom string) sdk.Coin {
	derivModAddr := authtypes.NewModuleAddress(types.ModuleName)
	return k.bankKeeper.GetBalance(ctx, derivModAddr, denom)
}

// Return the current target amount of the asset in the pool.
func (k Keeper) GetAssetTargetAmount(ctx sdk.Context, denom string) (sdk.Coin, error) {
	mc := k.GetPoolMarketCap(ctx)
	asset := k.GetPoolAcceptedAssetConfByDenom(ctx, denom)

	price, err := k.GetAssetPrice(ctx, denom)
	if err != nil {
		return sdk.Coin{}, err
	}

	targetAmount := types.CalcTargetAmountInPool(asset.TargetWeight, price.Price, mc.Total)
	return sdk.NewCoin(denom, targetAmount), nil
}

func (k Keeper) GetPoolMarketCapSnapshot(ctx sdk.Context, height int64) types.PoolMarketCap {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.AddressPoolMarketCapSnapshotKeyPrefix(height))
	marketCap := types.PoolMarketCap{}
	if err := k.cdc.Unmarshal(bz, &marketCap); err != nil {
		return types.PoolMarketCap{}
	}

	return marketCap
}

func (k Keeper) SetPoolMarketCapSnapshot(ctx sdk.Context, height int64, marketCap types.PoolMarketCap) error {
	bz, err := k.cdc.Marshal(&marketCap)
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AddressPoolMarketCapSnapshotKeyPrefix(height), bz)

	return nil
}

func (k Keeper) GetPoolQuoteTicker(ctx sdk.Context) string {
	return k.GetParams(ctx).PoolParams.QuoteTicker
}

func (k Keeper) GetPoolMarketCap(ctx sdk.Context) types.PoolMarketCap {
	assets := k.GetPoolAcceptedAssetsConf(ctx)

	breakdowns := []types.PoolMarketCap_Breakdown{}
	mc := sdk.NewDec(0)

	quoteTicker := k.GetPoolQuoteTicker(ctx)

	for _, asset := range assets {
		balance := k.GetAssetBalanceInPoolByDenom(ctx, asset.Denom)
		price, err := k.GetAssetPrice(ctx, asset.Denom)

		if err != nil {
			panic(fmt.Sprintf("not able to calculate market cap: %s", err.Error()))
		}

		breakdown := types.PoolMarketCap_Breakdown{
			Denom:  asset.Denom,
			Amount: balance.Amount,
			Price:  price.Price,
		}
		breakdowns = append(breakdowns, breakdown)
		mc = mc.Add(sdk.Dec(sdk.NewDecFromInt(balance.Amount)).Mul(price.Price))
	}

	return types.PoolMarketCap{
		QuoteTicker: quoteTicker,
		Total:       mc,
		Breakdown:   breakdowns,
	}
}

// IsPriceReady returns true if all assets have price feeded.
// This is used to decide to run setPoolMarketCapSnapshot to avoid unnecessary snapshot.
func (k Keeper) IsPriceReady(ctx sdk.Context) bool {
	assets := k.GetPoolAcceptedAssetsConf(ctx)

	for _, asset := range assets {
		_, err := k.GetAssetPrice(ctx, asset.Denom)
		if err != nil {
			_ = ctx.EventManager().EmitTypedEvent(&types.EventPriceIsNotFeeded{
				Asset: asset.String(),
			})

			return false
		}
	}

	return true
}
