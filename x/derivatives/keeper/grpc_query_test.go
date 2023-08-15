package keeper_test

import (
	"testing"
	"time"

	testkeeper "github.com/UnUniFi/chain/testutil/keeper"
	"github.com/UnUniFi/chain/x/derivatives/types"
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/require"
)

// TODO: add test for followings after full implementation
// LiquidityProviderTokenRealAPY
// LiquidityProviderTokenNominalAPY
// PerpetualFutures
// PerpetualFuturesMarket
// PerpetualOptions
// Pool

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.DerivativesKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

func (suite *KeeperTestSuite) TestQueryLiquidityProviderTokenRealAPY() {
	// setup snapshot for block height 1
	suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 1, sdk.NewInt(1000000))
	suite.keeper.SetPoolMarketCapSnapshot(suite.ctx, 1, types.PoolMarketCap{
		QuoteTicker: "uatom",
		Total:       sdk.NewDec(100000000),
		AssetInfo:   []types.PoolMarketCap_AssetInfo{},
	})

	suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 1, sdk.NewInt(1000000))
	// setup snapshot for block height 11
	suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 11, sdk.NewInt(2000000))
	suite.keeper.SetPoolMarketCapSnapshot(suite.ctx, 11, types.PoolMarketCap{
		QuoteTicker: "uatom",
		Total:       sdk.NewDec(300000000),
		AssetInfo:   []types.PoolMarketCap_AssetInfo{},
	})
	now := time.Now()
	future := time.Now().Add(time.Second * 11 * 100)
	suite.keeper.SaveBlockTimestamp(suite.ctx, 1, now)
	suite.keeper.SaveBlockTimestamp(suite.ctx, 11, future)
	req := &types.QueryLiquidityProviderTokenRealAPYRequest{
		BeforeHeight: 1,
		AfterHeight:  11,
	}

	suite.ctx = suite.ctx.WithBlockHeight(20)

	actual, err := suite.keeper.LiquidityProviderTokenRealAPY(suite.ctx, req)
	suite.Require().NoError(err)
	//rate: 0.5  duration: 100seconds
	apy := sdk.MustNewDecFromStr("14334.545454545454545455")
	expected := &types.QueryLiquidityProviderTokenRealAPYResponse{
		Apy: &apy,
	}
	suite.Require().Equal(
		expected,
		actual,
	)
}

func (suite *KeeperTestSuite) TestQueryLiquidityProviderTokenNominalAPY() {
	// setup snapshot for block height 1
	suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 1, sdk.NewInt(1000000))
	suite.keeper.SetPoolMarketCapSnapshot(suite.ctx, 1, types.PoolMarketCap{
		QuoteTicker: "uatom",
		Total:       sdk.NewDec(100000000),
		AssetInfo:   []types.PoolMarketCap_AssetInfo{},
	})

	suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 1, sdk.NewInt(1000000))
	// setup snapshot for block height 11
	suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 11, sdk.NewInt(2000000))
	suite.keeper.SetPoolMarketCapSnapshot(suite.ctx, 11, types.PoolMarketCap{
		QuoteTicker: "uatom",
		Total:       sdk.NewDec(300000000),
		AssetInfo:   []types.PoolMarketCap_AssetInfo{},
	})
	now := time.Now()
	future := time.Now().Add(time.Second * 11 * 100)
	suite.keeper.SaveBlockTimestamp(suite.ctx, 1, now)
	suite.keeper.SaveBlockTimestamp(suite.ctx, 11, future)

	suite.ctx = suite.ctx.WithBlockHeight(20)

	req := &types.QueryLiquidityProviderTokenNominalAPYRequest{
		BeforeHeight: 1,
		AfterHeight:  11,
	}
	actual, err := suite.keeper.LiquidityProviderTokenNominalAPY(suite.ctx, req)
	suite.Require().NoError(err)
	apy := sdk.MustNewDecFromStr("14334.545454545454545455")
	expected := &types.QueryLiquidityProviderTokenNominalAPYResponse{
		Apy: &apy,
	}
	suite.Require().Equal(
		expected,
		actual,
	)
}

