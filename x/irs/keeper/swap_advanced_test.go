package keeper_test

// Unable to create test for SwapUtToYt because it depends on the strategy contract
func (suite *KeeperTestSuite) TestSwapUtToYt() {
	// sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// strategyContract := sdk.AccAddress("strategy_contract_address")
	// ut := "uatom"
	// pt := types.PtDenom(types.TranchePool{Id: 1})
	// yt := types.YtDenom(types.TranchePool{Id: 1})
	// lp := types.LsDenom(types.TranchePool{Id: 1})

	// ptUt5k := sdk.NewCoins(sdk.NewCoin(ut, sdk.NewInt(5000)), sdk.NewCoin(pt, sdk.NewInt(5000))).Sort()

	// type param struct {
	// 	tokenIn          sdk.Coin
	// 	tokenOutDenom    string
	// 	requiredYtAmount sdk.Int
	// }

	// tests := []struct {
	// 	name        string
	// 	param       param
	// 	swapFee     sdk.Dec
	// 	expectPass  bool
	// 	expectedErr error
	// }{
	// 	{
	// 		name: "Too few TokenIn",
	// 		param: param{
	// 			tokenIn:          sdk.NewCoin(ut, sdk.NewInt(1)),
	// 			tokenOutDenom:    yt,
	// 			requiredYtAmount: sdk.NewInt(10_000),
	// 		},
	// 		swapFee:     sdk.ZeroDec(),
	// 		expectPass:  false,
	// 		expectedErr: types.ErrInsufficientFunds,
	// 	},
	// 	{
	// 		name: "Enough TokenIn",
	// 		param: param{
	// 			tokenIn:          sdk.NewCoin(ut, sdk.NewInt(100_000)),
	// 			tokenOutDenom:    yt,
	// 			requiredYtAmount: sdk.NewInt(10_000),
	// 		},
	// 		swapFee:    sdk.ZeroDec(),
	// 		expectPass: true,
	// 	},
	// 	{
	// 		name: "Enough TokenIn with SwapFee",
	// 		param: param{
	// 			tokenIn:          sdk.NewCoin(ut, sdk.NewInt(100_000)),
	// 			tokenOutDenom:    yt,
	// 			requiredYtAmount: sdk.NewInt(10_000),
	// 		},
	// 		swapFee:    sdk.MustNewDecFromStr("0.1"),
	// 		expectPass: true,
	// 	},
	// }

	// // setup environment
	// ctx := suite.ctx
	// bankKeeper := suite.app.BankKeeper
	// coins := sdk.Coins{sdk.NewInt64Coin(ut, 1_000_000)}
	// err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	// suite.Require().NoError(err)
	// err = bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sender, coins)
	// suite.Require().NoError(err)

	// existingShares := sdk.OneInt()
	// tranchePool := types.TranchePool{
	// 	Id:               1,
	// 	StrategyContract: strategyContract.String(),
	// 	StartTime:        uint64(ctx.BlockTime().Unix() - 86400*30), // 30 days ago
	// 	Maturity:         86400 * 180,                               // 180 days maturity
	// 	SwapFee:          sdk.NewDecWithPrec(3, 3),                  // 0.3%
	// 	ExitFee:          sdk.ZeroDec(),
	// 	TotalShares:      sdk.NewCoin(lp, existingShares),
	// 	PoolAssets:       ptUt5k,
	// }
	// irsKeeper := suite.app.IrsKeeper
	// irsKeeper.SetTranchePool(ctx, tranchePool)
	// poolAddr := types.GetVaultModuleAddress(tranchePool)
	// err = bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	// suite.Require().NoError(err)
	// err = bankKeeper.SendCoins(ctx, sender, poolAddr, coins)
	// suite.Require().NoError(err)

	// for _, tc := range tests {
	// 	err = irsKeeper.SwapUtToYt(ctx, sender, tranchePool, tc.param.requiredYtAmount, tc.param.tokenIn)
	// 	if tc.expectPass {
	// 		suite.Require().NoError(err)
	// 	} else {
	// 		suite.Require().EqualError(err, tc.expectedErr.Error())
	// 	}
	// }
}
