package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func CalculateMintFee(currentBalance, targetBalance, sentCoin sdk.Coin, feeRate sdk.Dec) sdk.Coin {
	increaseRate := sdk.NewDecFromInt(currentBalance.Amount).Sub(sdk.NewDecFromInt(targetBalance.Amount)).Quo(sdk.NewDecFromInt(targetBalance.Amount))
	if increaseRate.IsNegative() {
		increaseRate = sdk.NewDec(0)
	}

	mintFeeRate := feeRate.Mul(increaseRate.Add(sdk.NewDecWithPrec(1, 0)))
	fee := sdk.NewCoin(sentCoin.Denom, sdk.NewDec(sentCoin.Amount.Int64()).Mul(mintFeeRate).TruncateInt())

	return fee
}
