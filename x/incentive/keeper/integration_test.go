package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/lcnem/jpyx/app"
	"github.com/lcnem/jpyx/x/cdp"
	committeetypes "github.com/lcnem/jpyx/x/committee/types"
	"github.com/lcnem/jpyx/x/hard"
	"github.com/lcnem/jpyx/x/pricefeed"
)

func NewCDPGenStateMulti() app.GenesisState {
	cdpGenesis := cdp.GenesisState{
		Params: cdp.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("jpyx", 2000000000000),
			SurplusAuctionThreshold: cdp.DefaultSurplusThreshold,
			SurplusAuctionLot:       cdp.DefaultSurplusLot,
			DebtAuctionThreshold:    cdp.DefaultDebtThreshold,
			DebtAuctionLot:          cdp.DefaultDebtLot,
			CollateralParams: cdp.CollateralParams{
				{
					Denom:               "xrp",
					Type:                "xrp-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("2.0"),
					DebtLimit:           sdk.NewInt64Coin("jpyx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(7000000000),
					Prefix:              0x20,
					SpotMarketID:        "xrp:jpy",
					LiquidationMarketID: "xrp:jpy",
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
					SpotMarketID:        "btc:jpy",
					LiquidationMarketID: "btc:jpy",
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
					SpotMarketID:        "bnb:jpy",
					LiquidationMarketID: "bnb:jpy",
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
					SpotMarketID:        "bjpy:jpy",
					LiquidationMarketID: "bjpy:jpy",
					ConversionFactor:    i(8),
				},
			},
			DebtParam: cdp.DebtParam{
				Denom:            "jpyx",
				ReferenceAsset:   "jpy",
				ConversionFactor: i(6),
				DebtFloor:        i(10000000),
			},
		},
		StartingCdpID: cdp.DefaultCdpStartingID,
		DebtDenom:     cdp.DefaultDebtDenom,
		GovDenom:      cdp.DefaultGovDenom,
		CDPs:          cdp.CDPs{},
		PreviousAccumulationTimes: cdp.GenesisAccumulationTimes{
			cdp.NewGenesisAccumulationTime("btc-a", time.Time{}, sdk.OneDec()),
			cdp.NewGenesisAccumulationTime("xrp-a", time.Time{}, sdk.OneDec()),
			cdp.NewGenesisAccumulationTime("bjpy-a", time.Time{}, sdk.OneDec()),
			cdp.NewGenesisAccumulationTime("bnb-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: cdp.GenesisTotalPrincipals{
			cdp.NewGenesisTotalPrincipal("btc-a", sdk.ZeroInt()),
			cdp.NewGenesisTotalPrincipal("xrp-a", sdk.ZeroInt()),
			cdp.NewGenesisTotalPrincipal("bjpy-a", sdk.ZeroInt()),
			cdp.NewGenesisTotalPrincipal("bnb-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{cdp.ModuleName: cdp.ModuleCdc.MustMarshalJSON(cdpGenesis)}
}

func NewPricefeedGenStateMulti() app.GenesisState {
	pfGenesis := pricefeed.GenesisState{
		Params: pricefeed.Params{
			Markets: []pricefeed.Market{
				{MarketID: "kava:jpy", BaseAsset: "kava", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "btc:jpy", BaseAsset: "btc", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "xrp:jpy", BaseAsset: "xrp", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "bnb:jpy", BaseAsset: "bnb", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "bjpy:jpy", BaseAsset: "bjpy", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "zzz:jpy", BaseAsset: "zzz", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []pricefeed.PostedPrice{
			{
				MarketID:      "kava:jpy",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
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
			{
				MarketID:      "bnb:jpy",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("17.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "bjpy:jpy",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "zzz:jpy",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeed.ModuleName: pricefeed.ModuleCdc.MustMarshalJSON(pfGenesis)}
}

func NewHardGenStateMulti() app.GenesisState {
	loanToValue, _ := sdk.NewDecFromStr("0.6")
	borrowLimit := sdk.NewDec(1000000000000000)

	hardGS := hard.NewGenesisState(hard.NewParams(
		hard.MoneyMarkets{
			hard.NewMoneyMarket("jpyx", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "jpyx:jpy", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("ujsmn", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "kava:jpy", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
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

func NewAuthGenState(addresses []sdk.AccAddress, coins sdk.Coins) app.GenesisState {
	coinsList := []sdk.Coins{}
	for range addresses {
		coinsList = append(coinsList, coins)
	}

	// Load up our primary user address
	if len(addresses) >= 4 {
		coinsList[3] = sdk.NewCoins(
			sdk.NewCoin("bnb", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("ujsmn", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("btcb", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("xrp", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("zzz", sdk.NewInt(1000000000000000)),
		)
	}

	return app.NewAuthGenState(addresses, coinsList)
}

func NewStakingGenesisState() app.GenesisState {
	genState := staking.DefaultGenesisState()
	genState.Params.BondDenom = "ujsmn"
	return app.GenesisState{
		staking.ModuleName: staking.ModuleCdc.MustMarshalJSON(genState),
	}
}

func (suite *KeeperTestSuite) SetupWithGenState() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	_, allAddrs := app.GeneratePrivKeyAddressPairs(10)
	suite.addrs = allAddrs[:5]
	for _, a := range allAddrs[5:] {
		suite.validatorAddrs = append(suite.validatorAddrs, sdk.ValAddress(a))
	}

	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{Height: 1, Time: tmtime.Now()})

	tApp.InitializeFromGenesisStates(
		NewAuthGenState(allAddrs, cs(c("ujsmn", 5_000_000))),
		NewStakingGenesisState(),
		NewPricefeedGenStateMulti(),
		NewCDPGenStateMulti(),
		NewHardGenStateMulti(),
	)

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

	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = tApp.GetIncentiveKeeper()
	suite.hardKeeper = tApp.GetHardKeeper()
	suite.stakingKeeper = tApp.GetStakingKeeper()
	suite.committeeKeeper = committeeModKeeper
}
