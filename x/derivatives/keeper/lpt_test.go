package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/derivatives/types"
	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

// TODO: impl more various situations for the test cases
func (suite *KeeperTestSuite) TestInitialLiquidityProviderTokenSupply() {
	mockPrice := sdk.OneDec()
	mockAssetPrice := &pftypes.CurrentPrice{
		MarketId: "uatom:usd",
		Price:    mockPrice,
	}

	mockDepositingTokenAmount := sdk.OneDec()
	mockAssetMarketCap := mockPrice.Mul(mockDepositingTokenAmount)

	initialLPTSupply, fee, err := suite.app.DerivativesKeeper.InitialLiquidityProviderTokenSupply(suite.ctx, mockAssetPrice, mockAssetMarketCap, TestBaseTokenDenom)
	suite.Require().Equal(sdk.NewInt(2), initialLPTSupply.Amount)
	suite.Require().Equal(fee, sdk.NewCoin(types.LiquidityProviderTokenDenom, sdk.ZeroInt()))
	suite.Require().Nil(err)
}

func (suite *KeeperTestSuite) TestDetermineMintingLPTokenAmount() {
	// when no liquidity provider token's available
	mintAmount, mintFee, err := suite.keeper.DetermineMintingLPTokenAmount(suite.ctx, sdk.NewInt64Coin("uatom", 10000))
	suite.Require().NoError(err)
	suite.Require().Equal(mintAmount.String(), "20000udlp")
	suite.Require().Equal(mintFee.String(), "0udlp")

	// set price for asset
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:uusdc", sdk.NewDec(13), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:uusdc")
	suite.Require().NoError(err)

	// add pool asset and balance
	suite.keeper.AddPoolAsset(suite.ctx, types.PoolParams_Asset{
		Denom:        "uatom",
		TargetWeight: sdk.OneDec(),
	})
	suite.keeper.SetAssetBalance(suite.ctx, sdk.NewInt64Coin("uatom", 1000000))

	// set lp token supply
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(types.LiquidityProviderTokenDenom, 1000000)})
	suite.Require().NoError(err)

	// when liquidity provider token's available
	mintAmount, mintFee, err = suite.keeper.DetermineMintingLPTokenAmount(suite.ctx, sdk.NewInt64Coin("uatom", 10000))
	suite.Require().NoError(err)
	suite.Require().Equal(mintAmount.String(), "10000udlp")
	suite.Require().Equal(mintFee.String(), "10udlp")
}

func (suite *KeeperTestSuite) TestLPTokenSupplySnapshotGetSet() {
	supply := suite.keeper.GetLPTokenSupplySnapshot(suite.ctx, 1)
	suite.Require().Equal(supply, sdk.ZeroInt())
	suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 1, sdk.NewInt(1000000))
	supply = suite.keeper.GetLPTokenSupplySnapshot(suite.ctx, 1)
	suite.Require().Equal(supply, sdk.NewInt(1000000))
}

// TODO: write test for
// GetLPTokenSupply(ctx sdk.Context) sdk.Int {
// GetLPTokenPrice(ctx sdk.Context) sdk.Dec {
// GetLPTokenAmount(ctx sdk.Context, amount sdk.Coin) (sdk.Coin, sdk.Coin, error) {
// GetRedeemDenomAmount(ctx sdk.Context, lptAmount sdk.Int, redeemDenom string) (sdk.Coin, sdk.Coin, error) {
// DecreaseRedeemDenomAmount(ctx sdk.Context, amount sdk.Coin) error {
// MintLiquidityProviderToken(ctx sdk.Context, msg *types.MsgMintLiquidityProviderToken) error {
// BurnLiquidityProviderToken(ctx sdk.Context, msg *types.MsgBurnLiquidityProviderToken) error {
// BurnCoin(ctx sdk.Context, burner sdk.AccAddress, amount sdk.Coin) error {
