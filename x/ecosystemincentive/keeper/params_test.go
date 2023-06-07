package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
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
				RewardParams: types.DefaultParams().RewardParams,
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
								RewardType: types.RewardType_FRONTEND_DEVELOPERS,
								Rate:       sdk.MustNewDecFromStr("-0.5"),
							},
						},
					},
				},
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
								RewardType: types.RewardType_FRONTEND_DEVELOPERS,
								Rate:       sdk.MustNewDecFromStr("10"),
							},
						},
					},
				},
			},
			expErr: true,
		},
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
