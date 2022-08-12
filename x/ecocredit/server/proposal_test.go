package server

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	dbm "github.com/tendermint/tm-db"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
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
	accountKeeper.EXPECT().GetModuleAddress(ecocredit.ModuleName).Return(sdk.AccAddress{}).Times(1)
	accountKeeper.EXPECT().GetModuleAddress(basket.BasketSubModuleName).Return(sdk.AccAddress{}).Times(1)
	s.server = newServer(storeKey, paramtypes.Subspace{}, accountKeeper, bankKeeper, sdk.AccAddress(""))
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
	return s
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
