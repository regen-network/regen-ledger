package testsuite

import (
	"os"
	"testing"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/regen-network/regen-ledger/v5/app"
)

func TestSimAppExportAndBlockedAddrs(t *testing.T) {
	encCfg := app.MakeEncodingConfig()
	db := dbm.NewMemDB()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	regenApp := NewAppWithCustomOptions(t, false, SetupOptions{
		Logger:             logger,
		DB:                 db,
		InvCheckPeriod:     0,
		EncConfig:          encCfg,
		HomePath:           app.DefaultNodeHome,
		SkipUpgradeHeights: map[int64]bool{},
		AppOpts:            EmptyAppOptions{},
	})

	for acc := range app.GetMaccPerms() {
		require.Equal(t, true, regenApp.BankKeeper.BlockedAddr(regenApp.AccountKeeper.GetModuleAddress(acc)),
			"ensure that all module account addresses are properly blocked in bank keeper")
	}

	regenApp.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := app.NewRegenApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, EmptyAppOptions{})
	_, err := app2.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
