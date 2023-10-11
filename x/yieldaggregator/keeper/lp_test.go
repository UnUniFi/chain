package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (suite *KeeperTestSuite) TestVaultAmountUnbondingAmountInStrategies() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	atomHostDenom := "uatom"
	prefixedDenom := transfertypes.GetPrefixedDenom("transfer", "channel-0", atomHostDenom)
	atomIbcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()

	strategy := types.Strategy{
		Id:              1,
		Name:            "AtomLiquidStaking",
		ContractAddress: "x/ibc-staking",
		Denom:           atomIbcDenom,
	}
	suite.app.YieldaggregatorKeeper.SetStrategy(suite.ctx, strategy.Denom, strategy)

	vault := types.Vault{
		Id:                     1,
		Symbol:                 "ATOM",
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.ZeroDec(),
		WithdrawReserveRate:    sdk.ZeroDec(),
		StrategyWeights: []types.StrategyWeight{
			{Denom: atomIbcDenom, StrategyId: 1, Weight: sdk.OneDec()},
		},
	}

	amount := suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	// mint coins to be spent on liquid staking
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	coins := sdk.Coins{sdk.NewInt64Coin(atomIbcDenom, 1000000)}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, vaultModAddr, coins)
	suite.Require().NoError(err)

	// stake to strategy - calls liquid staking
	suite.SetupZoneAndEpoch(atomHostDenom, atomIbcDenom)
	err = suite.app.YieldaggregatorKeeper.StakeToStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000))
	suite.Require().NoError(err)

	amount = suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "1000000")

	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	// unstake from strategy - calls redeem stake
	err = suite.app.YieldaggregatorKeeper.UnstakeFromStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000), "")
	suite.Require().NoError(err)

	// check amounts after unstake
	amount = suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "1000000")
}

func (suite *KeeperTestSuite) TestVaultWithdrawalAmount() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	atomHostDenom := "uatom"
	prefixedDenom := transfertypes.GetPrefixedDenom("transfer", "channel-0", atomHostDenom)
	atomIbcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()

	vault := types.Vault{
		Id:                     1,
		Symbol:                 "ATOM",
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.ZeroDec(),
		WithdrawReserveRate:    sdk.ZeroDec(),
		StrategyWeights: []types.StrategyWeight{
			{Denom: atomIbcDenom, StrategyId: 1, Weight: sdk.OneDec()},
		},
	}

	amount := suite.app.YieldaggregatorKeeper.VaultWithdrawalAmount(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	// mint coins to be spent on liquid staking
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	coins := sdk.Coins{sdk.NewInt64Coin(atomIbcDenom, 1000000)}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, vaultModAddr, coins)
	suite.Require().NoError(err)

	amount = suite.app.YieldaggregatorKeeper.VaultWithdrawalAmount(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "1000000")
}

func (suite *KeeperTestSuite) TestVaultAmountTotal() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	atomHostDenom := "uatom"
	prefixedDenom := transfertypes.GetPrefixedDenom("transfer", "channel-0", atomHostDenom)
	atomIbcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()

	strategy := types.Strategy{
		Id:              1,
		Name:            "AtomLiquidStaking",
		ContractAddress: "x/ibc-staking",
		Denom:           atomIbcDenom,
	}
	suite.app.YieldaggregatorKeeper.SetStrategy(suite.ctx, strategy.Denom, strategy)

	vault := types.Vault{
		Id:                     1,
		Symbol:                 "ATOM",
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.ZeroDec(),
		WithdrawReserveRate:    sdk.ZeroDec(),
		StrategyWeights: []types.StrategyWeight{
			{Denom: atomIbcDenom, StrategyId: 1, Weight: sdk.OneDec()},
		},
	}

	amount := suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	// mint coins to be spent on liquid staking
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	coins := sdk.Coins{sdk.NewInt64Coin(atomIbcDenom, 10000000)}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, vaultModAddr, coins)
	suite.Require().NoError(err)

	amount = suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "10000000")

	// stake to strategy - calls liquid staking
	suite.SetupZoneAndEpoch(atomHostDenom, atomIbcDenom)
	err = suite.app.YieldaggregatorKeeper.StakeToStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000))
	suite.Require().NoError(err)

	amount = suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "10000000")

	// unstake from strategy - calls redeem stake
	err = suite.app.YieldaggregatorKeeper.UnstakeFromStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000), "")
	suite.Require().NoError(err)

	// check amounts after unstake
	amount = suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "10000000")
}

