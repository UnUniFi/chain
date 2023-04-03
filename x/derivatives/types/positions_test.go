package types_test

import (
	"fmt"
	"testing"

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
				Leverage:     1,
			},
			exp: false,
		},
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
		{
			name: "Fauilure due to invalid levarage",
			position: types.Position{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("0.000001"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			},
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("10"),
				Leverage:     31,
			},
			exp: false,
		},
		{
			name: "Fauilure due to invalid levarage",
			position: types.Position{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("0.000001"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			},
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("10"),
				Leverage:     1,
			},
			exp: false,
		},
		{
			name: "Success",
			position: types.Position{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: sdk.MustNewDecFromStr("0.000001"),
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			},
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     1,
			},
			exp: false,
		},
	}

	params := types.DefaultParams()
	// run testCases
	for _, tc := range testCases {
		sizeInMicro := tc.instance.Size_.MulInt64(types.OneMillionInt).TruncateInt()
		tc.instance.SizeInMicro = &sizeInMicro
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
	uusdcRate := sdk.MustNewDecFromStr("0.000001")

	testCases := []struct {
		name          string
		position      types.PerpetualFuturesPosition
		minMarginRate sdk.Dec
		closedRate    []sdk.Dec //first is base rate, second is quote rate
		exp           bool
	}{
		{
			name: "False: change from opened rate",
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
			minMarginRate: sdk.MustNewDecFromStr("0.5"),
			closedRate: []sdk.Dec{
				sdk.MustNewDecFromStr("0.00001"),
				uusdcRate,
			},
			exp: false,
		},
		{
			name: "True: Price down for long position",
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
			minMarginRate: sdk.MustNewDecFromStr("0.5"),
			// In this case, the margin maintanance rate is gonna be "0.5"
			// Which is the defined rate of the liquidation criteria
			closedRate: []sdk.Dec{
				sdk.MustNewDecFromStr("0.0000095"),
				uusdcRate,
			},
			exp: true,
		},
		{
			name: "True: Price up for short position",
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
					Size_:        sdk.MustNewDecFromStr("10"),
					Leverage:     10,
				},
			},
			minMarginRate: sdk.MustNewDecFromStr("0.5"),
			// In this case, the margin maintanance rate is gonna be "0.5"
			// Which is the defined rate of the liquidation criteria
			closedRate: []sdk.Dec{
				sdk.MustNewDecFromStr("0.0000106"),
				uusdcRate,
			},
			exp: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			quoteTicker := "usd"
			baseMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.BaseDenom, tc.closedRate[0])
			quoteMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.QuoteDenom, tc.closedRate[1])
			sizeInMicro := tc.position.PositionInstance.Size_.MulInt64(types.OneMillionInt).TruncateInt()
			tc.position.PositionInstance.SizeInMicro = &sizeInMicro

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
	uusdcRate := sdk.MustNewDecFromStr("0.000001")

	testCases := []struct {
		name     string
		position types.PerpetualFuturesPosition
		Rate     []CurrencyRate
		exp      sdk.Dec
	}{
		{
			name: "long position not change rate",
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
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			// not change rate
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("0.00001"),
				},
				{
					name: "uusdc/usd",
					rate: uusdcRate,
				},
			},
			exp: sdk.MustNewDecFromStr("1"),
		},
		{
			name: "long position down 10%",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				// In the case of Long, BaseDenom is RemainingMargin.
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
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
					rate: sdk.MustNewDecFromStr("0.000009"),
				},
				{
					name: "uusdc/usd",
					rate: uusdcRate,
				},
			},
			exp: sdk.MustNewDecFromStr("0.888888888888888889"),
		},
		{
			name: "long position up 10%",
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
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("0.000011"),
				},
				{
					name: "uusdc/usd",
					rate: uusdcRate,
				},
			},
			exp: sdk.MustNewDecFromStr("1.090909090909090909"),
		},
		{
			name: "short position not change rate",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
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
					rate: sdk.MustNewDecFromStr("0.00001"),
				},
				{
					name: "uusdc/usd",
					rate: uusdcRate,
				},
			},
			exp: sdk.MustNewDecFromStr("1"),
		},
		{
			name: "short position down 10%",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
					Leverage:     1,
				},
			},
			// down 10%
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("0.000009"),
				},
				{
					name: "uusdc/usd",
					rate: uusdcRate,
				},
			},
			exp: sdk.MustNewDecFromStr("1.222222222222222222"),
		},
		{
			name: "short position up 10%",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("5"),
					Leverage:     5,
				},
			},
			Rate: []CurrencyRate{
				{
					name: "uatom/usd",
					rate: sdk.MustNewDecFromStr("0.000011"),
				},
				{
					name: "uusdc/usd",
					rate: uusdcRate,
				},
			},
			exp: sdk.MustNewDecFromStr("0.454545454545454545"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			quoteTicker := "usd"
			baseMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.BaseDenom, tc.Rate[0].rate)
			quoteMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.QuoteDenom, tc.Rate[1].rate)

			sizeInMicro := tc.position.PositionInstance.Size_.MulInt64(types.OneMillionInt).TruncateInt()
			tc.position.PositionInstance.SizeInMicro = &sizeInMicro

			result := tc.position.MarginMaintenanceRate(baseMetricsRate, quoteMetricsRate)
			if !tc.exp.Equal(result) {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}

