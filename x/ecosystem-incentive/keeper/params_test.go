package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

func (suite *KeeperTestSuite) TestParams() {

	testCases := []struct {
		name   string
		input  types.Params
		expErr bool
	}{
		{
			name: "ordinal success case",
			input: types.Params{
				RewardParams:          types.DefaultParams().RewardParams,
				MaxIncentiveUnitIdLen: types.DefaultMaxIncentiveUnitIdLen,
			},
			expErr: false,
		},
		{
			name: "negative reward rate",
			input: types.Params{
				RewardParams: []*types.RewardParams{
					{
						ModuleName: "nftmarket",
						RewardRate: []types.RewardRate{
							{
								RewardType: types.RewardType_NFTMARKET_FRONTEND,
								Rate:       sdk.MustNewDecFromStr("-0.5"),
							},
						},
					},
				},
				MaxIncentiveUnitIdLen: types.DefaultMaxIncentiveUnitIdLen,
			},
			expErr: true,
		},
		{
			name: "invalid reward rate which is greater than 1",
			input: types.Params{
				RewardParams: []*types.RewardParams{
					{
						ModuleName: "nftmarket",
						RewardRate: []types.RewardRate{
							{
								RewardType: types.RewardType_NFTMARKET_FRONTEND,
								Rate:       sdk.MustNewDecFromStr("10"),
							},
						},
					},
				},
				MaxIncentiveUnitIdLen: types.DefaultMaxIncentiveUnitIdLen,
			},
			expErr: true,
		},
		// {
		// 	name:   "negative base proposer reward",
		// 	input:  types.Params{},
		// 	expErr: true,
		// },
		// {
		// 	name:   "bonus proposer reward > 1",
		// 	input:  types.Params{},
		// 	expErr: true,
		// },
		// {
		// 	name:   "negative bonus proposer reward",
		// 	input:  types.Params{},
		// 	expErr: true,
		// },
	}

	for _, tc := range testCases {
		suite.SetupTest()
		expected := suite.app.EcosystemincentiveKeeper.GetParams(suite.ctx)

		if tc.expErr {
			suite.Panics(func() {
				suite.app.EcosystemincentiveKeeper.SetParams(suite.ctx, tc.input)
			})
		} else {
			suite.app.EcosystemincentiveKeeper.SetParams(suite.ctx, tc.input)
			expected = suite.app.EcosystemincentiveKeeper.GetParams(suite.ctx)
		}

		params := suite.app.EcosystemincentiveKeeper.GetParams(suite.ctx)
		suite.Require().Equal(expected, params)
	}
}
