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

	metadataAtom := banktypes.Metadata{
		Description: "The native staking token of the Cosmos Hub.",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uatom",
				Exponent: 0,
				Aliases:  []string{"microatom"},
			},
			{
				Denom:    "atom",
				Exponent: 6,
				Aliases:  []string{"ATOM"},
			},
		},
		Base:    "uatom",
		Display: "atom",
		Symbol:  "ATOM",
	}

	metadataUsdc := banktypes.Metadata{
		Description: "USD Coin",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uusdc",
				Exponent: 0,
			},
			{
				Denom:    "usdc",
				Exponent: 18,
			},
		},
		Base:    "uusdc",
		Display: "usdc",
		Symbol:  "USDC",
	}

	bankKeeper.SetDenomMetaData(suite.ctx, metadataAtom)
	bankKeeper.SetDenomMetaData(suite.ctx, metadataUsdc)

	pricefeedKeeper := pricefeedkeeper.NewKeeper(
		appCodec,
		app.GetKey(pricefeedtypes.StoreKey),
		app.GetMemKey(pricefeedtypes.MemStoreKey),
		app.GetSubspace(pricefeedtypes.ModuleName),
		bankKeeper,
	)
	pfParams := pricefeedtypes.Params{
		Markets: []pricefeedtypes.Market{
			{MarketId: "USDC:USD", BaseAsset: "uusdc", QuoteAsset: "usd", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			{MarketId: "ATOM:USD", BaseAsset: "uatom", QuoteAsset: "usd", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			{MarketId: "ATOM:USDC", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
		},
	}
	pricefeedKeeper.SetParams(suite.ctx, pfParams)

	pricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "ATOM:USDC", sdk.MustNewDecFromStr("15.28"), suite.ctx.BlockTime().Add(1*time.Hour))

	pricefeedKeeper.SetCurrentPrices(suite.ctx, "ATOM:USDC")

	keeper := keeper.NewKeeper(appCodec, app.GetKey(types.StoreKey), app.GetKey(types.MemStoreKey), suite.app.GetSubspace(types.ModuleName), bankKeeper, pricefeedKeeper)
	suite.keeper = keeper
	suite.pricefeedKeeper = pricefeedKeeper
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
