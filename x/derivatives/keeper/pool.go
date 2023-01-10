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
	deposit_data := types.UserDeposit{
		Amount: msg.Amount.Amount,
		Denom:  msg.Amount.Denom,
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return err
	}

	k.DepositPoolAsset(ctx, depositor, deposit_data)
	return nil
}
