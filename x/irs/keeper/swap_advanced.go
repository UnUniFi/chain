package keeper

// TODO:
func (k Keeper) SwapUtToYt() {
	// Internally combine MintPtYtPair and SwapPtToUt

	// Borrow msg.AmountToBuy from AMM pool
	// Open position
	// Sell msg.AmountToBuy worth of PT
	// Return borrowed amount
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
