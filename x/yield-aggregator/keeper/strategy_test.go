package keeper_test

import (
	"time"

	"github.com/UnUniFi/chain/testutil/nullify"
	epochtypes "github.com/UnUniFi/chain/x/epochs/types"
	recordstypes "github.com/UnUniFi/chain/x/records/types"
	stakeibctypes "github.com/UnUniFi/chain/x/stakeibc/types"
	"github.com/UnUniFi/chain/x/yield-aggregator/keeper"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	transfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func createNStrategy(keeper *keeper.Keeper, ctx sdk.Context, denom string, n int) []types.Strategy {
	items := make([]types.Strategy, n)
	for i := range items {
		items[i] = types.Strategy{
			Denom:           denom,
			ContractAddress: "",
			Name:            "",
		}
		items[i].Id = keeper.AppendStrategy(ctx, denom, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestStrategyGet() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNStrategy(&keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		got, found := keeper.GetStrategy(ctx, vaultDenom, item.Id)
		suite.Require().True(found)
		suite.Require().Equal(
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func (suite *KeeperTestSuite) TestStrategyRemove() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNStrategy(&keeper, ctx, vaultDenom, 10)
	for _, item := range items {
		keeper.RemoveStrategy(ctx, vaultDenom, item.Id)
		_, found := keeper.GetStrategy(ctx, vaultDenom, item.Id)
		suite.Require().False(found)
	}
}

func (suite *KeeperTestSuite) TestStrategyGetAll() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNStrategy(&keeper, ctx, vaultDenom, 10)
	suite.Require().ElementsMatch(
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStrategy(ctx, vaultDenom)),
	)
}

func (suite *KeeperTestSuite) TestStrategyCount() {
	keeper, ctx := suite.app.YieldaggregatorKeeper, suite.ctx
	vaultDenom := "uatom"
	items := createNStrategy(&keeper, ctx, vaultDenom, 10)
	count := uint64(len(items))
	suite.Require().Equal(count, keeper.GetStrategyCount(ctx, vaultDenom))
}

func (suite *KeeperTestSuite) SetupZoneAndEpoch(hostDenom, ibcDenom string) stakeibctypes.HostZone {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now()

	// set host zone for env
	zone := stakeibctypes.HostZone{
		ChainId:               "hub-1",
		ConnectionId:          "connection-0",
		Bech32Prefix:          "cosmos",
		TransferChannelId:     "channel-0",
		Validators:            []*stakeibctypes.Validator{},
		BlacklistedValidators: []*stakeibctypes.Validator{},
		WithdrawalAccount:     nil,
		FeeAccount:            nil,
		DelegationAccount:     nil,
		RedemptionAccount:     nil,
		IBCDenom:              ibcDenom,
		HostDenom:             hostDenom,
		RedemptionRate:        sdk.NewDec(1),
		Address:               addr1.String(),
		StakedBal:             1000_000,
	}

	// set epoch tracker for env
	suite.app.StakeibcKeeper.SetEpochTracker(suite.ctx, stakeibctypes.EpochTracker{
		EpochIdentifier:    epochtypes.BASE_EPOCH,
		EpochNumber:        1,
		NextEpochStartTime: uint64(now.Unix()),
		Duration:           43200,
	})

	suite.app.StakeibcKeeper.SetEpochTracker(suite.ctx, stakeibctypes.EpochTracker{
		EpochIdentifier:    "day",
		EpochNumber:        1,
		NextEpochStartTime: uint64(now.Unix()),
		Duration:           86400,
	})

	suite.app.RecordsKeeper.SetEpochUnbondingRecord(suite.ctx, recordstypes.EpochUnbondingRecord{
		EpochNumber: 1,
		HostZoneUnbondings: []*recordstypes.HostZoneUnbonding{
			{
				StTokenAmount:         0,
				NativeTokenAmount:     0,
				Denom:                 zone.HostDenom,
				HostZoneId:            zone.ChainId,
				UnbondingTime:         0,
				Status:                recordstypes.HostZoneUnbonding_BONDED,
				UserRedemptionRecords: []string{},
			},
		},
	})

	// set deposit record for env
	suite.app.RecordsKeeper.SetDepositRecord(suite.ctx, recordstypes.DepositRecord{
		Id:                 1,
		Amount:             100,
		Denom:              ibcDenom,
		HostZoneId:         "hub-1",
		Status:             recordstypes.DepositRecord_STAKE,
		DepositEpochNumber: 1,
		Source:             recordstypes.DepositRecord_STRIDE,
	})
	suite.app.StakeibcKeeper.SetHostZone(suite.ctx, zone)
	return zone
}

// stake into strategy
func (suite *KeeperTestSuite) TestStakeToStrategy() {
	suite.SetupTest()
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	atomHostDenom := "uatom"
	prefixedDenom := transfertypes.GetPrefixedDenom("transfer", "channel-0", atomHostDenom)
	atomIbcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()
	zone := suite.SetupZoneAndEpoch(atomHostDenom, atomIbcDenom)

	strategy := types.Strategy{
		Id:              1,
		Name:            "AtomLiquidStaking",
		ContractAddress: "x/ibc-staking",
		Denom:           atomIbcDenom,
	}

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

	// mint coins to be spent on liquid staking
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	coins := sdk.Coins{sdk.NewInt64Coin(atomIbcDenom, 1000000)}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, vaultModAddr, coins)
	suite.Require().NoError(err)

	// stake to strategy - automatically calls liquid staking
	err = suite.app.YieldaggregatorKeeper.StakeToStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000))
	suite.Require().NoError(err)

	// check the changes
	record, found := suite.app.RecordsKeeper.GetDepositRecord(suite.ctx, 1)
	suite.Require().True(found)
	suite.Require().Equal(record.Amount, int64(1000_100))

	balance := suite.app.BankKeeper.GetAllBalances(suite.ctx, vaultModAddr)
	suite.Require().Equal(balance.String(), "1000000stuatom")

	bech32ZoneAddress, err := sdk.AccAddressFromBech32(zone.Address)
	suite.Require().NoError(err)
	balance = suite.app.BankKeeper.GetAllBalances(suite.ctx, bech32ZoneAddress)
	suite.Require().Equal(balance.String(), coins.String())
}

