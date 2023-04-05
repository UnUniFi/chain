package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/cdp/types"
)

// AttemptKeeperLiquidation liquidates the cdp with the input collateral type and owner if it is below the required collateralization ratio
// if the cdp is liquidated, the keeper that sent the transaction is rewarded a percentage of the collateral according to that collateral types'
// keeper reward percentage.
func (k Keeper) AttemptKeeperLiquidation(ctx sdk.Context, keeper, owner sdk.AccAddress, collateralType string) error {
	cdp, found := k.GetCdpByOwnerAndCollateralType(ctx, owner, collateralType)
	if !found {
		return sdkerrors.Wrapf(types.ErrCdpNotFound, "owner %s, denom %s", owner, collateralType)
	}
	k.hooks.BeforeCdpModified(ctx, cdp)
	cdp = k.SynchronizeInterest(ctx, cdp)

	err := k.ValidateLiquidation(ctx, cdp.Collateral, cdp.Type, cdp.Principal, cdp.AccumulatedFees)
	if err != nil {
		return err
	}
	cdp, err = k.payoutKeeperLiquidationReward(ctx, keeper, cdp)
	if err != nil {
		return err
	}
	return k.SeizeCollateral(ctx, cdp)
}

// SeizeCollateral liquidates the collateral in the input cdp.
// the following operations are performed:
// 1. Collateral for all deposits is sent from the cdp module to the liquidator module account
// 2. The liquidation penalty is applied
// 3. Debt coins are sent from the cdp module to the liquidator module account
// 4. The total amount of principal outstanding for that collateral type is decremented
// (this is the equivalent of saying that fees are no longer accumulated by a cdp once it gets liquidated)
func (k Keeper) SeizeCollateral(ctx sdk.Context, cdp types.Cdp) error {
	// Calculate the previous collateral ratio
	oldCollateralToDebtRatio := k.CalculateCollateralToDebtRatio(ctx, cdp.Collateral, cdp.Type, cdp.GetTotalPrincipal())

	// Move debt coins from cdp to liquidator account
	deposits := k.GetDeposits(ctx, cdp.Id)
	debt := cdp.GetTotalPrincipal().Amount
	debtDenomMap := k.GetDebtDenomMap(ctx)
	modAccountDebt := k.getModAccountDebt(ctx, types.ModuleName, debtDenomMap[cdp.Principal.Denom])
	debt = sdk.MinInt(debt, modAccountDebt)
	debtCoin := sdk.NewCoin(debtDenomMap[cdp.Principal.Denom], debt)
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.LiquidatorMacc, sdk.NewCoins(debtCoin))
	if err != nil {
		return err
	}

	// liquidate deposits and send collateral from cdp to liquidator
	for _, dep := range deposits {
		err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.LiquidatorMacc, sdk.NewCoins(dep.Amount))
		if err != nil {
			return err
		}
		k.DeleteDeposit(ctx, dep.CdpId, dep.Depositor.AccAddress())

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCdpLiquidation,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(types.AttributeKeyCdpID, fmt.Sprintf("%d", cdp.Id)),
				sdk.NewAttribute(types.AttributeKeyDeposit, dep.String()),
			),
		)
	}

	err = k.AuctionCollateral(ctx, deposits, cdp.Type, debt, cdp.Principal.Denom)
	if err != nil {
		return err
	}

	// Decrement total principal for this collateral type
	coinsToDecrement := cdp.GetTotalPrincipal()
	k.DecrementTotalPrincipal(ctx, cdp.Type, coinsToDecrement)

	// Delete Cdp from state
	k.RemoveCdpOwnerIndex(ctx, cdp)
	k.RemoveCdpCollateralRatioIndex(ctx, cdp.Type, cdp.Id, oldCollateralToDebtRatio)
	return k.DeleteCdp(ctx, cdp)
}

