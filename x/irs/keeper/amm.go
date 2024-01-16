package keeper

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) DepositToLiquidityPool(
	ctx sdk.Context,
	sender sdk.AccAddress,
	poolId uint64,
	shareOutAmount math.Int,
	tokenInMaxs sdk.Coins,
) (tokenIn sdk.Coins, sharesOut math.Int, err error) {
	// all pools handled within this method are pointer references, `JoinPool` directly updates the pools
	pool, found := k.GetTranchePool(ctx, poolId)
	if !found {
		return nil, sdk.ZeroInt(), types.ErrTrancheNotFound
	}

	// Ensure underlying token and pt token denoms are accurate when adding the liquidity for the first time
	depositInfo := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	utDenom := depositInfo.Denom
	ptDenom := types.PtDenom(pool)

	if !tokenInMaxs.AmountOf(ptDenom).IsPositive() {
		return nil, sdk.ZeroInt(), types.ErrNoPtDenomExists
	}

	if !tokenInMaxs.AmountOf(utDenom).IsPositive() {
		return nil, sdk.ZeroInt(), types.ErrNoUtDenomExists
	}

	// When liquidity is added to the empty pool
	if pool.TotalShares.IsZero() {
		pool.IncreaseLiquidity(types.OneShare, tokenInMaxs)
		err = k.applyJoinPoolStateChange(ctx, pool, sender, types.OneShare, tokenInMaxs)
		return tokenInMaxs, types.OneShare, err
	}

	// we do an abstract calculation on the lp liquidity coins needed to have
	// the designated amount of given shares of the pool without performing swap
	neededLpLiquidity, err := GetMaximalNoSwapLPAmount(ctx, pool, shareOutAmount)
	if err != nil {
		return nil, sdk.ZeroInt(), err
	}

	// check that needed lp liquidity does not exceed the given `tokenInMaxs` parameter. Return error if so.
	// if tokenInMaxs == 0, don't do this check.
	if tokenInMaxs.Len() != 0 {
		if !(neededLpLiquidity.DenomsSubsetOf(tokenInMaxs)) {
			return nil, sdk.ZeroInt(), types.ErrInSufficientTokenInMaxs
		} else if !(tokenInMaxs.DenomsSubsetOf(neededLpLiquidity)) {
			return nil, sdk.ZeroInt(), types.ErrInSufficientTokenInMaxs
		}
		if !(tokenInMaxs.IsAllGTE(neededLpLiquidity)) {
			return nil, sdk.ZeroInt(), types.ErrInSufficientTokenInMaxs
		}
	}

	sharesOut, err = pool.JoinPoolNoSwap(ctx, neededLpLiquidity, pool.SwapFee)
	if err != nil {
		return nil, sdk.ZeroInt(), err
	}
	// sanity check, don't return error as not worth halting the LP. We know its not too much.
	if sharesOut.LT(shareOutAmount) {
		ctx.Logger().Error(fmt.Sprintf("Expected to JoinPoolNoSwap >= %s shares, actually did %s shares",
			shareOutAmount, sharesOut))
	}

	err = k.applyJoinPoolStateChange(ctx, pool, sender, sharesOut, neededLpLiquidity)
	return neededLpLiquidity, sharesOut, err
}

// getMaximalNoSwapLPAmount returns the coins(lp liquidity) needed to get the specified amount of shares in the pool.
// Steps to getting the needed lp liquidity coins needed for the share of the pools are
// 1. calculate how much percent of the pool does given share account for(# of input shares / # of current total shares)
// 2. since we know how much % of the pool we want, iterate through all pool liquidity to calculate how much coins we need for
// each pool asset.
func getMaximalNoSwapLPAmount(ctx sdk.Context, pool types.TranchePool, shareOutAmount math.Int) (neededLpLiquidity sdk.Coins, err error) {
	totalSharesAmount := pool.TotalShares.Amount
	if totalSharesAmount.IsZero() {
		return sdk.Coins{}, sdkerrors.Wrapf(types.ErrInvalidMathApprox, "Total shares amount is zero")
	}
	// shareRatio is the desired number of shares, divided by the total number of
	// shares currently in the pool. It is intended to be used in scenarios where you want
	shareRatio := sdk.NewDecFromInt(shareOutAmount).QuoInt(totalSharesAmount)
	if shareRatio.LTE(sdk.ZeroDec()) {
		return sdk.Coins{}, sdkerrors.Wrapf(types.ErrInvalidMathApprox, "Too few shares out wanted. "+
			"(debug: getMaximalNoSwapLPAmount share ratio is zero or negative)")
	}

	poolLiquidity := pool.PoolAssets
	neededLpLiquidity = sdk.Coins{}

	for _, coin := range poolLiquidity {
		// (coin.Amt * shareRatio).Ceil()
		neededAmt := sdk.NewDecFromInt(coin.Amount).Mul(shareRatio).Ceil().RoundInt()
		neededCoin := sdk.Coin{Denom: coin.Denom, Amount: neededAmt}
		neededLpLiquidity = neededLpLiquidity.Add(neededCoin)
	}
	return neededLpLiquidity, nil
}

