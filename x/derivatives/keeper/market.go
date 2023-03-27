package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

// Unit is same to the quote ticker of pool metrics (USD in default)
func (k Keeper) GetAssetPrice(ctx sdk.Context, denom string) (*pftypes.CurrentPrice, error) {
	ticker, err := k.pricefeedKeeper.GetTicker(ctx, denom)
	if err != nil {
		return nil, err
	}
	quoteTicker := k.GetPoolQuoteTicker(ctx)

	price, err := k.GetPrice(ctx, ticker, quoteTicker)

	return &price, err
}

func (k Keeper) GetPrice(ctx sdk.Context, lhsTicker string, rhsTicker string) (pftypes.CurrentPrice, error) {
	marketId := fmt.Sprintf("%s:%s", lhsTicker, rhsTicker)
	return k.pricefeedKeeper.GetCurrentPrice(ctx, marketId)
}
