package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func (k Keeper) GetPairRate(ctx sdk.Context, pair types.Market) (*sdk.Dec, error) {
	marketId, err := k.pricefeedKeeper.GetMarketIdFromDenom(ctx, pair.Denom, pair.QuoteDenom)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s", marketId)
	price, err := k.pricefeedKeeper.GetCurrentPrice(ctx, marketId)

	return &price.Price, err
}

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
