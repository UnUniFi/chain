package types

import (
	"bytes"
	"encoding/json"
)

type XMemoInputs MemoInputs

type XMemoInputsExceptions struct {
	XMemoInputs
	Other *string // other won't raise an error
}

func ParseMemo(memoContent []byte) (*MemoInputs, error) {
	memoInputs := &MemoInputs{}

	if err := memoInputs.UnmarshalJSON(memoContent); err != nil {
		// TODO: emit Event
		return nil, err
	}

	// make exception if unknown fields exists
	return memoInputs, nil
}

// UnmarshalJSON should error if there are fields unexpected.
func (memo *MemoInputs) UnmarshalJSON(data []byte) error {
	var memoInputsE XMemoInputsExceptions
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields() // Force

	if err := dec.Decode(&memoInputsE); err != nil {
		return err
	}

	*memo = MemoInputs(memoInputsE.XMemoInputs)
	return nil
}

var AvailableVersions = []string{
	"v1",
}
