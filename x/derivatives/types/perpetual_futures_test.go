package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

// test positionInstance.MarginRequirement
func TestPositionInstance_MarginRequirement(t *testing.T) {
	uusdcRate := sdk.MustNewDecFromStr("0.000001")

	// make testCases
	testCases := []struct {
		name             string
		positionInstance types.PerpetualFuturesPositionInstance
		baseRate         sdk.Dec
		quoteRate        sdk.Dec
		exp              sdk.Int
	}{
		{
			name: "case1",
			positionInstance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     10,
			},
			baseRate:  sdk.MustNewDecFromStr("0.00001"),
			quoteRate: uusdcRate,
			exp:       sdk.NewIntFromUint64(1000000),
		},
		{
			name: "Max Levarage",
			positionInstance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     30,
			},
			baseRate:  sdk.MustNewDecFromStr("0.00001"),
			quoteRate: uusdcRate,
			exp:       sdk.NewIntFromUint64(333333),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			currencyRate := tc.baseRate.Quo(tc.quoteRate)
			sizeInMicro := tc.positionInstance.Size_.MulInt64(types.OneMillionInt).TruncateInt()
			tc.positionInstance.SizeInMicro = &sizeInMicro

			result := tc.positionInstance.MarginRequirement(currencyRate)
			if !tc.exp.Equal(result) {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}
