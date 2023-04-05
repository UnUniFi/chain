package keeper_test

import (
	"errors"
	"strings"
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"github.com/UnUniFi/chain/app"
	cdptypes "github.com/UnUniFi/chain/x/cdp/types"

	"github.com/UnUniFi/chain/x/incentive/types"
	ununifidisttypes "github.com/UnUniFi/chain/x/ununifidist/types"
)

func (suite *KeeperTestSuite) TestPayoutCdpMintingClaim() {
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
				rewardsPerSecond:         c("uguu", 122354),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				initialCollateral:        c("bnb", 1000000000000),
				initialPrincipal:         c("jpu", 10000000000),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              86400,
				expectedBalance:          cs(c("jpu", 10000000000), c("uguu", 10576385600)),
				expectedPeriods:          vestingtypes.Periods{vestingtypes.Period{Length: 32918400, Amount: cs(c("uguu", 10571385600))}},
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
				rewardsPerSecond:         c("uguu", 0),
				initialTime:              time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC),
				initialCollateral:        c("bnb", 1000000000000),
				initialPrincipal:         c("jpu", 10000000000),
				multipliers:              types.Multipliers{types.NewMultiplier(types.MultiplierName("small"), 1, d("0.25")), types.NewMultiplier(types.MultiplierName("large"), 12, d("1.0"))},
				multiplier:               types.MultiplierName("large"),
				timeElapsed:              86400,
				expectedBalance:          cs(c("jpu", 10000000000)),
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
				tc.args.multipliers,
				tc.args.initialTime.Add(time.Hour*24*365*5),
			)
			suite.keeper.SetParams(suite.ctx, params)
			suite.keeper.SetPreviousCdpMintingAccrualTime(suite.ctx, tc.args.ctype, tc.args.initialTime)
			suite.keeper.SetCdpMintingRewardFactor(suite.ctx, tc.args.ctype, sdk.ZeroDec())

			// setup account state
			sk := suite.app.GetBankKeeper()
			err := sk.MintCoins(suite.ctx, cdptypes.ModuleName, sdk.NewCoins(tc.args.initialCollateral))
			suite.Require().NoError(err)
			err = sk.SendCoinsFromModuleToAccount(suite.ctx, cdptypes.ModuleName, suite.addrs[0], sdk.NewCoins(tc.args.initialCollateral))
			suite.Require().NoError(err)

			// setup kavadist state
			err = sk.MintCoins(suite.ctx, ununifidisttypes.ModuleName, cs(c("uguu", 1000000000000)))
			suite.Require().NoError(err)

			// setup cdp state
			cdpKeeper := suite.app.GetCDPKeeper()
			err = cdpKeeper.AddCdp(suite.ctx, suite.addrs[0], tc.args.initialCollateral, tc.args.initialPrincipal, tc.args.ctype)
			suite.Require().NoError(err)

			claim, found := suite.keeper.GetCdpMintingClaim(suite.ctx, suite.addrs[0])
			suite.Require().True(found)
			suite.Require().Equal(sdk.ZeroDec(), claim.RewardIndexes[0].RewardFactor)

			updatedBlockTime := suite.ctx.BlockTime().Add(time.Duration(int(time.Second) * tc.args.timeElapsed))
			suite.ctx = suite.ctx.WithBlockTime(updatedBlockTime)
			rewardPeriod, found := suite.keeper.GetCdpMintingRewardPeriod(suite.ctx, tc.args.ctype)
			suite.Require().True(found)
			err = suite.keeper.AccumulateCdpMintingRewards(suite.ctx, rewardPeriod)
			suite.Require().NoError(err)

			err = suite.keeper.ClaimCdpMintingReward(suite.ctx, suite.addrs[0], string(tc.args.multiplier))

			if tc.errArgs.expectPass {
				suite.Require().NoError(err)
				ak := suite.app.GetAccountKeeper()
				acc := ak.GetAccount(suite.ctx, suite.addrs[0])
				bk := suite.app.GetBankKeeper()
				suite.Require().Equal(tc.args.expectedBalance, bk.GetAllBalances(suite.ctx, acc.GetAddress()))

				if tc.args.isPeriodicVestingAccount {
					vacc, ok := acc.(*vestingtypes.PeriodicVestingAccount)
					suite.Require().True(ok)
					suite.Require().Equal(tc.args.expectedPeriods, vestingtypes.Periods(vacc.VestingPeriods))
				}

				claim, found := suite.keeper.GetCdpMintingClaim(suite.ctx, suite.addrs[0])
				suite.Require().True(found)
				suite.Require().Equal(c("uguu", 0), claim.Reward)
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
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
					origVestingCoins: cs(c("uguu", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 2, Amount: cs(c("uguu", 6))},
				ctxTime:             time.Unix(101, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 3, Amount: cs(c("uguu", 6))},
					vestingtypes.Period{Length: 2, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
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
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
					origVestingCoins: cs(c("uguu", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("uguu", 6))},
				ctxTime:             time.Unix(80, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 7, Amount: cs(c("uguu", 6))},
					vestingtypes.Period{Length: 18, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
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
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
					origVestingCoins: cs(c("uguu", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("uguu", 6))},
				ctxTime:             time.Unix(101, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 3, Amount: cs(c("uguu", 6))},
					vestingtypes.Period{Length: 2, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
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
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
					origVestingCoins: cs(c("uguu", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("uguu", 6))},
				ctxTime:             time.Unix(125, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 12, Amount: cs(c("uguu", 6))}},
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
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
					origVestingCoins: cs(c("uguu", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 6))},
				ctxTime:             time.Unix(110, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 11))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
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
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
					origVestingCoins: cs(c("uguu", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 7, Amount: cs(c("uguu", 6))},
				ctxTime:             time.Unix(125, 0),
				mintModAccountCoins: false,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 12, Amount: cs(c("uguu", 6))}},
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
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
						vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))}},
					origVestingCoins: cs(c("uguu", 20)),
					startTime:        100,
					endTime:          120,
				},
				period:              vestingtypes.Period{Length: 50, Amount: cs(c("uguu", 6))},
				ctxTime:             time.Unix(110, 0),
				mintModAccountCoins: true,
				expectedPeriods: vestingtypes.Periods{
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 5, Amount: cs(c("uguu", 5))},
					vestingtypes.Period{Length: 40, Amount: cs(c("uguu", 6))}},
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
				err = sk.MintCoins(suite.ctx, ununifidisttypes.ModuleName, tc.args.period.Amount)
				suite.Require().NoError(err)
			}

			err = suite.keeper.SendTimeLockedCoinsToPeriodicVestingAccount(suite.ctx, ununifidisttypes.ModuleName, pva.GetAddress(), tc.args.period.Amount, tc.args.period.Length)
			if tc.errArgs.expectErr {
				suite.Require().Error(err)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			} else {
				suite.Require().NoError(err)

				acc := suite.getAccount(pva.GetAddress())
				vacc, ok := acc.(*vestingtypes.PeriodicVestingAccount)
				suite.Require().True(ok)
				suite.Require().Equal(tc.args.expectedPeriods, vestingtypes.Periods(vacc.VestingPeriods))
				suite.Require().Equal(tc.args.expectedStartTime, vacc.StartTime)
				suite.Require().Equal(tc.args.expectedEndTime, vacc.EndTime)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestSendCoinsToBaseAccount() {
	suite.SetupWithAccountState()
	// send coins to base account
	err := suite.keeper.SendTimeLockedCoinsToAccount(suite.ctx, ununifidisttypes.ModuleName, suite.addrs[1], cs(c("uguu", 100)), 5)
	suite.Require().NoError(err)
	acc := suite.getAccount(suite.addrs[1])
	vacc, ok := acc.(*vestingtypes.PeriodicVestingAccount)
	suite.True(ok)
	expectedPeriods := vestingtypes.Periods{
		vestingtypes.Period{Length: int64(5), Amount: cs(c("uguu", 100))},
	}
	bk := suite.app.GetBankKeeper()
	suite.Equal(expectedPeriods, vestingtypes.Periods(vacc.VestingPeriods))
	suite.Equal(cs(c("uguu", 100)), vacc.OriginalVesting)
	suite.Equal(cs(c("uguu", 500)), bk.GetAllBalances(suite.ctx, vacc.GetAddress()))
	suite.Equal(int64(105), vacc.EndTime)
	suite.Equal(int64(100), vacc.StartTime)

}

func (suite *KeeperTestSuite) TestSendCoinsToInvalidAccount() {
	suite.SetupWithAccountState()
	err := suite.keeper.SendTimeLockedCoinsToAccount(suite.ctx, ununifidisttypes.ModuleName, suite.addrs[2], cs(c("uguu", 100)), 5)
	suite.Require().True(errors.Is(err, types.ErrInvalidAccountType))
	macc := suite.getModuleAccount(cdptypes.ModuleName)
	err = suite.keeper.SendTimeLockedCoinsToAccount(suite.ctx, ununifidisttypes.ModuleName, macc.GetAddress(), cs(c("uguu", 100)), 5)
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
			cs(c("uguu", 400)),
			cs(c("uguu", 400)),
			cs(c("uguu", 400)),
			cs(c("uguu", 400)),
		})
	tApp.InitializeFromGenesisStates(
		authGS,
	)
	ak := tApp.GetAccountKeeper()
	bk := tApp.GetBankKeeper()
	macc := ak.GetModuleAccount(ctx, ununifidisttypes.ModuleName)
	err := bk.MintCoins(ctx, macc.GetName(), cs(c("uguu", 600)))
	suite.Require().NoError(err)

	// sets addrs[0] to be a periodic vesting account
	ak = tApp.GetAccountKeeper()
	acc := ak.GetAccount(ctx, addrs[0])
	bacc := authtypes.NewBaseAccount(acc.GetAddress(), acc.GetPubKey(), acc.GetAccountNumber(), acc.GetSequence())
	periods := vestingtypes.Periods{
		vestingtypes.Period{Length: int64(1), Amount: cs(c("uguu", 100))},
		vestingtypes.Period{Length: int64(2), Amount: cs(c("uguu", 100))},
		vestingtypes.Period{Length: int64(8), Amount: cs(c("uguu", 100))},
		vestingtypes.Period{Length: int64(5), Amount: cs(c("uguu", 100))},
	}
	bva := vestingtypes.NewBaseVestingAccount(bacc, cs(c("uguu", 400)), ctx.BlockTime().Unix()+16)
	// suite.Require().NoError(err2)
	pva := vestingtypes.NewPeriodicVestingAccountRaw(bva, ctx.BlockTime().Unix(), periods)
	ak.SetAccount(ctx, pva)

	// sets addrs[2] to be a validator vesting account
	acc = ak.GetAccount(ctx, addrs[2])
	bacc = authtypes.NewBaseAccount(acc.GetAddress(), acc.GetPubKey(), acc.GetAccountNumber(), acc.GetSequence())
	bva = vestingtypes.NewBaseVestingAccount(bacc, cs(c("uguu", 400)), ctx.BlockTime().Unix()+16)
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
