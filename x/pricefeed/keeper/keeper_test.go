package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/app"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// TestKeeper_SetGetMarket tests adding markets to the pricefeed, getting markets from the store
func TestKeeper_SetGetMarket(t *testing.T) {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{})
	keeper := tApp.GetPriceFeedKeeper()

	mp := pricefeedtypes.Params{
		Markets: pricefeedtypes.Markets{
			pricefeedtypes.Market{MarketId: "tstusd", BaseAsset: "tst", QuoteAsset: "usd", Oracles: []string{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	markets := keeper.GetMarkets(ctx)
	require.Equal(t, len(markets), 1)
	require.Equal(t, markets[0].MarketId, "tstusd")

	_, found := keeper.GetMarket(ctx, "tstusd")
	require.Equal(t, found, true)

	mp = pricefeedtypes.Params{
		Markets: pricefeedtypes.Markets{
			pricefeedtypes.Market{MarketId: "tstusd", BaseAsset: "tst", QuoteAsset: "usd", Oracles: []string{}, Active: true},
			pricefeedtypes.Market{MarketId: "tst2usd", BaseAsset: "tst2", QuoteAsset: "usd", Oracles: []string{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	markets = keeper.GetMarkets(ctx)
	require.Equal(t, len(markets), 2)
	require.Equal(t, markets[0].MarketId, "tstusd")
	require.Equal(t, markets[1].MarketId, "tst2usd")

	_, found = keeper.GetMarket(ctx, "nan")
	require.Equal(t, found, false)
}

// TestKeeper_GetSetPrice Test Posting the price by an oracle
func TestKeeper_GetSetPrice(t *testing.T) {
	_, addrs := app.GeneratePrivKeyAddressPairs(2)
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{})
	keeper := tApp.GetPriceFeedKeeper()

	mp := pricefeedtypes.Params{
		Markets: pricefeedtypes.Markets{
			pricefeedtypes.Market{MarketId: "tstusd", BaseAsset: "tst", QuoteAsset: "usd", Oracles: []string{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	// Set price by oracle 1
	_, err := keeper.SetPrice(
		ctx, addrs[0], "tstusd",
		sdk.MustNewDecFromStr("0.33"),
		time.Now().Add(1*time.Hour))
	require.NoError(t, err)
	// Get raw prices
	rawPrices, err := keeper.GetRawPrices(ctx, "tstusd")
	require.NoError(t, err)
	require.Equal(t, len(rawPrices), 1)
	require.Equal(t, rawPrices[0].Price.Equal(sdk.MustNewDecFromStr("0.33")), true)
	// Set price by oracle 2
	_, err = keeper.SetPrice(
		ctx, addrs[1], "tstusd",
		sdk.MustNewDecFromStr("0.35"),
		time.Now().Add(time.Hour*1))
	require.NoError(t, err)

	rawPrices, err = keeper.GetRawPrices(ctx, "tstusd")
	require.NoError(t, err)
	require.Equal(t, len(rawPrices), 2)
	require.Equal(t, rawPrices[1].Price.Equal(sdk.MustNewDecFromStr("0.35")), true)

	// Update Price by Oracle 1
	_, err = keeper.SetPrice(
		ctx, addrs[0], "tstusd",
		sdk.MustNewDecFromStr("0.37"),
		time.Now().Add(time.Hour*1))
	require.NoError(t, err)
	rawPrices, err = keeper.GetRawPrices(ctx, "tstusd")
	require.NoError(t, err)
	require.Equal(t, rawPrices[0].Price.Equal(sdk.MustNewDecFromStr("0.37")), true)
}

// TestKeeper_GetSetCurrentPrice Test Setting the median price of an Asset
func TestKeeper_GetSetCurrentPrice(t *testing.T) {
	_, addrs := app.GeneratePrivKeyAddressPairs(4)
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{})
	keeper := tApp.GetPriceFeedKeeper()

	mp := pricefeedtypes.Params{
		Markets: pricefeedtypes.Markets{
			pricefeedtypes.Market{MarketId: "tstusd", BaseAsset: "tst", QuoteAsset: "usd", Oracles: []string{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	keeper.SetPrice(
		ctx, addrs[0], "tstusd",
		sdk.MustNewDecFromStr("0.33"),
		time.Now().Add(time.Hour*1))
	keeper.SetPrice(
		ctx, addrs[1], "tstusd",
		sdk.MustNewDecFromStr("0.35"),
		time.Now().Add(time.Hour*1))
	keeper.SetPrice(
		ctx, addrs[2], "tstusd",
		sdk.MustNewDecFromStr("0.34"),
		time.Now().Add(time.Hour*1))
	// Set current price
	err := keeper.SetCurrentPrices(ctx, "tstusd")
	require.NoError(t, err)
	// Get Current price
	price, err := keeper.GetCurrentPrice(ctx, "tstusd")
	require.Nil(t, err)
	require.Equal(t, price.Price.Equal(sdk.MustNewDecFromStr("0.34")), true)

	// Even number of oracles
	keeper.SetPrice(
		ctx, addrs[3], "tstusd",
		sdk.MustNewDecFromStr("0.36"),
		time.Now().Add(time.Hour*1))
	err = keeper.SetCurrentPrices(ctx, "tstusd")
	require.NoError(t, err)
	price, err = keeper.GetCurrentPrice(ctx, "tstusd")
	require.Nil(t, err)
	require.Equal(t, price.Price.Equal(sdk.MustNewDecFromStr("0.345")), true)
}
