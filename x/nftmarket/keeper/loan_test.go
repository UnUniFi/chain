package keeper_test

import (
	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestDebtBasics() {
	debts := []types.Loan{
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "1",
			},
			Loan: sdk.NewInt64Coin("uguu", 1000000),
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "2",
			},
			Loan: sdk.NewInt64Coin("uguu", 1000000),
		},
	}

	for _, debt := range debts {
		suite.app.NftmarketKeeper.SetDebt(suite.ctx, debt)
	}

	for _, debt := range debts {
		loan := suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, debt.NftId.IdBytes())
		suite.Require().Equal(loan, debt)
	}

	// check all debts
	allDebts := suite.app.NftmarketKeeper.GetAllDebts(suite.ctx)
	suite.Require().Len(allDebts, len(debts))

	// delete all the debts
	for _, debt := range debts {
		suite.app.NftmarketKeeper.DeleteDebt(suite.ctx, debt.NftId.IdBytes())
	}

	// check all debts
	allDebts = suite.app.NftmarketKeeper.GetAllDebts(suite.ctx)
	suite.Require().Len(allDebts, 0)
}

func (suite *KeeperTestSuite) TestIncreaseDecreaseDebt() {
	nftIdentifier := types.NftIdentifier{
		ClassId: "1",
		NftId:   "1",
	}

	loan := suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.Coin{})

	suite.app.NftmarketKeeper.IncreaseDebt(suite.ctx, nftIdentifier, sdk.NewInt64Coin("uguu", 1000000))
	loan = suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.NewInt64Coin("uguu", 1000000))

	suite.app.NftmarketKeeper.DecreaseDebt(suite.ctx, nftIdentifier, sdk.NewInt64Coin("uguu", 500000))
	loan = suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.NewInt64Coin("uguu", 500000))

	suite.app.NftmarketKeeper.DecreaseDebt(suite.ctx, nftIdentifier, sdk.NewInt64Coin("uguu", 500000))
	loan = suite.app.NftmarketKeeper.GetDebtByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Equal(loan.Loan, sdk.NewInt64Coin("uguu", 0))
}

// TODO: add test for Borrow
// TODO: add test for Repay
// TODO: add test for Liquidate
