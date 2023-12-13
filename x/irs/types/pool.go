package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/osmomath"
)

// JoinPoolNoSwap calculates the number of shares needed for an all-asset join given tokensIn with spreadFactor applied.
// It updates the liquidity if the pool is joined successfully. If not, returns error.
func (p *TranchePool) JoinPoolNoSwap(ctx sdk.Context, tokensIn sdk.Coins, spreadFactor sdk.Dec) (numShares sdk.Int, err error) {
	numShares, tokensJoined, err := p.CalcJoinPoolNoSwapShares(ctx, tokensIn, spreadFactor)
	if err != nil {
		return sdk.Int{}, err
	}

	// update pool with the calculated share and liquidity needed to join pool
	p.IncreaseLiquidity(numShares, tokensJoined)
	return numShares, nil
}

// ensureDenomInPool check to make sure the input denoms exist in the provided pool asset map
func ensureDenomInPool(poolAssetsByDenom sdk.Coins, tokensIn sdk.Coins) error {
	for _, coin := range tokensIn {
		if !poolAssetsByDenom.AmountOf(coin.Denom).IsPositive() {
			return ErrDenomNotFoundInPool
		}
	}

	return nil
}

// MaximalExactRatioJoin calculates the maximal amount of tokens that can be joined whilst maintaining pool asset's ratio
// returning the number of shares that'd be and how many coins would be left over.
//
//	e.g) suppose we have a pool of 10 foo tokens and 10 bar tokens, with the total amount of 100 shares.
//		 if `tokensIn` provided is 1 foo token and 2 bar tokens, `MaximalExactRatioJoin`
//		 would be returning (10 shares, 1 bar token, nil)
//
// This can be used when `tokensIn` are not guaranteed the same ratio as assets in the pool.
// Calculation for this is done in the following steps.
//  1. iterate through all the tokens provided as an argument, calculate how much ratio it accounts for the asset in the pool
//  2. get the minimal share ratio that would work as the benchmark for all tokens.
//  3. calculate the number of shares that could be joined (total share * min share ratio), return the remaining coins
func MaximalExactRatioJoin(p *TranchePool, ctx sdk.Context, tokensIn sdk.Coins) (numShares sdk.Int, remCoins sdk.Coins, err error) {
	coinShareRatios := make([]sdk.Dec, len(tokensIn))
	minShareRatio := sdk.MaxSortableDec
	maxShareRatio := sdk.ZeroDec()

	poolLiquidity := sdk.Coins(p.PoolAssets)
	totalShares := p.GetTotalShares().Amount

	for i, coin := range tokensIn {
		// Note: QuoInt implements floor division, unlike Quo
		// This is because it calls the native golang routine big.Int.Quo
		// https://pkg.go.dev/math/big#Int.Quo
		shareRatio := sdk.NewDecFromInt(coin.Amount).QuoInt(poolLiquidity.AmountOf(coin.Denom))
		if shareRatio.LT(minShareRatio) {
			minShareRatio = shareRatio
		}
		if shareRatio.GT(maxShareRatio) {
			maxShareRatio = shareRatio
		}
		coinShareRatios[i] = shareRatio
	}

	if minShareRatio.Equal(sdk.MaxSortableDec) {
		return numShares, remCoins, errors.New("unexpected error in MaximalExactRatioJoin")
	}

	remCoins = sdk.Coins{}
	// critically we round down here (TruncateInt), to ensure that the returned LP shares
	// are always less than or equal to % liquidity added.
	numShares = minShareRatio.MulInt(totalShares).TruncateInt()

	// if we have multiple share values, calculate remainingCoins
	if !minShareRatio.Equal(maxShareRatio) {
		// we have to calculate remCoins
		for i, coin := range tokensIn {
			// if coinShareRatios[i] == minShareRatio, no remainder
			if coinShareRatios[i].Equal(minShareRatio) {
				continue
			}

			usedAmount := minShareRatio.MulInt(poolLiquidity.AmountOfNoDenomValidation(coin.Denom)).Ceil().TruncateInt()
			newAmt := coin.Amount.Sub(usedAmount)
			// if newAmt is non-zero, add to RemCoins. (It could be zero due to rounding)
			if !newAmt.IsZero() {
				remCoins = remCoins.Add(sdk.Coin{Denom: coin.Denom, Amount: newAmt})
			}
		}
	}

	return numShares, remCoins, nil
}

