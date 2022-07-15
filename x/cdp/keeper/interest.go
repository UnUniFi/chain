package keeper

import (
	"fmt"
	"math"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/cdp/types"
)

var (
	scalingFactor  = 1e18
	secondsPerYear = 31536000
)

// AccumulateInterest calculates the new interest that has accrued for the input collateral type based on the total amount of principal
// that has been created with that collateral type and the amount of time that has passed since interest was last accumulated
func (k Keeper) AccumulateInterest(ctx sdk.Context, ctype string) error {
	previousAccrualTime, found := k.GetPreviousAccrualTime(ctx, ctype)
	if !found {
		k.SetPreviousAccrualTime(ctx, ctype, ctx.BlockTime())
		return nil
	}

	timeElapsed := int64(math.RoundToEven(
		ctx.BlockTime().Sub(previousAccrualTime).Seconds(),
	))
	if timeElapsed == 0 {
		return nil
	}

	totalPrincipalPrior := k.GetTotalPrincipal(ctx, ctype, types.DefaultStableDenom)
	if totalPrincipalPrior.IsZero() || totalPrincipalPrior.IsNegative() {
		k.SetPreviousAccrualTime(ctx, ctype, ctx.BlockTime())
		return nil
	}

	interestFactorPrior, foundInterestFactorPrior := k.GetInterestFactor(ctx, ctype)
	if !foundInterestFactorPrior {
		k.SetInterestFactor(ctx, ctype, sdk.OneDec())
		// set previous accrual time exit early because interest accumulated will be zero
		k.SetPreviousAccrualTime(ctx, ctype, ctx.BlockTime())
		return nil
	}

	borrowRateSpy := k.getFeeRate(ctx, ctype)
	if borrowRateSpy.Equal(sdk.OneDec()) {
		k.SetPreviousAccrualTime(ctx, ctype, ctx.BlockTime())
		return nil
	}
	interestFactor := CalculateInterestFactor(borrowRateSpy, sdk.NewInt(timeElapsed))
	interestAccumulated := (interestFactor.Mul(sdk.NewDecFromInt(totalPrincipalPrior))).RoundInt().Sub(totalPrincipalPrior)
	if interestAccumulated.IsZero() {
		// in the case accumulated interest rounds to zero, exit early without updating accrual time
		return nil
	}
	debtDenomMap := k.GetDebtDenomMap(ctx)
	err := k.MintDebtCoins(ctx, types.ModuleName, debtDenomMap[types.DefaultStableDenom], sdk.NewCoin(types.DefaultStableDenom, interestAccumulated))
	if err != nil {
		return err
	}

	dp, found := k.GetDebtParam(ctx, types.DefaultStableDenom)
	if !found {
		panic(fmt.Sprintf("Debt parameters for %s not found", types.DefaultStableDenom))
	}

	newFeesSurplus := interestAccumulated

	// mint surplus coins to the liquidator module account.
	if newFeesSurplus.IsPositive() {
		err := k.bankKeeper.MintCoins(ctx, types.LiquidatorMacc, sdk.NewCoins(sdk.NewCoin(dp.Denom, newFeesSurplus)))
		if err != nil {
			return err
		}
	}

	interestFactorNew := interestFactorPrior.Mul(interestFactor)
	totalPrincipalNew := totalPrincipalPrior.Add(interestAccumulated)

	k.SetTotalPrincipal(ctx, ctype, types.DefaultStableDenom, totalPrincipalNew)
	k.SetInterestFactor(ctx, ctype, interestFactorNew)
	k.SetPreviousAccrualTime(ctx, ctype, ctx.BlockTime())

	return nil
}

