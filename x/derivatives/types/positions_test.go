package types_test

import (
	"fmt"
	"testing"
	time "time"

	"cosmossdk.io/math"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

// IsValid test. This is the general validation for the creation of a position (only perpetual futures position)
func TestPosition_IsValid(t *testing.T) {
	// make testCases
	testCases := []struct {
		name     string
		position types.Position
		instance types.PerpetualFuturesPositionInstance
		exp      bool
	}{
		{
			name: "Failure due to invalid margin asset",
			position: types.Position{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("0.000001"),
				// not market base or quote asset
				RemainingMargin: sdk.NewCoin("ubtc", sdk.NewInt(1)),
			},
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("10"),
				Leverage:     5,
			},
			exp: false,
		},
		{
			name: "Failure due to lack of margin using BaseDenom token",
			position: types.Position{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("0.000001"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(100000)),
			},
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("10"),
				Leverage:     10,
			},
			exp: false,
		},
		// below test case fails now because the current implementation doesn't calculate
		// the margin requirement in a proper way
		{
			name: "Failure due to lack of margin using QuoteDenom token",
			position: types.Position{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("0.000001"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
			},
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("10"),
				Leverage:     1,
			},
			exp: false,
		},
		{
			name: "Fauilure due to zero margin",
			position: types.Position{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("0.000001"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(0)),
			},
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("10"),
				Leverage:     10,
			},
			exp: false,
		},
	}

	params := types.DefaultParams()
	// run testCases
	for _, tc := range testCases {
		any, err := codecTypes.NewAnyWithValue(&tc.instance)
		if err != nil {
			t.Error(err)
		}
		tc.position.PositionInstance = *any

		t.Run(tc.name, func(t *testing.T) {
			err := tc.position.IsValid(params)
			if tc.exp {
				if err != nil {
					t.Errorf("expected %v, got %v", tc.exp, err)
				}
			} else {
				if err == nil {
					t.Errorf("expected %v, got %v", tc.exp, err)
				}
			}
		})
	}
}

// make position.NeedLiquidationPerpetualFutures test
func TestPosition_NeedLiquidationPerpetualFutures(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")
	testCases := []struct {
		name          string
		position      types.PerpetualFuturesPosition
		minMarginRate sdk.Dec
		closedRate    []sdk.Dec //first is base rate, second is quote rate
		exp           bool
	}{
		{
			name: "not_change_rate_is_not_need_liquidation",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("10"),
					Leverage:     5,
				},
			},
			minMarginRate: sdk.MustNewDecFromStr("0.5"),
			closedRate: []sdk.Dec{
				sdk.MustNewDecFromStr("100"),
				sdk.MustNewDecFromStr("100"),
			},
			exp: false,
		},
		{
			name: "down_rate_is_need_liquidation",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			minMarginRate: sdk.MustNewDecFromStr("0.5"),
			closedRate: []sdk.Dec{
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("100"),
			},
			exp: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			quoteTicker := "usd"
			baseMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.BaseDenom, tc.closedRate[0])
			quoteMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.QuoteDenom, tc.closedRate[1])

			result := tc.position.NeedLiquidation(tc.minMarginRate, baseMetricsRate, quoteMetricsRate)
			if tc.exp != result {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}

type CurrencyRate struct {
	name string
	rate sdk.Dec
}

func TestPosition_MarginMaintenanceRate(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")

	testCases := []struct {
		name     string
		position types.PerpetualFuturesPosition
		Rate     []CurrencyRate
		exp      sdk.Dec
	}{
		{
			name: "long position not change rate",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:     time.Now().UTC(),
				OpenedHeight: 1,
				// atom/usd = 0.4
				OpenedBaseRate:  sdk.MustNewDecFromStr("400"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("400"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			// not change rate
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("400"),
				},
				{
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("400"),
				},
			},
			exp: sdk.MustNewDecFromStr("1"),
		},
		{
			name: "long position down 10%",
			position: types.PerpetualFuturesPosition{
				Id:      "1",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(100)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			// down 10%
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("90"),
				},
				{
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("100"),
				},
			},
			exp: sdk.MustNewDecFromStr("99.888888888888888889"),
		},
		{
			name: "long position up 10%",
			position: types.PerpetualFuturesPosition{
				Id:      "3",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(100)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("110"),
				},
				{
					// up 10%
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("100"),
				},
			},
			exp: sdk.MustNewDecFromStr("102.272727272727272727"),
		},
		{
			name: "short position not change rate",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("400"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("400"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("ubtc", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			// not change rate
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("400"),
				},
				{
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("400"),
				},
			},
			exp: sdk.MustNewDecFromStr("1"),
		},
		{
			name: "short position down 10%",
			position: types.PerpetualFuturesPosition{
				Id:      "1",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("ubtc", sdk.NewInt(100)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			// down 10%
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("90"),
				},
				{
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("100"),
				},
			},
			exp: sdk.MustNewDecFromStr("113.888888888888888889"),
		},
		{
			name: "short position up 10%",
			position: types.PerpetualFuturesPosition{
				Id:      "3",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "ubtc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("100"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("100"),
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("ubtc", sdk.NewInt(100)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("110"),
				},
				{
					// up 10%
					name: "ubtc/usd",
					rate: sdk.MustNewDecFromStr("100"),
				},
			},
			exp: sdk.MustNewDecFromStr("88.636363636363636364"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			quoteTicker := "usd"
			baseMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.BaseDenom, tc.Rate[0].rate)
			quoteMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.QuoteDenom, tc.Rate[1].rate)
			result := tc.position.MarginMaintenanceRate(baseMetricsRate, quoteMetricsRate)
			if !tc.exp.Equal(result) {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}

// make PerpetualFuturesPosition.CalcProfitAndLoss test
func TestPerpetualFuturesPosition_CalcProfitAndLoss(t *testing.T) {
	testCases := []struct {
		name        string
		position    types.PerpetualFuturesPosition
		closedRates []sdk.Dec
		exp         math.Int
	}{
		{
			name: "Long position profit",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("11.1"),
				sdk.MustNewDecFromStr("11.1"),
			},
			exp: sdk.NewInt(110000000),
		},
		{
			name: "Long position loss",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("9.1"),
				sdk.MustNewDecFromStr("9.1"),
			},
			exp: sdk.NewInt(-90000000),
		},
		{
			name: "Short position profit",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("9.1"),
				sdk.MustNewDecFromStr("9.1"),
			},
			exp: sdk.NewInt(90000000),
		},
		{
			name: "Short position loss",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("12.1"),
				sdk.MustNewDecFromStr("12.1"),
			},
			exp: sdk.NewInt(-210000000),
		},
		{
			name: "Profit Long position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("20"),
				sdk.MustNewDecFromStr("20"),
			},
			exp: sdk.NewInt(500000),
		},
		{
			name: "Profit Short position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("20"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     20,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("10"),
				sdk.MustNewDecFromStr("10"),
			},
			exp: sdk.NewInt(1000000),
		},
		{
			name: "Loss Long position in base denom margin ",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("20"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     20,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("10"),
				sdk.MustNewDecFromStr("10"),
			},
			exp: sdk.NewInt(-1000000),
		},
		{
			name: "Loss Short position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("20"),
				sdk.MustNewDecFromStr("20"),
			},
			exp: sdk.NewInt(-500000),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			quoteTicker := "usd"
			baseMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.BaseDenom, tc.closedRates[0])
			quoteMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.QuoteDenom, tc.closedRates[1])
			resultDec := tc.position.ProfitAndLossInMetrics(baseMetricsRate, quoteMetricsRate)
			result := types.NormalToMicroInt(resultDec)
			fmt.Println(result)
			if !tc.exp.Equal(result) {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}