func (suite *KeeperTestSuite) TestQueryPerpetualFutures() {
	// suite.SetParams()
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	coins := sdk.Coins{sdk.NewCoin("uatom", sdk.NewInt(5000000)), sdk.NewCoin("uusdc", sdk.NewInt(50000000))}
	_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, owner, coins)
	positions := []struct {
		positionId           string
		margin               sdk.Coin
		instance             types.PerpetualFuturesPositionInstance
		availableAssetInPool sdk.Coin
	}{
		{
			positionId: "0",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(2000000)),
		},
		{
			positionId: "1",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
		},
	}
	for _, testPosition := range positions {
		err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{testPosition.availableAssetInPool})
		suite.Require().NoError(err)

		position, err := suite.keeper.OpenPerpetualFuturesPosition(suite.ctx, testPosition.positionId, owner.String(), testPosition.margin, market, testPosition.instance)
		suite.Require().NoError(err)
		suite.Require().NotNil(position)

		_ = suite.keeper.SetPosition(suite.ctx, *position)
		_ = suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{testPosition.margin})
	}

	req := &types.QueryPerpetualFuturesRequest{}

	actual, err := suite.keeper.PerpetualFutures(suite.ctx, req)
	suite.Require().NoError(err)
	expected := &types.QueryPerpetualFuturesResponse{
		MetricsQuoteTicker: "usd",
		LongPositions:      sdk.NewDec(20000000), //position_size * 1000000 * price_per_quote
		ShortPositions:     sdk.NewDec(10000000), //position_size*1000000 * price_per_quote
	}
	suite.Require().Equal(
		expected,
		actual,
	)
}

func (suite *KeeperTestSuite) TestQueryPerpetualFuturesMarket() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	coins := sdk.Coins{sdk.NewCoin("uatom", sdk.NewInt(5000000)), sdk.NewCoin("uusdc", sdk.NewInt(50000000))}
	_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, owner, coins)
	positions := []struct {
		positionId           string
		margin               sdk.Coin
		instance             types.PerpetualFuturesPositionInstance
		availableAssetInPool sdk.Coin
	}{
		{
			positionId: "0",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(2000000)),
		},
		{
			positionId: "1",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
		},
	}
	for _, testPosition := range positions {
		err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{testPosition.availableAssetInPool})
		suite.Require().NoError(err)

		position, err := suite.keeper.OpenPerpetualFuturesPosition(suite.ctx, testPosition.positionId, owner.String(), testPosition.margin, market, testPosition.instance)
		suite.Require().NoError(err)
		suite.Require().NotNil(position)

		_ = suite.keeper.SetPosition(suite.ctx, *position)
		_ = suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{testPosition.margin})
	}
	req := &types.QueryPerpetualFuturesMarketRequest{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	actual, err := suite.keeper.PerpetualFuturesMarket(suite.ctx, req)
	//error
	suite.Require().NoError(err)
	price := sdk.NewDec(0)
	longPositions := sdk.NewDec(20000000)
	shortPositions := sdk.NewDec(10000000)
	expected := &types.QueryPerpetualFuturesMarketResponse{
		Price:              &price,
		MetricsQuoteTicker: "usd",
		LongPositions:      &longPositions,
		ShortPositions:     &shortPositions,
	}
	suite.Require().Equal(
		expected,
		actual,
	)
}

//	func (suite *KeeperTestSuite) TestPerpetualOptions() {
//		req := &types.QueryPerpetualOptionsRequest{}
//		actual, err := suite.keeper.PerpetualOptions(suite.ctx, req)
//		suite.Require().NoError(err)
//		expected := &types.QueryPerpetualOptionsResponse{}
//		suite.Require().Equal(
//			actual,
//			expected,
//		)
//	}
func (suite *KeeperTestSuite) TestPool() {
	req := &types.QueryPoolRequest{}
	actual, err := suite.keeper.Pool(suite.ctx, req)
	suite.Require().NoError(err)
	expected := &types.QueryPoolResponse{
		MetricsQuoteTicker: "usd",
		PoolMarketCap: &types.PoolMarketCap{
			QuoteTicker: "usd",
			Total:       sdk.NewDec(0),
			AssetInfo: []types.PoolMarketCap_AssetInfo{
				{
					Denom:  "uatom",
					Amount: sdk.NewInt(0),
					Price:  sdk.MustNewDecFromStr("0.000010000000000000"),
				},
				{
					Denom:  "uusdc",
					Amount: sdk.NewInt(0),
					Price:  sdk.MustNewDecFromStr("0.000001000000000000"),
				},
			},
		},
	}
	suite.Require().Equal(
		actual,
		expected,
	)
}
