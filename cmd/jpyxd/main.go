package main

import (
	"os"

	"github.com/UnUniFi/chain/app"
	"github.com/UnUniFi/chain/cmd/jpyxd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
