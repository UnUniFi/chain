package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	osmomath "github.com/UnUniFi/chain/osmomath"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k Keeper) VaultAmountInStrategies(ctx sdk.Context, vault types.Vault) sdk.Int {
	amountInStrategies := sdk.ZeroInt()

	// calculate amount in strategies
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, vault.Denom, strategyWeight.StrategyId)
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

func (k Keeper) VaultUnbondingAmountInStrategies(ctx sdk.Context, vault types.Vault) sdk.Int {
	unbondingAmount := sdk.ZeroInt()

	// calculate amount in strategies
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, vault.Denom, strategyWeight.StrategyId)
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

func (k Keeper) VaultWithdrawalAmount(ctx sdk.Context, vault types.Vault) sdk.Int {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	balance := k.bankKeeper.GetBalance(ctx, vaultModAddr, vault.Denom)
	return balance.Amount
}

func (k Keeper) VaultAmountTotal(ctx sdk.Context, vault types.Vault) sdk.Int {
	amountInStrategies := k.VaultAmountInStrategies(ctx, vault)
	amountInVault := k.VaultWithdrawalAmount(ctx, vault)
	amountUnbonding := k.VaultUnbondingAmountInStrategies(ctx, vault)

	totalAmount := amountInStrategies.Add(amountInVault).Add(amountUnbonding)
	return totalAmount
}

// lpAmount = lpSupply * (principalAmountToMint / principalAmountInVault)
// If principalAmountInVault is zero, lpAmount = principalAmountToMint
func (k Keeper) EstimateMintAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, principalAmount sdk.Int) sdk.Coin {
	lpDenom := types.GetLPTokenDenom(vaultId)
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return sdk.NewCoin(lpDenom, sdk.ZeroInt())
	}

	totalVaultAmount := k.VaultAmountTotal(ctx, vault)
	if totalVaultAmount.IsZero() {
		return sdk.NewCoin(lpDenom, principalAmount)
	}

	lpSupply := k.bankKeeper.GetSupply(ctx, lpDenom).Amount
	lpAmount := lpSupply.Mul(principalAmount).Quo(totalVaultAmount)

	return sdk.NewCoin(lpDenom, lpAmount)
}

// calculate principalAmount
// principalAmount = principalAmountInVault * (lpAmountToBurn / lpSupply)
func (k Keeper) EstimateRedeemAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, lpAmount sdk.Int) sdk.Coin {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return sdk.NewCoin(vaultDenom, sdk.ZeroInt())
	}
	principalInVault := k.VaultAmountTotal(ctx, vault)
	lpDenom := types.GetLPTokenDenom(vaultId)
	lpSupply := k.bankKeeper.GetSupply(ctx, lpDenom).Amount

	if lpSupply.IsZero() {
		return sdk.NewCoin(vaultDenom, sdk.ZeroInt())
	}
	principalAmount := principalInVault.Mul(lpAmount).Quo(lpSupply)

	return sdk.NewCoin(vaultDenom, principalAmount)
}

func (k Keeper) DepositAndMintLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, principalAmount sdk.Int) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return types.ErrInvalidVaultId
	}

	// calculate lp token amount
	lp := k.EstimateMintAmountInternal(ctx, vault.Denom, vaultId, principalAmount)

	// transfer coins after lp amount calculation
	vaultModName := types.GetVaultModuleAccountName(vaultId)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	principal := sdk.NewCoin(vault.Denom, principalAmount)
	err := k.bankKeeper.SendCoins(ctx, address, vaultModAddr, sdk.NewCoins(principal))
	if err != nil {
		return err
	}

	// mint and trasnfer lp token
	err = k.bankKeeper.MintCoins(ctx, vaultModName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoins(ctx, vaultModAddr, address, sdk.NewCoins(lp))
	if err != nil {
		return err
	}

	// Allocate funds through strategy
	totalAmount := k.VaultAmountTotal(ctx, vault)
	stratAmount := k.VaultAmountInStrategies(ctx, vault)
	newStrategyAmount := totalAmount.ToDec().Mul(sdk.OneDec().Sub(vault.WithdrawReserveRate)).RoundInt()
	amountToInvest := newStrategyAmount.Sub(stratAmount)
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, vault.Denom, strategyWeight.StrategyId)
		if !found {
			continue
		}
		strategyAmount := amountToInvest.ToDec().Mul(strategyWeight.Weight).RoundInt()
		err = k.StakeToStrategy(ctx, vault, strategy, strategyAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) BurnLPTokenAndRedeem(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, lpAmount sdk.Int) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return types.ErrInvalidVaultId
	}

	principal := k.EstimateRedeemAmountInternal(ctx, vault.Denom, vaultId, lpAmount)

	// burn lp tokens after calculating withdrawal amount
	vaultModName := types.GetVaultModuleAccountName(vaultId)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	lpDenom := types.GetLPTokenDenom(vaultId)
	lp := sdk.NewCoin(lpDenom, lpAmount)
	err := k.bankKeeper.SendCoins(ctx, address, vaultModAddr, sdk.NewCoins(lp))
	if err != nil {
		return err
	}
	err = k.bankKeeper.BurnCoins(ctx, vaultModName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}

	// Unstake funds from Strategy
	amountToUnbond := principal.Amount
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, vault.Denom, strategyWeight.StrategyId)
		if !found {
			continue
		}
		strategyAmount := amountToUnbond.ToDec().Mul(strategyWeight.Weight).RoundInt()
		err = k.UnstakeFromStrategy(ctx, vault, strategy, strategyAmount)
		if err != nil {
			return err
		}
	}

	// implement fees on withdrawal
	amountInVault := k.VaultWithdrawalAmount(ctx, vault)
	amountUnbonding := k.VaultUnbondingAmountInStrategies(ctx, vault)

	// reserveMaintenanceRate := amountInVault.ToDec().Quo(amountInVault.Add(amountUnbonding).ToDec())
	reserveMaintenanceRate := sdk.ZeroDec()
	if amountInVault.GT(amountToUnbond) {
		reserveMaintenanceRate = amountInVault.Sub(amountToUnbond).ToDec().Quo(amountInVault.Add(amountUnbonding).ToDec())
	}
	// reserve_maintenance_rate = max(0, withdraw_reserve - amount_to_withdraw) / (withdraw_reserve + tokens_in_unbonding_period)

	e := osmomath.NewDecWithPrec(2718281, 6) // 2.718281
	eInv := osmomath.OneDec().Quo(e)         // e^-1
	withdrawFeeRate := eInv.Power(osmomath.BigDecFromSDKDec(reserveMaintenanceRate).MulInt64(10)).SDKDec()

	// withdraw_reserve + tokens_in_unbonding_period + bonding_amount = total_amount
	// reserve_maintenance_rate = max(0, withdraw_reserve - amount_to_withdraw) / (withdraw_reserve + tokens_in_unbonding_period)
	// withdraw_fee_rate = e^(-10 * (reserve_maintenance_rate))
	// withdraw_fee = withdraw_fee_rate * amount_to_withdraw
	// If reserve_maintenance_rate is close to 1, withdraw_fee_rate will be close to 0 and vice versa

	withdrawFee := principal.Amount.ToDec().Mul(withdrawFeeRate).RoundInt()
	withdrawAmount := principal.Amount.Sub(withdrawFee)
	withdrawCoin := sdk.NewCoin(principal.Denom, withdrawAmount)

	err = k.bankKeeper.SendCoins(ctx, vaultModAddr, address, sdk.NewCoins(withdrawCoin))
	if err != nil {
		return err
	}

	return nil
}
