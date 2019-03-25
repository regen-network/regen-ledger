package esp

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/util"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

var cdc *codec.Codec
var ctx sdk.Context
//var keeper Keeper

func setupTestInput() {
	db := dbm.NewMemDB()

	cdc = codec.New()

	espKey := sdk.NewKVStoreKey("espKey")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(espKey, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	//keeper = NewKeeper(espKey, agentKeeper, geoKeeper, cdc)
	ctx = sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
}

func TestMain(m *testing.M) {
	util.GodogMain(m, "esp", FeatureContext)
}

func FeatureContext(s *godog.Suite) {
	s.BeforeFeature(func(*gherkin.Feature) {
		setupTestInput()
	})
}

