package utils

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/v3/assert"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

type baseSuite struct {
	t          *testing.T
	db         ormdb.ModuleDB
	stateStore api.StateStore
	ctx        context.Context
	addr       sdk.AccAddress
}

func setupBase(t *testing.T) *baseSuite {
	s := &baseSuite{t: t}
	var err error
	s.db, err = ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	s.stateStore, err = api.NewStateStore(s.db)
	assert.NilError(t, err)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	sdkCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(sdkCtx)
	_, _, s.addr = testdata.KeyTestPubAddr()

	return s
}

func TestUtils_GetBalance(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		BatchKey:       1,
		Address:        s.addr,
		TradableAmount: "100",
		RetiredAmount:  "100",
		EscrowedAmount: "100",
	}))

	bal, err := GetBalance(s.ctx, s.stateStore.BatchBalanceTable(), s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, "100", bal.TradableAmount)
	assert.Equal(t, "100", bal.RetiredAmount)
	assert.Equal(t, "100", bal.EscrowedAmount)

	noBalAddr := sdk.AccAddress("foobar")
	bal, err = GetBalance(s.ctx, s.stateStore.BatchBalanceTable(), noBalAddr, 1)
	assert.NilError(t, err)
	assert.Equal(t, bal.TradableAmount, "0")
	assert.Equal(t, bal.RetiredAmount, "0")
	assert.Equal(t, bal.EscrowedAmount, "0")
}

func TestUtils_GetNonNegativeFixedDecs(t *testing.T) {
	t.Parallel()
	precision := uint32(5)
	decStrs := []string{"100.32", "100.50", "100", "302"}
	decs, err := GetNonNegativeFixedDecs(precision, decStrs...)
	assert.NilError(t, err)
	assert.Equal(t, len(decStrs), len(decs))
	for i, ds := range decStrs {
		assert.Check(t, decs[i].String() == ds)
		assert.Check(t, decs[i].NumDecimalPlaces() <= precision)
	}

	// check error when one of the decimals has more places than the precision
	_, err = GetNonNegativeFixedDecs(2, "10.10", "10.31", "10.432")
	assert.ErrorContains(t, err, "10.432 exceeds maximum decimal places: 2")
}

func TestUtils_GetCreditTypeFromBatchDenom(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	creditType := &api.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "foo",
		Precision:    6,
	}
	assert.NilError(t, s.stateStore.CreditTypeTable().Insert(s.ctx, creditType))
	assert.NilError(t, s.stateStore.ClassTable().Insert(s.ctx, &api.Class{
		Id:               "C01",
		Admin:            s.addr,
		Metadata:         "foo",
		CreditTypeAbbrev: "C",
	}))
	batchDenom := "C01-000000-0000000-001"
	ct, err := GetCreditTypeFromBatchDenom(s.ctx, s.stateStore, batchDenom)
	assert.NilError(t, err)
	assert.DeepEqual(t, ct, creditType, cmpopts.IgnoreUnexported(api.CreditType{}))

	invalidDenom := "C02-0000000-0000000-001"
	ct, err = GetCreditTypeFromBatchDenom(s.ctx, s.stateStore, invalidDenom)
	assert.ErrorContains(t, err, "could not get class with ID C02")
}
