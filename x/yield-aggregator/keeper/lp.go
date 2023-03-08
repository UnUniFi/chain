package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

// lpAmount = lpSupply * (principalAmountToMint / principalAmountInVault)
// If principalAmountInVault is zero, lpAmount = principalAmountToMint
func (k Keeper) EstimateMintAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, principalAmount sdk.Int) sdk.Coin {
	lpDenom := types.GetLPTokenDenom(vaultId)
	principalAmountInVault := k.GetAmountFromStrategy(ctx, vaultDenom, vaultId)

	if principalAmountInVault.IsZero() {
		return sdk.NewCoin(lpDenom, principalAmount)
	}

	lpSupply := k.bankKeeper.GetSupply(ctx, lpDenom).Amount
	lpAmount := lpSupply.Mul(principalAmount).Quo(principalAmountInVault)

	return sdk.NewCoin(lpDenom, lpAmount)
}

// calculate principalAmount
// principalAmount = principalAmountInVault * (lpAmountToBurn / lpSupply)
func (k Keeper) EstimateRedeemAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, lpAmount sdk.Int) sdk.Coin {
	principalAmountInVault := k.GetAmountFromStrategy(ctx, vaultDenom, vaultId)
	lpDenom := types.GetLPTokenDenom(vaultId)
	lpSupply := k.bankKeeper.GetSupply(ctx, lpDenom).Amount

	if lpSupply.IsZero() {
		return sdk.NewCoin(vaultDenom, sdk.ZeroInt())
	}
	principalAmount := principalAmountInVault.Mul(lpAmount).Quo(lpSupply)

	return sdk.NewCoin(vaultDenom, principalAmount)
}

func (k Keeper) MintLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, principalAmount sdk.Int) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		// TODO
		panic("vault not found")
	}

	moduleName := types.GetModuleAccountName(vaultId)

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

	return nil
}

func (k Keeper) BurnLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, lpAmount sdk.Int) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		// TODO
		panic("vault not found")
	}

	moduleName := types.GetModuleAccountName(vaultId)

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

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleName, address, sdk.NewCoins(principal))
	if err != nil {
		return err
	}

	return nil
}
