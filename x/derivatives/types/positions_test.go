package types_test

import (
	"fmt"
	"testing"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

// make position.NeedLiquidationPerpetualFutures test
func TestPosition_NeedLiquidationPerpetualFutures(t *testing.T) {
	// owner := sdk.MustAccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")
	owner, _ := sdk.AccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")
	testCases := []struct {
		name          string
		position      types.PerpetualFuturesPosition
		minMarginRate sdk.Dec
		closedRate    sdk.Dec
		exp           bool
	}{
		{
			name: "default is valid",
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
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("10"),
					Leverage:     5,
				},
			},
			minMarginRate: sdk.MustNewDecFromStr("0.1"),
			closedRate:    sdk.MustNewDecFromStr("100"),
			exp:           true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.position.NeedLiquidationPerpetualFutures(tc.minMarginRate, tc.closedRate)
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

func TestPosition_GetMarginMaintenanceRate(t *testing.T) {
	// owner := sdk.MustAccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")
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
				RemainingMargin: sdk.NewCoin("uatom", sdk.NewInt(1)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
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
					rate: sdk.MustNewDecFromStr("90"),
				},
			},
			exp: sdk.MustNewDecFromStr("0.9"),
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
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1)),
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
					rate: sdk.MustNewDecFromStr("110"),
				},
			},
			exp: sdk.MustNewDecFromStr("1.1"),
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
				RemainingMargin: sdk.NewCoin("ubtc", sdk.NewInt(1)),
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
					rate: sdk.MustNewDecFromStr("90"),
				},
			},
			exp: sdk.MustNewDecFromStr("1.111111111111111111"),
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
				RemainingMargin: sdk.NewCoin("ubtc", sdk.NewInt(1)),
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
					rate: sdk.MustNewDecFromStr("110"),
				},
			},
			exp: sdk.MustNewDecFromStr("0.909090909090909091"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.position.GetMarginMaintenanceRate(tc.Rate[0].rate, tc.Rate[1].rate)
			if !tc.exp.Equal(result) {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}

// make PerpetualFuturesPosition.CalcProfit test
func TestPerpetualFuturesPosition_CalcProfit(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")
	testCases := []struct {
		name       string
		position   types.PerpetualFuturesPosition
		closedRate sdk.Dec
		exp        types.Revenue
	}{
		{
			name: "Long position profit",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRate: sdk.MustNewDecFromStr("11.1"),
			exp: types.Revenue{
				RevenueType: types.RevenueType_PROFIT,
				Amount:      sdk.NewCoin("uatom", sdk.NewInt(550)),
			},
		},
		{
			name: "Long position loss",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_LONG,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRate: sdk.MustNewDecFromStr("9.1"),
			exp: types.Revenue{
				RevenueType: types.RevenueType_LOSS,
				Amount:      sdk.NewCoin("uatom", sdk.NewInt(450)),
			},
		},
		{
			name: "Short position profit",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRate: sdk.MustNewDecFromStr("9.1"),
			exp: types.Revenue{
				RevenueType: types.RevenueType_PROFIT,
				Amount:      sdk.NewCoin("uatom", sdk.NewInt(450)),
			},
		},
		{
			name: "Short position loss",
			position: types.PerpetualFuturesPosition{
				Id:      "0",
				Address: owner.Bytes(),
				Market: types.Market{
					BaseDenom:  "uatom",
					QuoteDenom: "uusdc",
				},
				OpenedAt:        time.Now().UTC(),
				OpenedHeight:    1,
				OpenedBaseRate:  sdk.MustNewDecFromStr("10"),
				RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
				PositionInstance: types.PerpetualFuturesPositionInstance{
					PositionType: types.PositionType_SHORT,
					Size_:        sdk.MustNewDecFromStr("100"),
					Leverage:     5,
				},
			},
			closedRate: sdk.MustNewDecFromStr("12.1"),
			exp: types.Revenue{
				RevenueType: types.RevenueType_LOSS,
				Amount:      sdk.NewCoin("uatom", sdk.NewInt(1050)),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.position.CalcProfit(tc.closedRate)
			fmt.Println(result)
			if !tc.exp.Equal(result) {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}
