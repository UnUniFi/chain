package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (suite *KeeperTestSuite) TestDepositToLiquidityPool() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	poolId := uint64(1)
	share := sdk.NewInt64Coin("uatom", 200000)
	noTokenInMaxs := sdk.Coins{}
	tokenInMaxs := sdk.Coins{sdk.NewInt64Coin("uatom", 300000), sdk.NewInt64Coin("uosmo", 300000)}
	pool := types.TranchePool{
		Id:               poolId,
		StrategyContract: "address",
		StartTime:        1698796800,
		Maturity:         1572800,
		SwapFee:          sdk.ZeroDec(),
		ExitFee:          sdk.OneDec(),
		PoolAssets: sdk.Coins{
			sdk.NewInt64Coin("uatom", 1000000),
			sdk.NewInt64Coin("uosmo", 1000000),
		},
	}

	_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tokenInMaxs)
	_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, sender, tokenInMaxs)

	// no pool
	_, _, err := suite.app.IrsKeeper.DepositToLiquidityPool(suite.ctx, sender, poolId, share.Amount, tokenInMaxs)
	suite.Require().Error(err)
	suite.app.IrsKeeper.SetTranchePool(suite.ctx, pool)
	// todo: no tokenInMaxs
	_, _, err = suite.app.IrsKeeper.DepositToLiquidityPool(suite.ctx, sender, poolId, share.Amount, noTokenInMaxs)
	suite.Require().NoError(err)

	_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tokenInMaxs)
	_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, sender, tokenInMaxs)

	// share < tokenInMaxs
	// tokenInMaxs < share
}

func (suite *KeeperTestSuite) TestMintPoolShareToAccount() {
	address := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	poolId := uint64(1)
	shareAmount := sdk.NewInt(200000)
	pool := types.TranchePool{
		Id:               poolId,
		StrategyContract: "address",
		StartTime:        1698796800,
		Maturity:         1572800,
		SwapFee:          sdk.ZeroDec(),
		ExitFee:          sdk.OneDec(),
		PoolAssets: sdk.Coins{
			sdk.NewInt64Coin("uatom", 1000000),
			sdk.NewInt64Coin("uosmo", 1000000),
		},
	}
	pool.TotalShares = sdk.NewInt64Coin(types.LsDenom(pool), 1000000)

	err := suite.app.IrsKeeper.MintPoolShareToAccount(suite.ctx, pool, address, shareAmount)
	suite.Require().NoError(err)
	balances := suite.app.BankKeeper.GetAllBalances(suite.ctx, address)
	suite.Require().Equal(shareAmount, balances[0].Amount)
}

func (suite *KeeperTestSuite) TestBurnPoolShareFromAccount() {
	address := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	poolId := uint64(1)
	shareAmount := sdk.NewInt(200000)
	burnAmount := sdk.NewInt(100000)
	pool := types.TranchePool{
		Id:               poolId,
		StrategyContract: "address",
		StartTime:        1698796800,
		Maturity:         1572800,
		SwapFee:          sdk.ZeroDec(),
		ExitFee:          sdk.OneDec(),
		PoolAssets: sdk.Coins{
			sdk.NewInt64Coin("uatom", 1000000),
			sdk.NewInt64Coin("uosmo", 1000000),
		},
	}

	suite.app.IrsKeeper.SetTranchePool(suite.ctx, pool)
	err := suite.app.IrsKeeper.MintPoolShareToAccount(suite.ctx, pool, address, shareAmount)
	suite.Require().NoError(err)
	err = suite.app.IrsKeeper.BurnPoolShareFromAccount(suite.ctx, pool, address, burnAmount)
	suite.Require().NoError(err)
	balances := suite.app.BankKeeper.GetAllBalances(suite.ctx, address)
	suite.Require().Equal(shareAmount.Sub(burnAmount), balances[0].Amount)
}
