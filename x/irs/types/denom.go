package types

import "fmt"

func LsDenom(strategyContract string, maturity uint64) string {
	return fmt.Sprintf("%s/vaults/%s/maturity/%d/ls", ModuleName, strategyContract, maturity)
}

func PtDenom(strategyContract string, maturity uint64) string {
	return fmt.Sprintf("%s/vaults/%s/maturity/%d/pt", ModuleName, strategyContract, maturity)
}

func YtDenom(strategyContract string, maturity uint64) string {
	return fmt.Sprintf("%s/vaults/%s/maturity/%d/yt", ModuleName, strategyContract, maturity)
}
