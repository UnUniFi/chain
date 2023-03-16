package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/UnUniFi/chain/testutil/nullify"
	"github.com/UnUniFi/chain/x/yield-aggregator/keeper"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

func createNVault(keeper *keeper.Keeper, ctx sdk.Context, denom string, n int) []types.Vault {
	items := make([]types.Vault, n)
	for i := range items {
		addr := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
		items[i] = types.Vault{
			Denom:                  denom,
			WithdrawCommissionRate: sdk.MustNewDecFromStr("0.001"),
			WithdrawReserveRate:    sdk.MustNewDecFromStr("0.001"),
			Owner:                  addr.String(),
			OwnerDeposit:           sdk.NewInt64Coin("uguu", 1000_000),
			StrategyWeights: []types.StrategyWeight{
				{
					StrategyId: 1,
					Weight:     sdk.OneDec(),
				},
			},
		}
		items[i].Id = keeper.AppendVault(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestVaultGet() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNVault(&keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		got, found := keeper.GetVault(ctx, item.Id)
		suite.Require().True(found)
		suite.Require().Equal(
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func (suite *KeeperTestSuite) TestVaultRemove() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNVault(&keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		keeper.RemoveVault(ctx, item.Id)
		_, found := keeper.GetVault(ctx, item.Id)
		suite.Require().False(found)
	}
}

func (suite *KeeperTestSuite) TestVaultGetAll() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNVault(&keeper, ctx, vaultDenom, 10)
	suite.Require().ElementsMatch(
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVault(ctx)),
	)
}

func (suite *KeeperTestSuite) TestVaultCount() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNVault(&keeper, ctx, vaultDenom, 10)
	count := uint64(len(items))
	suite.Require().Equal(count, keeper.GetVaultCount(ctx))
}
