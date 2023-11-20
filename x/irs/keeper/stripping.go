package keeper

func (k Keeper) MintPtYtPair() {
	// TODO:
	// usedUnderlying - amount of ATOM used
	// shares - the amount of shares received from stATOM-ATOM strategy
	// valueSupplied - The outstanding amount of underlying which can be redeemed from the contract from Principal Tokens
	// holdingsValue - total underlying value on the contract
	// YT mint amount - usedUnderlying
	// PT mint amount - usedUnderlying * (1-(holdingsValue-valueSupplied)/interestSupply)

}
