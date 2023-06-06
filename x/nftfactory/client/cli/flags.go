package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagSymbol      = "symbol"
	FlagDescription = "description"
	FlagClassUri    = "class-uri"
)

var (
	FsCreateClass = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreateClass.String(FlagSymbol, "", "Class symbol")
	FsCreateClass.String(FlagDescription, "", "Description for denom")
	FsCreateClass.String(FlagClassUri, "", "Content URI for class")
}
