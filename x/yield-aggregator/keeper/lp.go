package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

// lpAmount = lpSupply * (principalAmountToMint / principalAmountInVault)
// If principalAmountInVault is zero, lpAmount = principalAmountToMint
func (k Keeper) EstimateMintAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, principalAmount sdk.Int) sdk.Coin {
	lpDenom := types.GetLPTokenDenom(vaultId)
	strategy, found := k.GetStrategy(ctx, vaultDenom, vaultId)
	if !found {
		return sdk.NewCoin(lpDenom, sdk.ZeroInt())
	}
	principalInVault, err := k.GetAmountFromStrategy(ctx, strategy)
	if err != nil {
		return sdk.NewCoin(lpDenom, sdk.ZeroInt())
	}

	if principalInVault.IsZero() {
		return sdk.NewCoin(lpDenom, principalAmount)
	}

	lpSupply := k.bankKeeper.GetSupply(ctx, lpDenom).Amount
	lpAmount := lpSupply.Mul(principalAmount).Quo(principalInVault.Amount)

	return sdk.NewCoin(lpDenom, lpAmount)
}

// calculate principalAmount
// principalAmount = principalAmountInVault * (lpAmountToBurn / lpSupply)
func (k Keeper) EstimateRedeemAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, lpAmount sdk.Int) sdk.Coin {
	strategy, found := k.GetStrategy(ctx, vaultDenom, vaultId)
	if !found {
		return sdk.NewCoin(vaultDenom, sdk.ZeroInt())
	}
	principalInVault, err := k.GetAmountFromStrategy(ctx, strategy)
	if err != nil {
		return sdk.NewCoin(vaultDenom, sdk.ZeroInt())
	}
	lpDenom := types.GetLPTokenDenom(vaultId)
	lpSupply := k.bankKeeper.GetSupply(ctx, lpDenom).Amount

	if lpSupply.IsZero() {
		return sdk.NewCoin(vaultDenom, sdk.ZeroInt())
	}
	principalAmount := principalInVault.Amount.Mul(lpAmount).Quo(lpSupply)

	return sdk.NewCoin(vaultDenom, principalAmount)
}

func (k Keeper) DepositAndMintLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, principalAmount sdk.Int) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		// TODO
		panic("vault not found")
	}

	moduleName := types.GetVaultModuleAccountName(vaultId)

	principal := sdk.NewCoin(vault.Denom, principalAmount)

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, moduleName, sdk.NewCoins(principal))

	if err != nil {
		return err
	}

	lp := k.EstimateMintAmountInternal(ctx, vault.Denom, vaultId, principalAmount)

	err = k.bankKeeper.MintCoins(ctx, moduleName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleName, address, sdk.NewCoins(lp))
	if err != nil {
		return err
	}

	// TODO: Allocate funds to Strategy module accounts

	return nil
}

func (k Keeper) BurnLPTokenAndRedeem(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, lpAmount sdk.Int) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		// TODO
		panic("vault not found")
	}

	moduleName := types.GetVaultModuleAccountName(vaultId)

	lpDenom := types.GetLPTokenDenom(vaultId)
	lp := sdk.NewCoin(lpDenom, lpAmount)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, moduleName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, moduleName, sdk.NewCoins(lp))
	if err != nil {
		return err
	}

	principal := k.EstimateRedeemAmountInternal(ctx, vault.Denom, vaultId, lpAmount)

	// TODO: Unstake funds from Strategy
	// TODO: Withdraw funds from "preparation for withdraw" of Strategy module account
	// TODO: If "preparation for withdraw" is soon to be short, withdraw fee will be imposed
	// withdraw_reserve + tokens_in_unbonding_period + bonding_amount = total_amount
	// reserve_maintenance_rate = withdraw_reserve / (withdraw_reserve + tokens_in_unbonding_period)
	// withdraw_fee_rate = e^(-10 * (reserve_maintenance_rate))
	// If reserve_maintenance_rate is close to 1, withdraw_fee_rate will be close to 0 and vice versa

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleName, address, sdk.NewCoins(principal))
	if err != nil {
		return err
	}

	return nil
}
