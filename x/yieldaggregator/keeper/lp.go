package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	osmomath "github.com/UnUniFi/chain/osmomath"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) VaultAmountInStrategies(ctx sdk.Context, vault types.Vault) sdkmath.Int {
	amountInStrategies := sdk.ZeroInt()

	// calculate amount in strategies
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, strategyWeight.Denom, strategyWeight.StrategyId)
		if !found {
			continue
		}
		amount, err := k.GetAmountFromStrategy(ctx, vault, strategy)
		if err != nil {
			continue
		}
		amountInStrategies = amountInStrategies.Add(amount.Amount)
	}
	return amountInStrategies
}

func (k Keeper) VaultUnbondingAmountInStrategies(ctx sdk.Context, vault types.Vault) sdkmath.Int {
	unbondingAmount := sdk.ZeroInt()

	// calculate amount in strategies
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, strategyWeight.Denom, strategyWeight.StrategyId)
		if !found {
			continue
		}
		amount, err := k.GetUnbondingAmountFromStrategy(ctx, vault, strategy)
		if err != nil {
			continue
		}
		unbondingAmount = unbondingAmount.Add(amount.Amount)
	}
	return unbondingAmount
}

func (k Keeper) VaultBalances(ctx sdk.Context, vault types.Vault) sdk.Coins {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	return k.bankKeeper.GetAllBalances(ctx, vaultModAddr)
}

func (k Keeper) VaultWithdrawalAmount(ctx sdk.Context, vault types.Vault) sdk.Int {
	amount := sdk.ZeroInt()
	vaultBalances := k.VaultBalances(ctx, vault)
	for _, balance := range vaultBalances {
		denomInfo := k.GetDenomInfo(ctx, balance.Denom)
		if denomInfo.Symbol == vault.Symbol {
			amount = amount.Add(balance.Amount)
		}
	}
	return amount
}

func (k Keeper) VaultAmountTotal(ctx sdk.Context, vault types.Vault) sdk.Int {
	amountInStrategies := k.VaultAmountInStrategies(ctx, vault)
	amountInVault := k.VaultWithdrawalAmount(ctx, vault)
	amountUnbonding := k.VaultUnbondingAmountInStrategies(ctx, vault)
	pendingDeposit := k.recordsKeeper.GetVaultPendingDeposit(ctx, vault.Id)

	totalAmount := amountInStrategies.Add(amountInVault).Add(amountUnbonding).Add(pendingDeposit)
	return totalAmount
}

// lpAmount = lpSupply * (principalAmountToMint / principalAmountInVault)
// If principalAmountInVault is zero, lpAmount = principalAmountToMint
func (k Keeper) EstimateMintAmountInternal(ctx sdk.Context, vaultId uint64, principalAmount sdk.Int) sdk.Coin {
	lpDenom := types.GetLPTokenDenom(vaultId)
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return sdk.NewCoin(lpDenom, sdk.ZeroInt())
	}

	totalVaultAmount := k.VaultAmountTotal(ctx, vault)
	lpSupply := k.bankKeeper.GetSupply(ctx, lpDenom).Amount
	if totalVaultAmount.IsZero() || lpSupply.IsZero() {
		return sdk.NewCoin(lpDenom, principalAmount)
	}

	lpAmount := lpSupply.Mul(principalAmount).Quo(totalVaultAmount)

	return sdk.NewCoin(lpDenom, lpAmount)
}

// calculate principalAmount
// principalAmount = principalAmountInVault * (lpAmountToBurn / lpSupply)
func (k Keeper) EstimateRedeemAmountInternal(ctx sdk.Context, vaultId uint64, lpAmount sdkmath.Int) sdk.Int {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return sdk.ZeroInt()
	}
	principalInVault := k.VaultAmountTotal(ctx, vault)
	lpDenom := types.GetLPTokenDenom(vaultId)
	lpSupply := k.bankKeeper.GetSupply(ctx, lpDenom).Amount

	if lpSupply.IsZero() {
		return sdk.ZeroInt()
	}
	principalAmount := principalInVault.Mul(lpAmount).Quo(lpSupply)

	return principalAmount
}

func (k Keeper) SendCoinsFromVault(ctx sdk.Context, vault types.Vault, address sdk.AccAddress, amount sdk.Int) error {
	vaultDenoms := vault.StrategyDenoms()
	vaultBalances := k.VaultBalances(ctx, vault)
	coins := sdk.Coins{}
	remainingAmount := amount
	for _, denom := range vaultDenoms {
		denomAmount := vaultBalances.AmountOf(denom)
		if denomAmount.IsZero() {
			continue
		}
		if remainingAmount.GT(denomAmount) {
			coins = coins.Add(sdk.NewCoin(denom, denomAmount))
			remainingAmount = remainingAmount.Sub(denomAmount)
		} else {
			coins = coins.Add(sdk.NewCoin(denom, remainingAmount))
			break
		}
	}

	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	return k.bankKeeper.SendCoins(ctx, vaultModAddr, address, coins)
}

