package types

import (
	"bytes"
	"encoding/json"
)

type incentiveInputs struct {
	Version         string `json:"version"`
	IncentiveUnitId string `json:"incentive-unit-id"`
}

type XIncentiveInputs incentiveInputs

type XIncentiveInputsExceptions struct {
	XIncentiveInputs
	Other *string // other won't raise an error
}

func ParseMemo(memoContent []byte) (*incentiveInputs, error) {
	incentiveInputs := &incentiveInputs{}

	if err := incentiveInputs.UnmarshalJSON(memoContent); err != nil {
		// TODO: emit Event
		return nil, err
	}

	// make exception if unknown fields exists
	return incentiveInputs, nil
}

// UnmarshalJSON should error if there are fields unexpected.
func (memo *incentiveInputs) UnmarshalJSON(data []byte) error {
	var incentiveE XIncentiveInputsExceptions
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields() // Force

	if err := dec.Decode(&incentiveE); err != nil {
		return err
	}

	*memo = incentiveInputs(incentiveE.XIncentiveInputs)
	return nil
}
