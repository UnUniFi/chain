package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func (k Keeper) GetMemo(ctx sdk.Context) (string, error) {
	txBytes := ctx.TxBytes()

	/// NOTE: this way requires txConfig by importing it into keeper struct
	txData, err := k.txCfg.TxDecoder()(txBytes)
	if err != nil {
		fmt.Printf("err: %v\n", err)

		txData, err = k.txCfg.TxJSONDecoder()(txBytes)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	txBldr, err := k.txCfg.WrapTxBuilder(txData)
	if err != nil {
		return "", err
	}
	memo := txBldr.GetTx().GetMemo()

	return memo, nil
}
