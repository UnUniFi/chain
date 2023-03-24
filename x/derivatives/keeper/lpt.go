package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func (k Keeper) GetLPTokenBaseMintFee(ctx sdk.Context) sdk.Dec {
	return k.GetParams(ctx).PoolParams.BaseLptMintFee
}

func (k Keeper) GetLPTokenBaseRedeemFee(ctx sdk.Context) sdk.Dec {
	return k.GetParams(ctx).PoolParams.BaseLptRedeemFee
}

func (k Keeper) GetLPTokenSupplySnapshot(ctx sdk.Context, height int64) sdk.Int {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.AddressLPTokenSupplySnapshotKeyPrefix(height))
	if bz == nil {
		return sdk.ZeroInt()
	}

	supply := sdk.Int{}
	supply.Unmarshal(bz)
	return supply
}

func (k Keeper) SetLPTokenSupplySnapshot(ctx sdk.Context, height int64, supply sdk.Int) error {
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

func (k Keeper) DetermineMintingLPTokenAmount(ctx sdk.Context, amount sdk.Coin) (sdk.Coin, error) {
	currentSupply := k.bankKeeper.GetSupply(ctx, types.LiquidityProviderTokenDenom)

	assetPrice, err := k.GetAssetPrice(ctx, amount.Denom)
	if err != nil {
		return sdk.Coin{}, err
	}

	assetMc := assetPrice.Price.Mul(sdk.NewDecFromInt(amount.Amount))
	if currentSupply.Amount.IsZero() {
		// TODO: we can eliminate unnecessary calculation -> assetPrice.Price
		return k.InitialLiquidityProviderTokenSupply(ctx, assetPrice, assetMc, amount.Denom)
	}

	lptPrice := k.GetLPTokenPrice(ctx)
	if lptPrice.IsZero() {
		return sdk.Coin{}, types.ErrZeroLpTokenPrice
	}

	mintAmount := assetMc.Quo(lptPrice)

	return sdk.NewCoin(types.LiquidityProviderTokenDenom, mintAmount.TruncateInt()), nil
}

func (k Keeper) LptDenomRate(ctx sdk.Context, redeemDenom string) (sdk.Dec, error) {
	lptPrice := k.GetLPTokenPrice(ctx)
	redeemAssetPrice, err := k.GetAssetPrice(ctx, redeemDenom)
	if err != nil {
		return sdk.Dec{}, err
	}
	return lptPrice.Quo(redeemAssetPrice.Price), nil
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

	if redeemAssetPrice.Price.IsNil() || redeemAssetPrice.Price.IsZero() {
		return sdk.Coin{}, sdk.Coin{}, types.ErrInvalidRedeemAmount
	}
	// redeem amount = lptPrice * lptAmount / redeemAssetPrice
	totalRedeemAmount := lptPrice.Mul(sdk.NewDecFromInt(lptAmount)).Quo(redeemAssetPrice.Price)

	if totalRedeemAmount.GT(sdk.NewDecFromInt(redeemAssetBalance.Amount)) {
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

	redeem := sdk.NewCoin(redeemDenom, totalRedeemAmount.TruncateInt())
	fee := sdk.NewCoin(redeemDenom, totalRedeemAmount.Mul(redeemFeeRate).TruncateInt())
	redeem, err = redeem.SafeSub(fee)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}
	return redeem, fee, nil
}

func (k Keeper) DecreaseRedeemDenomAmount(ctx sdk.Context, amount sdk.Coin) error {
	redeemAssetBalance := k.GetAssetBalance(ctx, amount.Denom)
	decreasedAmount, err := redeemAssetBalance.SafeSub(amount)
	if err != nil {
		return err
	}

	k.SetAssetBalance(ctx, decreasedAmount)
	return nil
}

// TODO: implement
// func (k Keeper) IncreaseRedeemDenomAmount(ctx sdk.Context, amount sdk.Coin) error {
// 	redeemAssetBalance := k.GetAssetBalance(ctx, amount.Denom)
// 	increasedAmount := redeemAssetBalance.Add(amount)

// 	err := k.SetAssetBalance(ctx, increasedAmount)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// Initial Liquidity Provider Token Supply is determined in following formulas
// initial_lp_token_price = Σ target_weight_of_ith_asset * price_of_ith_asset
// pool_marketcap = price_of_ith_asset * amount_of_ith_deopsited_asset
// initial_lp_supply = pool_marketcap / initial_lp_token_price
func (k Keeper) InitialLiquidityProviderTokenSupply(ctx sdk.Context, assetPrice *pftypes.CurrentPrice, assetMarketCap sdk.Dec, depositDenom string) (sdk.Coin, error) {
	assetInfo := k.GetPoolAssetByDenom(ctx, depositDenom)
	initialLPTokenPrice := assetPrice.Price.Mul(assetInfo.TargetWeight)
	initialLPTokenSupply := assetMarketCap.Quo(initialLPTokenPrice)

	return sdk.NewCoin(types.LiquidityProviderTokenDenom, initialLPTokenSupply.TruncateInt()),
		nil
}

func (k Keeper) MintLiquidityProviderToken(ctx sdk.Context, msg *types.MsgMintLiquidityProviderToken) error {
	depositor := msg.Sender.AccAddress()

	params := k.GetParams(ctx)
	// check if the deposit denom is valid and amount is positive
	if !types.IsValidDepositForPool(msg.Amount, params.PoolParams.AcceptedAssets) {
		return fmt.Errorf("invalid deposit token: %s", msg.Amount.Denom)
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, sdk.Coins{msg.Amount}); err != nil {
		return err
	}

	currentBalance := k.GetAssetBalance(ctx, msg.Amount.Denom)
	targetBalance, err := k.GetAssetTargetAmount(ctx, msg.Amount.Denom)
	if err != nil {
		return err
	}
	fee := types.CalculateMintFee(currentBalance, targetBalance, msg.Amount, params.PoolParams.BaseLptMintFee)

	// TODO: integrate into ecosystem-incentive module
	// temporarily, send mint fee to fee pool of the derivatives module
	if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.DerivativeFeeCollector, sdk.Coins{fee}); err != nil {
		return err
	}

	deposit, err := msg.Amount.SafeSub(fee)
	mintingDLP, err := k.DetermineMintingLPTokenAmount(ctx, deposit)
	if err != nil {
		return err
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{mintingDLP}); err != nil {
		return err
	}

	// send to user
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.Coins{mintingDLP}); err != nil {
		return err
	}

	k.DepositPoolAsset(ctx, depositor, deposit)
	return nil
}