// LiquidateCdps seizes collateral from all Cdps below the input liquidation ratio
func (k Keeper) LiquidateCdps(ctx sdk.Context, marketID string, collateralType string, liquidationRatio sdk.Dec, count sdk.Int) error {
	price, err := k.pricefeedKeeper.GetCurrentPrice(ctx, marketID)
	if err != nil {
		return err
	}
	priceDivLiqRatio := price.Price.Quo(liquidationRatio)
	if priceDivLiqRatio.IsZero() {
		priceDivLiqRatio = sdk.SmallestDec()
	}
	// price = $0.5
	// liquidation ratio = 1.5
	// normalizedRatio = (1/(0.5/1.5)) = 3
	normalizedRatio := sdk.OneDec().Quo(priceDivLiqRatio)
	cdpsToLiquidate := k.GetSliceOfCdpsByRatioAndType(ctx, count, normalizedRatio, collateralType)
	for _, c := range cdpsToLiquidate {
		k.hooks.BeforeCdpModified(ctx, c)
		err := k.SeizeCollateral(ctx, c)
		if err != nil {
			return err
		}
	}
	return nil
}

// ApplyLiquidationPenalty multiplies the input debt amount by the liquidation penalty
func (k Keeper) ApplyLiquidationPenalty(ctx sdk.Context, collateralType string, debt sdk.Int) sdk.Int {
	penalty := k.getLiquidationPenalty(ctx, collateralType)
	return sdk.NewDecFromInt(debt).Mul(penalty).RoundInt()
}

// ValidateLiquidation validate that adding the input principal puts the cdp below the liquidation ratio
func (k Keeper) ValidateLiquidation(ctx sdk.Context, collateral sdk.Coin, collateralType string, principal sdk.Coin, fees sdk.Coin) error {
	collateralizationRatio, err := k.CalculateCollateralizationRatio(ctx, collateral, collateralType, principal, fees, spot)
	if err != nil {
		return err
	}
	liquidationRatio := k.getLiquidationRatio(ctx, collateralType)
	if collateralizationRatio.GT(liquidationRatio) {
		return sdkerrors.Wrapf(types.ErrNotLiquidatable, "collateral %s, collateral ratio %s, liquidation ratio %s", collateral.Denom, collateralizationRatio, liquidationRatio)
	}
	return nil
}

func (k Keeper) getModAccountDebt(ctx sdk.Context, accountName string, deb_denom string) sdk.Int {
	macc := k.accountKeeper.GetModuleAccount(ctx, accountName)
	return k.bankKeeper.GetAllBalances(ctx, macc.GetAddress()).AmountOf(deb_denom)
}

func (k Keeper) payoutKeeperLiquidationReward(ctx sdk.Context, keeper sdk.AccAddress, cdp types.Cdp) (types.Cdp, error) {
	collateralParam, found := k.GetCollateral(ctx, cdp.Type)
	if !found {
		return types.Cdp{}, sdkerrors.Wrapf(types.ErrInvalidCollateral, "%s", cdp.Type)
	}
	reward := sdk.NewDecFromInt(cdp.Collateral.Amount).Mul(collateralParam.KeeperRewardPercentage).RoundInt()
	rewardCoin := sdk.NewCoin(cdp.Collateral.Denom, reward)
	paidReward := false
	deposits := k.GetDeposits(ctx, cdp.Id)
	for _, dep := range deposits {
		if dep.Amount.IsGTE(rewardCoin) {
			dep.Amount = dep.Amount.Sub(rewardCoin)
			k.SetDeposit(ctx, dep)
			paidReward = true
			break
		}
	}
	if !paidReward {
		return cdp, nil
	}
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, keeper, sdk.NewCoins(rewardCoin))
	if err != nil {
		return types.Cdp{}, err
	}

	cdp.Collateral = cdp.Collateral.Sub(rewardCoin)
	ratio := k.CalculateCollateralToDebtRatio(ctx, cdp.Collateral, cdp.Type, cdp.GetTotalPrincipal())
	err = k.UpdateCdpAndCollateralRatioIndex(ctx, cdp, ratio)
	if err != nil {
		return types.Cdp{}, err
	}
	return cdp, nil
}
