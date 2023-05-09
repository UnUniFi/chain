package keeper_test

// import (
// 	"time"

// 	"github.com/cometbft/cometbft/crypto/ed25519"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

// 	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

// 	"github.com/UnUniFi/chain/x/deprecated/yieldaggregatorv1/types"
// 	epochtypes "github.com/UnUniFi/chain/x/epochs/types"
// 	recordstypes "github.com/UnUniFi/chain/x/records/types"
// 	stakeibctypes "github.com/UnUniFi/chain/x/stakeibc/types"
// 	yieldfarmtypes "github.com/UnUniFi/chain/x/yieldfarm/types"
// )

// func (suite *KeeperTestSuite) TestInvestOnTarget() {
// 	// preparation before investment
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	coins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 1000))
// 	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
// 	suite.NoError(err)
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
// 	suite.NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().NoError(err)

// 	// execute investment
// 	err = suite.app.YieldaggregatorKeeper.InvestOnTarget(suite.ctx, addr1, types.AssetManagementTarget{
// 		AssetManagementAccountId: "UnunifiFarm",
// 		Id:                       "GUUStaking",
// 		IntegrateInfo: types.IntegrateInfo{
// 			Type: types.IntegrateType_GOLANG_MOD,
// 		},
// 	}, coins)
// 	suite.Require().NoError(err)

// 	// farming unit creation or update check
// 	unit := suite.app.YieldaggregatorKeeper.GetFarmingUnit(suite.ctx, addr1.String(), "UnunifiFarm", "GUUStaking")
// 	expectedFarmingUnit := types.FarmingUnit{
// 		AccountId:          "UnunifiFarm",
// 		TargetId:           "GUUStaking",
// 		Amount:             coins,
// 		FarmingStartTime:   suite.ctx.BlockTime().String(),
// 		UnbondingStarttime: time.Time{},
// 		Owner:              addr1.String(),
// 	}
// 	suite.Require().Equal(unit, expectedFarmingUnit)

// 	// check token transfer from yield aggregator module to yieldfarm module by amount
// 	moduleAddr := suite.app.AccountKeeper.GetModuleAddress(yieldfarmtypes.ModuleName)
// 	balance := suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, "uguu")
// 	suite.Require().Equal(balance, coins[0])
// 	moduleAddr = suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
// 	balance = suite.app.BankKeeper.GetBalance(suite.ctx, moduleAddr, "uguu")
// 	suite.Require().Equal(balance, sdk.NewInt64Coin("uguu", 0))

// 	// check farming unit account deposit amount increase
// 	farmerInfo := suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, expectedFarmingUnit.GetAddress())
// 	suite.Require().Equal(farmerInfo, yieldfarmtypes.FarmerInfo{
// 		Account: expectedFarmingUnit.GetAddress().String(),
// 		Amount:  coins,
// 		Rewards: sdk.Coins(nil),
// 	})
// }

// func (suite *KeeperTestSuite) TestBeginWithdrawFromTarget() {
// 	// try withdrawal when farming unit does not exist
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	coins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 1000))
// 	assetTarget := types.AssetManagementTarget{
// 		AssetManagementAccountId: "UnunifiFarm",
// 		Id:                       "GUUStaking",
// 		IntegrateInfo: types.IntegrateInfo{
// 			Type: types.IntegrateType_GOLANG_MOD,
// 		},
// 	}
// 	err := suite.app.YieldaggregatorKeeper.BeginWithdrawFromTarget(suite.ctx, addr1, assetTarget, coins)
// 	suite.Require().Error(err)

// 	// preparation for withdrawal
// 	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
// 	suite.NoError(err)
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
// 	suite.NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.InvestOnTarget(suite.ctx, addr1, assetTarget, coins)
// 	suite.Require().NoError(err)

// 	// withdraw partial amount
// 	partialCoins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 100))
// 	err = suite.app.YieldaggregatorKeeper.BeginWithdrawFromTarget(suite.ctx, addr1, assetTarget, partialCoins)
// 	suite.Require().NoError(err)

// 	// check farmerInfo change
// 	unit := suite.app.YieldaggregatorKeeper.GetFarmingUnit(suite.ctx, addr1.String(), "UnunifiFarm", "GUUStaking")
// 	farmerInfo := suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, unit.GetAddress())
// 	suite.Require().Equal(farmerInfo, yieldfarmtypes.FarmerInfo{
// 		Account: unit.GetAddress().String(),
// 		Amount:  coins.Sub(partialCoins),
// 		Rewards: sdk.Coins(nil),
// 	})

// 	// withdraw full amount
// 	err = suite.app.YieldaggregatorKeeper.BeginWithdrawFromTarget(suite.ctx, addr1, assetTarget, sdk.Coins{sdk.NewInt64Coin("uguu", 900)})
// 	suite.Require().NoError(err)

// 	// check farmerInfo change
// 	farmerInfo = suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, unit.GetAddress())
// 	suite.Require().Equal(farmerInfo, yieldfarmtypes.FarmerInfo{
// 		Account: unit.GetAddress().String(),
// 		Amount:  sdk.Coins(nil),
// 		Rewards: sdk.Coins(nil),
// 	})
// }

