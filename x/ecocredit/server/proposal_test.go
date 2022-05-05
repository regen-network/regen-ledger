package server

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/golang/mock/gomock"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
)

type baseSuite struct {
	sdkCtx sdk.Context
	ctx    context.Context
	server serverImpl
}

func setup(t *testing.T) baseSuite {
	s := baseSuite{}
	storeKey := sdk.NewKVStoreKey("proposal_test.go")
	ctrl := gomock.NewController(t)
	accountKeeper := mocks.NewMockAccountKeeper(ctrl)
	bankKeeper := mocks.NewMockBankKeeper(ctrl)
	distKeeper := mocks.NewMockDistributionKeeper(ctrl)
	s.server = newServer(storeKey, paramtypes.Subspace{}, accountKeeper, bankKeeper, distKeeper)
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
	return s
}

func TestProposal_CreditType(t *testing.T) {
	t.Parallel()
	s := setup(t)
	handler := NewProposalHandler(s.server)
	creditTypeProposal := core.CreditTypeProposal{
		Title:       "carbon type",
		Description: "i would like to add a carbon type",
		CreditType: &core.CreditType{
			Abbreviation: "FOO",
			Name:         "FOOBAR",
			Unit:         "metric ton c02 equivalent",
			Precision:    6,
		},
	}
	err := handler(s.sdkCtx, &creditTypeProposal)
	assert.NilError(t, err)
	res, err := s.server.coreKeeper.CreditTypes(s.ctx, &core.QueryCreditTypesRequest{})
	assert.NilError(t, err)
	assert.Check(t, len(res.CreditTypes) == 1)
	assert.DeepEqual(t, creditTypeProposal.CreditType, res.CreditTypes[0])
}

func TestProposal_AllowedDenom(t *testing.T) {
	t.Parallel()
	s := setup(t)
	handler := NewProposalHandler(s.server)
	proposal := marketplace.AllowDenomProposal{
		Title:       "regen token",
		Description: "i would like to use the regen token in the marketplace",
		Denom: &marketplace.AllowedDenom{
			BankDenom:    "uregen",
			DisplayDenom: "regen",
			Exponent:     18,
		},
	}
	err := handler(s.sdkCtx, &proposal)
	assert.NilError(t, err)
	res, err := s.server.marketplaceKeeper.AllowedDenoms(s.ctx, &marketplace.QueryAllowedDenomsRequest{})
	assert.NilError(t, err)
	assert.Check(t, len(res.AllowedDenoms) == 1)
	assert.DeepEqual(t, proposal.Denom, res.AllowedDenoms[0])
}

func TestProposal_Invalid(t *testing.T) {
	t.Parallel()
	s := setup(t)
	handler := NewProposalHandler(s.server)
	err := handler(s.sdkCtx, nil)
	assert.ErrorContains(t, err, "unrecognized proposal content type")
}
