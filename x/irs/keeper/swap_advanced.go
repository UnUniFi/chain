package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) SwapUtToYt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, requiredYtAmount sdk.Coin) error {
	// Take loan from IRS vault account
	moduleAddr := types.GetVaultModuleAddress(pool)
	err := k.bankKeeper.SendCoins(ctx, sender, moduleAddr, sdk.Coins{requiredYtAmount})
	if err != nil {
		return err
	}

	// Mint Pt and Yt
	ptAmount, err := k.MintPtYtPair(ctx, sender, pool, requiredYtAmount)
	if err != nil {
		return err
	}

	// Sell minted PT amount for underlying token
	ptDenom := types.PtDenom(pool)
	err = k.SwapPtToUt(ctx, sender, pool, sdk.NewCoin(ptDenom, ptAmount))
	if err != nil {
		return err
	}

	// Payback loan
	err = k.bankKeeper.SendCoins(ctx, sender, moduleAddr, sdk.Coins{requiredYtAmount})
	if err != nil {
		return err
	}

	return nil
}

// TODO:
func (k Keeper) SwapYtToUt() {
	// Internally combine SwapUtToPt and BurnPtYtPair

	// If matured, send required amount from unbonded from the share
	// Else
	// Put required amount of msg.PT from user wallet
	// Close position
	// Start redemption for strategy share
}