func (k Keeper) BurnLiquidityProviderToken(ctx sdk.Context, msg *types.MsgBurnLiquidityProviderToken) error {
	// todo:check validator address,amount,redeem denom
	// todo: use CacheCtx
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

	if redeemFee.IsPositive() {
		// send redeem fee to fee pool
		err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.DerivativeFeeCollector, sdk.Coins{redeemFee})
		if err != nil {
			return err
		}
	}

	// send redeem amount to the user
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.Coins{redeemAmount})
	if err != nil {
		return err
	}

	err = k.BurnCoin(ctx, sender, sdk.NewCoin(types.LiquidityProviderTokenDenom, amount))
	if err != nil {
		return err
	}

	err = k.CollectedFee(ctx, redeemFee)
	if err != nil {
		return err
	}

	err = k.DecreaseRedeemDenomAmount(ctx, redeemAmount.Add(redeemFee))
	if err != nil {
		return err
	}

	// todo emit event
	return nil
}

func (k Keeper) BurnCoin(ctx sdk.Context, burner sdk.AccAddress, amount sdk.Coin) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, burner, types.ModuleName, sdk.Coins{amount})
	if err != nil {
		return err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{amount})
	if err != nil {
		return err
	}
	return nil
}

// collect fee
func (k Keeper) CollectedFee(ctx sdk.Context, fee sdk.Coin) error {
	// todo: implement
	// LP fee = 70%
	// Protocol fee = 30%
	// fee.Amount.Mul(sdk.NewInt(0.7))
	// k.IncreaseRedeemDenomAmount(ctx, fee)
	return nil
}
