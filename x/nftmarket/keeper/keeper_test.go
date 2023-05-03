package keeper_test

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"

	simapp "github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/nftmarket/keeper"
	"github.com/UnUniFi/chain/x/nftmarket/types"
)

var (
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName: nil,
		distrtypes.ModuleName:      nil,
		minttypes.ModuleName:       {authtypes.Minter},
		nft.ModuleName:             nil,
		types.ModuleName:           nil,
		// types.NftTradingFee:        nil,
	}
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	app         *simapp.App
	addrs       []sdk.AccAddress
	queryClient types.QueryClient
	keeper      keeper.Keeper
	nftKeeper   nftkeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false

	app := simapp.Setup(suite.T(), ([]wasm.Option{})...)

	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.addrs = simapp.AddTestAddrsIncremental(app, suite.ctx, 3, sdk.NewInt(30000000))
	suite.app = app
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.NftmarketKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	hooks := dummyNftmarketHook{}
	app.NftmarketKeeper.SetHooks(&hooks)
	suite.nftKeeper = app.NFTKeeper
	suite.keeper = app.NftmarketKeeper
}
func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
