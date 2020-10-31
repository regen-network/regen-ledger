package main

import (
	"os"

	cmd "github.com/regen-network/regen-ledger/app/cmd/regen"
)

// In main we call the rootCmd
func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
