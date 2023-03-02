package simulation

import (
	"math/rand"

	"github.com/UnUniFi/chain/x/yield-aggregator/keeper"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDepositToVault(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDepositToVault{
			Sender: simAccount.Address.String(),
		}

		// TODO: Handling the DepositToVault simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "DepositToVault simulation not implemented"), nil, nil
	}
}
