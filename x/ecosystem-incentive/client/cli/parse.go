package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type XRegisterInputs registerInputs

type XRegisterInputsExceptions struct {
	XRegisterInputs
	Other *string // Other won't raise an error
}

// UnmarshalJSON should error if there are fields unexpected.
func (release *registerInputs) UnmarshalJSON(data []byte) error {
	var registerE XRegisterInputsExceptions
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields() // Force

	if err := dec.Decode(&registerE); err != nil {
		return err
	}

	*release = registerInputs(registerE.XRegisterInputs)
	return nil
}

func parseRegisterFlags(fs *pflag.FlagSet) (*registerInputs, error) {
	registerInputs := &registerInputs{}
	incentiveUnitFile, _ := fs.GetString(FlagRegisterFile)

	if incentiveUnitFile == "" {
		return nil, fmt.Errorf("must pass in a pool json using the --%s flag", FlagRegisterFile)
	}

	contents, err := os.ReadFile(incentiveUnitFile)
	if err != nil {
		return nil, err
	}

	// make exception if unknown fields exists
	if err := registerInputs.UnmarshalJSON(contents); err != nil {
		return nil, err
	}

	return registerInputs, nil
}
