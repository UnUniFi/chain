package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

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

func (suite *KeeperTestSuite) TestGetLPTokenSupply() {
	// get initial supply value
	supply := suite.keeper.GetLPTokenSupply(suite.ctx)
	suite.Require().Equal(supply, sdk.ZeroInt())

	// add lp token supply
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(types.LiquidityProviderTokenDenom, 1000000)})
	suite.Require().NoError(err)

	// get after mint
	supply = suite.keeper.GetLPTokenSupply(suite.ctx)
	suite.Require().Equal(supply, sdk.NewInt(1000000))
}

// FIXME: fix test case
func (suite *KeeperTestSuite) TestGetLPTokenPrice() {
	// set price for asset
	_, err := suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:uusdc", sdk.NewDec(13), suite.ctx.BlockTime().Add(time.Hour*3))
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

	// check current height rate
	currLptPrice := suite.keeper.GetLPTokenPrice(suite.ctx)
	suite.Require().Equal(currLptPrice.String(), "0.000015280000000000")
}

func (suite *KeeperTestSuite) TestGetRedeemDenomAmount() {
	// get uninitialized redeem amount
	lptAmount := sdk.NewInt(1000000)
	redeemAmount, redeemFee, err := suite.keeper.GetRedeemDenomAmount(suite.ctx, lptAmount, "uatom")
	suite.Require().Error(err)
	suite.Require().True(redeemAmount.IsNil())
	suite.Require().True(redeemFee.IsNil())

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

	// get initialized redeem amount
	redeemAmount, redeemFee, err = suite.keeper.GetRedeemDenomAmount(suite.ctx, lptAmount, "uatom")
	suite.Require().NoError(err)
	suite.Require().Equal(redeemAmount.String(), "999000uatom")
	suite.Require().Equal(redeemFee.String(), "1000uatom")
}

func (suite *KeeperTestSuite) TestDecreaseRedeemDenomAmount() {
	// try operation on uninitialized environment
	err := suite.keeper.DecreaseRedeemDenomAmount(suite.ctx, sdk.NewInt64Coin("uatom", 10000))
	suite.Require().Error(err)

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

	// try after initialization
	err = suite.keeper.DecreaseRedeemDenomAmount(suite.ctx, sdk.NewInt64Coin("uatom", 10000))
	suite.Require().NoError(err)
	assetBalance := suite.keeper.GetAssetBalance(suite.ctx, "uatom")
	suite.Require().Equal(assetBalance.String(), "990000uatom")
}

func (suite *KeeperTestSuite) TestBurnCoin() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	err := suite.keeper.BurnCoin(suite.ctx, owner, sdk.NewInt64Coin("uatom", 10000))
	suite.Require().Error(err)

	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin("uatom", 1000000)})
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, owner, sdk.Coins{sdk.NewInt64Coin("uatom", 1000000)})
	suite.Require().NoError(err)

	err = suite.keeper.BurnCoin(suite.ctx, owner, sdk.NewInt64Coin("uatom", 10000))
	suite.Require().NoError(err)

	balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "uatom")
	suite.Require().Equal(balance.String(), "990000uatom")
}

func (suite *KeeperTestSuite) TestMintBurnLiquidityProviderToken() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin("uatom", 1000000)})
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, owner, sdk.Coins{sdk.NewInt64Coin("uatom", 1000000)})
	suite.Require().NoError(err)

	// when no liquidity provider token's available
	err = suite.keeper.MintLiquidityProviderToken(suite.ctx, &types.MsgMintLiquidityProviderToken{
		Sender: owner.Bytes(),
		Amount: sdk.NewInt64Coin("uatom", 10000),
	})
	suite.Require().NoError(err)

	balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
	suite.Require().Equal(balance.String(), "20000udlp")

	// mint more lp tokens
	err = suite.keeper.MintLiquidityProviderToken(suite.ctx, &types.MsgMintLiquidityProviderToken{
		Sender: owner.Bytes(),
		Amount: sdk.NewInt64Coin("uatom", 10000),
	})
	suite.Require().NoError(err)

	derivativeFeeCollector := suite.app.AccountKeeper.GetModuleAddress(types.DerivativeFeeCollector)
	feeBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, derivativeFeeCollector)
	suite.Require().Equal(feeBalance.String(), "40udlp")

	balance = suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
	suite.Require().Equal(balance.String(), "39960udlp")

	err = suite.keeper.BurnLiquidityProviderToken(suite.ctx, &types.MsgBurnLiquidityProviderToken{
		Sender:      owner.Bytes(),
		Amount:      sdk.NewInt(20000),
		RedeemDenom: "uatom",
	})
	suite.Require().NoError(err)
	balance = suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
	suite.Require().Equal(balance.String(), "19960udlp")

	balance = suite.app.BankKeeper.GetBalance(suite.ctx, owner, "uatom")
	suite.Require().Equal(balance.String(), "989990uatom")

	feeBalance = suite.app.BankKeeper.GetAllBalances(suite.ctx, derivativeFeeCollector)
	suite.Require().Equal(feeBalance.String(), "10uatom,40udlp")
}
