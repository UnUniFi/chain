package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestEnsureDenomInPool(t *testing.T) {
	tests := []struct {
		name       string
		poolAssets []sdk.Coin
		TokensIn   []sdk.Coin
		err        bool
	}{
		{
			name: "exist in pool",
			poolAssets: []sdk.Coin{
				sdk.NewCoin("denom", sdk.NewInt(1000)),
			},
			TokensIn: []sdk.Coin{
				sdk.NewCoin("denom", sdk.NewInt(1000)),
			},
			err: false,
		},
		{
			name: "not exist in pool",
			poolAssets: []sdk.Coin{
				sdk.NewCoin("denom", sdk.NewInt(1000)),
			},
			TokensIn: []sdk.Coin{
				sdk.NewCoin("denom2", sdk.NewInt(1000)),
			},
			err: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ensureDenomInPool(tt.poolAssets, tt.TokensIn)
			if tt.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMaximalExactRatioJoin(t *testing.T)    {}
func TestCalcJoinPoolNoSwapShares(t *testing.T) {}

func TestIncreaseLiquidity(t *testing.T) {}
