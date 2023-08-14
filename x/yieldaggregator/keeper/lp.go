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

func (k Keeper) VaultUnbondingAmountInStrategies(ctx sdk.Context, vault types.Vault) sdkmath.Int {
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
func (k Keeper) EstimateMintAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, principalAmount sdkmath.Int) sdk.Coin {
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
func (k Keeper) EstimateRedeemAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, lpAmount sdkmath.Int) sdk.Coin {
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
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(lp))
	if err != nil {
		return err
	}

	// Allocate funds through strategy
	totalAmount := k.VaultAmountTotal(ctx, vault)
	stratAmount := k.VaultAmountInStrategies(ctx, vault)
	newStrategyAmount := sdk.NewDecFromInt(totalAmount).Mul(sdk.OneDec().Sub(vault.WithdrawReserveRate)).RoundInt()
	amountToInvest := newStrategyAmount.Sub(stratAmount)
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, vault.Denom, strategyWeight.StrategyId)
		if !found {
			continue
		}
		strategyAmount := sdk.NewDecFromInt(amountToInvest).Mul(strategyWeight.Weight).RoundInt()
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

	principal := k.EstimateRedeemAmountInternal(ctx, vault.Denom, vaultId, lpAmount)

	// burn lp tokens after calculating withdrawal amount
	vaultModName := types.GetVaultModuleAccountName(vaultId)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
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

	// Unstake funds from Strategy
	amountToUnbond := principal.Amount

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

	withdrawFee := sdk.NewDecFromInt(principal.Amount).Mul(withdrawFeeRate).RoundInt()
	withdrawAmount := principal.Amount.Sub(withdrawFee)

	withdrawModuleCommissionFee := sdk.NewDecFromInt(withdrawAmount).Mul(params.CommissionRate).RoundInt()
	withdrawVaultCommissionFee := sdk.NewDecFromInt(withdrawAmount).Mul(vault.WithdrawCommissionRate).RoundInt()
	withdrawAmountWithoutCommission := withdrawAmount.Sub(withdrawModuleCommissionFee).Sub(withdrawVaultCommissionFee)

	if withdrawModuleCommissionFee.IsPositive() {
		feeCollector, err := sdk.AccAddressFromBech32(params.FeeCollectorAddress)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoins(ctx, vaultModAddr, feeCollector, sdk.NewCoins(sdk.NewCoin(principal.Denom, withdrawModuleCommissionFee)))
		if err != nil {
			return err
		}
	}

	if withdrawVaultCommissionFee.IsPositive() {
		vaultOwner, err := sdk.AccAddressFromBech32(vault.Owner)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoins(ctx, vaultModAddr, vaultOwner, sdk.NewCoins(sdk.NewCoin(principal.Denom, withdrawVaultCommissionFee)))
		if err != nil {
			return err
		}
	}

	err = k.bankKeeper.SendCoins(ctx, vaultModAddr, address, sdk.NewCoins(sdk.NewCoin(principal.Denom, withdrawAmountWithoutCommission)))
	if err != nil {
		return err
	}

	return nil
}
