package types

import "fmt"

func GetLPTokenDenom(vaultId uint64) string {
	return fmt.Sprintf("yield-aggregator/vaults/%d", vaultId)
}

func GetModuleAccountName(vaultId uint64) string {
	return fmt.Sprintf("%s/vaults/%d", ModuleName, vaultId)
}
