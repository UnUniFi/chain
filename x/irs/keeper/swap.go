package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) SwapUtToPt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, tokenIn sdk.Coin) error {
	tokenOutDenom := pool.PoolAssets[0].Denom
	if tokenOutDenom == tokenIn.Denom {
		tokenOutDenom = pool.PoolAssets[1].Denom
	}
	_, err := k.SwapExactAmountIn(ctx, sender, pool, tokenIn, tokenOutDenom, sdk.ZeroInt(), pool.SwapFee)
	return err
}

func (k Keeper) SwapPtToUt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, tokenIn sdk.Coin) error {
	tokenOutDenom := pool.PoolAssets[0].Denom
	if tokenOutDenom == tokenIn.Denom {
		tokenOutDenom = pool.PoolAssets[1].Denom
	}
	_, err := k.SwapExactAmountIn(ctx, sender, pool, tokenIn, tokenOutDenom, sdk.ZeroInt(), pool.SwapFee)
	return err
}

func (k Keeper) SwapExactAmountIn(
	ctx sdk.Context,
	sender sdk.AccAddress,
	pool types.TranchePool,
	tokenIn sdk.Coin,
	tokenOutDenom string,
	tokenOutMinAmount sdk.Int,
	swapFee sdk.Dec,
) (tokenOutAmount sdk.Int, err error) {
	if tokenIn.Denom == tokenOutDenom {
		return sdk.Int{}, errors.New("cannot trade the same denomination in and out")
	}

	// Executes the swap in the pool and stores the output. Updates pool assets but
	// does not actually transfer any tokens to or from the pool.
	tokenOutCoin, err := pool.SwapOutAmtGivenIn(ctx, tokenIn, tokenOutDenom, swapFee)
	if err != nil {
		return sdk.Int{}, err
	}

	tokenOutAmount = tokenOutCoin.Amount

	if !tokenOutAmount.IsPositive() {
		return sdk.Int{}, sdkerrors.Wrapf(types.ErrInvalidMathApprox, "token amount must be positive")
	}

	if tokenOutAmount.LT(tokenOutMinAmount) {
		return sdk.Int{}, sdkerrors.Wrapf(types.ErrLimitMinAmount, "%s token is lesser than the minimum amount", tokenOutDenom)
	}

	// Settles balances between the tx sender and the pool to match the swap that was executed earlier.
	// Also emits a swap event and updates related liquidity metrics.
	err = k.UpdatePoolForSwap(ctx, pool, sender, tokenIn, tokenOutCoin)
	if err != nil {
		return sdk.Int{}, err
	}

	// Subtract swap out fee from the token out amount.
	return tokenOutAmount, nil
}

func (k Keeper) UpdatePoolForSwap(
	ctx sdk.Context,
	pool types.TranchePool,
	sender sdk.AccAddress,
	tokenIn sdk.Coin,
	tokenOut sdk.Coin,
) error {
	tokensIn := sdk.Coins{tokenIn}
	k.SetTranchePool(ctx, pool)

	poolAddr := types.GetVaultModuleAddress(pool)
	err := k.bankKeeper.SendCoins(ctx, sender, poolAddr, tokensIn)
	if err != nil {
		return err
	}

	return nil
}
