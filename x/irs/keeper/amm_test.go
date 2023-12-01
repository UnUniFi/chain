package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

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

func (s *KeeperTestSuite) TestDepositToLiquidityPool2() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	strategyContract := sdk.AccAddress("strategy_contract_address")
	ut := "uatom"
	pt := types.PtDenom(types.TranchePool{Id: 1})
	lp := types.LsDenom(types.TranchePool{Id: 1})

	ptUt5k := sdk.NewCoins(sdk.NewCoin(ut, sdk.NewInt(5000)), sdk.NewCoin(pt, sdk.NewInt(5000))).Sort()
	tests := []struct {
		name            string
		txSender        sdk.AccAddress
		sharesRequested sdk.Int
		tokenInMaxs     sdk.Coins
		expectPass      bool
	}{
		{
			name:            "basic join no swap",
			txSender:        sender,
			sharesRequested: types.OneShare.MulRaw(50),
			tokenInMaxs:     sdk.Coins{},
			expectPass:      true,
		},
		{
			name:            "join no swap with zero shares requested",
			txSender:        sender,
			sharesRequested: sdk.NewInt(0),
			tokenInMaxs:     sdk.Coins{},
			expectPass:      false,
		},
		{
			name:            "join no swap with negative shares requested",
			txSender:        sender,
			sharesRequested: sdk.NewInt(-1),
			tokenInMaxs:     sdk.Coins{},
			expectPass:      false,
		},
		{
			name:            "join no swap with insufficient funds",
			txSender:        sender,
			sharesRequested: sdk.NewInt(-1),
			tokenInMaxs: sdk.Coins{
				sdk.NewCoin("bar", sdk.NewInt(4999)), sdk.NewCoin("foo", sdk.NewInt(4999)),
			},
			expectPass: false,
		},
		{
			name:            "join no swap with exact tokenInMaxs",
			txSender:        sender,
			sharesRequested: types.OneShare.MulRaw(50),
			tokenInMaxs: sdk.Coins{
				ptUt5k[0], ptUt5k[1],
			},
			expectPass: true,
		},
		{
			name:            "join no swap with arbitrary extra token in tokenInMaxs",
			txSender:        sender,
			sharesRequested: types.OneShare.MulRaw(50),
			tokenInMaxs: sdk.Coins{
				ptUt5k[0], ptUt5k[1], sdk.NewCoin("baz", sdk.NewInt(5000)),
			},
			expectPass: false,
		},
		{
			name:            "join no swap with TokenInMaxs not containing every token in pool",
			txSender:        sender,
			sharesRequested: types.OneShare.MulRaw(50),
			tokenInMaxs: sdk.Coins{
				ptUt5k[0],
			},
			expectPass: false,
		},
	}

	for _, test := range tests {
		s.SetupTest()

		ctx := s.ctx
		irsKeeper := s.app.IrsKeeper
		bankKeeper := s.app.BankKeeper

		// Create the tranche pool at first
		tranchePool := types.TranchePool{
			Id:               1,
			StrategyContract: strategyContract.String(),
			StartTime:        uint64(ctx.BlockTime().Unix()),
			Maturity:         86400 * 180,
			SwapFee:          sdk.NewDecWithPrec(3, 3), // 0.3%
			ExitFee:          sdk.ZeroDec(),
			TotalShares:      sdk.Coin{},
			PoolAssets:       sdk.Coins{},
		}
		irsKeeper.SetTranchePool(ctx, tranchePool)

		coins := sdk.Coins{sdk.NewInt64Coin(ut, 1000000000)}.Add(sdk.NewInt64Coin(pt, 1000000000))
		err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
		s.Require().NoError(err)

		err = bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, test.txSender, coins)
		s.Require().NoError(err)

		balancesBefore := bankKeeper.GetAllBalances(ctx, test.txSender)
		_, _, err = irsKeeper.DepositToLiquidityPool(ctx, test.txSender, tranchePool.Id, test.sharesRequested, test.tokenInMaxs)
		if test.expectPass {
			s.Require().NoError(err, "test: %v", test.name)
			s.Require().Equal(test.sharesRequested.String(), bankKeeper.GetBalance(ctx, test.txSender, lp).Amount.String())
			balancesAfter := bankKeeper.GetAllBalances(ctx, test.txSender)
			deltaBalances, _ := balancesBefore.SafeSub(balancesAfter...)

			s.Require().Equal("5000", deltaBalances.AmountOf(pt).String())
			s.Require().Equal("5000", deltaBalances.AmountOf(ut).String())
		} else {
			s.Require().Error(err, "test: %v", test.name)
		}
	}
}

