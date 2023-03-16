package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/testutil/nullify"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

func (suite *KeeperTestSuite) TestStrategyQuerySingle() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	wctx := sdk.WrapSDKContext(ctx)
	vaultDenom := "uatom"
	msgs := createNStrategy(&keeper, ctx, vaultDenom, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetStrategyRequest
		response *types.QueryGetStrategyResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetStrategyRequest{Denom: vaultDenom, Id: msgs[0].Id},
			response: &types.QueryGetStrategyResponse{Strategy: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetStrategyRequest{Denom: vaultDenom, Id: msgs[1].Id},
			response: &types.QueryGetStrategyResponse{Strategy: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetStrategyRequest{Denom: vaultDenom, Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		suite.Run(tc.desc, func() {
			response, err := keeper.Strategy(wctx, tc.request)
			if tc.err != nil {
				suite.Require().ErrorIs(err, tc.err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestStrategyQueryPaginated() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	wctx := sdk.WrapSDKContext(ctx)
	vaultDenom := "uatom"
	msgs := createNStrategy(&keeper, ctx, vaultDenom, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllStrategyRequest {
		return &types.QueryAllStrategyRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	suite.Run("ByOffset", func() {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StrategyAll(wctx, request(nil, uint64(i), uint64(step), false))
			suite.Require().NoError(err)
			suite.Require().LessOrEqual(len(resp.Strategies), step)
			suite.Require().Subset(
				nullify.Fill(msgs),
				nullify.Fill(resp.Strategies),
			)
		}
	})
	suite.Run("ByKey", func() {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StrategyAll(wctx, request(next, 0, uint64(step), false))
			suite.Require().NoError(err)
			suite.Require().LessOrEqual(len(resp.Strategies), step)
			suite.Require().Subset(
				nullify.Fill(msgs),
				nullify.Fill(resp.Strategies),
			)
			next = resp.Pagination.NextKey
		}
	})
	suite.Run("Total", func() {
		resp, err := keeper.StrategyAll(wctx, request(nil, 0, 0, true))
		suite.Require().NoError(err)

		// TODO
		// require.Equal(t, len(msgs), int(resp.Pagination.Total))
		// require.ElementsMatch(t,
		// 	nullify.Fill(msgs),
		// 	nullify.Fill(resp.Strategies),
		// )
		println(resp)
		//
	})
	suite.Run("InvalidRequest", func() {
		_, err := keeper.StrategyAll(wctx, nil)
		suite.Require().ErrorIs(err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