// CalcReturningAmountAtClose test
func TestCalcReturningAmountAtClose(t *testing.T) {
	uusdcRate := sdk.MustNewDecFromStr("0.000001")

	testCases := []struct {
		name           string
		position       types.PerpetualFuturesPosition
		closedBaseRate sdk.Dec
		closeQuoteRate sdk.Dec
		expReturn      math.Int
		expLoss        math.Int
	}{
		{
			name: "Profit Long position in quote denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedBaseRate: sdk.MustNewDecFromStr("0.00002"),
			closeQuoteRate: uusdcRate,
			expReturn:      sdk.NewInt(11000000),
		},
		{
			name: "Profit Short position in quote denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00002"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     20,
				},
			},
			closedBaseRate: sdk.MustNewDecFromStr("0.00001"),
			closeQuoteRate: uusdcRate,
			expReturn:      sdk.NewInt(11000000),
		},
		{
			name: "Loss Long position in quote denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedBaseRate: sdk.MustNewDecFromStr("0.000009"),
			closeQuoteRate: uusdcRate,
			expReturn:      sdk.NewInt(0),
		},
		{
			name: "Loss Short position in quote denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			closedBaseRate: sdk.MustNewDecFromStr("0.0000105"),
			closeQuoteRate: uusdcRate,
			expReturn:      sdk.NewInt(500000),
		},
		{
			name: "Profit Long position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			closedBaseRate: sdk.MustNewDecFromStr("0.00002"),
			closeQuoteRate: uusdcRate,
			expReturn:      sdk.NewInt(1500000),
		},
		{
			name: "Profit Short position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00002"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			closedBaseRate: sdk.MustNewDecFromStr("0.00001"),
			closeQuoteRate: uusdcRate,
			expReturn:      sdk.NewInt(2000000),
		},
		{
			name: "Loss Long position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     10,
				},
			},
			closedBaseRate: sdk.MustNewDecFromStr("0.000009"),
			closeQuoteRate: uusdcRate,
			expReturn:      sdk.NewInt(888889),
		},
		{
			name: "Loss Short position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			closedBaseRate: sdk.MustNewDecFromStr("0.000011"),
			closeQuoteRate: uusdcRate,
			expReturn:      sdk.NewInt(909091),
		},
		{
			name: "Loss Exceeds Margin: Long position in base denom margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("10"),
					Leverage:     10,
				},
			},
			closedBaseRate: sdk.MustNewDecFromStr("0.000009"),
			closeQuoteRate: uusdcRate,
			expReturn:      sdk.NewInt(0),
			expLoss:        sdk.NewInt(-111111),
		},
	}

	quoteTicker := "usd"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			baseMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.BaseDenom, tc.closedBaseRate)
			quoteMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.QuoteDenom, tc.closeQuoteRate)

			returningAmount, lossToLP := tc.position.CalcReturningAmountAtClose(baseMetricsRate, quoteMetricsRate)
			fmt.Println(returningAmount, lossToLP)
			if !tc.expReturn.Equal(returningAmount) {
				t.Error(tc, "expected %v, got %v", tc.expReturn, returningAmount)
			}

			if !tc.expLoss.IsNil() {
				if !tc.expLoss.Equal(lossToLP) {
					t.Error(tc, "expected %v, got %v", tc.expLoss, lossToLP)
				}
			}
		})
	}
}
