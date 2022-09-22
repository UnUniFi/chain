package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	// Will be parsed to string
	FlagRegisterFile = "register-file"

	// Names of fields in incentive-unit json file
	IncentiveUnitId           = "incentive-id"
	IncentiveUnitSubjectAddrs = "subject-addrs"
	IncentiveUnitWeights      = "weights"
)

type registerInputs struct {
	IncentiveId  string   `json:"incentive-id"`
	SubjectAddrs []string `json:"subject-addrs"`
	Weights      []string `json:"weights"`
}

func FlagSetRegister() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagRegisterFile, "", "Register json file path (if this path is given, other register flags should not be used)")
	return fs
}
