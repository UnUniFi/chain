package keeper_test

import (
	"fmt"
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

	initialLPTSupply, err := suite.app.DerivativesKeeper.InitialLiquidityProviderTokenSupply(suite.ctx, mockAssetPrice, mockAssetMarketCap, TestBaseTokenDenom)
	suite.Require().Equal(sdk.NewInt(2), initialLPTSupply.Amount)
	suite.Require().Nil(err)
}

func (suite *KeeperTestSuite) TestDetermineMintingLPTokenAmount() {
	// when no liquidity provider token's available
	mintAmount, err := suite.keeper.DetermineMintingLPTokenAmount(suite.ctx, sdk.NewInt64Coin("uatom", 10000))
	suite.Require().NoError(err)
	suite.Require().Equal(mintAmount.String(), "20000udlp")

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
	suite.keeper.SetAssetBalance(suite.ctx, sdk.NewInt64Coin("uatom", 1000000))

	// set lp token supply
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(types.LiquidityProviderTokenDenom, 1000000)})
	suite.Require().NoError(err)

	// when liquidity provider token's available
	mintAmount, err = suite.keeper.DetermineMintingLPTokenAmount(suite.ctx, sdk.NewInt64Coin("uatom", 10000))
	suite.Require().NoError(err)
	suite.Require().Equal(mintAmount.String(), "10000udlp")
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
	suite.keeper.SetAssetBalance(suite.ctx, sdk.NewInt64Coin("uatom", 1000000))

	// set lp token supply
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(types.LiquidityProviderTokenDenom, 1000000)})
	suite.Require().NoError(err)

	// check current height rate
	currLptPrice := suite.keeper.GetLPTokenPrice(suite.ctx)
	suite.Require().Equal(currLptPrice.String(), "0.000010000000000000")
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

func (suite *KeeperTestSuite) TestMintLiquidityProviderToken() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testCases := []struct {
		name        string
		sendCoin    sdk.Coin
		expectedErr bool
	}{
		{
			name:        "Success: mint lp token for the first time",
			sendCoin:    sdk.NewInt64Coin("uatom", 1000000),
			expectedErr: false,
		},
		{
			name:        "Success: mint lp token for the second time with the same denom as the first time",
			sendCoin:    sdk.NewInt64Coin("uatom", 1000000),
			expectedErr: false,
		},
		{
			name:        "Success: mint lp token with different denom",
			sendCoin:    sdk.NewInt64Coin("uusdc", 10000000),
			expectedErr: false,
		},
		{
			name:        "Success: mint lp token with different market value",
			sendCoin:    sdk.NewInt64Coin("uusdc", 1000000),
			expectedErr: false,
		},
		{
			name:        "Error: mint lp token with invalid denom",
			sendCoin:    sdk.NewInt64Coin("uinvalid", 10000000),
			expectedErr: true,
		},
	}

	derivativeFeeCollector := suite.app.AccountKeeper.GetModuleAddress(types.DerivativeFeeCollector)
	poolAddress := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)

	for i, tc := range testCases {
		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.sendCoin})
		_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, owner, sdk.Coins{tc.sendCoin})

		err := suite.keeper.MintLiquidityProviderToken(suite.ctx, &types.MsgDepositToPool{
			Sender: owner.Bytes(),
			Amount: tc.sendCoin,
		})

		if tc.expectedErr {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)

			if i == 0 {
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
				suite.Require().Equal("2000000udlp", balance.String())

				feeBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, derivativeFeeCollector)
				suite.Require().Equal(sdk.Coins{}, feeBalance)

				suite.CheckLptPrice(poolAddress)
			} else if i == 1 {
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
				suite.Require().Equal("3996000udlp", balance.String())

				feeBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, derivativeFeeCollector)
				suite.Require().Equal("2000uatom", feeBalance.String())

				suite.CheckLptPrice(poolAddress)
			} else if i == 2 {
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
				suite.Require().Equal("5994000udlp", balance.String())

				feeBalance := suite.app.BankKeeper.GetBalance(suite.ctx, derivativeFeeCollector, tc.sendCoin.Denom)
				suite.Require().Equal("10000uusdc", feeBalance.String())

				suite.CheckLptPrice(poolAddress)
			} else if i == 3 {
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
				suite.Require().Equal("6193800udlp", balance.String())

				feeBalance := suite.app.BankKeeper.GetBalance(suite.ctx, derivativeFeeCollector, tc.sendCoin.Denom)
				suite.Require().Equal("11000uusdc", feeBalance.String())

				suite.CheckLptPrice(poolAddress)
			}
		}
	}
}

func (suite *KeeperTestSuite) CheckLptPrice(poolAddress sdk.AccAddress) {
	poolBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, poolAddress)
	poolTotalMarketCap := sdk.NewDecFromInt(poolBalance.AmountOf("uatom").MulRaw(10).Add(poolBalance.AmountOf("uusdc"))).Quo(sdk.MustNewDecFromStr("1000000"))

	lptSupply := suite.keeper.GetLPTokenSupply(suite.ctx)
	expectingLPTPrice := poolTotalMarketCap.Quo(sdk.NewDecFromInt(lptSupply))
	dlpPrice := suite.keeper.GetLPTokenPrice(suite.ctx)
	fmt.Println(dlpPrice)
	suite.Require().Equal(expectingLPTPrice, dlpPrice)
}
