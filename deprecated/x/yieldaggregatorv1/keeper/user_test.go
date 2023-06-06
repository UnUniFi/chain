package keeper_test

// import (
// 	"github.com/cometbft/cometbft/crypto/ed25519"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

// 	"github.com/UnUniFi/chain/deprecated/x/yieldaggregatorv1/types"
// )

// func (suite *KeeperTestSuite) TestUserDepositGetSet() {
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	addr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

// 	// get initial user deposit
// 	deposit := suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr1)
// 	suite.Require().Equal(deposit, sdk.Coins{})

// 	// set user deposit
// 	coins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 10000))
// 	suite.app.YieldaggregatorKeeper.SetUserDeposit(suite.ctx, addr1, coins)
// 	suite.app.YieldaggregatorKeeper.SetUserDeposit(suite.ctx, addr2, coins)

// 	// check user deposit
// 	deposit = suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr1)
// 	suite.Require().Equal(deposit, coins)
// 	deposit = suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr2)
// 	suite.Require().Equal(deposit, coins)
// 	deposits := suite.app.YieldaggregatorKeeper.GetAllUserDeposits(suite.ctx)
// 	suite.Require().Len(deposits, 2)

// 	// delete user deposit and check
// 	suite.app.YieldaggregatorKeeper.DeleteUserDeposit(suite.ctx, addr2)
// 	deposit = suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr2)
// 	suite.Require().Equal(deposit, sdk.Coins{})

// 	// increase user deposit and check
// 	suite.app.YieldaggregatorKeeper.IncreaseUserDeposit(suite.ctx, addr2, coins)
// 	deposit = suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr2)
// 	suite.Require().Equal(deposit, coins)

// 	// decrease user deposit and check
// 	suite.app.YieldaggregatorKeeper.DecreaseUserDeposit(suite.ctx, addr2, coins)
// 	deposit = suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr2)
// 	suite.Require().Equal(deposit, sdk.Coins(nil))

// 	// increase user deposit and check again
// 	suite.app.YieldaggregatorKeeper.IncreaseUserDeposit(suite.ctx, addr2, coins)
// 	deposit = suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr2)
// 	suite.Require().Equal(deposit, coins)
// }

// func (suite *KeeperTestSuite) TestDeposit() {
// 	// try deposit when not enough balance
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	coins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 1000))
// 	err := suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().Error(err)

// 	// deposit success when enough balance exists
// 	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
// 	suite.NoError(err)
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
// 	suite.NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().NoError(err)

// 	// check balance changes for deposit user + module account
// 	balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr1, "uguu")
// 	suite.Require().Equal(balance, sdk.NewInt64Coin("uguu", 0))
// 	moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
// 	balance = suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, "uguu")
// 	suite.Require().Equal(balance, coins[0])

// 	// check user deposit increase
// 	deposit := suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr1)
// 	suite.Require().Equal(deposit, coins)
// }

// func (suite *KeeperTestSuite) TestWithdraw() {
// 	// try withdraw when not enough deposit exists
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	coins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 1000))
// 	err := suite.app.YieldaggregatorKeeper.Withdraw(suite.ctx, &types.MsgWithdraw{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().Error(err)

// 	// withdraw when enough deposit exists
// 	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
// 	suite.NoError(err)
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
// 	suite.NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.Withdraw(suite.ctx, &types.MsgWithdraw{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().NoError(err)

// 	// check balance changes for withdrawn user + module account
// 	balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr1, "uguu")
// 	suite.Require().Equal(balance, coins[0])
// 	moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
// 	balance = suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, "uguu")
// 	suite.Require().Equal(balance, sdk.NewInt64Coin("uguu", 0))

// 	// check user deposit decrease
// 	deposit := suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr1)
// 	suite.Require().Nil(deposit)
// }
