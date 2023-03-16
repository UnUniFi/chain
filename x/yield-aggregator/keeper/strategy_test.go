package keeper_test

import (
	"github.com/UnUniFi/chain/testutil/nullify"
	"github.com/UnUniFi/chain/x/yield-aggregator/keeper"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createNStrategy(keeper *keeper.Keeper, ctx sdk.Context, denom string, n int) []types.Strategy {
	items := make([]types.Strategy, n)
	for i := range items {
		items[i] = types.Strategy{
			Denom:           denom,
			ContractAddress: "",
			Name:            "",
		}
		items[i].Id = keeper.AppendStrategy(ctx, denom, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestStrategyGet() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNStrategy(&keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		got, found := keeper.GetStrategy(ctx, vaultDenom, item.Id)
		suite.Require().True(found)
		suite.Require().Equal(
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func (suite *KeeperTestSuite) TestStrategyRemove() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNStrategy(&keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		keeper.RemoveStrategy(ctx, vaultDenom, item.Id)
		_, found := keeper.GetStrategy(ctx, vaultDenom, item.Id)
		suite.Require().False(found)
	}
}

func (suite *KeeperTestSuite) TestStrategyGetAll() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNStrategy(&keeper, ctx, vaultDenom, 10)
	suite.Require().ElementsMatch(
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStrategy(ctx, vaultDenom)),
	)
}

func (suite *KeeperTestSuite) TestStrategyCount() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNStrategy(&keeper, ctx, vaultDenom, 10)
	count := uint64(len(items))
	suite.Require().Equal(count, keeper.GetStrategyCount(ctx, vaultDenom))
}

// TODO: add test for StakeToStrategy
// TODO: add test for UnstakeFromStrategy
// TODO: add test for GetAmountFromStrategy
// TODO: add test for GetUnbondingAmountFromStrategy
