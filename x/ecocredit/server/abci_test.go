package server

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/golang/mock/gomock"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	coretypes "github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/core"
)

type TestSuite struct {
	module     module.AppModule
	keeper     ProposalKeeper
	stateStore api.StateStore
	querier    sdk.Querier
	handler    govtypes.Handler
	sdkCtx     sdk.Context
	ctx        context.Context
}

func setupTest(t *testing.T) TestSuite {
	s := TestSuite{}
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	storeKey := sdk.NewKVStoreKey("test")
	cms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	ormDB, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	s.stateStore, err = api.NewStateStore(ormDB)
	ctrl := gomock.NewController(t)
	bankKeeper := mocks.NewMockBankKeeper(ctrl)
	paramsKeeper := mocks.NewMockParamKeeper(ctrl)
	s.keeper = core.NewKeeper(s.stateStore, bankKeeper, paramsKeeper)
	s.handler = NewCreditTypeProposalHandler(s.keeper)
	return s
}

func TestCreditTypeProposal_BasicValid(t *testing.T) {
	t.Parallel()
	s := setupTest(t)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Name:         "biodiversity",
		Unit:         "meters squared",
		Precision:    6,
	}
	err := s.handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "dude wheres my credit type?",
		Description: "and then?",
		CreditType:  ct,
	})
	assert.NilError(t, err)
	ct2, err := s.stateStore.CreditTypeTable().Get(s.ctx, "BIO")
	assert.NilError(t, err)
	assertCreditTypesEqual(t, ct, ct2)
}

func TestCreditTypeProposal_InvalidPrecision(t *testing.T) {
	t.Parallel()
	s := setupTest(t)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Name:         "biodiversity",
		Unit:         "meters squared",
		Precision:    3,
	}
	err := s.handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "My New Credit Type",
		Description: "its very cool and awesome",
		CreditType:  ct,
	})
	assert.ErrorContains(t, err, "credit type precision is currently locked to 6")
}

func TestCreditTypeProposal_InvalidAbbreviation(t *testing.T) {
	t.Parallel()
	s := setupTest(t)
	ct := &coretypes.CreditType{
		Abbreviation: "biO",
		Name:         "biodiversity",
		Unit:         "meters squared",
		Precision:    6,
	}
	err := s.handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "My New Credit Type",
		Description: "its very cool and awesome",
		CreditType:  ct,
	})
	assert.ErrorContains(t, err, "credit type abbreviation must be 1-3 uppercase latin letters")
}

func TestCreditTypeProposal_NoName(t *testing.T) {
	t.Parallel()
	s := setupTest(t)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Unit:         "meters squared",
		Precision:    6,
	}
	err := s.handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "My New Credit Type",
		Description: "its very cool and awesome",
		CreditType:  ct,
	})
	assert.ErrorContains(t, err, "name cannot be empty")
}

func TestCreditTypeProposal_NoUnit(t *testing.T) {
	t.Parallel()
	s := setupTest(t)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Name:         "FooBar",
		Precision:    6,
	}
	err := s.handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "My New Credit Type",
		Description: "its very cool and awesome",
		CreditType:  ct,
	})
	assert.ErrorContains(t, err, "unit cannot be empty")
}

func TestCreditTypeProposal_Duplicate(t *testing.T) {
	t.Parallel()
	s := setupTest(t)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Name:         "FooBar",
		Unit:         "many",
		Precision:    6,
	}
	var pulsarCreditType api.CreditType
	assert.NilError(t, ormutil.GogoToPulsarSlow(ct, &pulsarCreditType))
	assert.NilError(t, s.stateStore.CreditTypeTable().Insert(s.ctx, &pulsarCreditType))
	err := s.handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "My New Credit Type",
		Description: "its very cool and awesome",
		CreditType:  ct,
	})
	assert.ErrorContains(t, err, "could not insert credit type with abbreviation BIO")
}

func assertCreditTypesEqual(t *testing.T, ct *coretypes.CreditType, ct2 *api.CreditType) {
	assert.Check(t, ct != nil)
	assert.Check(t, ct2 != nil)

	assert.Equal(t, ct.Abbreviation, ct2.Abbreviation)
	assert.Equal(t, ct.Name, ct2.Name)
	assert.Equal(t, ct.Unit, ct2.Unit)
	assert.Equal(t, ct.Precision, ct2.Precision)
}
