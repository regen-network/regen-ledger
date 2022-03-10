package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestCreateClass_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(ctx interface{}, p *ecocredit.Params) {
		p.AllowlistEnabled = false
		p.CreditClassFee = sdk.NewCoins(sdk.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).Times(1)

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(any, any, any, any).Return(nil).Times(1)
	s.bankKeeper.EXPECT().BurnCoins(any, any, any).Return(nil).Times(1)

	res, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         nil,
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err, "error creating class: %+w", err)
	assert.Equal(t, res.ClassId, "C01")

	// check class info
	ci, err := s.stateStore.ClassInfoStore().GetByName(s.ctx, res.ClassId)
	assert.NilError(t, err)
	assert.Equal(t, res.ClassId, ci.Name)

	// check class issuer
	_, err = s.stateStore.ClassIssuerStore().Get(s.ctx, ci.Id, s.addr)
	assert.NilError(t, err)

	// check sequence number
	seq, err := s.stateStore.ClassSequenceStore().Get(s.ctx, "C")
	assert.NilError(t, err)
	assert.Equal(t, uint64(2), seq.NextClassId)
}

func TestCreateClass_Unauthorized(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	any := gomock.Any()

	// allowlist = true and sender is not in allowlist
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(ctx interface{}, p *ecocredit.Params) {
		p.AllowlistEnabled = true
		p.AllowedClassCreators = append(p.AllowedClassCreators, "foo")
	}).Times(1)
	_, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         nil,
		CreditTypeAbbrev: "C",
	})
	assert.ErrorContains(t, err, "is not allowed to create credit classes")
}

func TestCreateClass_Sequence(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(ctx interface{}, p *ecocredit.Params) {
		p.AllowlistEnabled = false
		p.CreditClassFee = sdk.NewCoins(sdk.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).Times(2)

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(any, any, any, any).Return(nil).Times(2)
	s.bankKeeper.EXPECT().BurnCoins(any, any, any).Return(nil).Times(2)

	res, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         nil,
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err, "error creating class: %+w", err)

	res2, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         nil,
		CreditTypeAbbrev: "C",
	})
	assert.NilError(t, err, "error creating class: %+w", err)

	assert.Equal(t, res.ClassId, "C01")
	assert.Equal(t, res2.ClassId, "C02")
}
