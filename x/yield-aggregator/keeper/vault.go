package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

func (k Keeper) GetVault(denom string) types.Vault {
	panic("implement me")
}

func (k Keeper) GetVaults() []types.Vault {
	panic("implement me")
}

func (k Keeper) SetVault(vault types.Vault) {

}

func (k Keeper) DeleteVault(denom string) {

}

func (k Keeper) GetAPY(denom string) sdk.Dec {
	strategies := k.GetStrategies(denom)
	sum := sdk.ZeroDec()

	for _, strategy := range strategies {
		sum = sum.Add(strategy.Weight.Mul(strategy.Metrics.Apr))
	}

	return sum
}

func (k Keeper) DepositToVault(sender string, amount sdk.Coin) {
	strategies := k.GetStrategies(amount.Denom)

	for _, strategy := range strategies {
		allocation := strategy.Weight.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()
		k.StakeToStrategy(amount.Denom, strategy.Id, allocation)
	}
}

func (k Keeper) WithdrawFromVault(sender string, principalDenom string, LpTokenAmount sdk.Int) {

}
