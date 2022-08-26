package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	"github.com/tendermint/tendermint/libs/log"
)

func TestSimAppExportAndBlockedAddrs(t *testing.T) {
	encCfg := MakeEncodingConfig()
	db := dbm.NewMemDB()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	app := NewAppWithCustomOptions(t, false, SetupOptions{
		Logger:             logger,
		DB:                 db,
		InvCheckPeriod:     0,
		EncConfig:          encCfg,
		HomePath:           DefaultNodeHome,
		SkipUpgradeHeights: map[int64]bool{},
		AppOpts:            EmptyAppOptions{},
	})

	for acc := range maccPerms {
		require.Equal(t, true, app.BankKeeper.BlockedAddr(app.AccountKeeper.GetModuleAddress(acc)),
			"ensure that all module account addresses are properly blocked in bank keeper")
	}

	app.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewRegenApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, encCfg, EmptyAppOptions{})
	_, err := app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

func TestGetMaccPerms(t *testing.T) {
	dup := GetMaccPerms()
	require.Equal(t, maccPerms, dup, "duplicated module account permissions differed from actual module account permissions")
}
