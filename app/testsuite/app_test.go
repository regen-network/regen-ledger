package testsuite

import (
	"fmt"
	"os"
	"testing"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v7/testing/simapp"
	regenapp "github.com/regen-network/regen-ledger/v6/app"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/libs/log"
)

func TestSimAppExportAndBlockedAddrs(t *testing.T) {
	db := dbm.NewMemDB()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	app := NewAppWithCustomOptions(t, false, SetupOptions{
		Logger:  logger,
		DB:      db,
		AppOpts: simtestutil.NewAppOptionsWithFlagHome(t.TempDir()),
	})

	// BlockedAddresses returns a map of addresses in app v1 and a map of modules name in app v2.
	for acc := range simapp.BlockedAddresses() {
		var addr sdk.AccAddress
		if modAddr, err := sdk.AccAddressFromBech32(acc); err == nil {
			addr = modAddr
		} else {
			addr = app.AccountKeeper.GetModuleAddress(acc)
		}

		require.True(
			t,
			app.BankKeeper.BlockedAddr(addr),
			fmt.Sprintf("ensure that blocked addresses are properly set in bank keeper: %s should be blocked", acc),
		)
	}

	app.Commit()

	logger2 := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	// Making a new app object with the db, so that initchain hasn't been called
	app2 := regenapp.NewRegenApp(logger2, db, nil, true, 0, EmptyAppOptions{}, emptyWasmOption)
	_, err := app2.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
