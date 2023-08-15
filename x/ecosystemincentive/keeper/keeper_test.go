package keeper_test

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/stretchr/testify/suite"

	simapp "github.com/UnUniFi/chain/app"
	nftbackedloankeeper "github.com/UnUniFi/chain/x/nftbackedloan/keeper"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

var (
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:    nil,
		distrtypes.ModuleName:         nil,
		minttypes.ModuleName:          {authtypes.Minter},
		nft.ModuleName:                nil,
		nftbackedloantypes.ModuleName: nil,
		// nftbackedloantypes.NftTradingFee: nil,
	}
)

type KeeperTestSuite struct {
	suite.Suite

	ctx                 sdk.Context
	app                 *simapp.App
	addrs               []sdk.AccAddress
	nftbackedloanKeeper nftbackedloankeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false

	app := simapp.Setup(suite.T(), ([]wasm.Option{})...)

	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.app = app
	suite.addrs = simapp.AddTestAddrsIncremental(app, suite.ctx, 3, sdk.NewInt(30000000))
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