// CalcJoinPoolNoSwapShares calculates the number of shares created to execute an all-asset pool join with the provided amount of `tokensIn`.
// The input tokens must contain the same tokens as in the pool.
//
// Returns the number of shares created, the amount of coins actually joined into the pool, (in case of not being able to fully join),
// and the remaining tokens in `tokensIn` after joining. If an all-asset join is not possible, returns an error.
//
// Since CalcJoinPoolNoSwapShares is non-mutative, the steps for updating pool shares / liquidity are
// more complex / don't just alter the state.
// We should simplify this logic further in the future using multi-join equations.
func (p *TranchePool) CalcJoinPoolNoSwapShares(ctx sdk.Context, tokensIn sdk.Coins, spreadFactor sdk.Dec) (numShares sdk.Int, tokensJoined sdk.Coins, err error) {
	err = ensureDenomInPool(p.PoolAssets, tokensIn)
	if err != nil {
		return sdk.ZeroInt(), sdk.NewCoins(), err
	}

	// ensure that there aren't too many or too few assets in `tokensIn`
	if tokensIn.Len() != len(p.PoolAssets) {
		return sdk.ZeroInt(), sdk.NewCoins(), errors.New("no-swap joins require LP'ing with all assets in pool")
	}

	// execute a no-swap join with as many tokens as possible given a perfect ratio:
	// * numShares is how many shares are perfectly matched.
	// * remainingTokensIn is how many coins we have left to join that have not already been used.
	numShares, remainingTokensIn, err := MaximalExactRatioJoin(p, ctx, tokensIn)
	if err != nil {
		return sdk.ZeroInt(), sdk.NewCoins(), err
	}

	// ensure that no more tokens have been joined than is possible with the given `tokensIn`
	tokensJoined = tokensIn.Sub(remainingTokensIn...)
	if tokensJoined.IsAnyGT(tokensIn) {
		return sdk.ZeroInt(), sdk.NewCoins(), errors.New("an error has occurred, more coins joined than token In")
	}

	return numShares, tokensJoined, nil
}

func (p *TranchePool) IncreaseLiquidity(sharesOut sdk.Int, coinsIn sdk.Coins) {
	p.PoolAssets = sdk.Coins(p.PoolAssets).Add(coinsIn...)
	p.TotalShares.Amount = p.TotalShares.Amount.Add(sharesOut)
}

func (p *TranchePool) ExitPool(ctx sdk.Context, exitingShares sdk.Int, exitFee sdk.Dec) (exitingCoins sdk.Coins, err error) {
	exitingCoins, err = p.CalcExitPoolCoinsFromShares(ctx, exitingShares, exitFee)
	if err != nil {
		return sdk.Coins{}, err
	}

	if err := p.exitPool(ctx, exitingCoins, exitingShares); err != nil {
		return sdk.Coins{}, err
	}

	return exitingCoins, nil
}

// exitPool exits the pool given exitingCoins and exitingShares.
// updates the pool's liquidity and totalShares.
func (p *TranchePool) exitPool(ctx sdk.Context, exitingCoins sdk.Coins, exitingShares sdk.Int) error {
	balances := sdk.Coins(p.PoolAssets).Sub(exitingCoins...)
	p.PoolAssets = balances

	totalShares := p.TotalShares.Amount
	p.TotalShares = sdk.NewCoin(p.TotalShares.Denom, totalShares.Sub(exitingShares))

	return nil
}

func (p *TranchePool) CalcExitPoolCoinsFromShares(ctx sdk.Context, exitingShares sdk.Int, exitFee sdk.Dec) (exitedCoins sdk.Coins, err error) {
	return CalcExitPool(ctx, p, exitingShares, exitFee)
}

