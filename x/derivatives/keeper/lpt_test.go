package keeper_test

import (
	"fmt"
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

// TODO: impl more various situations for the test cases
func (suite *KeeperTestSuite) TestInitialLiquidityProviderTokenSupply() {
	mockAmount := sdk.NewInt64Coin("uatom", 1000000)
	// mockPrice := sdk.OneDec()
	// mockAssetPrice := &pftypes.CurrentPrice{
	// 	MarketId: "uatom:usd",
	// 	Price:    mockPrice,
	// }

	//mockDepositingTokenAmount := sdk.OneDec()
	//mockAssetMarketCap := mockPrice.Mul(mockDepositingTokenAmount)

	initialLPTSupply, err := suite.app.DerivativesKeeper.InitialLiquidityProviderTokenSupply(suite.ctx, mockAmount)
	suite.Require().Equal(mockAmount.Amount, initialLPTSupply.Amount)
	suite.Require().Nil(err)
}

func (suite *KeeperTestSuite) TestDetermineMintingLPTokenAmount() {
	// when no liquidity provider token's available
	mintAmount, err := suite.keeper.DetermineMintingLPTokenAmount(suite.ctx, sdk.NewInt64Coin("uatom", 10000))
	suite.Require().NoError(err)
	suite.Require().Equal(mintAmount.String(), "10000udlp")

	// set price for asset
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.MustNewDecFromStr("0.000013"), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []string{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)

	// set lp token supply
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(types.LiquidityProviderTokenDenom, 1000000)})
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin("uatom", 500000)})
	suite.Require().NoError(err)

	// when liquidity provider token's available
	mintAmount, err = suite.keeper.DetermineMintingLPTokenAmount(suite.ctx, sdk.NewInt64Coin("uatom", 15000))
	suite.Require().NoError(err)
	// Rate udlp:uatom = 2:1
	suite.Require().Equal(mintAmount.String(), "30000udlp")
}

func (suite *KeeperTestSuite) TestLPTokenSupplySnapshotGetSet() {
	supply := suite.keeper.GetLPTokenSupplySnapshot(suite.ctx, 1)
	suite.Require().Equal(supply, sdk.ZeroInt())
	err := suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 1, sdk.NewInt(1000000))
	suite.Require().NoError(err)
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

func (suite *KeeperTestSuite) TestGetLPTokenPrice() {
	// set price for asset
	_, err := suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.MustNewDecFromStr("0.00002"), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []string{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)

	// set lp token supply 1dlp=1atom
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(types.LiquidityProviderTokenDenom, 1000000)})
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin("uatom", 1000000)})
	suite.Require().NoError(err)

	// check current height rate
	currLptPrice := suite.keeper.GetLPTokenPrice(suite.ctx)
	// 1udlp = 1uatom = 0.00002USD
	suite.Require().Equal(currLptPrice, sdk.MustNewDecFromStr("0.00002"))
}

func (suite *KeeperTestSuite) TestGetRedeemDenomAmount() {
	// get uninitialized redeem amount
	lptAmount := sdk.NewInt(1000000)
	redeemAmount, redeemFee, err := suite.keeper.GetRedeemDenomAmount(suite.ctx, lptAmount, "uatom")
	suite.Require().Error(err)
	suite.Require().True(redeemAmount.IsNil())
	suite.Require().True(redeemFee.IsNil())

	// set price for asset
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.MustNewDecFromStr("0.000013"), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []string{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)

	// set lp token supply
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(types.LiquidityProviderTokenDenom, 1000000)})
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin("uatom", 600000)})
	suite.Require().NoError(err)

	// get initialized redeem amount
	redeemAmount, redeemFee, err = suite.keeper.GetRedeemDenomAmount(suite.ctx, lptAmount, "uatom")
	suite.Require().NoError(err)
	suite.Require().Equal(redeemAmount.String(), "599400uatom")
	suite.Require().Equal(redeemFee.String(), "600uatom")
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
			Sender: owner.String(),
			Amount: tc.sendCoin,
		})

		if tc.expectedErr {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)

			if i == 0 {
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
				// Initial supply
				// 1000000 - 0fee= 1000000
				suite.Require().Equal("1000000udlp", balance.String())

				suite.CheckLptPrice(poolAddress)
			} else if i == 1 {
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
				// 1000000 - 2000fee = 998000
				suite.Require().Equal("1998000udlp", balance.String())

				feeBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, derivativeFeeCollector)
				suite.Require().Equal("2000uatom", feeBalance.String())

				suite.CheckLptPrice(poolAddress)
			} else if i == 2 {
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
				// usdc:lpt=10:2
				// (10000000 - 10000) / 10  = 999000
				suite.Require().Equal("2997000udlp", balance.String())

				feeBalance := suite.app.BankKeeper.GetBalance(suite.ctx, derivativeFeeCollector, tc.sendCoin.Denom)
				suite.Require().Equal("10000uusdc", feeBalance.String())

				suite.CheckLptPrice(poolAddress)
			} else if i == 3 {
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "udlp")
				// usdc:lpt=10:2
				// 1000000 - 1000 / 10 = 99900
				suite.Require().Equal("3096900udlp", balance.String())

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
