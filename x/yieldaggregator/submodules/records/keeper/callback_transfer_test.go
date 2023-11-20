package keeper_test

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/gogo/protobuf/proto"

	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/keeper"
	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
)

func (suite *KeeperTestSuite) TestVaultTransferCallback() {
	vaultAddr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	recvAddr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	coin := sdk.NewInt64Coin("uatom1", 1000000)
	coins := sdk.Coins{coin}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, vaultAddr, coins)
	suite.Require().NoError(err)

	contractAddr := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	suite.app.IBCKeeper.ChannelKeeper.SetNextSequenceSend(suite.ctx, transfertypes.ModuleName, "channel-0", 1)

	timeoutTimestamp := uint64(suite.ctx.BlockTime().UnixNano()) + 60000000000
	err = suite.app.RecordsKeeper.VaultTransfer(suite.ctx, 1, contractAddr, transfertypes.NewMsgTransfer(
		transfertypes.PortID,
		"channel-0",
		coin,
		vaultAddr.String(),
		recvAddr.String(),
		clienttypes.Height{},
		timeoutTimestamp,
		"",
	))
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "channel not found")

	data := transfertypes.FungibleTokenPacketData{
		Denom:    "uatom1",
		Amount:   coin.Amount.String(),
		Receiver: recvAddr.String(),
		Sender:   vaultAddr.String(),
		Memo:     "",
	}
	ftpdBz, err := transfertypes.ModuleCdc.MarshalJSON(&data)
	suite.Require().NoError(err)

	callbackData := types.VaultTransferCallback{
		VaultId:          1,
		StrategyContract: contractAddr.String(),
	}
	callbackBz, err := proto.Marshal(&callbackData)
	suite.Require().NoError(err)

	err = keeper.VaultTransferCallback(suite.app.RecordsKeeper, suite.ctx, channeltypes.Packet{
		Data: ftpdBz,
	}, &channeltypes.Acknowledgement{}, callbackBz)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "failed to Sudo")
}
