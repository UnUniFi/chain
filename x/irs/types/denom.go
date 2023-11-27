package types

import "fmt"

func LsDenom(pool TranchePool) string {
	return fmt.Sprintf("%s/tranche/%d/ls", ModuleName, pool.Id)
}

func PtDenom(pool TranchePool) string {
	return fmt.Sprintf("%s/tranche/%d/pt", ModuleName, pool.Id)
}

func YtDenom(pool TranchePool) string {
	return fmt.Sprintf("%s/tranche/%d/yt", ModuleName, pool.Id)
}
