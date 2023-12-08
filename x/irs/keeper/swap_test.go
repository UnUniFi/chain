package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

// SwapUtToPt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, tokenIn sdk.Coin) error {
// SwapPtToUt(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, tokenIn sdk.Coin) error {
// UpdatePoolForSwap(
// 	ctx sdk.Context,
// 	pool types.TranchePool,
// 	sender sdk.AccAddress,
// 	tokenIn sdk.Coin,
// 	tokenOut sdk.Coin,
// ) error

func (s *KeeperTestSuite) TestSwapExactAmountIn() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	strategyContract := sdk.AccAddress("strategy_contract_address")
	ut := "uatom"
	pt := types.PtDenom(types.TranchePool{Id: 1})
	lp := types.LsDenom(types.TranchePool{Id: 1})

	ptUt5k := sdk.NewCoins(sdk.NewCoin(ut, sdk.NewInt(5000)), sdk.NewCoin(pt, sdk.NewInt(5000))).Sort()

	type param struct {
		tokenIn           sdk.Coin
		tokenOutDenom     string
		tokenOutMinAmount sdk.Int
		expectedTokenOut  sdk.Int
	}

	tests := []struct {
		name       string
		param      param
		swapFee    sdk.Dec
		expectPass bool
	}{
		{
			name: "Proper swap",
			param: param{
				tokenIn:           sdk.NewCoin(pt, sdk.NewInt(100_000)),
				tokenOutDenom:     ut,
				tokenOutMinAmount: sdk.NewInt(1),
				expectedTokenOut:  sdk.NewInt(4992),
			},
			swapFee:    sdk.ZeroDec(),
			expectPass: true,
		},
		{
			name: "Proper swap2",
			param: param{
				tokenIn:           sdk.NewCoin(ut, sdk.NewInt(2_000)),
				tokenOutDenom:     pt,
				tokenOutMinAmount: sdk.NewInt(1),
				expectedTokenOut:  sdk.NewInt(1499),
			},
			swapFee:    sdk.ZeroDec(),
			expectPass: true,
		},
		{
			name: "too much token in",
			param: param{
				tokenIn:           sdk.NewCoin(ut, sdk.NewInt(999_999)),
				tokenOutDenom:     pt,
				tokenOutMinAmount: sdk.NewInt(1),
				expectedTokenOut:  sdk.NewInt(1),
			},
			swapFee:    sdk.ZeroDec(),
			expectPass: false,
		},
		{
			name: "Proper swap with swap fee",
			param: param{
				tokenIn:           sdk.NewCoin(pt, sdk.NewInt(100_000)),
				tokenOutDenom:     ut,
				tokenOutMinAmount: sdk.NewInt(1),
				expectedTokenOut:  sdk.NewInt(4987),
			},
			swapFee:    sdk.MustNewDecFromStr("0.1"),
			expectPass: true,
		},
		{
			name: "out is lesser than min amount",
			param: param{
				tokenIn:           sdk.NewCoin(pt, sdk.NewInt(2451783)),
				tokenOutDenom:     ut,
				tokenOutMinAmount: sdk.NewInt(9000000),
			},
			swapFee:    sdk.ZeroDec(),
			expectPass: false,
		},
		{
			name: "in and out denom are same",
			param: param{
				tokenIn:           sdk.NewCoin(ut, sdk.NewInt(100_000)),
				tokenOutDenom:     ut,
				tokenOutMinAmount: sdk.NewInt(1),
			},
			swapFee:    sdk.ZeroDec(),
			expectPass: false,
		},
		{
			name: "unknown in denom",
			param: param{
				tokenIn:           sdk.NewCoin("unknown", sdk.NewInt(2451783)),
				tokenOutDenom:     ut,
				tokenOutMinAmount: sdk.NewInt(1),
			},
			swapFee:    sdk.ZeroDec(),
			expectPass: false,
		},
		{
			name: "unknown out denom",
			param: param{
				tokenIn:           sdk.NewCoin(pt, sdk.NewInt(2451783)),
				tokenOutDenom:     "unknown",
				tokenOutMinAmount: sdk.NewInt(1),
			},
			swapFee:    sdk.ZeroDec(),
			expectPass: false,
		},
	}

	for _, test := range tests {
		test := test
		s.Run(test.name, func() {
			// Init suite for each test.
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
				StartTime:        uint64(ctx.BlockTime().Unix() - 86400*30), // 30 days ago
				Maturity:         86400 * 180,                               // 180 days maturity
				SwapFee:          sdk.NewDecWithPrec(3, 3),                  // 0.3%
				ExitFee:          sdk.ZeroDec(),
				TotalShares:      sdk.NewCoin(lp, existingShares),
				PoolAssets:       poolTokens,
			}
			irsKeeper.SetTranchePool(ctx, tranchePool)

			coins := sdk.Coins{sdk.NewInt64Coin(ut, 1000000000)}.Add(sdk.NewInt64Coin(pt, 1000000000))
			err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
			s.Require().NoError(err)

			err = bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sender, coins)
			s.Require().NoError(err)

			_, _, err = irsKeeper.DepositToLiquidityPool(ctx, sender, tranchePool.Id, sharesRequested, ptUt5k)
			s.Require().NoError(err)

			tranchePool, found := irsKeeper.GetTranchePool(ctx, tranchePool.Id)
			s.Require().True(found)

			if test.expectPass {
				tokenOutAmount, err := irsKeeper.SwapExactAmountIn(ctx, sender, tranchePool, test.param.tokenIn, test.param.tokenOutDenom, test.param.tokenOutMinAmount, test.swapFee)
				s.NoError(err, "test: %v", test.name)
				s.Require().Equal(test.param.expectedTokenOut.String(), tokenOutAmount.String())

				if !test.swapFee.IsNil() {
					return
				}
			} else {
				_, err := irsKeeper.SwapExactAmountIn(ctx, sender, tranchePool, test.param.tokenIn, test.param.tokenOutDenom, test.param.tokenOutMinAmount, test.swapFee)
				s.Error(err, "test: %v", test.name)
			}
		})
	}
}
