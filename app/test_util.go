// Mostly copied from https://github.com/cosmos/gaia/tree/master/app
package app

import (
	"io"

	"github.com/tendermint/tendermint/libs/log"

	dbm "github.com/tendermint/tendermint/libs/db"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

// used for debugging by gaia/cmd/gaiadebug
// NOTE to not use this function with non-test code
func NewXrnAppUNSAFE(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp),
) (app *XrnApp, keyMain, keyStaking *sdk.KVStoreKey, stakingKeeper staking.Keeper) {

	app = NewXrnApp(logger, db, traceStore, loadLatest, invCheckPeriod, baseAppOptions...)
	return app, app.keyMain, app.keyStaking, app.stakingKeeper
}
