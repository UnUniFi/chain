package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

func (k Keeper) MintLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, principalAmount sdk.Int) error {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		// TODO
		panic("vault not found")
	}

	moduleName := types.GetModuleAccountName(vaultId)

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, moduleName, sdk.NewCoins(sdk.NewCoin(vault.Denom, principalAmount)))
	if err != nil {
		return err
	}

	lpDenom := types.GetLPTokenDenom(vaultId)
	// TODO: calculate lpAmount
	// lpAmount = lpSupplyInPool * (principalAmountToMint / principalSupplyInPool)
	// If principalSupplyInPool is zero, lpAmount = principalAmountToMint
	lpAmount := sdk.NewInt(0)
	err = k.bankKeeper.MintCoins(ctx, moduleName, sdk.NewCoins(sdk.NewCoin(lpDenom, lpAmount)))
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleName, address, sdk.NewCoins(sdk.NewCoin(lpDenom, lpAmount)))
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
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, moduleName, sdk.NewCoins(sdk.NewCoin(lpDenom, lpAmount)))
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, moduleName, sdk.NewCoins(sdk.NewCoin(lpDenom, lpAmount)))
	if err != nil {
		return err
	}

	// TOOD: calculate principalAmount
	// principalAmount = principalSupplyInPool * (lpAmountToBurn / lpSupplyInPool)
	principalAmount := sdk.NewInt(0)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleName, address, sdk.NewCoins(sdk.NewCoin(vault.Denom, principalAmount)))
	if err != nil {
		return err
	}

	return nil
}
