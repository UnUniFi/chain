package cdp_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/lcnem/jpyx/app"
	jpyxtypes "github.com/lcnem/jpyx/types"
	cdptypes "github.com/lcnem/jpyx/x/cdp/types"
	pricefeedtypes "github.com/lcnem/jpyx/x/pricefeed/types"
)

// Avoid cluttering test cases with long function names
func i(in int64) sdk.Int                    { return sdk.NewInt(in) }
func d(str string) sdk.Dec                  { return sdk.MustNewDecFromStr(str) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }

func NewPricefeedGenState(asset string, price sdk.Dec) app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: asset + ":usd", BaseAsset: asset, QuoteAsset: "usd", Oracles: jpyxtypes.StringAccAddresses([]sdk.AccAddress{}), Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      asset + ":usd",
				OracleAddress: jpyxtypes.StringAccAddress(sdk.AccAddress{}),
				Price:         price,
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: pricefeedtypes.ModuleCdc.MustMarshalJSON(&pfGenesis)}
}

func NewCDPGenState(asset string, liquidationRatio sdk.Dec) app.GenesisState {
	cdpGenesis := cdptypes.GenesisState{
		Params: cdptypes.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("usdx", 1000000000000),
			SurplusAuctionThreshold: cdptypes.DefaultSurplusThreshold,
			SurplusAuctionLot:       cdptypes.DefaultSurplusLot,
			DebtAuctionThreshold:    cdptypes.DefaultDebtThreshold,
			DebtAuctionLot:          cdptypes.DefaultDebtLot,
			// SavingsDistributionFrequency: cdptypes.DefaultSavingsDistributionFrequency,
			CollateralParams: cdptypes.CollateralParams{
				{
					Denom:               asset,
					Type:                asset + "-a",
					LiquidationRatio:    liquidationRatio,
					DebtLimit:           sdk.NewInt64Coin("usdx", 1000000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(1000000000),
					Prefix:              0x20,
					ConversionFactor:    i(6),
					SpotMarketId:        asset + ":usd",
					LiquidationMarketId: asset + ":usd",
				},
			},
			DebtParam: cdptypes.DebtParam{
				Denom:            "usdx",
				ReferenceAsset:   "usd",
				ConversionFactor: i(6),
				DebtFloor:        i(10000000),
				// SavingsRate:      d("0.95"),
			},
		},
		StartingCdpId: cdptypes.DefaultCdpStartingID,
		DebtDenom:     cdptypes.DefaultDebtDenom,
		GovDenom:      cdptypes.DefaultGovDenom,
		Cdps:          cdptypes.Cdps{},
		// PreviousDistributionTime: cdptypes.DefaultPreviousDistributionTime,
	}
	return app.GenesisState{cdptypes.ModuleName: cdptypes.ModuleCdc.MustMarshalJSON(&cdpGenesis)}
}

func NewPricefeedGenStateMulti() app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: jpyxtypes.StringAccAddresses([]sdk.AccAddress{}), Active: true},
				{MarketId: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: jpyxtypes.StringAccAddresses([]sdk.AccAddress{}), Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      "btc:usd",
				OracleAddress: jpyxtypes.StringAccAddress(sdk.AccAddress{}),
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "xrp:usd",
				OracleAddress: jpyxtypes.StringAccAddress(sdk.AccAddress{}),
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: pricefeedtypes.ModuleCdc.MustMarshalJSON(&pfGenesis)}
}
func NewCDPGenStateMulti() app.GenesisState {
	cdpGenesis := cdptypes.GenesisState{
		Params: cdptypes.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("usdx", 1000000000000),
			SurplusAuctionThreshold: cdptypes.DefaultSurplusThreshold,
			SurplusAuctionLot:       cdptypes.DefaultSurplusLot,
			DebtAuctionThreshold:    cdptypes.DefaultDebtThreshold,
			DebtAuctionLot:          cdptypes.DefaultDebtLot,
			// SavingsDistributionFrequency: cdptypes.DefaultSavingsDistributionFrequency,
			CollateralParams: cdptypes.CollateralParams{
				{
					Denom:               "xrp",
					Type:                "xrp-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("2.0"),
					DebtLimit:           sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(7000000000),
					Prefix:              0x20,
					SpotMarketId:        "xrp:usd",
					LiquidationMarketId: "xrp:usd",
					ConversionFactor:    i(6),
				},
				{
					Denom:               "btc",
					Type:                "btc-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("1.5"),
					DebtLimit:           sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000000782997609"), // %2.5 apr
					LiquidationPenalty:  d("0.025"),
					AuctionSize:         i(10000000),
					Prefix:              0x21,
					SpotMarketId:        "btc:usd",
					LiquidationMarketId: "btc:usd",
					ConversionFactor:    i(8),
				},
			},
			DebtParam: cdptypes.DebtParam{
				Denom:            "usdx",
				ReferenceAsset:   "usd",
				ConversionFactor: i(6),
				DebtFloor:        i(10000000),
				// SavingsRate:      d("0.95"),
			},
		},
		StartingCdpId: cdptypes.DefaultCdpStartingID,
		DebtDenom:     cdptypes.DefaultDebtDenom,
		GovDenom:      cdptypes.DefaultGovDenom,
		Cdps:          cdptypes.Cdps{},
		// PreviousDistributionTime: cdp.DefaultPreviousDistributionTime,
	}
	return app.GenesisState{cdptypes.ModuleName: cdptypes.ModuleCdc.MustMarshalJSON(&cdpGenesis)}
}

func cdps() (cdps cdptypes.Cdps) {
	_, addrs := app.GeneratePrivKeyAddressPairs(3)
	c1 := cdptypes.NewCdp(uint64(1), addrs[0], sdk.NewCoin("xrp", sdk.NewInt(100000000)), "xrp-a", sdk.NewCoin("usdx", sdk.NewInt(8000000)), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	c2 := cdptypes.NewCdp(uint64(2), addrs[1], sdk.NewCoin("xrp", sdk.NewInt(100000000)), "xrp-a", sdk.NewCoin("usdx", sdk.NewInt(10000000)), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	c3 := cdptypes.NewCdp(uint64(3), addrs[1], sdk.NewCoin("btc", sdk.NewInt(1000000000)), "btc-a", sdk.NewCoin("usdx", sdk.NewInt(10000000)), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	c4 := cdptypes.NewCdp(uint64(4), addrs[2], sdk.NewCoin("xrp", sdk.NewInt(1000000000)), "xrp-a", sdk.NewCoin("usdx", sdk.NewInt(50000000)), tmtime.Canonical(time.Now()), sdk.NewDec(0))
	cdps = append(cdps, c1, c2, c3, c4)
	return
}
