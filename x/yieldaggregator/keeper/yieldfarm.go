package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// TODO: implementation should be following the type of asset management target
func (k Keeper) InvestOnTarget(ctx sdk.Context, target types.AssetManagementTarget, farmUnit types.FarmingUnit) error {
	// set farming unit
	k.SetFarmingUnit(ctx, farmUnit)

	// move tokens to farm target
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		addr, err := sdk.AccAddressFromBech32(farmUnit.Owner)
		if err != nil {
			return err
		}
		err = k.yieldfarmKeeper.Deposit(ctx, addr, farmUnit.Amount)
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
		addr, err := sdk.AccAddressFromBech32(farmUnit.Owner)
		if err != nil {
			return err
		}
		err = k.yieldfarmKeeper.Withdraw(ctx, addr, farmUnit.Amount)
		if err != nil {
			return err
		}
	case types.IntegrateType_COSMWASM:
		// TODO: implement begin withdraw flow in case of cosmwasm
	}
	return nil
}

func (k Keeper) ClaimWithdrawFromTarget(ctx sdk.Context, target types.AssetManagementTarget, unit types.FarmingUnit) {
	// TODO: check unbonding time passed
	// TODO: destroy farming unit and increase users' deposit balance
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
	case types.IntegrateType_COSMWASM:
	}
}

func (k Keeper) ClaimRewardsFromTarget(ctx sdk.Context, target types.AssetManagementTarget) error {
	// TODO: claim and assign rewards to farm units
	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		addr, err := sdk.AccAddressFromBech32(target.AccountAddress)
		if err != nil {
			return err
		}
		amounts := k.yieldfarmKeeper.ClaimRewards(ctx, addr)
		_ = amounts
	case types.IntegrateType_COSMWASM:
	}
	return nil
}
