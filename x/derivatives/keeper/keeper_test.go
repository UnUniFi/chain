package keeper_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	pricefeedkeeper "github.com/UnUniFi/chain/x/pricefeed/keeper"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"

	simapp "github.com/UnUniFi/chain/app"
	appparams "github.com/UnUniFi/chain/app/params"
	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/derivatives/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx             sdk.Context
	app             *simapp.App
	addrs           []sdk.AccAddress
	queryClient     types.QueryClient
	keeper          keeper.Keeper
	pricefeedKeeper pricefeedkeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false

	app := simapp.Setup(suite.T(), isCheckTx)

	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.app = app
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.DerivativesKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	encodingConfig := appparams.MakeEncodingConfig()
	appCodec := encodingConfig.Marshaler

	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		app.GetKey(banktypes.StoreKey),
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)

	pricefeedKeeper := pricefeedkeeper.NewKeeper(
		appCodec,
		app.GetKey(pricefeedtypes.StoreKey),
		app.GetMemKey(pricefeedtypes.MemStoreKey),
		app.GetSubspace(pricefeedtypes.ModuleName),
	)
	pfParams := pricefeedtypes.Params{
		Markets: []pricefeedtypes.Market{
			{MarketId: "btc:usdc", BaseAsset: "btc", QuoteAsset: "usdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			{MarketId: "usdc:usdc", BaseAsset: "usdc", QuoteAsset: "usdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			{MarketId: "bnb:usdc", BaseAsset: "bnb", QuoteAsset: "usdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			{MarketId: "bjpy:usdc", BaseAsset: "bjpy", QuoteAsset: "usdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
		},
		DenomTickerPairs: []pricefeedtypes.DenomTickerPair{
			{Denom: "btc", Ticker: "btc"},
			{Denom: "usdc", Ticker: "usdc"},
		},
	}
	pricefeedKeeper.SetParams(suite.ctx, pfParams)
	pricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "btc:usdc", sdk.MustNewDecFromStr("8000.00"), suite.ctx.BlockTime().Add(1*time.Hour))
	pricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "usdc:usdc", sdk.MustNewDecFromStr("1.00"), suite.ctx.BlockTime().Add(1*time.Hour))
	pricefeedKeeper.SetCurrentPrices(suite.ctx, "btc:usdc")
	pricefeedKeeper.SetCurrentPrices(suite.ctx, "usdc:usdc")
	keeper := keeper.NewKeeper(appCodec, app.GetKey(types.StoreKey), app.GetKey(types.MemStoreKey), suite.app.GetSubspace(types.ModuleName), bankKeeper, pricefeedKeeper)
	suite.keeper = keeper
	suite.pricefeedKeeper = pricefeedKeeper
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
