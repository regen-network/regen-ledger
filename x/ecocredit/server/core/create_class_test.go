package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestCreateClass_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	any := gomock.Any()
	ccFee := &sdk.Coin{Denom: "foo", Amount: sdk.NewInt(20)}
	setupParams(s, false, nil, *ccFee, "C")

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(any, any, any, sdk.Coins{*ccFee}).Return(nil).Times(1)
	s.bankKeeper.EXPECT().BurnCoins(any, any, any).Return(nil).Times(1)

	res, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              ccFee,
	})
	assert.NilError(t, err, "error creating class: %+w", err)
	assert.Equal(t, res.ClassId, "C01")

	// check class info
	ci, err := s.stateStore.ClassInfoTable().GetByName(s.ctx, res.ClassId)
	assert.NilError(t, err)
	assert.Equal(t, res.ClassId, ci.Name)

	// check class issuer
	_, err = s.stateStore.ClassIssuerTable().Get(s.ctx, ci.Id, s.addr)
	assert.NilError(t, err)

	// check sequence number
	seq, err := s.stateStore.ClassSequenceTable().Get(s.ctx, "C")
	assert.NilError(t, err)
	assert.Equal(t, uint64(2), seq.NextClassId)
}

func TestCreateClass_Unauthorized(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	// allowlist = true and sender is not in allowlist
	setupParams(s, true, []sdk.AccAddress{sdk.AccAddress("foo")}, sdk.NewInt64Coin("foo", 3), "C")
	_, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.ErrorContains(t, err, "class creation is currently limited to addresses in the governance controlled class creation allowlist")
}

func TestCreateClass_Sequence(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	ccFee := &sdk.Coin{Denom: "foo", Amount: sdk.NewInt(20)}
	gmAny := gomock.Any()
	setupParams(s, true, []sdk.AccAddress{s.addr}, *ccFee, "C")

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(2)
	s.bankKeeper.EXPECT().BurnCoins(gmAny, gmAny, gmAny).Return(nil).Times(2)

	res, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              ccFee,
	})
	assert.NilError(t, err, "error creating class: %+w", err)

	res2, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              ccFee,
	})
	assert.NilError(t, err, "error creating class: %+w", err)

	assert.Equal(t, res.ClassId, "C01")
	assert.Equal(t, res2.ClassId, "C02")
}

func TestCreateClass_Fees(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	setupParams(s, false, nil, sdk.NewInt64Coin("foo", 20), "C")

	// wrong denom
	_, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              &sdk.Coin{Denom: "bar", Amount: sdk.NewInt(10)},
	})
	assert.ErrorContains(t, err, "bar is not allowed to be used in credit class fees")

	// fee too low
	_, err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              &sdk.Coin{Denom: "foo", Amount: sdk.NewInt(10)},
	})
	assert.ErrorContains(t, err, "expected 20foo for fee, got 10foo")
}

func setupParams(s *baseSuite, allowlistEnabled bool, allowedCreators []sdk.AccAddress, classFee sdk.Coin, ctAbbrev string) {
	assert.NilError(s.t, s.stateStore.AllowlistEnabledTable().Save(s.ctx, &api.AllowlistEnabled{Enabled: allowlistEnabled}))
	assert.NilError(s.t, s.stateStore.CreditClassFeeTable().Insert(s.ctx, &api.CreditClassFee{Amount: classFee.Amount.String(), Denom: classFee.Denom}))
	assert.NilError(s.t, s.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{Abbreviation: ctAbbrev}))
	for _, creator := range allowedCreators {
		assert.NilError(s.t, s.stateStore.AllowedClassCreatorTable().Insert(s.ctx, &api.AllowedClassCreator{Address: creator}))
	}
}
