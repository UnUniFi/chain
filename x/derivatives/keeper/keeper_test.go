package keeper_test

import (
	"testing"
	"time"

	"github.com/CosmWasm/wasmd/x/wasm"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"
	pricefeedkeeper "github.com/UnUniFi/chain/x/pricefeed/keeper"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/baseapp"

	simapp "github.com/UnUniFi/chain/app"
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
	app := simapp.Setup(suite.T(), ([]wasm.Option{})...)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.DerivativesKeeper)
	queryClient := types.NewQueryClient(queryHelper)
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.app = app
	suite.queryClient = queryClient

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

	app.BankKeeper.SetDenomMetaData(suite.ctx, metadataAtom)
	app.BankKeeper.SetDenomMetaData(suite.ctx, metadataUsdc)

	pfParams := pricefeedtypes.Params{
		Markets: []pricefeedtypes.Market{
			{MarketId: "uusdc:usd", BaseAsset: TestQuoteTokenDenom, QuoteAsset: TestQuoteTokenDenom, Oracles: []string{}, Active: true},
			{MarketId: "uatom:usd", BaseAsset: TestBaseTokenDenom, QuoteAsset: TestQuoteTokenDenom, Oracles: []string{}, Active: true},
		},
	}
	app.PricefeedKeeper.SetParams(suite.ctx, pfParams)

	_, _ = app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.MustNewDecFromStr("0.00001"), suite.ctx.BlockTime().Add(1*time.Hour))
	_, _ = app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uusdc:usd", sdk.MustNewDecFromStr("0.000001"), suite.ctx.BlockTime().Add(1*time.Hour))

	_ = app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	_ = app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uusdc:usd")

	params := types.DefaultParams()
	params.PoolParams.AcceptedAssetsConf = []types.PoolAssetConf{
		{Denom: "uatom", TargetWeight: sdk.MustNewDecFromStr("0.5")},
		{Denom: "uusdc", TargetWeight: sdk.MustNewDecFromStr("0.5")},
	}
	params.PerpetualFutures.Markets = []*types.Market{
		{BaseDenom: TestBaseTokenDenom, QuoteDenom: TestQuoteTokenDenom},
	}

	app.DerivativesKeeper.SetParams(suite.ctx, params)

	suite.keeper = app.DerivativesKeeper
	suite.pricefeedKeeper = app.PricefeedKeeper

	nftfactoryParams := nftfactorytypes.DefaultParams()
	app.NftfactoryKeeper.SetParamSet(suite.ctx, nftfactoryParams)

	derivativesAddr := app.DerivativesKeeper.GetModuleAddress()
	_ = app.NftfactoryKeeper.CreateClass(suite.ctx, "derivatives/perpetual_futures/positions", &nftfactorytypes.MsgCreateClass{
		Sender:            derivativesAddr.String(),
		Name:              "derivatives/perpetual_futures/positions",
		BaseTokenUri:      "ipfs://testcid/",
		TokenSupplyCap:    100000,
		MintingPermission: 0,
		Symbol:            "",
		Description:       "",
		ClassUri:          "",
	})
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
