package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

func (k Keeper) GetLPTokenDenom(vaultId uint64) string {
	return fmt.Sprintf("yield-aggregator/vaults/%d", vaultId)
}

func (k Keeper) MintLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, amount sdk.Int) {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		panic("vault not found")
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, sdk.NewCoins(sdk.NewCoin(vault.Denom, amount)))
	if err != nil {
		panic(err)
	}

	lpDenom := k.GetLPTokenDenom(vaultId)
	// TODO: calculate lpAmount
	lpAmount := sdk.NewInt(0)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(lpDenom, lpAmount)))
	if err != nil {
		panic(err)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(sdk.NewCoin(lpDenom, lpAmount)))
	if err != nil {
		panic(err)
	}
}

func (k Keeper) BurnLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, amount sdk.Int) {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		panic("vault not found")
	}

	lpDenom := k.GetLPTokenDenom(vaultId)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, sdk.NewCoins(sdk.NewCoin(lpDenom, amount)))
	if err != nil {
		panic(err)
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(lpDenom, amount)))
	if err != nil {
		panic(err)
	}

	// TOOD: calculate principalAmount
	principalAmount := sdk.NewInt(0)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(sdk.NewCoin(vault.Denom, principalAmount)))
	if err != nil {
		panic(err)
	}
}
