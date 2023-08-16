package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/testutil/nullify"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (suite *KeeperTestSuite) TestVaultQuerySingle() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	wctx := sdk.WrapSDKContext(ctx)
	vaultDenom := "uatom"
	msgs := createNVault(&keeper, ctx, vaultDenom, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVaultRequest
		response *types.QueryGetVaultResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetVaultRequest{Id: msgs[0].Id},
			response: &types.QueryGetVaultResponse{Vault: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetVaultRequest{Id: msgs[1].Id},
			response: &types.QueryGetVaultResponse{Vault: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetVaultRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		suite.Run(tc.desc, func() {
			response, err := keeper.Vault(wctx, tc.request)
			if tc.err != nil {
				suite.Require().ErrorIs(err, tc.err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(
					nullify.Fill(tc.response.Vault),
					nullify.Fill(response.Vault),
				)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestVaultQueryPaginated() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	wctx := sdk.WrapSDKContext(ctx)
	vaultDenom := "uatom"
	msgs := createNVault(&keeper, ctx, vaultDenom, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVaultRequest {
		return &types.QueryAllVaultRequest{
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
			resp, err := keeper.VaultAll(wctx, request(nil, uint64(i), uint64(step), false))
			suite.Require().NoError(err)
			suite.Require().LessOrEqual(len(resp.Vaults), step)
			vaults := []types.Vault{}
			for _, vault := range resp.Vaults {
				vaults = append(vaults, vault.Vault)
			}
			suite.Require().Subset(
				nullify.Fill(msgs),
				nullify.Fill(vaults),
			)
		}
	})
	suite.Run("ByKey", func() {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VaultAll(wctx, request(next, 0, uint64(step), false))
			suite.Require().NoError(err)
			suite.Require().LessOrEqual(len(resp.Vaults), step)
			vaults := []types.Vault{}
			for _, vault := range resp.Vaults {
				vaults = append(vaults, vault.Vault)
			}
			suite.Require().Subset(
				nullify.Fill(msgs),
				nullify.Fill(vaults),
			)
			next = resp.Pagination.NextKey
		}
	})
	suite.Run("Total", func() {
		resp, err := keeper.VaultAll(wctx, request(nil, 0, 0, true))
		suite.Require().NoError(err)

		// TODO
		// suite.Require().Equal(len(msgs), int(resp.Pagination.Total))
		// suite.Require().ElementsMatch(t,
		// 	nullify.Fill(msgs),
		// 	nullify.Fill(resp.Strategies),
		// )
		println(resp)
		//
	})
	suite.Run("InvalidRequest", func() {
		_, err := keeper.VaultAll(wctx, nil)
		suite.Require().ErrorIs(err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
