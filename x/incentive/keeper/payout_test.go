package keeper_test

import (
	"errors"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/lcnem/jpyx/app"
	cdptypes "github.com/lcnem/jpyx/x/cdp/types"

	"github.com/lcnem/jpyx/x/incentive/types"
	jsmndisttypes "github.com/lcnem/jpyx/x/jsmndist/types"
)

func (suite *KeeperTestSuite) TestPayoutJpyxMintingClaim() {
	type args struct {
		ctype                    string
		rewardsPerSecond         sdk.Coin
		initialTime              time.Time
		initialCollateral        sdk.Coin
		initialPrincipal         sdk.Coin
		multipliers              types.Multipliers
		multiplier               types.MultiplierName
		timeElapsed              int
		expectedBalance          sdk.Coins
		expectedPeriods          vestingtypes.Periods
		isPeriodicVestingAccount bool
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}
	type test struct {
		name    string
		args    args
		errArgs errArgs
	}
	testCases := []test{
		{
			"valid 1 day",
			args{
				ctype:                    "bnb-a",
				rewardsPerSecond:         c("ukava", 122354),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				initialCollateral:        c("bnb", 1000000000000),
				initialPrincipal:         c("usdx", 10000000000),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              86400,
				expectedBalance:          cs(c("usdx", 10000000000), c("ukava", 10576385600)),
				expectedPeriods:          vestingtypes.Periods{vestingtypes.Period{Length: 32918400, Amount: cs(c("ukava", 10571385600))}},
				isPeriodicVestingAccount: true,
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"invalid zero rewards",
			args{
				ctype:                    "bnb-a",
				rewardsPerSecond:         c("ukava", 0),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				initialCollateral:        c("bnb", 1000000000000),
				initialPrincipal:         c("usdx", 10000000000),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              86400,
				expectedBalance:          cs(c("usdx", 10000000000)),
				expectedPeriods:          vestingtypes.Periods{},
				isPeriodicVestingAccount: false,
			},
			errArgs{
				expectPass: false,
				contains:   "claim amount rounds to zero",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupWithGenState()
			suite.ctx = suite.ctx.WithBlockTime(tc.args.initialTime)

			// setup incentive state
			params := types.NewParams(
				types.RewardPeriods{types.NewRewardPeriod(true, tc.args.ctype, tc.args.initialTime, tc.args.initialTime.Add(time.Hour*24*365*4), tc.args.rewardsPerSecond)},
				types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, tc.args.ctype, tc.args.initialTime, tc.args.initialTime.Add(time.Hour*24*365*4), cs(tc.args.rewardsPerSecond))},
				types.MultiRewardPeriods{types.NewMultiRewardPeriod(true, tc.args.ctype, tc.args.initialTime, tc.args.initialTime.Add(time.Hour*24*365*4), cs(tc.args.rewardsPerSecond))},
				types.RewardPeriods{types.NewRewardPeriod(true, tc.args.ctype, tc.args.initialTime, tc.args.initialTime.Add(time.Hour*24*365*4), tc.args.rewardsPerSecond)},
				tc.args.multipliers,
				tc.args.initialTime.Add(time.Hour*24*365*5),
			)
			suite.keeper.SetParams(suite.ctx, params)
			suite.keeper.SetPreviousJpyxMintingAccrualTime(suite.ctx, tc.args.ctype, tc.args.initialTime)
			suite.keeper.SetJpyxMintingRewardFactor(suite.ctx, tc.args.ctype, sdk.ZeroDec())

			// setup account state
			sk := suite.app.GetBankKeeper()
			err := sk.MintCoins(suite.ctx, cdptypes.ModuleName, sdk.NewCoins(tc.args.initialCollateral))
			suite.Require().NoError(err)
			err = sk.SendCoinsFromModuleToAccount(suite.ctx, cdptypes.ModuleName, suite.addrs[0], sdk.NewCoins(tc.args.initialCollateral))
			suite.Require().NoError(err)

			// setup kavadist state
			err = sk.MintCoins(suite.ctx, jsmndisttypes.ModuleName, cs(c("ukava", 1000000000000)))
			suite.Require().NoError(err)

			// setup cdp state
			cdpKeeper := suite.app.GetCDPKeeper()
			err = cdpKeeper.AddCdp(suite.ctx, suite.addrs[0], tc.args.initialCollateral, tc.args.initialPrincipal, tc.args.ctype)
			suite.Require().NoError(err)

			claim, found := suite.keeper.GetJpyxMintingClaim(suite.ctx, suite.addrs[0])
			suite.Require().True(found)
			suite.Require().Equal(sdk.ZeroDec(), claim.RewardIndexes[0].RewardFactor)

			updatedBlockTime := suite.ctx.BlockTime().Add(time.Duration(int(time.Second) * tc.args.timeElapsed))
			suite.ctx = suite.ctx.WithBlockTime(updatedBlockTime)
			rewardPeriod, found := suite.keeper.GetJpyxMintingRewardPeriod(suite.ctx, tc.args.ctype)
			suite.Require().True(found)
			err = suite.keeper.AccumulateJpyxMintingRewards(suite.ctx, rewardPeriod)
			suite.Require().NoError(err)

			err = suite.keeper.ClaimJpyxMintingReward(suite.ctx, suite.addrs[0], string(tc.args.multiplier))

			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
				ak := suite.app.GetAccountKeeper()
				acc := ak.GetAccount(suite.ctx, suite.addrs[0])
				bk := suite.app.GetBankKeeper()
				suite.Require().Equal(tc.args.expectedBalance, bk.GetAllBalances(suite.ctx, acc.GetAddress()))

				if tc.args.isPeriodicVestingAccount {
					vacc, ok := acc.(*vestingtypes.PeriodicVestingAccount)
					suite.Require().True(ok)
					suite.Require().Equal(tc.args.expectedPeriods, vacc.VestingPeriods)
				}

				claim, found := suite.keeper.GetJpyxMintingClaim(suite.ctx, suite.addrs[0])
				suite.Require().True(found)
				suite.Require().Equal(c("ukava", 0), claim.Reward)
			} else {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestPayoutHardLiquidityProviderClaim() {
	type args struct {
		deposit                  sdk.Coins
		borrow                   sdk.Coins
		rewardsPerSecond         sdk.Coins
		initialTime              time.Time
		multipliers              types.Multipliers
		multiplier               types.MultiplierName
		timeElapsed              int64
		expectedRewards          sdk.Coins
		expectedPeriods          vestingtypes.Periods
		isPeriodicVestingAccount bool
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}
	type test struct {
		name    string
		args    args
		errArgs errArgs
	}
	testCases := []test{
		{
			"single reward denom: valid 1 day",
			args{
				deposit:                  cs(c("bnb", 10000000000)),
				borrow:                   cs(c("bnb", 5000000000)),
				rewardsPerSecond:         cs(c("hard", 122354)),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              86400,
				expectedRewards:          cs(c("hard", 21142771200)), // 10571385600 (deposit reward) + 10571385600 (borrow reward)
				expectedPeriods:          vestingtypes.Periods{vestingtypes.Period{Length: 32918400, Amount: cs(c("hard", 21142771200))}},
				isPeriodicVestingAccount: true,
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"single reward denom: valid 10 days",
			args{
				deposit:                  cs(c("bnb", 10000000000)),
				borrow:                   cs(c("bnb", 5000000000)),
				rewardsPerSecond:         cs(c("hard", 122354)),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              864000,
				expectedRewards:          cs(c("hard", 211427712000)), // 105713856000 (deposit reward) + 105713856000 (borrow reward)
				expectedPeriods:          vestingtypes.Periods{vestingtypes.Period{Length: 32140800, Amount: cs(c("hard", 211427712000))}},
				isPeriodicVestingAccount: true,
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"invalid zero rewards",
			args{
				deposit:                  cs(c("bnb", 10000000000)),
				borrow:                   cs(c("bnb", 5000000000)),
				rewardsPerSecond:         cs(c("hard", 0)),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              86400,
				expectedRewards:          cs(c("hard", 0)),
				expectedPeriods:          vestingtypes.Periods{},
				isPeriodicVestingAccount: false,
			},
			errArgs{
				expectPass: false,
				contains:   "claim amount rounds to zero",
			},
		},
		{
			"multiple reward denoms: valid 1 day",
			args{
				deposit:                  cs(c("bnb", 10000000000)),
				borrow:                   cs(c("bnb", 5000000000)),
				rewardsPerSecond:         cs(c("hard", 122354), c("ukava", 122354)),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              86400,
				expectedRewards:          cs(c("hard", 21142771200), c("ukava", 21142771200)), // 10571385600 (deposit reward) + 10571385600 (borrow reward)
				expectedPeriods:          vestingtypes.Periods{vestingtypes.Period{Length: 32918400, Amount: cs(c("hard", 21142771200), c("ukava", 21142771200))}},
				isPeriodicVestingAccount: true,
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"multiple reward denoms: valid 10 days",
			args{
				deposit:                  cs(c("bnb", 10000000000)),
				borrow:                   cs(c("bnb", 5000000000)),
				rewardsPerSecond:         cs(c("hard", 122354), c("ukava", 122354)),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              864000,
				expectedRewards:          cs(c("hard", 211427712000), c("ukava", 211427712000)), // 105713856000 (deposit reward) + 105713856000 (borrow reward)
				expectedPeriods:          vestingtypes.Periods{vestingtypes.Period{Length: 32140800, Amount: cs(c("hard", 211427712000), c("ukava", 211427712000))}},
				isPeriodicVestingAccount: true,
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"multiple reward denoms with different rewards per second: valid 1 day",
			args{
				deposit:                  cs(c("bnb", 10000000000)),
				borrow:                   cs(c("bnb", 5000000000)),
				rewardsPerSecond:         cs(c("hard", 122354), c("ukava", 222222)),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              86400,
				expectedRewards:          cs(c("hard", 21142771200), c("ukava", 38399961600)),
				expectedPeriods:          vestingtypes.Periods{vestingtypes.Period{Length: 32918400, Amount: cs(c("hard", 21142771200), c("ukava", 38399961600))}},
				isPeriodicVestingAccount: true,
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupWithGenState()
			suite.ctx = suite.ctx.WithBlockTime(tc.args.initialTime)

			// setup kavadist state
			sk := suite.app.GetBankKeeper()
			err := sk.MintCoins(suite.ctx, jsmndisttypes.ModuleName, cs(c("hard", 1000000000000000000), c("ukava", 1000000000000000000)))
			suite.Require().NoError(err)

			// Set up generic reward periods
			var multiRewardPeriods types.MultiRewardPeriods
			var rewardPeriods types.RewardPeriods
			for _, coin := range tc.args.deposit {
				if len(tc.args.rewardsPerSecond) > 0 {
					rewardPeriod := types.NewRewardPeriod(true, coin.Denom, tc.args.initialTime, tc.args.initialTime.Add(time.Hour*24*365*4), tc.args.rewardsPerSecond[0])
					rewardPeriods = append(rewardPeriods, rewardPeriod)
				}
				multiRewardPeriod := types.NewMultiRewardPeriod(true, coin.Denom, tc.args.initialTime, tc.args.initialTime.Add(time.Hour*24*365*4), tc.args.rewardsPerSecond)
				multiRewardPeriods = append(multiRewardPeriods, multiRewardPeriod)
			}

			// Set up generic reward periods
			params := types.NewParams(
				rewardPeriods, multiRewardPeriods, multiRewardPeriods, rewardPeriods,
				types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				tc.args.initialTime.Add(time.Hour*24*365*5),
			)
			suite.keeper.SetParams(suite.ctx, params)

			/*
				// Set each denom's previous accrual time and supply reward factor
				if len(tc.args.rewardsPerSecond) > 0 {
					for _, coin := range tc.args.deposit {
						suite.keeper.SetPreviousHardSupplyRewardAccrualTime(suite.ctx, coin.Denom, tc.args.initialTime)
						var rewardIndexes types.RewardIndexes
						for _, rewardCoin := range tc.args.rewardsPerSecond {
							rewardIndex := types.NewRewardIndex(rewardCoin.Denom, sdk.ZeroDec())
							rewardIndexes = append(rewardIndexes, rewardIndex)
						}
						suite.keeper.SetHardSupplyRewardIndexes(suite.ctx, coin.Denom, rewardIndexes)
					}
				}
			*/

			/*
				// Set each denom's previous accrual time and borrow reward factor
				if len(tc.args.rewardsPerSecond) > 0 {
					for _, coin := range tc.args.borrow {
						suite.keeper.SetPreviousHardBorrowRewardAccrualTime(suite.ctx, coin.Denom, tc.args.initialTime)
						var rewardIndexes types.RewardIndexes
						for _, rewardCoin := range tc.args.rewardsPerSecond {
							rewardIndex := types.NewRewardIndex(rewardCoin.Denom, sdk.ZeroDec())
							rewardIndexes = append(rewardIndexes, rewardIndex)
						}
						suite.keeper.SetHardBorrowRewardIndexes(suite.ctx, coin.Denom, rewardIndexes)
					}
				}

				hardKeeper := suite.app.GetHardKeeper()
				userAddr := suite.addrs[3]

				// User deposits and borrows
				err = hardKeeper.Deposit(suite.ctx, userAddr, tc.args.deposit)
				suite.Require().NoError(err)
				err = hardKeeper.Borrow(suite.ctx, userAddr, tc.args.borrow)
				suite.Require().NoError(err)

				// Check that Hard hooks initialized a HardLiquidityProviderClaim that has 0 rewards
				claim, found := suite.keeper.GetHardLiquidityProviderClaim(suite.ctx, suite.addrs[3])
				suite.Require().True(found)
				for _, coin := range tc.args.deposit {
					suite.Require().Equal(sdk.ZeroInt(), claim.Reward.AmountOf(coin.Denom))
				}
			*/

			// Set up future runtime context
			runAtTime := time.Unix(suite.ctx.BlockTime().Unix()+(tc.args.timeElapsed), 0)
			runCtx := suite.ctx.WithBlockTime(runAtTime)

			/*
				// Run Hard begin blocker
				hard.BeginBlocker(runCtx, suite.hardKeeper)

				// Accumulate supply rewards for each deposit denom
				for _, coin := range tc.args.deposit {
					rewardPeriod, found := suite.keeper.GetHardSupplyRewardPeriods(runCtx, coin.Denom)
					suite.Require().True(found)
					err = suite.keeper.AccumulateHardSupplyRewards(runCtx, rewardPeriod)
					suite.Require().NoError(err)
				}

				// Accumulate borrow rewards for each deposit denom
				for _, coin := range tc.args.borrow {
					rewardPeriod, found := suite.keeper.GetHardBorrowRewardPeriods(runCtx, coin.Denom)
					suite.Require().True(found)
					err = suite.keeper.AccumulateHardBorrowRewards(runCtx, rewardPeriod)
					suite.Require().NoError(err)
				}

				// Sync hard supply rewards
				deposit, found := suite.hardKeeper.GetDeposit(suite.ctx, suite.addrs[3])
				suite.Require().True(found)
				suite.keeper.SynchronizeHardSupplyReward(suite.ctx, deposit)

				// Sync hard borrow rewards
				borrow, found := suite.hardKeeper.GetBorrow(suite.ctx, suite.addrs[3])
				suite.Require().True(found)
				suite.keeper.SynchronizeHardBorrowReward(suite.ctx, borrow)
			*/

			// Fetch pre-claim balances
			ak := suite.app.GetAccountKeeper()
			preClaimAcc := ak.GetAccount(runCtx, suite.addrs[3])

			// err = suite.keeper.ClaimHardReward(runCtx, suite.addrs[3], tc.args.multiplier)
			if tc.errArgs.expectPass {
				suite.Require().NoError(err)

				// Check that user's balance has increased by expected reward amount
				postClaimAcc := ak.GetAccount(suite.ctx, suite.addrs[3])
				suite.Require().Equal(sk.GetAllBalances(suite.ctx, preClaimAcc.GetAddress()).Add(tc.args.expectedRewards...), sk.GetAllBalances(suite.ctx, postClaimAcc.GetAddress()))

				if tc.args.isPeriodicVestingAccount {
					vacc, ok := postClaimAcc.(*vestingtypes.PeriodicVestingAccount)
					suite.Require().True(ok)
					suite.Require().Equal(tc.args.expectedPeriods, vacc.VestingPeriods)
				}

				/*
					// Check that each claim reward coin's amount has been reset to 0
					claim, found := suite.keeper.GetHardLiquidityProviderClaim(runCtx, suite.addrs[3])
					suite.Require().True(found)
					for _, claimRewardCoin := range claim.Reward {
						suite.Require().Equal(c(claimRewardCoin.Denom, 0), claimRewardCoin)
					}
				*/
			} else {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestSendCoinsToPeriodicVestingAccount() {
	type accountArgs struct {
		periods          vestingtypes.Periods
		origVestingCoins sdk.Coins
		startTime        int64
		endTime          int64
	}
	type args struct {
		accArgs             accountArgs
		period              vestingtypes.Period
		ctxTime             time.Time
		mintModAccountCoins bool
		expectedPeriods     vestingtypes.Periods
		expectedStartTime   int64
		expectedEndTime     int64
	}
	type errArgs struct {
		expectErr bool
		contains  string
	}
	type testCase struct {
		name    string
		args    args
		errArgs errArgs
	}
	type testCases []testCase

	tests := testCases{
		{
			name: "insert period at beginning schedule",
			args: args{
				accArgs: accountArgs{
					periods: vestingtypes.Periods{
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
					origVestingCoins: cs(c("ukava", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 2, Amount: cs(c("ukava", 6))},
				ctxTime:             time.Unix(101, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 3, Amount: cs(c("ukava", 6))},
					vestingtypes.Period{Length: 2, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
				expectedStartTime: 100,
				expectedEndTime:   120,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "insert period at beginning with new start time",
			args: args{
				accArgs: accountArgs{
					periods: vestingtypes.Periods{
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
					origVestingCoins: cs(c("ukava", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("ukava", 6))},
				ctxTime:             time.Unix(80, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 7, Amount: cs(c("ukava", 6))},
					vestingtypes.Period{Length: 18, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
				expectedStartTime: 80,
				expectedEndTime:   120,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "insert period in middle of schedule",
			args: args{
				accArgs: accountArgs{
					periods: vestingtypes.Periods{
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
					origVestingCoins: cs(c("ukava", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("ukava", 6))},
				ctxTime:             time.Unix(101, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 3, Amount: cs(c("ukava", 6))},
					vestingtypes.Period{Length: 2, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
				expectedStartTime: 100,
				expectedEndTime:   120,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "append to end of schedule",
			args: args{
				accArgs: accountArgs{
					periods: vestingtypes.Periods{
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
					origVestingCoins: cs(c("ukava", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("ukava", 6))},
				ctxTime:             time.Unix(125, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 12, Amount: cs(c("ukava", 6))}},
				expectedStartTime: 100,
				expectedEndTime:   132,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "add coins to existing period",
			args: args{
				accArgs: accountArgs{
					periods: vestingtypes.Periods{
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
					origVestingCoins: cs(c("ukava", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 6))},
				ctxTime:             time.Unix(110, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 11))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
				expectedStartTime: 100,
				expectedEndTime:   120,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
		{
			name: "insufficient mod account balance",
			args: args{
				accArgs: accountArgs{
					periods: vestingtypes.Periods{
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
					origVestingCoins: cs(c("ukava", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("ukava", 6))},
				ctxTime:             time.Unix(125, 0),
				mintModAccountCoins: false,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 12, Amount: cs(c("ukava", 6))}},
				expectedStartTime: 100,
				expectedEndTime:   132,
			},
			errArgs: errArgs{
				expectErr: true,
				contains:  "insufficient funds",
			},
		},
		{
			name: "add large period mid schedule",
			args: args{
				accArgs: accountArgs{
					periods: vestingtypes.Periods{
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))}},
					origVestingCoins: cs(c("ukava", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 50, Amount: cs(c("ukava", 6))},
				ctxTime:             time.Unix(110, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("ukava", 5))},
					vestingtypes.Period{Length: 40, Amount: cs(c("ukava", 6))}},
				expectedStartTime: 100,
				expectedEndTime:   160,
			},
			errArgs: errArgs{
				expectErr: false,
				contains:  "",
			},
		},
	}
	for _, tc := range tests {
		suite.Run(tc.name, func() {
			// create the periodic vesting account
			pva, err := suite.createPeriodicVestingAccount(tc.args.accArgs.origVestingCoins, tc.args.accArgs.periods, tc.args.accArgs.startTime, tc.args.accArgs.endTime)
			suite.Require().NoError(err)

			// setup store state with account and kavadist module account
			suite.ctx = suite.ctx.WithBlockTime(tc.args.ctxTime)
			ak := suite.app.GetAccountKeeper()
			ak.SetAccount(suite.ctx, pva)
			// mint module account coins if required
			if tc.args.mintModAccountCoins {
				sk := suite.app.GetBankKeeper()
				err = sk.MintCoins(suite.ctx, jsmndisttypes.ModuleName, tc.args.period.Amount)
				suite.Require().NoError(err)
			}

			err = suite.keeper.SendTimeLockedCoinsToPeriodicVestingAccount(suite.ctx, jsmndisttypes.ModuleName, pva.GetAddress(), tc.args.period.Amount, tc.args.period.Length)
			if tc.errArgs.expectErr {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			} else {
				suite.Require().NoError(err)

				acc := suite.getAccount(pva.GetAddress())
				vacc, ok := acc.(*vestingtypes.PeriodicVestingAccount)
				suite.Require().True(ok)
				suite.Require().Equal(tc.args.expectedPeriods, vacc.VestingPeriods)
				suite.Require().Equal(tc.args.expectedStartTime, vacc.StartTime)
				suite.Require().Equal(tc.args.expectedEndTime, vacc.EndTime)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestSendCoinsToBaseAccount() {
	suite.SetupWithAccountState()
	// send coins to base account
	err := suite.keeper.SendTimeLockedCoinsToAccount(suite.ctx, jsmndisttypes.ModuleName, suite.addrs[1], cs(c("ukava", 100)), 5)
	suite.Require().NoError(err)
	acc := suite.getAccount(suite.addrs[1])
	vacc, ok := acc.(*vestingtypes.PeriodicVestingAccount)
	suite.True(ok)
	expectedPeriods := vestingtypes.Periods{
		vestingtypes.Period{Length: int64(5), Amount: cs(c("ukava", 100))},
	}
	bk := suite.app.GetBankKeeper()
	suite.Equal(expectedPeriods, vacc.VestingPeriods)
	suite.Equal(cs(c("ukava", 100)), vacc.OriginalVesting)
	suite.Equal(cs(c("ukava", 500)), bk.GetAllBalances(suite.ctx, vacc.GetAddress()))
	suite.Equal(int64(105), vacc.EndTime)
	suite.Equal(int64(100), vacc.StartTime)

}

func (suite *KeeperTestSuite) TestSendCoinsToInvalidAccount() {
	suite.SetupWithAccountState()
	err := suite.keeper.SendTimeLockedCoinsToAccount(suite.ctx, jsmndisttypes.ModuleName, suite.addrs[2], cs(c("ukava", 100)), 5)
	suite.Require().True(errors.Is(err, types.ErrInvalidAccountType))
	macc := suite.getModuleAccount(cdptypes.ModuleName)
	err = suite.keeper.SendTimeLockedCoinsToAccount(suite.ctx, jsmndisttypes.ModuleName, macc.GetAddress(), cs(c("ukava", 100)), 5)
	suite.Require().True(errors.Is(err, types.ErrInvalidAccountType))
}

func (suite *KeeperTestSuite) SetupWithAccountState() {
	// creates a new app state with 4 funded addresses and 1 module account
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: time.Unix(100, 0)})
	_, addrs := app.GeneratePrivKeyAddressPairs(4)
	authGS := app.NewAuthGenState(
		tApp, 
		addrs,
		[]sdk.Coins{
			cs(c("ukava", 400)),
			cs(c("ukava", 400)),
			cs(c("ukava", 400)),
			cs(c("ukava", 400)),
		})
	tApp.InitializeFromGenesisStates(
		authGS,
	)
	ak := tApp.GetAccountKeeper()
	bk := tApp.GetBankKeeper()
	macc := ak.GetModuleAccount(ctx, jsmndisttypes.ModuleName)
	err := bk.MintCoins(ctx, macc.GetName(), cs(c("ukava", 600)))
	suite.Require().NoError(err)

	// sets addrs[0] to be a periodic vesting account
	ak = tApp.GetAccountKeeper()
	acc := ak.GetAccount(ctx, addrs[0])
	bacc := authtypes.NewBaseAccount(acc.GetAddress(), acc.GetPubKey(), acc.GetAccountNumber(), acc.GetSequence())
	periods := vestingtypes.Periods{
		vestingtypes.Period{Length: int64(1), Amount: cs(c("ukava", 100))},
		vestingtypes.Period{Length: int64(2), Amount: cs(c("ukava", 100))},
		vestingtypes.Period{Length: int64(8), Amount: cs(c("ukava", 100))},
		vestingtypes.Period{Length: int64(5), Amount: cs(c("ukava", 100))},
	}
	bva := vestingtypes.NewBaseVestingAccount(bacc, cs(c("ukava", 400)), ctx.BlockTime().Unix()+16)
	// suite.Require().NoError(err2)
	pva := vestingtypes.NewPeriodicVestingAccountRaw(bva, ctx.BlockTime().Unix(), periods)
	ak.SetAccount(ctx, pva)

	// sets addrs[2] to be a validator vesting account
	acc = ak.GetAccount(ctx, addrs[2])
	bacc = authtypes.NewBaseAccount(acc.GetAddress(), acc.GetPubKey(), acc.GetAccountNumber(), acc.GetSequence())
	bva = vestingtypes.NewBaseVestingAccount(bacc, cs(c("ukava", 400)), ctx.BlockTime().Unix()+16)
	// suite.Require().NoError(err2)
	// vva := validatorvesting.NewValidatorVestingAccountRaw(bva, ctx.BlockTime().Unix(), periods, sdk.ConsAddress{}, nil, 90)
	// ak.SetAccount(ctx, vva)
	ak.SetAccount(ctx, bva)
	suite.app = tApp
	suite.keeper = tApp.GetIncentiveKeeper()
	suite.ctx = ctx
	suite.addrs = addrs
}

func (suite *KeeperTestSuite) TestGetPeriodLength() {
	type args struct {
		blockTime      time.Time
		multiplier     types.Multiplier
		expectedLength int64
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}
	type periodTest struct {
		name    string
		args    args
		errArgs errArgs
	}
	testCases := []periodTest{
		{
			name: "first half of month",
			args: args{
				blockTime:      time.Date(2020, 11, 2, 15, 0, 0, 0, time.UTC),
				multiplier:     types.NewMultiplier(types.Medium, 6, sdk.MustNewDecFromStr("0.333333")),
				expectedLength: time.Date(2021, 5, 15, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 11, 2, 15, 0, 0, 0, time.UTC).Unix(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "first half of month long lockup",
			args: args{
				blockTime:      time.Date(2020, 11, 2, 15, 0, 0, 0, time.UTC),
				multiplier:     types.NewMultiplier(types.Medium, 24, sdk.MustNewDecFromStr("0.333333")),
				expectedLength: time.Date(2022, 11, 15, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 11, 2, 15, 0, 0, 0, time.UTC).Unix(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "second half of month",
			args: args{
				blockTime:      time.Date(2020, 12, 31, 15, 0, 0, 0, time.UTC),
				multiplier:     types.NewMultiplier(types.Medium, 6, sdk.MustNewDecFromStr("0.333333")),
				expectedLength: time.Date(2021, 7, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 12, 31, 15, 0, 0, 0, time.UTC).Unix(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "second half of month long lockup",
			args: args{
				blockTime:      time.Date(2020, 12, 31, 15, 0, 0, 0, time.UTC),
				multiplier:     types.NewMultiplier(types.Large, 24, sdk.MustNewDecFromStr("0.333333")),
				expectedLength: time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 12, 31, 15, 0, 0, 0, time.UTC).Unix(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "end of feb",
			args: args{
				blockTime:      time.Date(2021, 2, 28, 15, 0, 0, 0, time.UTC),
				multiplier:     types.NewMultiplier(types.Medium, 6, sdk.MustNewDecFromStr("0.333333")),
				expectedLength: time.Date(2021, 9, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2021, 2, 28, 15, 0, 0, 0, time.UTC).Unix(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "leap year",
			args: args{
				blockTime:      time.Date(2020, 2, 29, 15, 0, 0, 0, time.UTC),
				multiplier:     types.NewMultiplier(types.Medium, 6, sdk.MustNewDecFromStr("0.333333")),
				expectedLength: time.Date(2020, 9, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 2, 29, 15, 0, 0, 0, time.UTC).Unix(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "leap year long lockup",
			args: args{
				blockTime:      time.Date(2020, 2, 29, 15, 0, 0, 0, time.UTC),
				multiplier:     types.NewMultiplier(types.Large, 24, sdk.MustNewDecFromStr("1")),
				expectedLength: time.Date(2022, 3, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 2, 29, 15, 0, 0, 0, time.UTC).Unix(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "exactly half of month",
			args: args{
				blockTime:      time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				multiplier:     types.NewMultiplier(types.Medium, 6, sdk.MustNewDecFromStr("0.333333")),
				expectedLength: time.Date(2021, 7, 1, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC).Unix(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			name: "just before half of month",
			args: args{
				blockTime:      time.Date(2020, 12, 15, 13, 59, 59, 0, time.UTC),
				multiplier:     types.NewMultiplier(types.Medium, 6, sdk.MustNewDecFromStr("0.333333")),
				expectedLength: time.Date(2021, 6, 15, 14, 0, 0, 0, time.UTC).Unix() - time.Date(2020, 12, 15, 13, 59, 59, 0, time.UTC).Unix(),
			},
			errArgs: errArgs{
				expectPass: true,
				contains:   "",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			ctx := suite.ctx.WithBlockTime(tc.args.blockTime)
			length, err := suite.keeper.GetPeriodLength(ctx, tc.args.multiplier)
			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.args.expectedLength, length)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
