package main

import (
	"os"

	"github.com/cheqd/cosmos-sdk/server"
	svrcmd "github.com/cheqd/cosmos-sdk/server/cmd"

	app "github.com/cosmos/gaia/v7/app"
	"github.com/cosmos/gaia/v7/cmd/gaiad/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
