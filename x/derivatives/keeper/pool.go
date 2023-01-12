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
	store.Set(types.AssetKeyPrefix(asset.GetDenom()), bz)
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
	store.Set(types.AddressDepositKeyPrefix(depositor), bz)
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
		k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.NewCoin("DLP", sdk.NewInt(newSupply.Int64()))})
	}

	k.DepositPoolAsset(ctx, depositor, depositData)
	return nil
}

func (k Keeper) BurnLiquidityProviderToken(ctx sdk.Context, msg *types.MsgBurnLiquidityProviderToken) error {

	return nil
}

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) error {
	return nil
}

func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) error {
	return nil
}
