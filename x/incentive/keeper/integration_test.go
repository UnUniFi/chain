package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/lcnem/jpyx/app"
	jpyxtypes "github.com/lcnem/jpyx/types"
	cdptypes "github.com/lcnem/jpyx/x/cdp/types"
	pricefeedtypes "github.com/lcnem/jpyx/x/pricefeed/types"
)

func NewCDPGenStateMulti() app.GenesisState {
	cdpGenesis := cdptypes.GenesisState{
		Params: cdptypes.Params{
			GlobalDebtLimit:         sdk.NewInt64Coin("usdx", 2000000000000),
			SurplusAuctionThreshold: cdptypes.DefaultSurplusThreshold,
			SurplusAuctionLot:       cdptypes.DefaultSurplusLot,
			DebtAuctionThreshold:    cdptypes.DefaultDebtThreshold,
			DebtAuctionLot:          cdptypes.DefaultDebtLot,
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
				{
					Denom:               "bnb",
					Type:                "bnb-a",
					LiquidationRatio:    sdk.MustNewDecFromStr("1.5"),
					DebtLimit:           sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:        sdk.MustNewDecFromStr("1.000000001547125958"), // %5 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(50000000000),
					Prefix:              0x22,
					SpotMarketId:        "bnb:usd",
					LiquidationMarketId: "bnb:usd",
					ConversionFactor:    i(8),
				},
				{
					Denom:               "busd",
					Type:                "busd-a",
					LiquidationRatio:    d("1.01"),
					DebtLimit:           sdk.NewInt64Coin("usdx", 500000000000),
					StabilityFee:        sdk.OneDec(), // %0 apr
					LiquidationPenalty:  d("0.05"),
					AuctionSize:         i(10000000000),
					Prefix:              0x23,
					SpotMarketId:        "busd:usd",
					LiquidationMarketId: "busd:usd",
					ConversionFactor:    i(8),
				},
			},
			DebtParam: cdptypes.DebtParam{
				Denom:            "usdx",
				ReferenceAsset:   "usd",
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
			cdptypes.NewGenesisAccumulationTime("busd-a", time.Time{}, sdk.OneDec()),
			cdptypes.NewGenesisAccumulationTime("bnb-a", time.Time{}, sdk.OneDec()),
		},
		TotalPrincipals: cdptypes.GenesisTotalPrincipals{
			cdptypes.NewGenesisTotalPrincipal("btc-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("xrp-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("busd-a", sdk.ZeroInt()),
			cdptypes.NewGenesisTotalPrincipal("bnb-a", sdk.ZeroInt()),
		},
	}
	return app.GenesisState{cdptypes.ModuleName: cdptypes.ModuleCdc.MustMarshalJSON(&cdpGenesis)}
}

func NewPricefeedGenStateMulti() app.GenesisState {
	pfGenesis := pricefeedtypes.GenesisState{
		Params: pricefeedtypes.Params{
			Markets: []pricefeedtypes.Market{
				{MarketId: "kava:usd", BaseAsset: "kava", QuoteAsset: "usd", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
				{MarketId: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
				{MarketId: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
				{MarketId: "bnb:usd", BaseAsset: "bnb", QuoteAsset: "usd", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
				{MarketId: "busd:usd", BaseAsset: "busd", QuoteAsset: "usd", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
				{MarketId: "zzz:usd", BaseAsset: "zzz", QuoteAsset: "usd", Oracles: []jpyxtypes.StringAccAddress{}, Active: true},
			},
		},
		PostedPrices: []pricefeedtypes.PostedPrice{
			{
				MarketId:      "kava:usd",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "btc:usd",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "xrp:usd",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "bnb:usd",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("17.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "busd:usd",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.OneDec(),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketId:      "zzz:usd",
				OracleAddress: jpyxtypes.StringAccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{pricefeedtypes.ModuleName: pricefeedtypes.ModuleCdc.MustMarshalJSON(&pfGenesis)}
}

/* Question: Not necessary for jpyx because hard module does not exist?
func NewHardGenStateMulti() app.GenesisState {
	loanToValue, _ := sdk.NewDecFromStr("0.6")
	borrowLimit := sdk.NewDec(1000000000000000)

	hardGS := hard.NewGenesisState(hard.NewParams(
		hard.MoneyMarkets{
			hard.NewMoneyMarket("usdx", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "usdx:usd", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("ukava", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "kava:usd", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("bnb", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "bnb:usd", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("btcb", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "btc:usd", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("xrp", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "xrp:usd", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
			hard.NewMoneyMarket("zzz", hard.NewBorrowLimit(false, borrowLimit, loanToValue), "zzz:usd", sdk.NewInt(1000000), hard.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
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
			sdk.NewCoin("ukava", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("btcb", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("xrp", sdk.NewInt(1000000000000000)),
			sdk.NewCoin("zzz", sdk.NewInt(1000000000000000)),
		)
	}

	return app.NewAuthGenState(tApp, addresses, coinsList)
}

func NewStakingGenesisState() app.GenesisState {
	genState := stakingtypes.DefaultGenesisState()
	genState.Params.BondDenom = "ukava"
	return app.GenesisState{
		stakingtypes.ModuleName: stakingtypes.ModuleCdc.MustMarshalJSON(genState),
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
		NewAuthGenState(tApp, allAddrs, cs(c("ukava", 5_000_000))),
		NewStakingGenesisState(),
		NewPricefeedGenStateMulti(),
		NewCDPGenStateMulti(),
		// NewHardGenStateMulti(),
	)

	/* Question: Not necessary for jpyx because committee module does not exist?
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
