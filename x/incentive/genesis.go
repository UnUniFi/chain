package incentive

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/x/incentive/keeper"
	"github.com/lcnem/jpyx/x/incentive/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, supplyKeeper types.SupplyKeeper, cdpKeeper types.CdpKeeper, gs types.GenesisState) {

	// check if the module account exists
	moduleAcc := supplyKeeper.GetModuleAccount(ctx, types.IncentiveMacc)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.IncentiveMacc))
	}

	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	for _, rp := range gs.Params.JPYXMintingRewardPeriods {
		_, found := cdpKeeper.GetCollateral(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("jpyx minting collateral type %s not found in cdp collateral types", rp.CollateralType))
		}
		k.SetJPYXMintingRewardFactor(ctx, rp.CollateralType, sdk.ZeroDec())
	}

	for _, mrp := range gs.Params.HardSupplyRewardPeriods {
		newRewardIndexes := types.RewardIndexes{}
		for _, rc := range mrp.RewardsPerSecond {
			ri := types.NewRewardIndex(rc.Denom, sdk.ZeroDec())
			newRewardIndexes = append(newRewardIndexes, ri)
		}
		k.SetHardSupplyRewardIndexes(ctx, mrp.CollateralType, newRewardIndexes)
	}

	for _, mrp := range gs.Params.HardBorrowRewardPeriods {
		newRewardIndexes := types.RewardIndexes{}
		for _, rc := range mrp.RewardsPerSecond {
			ri := types.NewRewardIndex(rc.Denom, sdk.ZeroDec())
			newRewardIndexes = append(newRewardIndexes, ri)
		}
		k.SetHardBorrowRewardIndexes(ctx, mrp.CollateralType, newRewardIndexes)
	}

	for _, rp := range gs.Params.HardDelegatorRewardPeriods {
		k.SetHardDelegatorRewardFactor(ctx, rp.CollateralType, sdk.ZeroDec())
	}

	k.SetParams(ctx, gs.Params)

	for _, gat := range gs.JPYXAccumulationTimes {
		k.SetPreviousJPYXMintingAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}

	for _, gat := range gs.HardSupplyAccumulationTimes {
		k.SetPreviousHardSupplyRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}

	for _, gat := range gs.HardBorrowAccumulationTimes {
		k.SetPreviousHardBorrowRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}

	for _, gat := range gs.HardDelegatorAccumulationTimes {
		k.SetPreviousHardDelegatorRewardAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}

	for i, claim := range gs.JPYXMintingClaims {
		for j, ri := range claim.RewardIndexes {
			if ri.RewardFactor != sdk.ZeroDec() {
				gs.JPYXMintingClaims[i].RewardIndexes[j].RewardFactor = sdk.ZeroDec()
			}
		}
		k.SetJPYXMintingClaim(ctx, claim)
	}

	for i, claim := range gs.HardLiquidityProviderClaims {
		for j, mri := range claim.SupplyRewardIndexes {
			for k, ri := range mri.RewardIndexes {
				if ri.RewardFactor != sdk.ZeroDec() {
					gs.HardLiquidityProviderClaims[i].SupplyRewardIndexes[j].RewardIndexes[k].RewardFactor = sdk.ZeroDec()
				}
			}
		}
		for j, mri := range claim.BorrowRewardIndexes {
			for k, ri := range mri.RewardIndexes {
				if ri.RewardFactor != sdk.ZeroDec() {
					gs.HardLiquidityProviderClaims[i].BorrowRewardIndexes[j].RewardIndexes[k].RewardFactor = sdk.ZeroDec()
				}
			}
		}
		for j, ri := range claim.DelegatorRewardIndexes {
			if ri.RewardFactor != sdk.ZeroDec() {
				gs.HardLiquidityProviderClaims[i].DelegatorRewardIndexes[j].RewardFactor = sdk.ZeroDec()
			}
		}
		k.SetHardLiquidityProviderClaim(ctx, claim)
	}
}

