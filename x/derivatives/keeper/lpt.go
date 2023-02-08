package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetLPTokenBaseMintFee(ctx sdk.Context) sdk.Dec {
	return k.GetParams(ctx).Pool.BaseLptMintFee
}

func (k Keeper) GetLPTokenBaseRedeemFee(ctx sdk.Context) sdk.Dec {
	return k.GetParams(ctx).Pool.BaseLptRedeemFee
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

func (k Keeper) GetLPTokenSupply(ctx sdk.Context) sdk.Int {
	return k.bankKeeper.GetSupply(ctx, types.LiquidityProviderTokenDenom).Amount
}

func (k Keeper) GetLPTokenPrice(ctx sdk.Context) sdk.Dec {
	return k.GetPoolMarketCap(ctx).CalculateLPTokenPrice(k.GetLPTokenSupply(ctx))
}

func (k Keeper) GetRedeemDenomAmount(ctx sdk.Context, lptAmount sdk.Int, redeemDenom string) (sdk.Coin, sdk.Coin, error) {
	lptPrice := k.GetLPTokenPrice(ctx)
	redeemAssetPrice, err := k.GetAssetPrice(ctx, redeemDenom)

	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}

	totalSupply := k.bankKeeper.GetSupply(ctx, types.LiquidityProviderTokenDenom)
	if totalSupply.Amount.IsZero() {
		return sdk.Coin{}, sdk.Coin{}, types.ErrNoLiquidityProviderToken
	}

	redeemAssetBalance := k.GetAssetBalance(ctx, redeemDenom)

	// redeem amount = lptPrice * lptAmount / redeemAssetPrice
	redeemAmount := lptPrice.Mul(sdk.NewDecFromInt(lptAmount)).Quo(redeemAssetPrice.Price)

	if redeemAmount.GT(sdk.NewDecFromInt(redeemAssetBalance.Amount)) {
		return sdk.Coin{}, sdk.Coin{}, types.ErrInvalidRedeemAmount
	}

	targetAmount, err := k.GetAssetTargetAmount(ctx, redeemDenom)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}

	increaseRate := sdk.NewDecFromInt(targetAmount.Amount).Sub(sdk.NewDecFromInt(redeemAssetBalance.Amount)).Quo(sdk.NewDecFromInt(targetAmount.Amount))
	if increaseRate.IsNegative() {
		increaseRate = sdk.NewDec(0)
	}

	redeemFeeRate := k.GetLPTokenBaseRedeemFee(ctx).Mul(increaseRate.Add(sdk.NewDecWithPrec(1, 0)))

	return sdk.NewCoin(redeemDenom, redeemAmount.TruncateInt()), sdk.NewCoin(redeemDenom, redeemFeeRate.TruncateInt()), nil
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

	assetMc := price.Price.Mul(sdk.NewDecFromInt(msg.Amount.Amount))

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
	redeemDenom := msg.RedeemDenom

	userBalance := k.bankKeeper.GetBalance(ctx, sender, types.LiquidityProviderTokenDenom)
	if userBalance.Amount.LT(amount) {
		return types.ErrInvalidRedeemAmount
	}

	redeemAmount, redeemFee, err := k.GetRedeemDenomAmount(ctx, amount, redeemDenom)
	if err != nil {
		panic("failed to get redeemable amount")
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.Coins{redeemAmount.Sub(redeemFee)})

	// send redeem fee to fee collector

	k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{sdk.NewCoin(types.LiquidityProviderTokenDenom, amount)})

	return nil
}