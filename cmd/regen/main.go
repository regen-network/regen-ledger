package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/regen-network/regen-ledger/v7/app"
	"github.com/regen-network/regen-ledger/v7/app/client/cli"
)

func main() {
	rootCmd := cli.NewRootCmd()
	if err := cmd.Execute(rootCmd, app.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
