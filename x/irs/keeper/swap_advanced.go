package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) SwapUtToYt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, requiredYtAmount sdk.Int, tokenIn sdk.Coin) error {
	// Check if TokenIn is enough to cover to payback loan
	depositInfo := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	if tokenIn.Denom != depositInfo.Denom {
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
	poolAddr := types.GetVaultModuleAddress(pool)
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

func (k Keeper) CalculateRequiredUtSwapToYt(ctx sdk.Context, pool types.TranchePool, requiredYtAmount sdk.Int) (sdk.Coin, error) {
	depositInfo := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	loan := sdk.NewCoin(depositInfo.Denom, requiredYtAmount)
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
	requiredUt := loan.Sub(estimatedUt)
	return requiredUt, nil
}

func (k Keeper) SwapYtToUt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, requiredUtAmount sdk.Int, token sdk.Coin) error {
	// Internally RedeemYtAtMaturity or RedeemPtYtPair

	// If matured, unstake from strategy
	// Else, redeem PT & YT pair
	if uint64(ctx.BlockTime().Unix()) > pool.StartTime+pool.Maturity {
		redeemAmount, err := k.CalculateRedeemYtAmount(ctx, pool, token)
		if err != nil {
			return err
		}
		if redeemAmount.LT(requiredUtAmount) {
			return types.ErrInsufficientFunds
		}
		err = k.RedeemYtAtMaturity(ctx, sender, pool, token)
		if err != nil {
			return err
		}
	} else {
		depositInfo := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
		estimateSwapPt, err := k.SimulateSwapPoolTokens(ctx, pool, sdk.NewCoin(depositInfo.Denom, token.Amount))
		if err != nil {
			return err
		}

		// 1. Take PT loan from IRS vault account (pool => sender)
		poolAddr := types.GetVaultModuleAddress(pool)
		loanPtAmount, err := k.CalculateRedeemRequiredPtAmount(ctx, pool, token.Amount)
		if err != nil {
			return err
		}
		if estimateSwapPt.Amount.LT(loanPtAmount) {
			return types.ErrInsufficientFunds
		}

		ptDenom := types.PtDenom(pool)
		loan := sdk.NewCoin(ptDenom, loanPtAmount)
		err = k.bankKeeper.SendCoins(ctx, poolAddr, sender, sdk.NewCoins(loan))
		if err != nil {
			return err
		}

		// 2. Redeem PT & YT pair
		err = k.RedeemPtYtPair(ctx, sender, pool, token.Amount, sdk.NewCoins(token, loan))
		if err != nil {
			return err
		}

		// 3. Swap UT to PT
		afterSwapPt, err := k.SwapPoolTokens(ctx, sender, pool, sdk.NewCoin(depositInfo.Denom, token.Amount))
		if err != nil {
			return err
		}

		// 4. Payback loan
		err = k.bankKeeper.SendCoins(ctx, sender, poolAddr, sdk.NewCoins(loan))
		if err != nil {
			return err
		}

		// 5. Swap rest PT to UT
		_, err = k.SwapPoolTokens(ctx, sender, pool, afterSwapPt.Sub(loan))
		if err != nil {
			return err
		}
	}
	return nil
}
