package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

// test positionInstance.MarginRequirement
func TestPositionInstance_MarginRequirement(t *testing.T) {
	// make testCases
	testCases := []struct {
		name             string
		positionInstance types.PerpetualFuturesPositionInstance
		currencyRate     sdk.Dec
		exp              sdk.Dec
	}{
		{
			name: "case1",
			positionInstance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("10000"),
				Leverage:     25,
			},
			currencyRate: sdk.MustNewDecFromStr("100"),
			exp:          sdk.MustNewDecFromStr("40000"),
		},
		{
			name: "case2",
			positionInstance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("10000"),
				Leverage:     50,
			},
			currencyRate: sdk.MustNewDecFromStr("100"),
			exp:          sdk.MustNewDecFromStr("20000"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.positionInstance.MarginRequirement(tc.currencyRate)
			if !tc.exp.Equal(result) {
				t.Error(tc, "expected %v, got %v", tc.exp, result)
			}
		})
	}
}
