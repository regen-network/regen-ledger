package main

import (
	"os"

	cmd "github.com/regen-network/regen-ledger/v4/app/regen/cmd"
)

// In main we call the rootCmd
func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
