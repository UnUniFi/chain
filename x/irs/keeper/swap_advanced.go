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
	err := k.CheckEnoughUtTokenIn(ctx, pool, tokenIn, loan)
	if err != nil {
		return err
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
	err = k.SwapPoolTokens(ctx, sender, pool, sdk.NewCoin(ptDenom, ptAmount))
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

func (k Keeper) CheckEnoughUtTokenIn(ctx sdk.Context, pool types.TranchePool, tokenIn sdk.Coin, loan sdk.Coin) error {
	ptDenom := types.PtDenom(pool)
	// estimation 2. PT amount to mint
	estimatedPtAmount, err := k.CalculateMintPtAmount(ctx, pool, loan)
	if err != nil {
		return err
	}
	// estimation 3. UT amount to get by selling PT
	estimatedUt, err := k.SimulateSwapPoolTokens(ctx, pool, sdk.NewCoin(ptDenom, estimatedPtAmount))
	if err != nil {
		return err
	}
	// Check if estimated UT + TokenIn is enough to payback loan
	if estimatedUt.Add(tokenIn).IsLT(loan) {
		return types.ErrInsufficientFunds
	}
	return nil
}

// TODO:
func (k Keeper) SwapYtToUt() {
	// Internally combine SwapUtToPt and BurnPtYtPair

	// If matured, send required amount from unbonded from the share
	// Else
	// Put required amount of msg.PT from user wallet
	// Close position
	// Start redemption for strategy share
}
