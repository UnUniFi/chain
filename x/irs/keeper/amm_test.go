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
		TotalShares:      sdk.NewInt64Coin("uatom", 1000000),
		PoolAssets: sdk.Coins{
			sdk.NewInt64Coin("uatom", 1000000),
			sdk.NewInt64Coin("uosmo", 1000000),
		},
	}

	_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tokenInMaxs)
	_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, sender, tokenInMaxs)

	_, _, err := suite.app.IrsKeeper.DepositToLiquidityPool(suite.ctx, sender, poolId, share.Amount, tokenInMaxs)
	suite.Require().Error(err)
	suite.app.IrsKeeper.SetTranchePool(suite.ctx, pool)
	_, _, err = suite.app.IrsKeeper.DepositToLiquidityPool(suite.ctx, sender, poolId, share.Amount, noTokenInMaxs)
	suite.Require().NoError(err)
	tokenIn, shareOut, err := suite.app.IrsKeeper.DepositToLiquidityPool(suite.ctx, sender, poolId, share.Amount, tokenInMaxs)
	suite.Require().NoError(err)
	suite.Require().Equal(share.Amount, shareOut)
	suite.Require().Equal(tokenInMaxs, tokenIn)
}
