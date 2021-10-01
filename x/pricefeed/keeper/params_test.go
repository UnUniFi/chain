package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/lcnem/jpyx/app"
	"github.com/lcnem/jpyx/types"
	"github.com/lcnem/jpyx/x/pricefeed/keeper"
	pricefeedtypes "github.com/lcnem/jpyx/x/pricefeed/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	addrs  []types.StringAccAddress
	ctx    sdk.Context
}

func (suite *KeeperTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	_, addrs := app.GeneratePrivKeyAddressPairs(10)
	tApp.InitializeFromGenesisStates(
		NewPricefeedGenStateMulti(tApp),
	)
	suite.keeper = tApp.GetPriceFeedKeeper()
	suite.ctx = ctx
	suite.addrs = types.StringAccAddresses(addrs)
}

func (suite *KeeperTestSuite) TestGetSetOracles() {
	params := suite.keeper.GetParams(suite.ctx)
	suite.Equal([]types.StringAccAddress(nil), params.Markets[0].Oracles)

	params.Markets[0].Oracles = suite.addrs
	suite.NotPanics(func() { suite.keeper.SetParams(suite.ctx, params) })
	params = suite.keeper.GetParams(suite.ctx)
	suite.Equal(suite.addrs, params.Markets[0].Oracles)

	addr, err := suite.keeper.GetOracle(suite.ctx, params.Markets[0].MarketId, suite.addrs[0].AccAddress())
	suite.NoError(err)
	suite.Equal(suite.addrs[0].AccAddress(), addr)
}

func (suite *KeeperTestSuite) TestGetAuthorizedAddresses() {
	_, oracles := app.GeneratePrivKeyAddressPairs(5)
	params := pricefeedtypes.Params{
		Markets: []pricefeedtypes.Market{
			{MarketId: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: types.StringAccAddresses(oracles[:3]), Active: true},
			{MarketId: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: types.StringAccAddresses(oracles[2:]), Active: true},
			{MarketId: "xrp:usd:30", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: nil, Active: true},
		},
	}
	suite.keeper.SetParams(suite.ctx, params)

	actualOracles := suite.keeper.GetAuthorizedAddresses(suite.ctx)

	suite.Require().ElementsMatch(oracles, actualOracles)
}
func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
