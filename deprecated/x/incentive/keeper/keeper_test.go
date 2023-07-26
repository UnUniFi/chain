package keeper_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/suite"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	// authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
// 	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

// 	// supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"

// 	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
// 	tmtime "github.com/cometbft/cometbft/types/time"

// 	"github.com/UnUniFi/chain/app"
// 	// committeekeeper "github.com/UnUniFi/chain/x/committee/keeper"
// 	// hardkeeper "github.com/UnUniFi/chain/x/hard/keeper"

// 	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

// 	"github.com/UnUniFi/chain/deprecated/x/incentive/keeper"
// 	"github.com/UnUniFi/chain/deprecated/x/incentive/types"
// )

// // Test suite used for all keeper tests
// type KeeperTestSuite struct {
// 	suite.Suite

// 	keeper keeper.Keeper
// 	// hardKeeper      hardkeeper.Keeper
// 	stakingKeeper *stakingkeeper.Keeper
// 	// committeeKeeper committeekeeper.Keeper
// 	app            app.TestApp
// 	ctx            sdk.Context
// 	addrs          []sdk.AccAddress
// 	validatorAddrs []sdk.ValAddress
// }

// // The default state used by each test
// func (suite *KeeperTestSuite) SetupTest() {
// 	tApp := app.NewTestApp()
// 	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

// 	tApp.InitializeFromGenesisStates()

// 	_, addrs := app.GeneratePrivKeyAddressPairs(5)
// 	keeper := tApp.GetIncentiveKeeper()
// 	suite.app = tApp
// 	suite.ctx = ctx
// 	suite.keeper = keeper
// 	suite.addrs = addrs
// }

// func (suite *KeeperTestSuite) getAccount(addr sdk.AccAddress) authtypes.AccountI {
// 	ak := suite.app.GetAccountKeeper()
// 	return ak.GetAccount(suite.ctx, addr)
// }

// func (suite *KeeperTestSuite) getModuleAccount(name string) authtypes.ModuleAccountI {
// 	sk := suite.app.GetAccountKeeper()
// 	return sk.GetModuleAccount(suite.ctx, name)
// }

// func (suite *KeeperTestSuite) TestGetSetDeleteCdpMintingClaim() {
// 	c := types.NewCdpMintingClaim(suite.addrs[0], c("uguu", 1000000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
// 	_, found := suite.keeper.GetCdpMintingClaim(suite.ctx, suite.addrs[0])
// 	suite.Require().False(found)
// 	suite.Require().NotPanics(func() {
// 		suite.keeper.SetCdpMintingClaim(suite.ctx, c)
// 	})
// 	testC, found := suite.keeper.GetCdpMintingClaim(suite.ctx, suite.addrs[0])
// 	suite.Require().True(found)
// 	suite.Require().Equal(c, testC)
// 	suite.Require().NotPanics(func() {
// 		suite.keeper.DeleteCdpMintingClaim(suite.ctx, suite.addrs[0])
// 	})
// 	_, found = suite.keeper.GetCdpMintingClaim(suite.ctx, suite.addrs[0])
// 	suite.Require().False(found)
// }

// func (suite *KeeperTestSuite) TestIterateJPYXMintingClaims() {
// 	for i := 0; i < len(suite.addrs); i++ {
// 		c := types.NewCdpMintingClaim(suite.addrs[i], c("uguu", 100000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
// 		suite.Require().NotPanics(func() {
// 			suite.keeper.SetCdpMintingClaim(suite.ctx, c)
// 		})
// 	}
// 	claims := types.CdpMintingClaims{}
// 	suite.keeper.IterateCdpMintingClaims(suite.ctx, func(c types.CdpMintingClaim) bool {
// 		claims = append(claims, c)
// 		return false
// 	})
// 	suite.Require().Equal(len(suite.addrs), len(claims))

// 	claims = suite.keeper.GetAllCdpMintingClaims(suite.ctx)
// 	suite.Require().Equal(len(suite.addrs), len(claims))
// }

// func (suite *KeeperTestSuite) fundAccount(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
// 	err := suite.app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
// 	if err != nil {
// 		return err
// 	}
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, coins)
// 	return err
// }

// func (suite *KeeperTestSuite) createPeriodicVestingAccount(origVesting sdk.Coins, periods vestingtypes.Periods, startTime, endTime int64) (*vestingtypes.PeriodicVestingAccount, error) {
// 	_, addr := app.GeneratePrivKeyAddressPairs(1)
// 	bacc := authtypes.NewBaseAccountWithAddress(addr[0])
// 	suite.fundAccount(suite.ctx, bacc.GetAddress(), origVesting)
// 	bva := vestingtypes.NewBaseVestingAccount(bacc, origVesting, endTime)
// 	err := bva.Validate()
// 	if err != nil {
// 		return &vestingtypes.PeriodicVestingAccount{}, err
// 	}
// 	pva := vestingtypes.NewPeriodicVestingAccountRaw(bva, startTime, periods)
// 	err = pva.Validate()
// 	if err != nil {
// 		return &vestingtypes.PeriodicVestingAccount{}, err
// 	}
// 	return pva, nil
// }

// // Avoid cluttering test cases with long function names
// func i(in int64) sdk.Int                    { return sdk.NewInt(in) }
// func d(str string) sdk.Dec                  { return sdk.MustNewDecFromStr(str) }
// func c(denom string, amount int64) sdk.Coin { return sdk.NewInt64Coin(denom, amount) }
// func cs(coins ...sdk.Coin) sdk.Coins        { return sdk.NewCoins(coins...) }

// func TestKeeperTestSuite(t *testing.T) {
// 	suite.Run(t, new(KeeperTestSuite))
// }