// func (suite *KeeperTestSuite) TestClaimWithdrawFromTarget() {
// 	now := time.Now().UTC()
// 	suite.ctx = suite.ctx.WithBlockTime(now)
// 	// try claim withdraw when farming unit does not exist
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	coins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 1000))
// 	assetTarget := types.AssetManagementTarget{
// 		AssetManagementAccountId: "UnunifiFarm",
// 		Id:                       "GUUStaking",
// 		IntegrateInfo: types.IntegrateInfo{
// 			Type: types.IntegrateType_GOLANG_MOD,
// 		},
// 		UnbondingTime: time.Hour,
// 	}
// 	err := suite.app.YieldaggregatorKeeper.ClaimWithdrawFromTarget(suite.ctx, addr1, assetTarget)
// 	suite.Require().Error(err)

// 	// try claim before unbonding time pass
// 	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
// 	suite.NoError(err)
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
// 	suite.NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.InvestOnTarget(suite.ctx, addr1, assetTarget, coins)
// 	suite.Require().NoError(err)
// 	suite.Require().NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.BeginWithdrawFromTarget(suite.ctx, addr1, assetTarget, sdk.Coins{})
// 	suite.Require().NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.ClaimWithdrawFromTarget(suite.ctx, addr1, assetTarget)
// 	suite.Require().Error(err)

// 	// claim after unbonding time pass
// 	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Hour))
// 	err = suite.app.YieldaggregatorKeeper.ClaimWithdrawFromTarget(suite.ctx, addr1, assetTarget)
// 	suite.Require().NoError(err)

// 	// check user deposit increase
// 	deposit := suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr1)
// 	suite.Require().True(deposit.IsAllGTE(coins))

// 	// check farmerInfo zero
// 	unit := suite.app.YieldaggregatorKeeper.GetFarmingUnit(suite.ctx, addr1.String(), "UnunifiFarm", "GUUStaking")
// 	farmerInfo := suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, unit.GetAddress())
// 	suite.Require().Equal(farmerInfo, yieldfarmtypes.FarmerInfo{
// 		Account: unit.GetAddress().String(),
// 		Amount:  sdk.Coins(nil),
// 		Rewards: sdk.Coins(nil),
// 	})
// }

// func (suite *KeeperTestSuite) TestClaimRewardsFromTarget() {
// 	now := time.Now().UTC()
// 	suite.ctx = suite.ctx.WithBlockTime(now)
// 	// try claim withdraw when farming unit does not exist
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	coins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 1000))
// 	assetTarget := types.AssetManagementTarget{
// 		AssetManagementAccountId: "UnunifiFarm",
// 		Id:                       "GUUStaking",
// 		IntegrateInfo: types.IntegrateInfo{
// 			Type: types.IntegrateType_GOLANG_MOD,
// 		},
// 		UnbondingTime: time.Hour,
// 	}
// 	err := suite.app.YieldaggregatorKeeper.ClaimRewardsFromTarget(suite.ctx, addr1, assetTarget)
// 	suite.Require().Error(err)

// 	// preparation
// 	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
// 	suite.NoError(err)
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
// 	suite.NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.InvestOnTarget(suite.ctx, addr1, assetTarget, coins)
// 	suite.Require().NoError(err)
// 	suite.Require().NoError(err)

// 	// claim after some time
// 	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Hour))
// 	err = suite.app.YieldaggregatorKeeper.ClaimRewardsFromTarget(suite.ctx, addr1, assetTarget)
// 	suite.Require().NoError(err)

// 	// check user deposit increase
// 	deposit := suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr1)
// 	suite.Require().Equal(deposit, sdk.Coins{sdk.NewInt64Coin("uguu", 1000)})

// 	// check farmerInfo stays same
// 	unit := suite.app.YieldaggregatorKeeper.GetFarmingUnit(suite.ctx, addr1.String(), "UnunifiFarm", "GUUStaking")
// 	farmerInfo := suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, unit.GetAddress())
// 	suite.Require().Equal(farmerInfo, yieldfarmtypes.FarmerInfo{
// 		Account: unit.GetAddress().String(),
// 		Amount:  coins,
// 		Rewards: sdk.Coins(nil),
// 	})
// }

// func (suite *KeeperTestSuite) TestClaimAllFarmUnitRewards() {
// 	now := time.Now().UTC()
// 	suite.ctx = suite.ctx.WithBlockTime(now)
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	coins := sdk.NewCoins(sdk.NewInt64Coin("uguu", 1000))
// 	assetTarget := types.AssetManagementTarget{
// 		AssetManagementAccountId: "UnunifiFarm",
// 		Id:                       "GUUStaking",
// 		IntegrateInfo: types.IntegrateInfo{
// 			Type: types.IntegrateType_GOLANG_MOD,
// 		},
// 		UnbondingTime: time.Hour,
// 	}

// 	// preparation
// 	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
// 	suite.NoError(err)
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
// 	suite.NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.InvestOnTarget(suite.ctx, addr1, assetTarget, coins)
// 	suite.Require().NoError(err)
// 	suite.Require().NoError(err)

