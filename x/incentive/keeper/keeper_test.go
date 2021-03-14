package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/lcnem/jpyx/app"
	committeekeeper "github.com/lcnem/jpyx/x/committee/keeper"
	hardkeeper "github.com/lcnem/jpyx/x/hard/keeper"
	"github.com/lcnem/jpyx/x/incentive/keeper"
	"github.com/lcnem/jpyx/x/incentive/types"
)

// Test suite used for all keeper tests
type KeeperTestSuite struct {
	suite.Suite

	keeper          keeper.Keeper
	hardKeeper      hardkeeper.Keeper
	stakingKeeper   stakingkeeper.Keeper
	committeeKeeper committeekeeper.Keeper
	app             app.TestApp
	ctx             sdk.Context
	addrs           []sdk.AccAddress
	validatorAddrs  []sdk.ValAddress
}

// The default state used by each test
func (suite *KeeperTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{Height: 1, Time: tmtime.Now()})

	tApp.InitializeFromGenesisStates()

	_, addrs := app.GeneratePrivKeyAddressPairs(5)
	keeper := tApp.GetIncentiveKeeper()
	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = keeper
	suite.addrs = addrs
}

func (suite *KeeperTestSuite) getAccount(addr sdk.AccAddress) authexported.Account {
	ak := suite.app.GetAccountKeeper()
	return ak.GetAccount(suite.ctx, addr)
}

func (suite *KeeperTestSuite) getModuleAccount(name string) supplyexported.ModuleAccountI {
	sk := suite.app.GetSupplyKeeper()
	return sk.GetModuleAccount(suite.ctx, name)
}

func (suite *KeeperTestSuite) TestGetSetDeleteJPYXMintingClaim() {
	c := types.NewJPYXMintingClaim(suite.addrs[0], c("ujsmn", 1000000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
	_, found := suite.keeper.GetJPYXMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
	suite.Require().NotPanics(func() {
		suite.keeper.SetJPYXMintingClaim(suite.ctx, c)
	})
	testC, found := suite.keeper.GetJPYXMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().True(found)
	suite.Require().Equal(c, testC)
	suite.Require().NotPanics(func() {
		suite.keeper.DeleteJPYXMintingClaim(suite.ctx, suite.addrs[0])
	})
	_, found = suite.keeper.GetJPYXMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestIterateJPYXMintingClaims() {
	for i := 0; i < len(suite.addrs); i++ {
		c := types.NewJPYXMintingClaim(suite.addrs[i], c("ujsmn", 100000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
		suite.Require().NotPanics(func() {
			suite.keeper.SetJPYXMintingClaim(suite.ctx, c)
		})
	}
	claims := types.JPYXMintingClaims{}
	suite.keeper.IterateJPYXMintingClaims(suite.ctx, func(c types.JPYXMintingClaim) bool {
		claims = append(claims, c)
		return false
	})
	suite.Require().Equal(len(suite.addrs), len(claims))

	claims = suite.keeper.GetAllJPYXMintingClaims(suite.ctx)
	suite.Require().Equal(len(suite.addrs), len(claims))
}

func createPeriodicVestingAccount(origVesting sdk.Coins, periods vesting.Periods, startTime, endTime int64) (*vesting.PeriodicVestingAccount, error) {
	_, addr := app.GeneratePrivKeyAddressPairs(1)
	bacc := auth.NewBaseAccountWithAddress(addr[0])
	bacc.Coins = origVesting
	bva, err := vesting.NewBaseVestingAccount(&bacc, origVesting, endTime)
	if err != nil {
		return &vesting.PeriodicVestingAccount{}, err
	}
	pva := vesting.NewPeriodicVestingAccountRaw(bva, startTime, periods)
	err = pva.Validate()
	if err != nil {
		return &vesting.PeriodicVestingAccount{}, err
	}
	return pva, nil
}

// Avoid cluttering test cases with long function names
func i(in int64) sdk.Int                    { return sdk.NewInt(in) }
func d(str string) sdk.Dec                  { return sdk.MustNewDecFromStr(str) }
func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
