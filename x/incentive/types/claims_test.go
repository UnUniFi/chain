package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	jpyxtypes "github.com/lcnem/jpyx/types"
)

func TestClaimsValidate(t *testing.T) {
	owner := sdk.AccAddress(crypto.AddressHash([]byte("KavaTestUser1")))

	testCases := []struct {
		msg     string
		claims  JpyxMintingClaims
		expPass bool
	}{
		{
			"valid",
			JpyxMintingClaims{
				NewJpyxMintingClaim(owner, sdk.NewCoin("bnb", sdk.OneInt()), RewardIndexes{NewRewardIndex("bnb-a", sdk.ZeroDec())}),
			},
			true,
		},
		{
			"invalid owner",
			JpyxMintingClaims{
				JpyxMintingClaim{
					BaseClaim: &BaseClaim{
						Owner: nil,
					},
				},
			},
			false,
		},
		{
			"invalid reward",
			JpyxMintingClaims{
				{
					BaseClaim: &BaseClaim{
						Owner:  jpyxtypes.StringAccAddress(owner),
						Reward: sdk.Coin{Denom: "", Amount: sdk.ZeroInt()},
					},
				},
			},
			false,
		},
		{
			"invalid collateral type",
			JpyxMintingClaims{
				{
					BaseClaim: &BaseClaim{
						Owner:  jpyxtypes.StringAccAddress(owner),
						Reward: sdk.NewCoin("bnb", sdk.OneInt()),
					},
					RewardIndexes: []RewardIndex{{"", sdk.ZeroDec()}},
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		err := tc.claims.Validate()
		if tc.expPass {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}