func (suite *KeeperTestSuite) TestEstimateMintRedeemAmountInternal() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	atomHostDenom := "uatom"
	prefixedDenom := transfertypes.GetPrefixedDenom("transfer", "channel-0", atomHostDenom)
	atomIbcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()
	lpDenom := types.GetLPTokenDenom(1)
	_ = suite.SetupZoneAndEpoch(atomHostDenom, atomIbcDenom)

	strategy := types.Strategy{
		Id:              1,
		Name:            "AtomLiquidStaking",
		ContractAddress: "x/ibc-staking",
		Denom:           atomIbcDenom,
	}
	suite.app.YieldaggregatorKeeper.SetStrategy(suite.ctx, strategy.Denom, strategy)

	suite.app.YieldaggregatorKeeper.SetDenomInfo(suite.ctx, types.DenomInfo{
		Denom:  atomIbcDenom,
		Symbol: "ATOM",
	})
	vault := types.Vault{
		Id:                     1,
		Symbol:                 "ATOM",
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.NewDecWithPrec(1, 1), // 10%
		WithdrawReserveRate:    sdk.NewDecWithPrec(1, 1), // 10%
		StrategyWeights: []types.StrategyWeight{
			{Denom: atomIbcDenom, StrategyId: 1, Weight: sdk.OneDec()},
		},
	}
	suite.app.YieldaggregatorKeeper.SetVault(suite.ctx, vault)

	// mint coins to be spent on liquid staking
	coins := sdk.Coins{sdk.NewInt64Coin(atomIbcDenom, 1000000)}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
	suite.Require().NoError(err)

	// try execution after setup
	err = suite.app.YieldaggregatorKeeper.DepositAndMintLPToken(suite.ctx, addr1, 1, sdk.NewInt64Coin(atomIbcDenom, 100000))
	suite.Require().NoError(err)

	estMintAmount := suite.app.YieldaggregatorKeeper.EstimateMintAmountInternal(suite.ctx, vault.Id, sdk.NewInt(100000))
	suite.Require().Equal(estMintAmount.String(), "100000"+lpDenom)

	estBurnAmount := suite.app.YieldaggregatorKeeper.EstimateRedeemAmountInternal(suite.ctx, vault.Id, sdk.NewInt(100000))
	suite.Require().Equal(estBurnAmount.String(), "100000")
}