// func (s *KeeperTestSuite) TestWithdrawFromLiquidityPool() {
// 	ptUt5k := sdk.NewCoins(sdk.NewCoin("bar", sdk.NewInt(5000)), sdk.NewCoin("foo", sdk.NewInt(5000)))
// 	tests := []struct {
// 		name         string
// 		txSender     sdk.AccAddress
// 		sharesIn     sdk.Int
// 		tokenOutMins sdk.Coins
// 		emptySender  bool
// 		expectPass   bool
// 	}{
// 		{
// 			name:         "attempt exit pool with no pool share balance",
// 			txSender:     s.TestAccs[0],
// 			sharesIn:     types.OneShare.MulRaw(50),
// 			tokenOutMins: sdk.Coins{},
// 			emptySender:  true,
// 			expectPass:   false,
// 		},
// 		{
// 			name:         "exit half pool with correct pool share balance",
// 			txSender:     s.TestAccs[0],
// 			sharesIn:     types.OneShare.MulRaw(50),
// 			tokenOutMins: sdk.Coins{},
// 			emptySender:  false,
// 			expectPass:   true,
// 		},
// 		{
// 			name:         "attempt exit pool requesting 0 share amount",
// 			txSender:     s.TestAccs[0],
// 			sharesIn:     sdk.NewInt(0),
// 			tokenOutMins: sdk.Coins{},
// 			emptySender:  false,
// 			expectPass:   false,
// 		},
// 		{
// 			name:         "attempt exit pool requesting negative share amount",
// 			txSender:     s.TestAccs[0],
// 			sharesIn:     sdk.NewInt(-1),
// 			tokenOutMins: sdk.Coins{},
// 			emptySender:  false,
// 			expectPass:   false,
// 		},
// 		{
// 			name:     "attempt exit pool with tokenOutMins above actual output",
// 			txSender: s.TestAccs[0],
// 			sharesIn: types.OneShare.MulRaw(50),
// 			tokenOutMins: sdk.Coins{
// 				sdk.NewCoin("foo", sdk.NewInt(5001)),
// 			},
// 			emptySender: false,
// 			expectPass:  false,
// 		},
// 		{
// 			name:     "attempt exit pool requesting tokenOutMins at exactly the actual output",
// 			txSender: s.TestAccs[0],
// 			sharesIn: types.OneShare.MulRaw(50),
// 			tokenOutMins: sdk.Coins{
// 				ptUt5k[1],
// 			},
// 			emptySender: false,
// 			expectPass:  true,
// 		},
// 	}

// 	for _, test := range tests {
// 		s.Run(test.name, func() {
// 			s.SetupTest()
// 			ctx := s.Ctx

// 			irsKeeper := s.App.irsKeeper
// 			bankKeeper := s.App.BankKeeper
// 			poolmanagerKeeper := s.App.PoolManagerKeeper

// 			// Mint assets to the pool creator
// 			s.FundAcc(test.txSender, defaultAcctFunds)

// 			// Create the pool at first
// 			msg := balancer.NewMsgCreateBalancerPool(test.txSender, balancer.PoolParams{
// 				SwapFee: sdk.NewDecWithPrec(1, 2),
// 				ExitFee: sdk.NewDec(0),
// 			}, defaultPoolAssets, defaultFutureGovernor)
// 			poolId, err := poolmanagerKeeper.CreatePool(ctx, msg)
// 			s.Require().NoError(err)

// 			// If we are testing insufficient pool share balances, switch tx sender from pool creator to empty account
// 			if test.emptySender {
// 				test.txSender = sender
// 			}

// 			balancesBefore := bankKeeper.GetAllBalances(s.Ctx, test.txSender)
// 			_, err = irsKeeper.WithdrawFromLiquidityPool(ctx, test.txSender, poolId, test.sharesIn, test.tokenOutMins)

// 			if test.expectPass {
// 				s.Require().NoError(err, "test: %v", test.name)
// 				s.Require().Equal(test.sharesIn.String(), bankKeeper.GetBalance(s.Ctx, test.txSender, "gamm/pool/1").Amount.String())
// 				balancesAfter := bankKeeper.GetAllBalances(s.Ctx, test.txSender)
// 				deltaBalances, _ := balancesBefore.SafeSub(balancesAfter)
// 				// The pool was created with the 10000foo, 10000bar, and the pool share was minted as 100*OneShare gamm/pool/1.
// 				// Thus, to refund the 50*OneShare gamm/pool/1, (10000foo, 10000bar) * (1 / 2) balances should be refunded.
// 				s.Require().Equal("-5000", deltaBalances.AmountOf("foo").String())
// 				s.Require().Equal("-5000", deltaBalances.AmountOf("bar").String())

// 				liquidity, err := irsKeeper.GetTotalLiquidity(ctx)
// 				s.Require().NoError(err)
// 				s.Require().Equal("5000bar,5000foo", liquidity.String())
// 			} else {
// 				s.Require().Error(err, "test: %v", test.name)
// 			}
// 		})
// 	}
// }
