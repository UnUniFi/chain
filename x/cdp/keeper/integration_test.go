package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/UnUniFi/chain/app"
	ununifitypes "github.com/UnUniFi/chain/types"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

// Avoid cluttering test cases with long function names
func i(in int64) sdk.Int                    { return sdk.NewInt(in) }
func d(str string) sdk.Dec                  { return sdk.MustNewDecFromStr(str) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }

func NewPricefeedGenState(tApp app.TestApp, asset string, price sdk.Dec) app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: asset + ":jpy", BaseAsset: asset, QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: asset + ":eur", BaseAsset: asset, QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      asset + ":jpy",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         price,
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      asset + ":eur",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         price,
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&pfGenesis)}
}

func NewCDPGenState(tApp app.TestApp, asset string, liquidationRatio sdk.Dec) app.GenesisState {
	cdpGenesis := cdptypes.GenesisState{
		Params: cdptypes.Params{
			CollateralParams: cdptypes.CollateralParams{
				{
					Denom:                            asset,
					Type:                             asset + "-a",
					LiquidationRatio:                 liquidationRatio,
					DebtLimit:                        sdk.NewInt64Coin("jpu", 1000000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(100),
					Prefix:                           0x20,
					SpotMarketId:                     asset + ":jpy",
					LiquidationMarketId:              asset + ":jpy",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(6),
				},
			},
			DebtParams: cdptypes.DebtParams{
				{
					Denom:                   "jpu",
					ReferenceAsset:          "jpy",
					ConversionFactor:        i(6),
					DebtFloor:               i(10000000),
					GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 1000000000000),
					DebtDenom:               "deptjpu",
					SurplusAuctionThreshold: sdk.NewInt(500000000000),
					SurplusAuctionLot:       sdk.NewInt(10000000000),
					DebtAuctionThreshold:    sdk.NewInt(100000000000),
					DebtAuctionLot:          sdk.NewInt(10000000000),
					CircuitBreaker:          false,
				},
			},
		},
		StartingCdpId: cdptypes.DefaultCdpStartingID,
		GovDenom:      cdptypes.DefaultGovDenom,
		Cdps:          cdptypes.Cdps{},
		PreviousAccumulationTimes: cdptypes.GenesisAccumulationTimes{
			cdptypes.NewGenesisAccumulationTime(asset+"-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: cdptypes.GenesisTotalPrincipals{
			cdptypes.NewGenesisTotalPrincipal(asset+"-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{cdptypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&cdpGenesis)}
}

func NewPricefeedGenStateMulti(tApp app.TestApp) app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: "btc:jpy", BaseAsset: "btc", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "xrp:jpy", BaseAsset: "xrp", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "bnb:jpy", BaseAsset: "bnb", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "bjpy:jpy", BaseAsset: "bjpy", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},

				{MarketId: "btc:eur", BaseAsset: "btc", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "xrp:eur", BaseAsset: "xrp", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "bnb:eur", BaseAsset: "bnb", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "bjpy:eur", BaseAsset: "bjpy", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      "btc:jpy",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "xrp:jpy",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "bnb:jpy",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("17.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "bjpy:jpy",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},

			{
				MarketId:      "btc:eur",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "xrp:eur",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "bnb:eur",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("17.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "bjpy:eur",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&pfGenesis)}
}
func NewCDPGenStateMulti(tApp app.TestApp) app.GenesisState {
	cdpGenesis := cdptypes.GenesisState{
		Params: cdptypes.Params{
			CollateralParams: cdptypes.CollateralParams{
				{
					Denom:                            "xrp",
					Type:                             "xrp-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("2.0"),
					DebtLimit:                        sdk.NewInt64Coin("jpu", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(7000000000),
					Prefix:                           0x20,
					SpotMarketId:                     "xrp:jpy",
					LiquidationMarketId:              "xrp:jpy",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(6),
				},
				{
					Denom:                            "btc",
					Type:                             "btc-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
					DebtLimit:                        sdk.NewInt64Coin("jpu", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000000782997609"), // %2.5 apr
					LiquidationPenalty:               d("0.025"),
					AuctionSize:                      i(10000000),
					Prefix:                           0x21,
					SpotMarketId:                     "btc:jpy",
					LiquidationMarketId:              "btc:jpy",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
				{
					Denom:                            "bnb",
					Type:                             "bnb-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
					DebtLimit:                        sdk.NewInt64Coin("jpu", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(50000000000),
					Prefix:                           0x22,
					SpotMarketId:                     "bnb:jpy",
					LiquidationMarketId:              "bnb:jpy",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
				{
					Denom:                            "bjpy",
					Type:                             "bjpy-a",
					LiquidationRatio:                 d("1.01"),
					DebtLimit:                        sdk.NewInt64Coin("jpu", 500000000000),
					StabilityFee:                     sdk.OneDec(), // %0 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(10000000000),
					Prefix:                           0x23,
					SpotMarketId:                     "bjpy:jpy",
					LiquidationMarketId:              "bjpy:jpy",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
				{
					Denom:                            "xrp",
					Type:                             "xrp-b",
					LiquidationRatio:                 sdk.MustNewDecFromStr("2.0"),
					DebtLimit:                        sdk.NewInt64Coin("euu", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(7000000000),
					Prefix:                           0x24,
					SpotMarketId:                     "xrp:eur",
					LiquidationMarketId:              "xrp:eur",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(6),
				},
				{
					Denom:                            "btc",
					Type:                             "btc-b",
					LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
					DebtLimit:                        sdk.NewInt64Coin("euu", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000000782997609"), // %2.5 apr
					LiquidationPenalty:               d("0.025"),
					AuctionSize:                      i(10000000),
					Prefix:                           0x25,
					SpotMarketId:                     "btc:eur",
					LiquidationMarketId:              "btc:eur",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
				{
					Denom:                            "bnb",
					Type:                             "bnb-b",
					LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
					DebtLimit:                        sdk.NewInt64Coin("euu", 500000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(50000000000),
					Prefix:                           0x26,
					SpotMarketId:                     "bnb:eur",
					LiquidationMarketId:              "bnb:eur",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
				{
					Denom:                            "bjpy",
					Type:                             "bjpy-b",
					LiquidationRatio:                 d("1.01"),
					DebtLimit:                        sdk.NewInt64Coin("euu", 500000000000),
					StabilityFee:                     sdk.OneDec(), // %0 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(10000000000),
					Prefix:                           0x27,
					SpotMarketId:                     "bjpy:eur",
					LiquidationMarketId:              "bjpy:eur",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
			},
			DebtParams: cdptypes.DebtParams{
				{
					Denom:                   "jpu",
					ReferenceAsset:          "jpy",
					ConversionFactor:        i(6),
					DebtFloor:               i(10000000),
					GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 2000000000000),
					DebtDenom:               "debtjpu",
					SurplusAuctionThreshold: sdk.NewInt(500000000000),
					SurplusAuctionLot:       sdk.NewInt(10000000000),
					DebtAuctionThreshold:    sdk.NewInt(100000000000),
					DebtAuctionLot:          sdk.NewInt(10000000000),
					CircuitBreaker:          false,
				},
				{
					Denom:                   "euu",
					ReferenceAsset:          "eur",
					ConversionFactor:        i(6),
					DebtFloor:               i(10000000),
					GlobalDebtLimit:         sdk.NewInt64Coin("euu", 2000000000000),
					DebtDenom:               "debteuu",
					SurplusAuctionThreshold: sdk.NewInt(500000000000),
					SurplusAuctionLot:       sdk.NewInt(10000000000),
					DebtAuctionThreshold:    sdk.NewInt(100000000000),
					DebtAuctionLot:          sdk.NewInt(10000000000),
					CircuitBreaker:          false,
				},
			},
		},
		StartingCdpId: cdptypes.DefaultCdpStartingID,
		GovDenom:      cdptypes.DefaultGovDenom,
		Cdps:          cdptypes.Cdps{},
		PreviousAccumulationTimes: cdptypes.GenesisAccumulationTimes{
			cdptypes.NewGenesisAccumulationTime("btc-a", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("xrp-a", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("bjpy-a", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("bnb-a", time.Time{}, sdk.OneDec()),

			cdptypes.NewGenesisAccumulationTime("btc-b", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("xrp-b", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("bjpy-b", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("bnb-b", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: cdptypes.GenesisTotalPrincipals{
			cdptypes.NewGenesisTotalPrincipal("btc-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("xrp-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("bjpy-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("bnb-a", sdk.ZeroInt()),

			cdptypes.NewGenesisTotalPrincipal("btc-b", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("xrp-b", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("bjpy-b", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("bnb-b", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{cdptypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&cdpGenesis)}
}

func NewCDPGenStateHighDebtLimit(tApp app.TestApp) app.GenesisState {
	cdpGenesis := cdptypes.GenesisState{
		Params: cdptypes.Params{
			CollateralParams: cdptypes.CollateralParams{
				{
					Denom:                            "xrp",
					Type:                             "xrp-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("2.0"),
					DebtLimit:                        sdk.NewInt64Coin("jpu", 50000000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:               d("0.05"),
					AuctionSize:                      i(7000000000),
					Prefix:                           0x20,
					SpotMarketId:                     "xrp:jpy",
					LiquidationMarketId:              "xrp:jpy",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(6),
				},
				{
					Denom:                            "btc",
					Type:                             "btc-a",
					LiquidationRatio:                 sdk.MustNewDecFromStr("1.5"),
					DebtLimit:                        sdk.NewInt64Coin("jpu", 50000000000000),
					StabilityFee:                     sdk.MustNewDecFromStr("1.000000000782997609"), // %2.5 apr
					LiquidationPenalty:               d("0.025"),
					AuctionSize:                      i(10000000),
					Prefix:                           0x21,
					SpotMarketId:                     "btc:jpy",
					LiquidationMarketId:              "btc:jpy",
					KeeperRewardPercentage:           d("0.01"),
					CheckCollateralizationIndexCount: i(10),
					ConversionFactor:                 i(8),
				},
			},
			DebtParams: cdptypes.DebtParams{
				{
					Denom:                   "jpu",
					ReferenceAsset:          "jpy",
					ConversionFactor:        i(6),
					DebtFloor:               i(10000000),
					GlobalDebtLimit:         sdk.NewInt64Coin("jpu", 100000000000000),
					DebtDenom:               "debtjpu",
					SurplusAuctionThreshold: sdk.NewInt(500000000000),
					SurplusAuctionLot:       sdk.NewInt(10000000000),
					DebtAuctionThreshold:    sdk.NewInt(100000000000),
					DebtAuctionLot:          sdk.NewInt(10000000000),
					CircuitBreaker:          false,
				},
			},
		},
		StartingCdpId: cdptypes.DefaultCdpStartingID,
		GovDenom:      cdptypes.DefaultGovDenom,
		Cdps:          cdptypes.Cdps{},
		PreviousAccumulationTimes: cdptypes.GenesisAccumulationTimes{
			cdptypes.NewGenesisAccumulationTime("btc-a", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("xrp-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: cdptypes.GenesisTotalPrincipals{
			cdptypes.NewGenesisTotalPrincipal("btc-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("xrp-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{cdptypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&cdpGenesis)}
}

func cdps() (cdps cdptypes.Cdps) {
	_, addrs := app.GeneratePrivKeyAddressPairs(6)
	c1 := cdptypes.NewCdp(uint64(1), addrs[0], sdk.NewCoin("xrp", sdk.NewInt(10000000)), "xrp-a", sdk.NewCoin("jpu", sdk.NewInt(8000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c2 := cdptypes.NewCdp(uint64(2), addrs[1], sdk.NewCoin("xrp", sdk.NewInt(100000000)), "xrp-a", sdk.NewCoin("jpu", sdk.NewInt(10000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c3 := cdptypes.NewCdp(uint64(3), addrs[1], sdk.NewCoin("btc", sdk.NewInt(1000000000)), "btc-a", sdk.NewCoin("jpu", sdk.NewInt(10000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c4 := cdptypes.NewCdp(uint64(4), addrs[2], sdk.NewCoin("xrp", sdk.NewInt(1000000000)), "xrp-a", sdk.NewCoin("jpu", sdk.NewInt(500000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c5 := cdptypes.NewCdp(uint64(1), addrs[3], sdk.NewCoin("xrp", sdk.NewInt(10000000)), "xrp-b", sdk.NewCoin("euu", sdk.NewInt(8000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c6 := cdptypes.NewCdp(uint64(2), addrs[4], sdk.NewCoin("xrp", sdk.NewInt(100000000)), "xrp-b", sdk.NewCoin("euu", sdk.NewInt(10000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c7 := cdptypes.NewCdp(uint64(3), addrs[4], sdk.NewCoin("btc", sdk.NewInt(1000000000)), "btc-b", sdk.NewCoin("euu", sdk.NewInt(10000000)), tmtime.Canonical(time.Now()), sdk.OneDec())
	c8 := cdptypes.NewCdp(uint64(4), addrs[5], sdk.NewCoin("xrp", sdk.NewInt(1000000000)), "xrp-b", sdk.NewCoin("euu", sdk.NewInt(500000000)), tmtime.Canonical(time.Now()), sdk.OneDec())

	cdps = append(cdps, c1, c2, c3, c4, c5, c6, c7, c8)
	return
}