func (k Keeper) DepositAndMintLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, principalCoin sdk.Coin) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return types.ErrInvalidVaultId
	}

	denomInfo := k.GetDenomInfo(ctx, principalCoin.Denom)
	if denomInfo.Symbol != vault.Symbol {
		return types.ErrDenomDoesNotMatchVaultSymbol
	}

	// calculate lpCoin token amount
	lpCoin := k.EstimateMintAmountInternal(ctx, vaultId, principalCoin.Amount)

	// transfer coins after lp amount calculation
	vaultModName := types.GetVaultModuleAccountName(vaultId)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	err := k.bankKeeper.SendCoins(ctx, address, vaultModAddr, sdk.NewCoins(principalCoin))
	if err != nil {
		return err
	}

	// mint and transfer lp token
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(lpCoin))
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(lpCoin))
	if err != nil {
		return err
	}

	// Allocate funds through strategy
	totalAmount := k.VaultAmountTotal(ctx, vault)
	stratAmount := k.VaultAmountInStrategies(ctx, vault)
	newStrategyAmount := sdk.NewDecFromInt(totalAmount).Mul(sdk.OneDec().Sub(vault.WithdrawReserveRate)).RoundInt()
	amountToInvest := newStrategyAmount.Sub(stratAmount)
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, strategyWeight.Denom, strategyWeight.StrategyId)
		if !found {
			continue
		}
		strategyAmount := sdk.NewDecFromInt(amountToInvest).Mul(strategyWeight.Weight).RoundInt()
		if !strategyAmount.IsPositive() {
			continue
		}
		err = k.StakeToStrategy(ctx, vault, strategy, strategyAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) BurnLPTokenAndRedeem(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, lpAmount sdkmath.Int) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return types.ErrInvalidVaultId
	}

	params, err := k.GetParams(ctx)
	if err != nil {
		return err
	}

	principal := k.EstimateRedeemAmountInternal(ctx, vaultId, lpAmount)

	// burn lp tokens after calculating withdrawal amount
	lpDenom := types.GetLPTokenDenom(vaultId)
	lp := sdk.NewCoin(lpDenom, lpAmount)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}

	// Unstake funds from the vault
	amountToUnbond := principal

	// implement fees on withdrawal
	amountInVault := k.VaultWithdrawalAmount(ctx, vault)
	amountUnbonding := k.VaultUnbondingAmountInStrategies(ctx, vault)

	// reserveMaintenanceRate := sdk.NewDecFromInt(amountInVault).Quo(sdk.NewDecFromInt(amountInVault.Add(amountUnbonding)))
	reserveMaintenanceRate := sdk.ZeroDec()
	if amountInVault.GT(amountToUnbond) {
		reserveMaintenanceRate = sdk.NewDecFromInt(amountInVault.Sub(amountToUnbond)).Quo(sdk.NewDecFromInt(amountInVault.Add(amountUnbonding)))
	}
	// reserve_maintenance_rate = max(0, withdraw_reserve - amount_to_withdraw) / (withdraw_reserve + tokens_in_unbonding_period)

	e := osmomath.NewDecWithPrec(2718281, 6) // 2.718281
	withdrawFeeRate := osmomath.OneDec().
		Quo(e.Power(osmomath.BigDecFromSDKDec(reserveMaintenanceRate).MulInt64(10))).
		SDKDec()

	// withdraw_reserve + tokens_in_unbonding_period + bonding_amount = total_amount
	// reserve_maintenance_rate = max(0, withdraw_reserve - amount_to_withdraw) / (withdraw_reserve + tokens_in_unbonding_period)
	// withdraw_fee_rate = e^(-10 * (reserve_maintenance_rate))
	// withdraw_fee = withdraw_fee_rate * amount_to_withdraw
	// If reserve_maintenance_rate is close to 1, withdraw_fee_rate will be close to 0 and vice versa

	withdrawFee := sdk.NewDecFromInt(amountToUnbond).Mul(withdrawFeeRate).RoundInt()
	withdrawAmount := amountToUnbond.Sub(withdrawFee)

	withdrawModuleCommissionFee := sdk.NewDecFromInt(withdrawAmount).Mul(params.CommissionRate).RoundInt()
	withdrawVaultCommissionFee := sdk.NewDecFromInt(withdrawAmount).Mul(vault.WithdrawCommissionRate).RoundInt()
	withdrawAmountWithoutCommission := withdrawAmount.Sub(withdrawModuleCommissionFee).Sub(withdrawVaultCommissionFee)

	if withdrawModuleCommissionFee.IsPositive() {
		feeCollector, err := sdk.AccAddressFromBech32(params.FeeCollectorAddress)
		if err != nil {
			return err
		}
		err = k.SendCoinsFromVault(ctx, vault, feeCollector, withdrawModuleCommissionFee)
		if err != nil {
			return err
		}
	}

	if withdrawVaultCommissionFee.IsPositive() {
		vaultOwner, err := sdk.AccAddressFromBech32(vault.Owner)
		if err != nil {
			return err
		}
		err = k.SendCoinsFromVault(ctx, vault, vaultOwner, withdrawVaultCommissionFee)
		if err != nil {
			return err
		}
	}

	err = k.SendCoinsFromVault(ctx, vault, address, withdrawAmountWithoutCommission)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) BurnLPTokenAndBeginUnbonding(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, lpAmount sdkmath.Int) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return types.ErrInvalidVaultId
	}

	principal := k.EstimateRedeemAmountInternal(ctx, vaultId, lpAmount)

	// burn lp tokens after calculating withdrawal amount
	lpDenom := types.GetLPTokenDenom(vaultId)
	lp := sdk.NewCoin(lpDenom, lpAmount)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}

	// Unstake funds from Strategy
	amountToUnbond := principal
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, strategyWeight.Denom, strategyWeight.StrategyId)
		if !found {
			continue
		}
		strategyAmount := sdk.NewDecFromInt(amountToUnbond).Mul(strategyWeight.Weight).RoundInt()

		err = k.UnstakeFromStrategy(ctx, vault, strategy, strategyAmount, address.String())
		if err != nil {
			return err
		}
	}

	return nil
}
