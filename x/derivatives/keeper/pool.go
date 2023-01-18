package keeper

import (
	"fmt"
	"math/big"

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

func (k Keeper) GetUserDeposits(ctx sdk.Context, depositor sdk.AccAddress) []types.UserDeposit {
	store := ctx.KVStore(k.storeKey)

	deposits := []types.UserDeposit{}
	it := sdk.KVStorePrefixIterator(store, types.AddressDepositKeyPrefix(depositor))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		deposit := types.UserDeposit{}
		k.cdc.Unmarshal(it.Value(), &deposit)

		deposits = append(deposits, deposit)
	}

	return deposits
}

func (k Keeper) DepositPoolAsset(ctx sdk.Context, depositor sdk.AccAddress, deposit_data types.UserDeposit) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&deposit_data)

	store.Set(types.AddressAssetDepositKeyPrefix(depositor, deposit_data.Denom), bz)

	key := types.AssetDepositKeyPrefix(deposit_data.Denom)
	coinBz := store.Get(key)
	coin := sdk.Coin{}
	k.cdc.MustUnmarshal(coinBz, &coin)
	coin.Amount.Add(deposit_data.Amount)

	coinBz = k.cdc.MustMarshal(&coin)
	store.Set(key, coinBz)
}

func (k Keeper) MintLiquidityProviderToken(ctx sdk.Context, msg *types.MsgMintLiquidityProviderToken) error {
	depositor := msg.Sender.AccAddress()
	depositData := types.UserDeposit{
		Amount: msg.Amount.Amount,
		Denom:  msg.Amount.Denom,
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return err
	}

	marketId := fmt.Sprintf("%s:%s", msg.Amount.Denom, "USDC")
	price, err := k.pricefeedKeeper.GetCurrentPrice(ctx, marketId)
	if err != nil {
		return err
	}

	dlpMarketId := fmt.Sprintf("%s:%s", "DLP", "USDC")
	assetMc := price.Price.Mul(sdk.Dec(msg.Amount.Amount))

	// currently mint to module and need to send it to msg.sender
	currentSupply := k.bankKeeper.GetSupply(ctx, "DLP")
	if currentSupply.Amount.IsZero() {
		// first deposit should mint 1 million tokens
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.NewCoin("DLP", sdk.NewInt(1000000))})
		initialDlpPrice := *(assetMc.BigInt().Div(assetMc.BigInt(), big.NewInt(1000000)))
		k.pricefeedKeeper.SetCurrentPrice(ctx, dlpMarketId, pftypes.CurrentPrice{Price: sdk.Dec(initialDlpPrice)})
	} else {
		dlpPrice, err := k.pricefeedKeeper.GetCurrentPrice(ctx, dlpMarketId)
		if err != nil {
			return err
		}

		newSupply := *(assetMc.BigInt().Div(assetMc.BigInt(), dlpPrice.Price.BigInt()))
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.NewCoin("DLP", sdk.NewInt(newSupply.Int64()))})
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

	userBalance := k.bankKeeper.GetBalance(ctx, sender, "DLP")
	if userBalance.Amount.LT(amount) {
		return types.ErrInvalidRedeemAmount
	}

	totalSupply := k.bankKeeper.GetSupply(ctx, "DLP")

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

	k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{sdk.NewCoin("DLP", amount)})

	return nil
}
