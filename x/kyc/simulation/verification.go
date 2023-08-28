package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/UnUniFi/chain/x/kyc/keeper"
	"github.com/UnUniFi/chain/x/kyc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateVerification(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateVerification{
			Sender:   simAccount.Address.String(),
			Customer: strconv.Itoa(i),
		}

		_, found := k.GetVerification(ctx, msg.Customer)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, "", "Verification already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         "",
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateVerification(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount      = simtypes.Account{}
			verification    = types.Verification{}
			msg             = &types.MsgUpdateVerification{}
			allVerification = k.GetAllVerification(ctx)
			found           = false
		)
		for _, obj := range allVerification {
			simAccount, found = FindAccount(accs, obj.Address)
			if found {
				verification = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, "", "verification creator not found"), nil, nil
		}
		msg.Sender = simAccount.Address.String()

		msg.Customer = verification.Address

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         "",
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteVerification(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount      = simtypes.Account{}
			verification    = types.Verification{}
			msg             = &types.MsgUpdateVerification{}
			allVerification = k.GetAllVerification(ctx)
			found           = false
		)
		for _, obj := range allVerification {
			simAccount, found = FindAccount(accs, obj.Address)
			if found {
				verification = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, "", "verification creator not found"), nil, nil
		}
		msg.Sender = simAccount.Address.String()

		msg.Customer = verification.Address

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         "",
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
