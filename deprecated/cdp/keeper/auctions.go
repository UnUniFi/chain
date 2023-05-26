package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/deprecated/cdp/types"
)

const (
	// factor for setting the initial value of gov tokens to sell at debt auctions -- assuming stable token is ~1 eur, this starts the auction with a price of $0.01 KAVA
	dump = 100
)

// AuctionCollateral creates auctions from the input deposits which attempt to raise the corresponding amount of debt
func (k Keeper) AuctionCollateral(ctx sdk.Context, deposits types.Deposits, collateralType string, debt sdk.Int, bidDenom string) error {

	auctionSize := k.getAuctionSize(ctx, collateralType)
	totalCollateral := deposits.SumCollateral()
	for _, deposit := range deposits {
		debtCoveredByDeposit := (sdk.NewDecFromInt(deposit.Amount.Amount).Quo(sdk.NewDecFromInt(totalCollateral))).Mul(sdk.NewDecFromInt(debt)).RoundInt()
		err := k.CreateAuctionsFromDeposit(ctx, deposit.Amount, collateralType, deposit.Depositor.AccAddress(), debtCoveredByDeposit, auctionSize, bidDenom)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateAuctionsFromDeposit creates auctions from the input deposit
func (k Keeper) CreateAuctionsFromDeposit(
	ctx sdk.Context, collateral sdk.Coin, collateralType string, returnAddr sdk.AccAddress, debt, auctionSize sdk.Int,
	principalDenom string) error {

	// number of auctions of auctionSize
	numberOfAuctions := collateral.Amount.Quo(auctionSize)
	debtPerAuction := debt.Mul(auctionSize).Quo(collateral.Amount)

	// last auction for remaining collateral (collateral < auctionSize)
	lastAuctionCollateral := collateral.Amount.Mod(auctionSize)
	lastAuctionDebt := debt.Mul(lastAuctionCollateral).Quo(collateral.Amount)

	// amount of debt that has not been allocated due to
	// rounding error (unallocated debt is less than numberOfAuctions + 1)
	unallocatedDebt := debt.Sub(numberOfAuctions.Mul(debtPerAuction).Add(lastAuctionDebt))

	// rounding error for whole and last auctions in units of collateral
	// higher value means a larger truncation
	wholeAuctionError := debt.Mul(auctionSize).Mod(collateral.Amount)
	lastAuctionError := debt.Mul(lastAuctionCollateral).Mod(collateral.Amount)

	// if last auction has larger rounding error, then allocate one debt to last auction first
	// follows the largest remainder method https://en.wikipedia.org/wiki/Largest_remainder_method
	if lastAuctionError.GT(wholeAuctionError) {
		lastAuctionDebt = lastAuctionDebt.Add(sdk.OneInt())
		unallocatedDebt = unallocatedDebt.Sub(sdk.OneInt())
	}

	debtDenomMap := k.GetDebtDenomMap(ctx)
	numAuctions := numberOfAuctions.Int64()

	// create whole auctions
	for i := int64(0); i < numAuctions; i++ {
		debtAmount := debtPerAuction

		// distribute unallocated debt left over starting with first auction created
		if unallocatedDebt.IsPositive() {
			debtAmount = debtAmount.Add(sdk.OneInt())
			unallocatedDebt = unallocatedDebt.Sub(sdk.OneInt())
		}

		penalty := k.ApplyLiquidationPenalty(ctx, collateralType, debtAmount)

		_, err := k.auctionKeeper.StartCollateralAuction(
			ctx, types.LiquidatorMacc, sdk.NewCoin(collateral.Denom, auctionSize),
			sdk.NewCoin(principalDenom, debtAmount.Add(penalty)), []sdk.AccAddress{returnAddr},
			[]sdk.Int{auctionSize}, sdk.NewCoin(debtDenomMap[principalDenom], debtAmount),
		)

		if err != nil {
			return err
		}
	}

	// skip last auction if there is no collateral left to auction
	if !lastAuctionCollateral.IsPositive() {
		return nil
	}

	// if the last auction had a larger rounding error than whole auctions,
	// then unallocatedDebt will be zero since we will have already distributed
	// all of the unallocated debt
	if unallocatedDebt.IsPositive() {
		lastAuctionDebt = lastAuctionDebt.Add(sdk.OneInt())
		unallocatedDebt = unallocatedDebt.Sub(sdk.OneInt())
	}

	penalty := k.ApplyLiquidationPenalty(ctx, collateralType, lastAuctionDebt)

	_, err := k.auctionKeeper.StartCollateralAuction(
		ctx, types.LiquidatorMacc, sdk.NewCoin(collateral.Denom, lastAuctionCollateral),
		sdk.NewCoin(principalDenom, lastAuctionDebt.Add(penalty)), []sdk.AccAddress{returnAddr},
		[]sdk.Int{lastAuctionCollateral}, sdk.NewCoin(debtDenomMap[principalDenom], lastAuctionDebt),
	)

	return err
}

// NetSurplusAndDebt burns surplus and debt coins equal to the minimum of surplus and debt balances held by the liquidator module account
// for example, if there is 1000 debt and 100 surplus, 100 surplus and 100 debt are burned, netting to 900 debt
func (k Keeper) NetSurplusAndDebt(ctx sdk.Context) error {
	dps := k.GetParams(ctx).DebtParams
	debtDenomMap := k.GetDebtDenomMap(ctx)
	for _, dp := range dps {
		totalSurplus := k.GetTotalSurplus(ctx, types.LiquidatorMacc, dp.Denom)
		debt := k.GetTotalDebt(ctx, types.LiquidatorMacc, debtDenomMap[dp.Denom])
		netAmount := sdk.MinInt(totalSurplus, debt)
		if netAmount.IsZero() {
			return nil
		}
		// burn debt coins equal to netAmount
		err := k.bankKeeper.BurnCoins(ctx, types.LiquidatorMacc, sdk.NewCoins(sdk.NewCoin(debtDenomMap[dp.Denom], netAmount)))
		if err != nil {
			return err
		}

		// burn stable coins equal to min(balance, netAmount)
		macc := k.accountKeeper.GetModuleAccount(ctx, types.LiquidatorMacc)
		balance := k.bankKeeper.GetAllBalances(ctx, macc.GetAddress()).AmountOf(dp.Denom)
		burnAmount := sdk.MinInt(balance, netAmount)
		k.bankKeeper.BurnCoins(ctx, types.LiquidatorMacc, sdk.NewCoins(sdk.NewCoin(dp.Denom, burnAmount)))
	}
	return nil
}

// GetTotalSurplus returns the total amount of surplus tokens held by the liquidator module account
func (k Keeper) GetTotalSurplus(ctx sdk.Context, accountName string, denom string) sdk.Int {
	return k.getTotalDenom(ctx, accountName, denom)
}

// GetTotalDebt returns the total amount of debt tokens held by the liquidator module account
func (k Keeper) GetTotalDebt(ctx sdk.Context, accountName string, deb_denom string) sdk.Int {
	return k.getTotalDenom(ctx, accountName, deb_denom)
}

func (k Keeper) getTotalDenom(ctx sdk.Context, accountName string, denom string) sdk.Int {
	acc := k.accountKeeper.GetModuleAccount(ctx, accountName)
	return k.bankKeeper.GetAllBalances(ctx, acc.GetAddress()).AmountOf(denom)
}

// RunSurplusAndDebtAuctions nets the surplus and debt balances and then creates surplus or debt auctions if the remaining balance is above the auction threshold parameter
func (k Keeper) RunSurplusAndDebtAuctions(ctx sdk.Context) error {
	if err := k.NetSurplusAndDebt(ctx); err != nil {
		return err
	}
	debtDenomMap := k.GetDebtDenomMap(ctx)
	dps := k.GetParams(ctx).DebtParams
	for _, dp := range dps {
		remainingDebt := k.GetTotalDebt(ctx, types.LiquidatorMacc, debtDenomMap[dp.Denom])
		if remainingDebt.GTE(dp.DebtAuctionThreshold) {
			debtLot := sdk.NewCoin(debtDenomMap[dp.Denom], dp.DebtAuctionLot)
			bidCoin := sdk.NewCoin(dp.Denom, debtLot.Amount)
			initialLot := sdk.NewCoin(k.GetGovDenom(ctx), debtLot.Amount.Mul(sdk.NewInt(dump)))

			_, err := k.auctionKeeper.StartDebtAuction(ctx, types.LiquidatorMacc, bidCoin, initialLot, debtLot)
			if err != nil {
				return err
			}
		}

		macc := k.accountKeeper.GetModuleAccount(ctx, types.LiquidatorMacc)
		surplus := k.bankKeeper.GetAllBalances(ctx, macc.GetAddress()).AmountOf(dp.Denom)
		if !surplus.GTE(dp.SurplusAuctionThreshold) {
			continue
		}

		surplusLot := sdk.NewCoin(dp.Denom, sdk.MinInt(dp.SurplusAuctionLot, surplus))
		_, err := k.auctionKeeper.StartSurplusAuction(ctx, types.LiquidatorMacc, surplusLot, k.GetGovDenom(ctx))
		if err != nil {
			return err
		}
	}
	return nil
}
