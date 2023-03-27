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

var (
	TestBaseTokenDenom  = "uatom"
	TestQuoteTokenDenom = "uusdc"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *simapp.App
	// addrs           []sdk.AccAddress
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
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    TestBaseTokenDenom,
				Exponent: 6,
			},
		},
		Base:   TestBaseTokenDenom,
		Symbol: TestBaseTokenDenom,
	}

	metadataUsdc := banktypes.Metadata{
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    TestQuoteTokenDenom,
				Exponent: 6,
			},
		},
		Base:   TestQuoteTokenDenom,
		Symbol: TestQuoteTokenDenom,
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
			{MarketId: "uusdc:usd", BaseAsset: TestQuoteTokenDenom, QuoteAsset: TestQuoteTokenDenom, Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			{MarketId: "uatom:usd", BaseAsset: TestBaseTokenDenom, QuoteAsset: TestQuoteTokenDenom, Oracles: []ununifitypes.StringAccAddress{}, Active: true},
		},
	}
	pricefeedKeeper.SetParams(suite.ctx, pfParams)

	_, _ = pricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.MustNewDecFromStr("0.00001"), suite.ctx.BlockTime().Add(1*time.Hour))
	_, _ = pricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uusdc:usd", sdk.MustNewDecFromStr("0.000001"), suite.ctx.BlockTime().Add(1*time.Hour))

	_ = pricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	_ = pricefeedKeeper.SetCurrentPrices(suite.ctx, "uusdc:usd")

	keeper := keeper.NewKeeper(appCodec, app.GetKey(types.StoreKey), app.GetKey(types.MemStoreKey), suite.app.GetSubspace(types.ModuleName), bankKeeper, pricefeedKeeper)

	params := types.DefaultParams()
	params.PoolParams.AcceptedAssets = []*types.PoolParams_Asset{
		{Denom: "uatom", TargetWeight: sdk.MustNewDecFromStr("0.5")},
		{Denom: "uusdc", TargetWeight: sdk.MustNewDecFromStr("0.5")},
	}
	params.PerpetualFutures.Markets = []*types.Market{
		{BaseDenom: TestBaseTokenDenom, QuoteDenom: TestQuoteTokenDenom},
	}

	keeper.SetParams(suite.ctx, params)

	for _, asset := range params.PoolParams.AcceptedAssets {
		keeper.AddPoolAsset(suite.ctx, *asset)
	}
	suite.keeper = keeper
	suite.pricefeedKeeper = pricefeedKeeper
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
