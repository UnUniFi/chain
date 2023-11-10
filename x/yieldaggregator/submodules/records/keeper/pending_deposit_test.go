package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
)

func (suite *KeeperTestSuite) TestVaultPendingDepositStore() {
	deposits := []types.PendingDeposit{
		{
			VaultId: 1,
			Amount:  sdk.OneInt(),
		},
		{
			VaultId: 2,
			Amount:  sdk.NewInt(2),
		},
	}

	for _, deposit := range deposits {
		suite.app.RecordsKeeper.SetVaultPendingDeposit(suite.ctx, deposit.VaultId, deposit.Amount)
	}

	for _, deposit := range deposits {
		r := suite.app.RecordsKeeper.GetVaultPendingDeposit(suite.ctx, deposit.VaultId)
		suite.Require().Equal(r, deposit.Amount)
	}

	storedInfos := suite.app.RecordsKeeper.GetAllVaultPendingDeposits(suite.ctx)
	suite.Require().Len(storedInfos, 2)

	suite.app.RecordsKeeper.IncreaseVaultPendingDeposit(suite.ctx, 1, sdk.NewInt(100))
	r := suite.app.RecordsKeeper.GetVaultPendingDeposit(suite.ctx, 1)
	suite.Require().Equal(r, sdk.NewInt(101))

	suite.app.RecordsKeeper.DecreaseVaultPendingDeposit(suite.ctx, 1, sdk.NewInt(50))
	r = suite.app.RecordsKeeper.GetVaultPendingDeposit(suite.ctx, 1)
	suite.Require().Equal(r, sdk.NewInt(51))
}
