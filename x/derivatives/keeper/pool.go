package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
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

func (k Keeper) GetAssetBalance(ctx sdk.Context, asset types.Pool_Asset) sdk.Coin {
	store := ctx.KVStore(k.storeKey)

	coin := sdk.Coin{}
	bz := store.Get(types.AssetDepositKeyPrefix(asset.Denom))
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

func (k Keeper) DepositPoolAsset(ctx sdk.Context, depositor sdk.AccAddress, deposit_data sdk.Coin) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&deposit_data)

	store.Set(types.AddressAssetPoolDepositKeyPrefix(depositor, deposit_data.Denom), bz)

	key := types.AssetDepositKeyPrefix(deposit_data.Denom)
	coinBz := store.Get(key)
	coin := sdk.Coin{}
	k.cdc.MustUnmarshal(coinBz, &coin)
	coin.Amount.Add(deposit_data.Amount)

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

func (k Keeper) GetLPTokenSupplySnapshot(ctx sdk.Context, height int64) sdk.Int {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.AddressLPTokenSupplySnapshotKeyPrefix(height))
	supply := sdk.Int{}
	supply.Unmarshal(bz)

	return supply
}

func (k Keeper) SetLPTokenSupplySnapshot(ctx sdk.Context, height int64, supply sdk.Dec) error {
	bz, err := supply.Marshal()
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AddressLPTokenSupplySnapshotKeyPrefix(height), bz)

	return nil
}

func (k Keeper) GetQuoteTicker(ctx sdk.Context) string {
	return k.GetParams(ctx).Pool.QuoteTicker
}

func (k Keeper) GetPairRate(ctx sdk.Context, pair types.Market) (*sdk.Dec, error) {
	marketId, err := k.pricefeedKeeper.GetMarketIdFromDenom(ctx, pair.Denom, pair.QuoteDenom)
	if err != nil {
		return nil, err
	}
	price, err := k.pricefeedKeeper.GetCurrentPrice(ctx, marketId)

	return &price.Price, err
}

func (k Keeper) GetAssetPrice(ctx sdk.Context, denom string) (*pftypes.CurrentPrice, error) {
	ticker, err := k.pricefeedKeeper.GetTicker(ctx, denom)
	if err != nil {
		return nil, err
	}
	quoteTicker := k.GetQuoteTicker(ctx)

	price, err := k.GetPrice(ctx, ticker, quoteTicker)

	return &price, err
}

func (k Keeper) GetPoolMarketCap(ctx sdk.Context) types.PoolMarketCap {
	assets := k.GetPoolAssets(ctx)

	breakdowns := []types.PoolMarketCap_Breakdown{}
	mc := sdk.NewDec(0)

	quoteTicker := k.GetQuoteTicker(ctx)

	for _, asset := range assets {
		balance := k.GetAssetBalance(ctx, asset)
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

func (k Keeper) GetLPTokenSupply(ctx sdk.Context) sdk.Int {
	return k.bankKeeper.GetSupply(ctx, types.LiquidityProviderTokenDenom).Amount
}

func (k Keeper) GetLPTokenPrice(ctx sdk.Context) sdk.Dec {
	return k.GetPoolMarketCap(ctx).CalculateLPTokenPrice(k.GetLPTokenSupply(ctx))
}

func (k Keeper) GetPrice(ctx sdk.Context, lhsTicker string, rhsTicker string) (pftypes.CurrentPrice, error) {
	marketId := fmt.Sprintf("%s:%s", lhsTicker, rhsTicker)
	return k.pricefeedKeeper.GetCurrentPrice(ctx, marketId)
}

func (k Keeper) GetImaginaryFundingRate(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get([]byte(types.KeyPrefixImaginaryFundingRate))
	fundingRate := sdk.MustNewDecFromStr(string(bz))

	return fundingRate
}

func (k Keeper) SetImaginaryFundingRate(ctx sdk.Context, fundingRate sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := []byte(fundingRate.String())

	store.Set([]byte(types.KeyPrefixImaginaryFundingRate), bz)
}

func (k Keeper) MintLiquidityProviderToken(ctx sdk.Context, msg *types.MsgMintLiquidityProviderToken) error {
	depositor := msg.Sender.AccAddress()
	depositData := sdk.Coin{
		Amount: msg.Amount.Amount,
		Denom:  msg.Amount.Denom,
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return err
	}

	price, err := k.GetAssetPrice(ctx, msg.Amount.Denom)
	if err != nil {
		return err
	}

	assetMc := price.Price.Mul(sdk.Dec(msg.Amount.Amount))

	// currently mint to module and need to send it to msg.sender
	currentSupply := k.bankKeeper.GetSupply(ctx, types.LiquidityProviderTokenDenom)
	if currentSupply.Amount.IsZero() {
		// first deposit should mint 1 token
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.NewCoin(types.LiquidityProviderTokenDenom, sdk.NewInt(1000000))})
	} else {
		dlpPrice := k.GetLPTokenPrice(ctx)

		newSupply := assetMc.Quo(dlpPrice)
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.NewCoin(types.LiquidityProviderTokenDenom, newSupply.RoundInt())})
		if err != nil {
			return err
		}
	}

	k.DepositPoolAsset(ctx, depositor, depositData)
	return nil
}

func (k Keeper) BurnLiquidityProviderToken(ctx sdk.Context, msg *types.MsgBurnLiquidityProviderToken) error {
	sender := msg.Sender.AccAddress()
	amount := msg.Amount

	userBalance := k.bankKeeper.GetBalance(ctx, sender, types.LiquidityProviderTokenDenom)
	if userBalance.Amount.LT(amount) {
		return types.ErrInvalidRedeemAmount
	}

	totalSupply := k.bankKeeper.GetSupply(ctx, types.LiquidityProviderTokenDenom)

	assets := k.GetPoolAssets(ctx)

	for _, asset := range assets {
		coinBalance := k.GetAssetBalance(ctx, asset)
		tempAmount := coinBalance.Amount.Mul(userBalance.Amount)
		balanceToRedeem := tempAmount.BigInt().Div(tempAmount.BigInt(), totalSupply.Amount.BigInt())

		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.Coins{sdk.NewCoin(asset.Denom, sdk.NewInt(balanceToRedeem.Int64()))})

		if err != nil {
			return err
		}
	}

	k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{sdk.NewCoin(types.LiquidityProviderTokenDenom, amount)})

	return nil
}
