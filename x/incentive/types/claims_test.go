package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"

	ununifitypes "github.com/UnUniFi/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestClaimsValidate(t *testing.T) {
	owner := sdk.AccAddress(crypto.AddressHash([]byte("KavaTestUser1")))

	testCases := []struct {
		msg     string
		claims  CdpMintingClaims
		expPass bool
	}{
		{
			"valid",
			CdpMintingClaims{
				NewCdpMintingClaim(owner, sdk.NewCoin("bnb", sdk.OneInt()), RewardIndexes{NewRewardIndex("bnb-a", sdk.ZeroDec())}),
			},
			true,
		},
		{
			"invalid owner",
			CdpMintingClaims{
				CdpMintingClaim{
					BaseClaim: &BaseClaim{
						Owner: nil,
					},
				},
			},
			false,
		},
		{
			"invalid reward",
			CdpMintingClaims{
				{
					BaseClaim: &BaseClaim{
						Owner:  ununifitypes.StringAccAddress(owner),
						Reward: sdk.Coin{Denom: "", Amount: sdk.ZeroInt()},
					},
				},
			},
			false,
		},
		{
			"invalid collateral type",
			CdpMintingClaims{
				{
					BaseClaim: &BaseClaim{
						Owner:  ununifitypes.StringAccAddress(owner),
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
