package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetPoolAssets(ctx sdk.Context) []types.Pool_Asset {
	store := ctx.KVStore(k.storeKey)

	assets := []types.Pool_Asset{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixDerivativesPoolAssets))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		asset := types.Pool_Asset{}
		k.cdc.Unmarshal(it.Value(), &asset)

		assets = append(assets, asset)
	}

	return assets
}

func (k Keeper) GetPoolAssetByDenom(ctx sdk.Context, denom string) types.Pool_Asset {
	store := ctx.KVStore(k.storeKey)

	asset := types.Pool_Asset{}
	bz := store.Get(types.AssetKeyPrefix(denom))
	k.cdc.MustUnmarshal(bz, &asset)

	return asset
}

func (k Keeper) AddPoolAsset(ctx sdk.Context, asset types.Pool_Asset) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&asset)
	store.Set(types.AssetKeyPrefix(asset.Denom), bz)

	coin := sdk.Coin{
		Denom:  asset.Denom,
		Amount: sdk.ZeroInt(),
	}
	coinBz := k.cdc.MustMarshal(&coin)
	store.Set(types.AssetDepositKeyPrefix(asset.Denom), coinBz)
}

func (k Keeper) IsAssetValid(ctx sdk.Context, iasset types.Pool_Asset) bool {
	assets := k.GetPoolAssets(ctx)

	for _, asset := range assets {
		if asset.Denom == iasset.Denom {
			return true
		}
	}

	return false
}

func (k Keeper) GetAssetBalance(ctx sdk.Context, denom string) sdk.Coin {
	store := ctx.KVStore(k.storeKey)

	coin := sdk.Coin{}
	bz := store.Get(types.AssetDepositKeyPrefix(denom))
	k.cdc.MustUnmarshal(bz, &coin)

	return coin
}

func (k Keeper) GetUserDeposits(ctx sdk.Context, depositor sdk.AccAddress) []sdk.Coin {
	store := ctx.KVStore(k.storeKey)

	deposits := []sdk.Coin{}
	it := sdk.KVStorePrefixIterator(store, types.AddressPoolDepositKeyPrefix(depositor))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		deposit := sdk.Coin{}
		k.cdc.Unmarshal(it.Value(), &deposit)

		deposits = append(deposits, deposit)
	}

	return deposits
}

func (k Keeper) DepositPoolAsset(ctx sdk.Context, depositor sdk.AccAddress, asset sdk.Coin) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&asset)

	store.Set(types.AddressAssetPoolDepositKeyPrefix(depositor, asset.Denom), bz)

	key := types.AssetDepositKeyPrefix(asset.Denom)
	coinBz := store.Get(key)
	coin := sdk.Coin{}
	k.cdc.MustUnmarshal(coinBz, &coin)
	coin.Amount = coin.Amount.Add(asset.Amount)

	coinBz = k.cdc.MustMarshal(&coin)
	store.Set(key, coinBz)
}

func (k Keeper) GetPoolMarketCapSnapshot(ctx sdk.Context, height int64) types.PoolMarketCap {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.AddressPoolMarketCapSnapshotKeyPrefix(height))
	marketCap := types.PoolMarketCap{}
	k.cdc.Unmarshal(bz, &marketCap)

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
	return k.GetParams(ctx).Pool.QuoteTicker
}

func (k Keeper) GetPoolMarketCap(ctx sdk.Context) types.PoolMarketCap {
	assets := k.GetPoolAssets(ctx)

	breakdowns := []types.PoolMarketCap_Breakdown{}
	mc := sdk.NewDec(0)

	quoteTicker := k.GetPoolQuoteTicker(ctx)

	for _, asset := range assets {
		balance := k.GetAssetBalance(ctx, asset.Denom)
		price, err := k.GetAssetPrice(ctx, asset.Denom)

		if err != nil {
			panic("not able to calculate market cap")
		}

		breakdown := types.PoolMarketCap_Breakdown{
			Denom:  asset.Denom,
			Amount: balance.Amount,
			Price:  price.Price,
		}
		breakdowns = append(breakdowns, breakdown)
		mc.Add(sdk.Dec(sdk.NewDecFromInt(balance.Amount)).Mul(price.Price))
	}

	return types.PoolMarketCap{
		QuoteTicker: quoteTicker,
		Total:       mc,
		Breakdown:   breakdowns,
	}
}
