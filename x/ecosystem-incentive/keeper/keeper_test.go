package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	appparams "github.com/UnUniFi/chain/app/params"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"

	simapp "github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/ecosystem-incentive/keeper"
	nftmarketkeeper "github.com/UnUniFi/chain/x/nftmarket/keeper"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

var (
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:   nil,
		distrtypes.ModuleName:        nil,
		minttypes.ModuleName:         {authtypes.Minter},
		nft.ModuleName:               nil,
		nftmarkettypes.ModuleName:    nil,
		nftmarkettypes.NftTradingFee: nil,
	}
)

type KeeperTestSuite struct {
	suite.Suite

	ctx             sdk.Context
	app             *simapp.App
	addrs           []sdk.AccAddress
	nftmarketKeeper nftmarketkeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false

	app := simapp.Setup(suite.T(), isCheckTx)

	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.app = app
	suite.addrs = simapp.AddTestAddrsIncremental(app, suite.ctx, 3, sdk.NewInt(30000000))

	encodingConfig := appparams.MakeEncodingConfig()
	appCodec := encodingConfig.Marshaler

	txCfg := encodingConfig.TxConfig
	accountKeeper := authkeeper.NewAccountKeeper(
		appCodec, app.GetKey(authtypes.StoreKey), app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms, sdk.Bech32MainPrefix,
	)
	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		app.GetKey(banktypes.StoreKey),
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)
	nftKeeper := nftkeeper.NewKeeper(app.GetKey(nft.StoreKey), appCodec, accountKeeper, bankKeeper)
	nftmarketkeeper := nftmarketkeeper.NewKeeper(appCodec, txCfg, app.GetKey(nftmarkettypes.StoreKey), app.GetKey(nftmarkettypes.MemStoreKey), suite.app.GetSubspace(nftmarkettypes.ModuleName), accountKeeper, bankKeeper, nftKeeper)
	hooks := keeper.Hooks{}
	nftmarketkeeper.SetHooks(&hooks)
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
