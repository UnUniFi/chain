package main

import (
	"os"

	"github.com/KimuraYu45z/test/app"
	"github.com/KimuraYu45z/test/cmd/testd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
