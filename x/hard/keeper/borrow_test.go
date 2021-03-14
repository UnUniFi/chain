package keeper_test

import (
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/lcnem/jpyx/app"
	"github.com/lcnem/jpyx/x/hard"
	"github.com/lcnem/jpyx/x/hard/types"
	"github.com/lcnem/jpyx/x/pricefeed"
)

const (
	JPYX_CF = 1000000
	KAVA_CF = 1000000
	BTCB_CF = 100000000
	BNB_CF  = 100000000
	BJPY_CF = 100000000
)

func (suite *KeeperTestSuite) TestBorrow() {

	type args struct {
		jpyxBorrowLimit           sdk.Dec
		priceKAVA                 sdk.Dec
		loanToValueKAVA           sdk.Dec
		priceBTCB                 sdk.Dec
		loanToValueBTCB           sdk.Dec
		priceBNB                  sdk.Dec
		loanToValueBNB            sdk.Dec
		borrower                  sdk.AccAddress
		depositCoins              []sdk.Coin
		previousBorrowCoins       sdk.Coins
		borrowCoins               sdk.Coins
		expectedAccountBalance    sdk.Coins
		expectedModAccountBalance sdk.Coins
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}
	type borrowTest struct {
		name    string
		args    args
		errArgs errArgs
	}
	testCases := []borrowTest{
		{
			"valid",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("100000000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("5.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.6"),
				priceBTCB:                 sdk.MustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.01"),
				priceBNB:                  sdk.MustNewDecFromStr("0.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              []sdk.Coin{sdk.NewCoin("ujsmn", sdk.NewInt(100*KAVA_CF))},
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(20*KAVA_CF))),
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(20*KAVA_CF)), sdk.NewCoin("btcb", sdk.NewInt(100*BTCB_CF)), sdk.NewCoin("bnb", sdk.NewInt(100*BNB_CF)), sdk.NewCoin("xyz", sdk.NewInt(1))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(1080*KAVA_CF)), sdk.NewCoin("jpyx", sdk.NewInt(200*JPYX_CF)), sdk.NewCoin("bjpy", sdk.NewInt(100*BJPY_CF))),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"invalid: loan-to-value limited",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("100000000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("5.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.6"),
				priceBTCB:                 sdk.MustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.01"),
				priceBNB:                  sdk.MustNewDecFromStr("0.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              []sdk.Coin{sdk.NewCoin("ujsmn", sdk.NewInt(20*KAVA_CF))},  // 20 KAVA x $5.00 price = $100
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(61*JPYX_CF))), // 61 JPYX x $1 price = $61
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "exceeds the allowable amount as determined by the collateralization ratio",
			},
		},
		{
			"valid: multiple deposits",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("100000000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("2.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.80"),
				priceBTCB:                 sdk.MustNewDecFromStr("10000.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.10"),
				priceBNB:                  sdk.MustNewDecFromStr("0.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(50*KAVA_CF)), sdk.NewCoin("btcb", sdk.NewInt(0.1*BTCB_CF))),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(180*JPYX_CF))),
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(50*KAVA_CF)), sdk.NewCoin("btcb", sdk.NewInt(99.9*BTCB_CF)), sdk.NewCoin("jpyx", sdk.NewInt(180*JPYX_CF)), sdk.NewCoin("bnb", sdk.NewInt(100*BNB_CF)), sdk.NewCoin("xyz", sdk.NewInt(1))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(1050*KAVA_CF)), sdk.NewCoin("jpyx", sdk.NewInt(20*JPYX_CF)), sdk.NewCoin("btcb", sdk.NewInt(0.1*BTCB_CF)), sdk.NewCoin("bjpy", sdk.NewInt(100*BJPY_CF))),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"invalid: multiple deposits",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("100000000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("2.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.80"),
				priceBTCB:                 sdk.MustNewDecFromStr("10000.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.10"),
				priceBNB:                  sdk.MustNewDecFromStr("0.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(50*KAVA_CF)), sdk.NewCoin("btcb", sdk.NewInt(0.1*BTCB_CF))),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(181*JPYX_CF))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "exceeds the allowable amount as determined by the collateralization ratio",
			},
		},
		{
			"valid: multiple previous borrows",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("100000000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("2.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.8"),
				priceBTCB:                 sdk.MustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.01"),
				priceBNB:                  sdk.MustNewDecFromStr("5.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.8"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("bnb", sdk.NewInt(30*BNB_CF)), sdk.NewCoin("ujsmn", sdk.NewInt(50*KAVA_CF))), // (50 KAVA x $2.00 price = $100) + (30 BNB x $5.00 price = $150) = $250
				previousBorrowCoins:       sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(99*JPYX_CF)), sdk.NewCoin("bjpy", sdk.NewInt(100*BJPY_CF))),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(1*JPYX_CF))),
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(50*KAVA_CF)), sdk.NewCoin("btcb", sdk.NewInt(100*BTCB_CF)), sdk.NewCoin("jpyx", sdk.NewInt(100*JPYX_CF)), sdk.NewCoin("bjpy", sdk.NewInt(100*BJPY_CF)), sdk.NewCoin("bnb", sdk.NewInt(70*BNB_CF)), sdk.NewCoin("xyz", sdk.NewInt(1))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(1050*KAVA_CF)), sdk.NewCoin("bnb", sdk.NewInt(30*BJPY_CF)), sdk.NewCoin("jpyx", sdk.NewInt(100*JPYX_CF))),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"invalid: over loan-to-value with multiple previous borrows",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("100000000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("2.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.8"),
				priceBTCB:                 sdk.MustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.01"),
				priceBNB:                  sdk.MustNewDecFromStr("5.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.8"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("bnb", sdk.NewInt(30*BNB_CF)), sdk.NewCoin("ujsmn", sdk.NewInt(50*KAVA_CF))), // (50 KAVA x $2.00 price = $100) + (30 BNB x $5.00 price = $150) = $250
				previousBorrowCoins:       sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(100*JPYX_CF)), sdk.NewCoin("bjpy", sdk.NewInt(100*BJPY_CF))),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(1*JPYX_CF))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "exceeds the allowable amount as determined by the collateralization ratio",
			},
		},
		{
			"invalid: no price for asset",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("100000000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("5.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.6"),
				priceBTCB:                 sdk.MustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.01"),
				priceBNB:                  sdk.MustNewDecFromStr("0.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(100*KAVA_CF))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("xyz", sdk.NewInt(1))),
				expectedAccountBalance:    sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(20*KAVA_CF)), sdk.NewCoin("btcb", sdk.NewInt(100*BTCB_CF)), sdk.NewCoin("bnb", sdk.NewInt(100*BNB_CF)), sdk.NewCoin("xyz", sdk.NewInt(1))),
				expectedModAccountBalance: sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(1080*KAVA_CF)), sdk.NewCoin("jpyx", sdk.NewInt(200*JPYX_CF)), sdk.NewCoin("bjpy", sdk.NewInt(100*BJPY_CF))),
			},
			errArgs{
				expectPass: false,
				contains:   "no price found for market",
			},
		},
		{
			"invalid: borrow exceed module account balance",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("100000000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("2.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.8"),
				priceBTCB:                 sdk.MustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.01"),
				priceBNB:                  sdk.MustNewDecFromStr("0.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(100*KAVA_CF))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("bjpy", sdk.NewInt(101*BJPY_CF))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "exceeds borrowable module account balance",
			},
		},
		{
			"invalid: over global asset borrow limit",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("20000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("2.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.8"),
				priceBTCB:                 sdk.MustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.01"),
				priceBNB:                  sdk.MustNewDecFromStr("0.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(50*KAVA_CF))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(25*JPYX_CF))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "fails global asset borrow limit validation",
			},
		},
		{
			"invalid: borrowing an individual coin type results in a borrow that's under the minimum JPY borrow limit",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("20000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("2.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.8"),
				priceBTCB:                 sdk.MustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.01"),
				priceBNB:                  sdk.MustNewDecFromStr("0.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(50*KAVA_CF))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(5*JPYX_CF))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "below the minimum borrow limit",
			},
		},
		{
			"invalid: borrowing multiple coins results in a borrow that's under the minimum JPY borrow limit",
			args{
				jpyxBorrowLimit:           sdk.MustNewDecFromStr("20000000"),
				priceKAVA:                 sdk.MustNewDecFromStr("2.00"),
				loanToValueKAVA:           sdk.MustNewDecFromStr("0.8"),
				priceBTCB:                 sdk.MustNewDecFromStr("0.00"),
				loanToValueBTCB:           sdk.MustNewDecFromStr("0.01"),
				priceBNB:                  sdk.MustNewDecFromStr("0.00"),
				loanToValueBNB:            sdk.MustNewDecFromStr("0.01"),
				borrower:                  sdk.AccAddress(crypto.AddressHash([]byte("test"))),
				depositCoins:              sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(50*KAVA_CF))),
				previousBorrowCoins:       sdk.NewCoins(),
				borrowCoins:               sdk.NewCoins(sdk.NewCoin("jpyx", sdk.NewInt(5*JPYX_CF)), sdk.NewCoin("ujsmn", sdk.NewInt(2*JPYX_CF))),
				expectedAccountBalance:    sdk.NewCoins(),
				expectedModAccountBalance: sdk.NewCoins(),
			},
			errArgs{
				expectPass: false,
				contains:   "below the minimum borrow limit",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// Initialize test app and set context
			tApp := app.NewTestApp()
			ctx := tApp.NewContext(true, abci.Header{Height: 1, Time: tmtime.Now()})

			// Auth module genesis state
			authGS := app.NewAuthGenState(
				[]sdk.AccAddress{tc.args.borrower},
				[]sdk.Coins{sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(100*KAVA_CF)),
					sdk.NewCoin("btcb", sdk.NewInt(100*BTCB_CF)), sdk.NewCoin("bnb", sdk.NewInt(100*BNB_CF)),
					sdk.NewCoin("xyz", sdk.NewInt(1)))})

			// hard module genesis state
			hardGS := types.NewGenesisState(types.NewParams(
				types.MoneyMarkets{
					types.NewMoneyMarket("jpyx", types.NewBorrowLimit(true, tc.args.jpyxBorrowLimit, sdk.MustNewDecFromStr("1")), "jpyx:jpy", sdk.NewInt(JPYX_CF), types.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
					types.NewMoneyMarket("bjpy", types.NewBorrowLimit(false, sdk.NewDec(100000000*BJPY_CF), sdk.MustNewDecFromStr("1")), "bjpy:jpy", sdk.NewInt(BJPY_CF), types.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
					types.NewMoneyMarket("ujsmn", types.NewBorrowLimit(false, sdk.NewDec(100000000*KAVA_CF), tc.args.loanToValueKAVA), "kava:jpy", sdk.NewInt(KAVA_CF), types.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
					types.NewMoneyMarket("btcb", types.NewBorrowLimit(false, sdk.NewDec(100000000*BTCB_CF), tc.args.loanToValueBTCB), "btcb:jpy", sdk.NewInt(BTCB_CF), types.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
					types.NewMoneyMarket("bnb", types.NewBorrowLimit(false, sdk.NewDec(100000000*BNB_CF), tc.args.loanToValueBNB), "bnb:jpy", sdk.NewInt(BNB_CF), types.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
					types.NewMoneyMarket("xyz", types.NewBorrowLimit(false, sdk.NewDec(1), tc.args.loanToValueBNB), "xyz:jpy", sdk.NewInt(1), types.NewInterestRateModel(sdk.MustNewDecFromStr("0.05"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10")), sdk.MustNewDecFromStr("0.05"), sdk.ZeroDec()),
				},
				sdk.NewDec(10),
			), types.DefaultAccumulationTimes, types.DefaultDeposits, types.DefaultBorrows,
				types.DefaultTotalSupplied, types.DefaultTotalBorrowed, types.DefaultTotalReserves,
			)

			// Pricefeed module genesis state
			pricefeedGS := pricefeed.GenesisState{
				Params: pricefeed.Params{
					Markets: []pricefeed.Market{
						{MarketID: "jpyx:jpy", BaseAsset: "jpyx", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "bjpy:jpy", BaseAsset: "bjpy", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "kava:jpy", BaseAsset: "kava", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "btcb:jpy", BaseAsset: "btcb", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "bnb:jpy", BaseAsset: "bnb", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
						{MarketID: "xyz:jpy", BaseAsset: "xyz", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
					},
				},
				PostedPrices: []pricefeed.PostedPrice{
					{
						MarketID:      "jpyx:jpy",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "bjpy:jpy",
						OracleAddress: sdk.AccAddress{},
						Price:         sdk.MustNewDecFromStr("1.00"),
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "kava:jpy",
						OracleAddress: sdk.AccAddress{},
						Price:         tc.args.priceKAVA,
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "btcb:jpy",
						OracleAddress: sdk.AccAddress{},
						Price:         tc.args.priceBTCB,
						Expiry:        time.Now().Add(1 * time.Hour),
					},
					{
						MarketID:      "bnb:jpy",
						OracleAddress: sdk.AccAddress{},
						Price:         tc.args.priceBNB,
						Expiry:        time.Now().Add(1 * time.Hour),
					},
				},
			}

			// Initialize test application
			tApp.InitializeFromGenesisStates(authGS,
				app.GenesisState{pricefeed.ModuleName: pricefeed.ModuleCdc.MustMarshalJSON(pricefeedGS)},
				app.GenesisState{types.ModuleName: types.ModuleCdc.MustMarshalJSON(hardGS)})

			// Mint coins to hard module account
			supplyKeeper := tApp.GetSupplyKeeper()
			hardMaccCoins := sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(1000*KAVA_CF)),
				sdk.NewCoin("jpyx", sdk.NewInt(200*JPYX_CF)), sdk.NewCoin("bjpy", sdk.NewInt(100*BJPY_CF)))
			supplyKeeper.MintCoins(ctx, types.ModuleAccountName, hardMaccCoins)

			keeper := tApp.GetHardKeeper()
			suite.app = tApp
			suite.ctx = ctx
			suite.keeper = keeper

			var err error

			// Run BeginBlocker once to transition MoneyMarkets
			hard.BeginBlocker(suite.ctx, suite.keeper)

			err = suite.keeper.Deposit(suite.ctx, tc.args.borrower, tc.args.depositCoins)
			suite.Require().NoError(err)

			// Execute user's previous borrows
			err = suite.keeper.Borrow(suite.ctx, tc.args.borrower, tc.args.previousBorrowCoins)
			if tc.args.previousBorrowCoins.IsZero() {
				suite.Require().True(strings.Contains(err.Error(), "cannot borrow zero coins"))
			} else {
				suite.Require().NoError(err)
			}

			// Now that our state is properly set up, execute the last borrow
			err = suite.keeper.Borrow(suite.ctx, tc.args.borrower, tc.args.borrowCoins)

			if tc.errArgs.expectPass {
				suite.Require().NoError(err)

				// Check borrower balance
				acc := suite.getAccount(tc.args.borrower)
				suite.Require().Equal(tc.args.expectedAccountBalance, acc.GetCoins())

				// Check module account balance
				mAcc := suite.getModuleAccount(types.ModuleAccountName)
				suite.Require().Equal(tc.args.expectedModAccountBalance, mAcc.GetCoins())

				// Check that borrow struct is in store
				_, f := suite.keeper.GetBorrow(suite.ctx, tc.args.borrower)
				suite.Require().True(f)
			} else {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestValidateBorrow() {

	blockDuration := time.Second * 3600 * 24 // long blocks to accumulate larger interest

	_, addrs := app.GeneratePrivKeyAddressPairs(5)
	borrower := addrs[0]
	initialBorrowerBalance := sdk.NewCoins(
		sdk.NewCoin("ujsmn", sdk.NewInt(1000*KAVA_CF)),
		sdk.NewCoin("jpyx", sdk.NewInt(1000*KAVA_CF)),
	)

	model := types.NewInterestRateModel(sdk.MustNewDecFromStr("1.0"), sdk.MustNewDecFromStr("2"), sdk.MustNewDecFromStr("0.8"), sdk.MustNewDecFromStr("10"))

	// Initialize test app and set context
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{Height: 1, Time: tmtime.Now()})

	// Auth module genesis state
	authGS := app.NewAuthGenState(
		[]sdk.AccAddress{borrower},
		[]sdk.Coins{initialBorrowerBalance})

	// Hard module genesis state
	hardGS := types.NewGenesisState(
		types.NewParams(
			types.MoneyMarkets{
				types.NewMoneyMarket("jpyx",
					types.NewBorrowLimit(false, sdk.NewDec(100000000*JPYX_CF), sdk.MustNewDecFromStr("1")), // Borrow Limit
					"jpyx:jpy",                     // Market ID
					sdk.NewInt(JPYX_CF),            // Conversion Factor
					model,                          // Interest Rate Model
					sdk.MustNewDecFromStr("1.0"),   // Reserve Factor (high)
					sdk.MustNewDecFromStr("0.05")), // Keeper Reward Percent
				types.NewMoneyMarket("ujsmn",
					types.NewBorrowLimit(false, sdk.NewDec(100000000*KAVA_CF), sdk.MustNewDecFromStr("0.8")), // Borrow Limit
					"kava:jpy",                     // Market ID
					sdk.NewInt(KAVA_CF),            // Conversion Factor
					model,                          // Interest Rate Model
					sdk.MustNewDecFromStr("1.0"),   // Reserve Factor (high)
					sdk.MustNewDecFromStr("0.05")), // Keeper Reward Percent
			},
			sdk.NewDec(10),
		),
		types.DefaultAccumulationTimes,
		types.DefaultDeposits,
		types.DefaultBorrows,
		types.DefaultTotalSupplied,
		types.DefaultTotalBorrowed,
		types.DefaultTotalReserves,
	)

	// Pricefeed module genesis state
	pricefeedGS := pricefeed.GenesisState{
		Params: pricefeed.Params{
			Markets: []pricefeed.Market{
				{MarketID: "jpyx:jpy", BaseAsset: "jpyx", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
				{MarketID: "kava:jpy", BaseAsset: "kava", QuoteAsset: "jpy", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []pricefeed.PostedPrice{
			{
				MarketID:      "jpyx:jpy",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("1.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				MarketID:      "kava:jpy",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("2.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}

	// Initialize test application
	tApp.InitializeFromGenesisStates(
		authGS,
		app.GenesisState{pricefeed.ModuleName: pricefeed.ModuleCdc.MustMarshalJSON(pricefeedGS)},
		app.GenesisState{types.ModuleName: types.ModuleCdc.MustMarshalJSON(hardGS)},
	)

	keeper := tApp.GetHardKeeper()
	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = keeper

	var err error

	// Run BeginBlocker once to transition MoneyMarkets
	hard.BeginBlocker(suite.ctx, suite.keeper)

	// Setup borrower with some collateral to borrow against, and some reserve in the protocol.
	depositCoins := sdk.NewCoins(
		sdk.NewCoin("ujsmn", sdk.NewInt(100*KAVA_CF)),
		sdk.NewCoin("jpyx", sdk.NewInt(100*JPYX_CF)),
	)
	err = suite.keeper.Deposit(suite.ctx, borrower, depositCoins)
	suite.Require().NoError(err)

	initialBorrowCoins := sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(70*KAVA_CF)))
	err = suite.keeper.Borrow(suite.ctx, borrower, initialBorrowCoins)
	suite.Require().NoError(err)

	runAtTime := suite.ctx.BlockTime().Add(blockDuration)
	suite.ctx = suite.ctx.WithBlockTime(runAtTime)
	hard.BeginBlocker(suite.ctx, suite.keeper)

	repayCoins := sdk.NewCoins(sdk.NewCoin("ujsmn", sdk.NewInt(100*KAVA_CF))) // repay everything including accumulated interest
	err = suite.keeper.Repay(suite.ctx, borrower, borrower, repayCoins)
	suite.Require().NoError(err)

	// Get the total borrowable amount from the protocol, taking into account the reserves.
	modAccBalance := suite.getModuleAccountAtCtx(types.ModuleAccountName, suite.ctx).GetCoins()
	reserves, found := suite.keeper.GetTotalReserves(suite.ctx)
	suite.Require().True(found)
	availableToBorrow := modAccBalance.Sub(reserves)

	// Test borrowing one over the available amount (try to borrow from the reserves)
	err = suite.keeper.Borrow(
		suite.ctx,
		borrower,
		sdk.NewCoins(sdk.NewCoin("ujsmn", availableToBorrow.AmountOf("ujsmn").Add(sdk.OneInt()))),
	)
	suite.Require().Error(err)

	// Test borrowing exactly the limit
	err = suite.keeper.Borrow(
		suite.ctx,
		borrower,
		sdk.NewCoins(sdk.NewCoin("ujsmn", availableToBorrow.AmountOf("ujsmn"))),
	)
	suite.Require().NoError(err)
}
