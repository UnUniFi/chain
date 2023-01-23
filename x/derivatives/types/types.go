package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (marketCap PoolMarketCap) CalculateLPTokenPrice(supply sdk.Int) sdk.Dec {
	return marketCap.Total.Quo(sdk.NewDecFromInt(supply))
}
