package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/pricefeed/types"

	ununifitypes "github.com/UnUniFi/chain/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetMarkets returns the markets from params
func (k Keeper) GetMarkets(ctx sdk.Context) []types.Market {
	return k.GetParams(ctx).Markets
}

func (k Keeper) GetDenomPairs(ctx sdk.Context) []types.DenomPair {
	return k.GetParams(ctx).DenomPairs
}

func (k Keeper) GetMarketDenom(ctx sdk.Context, internalDenom string) (string, error) {
	for _, pair := range k.GetDenomPairs(ctx) {
		if internalDenom == pair.InternalDenom {
			return pair.MarketDenom, nil
		}
	}
	return "", sdkerrors.Wrap(types.ErrInternalDenomNotFound, internalDenom)
}

// GetOracles returns the oracles in the pricefeed store
func (k Keeper) GetOracles(ctx sdk.Context, marketID string) ([]sdk.AccAddress, error) {
	for _, m := range k.GetMarkets(ctx) {
		if marketID == m.MarketId {
			return ununifitypes.AccAddresses(m.Oracles), nil
		}
	}
	return []sdk.AccAddress{}, sdkerrors.Wrap(types.ErrInvalidMarket, marketID)
}

// GetOracle returns the oracle from the store or an error if not found
func (k Keeper) GetOracle(ctx sdk.Context, marketID string, address sdk.AccAddress) (sdk.AccAddress, error) {
	oracles, err := k.GetOracles(ctx, marketID)
	if err != nil {
		return sdk.AccAddress{}, sdkerrors.Wrap(types.ErrInvalidMarket, marketID)
	}
	for _, addr := range oracles {
		if address.Equals(addr) {
			return addr, nil
		}
	}
	return sdk.AccAddress{}, sdkerrors.Wrap(types.ErrInvalidOracle, address.String())
}

// GetMarket returns the market if it is in the pricefeed system
func (k Keeper) GetMarket(ctx sdk.Context, marketID string) (types.Market, bool) {
	markets := k.GetMarkets(ctx)

	for i := range markets {
		if markets[i].MarketId == marketID {
			return markets[i], true
		}
	}
	return types.Market{}, false
}

// GetAuthorizedAddresses returns a list of addresses that have special authorization within this module, eg the oracles of all markets.
func (k Keeper) GetAuthorizedAddresses(ctx sdk.Context) []sdk.AccAddress {
	oracles := []sdk.AccAddress{}
	uniqueOracles := map[string]bool{}

	for _, m := range k.GetMarkets(ctx) {
		for _, o := range m.Oracles {
			// de-dup list of oracles
			if _, found := uniqueOracles[o.AccAddress().String()]; !found {
				oracles = append(oracles, o.AccAddress())
			}
			uniqueOracles[o.AccAddress().String()] = true
		}
	}
	return oracles
}
