package keeper_test

import (
	epochstypes "github.com/UnUniFi/chain/x/epochs/types"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	transfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) TestBeforeEpochStart() {
	// try execution with invalid vault id
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	err := suite.app.YieldaggregatorKeeper.DepositAndMintLPToken(suite.ctx, addr1, 1, sdk.NewInt(100000))
	suite.Require().Error(err)

	// try execution with invalid vault id
	err = suite.app.YieldaggregatorKeeper.BurnLPTokenAndRedeem(suite.ctx, addr1, 1, sdk.NewInt(100000))
	suite.Require().Error(err)

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

	vault := types.Vault{
		Id:                     1,
		Denom:                  atomIbcDenom,
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.NewDecWithPrec(1, 1), // 10%
		WithdrawReserveRate:    sdk.NewDecWithPrec(1, 1), // 10%
		StrategyWeights: []types.StrategyWeight{
			{StrategyId: 1, Weight: sdk.OneDec()},
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
	err = suite.app.YieldaggregatorKeeper.DepositAndMintLPToken(suite.ctx, addr1, 1, sdk.NewInt(100000))
	suite.Require().NoError(err)

	// burn execution
	err = suite.app.YieldaggregatorKeeper.BurnLPTokenAndRedeem(suite.ctx, addr1, 1, sdk.NewInt(1000))
	suite.Require().NoError(err)

	balance := suite.app.BankKeeper.GetBalance(suite.ctx, addr1, atomIbcDenom)
	suite.Require().Equal(balance.String(), "901000"+atomIbcDenom)
	balance = suite.app.BankKeeper.GetBalance(suite.ctx, addr1, lpDenom)
	suite.Require().Equal(balance.String(), "99000"+lpDenom)
	amount := suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "99000")
	amount = suite.app.YieldaggregatorKeeper.VaultWithdrawalAmount(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "9000")
	amount = suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "90000")
	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "0")

	// check after epoch execution
	suite.app.YieldaggregatorKeeper.BeforeEpochStart(suite.ctx, epochstypes.EpochInfo{
		Identifier: "day",
	})
	amount = suite.app.YieldaggregatorKeeper.VaultAmountTotal(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "99000")
	amount = suite.app.YieldaggregatorKeeper.VaultWithdrawalAmount(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "9000")
	amount = suite.app.YieldaggregatorKeeper.VaultAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "89100")
	amount = suite.app.YieldaggregatorKeeper.VaultUnbondingAmountInStrategies(suite.ctx, vault)
	suite.Require().Equal(amount.String(), "900")
}