func GetMaximalNoSwapLPAmount(ctx sdk.Context, pool types.TranchePool, shareOutAmount math.Int) (neededLpLiquidity sdk.Coins, err error) {
	return getMaximalNoSwapLPAmount(ctx, pool, shareOutAmount)
}

func (k Keeper) WithdrawFromLiquidityPool(
	ctx sdk.Context,
	sender sdk.AccAddress,
	poolId uint64,
	shareInAmount math.Int,
	tokenOutMins sdk.Coins,
) (exitCoins sdk.Coins, err error) {
	pool, found := k.GetTranchePool(ctx, poolId)
	if !found {
		return sdk.Coins{}, types.ErrTrancheNotFound
	}

	totalSharesAmount := pool.GetTotalShares()
	if shareInAmount.GT(totalSharesAmount.Amount) {
		return sdk.Coins{}, types.ErrInvalidTotalShares
	} else if shareInAmount.LTE(sdk.ZeroInt()) {
		return sdk.Coins{}, types.ErrInvalidTotalShares
	}
	exitFee := pool.ExitFee
	exitCoins, err = pool.ExitPool(ctx, shareInAmount, exitFee)
	if err != nil {
		return sdk.Coins{}, err
	}
	if !tokenOutMins.DenomsSubsetOf(exitCoins) || tokenOutMins.IsAnyGT(exitCoins) {
		return sdk.Coins{}, types.ErrInsufficientExitCoins
	}

	err = k.applyExitPoolStateChange(ctx, pool, sender, shareInAmount, exitCoins)
	if err != nil {
		return sdk.Coins{}, err
	}

	return exitCoins, nil
}

func (k Keeper) applyJoinPoolStateChange(ctx sdk.Context, pool types.TranchePool, joiner sdk.AccAddress, numShares math.Int, joinCoins sdk.Coins) error {
	err := k.bankKeeper.SendCoins(ctx, joiner, types.GetLiquidityPoolModuleAddress(pool), joinCoins)
	if err != nil {
		return err
	}

	err = k.MintPoolShareToAccount(ctx, pool, joiner, numShares)
	if err != nil {
		return err
	}

	k.SetTranchePool(ctx, pool)

	return nil
}

func (k Keeper) applyExitPoolStateChange(ctx sdk.Context, pool types.TranchePool, exiter sdk.AccAddress, numShares math.Int, exitCoins sdk.Coins) error {
	err := k.bankKeeper.SendCoins(ctx, types.GetLiquidityPoolModuleAddress(pool), exiter, exitCoins)
	if err != nil {
		return err
	}

	err = k.BurnPoolShareFromAccount(ctx, pool, exiter, numShares)
	if err != nil {
		return err
	}

	k.SetTranchePool(ctx, pool)

	return nil
}

// MintPoolShareToAccount attempts to mint shares of a GAMM denomination to the
// specified address returning an error upon failure. Shares are minted using
// the x/gamm module account.
func (k Keeper) MintPoolShareToAccount(ctx sdk.Context, pool types.TranchePool, addr sdk.AccAddress, amount math.Int) error {
	amt := sdk.NewCoins(sdk.NewCoin(types.LsDenom(pool), amount))

	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amt)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, amt)
	if err != nil {
		return err
	}

	return nil
}

// BurnPoolShareFromAccount burns `amount` of the given pools shares held by `addr`.
func (k Keeper) BurnPoolShareFromAccount(ctx sdk.Context, pool types.TranchePool, addr sdk.AccAddress, amount math.Int) error {
	amt := sdk.Coins{
		sdk.NewCoin(types.LsDenom(pool), amount),
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, amt)
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, amt)
	if err != nil {
		return err
	}

	return nil
}
