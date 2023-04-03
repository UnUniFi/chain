package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/UnUniFi/chain/app"
	ununifitypes "github.com/UnUniFi/chain/types"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func NewCDPGenStateMulti(tApp app.TestApp) app.GenesisState {
	cdpGenesis := cdptypes.GenesisState{
		Params: cdptypes.Params{
			CollateralParams: cdptypes.CollateralParams{
				{
					Denom:               "xrp",
					Type:                "xrp-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("2.0"),
					DebtLimit:           sdk.NewInt64Coin("jpu", 500000000000),
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
					DebtLimit:           sdk.NewInt64Coin("jpu", 500000000000),
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
					DebtLimit:           sdk.NewInt64Coin("jpu", 500000000000),
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
					DebtLimit:           sdk.NewInt64Coin("jpu", 500000000000),
					StabilityFee:        sdk.OneDec(), // %0 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(10000000000),
					Prefix:              0x23,
					SpotMarketId:        "bjpy:jpy",
					LiquidationMarketId: "bjpy:jpy",
					ConversionFactor:    i(8),
				},
				{
					Denom:               "xrp",
					Type:                "xrp-b",
					LiquidationRatio:    sdk.MustNewDecFromStr("2.0"),
					DebtLimit:           sdk.NewInt64Coin("euu", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(7000000000),
					Prefix:              0x24,
					SpotMarketId:        "xrp:eur",
					LiquidationMarketId: "xrp:eur",
					ConversionFactor:    i(6),
				},
				{
					Denom:               "btc",
					Type:                "btc-b",
					LiquidationRatio:    sdk.MustNewDecFromStr("1.5"),
					DebtLimit:           sdk.NewInt64Coin("euu", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000000782997609"), // %2.5 apr
					LiquidationPenalty:  d("0.025"),
					AuctionSize:         i(10000000),
					Prefix:              0x25,
					SpotMarketId:        "btc:eur",
					LiquidationMarketId: "btc:eur",
					ConversionFactor:    i(8),
				},
				{
					Denom:               "bnb",
					Type:                "bnb-b",
					LiquidationRatio:    sdk.MustNewDecFromStr("1.5"),
					DebtLimit:           sdk.NewInt64Coin("euu", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(50000000000),
					Prefix:              0x26,
					SpotMarketId:        "bnb:eur",
					LiquidationMarketId: "bnb:eur",
					ConversionFactor:    i(8),
				},
				{
					Denom:               "bjpy",
					Type:                "bjpy-b",
					LiquidationRatio:    d("1.01"),
					DebtLimit:           sdk.NewInt64Coin("euu", 500000000000),
					StabilityFee:        sdk.OneDec(), // %0 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(10000000000),
					Prefix:              0x27,
					SpotMarketId:        "bjpy:eur",
					LiquidationMarketId: "bjpy:eur",
					ConversionFactor:    i(8),
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

func NewPricefeedGenStateMulti(tApp app.TestApp) app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: "kava:jpy", BaseAsset: "kava", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "btc:jpy", BaseAsset: "btc", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "xrp:jpy", BaseAsset: "xrp", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "bnb:jpy", BaseAsset: "bnb", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "bjpy:jpy", BaseAsset: "bjpy", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "zzz:jpy", BaseAsset: "zzz", QuoteAsset: "jpy", Oracles: []ununifitypes.StringAccAddress{}, Active: true},

				{MarketId: "kava:eur", BaseAsset: "kava", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "btc:eur", BaseAsset: "btc", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "xrp:eur", BaseAsset: "xrp", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "bnb:eur", BaseAsset: "bnb", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "bjpy:eur", BaseAsset: "bjpy", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
				{MarketId: "zzz:eur", BaseAsset: "zzz", QuoteAsset: "eur", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      "kava:jpy",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
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
				MarketId:      "zzz:jpy",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "kava:eur",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
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
			{
				MarketId:      "zzz:eur",
				OracleAddress: ununifitypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: tApp.AppCodec().MustMarshalJSON(&pfGenesis)}
}

/* Question: Not necessary for jpu because hard module does not exist?
func NewHardGenStateMulti() app.GenesisState {
	loanToValue, _ := sdk.NewDecFromStr("0.6")
	borrowLimit := sdk.NewDec(1000000000000000)

	hardGS := hard.NewGenesisState(hard.NewParams(
		hard.MoneyMarkets{
			hard.NewMoneyMarket("jpu", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "jpu:jpy", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("uguu", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "kava:jpy", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("bnb", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "bnb:jpy", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("btcb", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "btc:jpy", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("xrp", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "xrp:jpy", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("zzz", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "zzz:jpy", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
		},
		sdk.NewDec(10),
	), hard.DefaultAccumulationTimes, hard.DefaultDeposits, hard.DefaultBorrows,
		hard.DefaultTotalSupplied, hard.DefaultTotalBorrowed, hard.DefaultTotalReserves,
	)

	return app.GenesisState{hard.ModuleName: hard.ModuleCdc.MustMarshalJSON(hardGS)}
}
*/

func NewAuthGenState(tApp app.TestApp, addresses []sdk.AccAddress, coins sdk.Coins) app.GenesisState {
	coinsList := []sdk.Coins{}
	for range addresses {
		coinsList = append(coinsList, coins)
	}

	// Load up our primary user address
	if len(addresses) >= 4 {
		coinsList[3] = sdk.NewCoins(
			sdk.NewCoin("bnb", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("uguu", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("btcb", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("xrp", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("zzz", sdk.NewInt(1000000000000000)),
		)
	}

	return app.NewAuthGenState(tApp, addresses, coinsList)
}

func NewStakingGenesisState(tApp app.TestApp) app.GenesisState {
	genState := stakingtypes.DefaultGenesisState()
	genState.Params.BondDenom = "uguu"
	return app.GenesisState{
		stakingtypes.ModuleName: tApp.AppCodec().MustMarshalJSON(genState),
	}
}

func (suite *KeeperTestSuite) SetupWithGenState() {
	_, allAddrs := app.GeneratePrivKeyAddressPairs(10)
	suite.addrs = allAddrs[:5]
	for _, a := range allAddrs[5:] {
		suite.validatorAddrs = append(suite.validatorAddrs, sdk.ValAddress(a))
	}

	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

	tApp.InitializeFromGenesisStates(
		NewAuthGenState(tApp, allAddrs, cs(c("uguu", 5_000_000))),
		NewStakingGenesisState(tApp),
		NewPricefeedGenStateMulti(tApp),
		NewCDPGenStateMulti(tApp),
		// NewHardGenStateMulti(),
	)

	/* Question: Not necessary for jpu because committee module does not exist?
	// Set up a god committee
	committeeModKeeper := tApp.GetCommitteeKeeper()
	godCommittee := committeetypes.Committee{
		ID:               1,
		Description:      "This committee is for testing.",
		Members:          suite.addrs[:2],
		Permissions:      []committeetypes.Permission{committeetypes.GodPermission{}},
		VoteThreshold:    d("0.667"),
		ProposalDuration: time.Hour * 24 * 7,
	}
	committeeModKeeper.SetCommittee(ctx, godCommittee)
	*/

	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = tApp.GetIncentiveKeeper()
	// suite.hardKeeper = tApp.GetHardKeeper()
	suite.stakingKeeper = tApp.GetStakingKeeper()
	// suite.committeeKeeper = committeeModKeeper
}
