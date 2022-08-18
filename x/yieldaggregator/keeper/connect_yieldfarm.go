package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) InvestOnTarget(ctx sdk.Context, target types.AssetManagementTarget, farmUnit types.FarmingUnit) error {
	// TODO: implementation should be following the type of asset management target
	// set farming unit
	k.SetFarmingUnit(ctx, farmUnit)

	// move tokens to farm target
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		address := farmUnit.GetAddress()
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, farmUnit.Amount)
		if err != nil {
			return err
		}
		err = k.yieldfarmKeeper.Deposit(ctx, address, farmUnit.Amount)
		if err != nil {
			return err
		}
	case types.IntegrateType_COSMWASM:
		// TODO: implement investment flow in case of cosmwasm
	}
	return nil
}

func (k Keeper) BeginWithdrawFromTarget(ctx sdk.Context, target types.AssetManagementTarget, farmUnit types.FarmingUnit) error {
	// TODO: request withdrawal from target by unit amount
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		address := farmUnit.GetAddress()
		err := k.yieldfarmKeeper.Withdraw(ctx, address, farmUnit.Amount)
		if err != nil {
			return err
		}
	case types.IntegrateType_COSMWASM:
		// TODO: implement begin withdraw flow in case of cosmwasm
	}
	return nil
}

func (k Keeper) ClaimWithdrawFromTarget(ctx sdk.Context, target types.AssetManagementTarget, farmUnit types.FarmingUnit) error {
	// TODO: check unbonding time passed
	// TODO: destroy farming unit and increase users' deposit balance
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		address := farmUnit.GetAddress()
		balances := k.bankKeeper.GetAllBalances(ctx, address)
		if balances.IsAllPositive() {
			err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, balances)
			if err != nil {
				return err
			}
		}
		farmUnit.Amount = sdk.Coins(farmUnit.Amount).Add(balances...)
		k.SetFarmingUnit(ctx, farmUnit)
	case types.IntegrateType_COSMWASM:
	}
	return nil
}

func (k Keeper) ClaimRewardsFromTarget(ctx sdk.Context, target types.AssetManagementTarget, farmUnit types.FarmingUnit) error {
	// claim and assign rewards to farm units
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		address := farmUnit.GetAddress()
		amounts := k.yieldfarmKeeper.ClaimRewards(ctx, address)
		_ = amounts
	case types.IntegrateType_COSMWASM:
	}
	return nil
}

func (k Keeper) ClaimAllFarmUnitRewards(ctx sdk.Context) {
	// iterate and run ClaimRewardsFromTarget
	farmUnits := k.GetAllFarmingUnits(ctx)
	for _, farmUnit := range farmUnits {
		target := k.GetAssetManagementTarget(ctx, farmUnit.AccountId, farmUnit.TargetId)
		err := k.ClaimRewardsFromTarget(ctx, target, farmUnit)
		if err != nil {
			continue
		}
	}
}
