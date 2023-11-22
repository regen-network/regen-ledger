package cli_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/regen-network/regen-ledger/v5/app"
	"github.com/regen-network/regen-ledger/v5/app/client/cli"
)

func TestInitCmd(t *testing.T) {
	tmp := os.TempDir()
	nodeHome := filepath.Join(tmp, "test_init_cmd")

	// clean up previous test home directory
	err := os.RemoveAll(nodeHome)
	require.NoError(t, err)

	// create new test home directory
	err = os.Mkdir(nodeHome, 0755)
	require.NoError(t, err)

	rootCmd, _ := cli.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",
		"test",
		fmt.Sprintf("--%s=%s", flags.FlagHome, nodeHome),
	})

	err = cmd.Execute(rootCmd, app.EnvPrefix, nodeHome)
	require.NoError(t, err)
}
