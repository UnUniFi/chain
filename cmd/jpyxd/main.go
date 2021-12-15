package main

import (
	"os"

	"github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/cmd/ununifid/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
