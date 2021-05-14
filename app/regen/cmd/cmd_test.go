package regen_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/regen-network/regen-ledger/app"
	"io/ioutil"
	"testing"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/stretchr/testify/require"

	cmd "github.com/regen-network/regen-ledger/app/regen/cmd"
)

func TestInitCmd(t *testing.T) {
	nodeHome, err := ioutil.TempDir(t.TempDir(), ".regen")
	require.NoError(t, err)

	rootCmd, _ := cmd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",          // Test the init cmd
		"regenapp-test", // Moniker
		fmt.Sprintf("--%s=%s", flags.FlagHome, nodeHome), // Set home flag
	})

	err = svrcmd.Execute(rootCmd, app.DefaultNodeHome)
	require.NoError(t, err)
}
