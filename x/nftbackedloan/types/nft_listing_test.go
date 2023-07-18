package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func TestAddCollectedAmount(t *testing.T) {
	testCases := []struct {
		name        string
		listing     types.NftListing
		amount      sdk.Coin
		expCoin     sdk.Coin
		expNegative bool
	}{
		{
			"Add 0 to positive",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(100)),
				CollectedAmountNegative: false,
			},
			sdk.NewCoin("uguu", sdk.NewInt(0)),
			sdk.NewCoin("uguu", sdk.NewInt(100)),
			false,
		},
		{
			"Add 0 to negative",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(60)),
				CollectedAmountNegative: true,
			},
			sdk.NewCoin("uguu", sdk.NewInt(0)),
			sdk.NewCoin("uguu", sdk.NewInt(60)),
			true,
		},
		{
			"Add to positive",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(100)),
				CollectedAmountNegative: false,
			},
			sdk.NewCoin("uguu", sdk.NewInt(100)),
			sdk.NewCoin("uguu", sdk.NewInt(200)),
			false,
		},
		{
			"Add to negative (add < negative)",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(100)),
				CollectedAmountNegative: true,
			},
			sdk.NewCoin("uguu", sdk.NewInt(50)),
			sdk.NewCoin("uguu", sdk.NewInt(50)),
			true,
		},
		{
			"Add to negative (add > negative)",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(100)),
				CollectedAmountNegative: true,
			},
			sdk.NewCoin("uguu", sdk.NewInt(200)),
			sdk.NewCoin("uguu", sdk.NewInt(100)),
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.listing.AddCollectedAmount(tc.amount)
			require.Equal(t, tc.expCoin, result.CollectedAmount)
			require.Equal(t, tc.expNegative, result.CollectedAmountNegative)
		})
	}
}

func TestSubCollectedAmount(t *testing.T) {
	testCases := []struct {
		name        string
		listing     types.NftListing
		amount      sdk.Coin
		expCoin     sdk.Coin
		expNegative bool
	}{
		{
			"Sub 0 from positive",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(100)),
				CollectedAmountNegative: false,
			},
			sdk.NewCoin("uguu", sdk.NewInt(0)),
			sdk.NewCoin("uguu", sdk.NewInt(100)),
			false,
		},
		{
			"Sub 0 from negative",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(60)),
				CollectedAmountNegative: true,
			},
			sdk.NewCoin("uguu", sdk.NewInt(0)),
			sdk.NewCoin("uguu", sdk.NewInt(60)),
			true,
		},
		{
			"Sub from negative",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(100)),
				CollectedAmountNegative: true,
			},
			sdk.NewCoin("uguu", sdk.NewInt(100)),
			sdk.NewCoin("uguu", sdk.NewInt(200)),
			true,
		},
		{
			"sub from positive (sub < positive)",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(100)),
				CollectedAmountNegative: false,
			},
			sdk.NewCoin("uguu", sdk.NewInt(50)),
			sdk.NewCoin("uguu", sdk.NewInt(50)),
			false,
		},
		{
			"sub from positive (sub > positive)",
			types.NftListing{
				CollectedAmount:         sdk.NewCoin("uguu", sdk.NewInt(100)),
				CollectedAmountNegative: false,
			},
			sdk.NewCoin("uguu", sdk.NewInt(150)),
			sdk.NewCoin("uguu", sdk.NewInt(50)),
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.listing.SubCollectedAmount(tc.amount)
			require.Equal(t, tc.expCoin, result.CollectedAmount)
			require.Equal(t, tc.expNegative, result.CollectedAmountNegative)
		})
	}
}
