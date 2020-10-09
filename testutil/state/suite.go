package state

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/regen-network/regen-ledger/testutil"
)

// TestSuite encapsulates a functional Akash nodes data stores for
// ephemeral testing.
type TestSuite struct {
	t       testing.TB
	ms      sdk.CommitMultiStore
	ctx     sdk.Context
	bkeeper bankkeeper.Keeper
}

// SetupTestSuite provides toolkit for accessing stores and keepers
// for complex data interactions.
func SetupTestSuite(t testing.TB, codec codec.Marshaler) *TestSuite {
	suite := &TestSuite{
		t: t,
	}

	db := dbm.NewMemDB()
	suite.ms = store.NewCommitMultiStore(db)

	err := suite.ms.LoadLatestVersion()
	require.NoError(t, err)
	suite.ctx = sdk.NewContext(suite.ms, tmproto.Header{}, true, testutil.Logger(t))

	return suite
}

// SetBlockHeight provides arbitrarily setting the chain's block height.
func (ts *TestSuite) SetBlockHeight(height int64) {
	ts.ctx = ts.ctx.WithBlockHeight(height)
}

// Store provides access to the underlying KVStore
func (ts *TestSuite) Store() sdk.CommitMultiStore {
	return ts.ms
}

// Context of the current mempool
func (ts *TestSuite) Context() sdk.Context {
	return ts.ctx
}

// BankKeeper key store
func (ts *TestSuite) BankKeeper() bankkeeper.Keeper {
	return ts.bkeeper
}
