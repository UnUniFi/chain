package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestStrategyDenoms(t *testing.T) {
	tests := []struct {
		name   string
		vault  Vault
		denoms []string
	}{
		{
			name: "invalid address",
			vault: Vault{
				Id:     1,
				Symbol: "ATOM",
				StrategyWeights: []StrategyWeight{
					{
						Denom:      "uatom1",
						StrategyId: 1,
						Weight:     sdk.OneDec(),
					},
					{
						Denom:      "uatom2",
						StrategyId: 1,
						Weight:     sdk.OneDec(),
					},
					{
						Denom:      "uatom1",
						StrategyId: 1,
						Weight:     sdk.OneDec(),
					},
				},
			},
			denoms: []string{"uatom1", "uatom2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			denoms := tt.vault.StrategyDenoms()
			require.Equal(t, denoms, tt.denoms)
		})
	}
}
