package keeper

import (
	"time"

	stakeibckeeper "github.com/UnUniFi/chain/x/stakeibc/keeper"
	stakeibctypes "github.com/UnUniFi/chain/x/stakeibc/types"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		switch target.IntegrateInfo.ModName {
		case "stakeibc":
			for _, token := range amount {
				msg := &stakeibctypes.MsgLiquidStake{
					Creator:   addr.String(),
					Amount:    token.Amount.Uint64(),
					HostDenom: token.Denom,
				}

				msgServer := stakeibckeeper.NewMsgServerImpl(k.stakeibcKeeper)
				_, err := msgServer.LiquidStake(
					sdk.WrapSDKContext(ctx),
					msg,
				)
				if err != nil {
					return err
				}
			}
		default:
			err = k.yieldfarmKeeper.Deposit(ctx, address, amount)
			if err != nil {
				return err
			}
		}
	case types.IntegrateType_COSMWASM:
		wasmMsg := `{"deposit_native_token":{}}`
		contractAddr := sdk.MustAccAddressFromBech32(target.AccountAddress)
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, farmingUnit.GetAddress(), []byte(wasmMsg), amount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) BeginWithdrawFromTarget(ctx sdk.Context, addr sdk.AccAddress, target types.AssetManagementTarget, amount sdk.Coins) error {
	farmingUnit := k.GetFarmingUnit(ctx, addr.String(), target.AssetManagementAccountId, target.Id)
	if farmingUnit.AccountId == "" {
		return types.ErrFarmingUnitDoesNotExist
	}
	farmingUnit.UnbondingStarttime = ctx.BlockTime()
	k.SetFarmingUnit(ctx, farmingUnit)

	switch target.IntegrateInfo.Type {
	case types.IntegrateType_GOLANG_MOD:
		address := farmingUnit.GetAddress()

		switch target.IntegrateInfo.ModName {
		case "stakeibc":
			hostZones := k.stakeibcKeeper.GetAllHostZone(ctx)
			for _, zone := range hostZones {
				msg := stakeibctypes.NewMsgRedeemStake(
					address.String(),
					amount.AmountOf(zone.IBCDenom).Uint64(),
					zone.ChainId,
					address.String(),
				)
				msgServer := stakeibckeeper.NewMsgServerImpl(k.stakeibcKeeper)
				_, err := msgServer.RedeemStake(
					sdk.WrapSDKContext(ctx),
					msg,
				)
				if err != nil {
					return err
				}
			}
		default:
			// request full withdraw from target if amount is empty
			if amount.String() == "" {
				farmerInfo := k.yieldfarmKeeper.GetFarmerInfo(ctx, address)
				amount = farmerInfo.Amount
			}
			err := k.yieldfarmKeeper.Withdraw(ctx, address, amount)
			if err != nil {
				return err
			}
		}
	case types.IntegrateType_COSMWASM:
		wasmMsg := `{"start_unbond":{}}`
		contractAddr := sdk.MustAccAddressFromBech32(target.AccountAddress)
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, farmingUnit.GetAddress(), []byte(wasmMsg), sdk.Coins{})
		if err != nil {
			return err
		}
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
		wasmMsg := `{"claim_unbond":{}}`
		contractAddr := sdk.MustAccAddressFromBech32(target.AccountAddress)
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, farmingUnit.GetAddress(), []byte(wasmMsg), sdk.Coins{})
		if err != nil {
			return err
		}

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
		wasmMsg := `{"claim_all_rewards":{}}`
		contractAddr := sdk.MustAccAddressFromBech32(target.AccountAddress)
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, farmingUnit.GetAddress(), []byte(wasmMsg), sdk.Coins{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) ClaimAllFarmUnitRewards(ctx sdk.Context) {
	// iterate and run ClaimRewardsFromTarget
	farmUnits := k.GetAllFarmingUnits(ctx)
	for _, farmUnit := range farmUnits {
		target := k.GetAssetManagementTarget(ctx, farmUnit.AccountId, farmUnit.TargetId)
		addr, err := sdk.AccAddressFromBech32(farmUnit.Owner)
		if err != nil {
			continue
		}
		err = k.ClaimRewardsFromTarget(ctx, addr, target)
		if err != nil {
			continue
		}
	}
}
