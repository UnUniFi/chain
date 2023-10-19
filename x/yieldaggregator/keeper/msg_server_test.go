package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (suite *KeeperTestSuite) TestMsgServerReinitVaultTransfer() {
	vault := types.Vault{
		Id:     1,
		Symbol: "ATOM",
		StrategyWeights: []types.StrategyWeight{
			{
				Denom:      "uatom1",
				StrategyId: 1,
				Weight:     sdk.OneDec(),
			},
			{
				Denom:      "uatom2",
				StrategyId: 1,
				Weight:     sdk.OneDec(),
			},
		},
	}

	contractAddr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	contractAddr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	suite.app.YieldaggregatorKeeper.SetVault(suite.ctx, vault)
	suite.app.YieldaggregatorKeeper.SetStrategy(suite.ctx, types.Strategy{
		Denom:           vault.StrategyWeights[0].Denom,
		Id:              vault.StrategyWeights[0].StrategyId,
		ContractAddress: contractAddr1.String(),
		Name:            "",
		Description:     "",
		GitUrl:          "",
	})
	suite.app.YieldaggregatorKeeper.SetStrategy(suite.ctx, types.Strategy{
		Denom:           vault.StrategyWeights[1].Denom,
		Id:              vault.StrategyWeights[1].StrategyId,
		ContractAddress: contractAddr2.String(),
		Name:            "",
		Description:     "",
		GitUrl:          "",
	})
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)

	coins := sdk.Coins{
		sdk.NewInt64Coin("uatom1", 1000000),
		sdk.NewInt64Coin("uatom2", 1000000),
	}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, vaultModAddr, coins)
	suite.Require().NoError(err)

	suite.app.YieldaggregatorKeeper.SetDenomInfo(suite.ctx, types.DenomInfo{
		Denom:  "uatom1",
		Symbol: "ATOM",
		Channels: []types.TransferChannel{
			{
				SendChainId: "cosmoshub-4",
				RecvChainId: "osmosis-1",
				ChannelId:   "channel-1",
			},
		},
	})
	suite.app.YieldaggregatorKeeper.SetDenomInfo(suite.ctx, types.DenomInfo{
		Denom:  "uatom2",
		Symbol: "ATOM",
		Channels: []types.TransferChannel{
			{
				SendChainId: "cosmoshub-4",
				RecvChainId: "neutron-1",
				ChannelId:   "channel-2",
			},
		},
	})

	suite.app.IBCKeeper.ChannelKeeper.SetNextSequenceSend(suite.ctx, transfertypes.ModuleName, "channel-1", 1)
	msgServer := keeper.NewMsgServerImpl(suite.app.YieldaggregatorKeeper)
	_, err = msgServer.ReinitVaultTransfer(sdk.WrapSDKContext(suite.ctx), &types.MsgReinitVaultTransfer{
		Sender:        authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		VaultId:       vault.Id,
		StrategyDenom: vault.StrategyWeights[0].Denom,
		StrategyId:    vault.StrategyWeights[0].StrategyId,
		Amount:        sdk.NewInt64Coin("uatom1", 1000),
	})
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "channel not found")
}