// CalculateInterestFactor calculates the simple interest scaling factor,
// which is equal to: (per-second interest rate ** number of seconds elapsed)
// Will return 1.000x, multiply by principal to get new principal with added interest
func CalculateInterestFactor(perSecondInterestRate sdk.Dec, secondsElapsed sdk.Int) sdk.Dec {
	scalingFactorUint := sdk.NewUint(uint64(scalingFactor))
	scalingFactorInt := sdk.NewInt(int64(scalingFactor))

	// Convert per-second interest rate to a uint scaled by 1e18
	interestMantissa := sdkmath.NewUintFromBigInt(perSecondInterestRate.MulInt(scalingFactorInt).RoundInt().BigInt())

	// Convert seconds elapsed to uint (*not scaled*)
	secondsElapsedUint := sdkmath.NewUintFromBigInt(secondsElapsed.BigInt())

	// Calculate the interest factor as a uint scaled by 1e18
	interestFactorMantissa := sdkmath.RelativePow(interestMantissa, secondsElapsedUint, scalingFactorUint)

	// Convert interest factor to an unscaled sdk.Dec
	return sdk.NewDecFromBigInt(interestFactorMantissa.BigInt()).QuoInt(scalingFactorInt)
}

// SynchronizeInterest updates the input cdp object to reflect the current accumulated interest, updates the cdp state in the store,
// and returns the updated cdp object
func (k Keeper) SynchronizeInterest(ctx sdk.Context, cdp types.Cdp) types.Cdp {
	globalInterestFactor, found := k.GetInterestFactor(ctx, cdp.Type)
	if !found {
		k.SetInterestFactor(ctx, cdp.Type, sdk.OneDec())
		cdp.InterestFactor = sdk.OneDec()
		cdp.FeesUpdated = ctx.BlockTime()
		k.SetCdp(ctx, cdp)
		return cdp
	}

	accumulatedInterest := k.CalculateNewInterest(ctx, cdp)
	prevAccrualTime, found := k.GetPreviousAccrualTime(ctx, cdp.Type)
	if !found {
		return cdp
	}
	if accumulatedInterest.IsZero() {
		// accumulated interest is zero if apy is zero or are if the total fees for all cdps round to zero
		if cdp.FeesUpdated.Equal(prevAccrualTime) {
			// if all fees are rounding to zero, don't update FeesUpdated
			return cdp
		}
		// if apy is zero, we need to update FeesUpdated
		cdp.FeesUpdated = prevAccrualTime
		k.SetCdp(ctx, cdp)
	}

	cdp.AccumulatedFees = cdp.AccumulatedFees.Add(accumulatedInterest)
	cdp.FeesUpdated = prevAccrualTime
	cdp.InterestFactor = globalInterestFactor
	collateralToDebtRatio := k.CalculateCollateralToDebtRatio(ctx, cdp.Collateral, cdp.Type, cdp.GetTotalPrincipal())
	k.UpdateCdpAndCollateralRatioIndex(ctx, cdp, collateralToDebtRatio)
	return cdp
}

// CalculateNewInterest returns the amount of interest that has accrued to the cdp since its interest was last synchronized
func (k Keeper) CalculateNewInterest(ctx sdk.Context, cdp types.Cdp) sdk.Coin {
	globalInterestFactor, found := k.GetInterestFactor(ctx, cdp.Type)
	if !found {
		return sdk.NewCoin(cdp.AccumulatedFees.Denom, sdk.ZeroInt())
	}
	cdpInterestFactor := globalInterestFactor.Quo(cdp.InterestFactor)
	if cdpInterestFactor.Equal(sdk.OneDec()) {
		return sdk.NewCoin(cdp.AccumulatedFees.Denom, sdk.ZeroInt())
	}
	accumulatedInterest := sdk.NewDecFromInt(cdp.GetTotalPrincipal().Amount).Mul(cdpInterestFactor).RoundInt().Sub(cdp.GetTotalPrincipal().Amount)
	return sdk.NewCoin(cdp.AccumulatedFees.Denom, accumulatedInterest)
}

// SynchronizeInterestForRiskyCdps synchronizes the interest for the slice of cdps with the lowest collateral:debt ratio
func (k Keeper) SynchronizeInterestForRiskyCdps(ctx sdk.Context, slice sdk.Int, targetRatio sdk.Dec, collateralType string) error {
	cdps := k.GetSliceOfCdpsByRatioAndType(ctx, slice, targetRatio, collateralType)
	for _, cdp := range cdps {
		k.SynchronizeInterest(ctx, cdp)
	}
	return nil
}
