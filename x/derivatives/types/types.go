package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (marketCap PoolMarketCap) CalculateLPTokenPrice(supply sdk.Dec) sdk.Dec {
	// TODO: gen-proto is needed for .Total
	return marketCap.Total.Quo(supply)
}
