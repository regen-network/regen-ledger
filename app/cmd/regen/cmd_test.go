package regen_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/regen-network/regen-ledger/app/cmd/regen"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := regen.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",          // Test the init cmd
		"regenapp-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
	})

	err := regen.Execute(rootCmd)
	require.NoError(t, err)
}
