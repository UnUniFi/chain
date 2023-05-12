package incentive

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/deprecated/incentive/keeper"
	"github.com/UnUniFi/chain/x/deprecated/incentive/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, cdpKeeper types.CdpKeeper, gs types.GenesisState) {

	// check if the module account exists
	moduleAcc := accountKeeper.GetModuleAccount(ctx, types.IncentiveMacc)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.IncentiveMacc))
	}

	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	for _, rp := range gs.Params.CdpMintingRewardPeriods {
		_, found := cdpKeeper.GetCollateral(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("cdp minting collateral type %s not found in cdp collateral types", rp.CollateralType))
		}
		k.SetCdpMintingRewardFactor(ctx, rp.CollateralType, sdk.ZeroDec())
	}

	k.SetParams(ctx, gs.Params)

	for _, gat := range gs.CdpAccumulationTimes {
		k.SetPreviousCdpMintingAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}

	for i, claim := range gs.CdpMintingClaims {
		for j, ri := range claim.RewardIndexes {
			if ri.RewardFactor != sdk.ZeroDec() {
				gs.CdpMintingClaims[i].RewardIndexes[j].RewardFactor = sdk.ZeroDec()
			}
		}
		k.SetCdpMintingClaim(ctx, claim)
	}

	k.SetGenesisDenoms(ctx, gs.Denoms)
}

// ExportGenesis export genesis state for incentive module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	params := k.GetParams(ctx)

	cdpClaims := k.GetAllCdpMintingClaims(ctx)

	synchronizedCdpClaims := types.CdpMintingClaims{}

	for _, jpuClaim := range cdpClaims {
		claim, err := k.SynchronizeCdpMintingClaim(ctx, jpuClaim)
		if err != nil {
			panic(err)
		}
		for i := range claim.RewardIndexes {
			claim.RewardIndexes[i].RewardFactor = sdk.ZeroDec()
		}
		synchronizedCdpClaims = append(synchronizedCdpClaims, claim)
	}

	var cdpMintingGats types.GenesisAccumulationTimes
	for _, rp := range params.CdpMintingRewardPeriods {
		pat, found := k.GetPreviousCdpMintingAccrualTime(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("expected previous jpu minting reward accrual time to be set in state for %s", rp.CollateralType))
		}
		gat := types.NewGenesisAccumulationTime(rp.CollateralType, pat)
		cdpMintingGats = append(cdpMintingGats, gat)
	}

	denoms, found := k.GetGenesisDenoms(ctx)
	if !found {
		denoms = types.DefaultGenesisDenoms()
	}

	return types.NewGenesisState(params, cdpMintingGats, synchronizedCdpClaims, denoms)
}
