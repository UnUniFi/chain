package incentive_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/app"
	jpyxtypes "github.com/lcnem/jpyx/types"
	cdptypes "github.com/lcnem/jpyx/x/cdp/types"
	incentivetypes "github.com/lcnem/jpyx/x/incentive/types"
	pricefeedtypes "github.com/lcnem/jpyx/x/pricefeed/types"
)

// Avoid cluttering test cases with long function names
func i(in int64) sdk.Int   { return sdk.NewInt(in) }
func d(str string) sdk.Dec { return sdk.MustNewDecFromStr(str) }

func NewCDPGenStateMulti() app.GenesisState {
	cdpGenesis := cdptypes.GenesisState{
		Params: cdptypes.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("jpyx", 2000000000000),
			SurplusAuctionThreshold: cdptypes.DefaultSurplusThreshold,
			SurplusAuctionLot:       cdptypes.DefaultSurplusLot,
			DebtAuctionThreshold:    cdptypes.DefaultDebtThreshold,
			DebtAuctionLot:          cdptypes.DefaultDebtLot,
			CollateralParams: cdptypes.CollateralParams{
				{
					Denom:               "xrp",
					Type:                "xrp-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("2.0"),
					DebtLimit:           sdk.NewInt64Coin("jpyx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(7000000000),
					Prefix:              0x20,
					SpotMarketId:        "xrp:jpy",
					LiquidationMarketId: "xrp:jpy",
					ConversionFactor:    i(6),
				},
				{
					Denom:               "btc",
					Type:                "btc-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("1.5"),
					DebtLimit:           sdk.NewInt64Coin("jpyx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000000782997609"), // %2.5 apr
					LiquidationPenalty:  d("0.025"),
					AuctionSize:         i(10000000),
					Prefix:              0x21,
					SpotMarketId:        "btc:jpy",
					LiquidationMarketId: "btc:jpy",
					ConversionFactor:    i(8),
				},
				{
					Denom:               "bnb",
					Type:                "bnb-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("1.5"),
					DebtLimit:           sdk.NewInt64Coin("jpyx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(50000000000),
					Prefix:              0x22,
					SpotMarketId:        "bnb:jpy",
					LiquidationMarketId: "bnb:jpy",
					ConversionFactor:    i(8),
				},
				{
					Denom:               "bjpy",
					Type:                "bjpy-a",
					LiquidationRatio:    d("1.01"),
					DebtLimit:           sdk.NewInt64Coin("jpyx", 500000000000),
					StabilityFee:        sdk.OneDec(), // %0 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(10000000000),
					Prefix:              0x23,
					SpotMarketId:        "bjpy:jpy",
					LiquidationMarketId: "bjpy:jpy",
					ConversionFactor:    i(8),
				},
			},
			DebtParam: cdptypes.DebtParam{
				Denom:            "jpyx",
				ReferenceAsset:   "jpy",
				ConversionFactor: i(6),
				DebtFloor:        i(10000000),
			},
		},
		StartingCdpId: cdptypes.DefaultCdpStartingID,
		DebtDenom:     cdptypes.DefaultDebtDenom,
		GovDenom:      cdptypes.DefaultGovDenom,
		Cdps:          cdptypes.Cdps{},
		PreviousAccumulationTimes: cdptypes.GenesisAccumulationTimes{
			cdptypes.NewGenesisAccumulationTime("btc-a", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("xrp-a", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("bjpy-a", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("bnb-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: cdptypes.GenesisTotalPrincipals{
			cdptypes.NewGenesisTotalPrincipal("btc-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("xrp-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("bjpy-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("bnb-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{cdptypes.ModuleName: cdptypes.ModuleCdc.MustMarshalJSON(&cdpGenesis)}
}

func NewPricefeedGenStateMulti() app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: "btc:jpy", BaseAsset: "btc", QuoteAsset: "jpy", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
				{MarketId: "xrp:jpy", BaseAsset: "xrp", QuoteAsset: "jpy", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
				{MarketId: "bnb:jpy", BaseAsset: "bnb", QuoteAsset: "jpy", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
				{MarketId: "bjpy:jpy", BaseAsset: "bjpy", QuoteAsset: "jpy", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      "btc:jpy",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "xrp:jpy",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "bnb:jpy",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("17.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "bjpy:jpy",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: pricefeedtypes.ModuleCdc.MustMarshalJSON(&pfGenesis)}
}

func NewIncentiveGenState(previousAccumTime, endTime time.Time, rewardPeriods ...incentivetypes.RewardPeriod) app.GenesisState {
	var accumulationTimes incentivetypes.GenesisAccumulationTimes
	for _, rp := range rewardPeriods {
		accumulationTimes = append(
			accumulationTimes,
			incentivetypes.NewGenesisAccumulationTime(
				rp.CollateralType,
				previousAccumTime,
			),
		)
	}
	genesis := incentivetypes.NewGenesisState(
		incentivetypes.NewParams(
			rewardPeriods,
			incentivetypes.Multipliers{
				incentivetypes.NewMultiplier(incentivetypes.Small, 1, d("0.25")),
				incentivetypes.NewMultiplier(incentivetypes.Large, 12, d("1.0")),
			},
			endTime,
		),
		incentivetypes.DefaultGenesisAccumulationTimes,
		incentivetypes.DefaultCdpClaims,
	)
	return app.GenesisState{incentivetypes.ModuleName: incentivetypes.ModuleCdc.MustMarshalJSON(&genesis)}
}

func NewCDPGenStateHighInterest() app.GenesisState {
	cdpGenesis := cdptypes.GenesisState{
		Params: cdptypes.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("jpyx", 2000000000000),
			SurplusAuctionThreshold: cdptypes.DefaultSurplusThreshold,
			SurplusAuctionLot:       cdptypes.DefaultSurplusLot,
			DebtAuctionThreshold:    cdptypes.DefaultDebtThreshold,
			DebtAuctionLot:          cdptypes.DefaultDebtLot,
			CollateralParams: cdptypes.CollateralParams{
				{
					Denom:               "bnb",
					Type:                "bnb-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("1.5"),
					DebtLimit:           sdk.NewInt64Coin("jpyx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000051034942716"), // 500% APR
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(50000000000),
					Prefix:              0x22,
					SpotMarketId:        "bnb:jpy",
					LiquidationMarketId: "bnb:jpy",
					ConversionFactor:    i(8),
				},
			},
			DebtParam: cdptypes.DebtParam{
				Denom:            "jpyx",
				ReferenceAsset:   "jpy",
				ConversionFactor: i(6),
				DebtFloor:        i(10000000),
			},
		},
		StartingCdpId: cdptypes.DefaultCdpStartingID,
		DebtDenom:     cdptypes.DefaultDebtDenom,
		GovDenom:      cdptypes.DefaultGovDenom,
		Cdps:          cdptypes.Cdps{},
		PreviousAccumulationTimes: cdptypes.GenesisAccumulationTimes{
			cdptypes.NewGenesisAccumulationTime("bnb-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: cdptypes.GenesisTotalPrincipals{
			cdptypes.NewGenesisTotalPrincipal("bnb-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{cdptypes.ModuleName: cdptypes.ModuleCdc.MustMarshalJSON(&cdpGenesis)}
}
