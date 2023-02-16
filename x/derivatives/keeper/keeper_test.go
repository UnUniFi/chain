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
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uatom",
				Exponent: 6,
			},
		},
		Base:   "uatom",
		Symbol: "uatom",
	}

	metadataUsdc := banktypes.Metadata{
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uusdc",
				Exponent: 6,
			},
		},
		Base:   "uusdc",
		Symbol: "uusdc",
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
			{MarketId: "uusdc:usd", BaseAsset: "uusdc", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			{MarketId: "uatom:usd", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
			{MarketId: "uatom:usdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
		},
	}
	pricefeedKeeper.SetParams(suite.ctx, pfParams)

	pricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usdc", sdk.MustNewDecFromStr("0.00001528"), suite.ctx.BlockTime().Add(1*time.Hour))
	pricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.MustNewDecFromStr("0.00001528"), suite.ctx.BlockTime().Add(1*time.Hour))
	pricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uusdc:usd", sdk.MustNewDecFromStr("0.000001"), suite.ctx.BlockTime().Add(1*time.Hour))

	pricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usdc")
	pricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	pricefeedKeeper.SetCurrentPrices(suite.ctx, "uusdc:usd")

	keeper := keeper.NewKeeper(appCodec, app.GetKey(types.StoreKey), app.GetKey(types.MemStoreKey), suite.app.GetSubspace(types.ModuleName), bankKeeper, pricefeedKeeper)

	params := types.DefaultParams()
	params.Pool.AcceptedAssets = []*types.Asset{
		{Denom: "uatom", TargetWeight: sdk.MustNewDecFromStr("0.5")},
		{Denom: "uusdc", TargetWeight: sdk.MustNewDecFromStr("0.5")},
	}
	params.PerpetualFutures.Markets = []*types.Market{
		{BaseDenom: "uatom", QuoteDenom: "uusdc"},
	}

	keeper.SetParams(suite.ctx, params)

	for _, asset := range params.Pool.AcceptedAssets {
		keeper.AddPoolAsset(suite.ctx, *asset)
	}
	suite.keeper = keeper
	suite.pricefeedKeeper = pricefeedKeeper
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
