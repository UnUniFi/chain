package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) SendMarginToMarginManager(ctx sdk.Context, sender sdk.AccAddress, margin sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.MarginManager, margin)
}

func (k Keeper) SendCoinFromMarginManagerToPool(ctx sdk.Context, amount sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.MarginManager, types.ModuleName, amount)
}

func (k Keeper) SendBackMargin(ctx sdk.Context, recipient sdk.AccAddress, amount sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, amount)
}
