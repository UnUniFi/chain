package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ecosystemincentivekeeper "github.com/UnUniFi/chain/x/ecosystemincentive/keeper"
)

type FrontendIncentiveDecorator struct {
	keeper ecosystemincentivekeeper.Keeper
}

func NewFrontendIncentiveDecorator(keeper ecosystemincentivekeeper.Keeper) FrontendIncentiveDecorator {
	return FrontendIncentiveDecorator{
		keeper: keeper,
	}
}

func (decorator FrontendIncentiveDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	memoTx, ok := tx.(sdk.TxWithMemo)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}
	msgs := memoTx.GetMsgs()
	memo := memoTx.GetMemo()

	// TODO
	println("memo: ", memo)
	println("msgs: ", msgs)

	return next(ctx, tx, simulate)
}