// CalcExitPool returns how many tokens should come out, when exiting k LP shares against a "standard" CFMM
func CalcExitPool(ctx sdk.Context, pool *TranchePool, exitingShares sdk.Int, exitFee sdk.Dec) (sdk.Coins, error) {
	totalShares := pool.TotalShares.Amount
	if exitingShares.GT(totalShares) {
		return sdk.Coins{}, ErrInvalidTotalShares
	}

	// refundedShares = exitingShares * (1 - exit fee)
	// with 0 exit fee optimization
	var refundedShares sdk.Dec
	if !exitFee.IsZero() {
		// exitingShares * (1 - exit fee)
		oneSubExitFee := sdk.OneDec().Sub(exitFee)
		refundedShares = oneSubExitFee.MulIntMut(exitingShares)
	} else {
		refundedShares = sdk.NewDecFromInt(exitingShares)
	}

	shareOutRatio := refundedShares.QuoInt(totalShares)
	// exitedCoins = shareOutRatio * pool liquidity
	exitedCoins := sdk.Coins{}
	poolLiquidity := pool.PoolAssets

	for _, asset := range poolLiquidity {
		// round down here, due to not wanting to over-exit
		exitAmt := shareOutRatio.MulInt(asset.Amount).TruncateInt()
		if exitAmt.LTE(sdk.ZeroInt()) {
			continue
		}
		if exitAmt.GT(asset.Amount) {
			return sdk.Coins{}, errors.New("too many shares out")
		}
		exitedCoins = exitedCoins.Add(sdk.NewCoin(asset.Denom, exitAmt))
	}

	return exitedCoins, nil
}

// SwapOutAmtGivenIn is a mutative method for CalcOutAmtGivenIn, which includes the actual swap.
func (p *TranchePool) SwapOutAmtGivenIn(
	ctx sdk.Context,
	tokenIn sdk.Coin,
	tokenOutDenom string,
	swapFee sdk.Dec,
) (tokenOut sdk.Coin, err error) {
	balancerOutCoin, err := p.CalcOutAmtGivenIn(ctx, tokenIn, tokenOutDenom, swapFee)
	if err != nil {
		return sdk.Coin{}, err
	}

	err = p.applySwap(tokenIn, balancerOutCoin, sdk.ZeroDec(), swapFee)
	if err != nil {
		return sdk.Coin{}, err
	}
	return balancerOutCoin, nil
}

// CalcOutAmtGivenIn calculates tokens to be swapped out given the provided
// amount and fee deducted, using solveConstantFunctionInvariant.
func (p TranchePool) CalcOutAmtGivenIn(
	ctx sdk.Context,
	tokenIn sdk.Coin,
	tokenOutDenom string,
	swapFee sdk.Dec,
) (sdk.Coin, error) {
	tokenAmountInAfterFee := sdk.NewDecFromInt(tokenIn.Amount).Mul(sdk.OneDec().Sub(swapFee))
	// Pool balance of tokenIn and tokenOut
	poolAssetIn, poolAssetOut, err := p.parsePoolAssetsByDenoms(tokenIn.Denom, tokenOutDenom)
	if err != nil {
		return sdk.Coin{}, err
	}
	poolTokenInBalance := sdk.NewDecFromInt(poolAssetIn.Amount)
	poolTokenOutBalance := sdk.NewDecFromInt(poolAssetOut.Amount)
	// Pool balance of tokenIn + tokenAmountInAfterFee
	poolPostSwapInBalance := poolTokenInBalance.Add(tokenAmountInAfterFee)

	// deduct swapfee on the tokensIn
	// delta balanceOut is positive(tokens inside the pool decreases)
	t := sdk.OneDec().Sub(sdk.NewDec(ctx.BlockTime().Unix() - int64(p.StartTime)).Quo(sdk.NewDec(int64(p.Maturity))))
	if !t.IsPositive() {
		return sdk.Coin{}, sdkerrors.Wrapf(ErrTrancheAlreadyMatured, "tranche has been already matured")
	}
	if t.GTE(sdk.OneDec()) {
		return sdk.Coin{}, ErrInvalidTrancheStartTime
	}
	tokenAmountOut := solveConstantFunctionInvariant(
		t,
		poolTokenInBalance,
		poolPostSwapInBalance,
		poolTokenOutBalance,
	)

	// We ignore the decimal component, as we round down the token amount out.
	tokenAmountOutInt := tokenAmountOut.TruncateInt()
	if !tokenAmountOutInt.IsPositive() {
		return sdk.Coin{}, sdkerrors.Wrapf(ErrInvalidMathApprox, "token amount must be positive")
	}

	return sdk.NewCoin(tokenOutDenom, tokenAmountOutInt), nil
}

