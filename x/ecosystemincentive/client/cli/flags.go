package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	// Will be parsed to string
	FlagRegisterFile = "register-file"

	// Names of fields in incentive-unit json file
	IncentiveUnitId           = "incentive_unit_id"
	IncentiveUnitSubjectAddrs = "subject_addrs"
	IncentiveUnitWeights      = "weights"
)

type registerInputs struct {
	IncentiveUnitId string   `json:"incentive_unit_id"`
	SubjectAddrs    []string `json:"subject_addrs"`
	Weights         []string `json:"weights"`
}

func FlagSetRegister() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagRegisterFile, "", "Register json file path (if this path is given, other register flags should not be used)")
	return fs
}
