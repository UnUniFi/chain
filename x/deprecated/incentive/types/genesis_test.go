package types

import (
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"

	ununifitypes "github.com/UnUniFi/chain/types"
)

func TestGenesisStateValidate(t *testing.T) {
	type args struct {
		params      Params
		genAccTimes GenesisAccumulationTimes
		claims      CdpMintingClaims
		denoms      *GenesisDenoms
	}

	type errArgs struct {
		expectPass bool
		contains   string
	}

	testCases := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			name: "default",
			args: args{
				params:      DefaultParams(),
				genAccTimes: DefaultGenesisAccumulationTimes,
				claims:      DefaultCdpClaims,
				denoms:      DefaultGenesisDenoms(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "valid",
			args: args{
				params: NewParams(
					RewardPeriods{
						NewRewardPeriod(
							true,
							"bnb-a",
							time.Date(2020, 10, 15, 14, 0, 0, 0, time.UTC),
							time.Date(2024, 10, 15, 14, 0, 0, 0, time.UTC),
							sdk.NewCoin("uguu", sdk.NewInt(25000)),
						),
					},
					Multipliers{
						NewMultiplier(Small, 1, sdk.MustNewDecFromStr("0.33")),
					},
					time.Date(2025, 10, 15, 14, 0, 0, 0, time.UTC),
				),
				genAccTimes: GenesisAccumulationTimes{GenesisAccumulationTime{
					CollateralType:           "bnb-a",
					PreviousAccumulationTime: time.Date(2020, 10, 15, 14, 0, 0, 0, time.UTC),
					//RewardFactor:             sdk.ZeroDec(),
				}},
				claims: CdpMintingClaims{
					{
						BaseClaim: &BaseClaim{
							Owner:  ununifitypes.StringAccAddress(sdk.AccAddress(crypto.AddressHash([]byte("KavaTestUser1")))),
							Reward: sdk.NewCoin("uguu", sdk.NewInt(100000000)),
						},
						RewardIndexes: []RewardIndex{
							{
								CollateralType: "bnb-a",
								RewardFactor:   sdk.ZeroDec(),
							},
						},
					},
				},
				denoms: DefaultGenesisDenoms(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		// {
		// 	name: "invalid genesis accumulation time",
		// 	args: args{
		// 		params: DefaultParams(),
		// 		genAccTimes: GenesisAccumulationTimes{
		// 			{
		// 				CollateralType: "btcb-a",
		// 				RewardFactor:   sdk.MustNewDecFromStr("-0.1"),
		// 			},
		// 		},
		// 		claims: DefaultCdpClaims,
		// 	},
		// 	errArgs: errArgs{
		// 		expectPass: false,
		// 		contains:   "reward factor should be ≥ 0.0",
		// 	},
		// },
		{
			name: "invalid claim",
			args: args{
				params:      DefaultParams(),
				genAccTimes: DefaultGenesisAccumulationTimes,
				claims: CdpMintingClaims{
					{
						BaseClaim: &BaseClaim{
							Owner:  ununifitypes.StringAccAddress(sdk.AccAddress{}),
							Reward: sdk.NewCoin("uguu", sdk.NewInt(100000000)),
						},
						RewardIndexes: []RewardIndex{
							{
								CollateralType: "bnb-a",
								RewardFactor:   sdk.ZeroDec(),
							},
						},
					},
				},
				denoms: DefaultGenesisDenoms(),
			},
			errArgs: errArgs{
				expectPass: false,
				contains:   "claim owner cannot be empty",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gs := NewGenesisState(tc.args.params, tc.args.genAccTimes, tc.args.claims, tc.args.denoms)
			err := gs.Validate()
			if tc.errArgs.expectPass {
				require.NoError(t, err, tc.name)
			} else {
				require.Error(t, err, tc.name)
				require.True(t, strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}
