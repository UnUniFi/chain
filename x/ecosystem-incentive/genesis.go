package ecosystemincentive

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystem-incentive/keeper"
	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

// InitGenesis initializes the capability module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	for _, incentiveUnit := range genState.IncentiveUnits {
		var subjectAddrs []ununifitypes.StringAccAddress
		var weights []sdk.Dec
		for i := 0; i < len(incentiveUnit.SubjectInfoLists); i++ {
			subjectAddrs = append(subjectAddrs, incentiveUnit.SubjectInfoLists[i].SubjectAddr)
			weights = append(weights, incentiveUnit.SubjectInfoLists[i].Weight)
		}

		if _, err := k.Register(ctx, &types.MsgRegister{
			IncentiveUnitId: incentiveUnit.Id,
			SubjectAddrs:    subjectAddrs,
			Weights:         weights,
		}); err != nil {
			panic(err)
		}
	}

	for _, rewardStore := range genState.RewardStores {
		if err := k.SetRewardStore(ctx, rewardStore); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.IncentiveUnits = k.GetAllIncentiveUnits(ctx)
	genesis.RewardStores = k.GetAllRewardStores(ctx)
	genesis.IncentiveUnitIdsByAddr = k.GetAllIncentiveUnitIdsByAddrs(ctx)

	return genesis
}
