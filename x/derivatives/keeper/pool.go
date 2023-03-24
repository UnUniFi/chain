package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetPoolAssets(ctx sdk.Context) []types.PoolParams_Asset {
	store := ctx.KVStore(k.storeKey)

	assets := []types.PoolParams_Asset{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixDerivativesPoolAssets))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		asset := types.PoolParams_Asset{}
		k.cdc.Unmarshal(it.Value(), &asset)

		assets = append(assets, asset)
	}

	return assets
}

func (k Keeper) GetPoolAssetByDenom(ctx sdk.Context, denom string) types.PoolParams_Asset {
	store := ctx.KVStore(k.storeKey)

	asset := types.PoolParams_Asset{}
	bz := store.Get(types.AssetKeyPrefix(denom))
	k.cdc.MustUnmarshal(bz, &asset)

	return asset
}

func (k Keeper) AddPoolAsset(ctx sdk.Context, asset types.PoolParams_Asset) {
	store := ctx.KVStore(k.storeKey)

	// TODO: remove below two lines as to change the way to handle PoolParams_Asset validation or reference
	bz := k.cdc.MustMarshal(&asset)
	store.Set(types.AssetKeyPrefix(asset.Denom), bz)

	coin := sdk.Coin{
		Denom:  asset.Denom,
		Amount: sdk.ZeroInt(),
	}
	coinBz := k.cdc.MustMarshal(&coin)
	store.Set(types.AssetDepositKeyPrefix(asset.Denom), coinBz)
}

func (k Keeper) IsAssetValid(ctx sdk.Context, iasset types.PoolParams_Asset) bool {
	assets := k.GetPoolAssets(ctx)

	for _, asset := range assets {
		if asset.Denom == iasset.Denom {
			return true
		}
	}

	return false
}

// TODO: The name GetAssetBalance is weird. We need to change the name like "GetAssetAmountInPool"
// TODO: Furthermore, is this really needed? Can we just use the bankKeeper function of getBalance for pool module account?
func (k Keeper) GetAssetBalance(ctx sdk.Context, denom string) sdk.Coin {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AssetDepositKeyPrefix(denom))
	if bz == nil {
		return sdk.NewCoin(denom, sdk.ZeroInt())
	}

	coin := sdk.Coin{}
	k.cdc.MustUnmarshal(bz, &coin)

	return coin
}

func (k Keeper) SetAssetBalance(ctx sdk.Context, coin sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&coin)
	store.Set(types.AssetDepositKeyPrefix(coin.Denom), bz)
}

func (k Keeper) GetAssetTargetAmount(ctx sdk.Context, denom string) (sdk.Coin, error) {
	mc := k.GetPoolMarketCap(ctx)
	asset := k.GetPoolAssetByDenom(ctx, denom)

	price, err := k.GetAssetPrice(ctx, denom)
	if err != nil {
		return sdk.Coin{}, err
	}

	targetAmount := mc.Total.Mul(asset.TargetWeight).Quo(price.Price)
	return sdk.NewCoin(denom, targetAmount.TruncateInt()), nil
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

func (k Keeper) GetUserDenomDepositAmount(ctx sdk.Context, depositer sdk.AccAddress, denom string) sdk.Int {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AddressAssetPoolDepositKeyPrefix(depositer, denom))
	if bz == nil {
		return sdk.ZeroInt()
	}

	coin := sdk.Coin{}
	k.cdc.MustUnmarshal(bz, &coin)
	return coin.Amount
}

func (k Keeper) SetUserDenomDepositAmount(ctx sdk.Context, addr sdk.AccAddress, denom string, amount sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	coin := sdk.Coin{
		Denom:  denom,
		Amount: amount,
	}
	bz := k.cdc.MustMarshal(&coin)
	store.Set(types.AddressAssetPoolDepositKeyPrefix(addr, denom), bz)
}

func (k Keeper) DepositPoolAsset(ctx sdk.Context, depositor sdk.AccAddress, asset sdk.Coin) {
	userDenomDepositAmount := k.GetUserDenomDepositAmount(ctx, depositor, asset.Denom)
	userDenomDepositAmount = userDenomDepositAmount.Add(asset.Amount)
	k.SetUserDenomDepositAmount(ctx, depositor, asset.Denom, userDenomDepositAmount)

	assetBalance := k.GetAssetBalance(ctx, asset.Denom)
	assetBalance.Amount = assetBalance.Amount.Add(asset.Amount)
	k.SetAssetBalance(ctx, assetBalance)
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
	assets := k.GetPoolAssets(ctx)

	breakdowns := []types.PoolMarketCap_Breakdown{}
	mc := sdk.NewDec(0)

	quoteTicker := k.GetPoolQuoteTicker(ctx)

	for _, asset := range assets {
		balance := k.GetAssetBalance(ctx, asset.Denom)
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

func (k Keeper) IsPriceReady(ctx sdk.Context) bool {
	assets := k.GetPoolAssets(ctx)

	for _, asset := range assets {
		_, err := k.GetAssetPrice(ctx, asset.Denom)
		if err != nil {
			ctx.EventManager().EmitTypedEvent(&types.EventPriceIsNotFeeded{
				Asset: asset.String(),
			})

			return false
		}
	}

	return true
}
