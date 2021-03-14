package incentive_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/lcnem/jpyx/app"
	"github.com/lcnem/jpyx/x/incentive"
	"github.com/lcnem/jpyx/x/incentive/types"
	"github.com/lcnem/jpyx/x/jsmndist"
)

func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }

type HandlerTestSuite struct {
	suite.Suite

	ctx     sdk.Context
	app     app.TestApp
	handler sdk.Handler
	keeper  incentive.Keeper
	addrs   []sdk.AccAddress
}

func (suite *HandlerTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{Height: 1, Time: tmtime.Now()})
	keeper := tApp.GetIncentiveKeeper()

	// Set up genesis state and initialize
	_, addrs := app.GeneratePrivKeyAddressPairs(3)
	coins := []sdk.Coins{}
	for j := 0; j < 3; j++ {
		coins = append(coins, cs(c("bnb", 10000000000), c("ujsmn", 10000000000)))
	}
	authGS := app.NewAuthGenState(addrs, coins)
	incentiveGS := incentive.NewGenesisState(
		incentive.NewParams(
			incentive.RewardPeriods{incentive.NewRewardPeriod(true, "bnb-a", time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC), time.Date(2024, 12, 15, 14, 0, 0, 0, time.UTC), c("ujsmn", 122354))},
			incentive.MultiRewardPeriods{incentive.NewMultiRewardPeriod(true, "bnb-a", time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC), time.Date(2024, 12, 15, 14, 0, 0, 0, time.UTC), cs(c("ujsmn", 122354)))},
			incentive.MultiRewardPeriods{incentive.NewMultiRewardPeriod(true, "bnb-a", time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC), time.Date(2024, 12, 15, 14, 0, 0, 0, time.UTC), cs(c("ujsmn", 122354)))},
			incentive.RewardPeriods{incentive.NewRewardPeriod(true, "bnb-a", time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC), time.Date(2024, 12, 15, 14, 0, 0, 0, time.UTC), c("ujsmn", 122354))},
			incentive.Multipliers{incentive.NewMultiplier(incentive.MultiplierName("small"), 1, d("0.25")), incentive.NewMultiplier(incentive.MultiplierName("large"), 12, d("1.0"))},
			time.Date(2025, 12, 15, 14, 0, 0, 0, time.UTC),
		),
		incentive.DefaultGenesisAccumulationTimes,
		incentive.DefaultGenesisAccumulationTimes,
		incentive.DefaultGenesisAccumulationTimes,
		incentive.DefaultGenesisAccumulationTimes,
		incentive.DefaultJPYXClaims,
		incentive.DefaultHardClaims,
	)
	tApp.InitializeFromGenesisStates(authGS, app.GenesisState{incentive.ModuleName: incentive.ModuleCdc.MustMarshalJSON(incentiveGS)}, NewCDPGenStateMulti(), NewPricefeedGenStateMulti())

	suite.addrs = addrs
	suite.handler = incentive.NewHandler(keeper)
	suite.keeper = keeper
	suite.app = tApp
	suite.ctx = ctx
}

func (suite *HandlerTestSuite) TestMsgJPYXMintingClaimReward() {
	suite.addJPYXMintingClaim()
	msg := incentive.NewMsgClaimJPYXMintingReward(suite.addrs[0], "small")
	res, err := suite.handler(suite.ctx, msg)
	suite.NoError(err)
	suite.Require().NotNil(res)
}

func (suite *HandlerTestSuite) TestMsgHardLiquidityProviderClaimReward() {
	suite.addHardLiquidityProviderClaim()
	msg := incentive.NewMsgClaimHardLiquidityProviderReward(suite.addrs[0], "small")
	res, err := suite.handler(suite.ctx, msg)
	suite.NoError(err)
	suite.Require().NotNil(res)
}

func (suite *HandlerTestSuite) addHardLiquidityProviderClaim() {
	sk := suite.app.GetSupplyKeeper()
	err := sk.MintCoins(suite.ctx, jsmndist.ModuleName, cs(c("ujsmn", 1000000000000)))
	suite.Require().NoError(err)
	rewardPeriod := types.RewardIndexes{types.NewRewardIndex("bnb-s", sdk.ZeroDec())}

	multiRewardIndex := types.NewMultiRewardIndex("bnb-s", rewardPeriod)
	multiRewardIndexes := types.MultiRewardIndexes{multiRewardIndex}
	c1 := incentive.NewHardLiquidityProviderClaim(suite.addrs[0], cs(c("ujsmn", 1000000)), multiRewardIndexes, multiRewardIndexes, rewardPeriod)
	suite.NotPanics(func() {
		suite.keeper.SetHardLiquidityProviderClaim(suite.ctx, c1)
	})
}

func (suite *HandlerTestSuite) addJPYXMintingClaim() {
	sk := suite.app.GetSupplyKeeper()
	err := sk.MintCoins(suite.ctx, jsmndist.ModuleName, cs(c("ujsmn", 1000000000000)))
	suite.Require().NoError(err)
	c1 := incentive.NewJPYXMintingClaim(suite.addrs[0], c("ujsmn", 1000000), types.RewardIndexes{types.NewRewardIndex("bnb-s", sdk.ZeroDec())})
	suite.NotPanics(func() {
		suite.keeper.SetJPYXMintingClaim(suite.ctx, c1)
	})
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

// Avoid cluttering test cases with long function names
func i(in int64) sdk.Int   { return sdk.NewInt(in) }
func d(str string) sdk.Dec { return sdk.MustNewDecFromStr(str) }