func (suite *KeeperTestSuite) TestMintBurnLPToken() {
	atomHostDenom := "uatom"
	prefixedDenom := transfertypes.GetPrefixedDenom("transfer", "channel-0", atomHostDenom)
	atomIbcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()

	// try execution with invalid vault id
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	err := suite.app.YieldaggregatorKeeper.DepositAndMintLPToken(suite.ctx, addr1, 1, sdk.NewInt64Coin(atomIbcDenom, 100000))
	suite.Require().Error(err)

	// try execution with invalid vault id
	err = suite.app.YieldaggregatorKeeper.BurnLPTokenAndRedeem(suite.ctx, addr1, 1, sdk.NewInt(100000))
	suite.Require().Error(err)

	lpDenom := types.GetLPTokenDenom(1)
	_ = suite.SetupZoneAndEpoch(atomHostDenom, atomIbcDenom)

	strategy := types.Strategy{
		Id:              1,
		Name:            "AtomLiquidStaking",
		ContractAddress: "x/ibc-staking",
		Denom:           atomIbcDenom,
	}
	suite.app.YieldaggregatorKeeper.SetStrategy(suite.ctx, strategy.Denom, strategy)

	suite.app.YieldaggregatorKeeper.SetDenomInfo(suite.ctx, types.DenomInfo{
		Denom:  atomIbcDenom,
		Symbol: "ATOM",
	})
	vault := types.Vault{
		Id:                     1,
		Symbol:                 "ATOM",
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.NewDecWithPrec(1, 1), // 10%
		WithdrawReserveRate:    sdk.NewDecWithPrec(1, 1), // 10%
		StrategyWeights: []types.StrategyWeight{
			{Denom: atomIbcDenom, StrategyId: 1, Weight: sdk.OneDec()},
		},
	}
	suite.app.YieldaggregatorKeeper.SetVault(suite.ctx, vault)

	// mint coins to be spent on liquid staking
	coins := sdk.Coins{sdk.NewInt64Coin(atomIbcDenom, 1000000)}
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
	suite.Require().NoError(err)

	// try execution after setup
	err = suite.app.YieldaggregatorKeeper.DepositAndMintLPToken(suite.ctx, addr1, 1, sdk.NewInt64Coin(atomIbcDenom, 100000))
	suite.Require().NoError(err)

	// check changes in user balance
	balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr1, atomIbcDenom)
	suite.Require().Equal(balance.String(), "900000"+atomIbcDenom)
	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, lpDenom)
	suite.Require().Equal(balance.String(), "100000"+lpDenom)

	// check changes in strategies, and reserve balance
	amount := suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "100000")

	amount = suite.app.YieldaggregatorKeeper.VaultWithdrawalAmount(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "10000")

	amount = suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "90000")

	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	// burn execution
	err = suite.app.YieldaggregatorKeeper.BurnLPTokenAndRedeem(suite.ctx, addr1, 1, sdk.NewInt(1000))
	suite.Require().NoError(err)

	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, atomIbcDenom)
	suite.Require().Equal(balance.String(), "901000"+atomIbcDenom)
	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, lpDenom)
	suite.Require().Equal(balance.String(), "99000"+lpDenom)
	amount = suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "99000")
	amount = suite.app.YieldaggregatorKeeper.VaultWithdrawalAmount(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "9000")
	amount = suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "90000")
	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	err = suite.app.YieldaggregatorKeeper.BurnLPTokenAndRedeem(suite.ctx, addr1, 1, sdk.NewInt(1000))
	suite.Require().NoError(err)

	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, atomIbcDenom)
	suite.Require().Equal(balance.String(), "902000"+atomIbcDenom)
	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, lpDenom)
	suite.Require().Equal(balance.String(), "98000"+lpDenom)
	amount = suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "98000")
	amount = suite.app.YieldaggregatorKeeper.VaultWithdrawalAmount(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "8000")
	amount = suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "90000")
	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	err = suite.app.YieldaggregatorKeeper.BurnLPTokenAndRedeem(suite.ctx, addr1, 1, sdk.NewInt(1000))
	suite.Require().NoError(err)

	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, atomIbcDenom)
	suite.Require().Equal(balance.String(), "903000"+atomIbcDenom)
	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, lpDenom)
	suite.Require().Equal(balance.String(), "97000"+lpDenom)
	amount = suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "97000")
	amount = suite.app.YieldaggregatorKeeper.VaultWithdrawalAmount(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "7000")
	amount = suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "90000")
	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	err = suite.app.YieldaggregatorKeeper.BurnLPTokenAndRedeem(suite.ctx, addr1, 1, sdk.NewInt(100000-3000))
	suite.Require().NoError(err)
	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, atomIbcDenom)
	suite.Require().Equal(balance.String(), "903000"+atomIbcDenom)
	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, lpDenom)
	suite.Require().Equal(balance.String(), "0"+lpDenom)
	amount = suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "97000")
	amount = suite.app.YieldaggregatorKeeper.VaultWithdrawalAmount(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "7000")
	amount = suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "90000")
	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")
}
