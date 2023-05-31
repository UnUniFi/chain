package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func GetMemo(txBytes []byte, txCfg client.TxConfig) string {
	/// NOTE: this way requires txConfig by importing it into keeper struct
	txData, err := txCfg.TxDecoder()(txBytes)
	if err != nil {
		fmt.Printf("err: %v\n", err)

		txData, err = txCfg.TxJSONDecoder()(txBytes)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	txBldr, err := txCfg.WrapTxBuilder(txData)
	if err != nil {
		return ""
	}
	memo := txBldr.GetTx().GetMemo()

	return memo
}
