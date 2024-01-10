package keeper

import (
	"errors"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/irs/types"
)

// SwapPoolTokens swaps tokens in a pool. UT => PT or PT => UT
func (k Keeper) SwapPoolTokens(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, tokenIn sdk.Coin) (sdk.Coin, error) {
	var tokenOutDenom string
	if tokenIn.Denom == pool.PoolAssets[0].Denom {
		tokenOutDenom = pool.PoolAssets[1].Denom
	} else if tokenIn.Denom == pool.PoolAssets[1].Denom {
		tokenOutDenom = pool.PoolAssets[0].Denom
	} else {
		return sdk.Coin{}, types.ErrInvalidDepositDenom
	}
	tokenOutAmount, err := k.SwapExactAmountIn(ctx, sender, pool, tokenIn, tokenOutDenom, sdk.ZeroInt(), pool.SwapFee)
	if err != nil {
		return sdk.Coin{}, err
	}
	return sdk.NewCoin(tokenOutDenom, tokenOutAmount), nil
}

// SimulateSwapPoolTokens simulates a swap in a pool & return TokenOut Amount value. UT => PT or PT => UT
func (k Keeper) SimulateSwapPoolTokens(ctx sdk.Context, pool types.TranchePool, tokenIn sdk.Coin) (sdk.Coin, error) {
	var tokenOutDenom string
	if tokenIn.Denom == pool.PoolAssets[0].Denom {
		tokenOutDenom = pool.PoolAssets[1].Denom
	} else if tokenIn.Denom == pool.PoolAssets[1].Denom {
		tokenOutDenom = pool.PoolAssets[0].Denom
	} else {
		return sdk.Coin{}, types.ErrInvalidDepositDenom
	}
	tokenOutAmount, err := k.CalculateSwapExactAmountIn(ctx, &pool, tokenIn, tokenOutDenom, sdk.ZeroInt(), pool.SwapFee)
	if err != nil {
		return sdk.Coin{}, err
	}
	return sdk.NewCoin(tokenOutDenom, tokenOutAmount), nil
}

func (k Keeper) SwapExactAmountIn(
	ctx sdk.Context,
	sender sdk.AccAddress,
	pool types.TranchePool,
	tokenIn sdk.Coin,
	tokenOutDenom string,
	tokenOutMinAmount math.Int,
	swapFee sdk.Dec,
) (tokenOutAmount math.Int, err error) {
	if tokenIn.Denom == tokenOutDenom {
		return math.Int{}, errors.New("cannot trade the same denomination in and out")
	}

	// check sender balance first
	poolAddr := types.GetLiquidityPoolModuleAddress(pool)
	tokensIn := sdk.Coins{tokenIn}
	err = k.bankKeeper.SendCoins(ctx, sender, poolAddr, tokensIn)
	if err != nil {
		return math.Int{}, err
	}

	tokenOutAmount, err = k.CalculateSwapExactAmountIn(ctx, &pool, tokenIn, tokenOutDenom, tokenOutMinAmount, swapFee)
	if err != nil {
		return math.Int{}, err
	}

	// Send out amount of tokens to the sender
	tokensOut := sdk.Coins{sdk.NewCoin(tokenOutDenom, tokenOutAmount)}
	err = k.bankKeeper.SendCoins(ctx, poolAddr, sender, tokensOut)
	if err != nil {
		return math.Int{}, err
	}

	// Settles balances between the tx sender and the pool to match the swap that was executed earlier.
	// Also emits a swap event and updates related liquidity metrics.
	k.SetTranchePool(ctx, pool)

	// Subtract swap out fee from the token out amount.
	return tokenOutAmount, nil
}

func (k Keeper) CalculateSwapExactAmountIn(
	ctx sdk.Context,
	pool *types.TranchePool,
	tokenIn sdk.Coin,
	tokenOutDenom string,
	tokenOutMinAmount math.Int,
	swapFee sdk.Dec,
) (tokenOutAmount math.Int, err error) {
	// Executes the swap in the pool and stores the output. Updates pool assets but
	// does not actually transfer any tokens to or from the pool.
	tokenOutCoin, err := pool.SwapOutAmtGivenIn(ctx, tokenIn, tokenOutDenom, swapFee)
	if err != nil {
		return math.Int{}, err
	}

	tokenOutAmount = tokenOutCoin.Amount

	if !tokenOutAmount.IsPositive() {
		return math.Int{}, sdkerrors.Wrapf(types.ErrInvalidMathApprox, "token amount must be positive")
	}

	if tokenOutAmount.LT(tokenOutMinAmount) {
		return math.Int{}, sdkerrors.Wrapf(types.ErrLimitMinAmount, "%s token is lesser than the minimum amount", tokenOutDenom)
	}

	return tokenOutAmount, nil
}
