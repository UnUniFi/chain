package keeper_test

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"

	sdknftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"

	nftkeeper "github.com/UnUniFi/chain/x/nft/keeper"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	simapp "github.com/UnUniFi/chain/app"
	appparams "github.com/UnUniFi/chain/app/params"
	"github.com/UnUniFi/chain/x/nftbackedloan/keeper"
	"github.com/UnUniFi/chain/x/nftbackedloan/types"
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
	types.RegisterQueryServer(queryHelper, app.NftbackedloanKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	encodingConfig := appparams.MakeEncodingConfig()
	appCodec := encodingConfig.Codec

	accountKeeper := authkeeper.NewAccountKeeper(
		appCodec,
		app.GetKey(authtypes.StoreKey),
		authtypes.ProtoBaseAccount,
		maccPerms,
		"ununifi",
		authtypes.NewModuleAddress("gov").String(),
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		app.GetKey(banktypes.StoreKey),
		app.AccountKeeper,
		app.BlockedAddrs(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	nftKeeper := nftkeeper.NewKeeper(sdknftkeeper.NewKeeper(app.GetKey(nft.StoreKey), appCodec, accountKeeper, bankKeeper), appCodec, authtypes.NewModuleAddress(govtypes.ModuleName).String())
	keeper := keeper.NewKeeper(appCodec, app.GetKey(types.StoreKey), app.GetKey(types.MemStoreKey), accountKeeper, bankKeeper, nftKeeper, authtypes.NewModuleAddress(govtypes.ModuleName).String())
	hooks := dummyNftmarketHook{}
	keeper.SetHooks(&hooks)
	suite.nftKeeper = nftKeeper
	suite.keeper = keeper
}
func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
