package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var testCoin = sdk.NewInt64Coin("test", 20)

// AuctionをAny型からAuction型に変換
var unpackAuctions, _ = UnpackGenesisAuctions(DefaultGenesis().Auctions)

func TestGenesisState_Validate(t *testing.T) {
	testCases := []struct {
		name       string
		nextID     uint64
		auctions   GenesisAuctions
		expectPass bool
	}{
		// {"default", DefaultGenesis().NextAuctionId, DefaultGenesis().Auctions, true},
		// {"invalid next ID", 54, GenesisAuctions{SurplusAuction{BaseAuction{ID: 105}}}, false},
		{"default", DefaultGenesis().NextAuctionId, unpackAuctions, true},
		{"invalid next ID", 54, GenesisAuctions{&SurplusAuction{BaseAuction{Id: 105}}}, false},
		{
			"repeated ID",
			1000,
			GenesisAuctions{
				// SurplusAuction{BaseAuction{ID: 105}},
				// DebtAuction{BaseAuction{ID: 105}, testCoin},
				&SurplusAuction{BaseAuction{Id: 105}},
				&DebtAuction{BaseAuction{Id: 105}, testCoin},
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// gs := NewGenesisState(tc.nextID, DefaultParams(), tc.auctions)
			// NewGenesisStateの引数tc.acutionをAny型に変換
			packAuction, _ := PackGenesisAuctions(tc.auctions)
			gs := NewGenesisState(tc.nextID, DefaultParams(), packAuction)

			err := gs.Validate()

			if tc.expectPass {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}

}
