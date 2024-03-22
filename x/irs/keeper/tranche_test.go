package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (suite *KeeperTestSuite) TestSetTranchePool() {
	id := uint64(1)
	pool := types.TranchePool{
		Id:               id,
		StrategyContract: "address",
		Denom:            "uatom",
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

	suite.app.IrsKeeper.SetTranchePool(suite.ctx, pool)
	pool2, found := suite.app.IrsKeeper.GetTranchePool(suite.ctx, id)
	suite.Require().True(found)

	suite.Require().Equal(pool, pool2)
}

func (suite *KeeperTestSuite) TestRemoveTranchePool() {
	id := uint64(1)
	pool := types.TranchePool{
		Id:               id,
		StrategyContract: "address",
		Denom:            "uatom",
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

	suite.app.IrsKeeper.SetTranchePool(suite.ctx, pool)
	suite.app.IrsKeeper.RemoveTranchePool(suite.ctx, pool)
	_, found := suite.app.IrsKeeper.GetTranchePool(suite.ctx, id)
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestGetTranchesByStrategy() {
	address := "address"
	pools := []types.TranchePool{
		{
			Id:               1,
			StrategyContract: address,
			Denom:            "uatom",
			StartTime:        1698796800,
			Maturity:         1572800,
			SwapFee:          sdk.ZeroDec(),
			ExitFee:          sdk.OneDec(),
			TotalShares:      sdk.NewInt64Coin("irs/tranche/1/ls", 1000000),
			PoolAssets: sdk.Coins{
				sdk.NewInt64Coin("uatom", 1000000),
				sdk.NewInt64Coin("uosmo", 1000000),
			},
		},
		{
			Id:               2,
			StrategyContract: address,
			Denom:            "uatom",
			StartTime:        1698796800,
			Maturity:         31622400,
			SwapFee:          sdk.ZeroDec(),
			ExitFee:          sdk.ZeroDec(),
			TotalShares:      sdk.NewInt64Coin("irs/tranche/2/ls", 500000),
			PoolAssets: sdk.Coins{
				sdk.NewInt64Coin("uatom", 1000000),
				sdk.NewInt64Coin("uosmo", 1000000),
			},
		},
	}

	for _, p := range pools {
		suite.app.IrsKeeper.SetTranchePool(suite.ctx, p)
	}
	pools2 := suite.app.IrsKeeper.GetTranchesByStrategy(suite.ctx, address)
	suite.Require().Equal(pools, pools2)
}

func (suite *KeeperTestSuite) TestGetAllTranchePool() {
	pools := []types.TranchePool{
		{
			Id:               1,
			StrategyContract: "address01",
			Denom:            "uatom",
			StartTime:        1698796800,
			Maturity:         1572800,
			SwapFee:          sdk.ZeroDec(),
			ExitFee:          sdk.OneDec(),
			TotalShares:      sdk.NewInt64Coin("irs/tranche/1/ls", 1000000),
			PoolAssets: sdk.Coins{
				sdk.NewInt64Coin("uatom", 1000000),
				sdk.NewInt64Coin("uosmo", 1000000),
			},
		},
		{
			Id:               2,
			StrategyContract: "address02",
			Denom:            "uatom",
			StartTime:        1698796800,
			Maturity:         31622400,
			SwapFee:          sdk.ZeroDec(),
			ExitFee:          sdk.ZeroDec(),
			TotalShares:      sdk.NewInt64Coin("irs/tranche/2/ls", 500000),
			PoolAssets: sdk.Coins{
				sdk.NewInt64Coin("uatom", 1000000),
				sdk.NewInt64Coin("uosmo", 1000000),
			},
		},
	}

	for _, p := range pools {
		suite.app.IrsKeeper.SetTranchePool(suite.ctx, p)
	}
	pools2 := suite.app.IrsKeeper.GetAllTranchePool(suite.ctx)
	suite.Require().Equal(pools, pools2)
}

func (suite *KeeperTestSuite) TestGetLastTrancheId() {
	id1 := uint64(1)
	id2 := uint64(2)
	pools := []types.TranchePool{
		{
			Id:               id1,
			StrategyContract: "address01",
			Denom:            "uatom",
			StartTime:        1698796800,
			Maturity:         1572800,
			SwapFee:          sdk.ZeroDec(),
			ExitFee:          sdk.OneDec(),
			TotalShares:      sdk.NewInt64Coin("irs/tranche/1/ls", 1000000),
			PoolAssets: sdk.Coins{
				sdk.NewInt64Coin("uatom", 1000000),
				sdk.NewInt64Coin("uosmo", 1000000),
			},
		},
		{
			Id:               id2,
			StrategyContract: "address02",
			Denom:            "uatom",
			StartTime:        1698796800,
			Maturity:         31622400,
			SwapFee:          sdk.ZeroDec(),
			ExitFee:          sdk.ZeroDec(),
			TotalShares:      sdk.NewInt64Coin("irs/tranche/2/ls", 500000),
			PoolAssets: sdk.Coins{
				sdk.NewInt64Coin("uatom", 1000000),
				sdk.NewInt64Coin("uosmo", 1000000),
			},
		},
	}

	for _, p := range pools {
		suite.app.IrsKeeper.SetTranchePool(suite.ctx, p)
	}
	id := suite.app.IrsKeeper.GetLastTrancheId(suite.ctx)
	suite.Require().Equal(id, id2)

	for _, p := range pools {
		suite.app.IrsKeeper.RemoveTranchePool(suite.ctx, p)
	}
	emptyId := suite.app.IrsKeeper.GetLastTrancheId(suite.ctx)
	suite.Require().Equal(emptyId, uint64(0))
}
