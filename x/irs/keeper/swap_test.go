package keeper_test

// sdk "github.com/cosmos/cosmos-sdk/types"

// SwapUtToPt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, tokenIn sdk.Coin) error {
// SwapPtToUt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, tokenIn sdk.Coin) error {
// UpdatePoolForSwap(
// 	ctx sdk.Context,
// 	pool types.TranchePool,
// 	sender sdk.AccAddress,
// 	tokenIn sdk.Coin,
// 	tokenOut sdk.Coin,
// ) error

// func (s *KeeperTestSuite) TestSwapExactAmountIn() {
// 	type param struct {
// 		tokenIn           sdk.Coin
// 		tokenOutDenom     string
// 		tokenOutMinAmount sdk.Int
// 		expectedTokenOut  sdk.Int
// 	}

// 	tests := []struct {
// 		name  string
// 		param param
// 		swapFee sdk.Dec
// 		expectPass                    bool
// 	}{
// 		{
// 			name: "Proper swap",
// 			param: param{
// 				tokenIn:           sdk.NewCoin("foo", sdk.NewInt(100000)),
// 				tokenOutDenom:     "bar",
// 				tokenOutMinAmount: sdk.NewInt(1),
// 				expectedTokenOut:  sdk.NewInt(49262),
// 			},
// 			expectPass: true,
// 		},
// 		{
// 			name: "Proper swap2",
// 			param: param{
// 				tokenIn:           sdk.NewCoin("bar", sdk.NewInt(2451783)),
// 				tokenOutDenom:     "baz",
// 				tokenOutMinAmount: sdk.NewInt(1),
// 				expectedTokenOut:  sdk.NewInt(1167843),
// 			},
// 			expectPass: true,
// 		},
// 		{
// 			name: "boundary valid spread factor given (= 0.5 pool spread factor)",
// 			param: param{
// 				tokenIn:           sdk.NewCoin("foo", sdk.NewInt(100000)),
// 				tokenOutDenom:     "bar",
// 				tokenOutMinAmount: sdk.NewInt(1),
// 				expectedTokenOut:  sdk.NewInt(46833),
// 			},
// 			swapFee:         sdk.MustNewDecFromStr("0.1"),
// 			expectPass:                    true,
// 		},
// 		{
// 			name: "invalid spread factor given (< 0.5 pool spread factor)",
// 			param: param{
// 				tokenIn:           sdk.NewCoin("foo", sdk.NewInt(100000)),
// 				tokenOutDenom:     "bar",
// 				tokenOutMinAmount: sdk.NewInt(1),
// 				expectedTokenOut:  sdk.NewInt(49262),
// 			},
// 			swapFee:         sdk.MustNewDecFromStr("0.1"),
// 			expectPass:                    false,
// 		},
// 		{
// 			name: "out is lesser than min amount",
// 			param: param{
// 				tokenIn:           sdk.NewCoin("bar", sdk.NewInt(2451783)),
// 				tokenOutDenom:     "baz",
// 				tokenOutMinAmount: sdk.NewInt(9000000),
// 			},
// 			expectPass: false,
// 		},
// 		{
// 			name: "in and out denom are same",
// 			param: param{
// 				tokenIn:           sdk.NewCoin("bar", sdk.NewInt(2451783)),
// 				tokenOutDenom:     "bar",
// 				tokenOutMinAmount: sdk.NewInt(1),
// 			},
// 			expectPass: false,
// 		},
// 		{
// 			name: "unknown in denom",
// 			param: param{
// 				tokenIn:           sdk.NewCoin("bara", sdk.NewInt(2451783)),
// 				tokenOutDenom:     "bar",
// 				tokenOutMinAmount: sdk.NewInt(1),
// 			},
// 			expectPass: false,
// 		},
// 		{
// 			name: "unknown out denom",
// 			param: param{
// 				tokenIn:           sdk.NewCoin("bar", sdk.NewInt(2451783)),
// 				tokenOutDenom:     "bara",
// 				tokenOutMinAmount: sdk.NewInt(1),
// 			},
// 			expectPass: false,
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		s.Run(test.name, func() {
// 			// Init suite for each test.
// 			s.SetupTest()
// 			swapFee := sdk.ZeroDec()
// 			if !test.swapFee.IsNil() {
// 				swapFee = test.swapFee
// 			}
// 			poolId := s.PrepareBalancerPoolWithPoolParams(balancer.PoolParams{
// 				SwapFee: swapFee,
// 				ExitFee: sdk.ZeroDec(),
// 			})
// 			keeper := s.App.IrsKeeper
// 			ctx := s.ctx
// 			pool, err := s.App.IrsKeeper.GetTranchePool(ctx, poolId)
// 			s.NoError(err)

// 			if test.expectPass {
// 				spotPriceBefore, err := keeper.CalculateSpotPrice(ctx, poolId, test.param.tokenIn.Denom, test.param.tokenOutDenom)
// 				s.NoError(err, "test: %v", test.name)

// 				prevGasConsumed := s.ctx.GasMeter().GasConsumed()
// 				tokenOutAmount, err := keeper.SwapExactAmountIn(ctx, s.TestAccs[0], pool, test.param.tokenIn, test.param.tokenOutDenom, test.param.tokenOutMinAmount, spreadFactor)
// 				s.NoError(err, "test: %v", test.name)
// 				s.Require().Equal(test.param.expectedTokenOut.String(), tokenOutAmount.String())
// 				gasConsumedForSwap := s.ctx.GasMeter().GasConsumed() - prevGasConsumed
// 				// We consume `types.GasFeeForSwap` directly, so the extra I/O operation mean we end up consuming more.
// 				s.Assert().Greater(gasConsumedForSwap, uint64(types.BalancerGasFeeForSwap))

// 				s.AssertEventEmitted(ctx, types.TypeEvtTokenSwapped, 1)

// 				spotPriceAfter, err := keeper.CalculateSpotPrice(ctx, poolId, test.param.tokenIn.Denom, test.param.tokenOutDenom)
// 				s.NoError(err, "test: %v", test.name)

// 				if !test.swapFee.IsNil() {
// 					return
// 				}

// 				// Ratio of the token out should be between the before spot price and after spot price.
// 				tradeAvgPrice := test.param.tokenIn.Amount.ToDec().Quo(tokenOutAmount.ToDec())
// 				s.True(tradeAvgPrice.GT(spotPriceBefore) && tradeAvgPrice.LT(spotPriceAfter), "test: %v", test.name)
// 			} else {
// 				_, err := keeper.SwapExactAmountIn(ctx, s.TestAccs[0], pool, test.param.tokenIn, test.param.tokenOutDenom, test.param.tokenOutMinAmount, spreadFactor)
// 				s.Error(err, "test: %v", test.name)
// 			}
// 		})
// 	}
// }
