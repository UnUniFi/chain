package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMarketValidate(t *testing.T) {
	mockPrivKey := ed25519.GenPrivKey()
	pubkey := mockPrivKey.PubKey()
	addr := sdk.AccAddress(pubkey.Address())

	testCases := []struct {
		msg     string
		market  Market
		expPass bool
	}{
		{
			"valid market",
			Market{
				MarketId:   "market",
				BaseAsset:  "xrp",
				QuoteAsset: "bnb",
				Oracles:    types.StringAccAddresses([]sdk.AccAddress{addr}),
				Active:     true,
			},
			true,
		},
		{
			"invalid id",
			Market{
				MarketId: " ",
			},
			false,
		},
		{
			"invalid base asset",
			Market{
				MarketId:  "market",
				BaseAsset: "XRP%($",
			},
			false,
		},
		{
			"invalid market",
			Market{
				MarketId:   "market",
				BaseAsset:  "xrp",
				QuoteAsset: "BNB%($",
			},
			false,
		},
		{
			"empty oracle address ",
			Market{
				MarketId:   "market",
				BaseAsset:  "xrp",
				QuoteAsset: "bnb",
				Oracles:    types.StringAccAddresses([]sdk.AccAddress{nil}),
			},
			false,
		},
		{
			"empty oracle address ",
			Market{
				MarketId:   "market",
				BaseAsset:  "xrp",
				QuoteAsset: "bnb",
				Oracles:    types.StringAccAddresses([]sdk.AccAddress{addr, addr}),
			},
			false,
		},
	}

	for _, tc := range testCases {
		err := tc.market.Validate()
		if tc.expPass {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}

func TestPostedPriceValidate(t *testing.T) {
	now := time.Now()
	mockPrivKey := ed25519.GenPrivKey()
	pubkey := mockPrivKey.PubKey()
	addr := sdk.AccAddress(pubkey.Address())

	testCases := []struct {
		msg         string
		postedPrice PostedPrice
		expPass     bool
	}{
		{
			"valid posted price",
			PostedPrice{
				MarketId:      "market",
				OracleAddress: types.StringAccAddress(addr),
				Price:         sdk.OneDec(),
				Expiry:        now,
			},
			true,
		},
		{
			"invalid id",
			PostedPrice{
				MarketId: " ",
			},
			false,
		},
		{
			"invalid oracle",
			PostedPrice{
				MarketId:      "market",
				OracleAddress: nil,
			},
			false,
		},
		{
			"invalid price",
			PostedPrice{
				MarketId:      "market",
				OracleAddress: types.StringAccAddress(addr),
				Price:         sdk.NewDec(-1),
			},
			false,
		},
		{
			"zero expiry time ",
			PostedPrice{
				MarketId:      "market",
				OracleAddress: types.StringAccAddress(addr),
				Price:         sdk.OneDec(),
				Expiry:        time.Time{},
			},
			false,
		},
	}

	for _, tc := range testCases {
		err := tc.postedPrice.Validate()
		if tc.expPass {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}
