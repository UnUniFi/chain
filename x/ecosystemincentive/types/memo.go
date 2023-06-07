package types

import (
	"bytes"
	"encoding/json"
)

type XTxMemoData TxMemoData

type XTxMemoDataExceptions struct {
	XTxMemoData
	Other *string // other won't raise an error
}

func ParseMemo(memoContent []byte) (*TxMemoData, error) {
	txMemoData := &TxMemoData{}

	if err := txMemoData.UnmarshalJSON(memoContent); err != nil {
		return nil, err
	}

	// make exception if unknown fields exists
	return txMemoData, nil
}

// UnmarshalJSON should error if there are fields unexpected.
func (memo *TxMemoData) UnmarshalJSON(data []byte) error {
	var txMemoDataE XTxMemoDataExceptions
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields() // Force

	if err := dec.Decode(&txMemoDataE); err != nil {
		return err
	}

	*memo = TxMemoData(txMemoDataE.XTxMemoData)
	return nil
}

var AvailableVersions = []string{
	"v1",
}
