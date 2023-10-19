package keeper_test

import (
	"encoding/json"
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	"github.com/UnUniFi/chain/testutil/nullify"
	epochtypes "github.com/UnUniFi/chain/x/epochs/types"
	"github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	recordstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
	stakeibctypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/types"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
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
				StTokenAmount:         sdk.ZeroInt(),
				NativeTokenAmount:     sdk.ZeroInt(),
				Denom:                 zone.HostDenom,
				HostZoneId:            zone.ChainId,
				UnbondingTime:         0,
				Status:                recordstypes.HostZoneUnbonding_UNBONDING_QUEUE,
				UserRedemptionRecords: []string{},
			},
		},
	})

	// set deposit record for env
	suite.app.RecordsKeeper.SetDepositRecord(suite.ctx, recordstypes.DepositRecord{
		Id:                 1,
		Amount:             sdk.NewInt(100),
		Denom:              ibcDenom,
		HostZoneId:         "hub-1",
		Status:             recordstypes.DepositRecord_DELEGATION_QUEUE,
		DepositEpochNumber: 1,
		Source:             recordstypes.DepositRecord_UNUNIFI,
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
		Symbol:                 "ATOM",
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.ZeroDec(),
		WithdrawReserveRate:    sdk.ZeroDec(),
		StrategyWeights: []types.StrategyWeight{
			{Denom: atomIbcDenom, StrategyId: 1, Weight: sdk.OneDec()},
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
		Symbol:                 "ATOM",
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.ZeroDec(),
		WithdrawReserveRate:    sdk.ZeroDec(),
		StrategyWeights: []types.StrategyWeight{
			{Denom: atomIbcDenom, StrategyId: 1, Weight: sdk.OneDec()},
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
	err = suite.app.YieldaggregatorKeeper.UnstakeFromStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000), "")
	suite.Require().NoError(err)

	// check the changes
	unbondingRecord, found := suite.app.RecordsKeeper.GetEpochUnbondingRecord(suite.ctx, 1)
	suite.Require().True(found)
	suite.Require().Len(unbondingRecord.HostZoneUnbondings, 1)
	suite.Require().Equal(unbondingRecord.HostZoneUnbondings[0].StTokenAmount, uint64(1000_000))
	suite.Require().Equal(unbondingRecord.HostZoneUnbondings[0].Status, recordstypes.HostZoneUnbonding_UNBONDING_QUEUE)
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
		Symbol:                 "ATOM",
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.ZeroDec(),
		WithdrawReserveRate:    sdk.ZeroDec(),
		StrategyWeights: []types.StrategyWeight{
			{Denom: atomIbcDenom, StrategyId: 1, Weight: sdk.OneDec()},
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
	err = suite.app.YieldaggregatorKeeper.UnstakeFromStrategy(suite.ctx, vault, strategy, sdk.NewInt(1000_000), "")
	suite.Require().NoError(err)

	// check amounts after unstake
	amount, err = suite.app.YieldaggregatorKeeper.GetAmountFromStrategy(suite.ctx, vault, strategy)
	suite.Require().NoError(err)
	suite.Require().Equal(amount.String(), "0"+atomIbcDenom)

	amount, err = suite.app.YieldaggregatorKeeper.GetUnbondingAmountFromStrategy(suite.ctx, vault, strategy)
	suite.Require().NoError(err)
	suite.Require().Equal(amount.String(), "1000000"+atomIbcDenom)
}

func (suite *KeeperTestSuite) TestCalculateTransferRoute() {
	currChannels := []types.TransferChannel{}
	tarChannels := []types.TransferChannel{}

	route := keeper.CalculateTransferRoute(currChannels, tarChannels)
	suite.Require().Equal(route, []types.TransferChannel{})

	// ATOM
	currChannels = []types.TransferChannel{
		{
			SendChainId: "ununifi-1",
			RecvChainId: "cosmoshub-1",
			ChannelId:   "channel-2", // back channel
		},
	}
	tarChannels = []types.TransferChannel{
		{
			SendChainId: "cosmoshub-1",
			RecvChainId: "osmosis-1",
			ChannelId:   "channel-1", // forward channel
		},
	}
	route = keeper.CalculateTransferRoute(currChannels, tarChannels)
	suite.Require().Equal(route, []types.TransferChannel{
		{
			SendChainId: "ununifi-1",
			RecvChainId: "cosmoshub-1",
			ChannelId:   "channel-2", // back channel
		},
		{
			SendChainId: "cosmoshub-1",
			RecvChainId: "osmosis-1",
			ChannelId:   "channel-1", // forward channel
		},
	})

	// ATOM.osmo
	currChannels = []types.TransferChannel{
		{
			SendChainId: "osmosis-1",
			RecvChainId: "cosmoshub-1",
			ChannelId:   "channel-3", // back channel
		},
		{
			SendChainId: "ununifi-1",
			RecvChainId: "osmosis-1",
			ChannelId:   "channel-3", // back channel
		},
	}
	tarChannels = []types.TransferChannel{
		{
			SendChainId: "cosmoshub-1",
			RecvChainId: "osmosis-1",
			ChannelId:   "channel-1", // forward channel
		},
	}
	route = keeper.CalculateTransferRoute(currChannels, tarChannels)
	suite.Require().Equal(route, []types.TransferChannel{
		{
			SendChainId: "ununifi-1",
			RecvChainId: "osmosis-1",
			ChannelId:   "channel-3", // back channel
		},
	})
}

func (suite *KeeperTestSuite) TestComposePacketForwardMetadata() {
	suite.app.YieldaggregatorKeeper.SetIntermediaryAccountInfo(suite.ctx, []types.ChainAddress{
		{
			ChainId: "cosmoshub-1",
			Address: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
		},
		{
			ChainId: "osmosis-1",
			Address: "osmo1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
		},
		{
			ChainId: "neutron-1",
			Address: "neutron1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
		},
	})

	receiver, metadata := suite.app.YieldaggregatorKeeper.ComposePacketForwardMetadata(suite.ctx, []types.TransferChannel{
		{
			SendChainId: "ununifi-1",
			RecvChainId: "osmosis-1",
			ChannelId:   "channel-3", // back channel
		},
	}, "osmo1aqvlxpk8dc4m2nkmxkf63a5zez9jkzgm6amkgddhfk0qj9j4rw3q662wuk")
	suite.Require().Equal(receiver, "osmo1aqvlxpk8dc4m2nkmxkf63a5zez9jkzgm6amkgddhfk0qj9j4rw3q662wuk")
	metadataJson, err := json.Marshal(metadata)
	suite.Require().NoError(err)
	suite.Require().Equal(string(metadataJson), "null")

	receiver, metadata = suite.app.YieldaggregatorKeeper.ComposePacketForwardMetadata(suite.ctx, []types.TransferChannel{
		{
			SendChainId: "ununifi-1",
			RecvChainId: "cosmoshub-1",
			ChannelId:   "channel-2", // back channel
		},
		{
			SendChainId: "cosmoshub-1",
			RecvChainId: "osmosis-1",
			ChannelId:   "channel-1", // forward channel
		},
	}, "osmo1aqvlxpk8dc4m2nkmxkf63a5zez9jkzgm6amkgddhfk0qj9j4rw3q662wuk")
	suite.Require().Equal(receiver, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t")
	metadataJson, err = json.Marshal(metadata)
	suite.Require().NoError(err)
	suite.Require().Equal(string(metadataJson), `{"forward":{"receiver":"osmo1aqvlxpk8dc4m2nkmxkf63a5zez9jkzgm6amkgddhfk0qj9j4rw3q662wuk","port":"transfer","channel":"channel-1","retries":2,"next":null}}`)

	receiver, metadata = suite.app.YieldaggregatorKeeper.ComposePacketForwardMetadata(suite.ctx, []types.TransferChannel{
		{
			SendChainId: "ununifi-1",
			RecvChainId: "neutron-1",
			ChannelId:   "channel-2", // back channel
		},
		{
			SendChainId: "neutron-1",
			RecvChainId: "cosmoshub-1",
			ChannelId:   "channel-3", // back channel
		},
		{
			SendChainId: "cosmoshub-1",
			RecvChainId: "osmosis-1",
			ChannelId:   "channel-1", // forward channel
		},
	}, "osmo1aqvlxpk8dc4m2nkmxkf63a5zez9jkzgm6amkgddhfk0qj9j4rw3q662wuk")
	suite.Require().Equal(receiver, "neutron1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t")
	metadataJson, err = json.Marshal(metadata)
	suite.Require().NoError(err)
	suite.Require().Equal(string(metadataJson), `{"forward":{"receiver":"cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t","port":"transfer","channel":"channel-3","retries":2,"next":{"forward":{"receiver":"osmo1aqvlxpk8dc4m2nkmxkf63a5zez9jkzgm6amkgddhfk0qj9j4rw3q662wuk","port":"transfer","channel":"channel-1","retries":2,"next":null}}}}`)
}

func (suite *KeeperTestSuite) TestExecuteVaultTransfer() {
	suite.app.YieldaggregatorKeeper.SetIntermediaryAccountInfo(suite.ctx, []types.ChainAddress{
		{
			ChainId: "cosmoshub-1",
			Address: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
		},
		{
			ChainId: "osmosis-1",
			Address: "osmo1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
		},
		{
			ChainId: "neutron-1",
			Address: "neutron1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
		},
	})

	denomInfos := []types.DenomInfo{
		{
			Denom:  "ibc/01AAFF",
			Symbol: "ATOM",
			Channels: []types.TransferChannel{
				{
					RecvChainId: "cosmoshub-4",
					SendChainId: "ununifi-1",
					ChannelId:   "channel-2",
				},
			},
		},
	}

	symbolInfos := []types.SymbolInfo{
		{
			Symbol:        "ATOM",
			NativeChainId: "cosmoshub-4",
			Channels: []types.TransferChannel{
				{
					SendChainId: "cosmoshub-4",
					RecvChainId: "", // osmosis-1 (since contract is not mocked target chain id is set as "")
					ChannelId:   "channel-1",
				},
			},
		},
	}

	for _, denomInfo := range denomInfos {
		suite.app.YieldaggregatorKeeper.SetDenomInfo(suite.ctx, denomInfo)
	}

	for _, symbolInfo := range symbolInfos {
		suite.app.YieldaggregatorKeeper.SetSymbolInfo(suite.ctx, symbolInfo)
	}

	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	vault := types.Vault{
		Id:                     1,
		Symbol:                 "ATOM",
		Owner:                  addr1.String(),
		OwnerDeposit:           sdk.NewInt64Coin("uguu", 100),
		WithdrawCommissionRate: sdk.ZeroDec(),
		WithdrawReserveRate:    sdk.ZeroDec(),
		StrategyWeights: []types.StrategyWeight{
			{Denom: denomInfos[0].Denom, StrategyId: 1, Weight: sdk.OneDec()},
		},
	}

	contractAddr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	strategy := types.Strategy{
		Id:              1,
		Name:            "AtomStaking",
		ContractAddress: contractAddr.String(),
		Denom:           denomInfos[0].Denom,
	}

	coin := sdk.NewInt64Coin(denomInfos[0].Denom, 1000)

	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, vaultModAddr, sdk.Coins{coin})
	suite.Require().NoError(err)

	suite.app.IBCKeeper.ChannelKeeper.SetNextSequenceSend(suite.ctx, transfertypes.ModuleName, denomInfos[0].Channels[0].ChannelId, 1)
	msg, err := suite.app.YieldaggregatorKeeper.ExecuteVaultTransfer(suite.ctx, vault, strategy, coin)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "channel not found")

	msgBz, err := suite.app.AppCodec().MarshalJSON(msg)
	suite.Require().NoError(err)
	suite.Require().Equal(string(msgBz), `{"source_port":"transfer","source_channel":"channel-2","token":{"denom":"ibc/01AAFF","amount":"1000"},"sender":"cosmos1rrxq3fae4jksmnmtd0k5z744am9hp307m8t429","receiver":"","timeout_height":{"revision_number":"0","revision_height":"0"},"timeout_timestamp":"11651379494838206464","memo":"{\"forward\":{\"port\":\"transfer\",\"channel\":\"channel-1\",\"retries\":2,\"next\":null}}"}`)
}
