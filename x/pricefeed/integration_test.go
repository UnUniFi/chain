package pricefeed_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/app"
	"github.com/lcnem/jpyx/x/pricefeed"
)

func NewPricefeedGenStateMulti() app.GenesisState {
	pfGenesis := pricefeed.GenesisState{
		Params: pricefeed.Params{
			Markets: []pricefeed.Market{
				{MarketID: "btc:jpy", BaseAsset: "btc", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "xrp:jpy", BaseAsset: "xrp", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []pricefeed.PostedPrice{
			{
				MarketID:      "btc:jpy",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "xrp:jpy",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeed.ModuleName: pricefeed.ModuleCdc.MustMarshalJSON(pfGenesis)}
}

func NewPricefeedGenStateWithOracles(addrs []sdk.AccAddress) app.GenesisState {
	pfGenesis := pricefeed.GenesisState{
		Params: pricefeed.Params{
			Markets: []pricefeed.Market{
				{MarketID: "btc:jpy", BaseAsset: "btc", QuoteAsset: "jpy", Oracles: addrs, Active: true},
				{MarketID: "xrp:jpy", BaseAsset: "xrp", QuoteAsset: "jpy", Oracles: addrs, Active: true},
			},
		},
		PostedPrices: []pricefeed.PostedPrice{
			{
				MarketID:      "btc:jpy",
				OracleAddress: addrs[0],
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "xrp:jpy",
				OracleAddress: addrs[0],
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeed.ModuleName: pricefeed.ModuleCdc.MustMarshalJSON(pfGenesis)}
}
