package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagName        = "name"
	FlagSymbol      = "symbol"
	FlagDescription = "description"
	FlagUri         = "uri"
	FlagUriHash     = "uri-hash"
)

var (
	FsCreateClass = flag.NewFlagSet("", flag.ContinueOnError)
	FsUpdateClass = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreateClass.String(FlagName, "", "Class name")
	FsCreateClass.String(FlagSymbol, "", "Class symbol")
	FsCreateClass.String(FlagDescription, "", "Description for denom")
	FsCreateClass.String(FlagUri, "", "Content URI for class")
	FsCreateClass.String(FlagUriHash, "", "Hash of content URI for class")

	FsUpdateClass.String(FlagName, "", "Class name")
	FsUpdateClass.String(FlagSymbol, "", "Class symbol")
	FsUpdateClass.String(FlagDescription, "", "Description for denom")
	FsUpdateClass.String(FlagUri, "", "Content URI for class")
	FsUpdateClass.String(FlagUriHash, "", "Hash of content URI for class")
}
