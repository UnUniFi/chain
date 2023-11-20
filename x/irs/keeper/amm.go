package keeper

// TODO:
func (k Keeper) DepositToLiquidityPool() {
	// BankKeeper.send Ut Sender→Pool
	// BankKeeper.send Pt Sender→Pool
	// BankKeeper.send Ls Pool→Sender
}

// TODO:
func (k Keeper) WithdrawFromLiquidityPool() {
	// BankKeeper.send Ls Sender→Pool
	// BankKeeper.send Ut Pool→Sender
	// BankKeeper.send Pt Pool→Sender
}
