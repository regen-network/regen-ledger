package main

import (
	"log"

	"github.com/spf13/cobra/doc"

	cmd "github.com/regen-network/regen-ledger/v3/app/regen/cmd"
)

// generate documentation for all regen app commands
func main() {
	rootCmd, _ := cmd.NewRootCmd()
	err := doc.GenMarkdownTree(rootCmd, "commands")
	if err != nil {
		log.Fatal(err)
	}
}
