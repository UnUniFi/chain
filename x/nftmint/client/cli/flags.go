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
	FsCreateClass.String(FlagSymbol, "", "Class Symbol")
	FsCreateClass.String(FlagDescription, "", "Description for denom")
	FsCreateClass.String(FlagClassUri, "", "Content uri for class")
}
