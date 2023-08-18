package types

import "fmt"

func GetLPTokenDenom(vaultId uint64) string {
	return fmt.Sprintf("%s/vaults/%d", ModuleName, vaultId)
}

func GetVaultModuleAccountName(vaultId uint64) string {
	return fmt.Sprintf("%s/vaults/%d", ModuleName, vaultId)
}

func GetStrategyModuleAccountName(strategyId uint64) string {
	return fmt.Sprintf("%s/strategies/%d", ModuleName, strategyId)
}
