package keeper_test

import (
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmtime "github.com/cometbft/cometbft/types/time"

	"github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/x/pricefeed/keeper"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	addrs  []string
	ctx    sdk.Context
}

func (suite *KeeperTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	_, addrs := app.GeneratePrivKeyAddressPairs(10)
	tApp.InitializeFromGenesisStates(
		NewPricefeedGenStateMulti(tApp),
	)
	var strAddrs []string
	for _, addr := range addrs {
		strAddrs = append(strAddrs, addr.String())
	}
	suite.keeper = tApp.GetPriceFeedKeeper()
	suite.ctx = ctx
	suite.addrs = strAddrs
}

func (suite *KeeperTestSuite) TestGetSetOracles() {
	params := suite.keeper.GetParams(suite.ctx)
	suite.Equal([]string{}, params.Markets[0].Oracles)

	params.Markets[0].Oracles = suite.addrs
	suite.NotPanics(func() { suite.keeper.SetParams(suite.ctx, params) })
	params = suite.keeper.GetParams(suite.ctx)
	suite.Equal(suite.addrs, params.Markets[0].Oracles)

	accAddr, err := sdk.AccAddressFromBech32(suite.addrs[0])
	suite.NoError(err)
	_, err = suite.keeper.GetOracle(suite.ctx, params.Markets[0].MarketId, accAddr)
	suite.NoError(err)
}

func (suite *KeeperTestSuite) TestGetAuthorizedAddresses() {
	_, oracles := app.GeneratePrivKeyAddressPairs(5)
	var strAddrs1 []string
	for _, addr := range oracles[:3] {
		strAddrs1 = append(strAddrs1, addr.String())
	}
	var strAddrs2 []string
	for _, addr := range oracles[2:] {
		strAddrs2 = append(strAddrs2, addr.String())
	}
	params := pricefeedtypes.Params{
		Markets: []pricefeedtypes.Market{
			{MarketId: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: strAddrs1, Active: true},
			{MarketId: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: strAddrs2, Active: true},
			{MarketId: "xrp:usd:30", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: nil, Active: true},
		},
	}
	suite.keeper.SetParams(suite.ctx, params)

	actualOracles, err := suite.keeper.GetAuthorizedAddresses(suite.ctx)
	suite.NoError(err)

	suite.Require().ElementsMatch(oracles, actualOracles)
}

// func TestKeeperTestSuite(t *testing.T) {
// 	suite.Run(t, new(KeeperTestSuite))
// }
