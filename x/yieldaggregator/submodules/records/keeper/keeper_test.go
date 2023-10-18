package keeper_test

import (
	"testing"

	"github.com/cometbft/cometbft/crypto/ed25519"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/CosmWasm/wasmd/x/wasm"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"

	simapp "github.com/UnUniFi/chain/app"
	icacallbackstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/icacallbacks/types"
	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/keeper"
	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *simapp.App
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false
	app := simapp.Setup(suite.T(), ([]wasm.Option{})...)
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.app = app
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestVaultTransfer() {
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
	transferCallback := types.VaultTransferCallback{
		VaultId:          1,
		StrategyContract: contractAddr.String(),
	}
	marshalledCallbackArgs, err := proto.Marshal(&transferCallback)
	suite.Require().NoError(err)

	callbacks := suite.app.IcacallbacksKeeper.GetAllCallbackData(suite.ctx)
	suite.Require().Len(callbacks, 1)
	suite.Require().Equal(callbacks[0], icacallbackstypes.CallbackData{
		CallbackKey:  icacallbackstypes.PacketID(transfertypes.ModuleName, "channel-0", 1),
		PortId:       transfertypes.ModuleName,
		ChannelId:    "channel-0",
		Sequence:     1,
		CallbackId:   keeper.CONTRACT_TRANSFER,
		CallbackArgs: marshalledCallbackArgs,
	})
}
