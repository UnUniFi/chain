package keeper_test

import (
	"fmt"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/irs/keeper"
	"github.com/UnUniFi/chain/x/irs/types"
)

func (suite *KeeperTestSuite) TestDepositToLiquidityPool__NotAvailableTranchePool() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	poolId := uint64(1)
	share := sdk.NewInt64Coin("uatom", 200000)
	tokenInMaxs := sdk.Coins{sdk.NewInt64Coin("uatom", 300000), sdk.NewInt64Coin("uosmo", 300000)}

	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tokenInMaxs)
	suite.Require().NoError(err)

	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, sender, tokenInMaxs)
	suite.Require().NoError(err)

	_, _, err = suite.app.IrsKeeper.DepositToLiquidityPool(suite.ctx, sender, poolId, share.Amount, tokenInMaxs)
	suite.Require().Error(err)
}

func (s *KeeperTestSuite) TestDepositToLiquidityPool() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	strategyContract := sdk.AccAddress("strategy_contract_address")
	ut := "uatom"
	pt := types.PtDenom(types.TranchePool{Id: 1})
	lp := types.LsDenom(types.TranchePool{Id: 1})

	ptUt100 := sdk.NewCoins(sdk.NewCoin(ut, sdk.NewInt(100)), sdk.NewCoin(pt, sdk.NewInt(100))).Sort()
	ptUt5k := sdk.NewCoins(sdk.NewCoin(ut, sdk.NewInt(5000)), sdk.NewCoin(pt, sdk.NewInt(5000))).Sort()
	tests := []struct {
		name            string
		txSender        sdk.AccAddress
		sharesRequested sdk.Int
		existingShares  sdk.Int
		poolTokens      sdk.Coins
		tokenInMaxs     sdk.Coins
		expectPass      bool
	}{
		{
			name:            "basic add liquidity",
			txSender:        sender,
			sharesRequested: types.OneShare.MulRaw(50),
			existingShares:  types.OneShare,
			poolTokens:      ptUt100,
			tokenInMaxs:     sdk.Coins{},
			expectPass:      true,
		},
		{
			name:            "add liquidity with zero shares requested",
			txSender:        sender,
			sharesRequested: sdk.NewInt(0),
			existingShares:  types.OneShare,
			poolTokens:      ptUt100,
			tokenInMaxs:     sdk.Coins{},
			expectPass:      false,
		},
		{
			name:            "add liquidity with negative shares requested",
			txSender:        sender,
			sharesRequested: sdk.NewInt(-1),
			existingShares:  types.OneShare,
			poolTokens:      ptUt100,
			tokenInMaxs:     sdk.Coins{},
			expectPass:      false,
		},
		{
			name:            "add liquidity with insufficient funds",
			txSender:        sender,
			sharesRequested: sdk.NewInt(-1),
			existingShares:  types.OneShare,
			poolTokens:      ptUt100,
			tokenInMaxs: sdk.Coins{
				sdk.NewCoin(pt, sdk.NewInt(4999)), sdk.NewCoin(ut, sdk.NewInt(4999)),
			},
			expectPass: false,
		},
		{
			name:            "add liquidity with exact tokenInMaxs",
			txSender:        sender,
			sharesRequested: types.OneShare.MulRaw(50),
			existingShares:  types.OneShare,
			poolTokens:      ptUt100,
			tokenInMaxs: sdk.Coins{
				ptUt5k[0], ptUt5k[1],
			},
			expectPass: true,
		},
		{
			name:            "add liquidity with arbitrary extra token in tokenInMaxs",
			txSender:        sender,
			sharesRequested: types.OneShare.MulRaw(50),
			existingShares:  types.OneShare,
			poolTokens:      ptUt100,
			tokenInMaxs: sdk.Coins{
				ptUt5k[0], ptUt5k[1], sdk.NewCoin("baz", sdk.NewInt(5000)),
			},
			expectPass: false,
		},
		{
			name:            "add liquidity with TokenInMaxs not containing every token in pool",
			txSender:        sender,
			sharesRequested: types.OneShare.MulRaw(50),
			existingShares:  types.OneShare,
			poolTokens:      ptUt100,
			tokenInMaxs: sdk.Coins{
				ptUt5k[0],
			},
			expectPass: false,
		},
		{
			name:            "adding liquidity to empty pool",
			txSender:        sender,
			sharesRequested: types.OneShare,
			existingShares:  sdk.ZeroInt(),
			poolTokens:      sdk.Coins{},
			tokenInMaxs:     ptUt5k,
			expectPass:      true,
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
			TotalShares:      sdk.NewCoin(lp, test.existingShares),
			PoolAssets:       test.poolTokens,
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

func (suite *KeeperTestSuite) TestGetMaximalNoSwapLPAmount() {
	strategyContract := sdk.AccAddress("strategy_contract_address")
	ut := "uatom"
	pt := types.PtDenom(types.TranchePool{Id: 1})
	lp := types.LsDenom(types.TranchePool{Id: 1})
	ptUt1000000 := sdk.NewCoins(sdk.NewCoin(ut, sdk.NewInt(1000000)), sdk.NewCoin(pt, sdk.NewInt(1000000))).Sort()

	tests := map[string]struct {
		poolAssets          sdk.Coins
		existingShares      sdk.Int
		shareOutAmount      sdk.Int
		expectedLpLiquidity sdk.Coins
		err                 error
	}{
		"Share ratio is zero": {
			poolAssets:     ptUt1000000,
			existingShares: sdk.NewInt(10_000_000),
			shareOutAmount: sdk.ZeroInt(),
			err:            types.ErrInvalidMathApprox,
		},

		"Share ratio is negative": {
			poolAssets:     ptUt1000000,
			existingShares: sdk.NewInt(10_000_000),
			shareOutAmount: sdk.NewInt(-1),
			err:            types.ErrInvalidMathApprox,
		},

		"Pass": {
			poolAssets:     ptUt1000000,
			existingShares: sdk.NewInt(10_000_000),

			// totalShare:   100_000_000_000_000_000_000
			// shareOutAmount: 8_000_000_000_000_000_000
			// shareRatio = shareOutAmount/totalShare = 0.08
			// Amount of tokens in poolAssets:
			// 		- defaultPoolAssets[1].Token.Amount: 10000
			//  	- defaultPoolAssets[0].Token.Amount: 10000
			shareOutAmount: sdk.NewInt(80_000),
			expectedLpLiquidity: sdk.Coins{
				sdk.NewInt64Coin(pt, 8_000), sdk.NewInt64Coin(ut, 8_000),
			},
		},

		"Pass with ceiling result": {
			poolAssets:     ptUt1000000,
			existingShares: sdk.NewInt(1_000_000),

			// totalShare:   100_000_000_000_000_000_000
			// shareOutAmount: 8_888_000_000_000_000_000
			// shareRatio = shareOutAmount/totalShare = 0.08888
			// Amount of tokens in poolAssets:
			// 		- defaultPoolAssets[1].Token.Amount: 10000
			//  	- defaultPoolAssets[0].Token.Amount: 10000
			shareOutAmount: sdk.NewInt(8_888),
			expectedLpLiquidity: sdk.Coins{
				sdk.NewInt64Coin(pt, 8_888), sdk.NewInt64Coin(ut, 8_888),
			},
		},

		"Error with empty pool": {
			poolAssets:     sdk.Coins{},
			existingShares: sdk.ZeroInt(),
			err:            types.ErrInvalidMathApprox,
		},
	}

	for name, tc := range tests {
		suite.Run(name, func() {
			suite.SetupTest()

			ctx := suite.ctx
			irsKeeper := suite.app.IrsKeeper

			// Create the tranche pool at first
			tranchePool := types.TranchePool{
				Id:               1,
				StrategyContract: strategyContract.String(),
				StartTime:        uint64(ctx.BlockTime().Unix()),
				Maturity:         86400 * 180,
				SwapFee:          sdk.NewDecWithPrec(3, 3), // 0.3%
				ExitFee:          sdk.ZeroDec(),
				TotalShares:      sdk.NewCoin(lp, tc.existingShares),
				PoolAssets:       tc.poolAssets,
			}
			irsKeeper.SetTranchePool(ctx, tranchePool)

			neededLpLiquidity, err := keeper.GetMaximalNoSwapLPAmount(ctx, tranchePool, tc.shareOutAmount)
			if tc.err != nil {
				suite.Require().Error(err)
				msgError := fmt.Sprintf("Too few shares out wanted. (debug: getMaximalNoSwapLPAmount share ratio is zero or negative): %s", tc.err)
				suite.Require().EqualError(err, msgError)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(neededLpLiquidity, tc.expectedLpLiquidity)
			}
		})
	}
}

func (s *KeeperTestSuite) TestWithdrawFromLiquidityPool() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	strategyContract := sdk.AccAddress("strategy_contract_address")
	ut := "uatom"
	pt := types.PtDenom(types.TranchePool{Id: 1})
	lp := types.LsDenom(types.TranchePool{Id: 1})

	ptUt5k := sdk.NewCoins(sdk.NewCoin(ut, sdk.NewInt(5000)), sdk.NewCoin(pt, sdk.NewInt(5000))).Sort()
	tests := []struct {
		name         string
		txSender     sdk.AccAddress
		sharesIn     sdk.Int
		tokenOutMins sdk.Coins
		expTokenOut  sdk.Coins
		expectPass   bool
	}{
		{
			name:         "exit half pool with correct pool share balance",
			txSender:     sender,
			sharesIn:     types.OneShare.Quo(sdk.NewInt(2)),
			tokenOutMins: sdk.Coins{},
			expTokenOut:  sdk.NewCoins(sdk.NewCoin(ut, sdk.NewInt(2500)), sdk.NewCoin(pt, sdk.NewInt(2500))).Sort(),
			expectPass:   true,
		},
		{
			name:         "attempt exit pool requesting 0 share amount",
			txSender:     sender,
			sharesIn:     sdk.NewInt(0),
			tokenOutMins: sdk.Coins{},
			expTokenOut:  sdk.Coins{},
			expectPass:   false,
		},
		{
			name:         "attempt exit pool requesting negative share amount",
			txSender:     sender,
			sharesIn:     sdk.NewInt(-1),
			tokenOutMins: sdk.Coins{},
			expTokenOut:  sdk.Coins{},
			expectPass:   false,
		},
		{
			name:     "attempt exit pool with tokenOutMins above actual output",
			txSender: sender,
			sharesIn: types.OneShare,
			tokenOutMins: sdk.Coins{
				sdk.NewCoin(ut, sdk.NewInt(5001)),
			},
			expTokenOut: sdk.Coins{},
			expectPass:  false,
		},
		{
			name:     "attempt exit pool requesting tokenOutMins at exactly the actual output",
			txSender: sender,
			sharesIn: types.OneShare,
			tokenOutMins: sdk.Coins{
				ptUt5k[1],
			},
			expTokenOut: ptUt5k,
			expectPass:  true,
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			s.SetupTest()
			ctx := s.ctx

			irsKeeper := s.app.IrsKeeper
			bankKeeper := s.app.BankKeeper

			sharesRequested := types.OneShare
			existingShares := sdk.ZeroInt()
			poolTokens := sdk.Coins{}

			// Create the tranche pool at first
			tranchePool := types.TranchePool{
				Id:               1,
				StrategyContract: strategyContract.String(),
				StartTime:        uint64(ctx.BlockTime().Unix()),
				Maturity:         86400 * 180,
				SwapFee:          sdk.NewDecWithPrec(3, 3), // 0.3%
				ExitFee:          sdk.ZeroDec(),
				TotalShares:      sdk.NewCoin(lp, existingShares),
				PoolAssets:       poolTokens,
			}
			irsKeeper.SetTranchePool(ctx, tranchePool)

			coins := sdk.Coins{sdk.NewInt64Coin(ut, 1000000000)}.Add(sdk.NewInt64Coin(pt, 1000000000))
			err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
			s.Require().NoError(err)

			err = bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, test.txSender, coins)
			s.Require().NoError(err)

			_, _, err = irsKeeper.DepositToLiquidityPool(ctx, test.txSender, tranchePool.Id, sharesRequested, ptUt5k)
			s.Require().NoError(err)

			balancesBefore := bankKeeper.GetAllBalances(s.ctx, test.txSender)
			_, err = irsKeeper.WithdrawFromLiquidityPool(ctx, test.txSender, tranchePool.Id, test.sharesIn, test.tokenOutMins)
			if test.expectPass {
				s.Require().NoError(err, "test: %v", test.name)
				balancesAfter := bankKeeper.GetAllBalances(s.ctx, test.txSender)
				deltaBalances, _ := balancesAfter.SafeSub(balancesBefore...)
				s.Require().Equal(test.expTokenOut.AmountOf(pt).String(), deltaBalances.AmountOf(pt).String())
				s.Require().Equal(test.expTokenOut.AmountOf(ut).String(), deltaBalances.AmountOf(ut).String())
			} else {
				s.Require().Error(err, "test: %v", test.name)
			}
		})
	}
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
