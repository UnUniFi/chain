package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/pricefeed/types"
)

func (k Keeper) GetTicker(ctx sdk.Context, denom string) (string, error) {
	metadata, exists := k.bankKeeper.GetDenomMetaData(ctx, denom)
	if !exists {
		return "", sdkerrors.Wrap(types.ErrInternalDenomNotFound, denom)
	}
	return metadata.Base, nil
}

func (k Keeper) GetMarketId(ctx sdk.Context, lhsTicker string, rhsTicker string) string {
	return fmt.Sprintf("%s:%s", lhsTicker, rhsTicker)
}

func (k Keeper) GetMarketIdFromDenom(ctx sdk.Context, lhsDenom string, rhsDenom string) (string, error) {
	lhsTicker, err := k.GetTicker(ctx, lhsDenom)
	if err != nil {
		return "", err
	}
	rhsTicker, err := k.GetTicker(ctx, rhsDenom)
	if err != nil {
		return "", err
	}

	return k.GetMarketId(ctx, lhsTicker, rhsTicker), nil
}