func (p *TranchePool) applySwap(tokenIn sdk.Coin, tokenOut sdk.Coin, swapFeeIn, swapFeeOut sdk.Dec) error {
	inTokensAfterFee := sdk.NewDecFromInt(tokenIn.Amount).Mul(sdk.OneDec().Sub(swapFeeIn)).TruncateInt()
	outTokensAfterFee := sdk.NewDecFromInt(tokenOut.Amount).Mul(sdk.OneDec().Sub(swapFeeOut)).TruncateInt()

	p.PoolAssets = sdk.Coins(p.PoolAssets).Add(sdk.NewCoin(tokenIn.Denom, inTokensAfterFee))
	p.PoolAssets = sdk.Coins(p.PoolAssets).Sub(sdk.NewCoin(tokenOut.Denom, outTokensAfterFee))

	return nil
}

func solveConstantFunctionInvariant(
	t,
	tokenBalanceFixedBefore,
	tokenBalanceFixedAfter,
	tokenBalanceUnknownBefore sdk.Dec,
) sdk.Dec {
	exp := osmomath.BigDecFromSDKDec(sdk.OneDec().Sub(t))
	// k = x1^(1-t) + y1^(1-t)
	x1 := osmomath.BigDecFromSDKDec(tokenBalanceFixedBefore)
	x1exp := x1.Power(exp)
	y1 := osmomath.BigDecFromSDKDec(tokenBalanceUnknownBefore)
	y1exp := y1.Power(exp)
	k := x1exp.Add(y1exp)
	// y2^(1-t) = k - x2^(1-t)
	y2exp := k.Sub(osmomath.BigDecFromSDKDec(tokenBalanceFixedAfter).Power(exp))
	// y2 = y2^(1-t)^(1/(1-t))
	if y2exp.IsNegative() {
		// TODO: return y1 or error
		// Plz verify that y can be 0 in other codes.
		return sdk.ZeroDec()
	}
	y2 := y2exp.Power(osmomath.BigDecFromSDKDec(sdk.OneDec()).Quo(exp))
	// TokenOut to be issued = y1 - y2
	return y1.Sub(y2).SDKDec()
}

func (p TranchePool) parsePoolAssetsByDenoms(tokenADenom, tokenBDenom string) (
	Aasset sdk.Coin, Basset sdk.Coin, err error,
) {
	Aasset, found1 := getPoolAssetByDenom(p.PoolAssets, tokenADenom)
	Basset, found2 := getPoolAssetByDenom(p.PoolAssets, tokenBDenom)

	if !found1 {
		return sdk.Coin{}, sdk.Coin{}, fmt.Errorf("(%s) does not exist in the pool", tokenADenom)
	}
	if !found2 {
		return sdk.Coin{}, sdk.Coin{}, fmt.Errorf("(%s) does not exist in the pool", tokenBDenom)
	}
	return Aasset, Basset, nil
}

func getPoolAssetByDenom(assets []sdk.Coin, denom string) (sdk.Coin, bool) {
	for _, asset := range assets {
		if asset.Denom == denom {
			return asset, true
		}
	}
	return sdk.Coin{}, false
}
