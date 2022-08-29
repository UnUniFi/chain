package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) InvestOnTarget(ctx sdk.Context, addr sdk.AccAddress, target types.AssetManagementTarget, amount sdk.Coins) error {
	farmingUnit := k.GetFarmingUnit(ctx, addr.String(), target.AssetManagementAccountId, target.Id)
	// set farming unit if does not exists
	if farmingUnit.AccountId == "" {
		farmingUnit = types.FarmingUnit{
			AccountId:          target.AssetManagementAccountId,
			TargetId:           target.Id,
			Amount:             amount,
			FarmingStartTime:   ctx.BlockTime().String(),
			UnbondingStarttime: time.Time{},
			Owner:              addr.String(),
		}
		k.SetFarmingUnit(ctx, farmingUnit)
	} else {
		farmingUnit.Amount = sdk.Coins(farmingUnit.Amount).Add(amount...)
	}

	// move tokens to farm target
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		address := farmingUnit.GetAddress()
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, amount)
		if err != nil {
			return err
		}
		err = k.yieldfarmKeeper.Deposit(ctx, address, amount)
		if err != nil {
			return err
		}
	case types.IntegrateType_COSMWASM:
		// TODO: implement investment flow in case of cosmwasm
	}
	return nil
}

func (k Keeper) BeginWithdrawFromTarget(ctx sdk.Context, addr sdk.AccAddress, target types.AssetManagementTarget, amount sdk.Coins) error {
	farmingUnit := k.GetFarmingUnit(ctx, addr.String(), target.AssetManagementAccountId, target.Id)
	if farmingUnit.AccountId == "" {
		return types.ErrFarmingUnitDoesNotExist
	}

	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		address := farmingUnit.GetAddress()

		// request full withdraw from target if amount is empty
		if amount.String() == "" {
			farmerInfo := k.yieldfarmKeeper.GetFarmerInfo(ctx, address)
			amount = farmerInfo.Amount
		}
		err := k.yieldfarmKeeper.Withdraw(ctx, address, amount)
		if err != nil {
			return err
		}
	case types.IntegrateType_COSMWASM:
		// TODO: implement begin withdraw flow in case of cosmwasm
	}
	return nil
}

func (k Keeper) ClaimWithdrawFromTarget(ctx sdk.Context, addr sdk.AccAddress, target types.AssetManagementTarget) error {
	farmingUnit := k.GetFarmingUnit(ctx, addr.String(), target.AssetManagementAccountId, target.Id)
	if farmingUnit.AccountId == "" {
		return types.ErrFarmingUnitDoesNotExist
	}

	// check unbonding time passed
	if farmingUnit.UnbondingStarttime.Add(target.UnbondingTime).After(ctx.BlockTime()) {
		return types.ErrUnbondingTimeNotPassed
	}

	// withdraw from farming unit and increase users' deposit balance
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		address := farmingUnit.GetAddress()
		balances := k.bankKeeper.GetAllBalances(ctx, address)
		if balances.IsAllPositive() {
			err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, balances)
			if err != nil {
				return err
			}
		}
		k.IncreaseUserDeposit(ctx, addr, balances)
	case types.IntegrateType_COSMWASM:
		// TODO: implement claim withdraw flow in case of cosmwasm
	}
	return nil
}

func (k Keeper) ClaimRewardsFromTarget(ctx sdk.Context, addr sdk.AccAddress, target types.AssetManagementTarget) error {
	farmingUnit := k.GetFarmingUnit(ctx, addr.String(), target.AssetManagementAccountId, target.Id)
	if farmingUnit.AccountId == "" {
		return types.ErrFarmingUnitDoesNotExist
	}

	// claim and assign rewards to farm units
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		address := farmingUnit.GetAddress()
		k.yieldfarmKeeper.ClaimRewards(ctx, address)
		balances := k.bankKeeper.GetAllBalances(ctx, address)
		if balances.IsAllPositive() {
			err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, balances)
			if err != nil {
				return err
			}
		}
		k.IncreaseUserDeposit(ctx, addr, balances)
	case types.IntegrateType_COSMWASM:
		// TODO: implement claim rewards flow in case of cosmwasm
	}
	return nil
}

func (k Keeper) ClaimAllFarmUnitRewards(ctx sdk.Context) {
	// iterate and run ClaimRewardsFromTarget
	farmUnits := k.GetAllFarmingUnits(ctx)
	for _, farmUnit := range farmUnits {
		target := k.GetAssetManagementTarget(ctx, farmUnit.AccountId, farmUnit.TargetId)
		addr := sdk.MustAccAddressFromBech32(farmUnit.Owner)
		err := k.ClaimRewardsFromTarget(ctx, addr, target)
		if err != nil {
			continue
		}
	}
}