// 	// claim after some time
// 	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Hour))
// 	suite.app.YieldaggregatorKeeper.ClaimAllFarmUnitRewards(suite.ctx)

// 	// check user deposit increase
// 	deposit := suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, addr1)
// 	suite.Require().Equal(deposit, sdk.Coins{sdk.NewInt64Coin("uguu", 1000)})

// 	// check farmerInfo stays same
// 	unit := suite.app.YieldaggregatorKeeper.GetFarmingUnit(suite.ctx, addr1.String(), "UnunifiFarm", "GUUStaking")
// 	farmerInfo := suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, unit.GetAddress())
// 	suite.Require().Equal(farmerInfo, yieldfarmtypes.FarmerInfo{
// 		Account: unit.GetAddress().String(),
// 		Amount:  coins,
// 		Rewards: sdk.Coins(nil),
// 	})
// }

// func (suite *KeeperTestSuite) TestClaimAllFarmUnitRewardsIBCStake() {
// 	now := time.Now()

// 	suite.SetupTest()
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

// 	// set epoch tracker for env
// 	suite.app.StakeibcKeeper.SetEpochTracker(suite.ctx, stakeibctypes.EpochTracker{
// 		EpochIdentifier:    epochtypes.BASE_EPOCH,
// 		EpochNumber:        1,
// 		NextEpochStartTime: uint64(now.Unix()),
// 		Duration:           43200,
// 	})

// 	atomHostDenom := "uatom"
// 	prefixedDenom := transfertypes.GetPrefixedDenom("transfer", "channel-0", atomHostDenom)
// 	atomIbcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()

// 	// set deposit record for env
// 	suite.app.RecordsKeeper.SetDepositRecord(suite.ctx, recordstypes.DepositRecord{
// 		Id:                 1,
// 		Amount:             100,
// 		Denom:              atomIbcDenom,
// 		HostZoneId:         "hub-1",
// 		Status:             recordstypes.DepositRecord_STAKE,
// 		DepositEpochNumber: 1,
// 		Source:             recordstypes.DepositRecord_STRIDE,
// 	})
// 	// set host zone for env
// 	zone := stakeibctypes.HostZone{
// 		ChainId:               "hub-1",
// 		ConnectionId:          "connection-0",
// 		Bech32Prefix:          "cosmos",
// 		TransferChannelId:     "channel-0",
// 		Validators:            []*stakeibctypes.Validator{},
// 		BlacklistedValidators: []*stakeibctypes.Validator{},
// 		WithdrawalAccount:     nil,
// 		FeeAccount:            nil,
// 		DelegationAccount:     nil,
// 		RedemptionAccount:     nil,
// 		IBCDenom:              atomIbcDenom,
// 		HostDenom:             atomHostDenom,
// 		RedemptionRate:        sdk.NewDec(1),
// 		Address:               addr1.String(),
// 	}
// 	suite.app.StakeibcKeeper.SetHostZone(suite.ctx, zone)

// 	// mint coins to be spent on liquid staking
// 	coins := sdk.Coins{sdk.NewInt64Coin(atomIbcDenom, 1000000)}
// 	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
// 	suite.Require().NoError(err)
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, addr1, coins)
// 	suite.Require().NoError(err)

// 	assetTarget := types.AssetManagementTarget{
// 		AssetManagementAccountId: "AtomFarm",
// 		Id:                       "AtomLiquidStaking",
// 		IntegrateInfo: types.IntegrateInfo{
// 			Type:    types.IntegrateType_GOLANG_MOD,
// 			ModName: "stakeibc",
// 		},
// 		UnbondingTime: time.Hour,
// 	}

// 	// preparation
// 	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress: addr1.Bytes(),
// 		Amount:      coins,
// 	})
// 	suite.Require().NoError(err)
// 	suite.app.YieldaggregatorKeeper.SetAssetManagementTarget(suite.ctx, assetTarget)
// 	err = suite.app.YieldaggregatorKeeper.InvestOnTarget(suite.ctx, addr1, assetTarget, coins)
// 	suite.Require().NoError(err)

// 	farmUnits := suite.app.YieldaggregatorKeeper.GetFarmingUnitsOfAddress(suite.ctx, addr1)
// 	suite.Require().Equal(len(farmUnits), 1)
// 	suite.Require().Equal(sdk.Coins(farmUnits[0].Amount).String(), "1000000"+atomIbcDenom)

// 	// claim after some time
// 	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Hour))
// 	zone.RedemptionRate = zone.RedemptionRate.Mul(sdk.NewDec(2)) // 2x redemption rate
// 	suite.app.StakeibcKeeper.SetHostZone(suite.ctx, zone)

// 	suite.app.YieldaggregatorKeeper.ClaimAllFarmUnitRewards(suite.ctx)

// 	// check after claim
// 	farmUnits = suite.app.YieldaggregatorKeeper.GetFarmingUnitsOfAddress(suite.ctx, addr1)
// 	suite.Require().Equal(len(farmUnits), 1)
// 	suite.Require().Equal(sdk.Coins(farmUnits[0].Amount).String(), "2000000"+atomIbcDenom)
// }
