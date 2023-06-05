package ecosystemincentive

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/keeper"
	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

// InitGenesis initializes the capability module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	for _, container := range genState.RecipientContainers {
		var subjectAddrs []string
		var weights []sdk.Dec
		for i := 0; i < len(container.WeightedAddresses); i++ {
			subjectAddrs = append(subjectAddrs, container.WeightedAddresses[i].Address)
			weights = append(weights, container.WeightedAddresses[i].Weight)
		}

		if _, err := k.Register(ctx, &types.MsgRegister{
			IncentiveUnitId: container.Id,
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
	genesis.RecipientContainers = k.GetAllIncentiveUnits(ctx)
	genesis.RewardStores = k.GetAllRewardStores(ctx)

	return genesis
}
