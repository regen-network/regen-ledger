package testsuite

import (
	"fmt"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v7/testing/simapp"
	regenapp "github.com/regen-network/regen-ledger/v5/app"
	"os"
	"testing"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/libs/log"
)

//func TestSimAppExportAndBlockedAddrs(t *testing.T) {
//	db := dbm.NewMemDB()
//	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
//	setupOptions := SetupOptions{
//		Logger:         logger,
//		DB:             db,
//		InvCheckPeriod: 0,
//		AppOpts:        EmptyAppOptions{},
//	}
//	regenApp := NewAppWithCustomOptions(t, false, setupOptions)
//
//	for acc := range app.GetMaccPerms() {
//		require.Equal(t, true, regenApp.BankKeeper.BlockedAddr(regenApp.AccountKeeper.GetModuleAddress(acc)),
//			"ensure that all module account addresses are properly blocked in bank keeper")
//	}
//
//	regenApp.Commit()
//
//	// Making a new app object with the db, so that initchain hasn't been called
//	app2 := app.NewRegenApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, setupOptions.InvCheckPeriod, EmptyAppOptions{})
//	_, err := app2.ExportAppStateAndValidators(false, []string{}, []string{})
//	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
//}
//

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
	app2 := regenapp.NewRegenApp(logger2, db, nil, true, 0, EmptyAppOptions{})
	_, err := app2.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