// ExportGenesis export genesis state for incentive module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	params := k.GetParams(ctx)

	jpyxClaims := k.GetAllJPYXMintingClaims(ctx)
	hardClaims := k.GetAllHardLiquidityProviderClaims(ctx)

	synchronizedJpyxClaims := types.JPYXMintingClaims{}
	synchronizedHardClaims := types.HardLiquidityProviderClaims{}

	for _, jpyxClaim := range jpyxClaims {
		claim, err := k.SynchronizeJPYXMintingClaim(ctx, jpyxClaim)
		if err != nil {
			panic(err)
		}
		for i := range claim.RewardIndexes {
			claim.RewardIndexes[i].RewardFactor = sdk.ZeroDec()
		}
		synchronizedJpyxClaims = append(synchronizedJpyxClaims, claim)
	}

	for _, hardClaim := range hardClaims {
		k.SynchronizeHardLiquidityProviderClaim(ctx, hardClaim.Owner)
		claim, found := k.GetHardLiquidityProviderClaim(ctx, hardClaim.Owner)
		if !found {
			panic("hard liquidity provider claim should always be found after synchronization")
		}
		for i, bri := range claim.BorrowRewardIndexes {
			for j := range bri.RewardIndexes {
				claim.BorrowRewardIndexes[i].RewardIndexes[j].RewardFactor = sdk.ZeroDec()
			}
		}
		for i, sri := range claim.SupplyRewardIndexes {
			for j := range sri.RewardIndexes {
				claim.SupplyRewardIndexes[i].RewardIndexes[j].RewardFactor = sdk.ZeroDec()
			}
		}
		for i := range claim.DelegatorRewardIndexes {
			claim.DelegatorRewardIndexes[i].RewardFactor = sdk.ZeroDec()
		}
		synchronizedHardClaims = append(synchronizedHardClaims, claim)
	}

	var jpyxMintingGats GenesisAccumulationTimes
	for _, rp := range params.JPYXMintingRewardPeriods {
		pat, found := k.GetPreviousJPYXMintingAccrualTime(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("expected previous jpyx minting reward accrual time to be set in state for %s", rp.CollateralType))
		}
		gat := types.NewGenesisAccumulationTime(rp.CollateralType, pat)
		jpyxMintingGats = append(jpyxMintingGats, gat)
	}

	var hardSupplyGats GenesisAccumulationTimes
	for _, rp := range params.HardSupplyRewardPeriods {
		pat, found := k.GetPreviousHardSupplyRewardAccrualTime(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("expected previous hard supply reward accrual time to be set in state for %s", rp.CollateralType))
		}
		gat := types.NewGenesisAccumulationTime(rp.CollateralType, pat)
		hardSupplyGats = append(hardSupplyGats, gat)
	}

	var hardBorrowGats GenesisAccumulationTimes
	for _, rp := range params.HardBorrowRewardPeriods {
		pat, found := k.GetPreviousHardBorrowRewardAccrualTime(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("expected previous hard borrow reward accrual time to be set in state for %s", rp.CollateralType))
		}
		gat := types.NewGenesisAccumulationTime(rp.CollateralType, pat)
		hardBorrowGats = append(hardBorrowGats, gat)
	}

	var hardDelegatorGats GenesisAccumulationTimes
	for _, rp := range params.HardDelegatorRewardPeriods {
		pat, found := k.GetPreviousHardDelegatorRewardAccrualTime(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("expected previous hard delegator reward accrual time to be set in state for %s", rp.CollateralType))
		}
		gat := types.NewGenesisAccumulationTime(rp.CollateralType, pat)
		hardDelegatorGats = append(hardDelegatorGats, gat)
	}

	return types.NewGenesisState(params, jpyxMintingGats, hardSupplyGats,
		hardBorrowGats, hardDelegatorGats, synchronizedJpyxClaims, synchronizedHardClaims)
}
