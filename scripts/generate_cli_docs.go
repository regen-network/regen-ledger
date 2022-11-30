package main

import (
	"log"

	"github.com/spf13/cobra/doc"

	"github.com/regen-network/regen-ledger/v5/app/client/cli"
)

// generate documentation for all regen app commands
func main() {
	rootCmd, _ := cli.NewRootCmd()
	err := doc.GenMarkdownTree(rootCmd, "commands")
	if err != nil {
		log.Fatal(err)
	}
}
