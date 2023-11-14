package types

import "fmt"

func GetLPTokenDenom(strategyContract string, maturity uint64) string {
	return fmt.Sprintf("%s/vaults/%s/maturity/%d/lp", ModuleName, strategyContract, maturity)
}

func GetFixedYieldTrancheDenom(strategyContract string, maturity uint64) string {
	return fmt.Sprintf("%s/vaults/%s/maturity/%d/fy", ModuleName, strategyContract, maturity)
}

func GetLeveragedYieldTrancheDenom(strategyContract string, maturity uint64) string {
	return fmt.Sprintf("%s/vaults/%s/maturity/%d/lvy", ModuleName, strategyContract, maturity)
}