// make PerpetualFuturesPosition.CalcProfitAndLoss test
func TestPerpetualFuturesPosition_CalcProfitAndLoss(t *testing.T) {
	uusdcRate := sdk.MustNewDecFromStr("0.000001")

	testCases := []struct {
		name        string
		position    types.PerpetualFuturesPosition
		closedRates []sdk.Dec
		exp         math.Int
	}{
		{
			name: "Long position profit in Base Denom Margin",
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
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("0.000011"),
				uusdcRate,
			},
			exp: sdk.NewInt(90909),
		},
		{
			name: "Long position profit in Quote Denom Margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("0.000011"),
				uusdcRate,
			},
			exp: sdk.NewInt(1000000),
		},
		{
			name: "Long position loss in Base Denom Margin",
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
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("0.000009"),
				uusdcRate,
			},
			exp: sdk.NewInt(-111111),
		},
		{
			name: "Long position loss in Quote Denom Margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("1"),
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("0.000009"),
				uusdcRate,
			},
			exp: sdk.NewInt(-1000000),
		},
		{
			name: "Short position profit in Base Denom Margin",
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
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("0.000009"),
				uusdcRate,
			},
			exp: sdk.NewInt(111111),
		},
		{
			name: "Short position profit in Quote Denom Margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("0.000009"),
				uusdcRate,
			},
			exp: sdk.NewInt(1000000),
		},
		{
			name: "Short position loss in Base Denom Margin",
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
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("0.000011"),
				uusdcRate,
			},
			exp: sdk.NewInt(-90909),
		},
		{
			name: "Short position loss in Quote Denom Margin",
			position: types.PerpetualFuturesPosition{
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedBaseRate:  sdk.MustNewDecFromStr("0.00001"),
				OpenedQuoteRate: uusdcRate,
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("1"),
				},
			},
			closedRates: []sdk.Dec{
				sdk.MustNewDecFromStr("0.000011"),
				uusdcRate,
			},
			exp: sdk.NewInt(-1000000),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			quoteTicker := "usd"
			baseMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.BaseDenom, tc.closedRates[0])
			quoteMetricsRate := types.NewMetricsRateType(quoteTicker, tc.position.Market.QuoteDenom, tc.closedRates[1])

			sizeInMicro := tc.position.PositionInstance.Size_.MulInt64(types.OneMillionInt).TruncateInt()
			tc.position.PositionInstance.SizeInMicro = &sizeInMicro

			// ProfitAndLoss returns PnL in the Margin Denom
			result := tc.position.ProfitAndLoss(baseMetricsRate, quoteMetricsRate)
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
			expReturn:      sdk.NewInt(1000000),
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
			sizeInMicro := types.NormalToMicroInt(tc.position.PositionInstance.Size_)
			tc.position.PositionInstance.SizeInMicro = &sizeInMicro

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
