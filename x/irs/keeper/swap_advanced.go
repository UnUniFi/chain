package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) SwapUtToYt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, requiredYtAmount math.Int, tokenIn sdk.Coin) error {
	// Check if TokenIn is enough to cover to payback loan
	if tokenIn.Denom != pool.Denom {
		return types.ErrInvalidDepositDenom
	}
	loan := sdk.NewCoin(tokenIn.Denom, requiredYtAmount)
	ptDenom := types.PtDenom(pool)
	requiredUt, err := k.CalculateRequiredUtSwapToYt(ctx, pool, requiredYtAmount)
	if err != nil {
		return err
	}
	if tokenIn.Amount.LT(requiredUt.Amount) {
		return types.ErrInsufficientFunds
	}

	// 1. Take loan from IRS vault account (pool => sender)
	poolAddr := types.GetLiquidityPoolModuleAddress(pool)
	err = k.bankKeeper.SendCoins(ctx, poolAddr, sender, sdk.NewCoins(loan))
	if err != nil {
		return err
	}

	// 2. Mint Pt and Yt
	ptAmount, err := k.MintPtYtPair(ctx, sender, pool, loan)
	if err != nil {
		return err
	}

	// 3. Sell minted PT amount for underlying token
	_, err = k.SwapPoolTokens(ctx, sender, pool, sdk.NewCoin(ptDenom, ptAmount))
	if err != nil {
		return err
	}

	// 4. Payback loan
	err = k.bankKeeper.SendCoins(ctx, sender, poolAddr, sdk.NewCoins(loan))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) CalculateRequiredUtSwapToYt(ctx sdk.Context, pool types.TranchePool, requiredYtAmount math.Int) (sdk.Coin, error) {
	loan := sdk.NewCoin(pool.Denom, requiredYtAmount)
	ptDenom := types.PtDenom(pool)
	// estimation 2. PT amount to mint
	estimatedPtAmount, err := k.CalculateMintPtAmount(ctx, pool, loan)
	if err != nil {
		return sdk.Coin{}, err
	}
	// estimation 3. UT amount to get by selling PT
	estimatedUt, err := k.SimulateSwapPoolTokens(ctx, pool, sdk.NewCoin(ptDenom, estimatedPtAmount))
	if err != nil {
		return sdk.Coin{}, err
	}
	if loan.IsLT(estimatedUt) {
		return sdk.NewInt64Coin(pool.Denom, 0), nil
	}
	requiredUt := loan.Sub(estimatedUt)
	return requiredUt, nil
}

func (k Keeper) CalculateSwapUtToYt(ctx sdk.Context, pool types.TranchePool, tokenIn sdk.Coin) (sdk.Coin, error) {
	if tokenIn.Denom != pool.Denom {
		return sdk.Coin{}, types.ErrInvalidDepositDenom
	}
	ytDenom := types.YtDenom(pool)
	ytSupply := k.bankKeeper.GetSupply(ctx, ytDenom)
	moduleAddr := types.GetVaultModuleAddress(pool)
	amountFromStrategy, err := k.GetAmountFromStrategy(ctx, moduleAddr, pool.StrategyContract)
	if err != nil {
		return sdk.Coin{}, err
	}
	if amountFromStrategy.IsZero() {
		return sdk.Coin{}, types.ErrInvalidMathApprox
	}
	swapCoin := sdk.NewCoin(tokenIn.Denom, sdk.NewInt(1_000_000))
	pt, err := k.SimulateSwapPoolTokens(ctx, pool, swapCoin)
	if err != nil {
		return sdk.Coin{}, err
	}
	ptRate := sdk.NewDecFromInt(pt.Amount).QuoInt(swapCoin.Amount)
	if ptRate.IsZero() {
		return sdk.Coin{}, types.ErrInvalidMathApprox
	}
	ytAmount := sdk.NewDecFromInt(tokenIn.Amount).MulInt(ytSupply.Amount).QuoInt(amountFromStrategy).Quo(ptRate).TruncateInt()
	return sdk.NewCoin(ytDenom, ytAmount), nil
}

func (k Keeper) SwapYtToUt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, requiredUtAmount math.Int, tokens sdk.Coins) error {
	err := k.RedeemPtYtPair(ctx, sender, pool, requiredUtAmount, tokens)
	if err != nil {
		return err
	}
	return nil
}

// // TODO: This implementation is better if there is no Redeem time lag
// func (k Keeper) SwapYtToUt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, requiredUtAmount math.Int, token sdk.Coin) error {
// 	depositInfo := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
// 	redeemUtAmount, err := k.CalculateRedeemAmount(ctx, pool, token)
// 	if err != nil {
// 		return err
// 	}
// 	redeemUt := sdk.NewCoin(depositInfo.Denom, redeemUtAmount)
// 	estimateSwapPt, err := k.SimulateSwapPoolTokens(ctx, pool, redeemUt)
// 	if err != nil {
// 		return err
// 	}

// 	// 1. Take PT loan from IRS vault account (pool => sender)
// 	poolAddr := types.GetVaultModuleAddress(pool)
// 	ptDenom := types.PtDenom(pool)
// 	loanPtAmount, err := k.CalculateRedeemRequiredAmount(ctx, pool, redeemUtAmount, ptDenom)
// 	if err != nil {
// 		return err
// 	}
// 	if estimateSwapPt.Amount.LT(loanPtAmount) {
// 		return types.ErrInsufficientFunds
// 	}

// 	loan := sdk.NewCoin(ptDenom, loanPtAmount)
// 	err = k.bankKeeper.SendCoins(ctx, poolAddr, sender, sdk.NewCoins(loan))
// 	if err != nil {
// 		return err
// 	}

// 	// 2. Redeem PT & YT pair
// 	// TODO: it contains time lag between 2 and 3
// 	err = k.RedeemPtYtPair(ctx, sender, pool, redeemUtAmount, sdk.NewCoins(token, loan))
// 	if err != nil {
// 		return err
// 	}

// 	// 3. Swap UT to PT
// 	afterSwapPt, err := k.SwapPoolTokens(ctx, sender, pool, redeemUt)
// 	if err != nil {
// 		return err
// 	}

// 	// 4. Payback loan
// 	err = k.bankKeeper.SendCoins(ctx, sender, poolAddr, sdk.NewCoins(loan))
// 	if err != nil {
// 		return err
// 	}

// 	// 5. Swap rest PT to UT
// 	_, err = k.SwapPoolTokens(ctx, sender, pool, afterSwapPt.Sub(loan))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
