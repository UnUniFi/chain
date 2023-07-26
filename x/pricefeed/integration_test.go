package pricefeed_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/app"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func NewPricefeedGenStateMulti(tApp app.TestApp) app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []string{""}, Active: true},
				{MarketId: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []string{""}, Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      "btc:usd",
				OracleAddress: "",
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "xrp:usd",
				OracleAddress: "",
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&pfGenesis)}
}

func NewPricefeedGenStateWithOracles(tApp app.TestApp, addrs []string) app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: addrs, Active: true},
				{MarketId: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: addrs, Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      "btc:usd",
				OracleAddress: addrs[0],
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "xrp:usd",
				OracleAddress: addrs[0],
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&pfGenesis)}
}
