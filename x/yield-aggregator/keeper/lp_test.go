package keeper_test

import (
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
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
		Denom:                  atomIbcDenom,
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.ZeroDec(),
		WithdrawReserveRate:    sdk.ZeroDec(),
		StrategyWeights: []types.StrategyWeight{
			{StrategyId: 1, Weight: sdk.OneDec()},
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
	err = suite.app.YieldaggregatorKeeper.UnstakeFromStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000))
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
		Denom:                  atomIbcDenom,
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.ZeroDec(),
		WithdrawReserveRate:    sdk.ZeroDec(),
		StrategyWeights: []types.StrategyWeight{
			{StrategyId: 1, Weight: sdk.OneDec()},
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

// TODO: add test for
// VaultAmountTotal(ctx sdk.Context, vault types.Vault) sdk.Int {
// EstimateMintAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, principalAmount sdk.Int) sdk.Coin {
// EstimateRedeemAmountInternal(ctx sdk.Context, vaultDenom string, vaultId uint64, lpAmount sdk.Int) sdk.Coin {
// DepositAndMintLPToken(ctx sdk.Context, address sdk.AccAddress, vaultId uint64, principalAmount sdk.Int) error {
// TODO: add test for BurnLPTokenAndRedeem
// Imagine
// withdraw reserve 100
// staked 900
// User A execute MsgWithdraw with amount 10.
// maintenance rate is max(0, 100 - 10) / (100 + 0) = 90 / 100 = 0.9
// withdraw reserve 90
// staked 900
// bonding 890
// unbonding 10
// Then User B execute MsgWithdraw with amount 10.
// maintenance rate is max(0, 90 - 10) / (90 + 10) = 80 / 100 = 0.8
// withdraw reserve 80
// staked 900
// bonding 880
// unbonding 20
// after the unbonding period of user A withdrawal, unbonded token will go to withdraw reserve. ( for simplification, don't think the rebalancing now)
// withdraw reserve 90
// staked 890
// bonding 880
// unbonding 10
// after the unbonding period of user B withdrawal, unbonded token will go to withdraw reserve. ( for simplification, don't think the rebalancing now)
// withdraw reserve 100
// staked 880
// bonding 880
// unbonding 0
// Then User C execute MsgWithdraw with amount 10.
// maintenance rate is max(0, 100 - 10) / (100 + 0) = 90 / 100 = 0.9
