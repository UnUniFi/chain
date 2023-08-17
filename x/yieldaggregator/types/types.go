package types

import "fmt"

// TODO: want to %s with types.Module name. state breaking
func GetLPTokenDenom(vaultId uint64) string {
	return fmt.Sprintf("yield-aggregator/vaults/%d", vaultId)
}

func GetVaultModuleAccountName(vaultId uint64) string {
	return fmt.Sprintf("%s/vaults/%d", ModuleName, vaultId)
}

func GetStrategyModuleAccountName(strategyId uint64) string {
	return fmt.Sprintf("%s/strategies/%d", ModuleName, strategyId)
}
