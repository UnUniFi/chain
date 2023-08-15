package keeper_test

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	nftkeeper "github.com/UnUniFi/chain/x/nft/keeper"

	"github.com/UnUniFi/chain/x/nftfactory/keeper"

	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/cosmos/cosmos-sdk/baseapp"

	simapp "github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/nftfactory/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx           sdk.Context
	app           *simapp.App
	addrs         []sdk.AccAddress
	queryClient   nft.QueryClient
	keeper        keeper.Keeper
	nftKeeper     nftkeeper.Keeper
	accountKeeper authkeeper.AccountKeeper
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false

	app := simapp.Setup(suite.T(), ([]wasm.Option{})...)

	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.app = app
	suite.addrs = simapp.AddTestAddrsIncremental(app, suite.ctx, 3, sdk.NewInt(30000000))

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	nft.RegisterQueryServer(queryHelper, app.UnUniFiNFTKeeper)
	suite.queryClient = nft.NewQueryClient(queryHelper)

	suite.keeper = app.NftfactoryKeeper
	suite.nftKeeper = app.UnUniFiNFTKeeper
	suite.accountKeeper = app.AccountKeeper
	params := types.DefaultParams()
	params.FeeCollectorAddress = suite.addrs[0].String()
	app.NftfactoryKeeper.SetParams(suite.ctx, &params)
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
