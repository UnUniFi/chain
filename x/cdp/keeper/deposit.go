package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/cdp/types"
)

// DepositCollateral adds collateral to a cdp
func (k Keeper) DepositCollateral(ctx sdk.Context, owner, depositor sdk.AccAddress, collateral sdk.Coin, collateralType string) error {
	// check that collateral exists and has a functioning pricefeed
	err := k.ValidateCollateral(ctx, collateral, collateralType)
	if err != nil {
		return err
	}
	cdp, found := k.GetCdpByOwnerAndCollateralType(ctx, owner, collateralType)
	if !found {
		return sdkerrors.Wrapf(types.ErrCdpNotFound, "owner %s, collateral %s", owner, collateralType)
	}
	err = k.ValidateBalance(ctx, collateral, depositor)
	if err != nil {
		return err
	}
	k.hooks.BeforeCdpModified(ctx, cdp)
	cdp = k.SynchronizeInterest(ctx, cdp)

	deposit, found := k.GetDeposit(ctx, cdp.Id, depositor)
	if found {
		deposit.Amount = deposit.Amount.Add(collateral)
	} else {
		deposit = types.NewDeposit(cdp.Id, depositor, collateral)
	}
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	k.SetDeposit(ctx, deposit)

	cdp.Collateral = cdp.Collateral.Add(collateral)
	collateralToDebtRatio := k.CalculateCollateralToDebtRatio(ctx, cdp.Collateral, cdp.Type, cdp.GetTotalPrincipal())

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCdpDeposit,
			sdk.NewAttribute(sdk.AttributeKeyAmount, collateral.String()),
			sdk.NewAttribute(types.AttributeKeyCdpID, fmt.Sprintf("%d", cdp.Id)),
		),
	)

	return k.UpdateCdpAndCollateralRatioIndex(ctx, cdp, collateralToDebtRatio)
}

// WithdrawCollateral removes collateral from a cdp if it does not put the cdp below the liquidation ratio
func (k Keeper) WithdrawCollateral(ctx sdk.Context, owner, depositor sdk.AccAddress, collateral sdk.Coin, collateralType string) error {
	err := k.ValidateCollateral(ctx, collateral, collateralType)
	if err != nil {
		return err
	}
	cdp, found := k.GetCdpByOwnerAndCollateralType(ctx, owner, collateralType)
	if !found {
		return sdkerrors.Wrapf(types.ErrCdpNotFound, "owner %s, collateral %s", owner, collateral.Denom)
	}
	deposit, found := k.GetDeposit(ctx, cdp.Id, depositor)
	if !found {
		return sdkerrors.Wrapf(types.ErrDepositNotFound, "depositor %s, collateral %s %s", depositor, collateral.Denom, collateralType)
	}
	if collateral.Amount.GT(deposit.Amount.Amount) {
		return sdkerrors.Wrapf(types.ErrInvalidWithdrawAmount, "collateral %s, deposit %s", collateral, deposit.Amount)
	}
	k.hooks.BeforeCdpModified(ctx, cdp)
	cdp = k.SynchronizeInterest(ctx, cdp)

	collateralizationRatio, err := k.CalculateCollateralizationRatio(ctx, cdp.Collateral.Sub(collateral), cdp.Type, cdp.Principal, cdp.AccumulatedFees, spot)
	if err != nil {
		return err
	}
	liquidationRatio := k.getLiquidationRatio(ctx, cdp.Type)
	if collateralizationRatio.LT(liquidationRatio) {
		return sdkerrors.Wrapf(types.ErrInvalidCollateralRatio, "collateral %s, collateral ratio %s, liquidation ration %s", collateral.Denom, collateralizationRatio, liquidationRatio)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoins(collateral))
	if err != nil {
		panic(err)
	}

	cdp.Collateral = cdp.Collateral.Sub(collateral)
	collateralToDebtRatio := k.CalculateCollateralToDebtRatio(ctx, cdp.Collateral, cdp.Type, cdp.GetTotalPrincipal())
	err = k.UpdateCdpAndCollateralRatioIndex(ctx, cdp, collateralToDebtRatio)
	if err != nil {
		return err
	}

	deposit.Amount = deposit.Amount.Sub(collateral)
	// delete deposits if amount is 0
	if deposit.Amount.IsZero() {
		k.DeleteDeposit(ctx, deposit.CdpId, deposit.Depositor.AccAddress())
	} else {
		k.SetDeposit(ctx, deposit)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCdpWithdrawal,
			sdk.NewAttribute(sdk.AttributeKeyAmount, collateral.String()),
			sdk.NewAttribute(types.AttributeKeyCdpID, fmt.Sprintf("%d", cdp.Id)),
		),
	)

	return nil
}

// GetDeposit returns the deposit of a depositor on a particular cdp from the store
func (k Keeper) GetDeposit(ctx sdk.Context, cdpID uint64, depositor sdk.AccAddress) (deposit types.Deposit, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	bz := store.Get(types.DepositKeySuffix(cdpID, depositor))
	if bz == nil {
		return deposit, false
	}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &deposit)
	return deposit, true

}

// SetDeposit sets the deposit in the store
func (k Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	bz := k.cdc.MustMarshalLengthPrefixed(&deposit)
	store.Set(types.DepositKeySuffix(deposit.CdpId, deposit.Depositor.AccAddress()), bz)

}

// DeleteDeposit deletes a deposit from the store
func (k Keeper) DeleteDeposit(ctx sdk.Context, cdpID uint64, depositor sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	store.Delete(types.DepositKeySuffix(cdpID, depositor))
}

// IterateDeposits iterates over the all the deposits of a cdp and performs a callback function
func (k Keeper) IterateDeposits(ctx sdk.Context, cdpID uint64, cb func(deposit types.Deposit) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	iterator := sdk.KVStorePrefixIterator(store, types.GetCdpIDBytes(cdpID))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var deposit types.Deposit
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &deposit)

		if cb(deposit) {
			break
		}
	}
}

// GetDeposits returns all the deposits to a cdp
func (k Keeper) GetDeposits(ctx sdk.Context, cdpID uint64) (deposits types.Deposits) {
	k.IterateDeposits(ctx, cdpID, func(deposit types.Deposit) bool {
		deposits = append(deposits, deposit)
		return false
	})
	return
}
