package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

func (k Keeper) GetStrategy(principalDenom string, id string) types.Strategy {
	panic("implement me")
}

func (k Keeper) GetStrategies(principalDenom string) []types.Strategy {
	panic("implement me")
}

func (k Keeper) SetStrategy(strategy types.Strategy) {

}

func (k Keeper) DeleteStrategy(principalDenom string, id string) {

}

func (k Keeper) StakeToStrategy(principalDenom string, id string, amount sdk.Int) {
	// call `stake` function of the strategy contract
}

func (k Keeper) UnstakeFromStrategy(principalDenom string, id string, amount sdk.Int) {
	// call `unstake` function of the strategy contract
}

func (k Keeper) GetAPRFromStrategy(principalDenom string, id string) {
	// call `get_apr` function of the strategy contract
}

func (k Keeper) GetPerformanceFeeRate(principalDenom string, id string) {
	// call `get_performance_fee_rate` function of the strategy contract
}