// unstake from strategy
func (suite *KeeperTestSuite) TestUnstakeFromStrategy() {
	suite.SetupTest()
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	atomHostDenom := "uatom"
	prefixedDenom := transfertypes.GetPrefixedDenom("transfer", "channel-0", atomHostDenom)
	atomIbcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()
	suite.SetupZoneAndEpoch(atomHostDenom, atomIbcDenom)

	strategy := types.Strategy{
		Id:              1,
		Name:            "AtomLiquidStaking",
		ContractAddress: "x/ibc-staking",
		Denom:           atomIbcDenom,
	}

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

	// mint coins to be spent on liquid staking
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	coins := sdk.Coins{sdk.NewInt64Coin(atomIbcDenom, 1000000)}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, vaultModAddr, coins)
	suite.Require().NoError(err)

	// stake to strategy - calls liquid staking
	err = suite.app.YieldaggregatorKeeper.StakeToStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000))
	suite.Require().NoError(err)

	// unstake from strategy - calls redeem stake
	err = suite.app.YieldaggregatorKeeper.UnstakeFromStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000))
	suite.Require().NoError(err)

	// check the changes
	unbondingRecord, found := suite.app.RecordsKeeper.GetEpochUnbondingRecord(suite.ctx, 1)
	suite.Require().True(found)
	suite.Require().Len(unbondingRecord.HostZoneUnbondings, 1)
	suite.Require().Equal(unbondingRecord.HostZoneUnbondings[0].StTokenAmount, uint64(1000_000))
	suite.Require().Equal(unbondingRecord.HostZoneUnbondings[0].Status, recordstypes.HostZoneUnbonding_BONDED)
	suite.Require().Equal(unbondingRecord.HostZoneUnbondings[0].NativeTokenAmount, uint64(1000_000))
	suite.Require().Len(unbondingRecord.HostZoneUnbondings[0].UserRedemptionRecords, 1)
}

// get amount put on the strategy
func (suite *KeeperTestSuite) TestGetAmountAndUnbondingAmountFromStrategy() {
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

	amount, err := suite.app.YieldaggregatorKeeper.GetAmountFromStrategy(suite.ctx, vault, strategy)
	suite.Require().NoError(err)
	suite.Require().Equal(amount.String(), "0"+atomIbcDenom)

	amount, err = suite.app.YieldaggregatorKeeper.GetUnbondingAmountFromStrategy(suite.ctx, vault, strategy)
	suite.Require().Error(err)

	// mint coins to be spent on liquid staking
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	coins := sdk.Coins{sdk.NewInt64Coin(atomIbcDenom, 1000000)}
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, vaultModAddr, coins)
	suite.Require().NoError(err)

	// stake to strategy - calls liquid staking
	suite.SetupZoneAndEpoch(atomHostDenom, atomIbcDenom)
	err = suite.app.YieldaggregatorKeeper.StakeToStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000))
	suite.Require().NoError(err)

	amount, err = suite.app.YieldaggregatorKeeper.GetAmountFromStrategy(suite.ctx, vault, strategy)
	suite.Require().NoError(err)
	suite.Require().Equal(amount.String(), "1000000"+atomIbcDenom)

	amount, err = suite.app.YieldaggregatorKeeper.GetUnbondingAmountFromStrategy(suite.ctx, vault, strategy)
	suite.Require().NoError(err)
	suite.Require().Equal(amount.String(), "0"+atomIbcDenom)

	// unstake from strategy - calls redeem stake
	err = suite.app.YieldaggregatorKeeper.UnstakeFromStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000))
	suite.Require().NoError(err)

	// check amounts after unstake
	amount, err = suite.app.YieldaggregatorKeeper.GetAmountFromStrategy(suite.ctx, vault, strategy)
	suite.Require().NoError(err)
	suite.Require().Equal(amount.String(), "0"+atomIbcDenom)

	amount, err = suite.app.YieldaggregatorKeeper.GetUnbondingAmountFromStrategy(suite.ctx, vault, strategy)
	suite.Require().NoError(err)
	suite.Require().Equal(amount.String(), "1000000"+atomIbcDenom)
}
