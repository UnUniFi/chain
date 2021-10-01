package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/app"
	"github.com/lcnem/jpyx/types"
	pricefeedtypes "github.com/lcnem/jpyx/x/pricefeed/types"
)

func NewPricefeedGenStateMulti(tApp app.TestApp) app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: types.StringAccAddresses([]sdk.AccAddress{}), Active: true},
				{MarketId: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: types.StringAccAddresses([]sdk.AccAddress{}), Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      "btc:usd",
				OracleAddress: types.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "xrp:usd",
				OracleAddress: types.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&pfGenesis)}
}
