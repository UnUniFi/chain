package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/pricefeed/types"
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

// GetOracles returns the oracles in the pricefeed store
func (k Keeper) GetOracles(ctx sdk.Context, marketID string) ([]string, error) {
	for _, m := range k.GetMarkets(ctx) {
		if marketID == m.MarketId {
			return m.Oracles, nil
		}
	}
	return nil, sdkerrors.Wrap(types.ErrInvalidMarket, marketID)
}

// GetOracle returns the oracle from the store or an error if not found
func (k Keeper) GetOracle(ctx sdk.Context, marketID string, address sdk.AccAddress) (sdk.AccAddress, error) {
	oracles, err := k.GetOracles(ctx, marketID)
	if err != nil {
		return nil, err
	}
	for _, oracle := range oracles {
		addr, err := sdk.AccAddressFromBech32(oracle)
		if err != nil {
			return sdk.AccAddress{}, err
		}
		if address.Equals(addr) {
			return addr, nil
		}
	}
	return nil, sdkerrors.Wrap(types.ErrInvalidOracle, address.String())
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
func (k Keeper) GetAuthorizedAddresses(ctx sdk.Context) ([]sdk.AccAddress, error) {
	oracles := []sdk.AccAddress{}
	uniqueOracles := map[string]bool{}

	for _, m := range k.GetMarkets(ctx) {
		for _, o := range m.Oracles {
			// de-dup list of oracles
			if _, found := uniqueOracles[o]; !found {
				oAddress, err := sdk.AccAddressFromBech32(o)
				if err != nil {
					return nil, err
				}
				oracles = append(oracles, oAddress)
			}
			uniqueOracles[o] = true
		}
	}
	return oracles, nil
}
