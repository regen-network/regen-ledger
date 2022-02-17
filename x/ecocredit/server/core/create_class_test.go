package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestValidState(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gomock.Any(), gomock.Any()).Do(func(any, p *ecocredit.Params)  {
		p.AllowlistEnabled = false
		p.CreditClassFee = sdk.NewCoins(sdk.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).AnyTimes()

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(any, any,any, any).Return(nil)
	s.bankKeeper.EXPECT().BurnCoins(any, any, any).Return(nil)

	res, err := s.k.CreateClass(s.ctx, &v1beta1.MsgCreateClass{
		Admin:          s.addr.String(),
		Issuers:        []string{s.addr.String()},
		Metadata:       nil,
		CreditTypeName: "carbon",
	})
	assert.NilError(t, err, "error creating class: %+w", err)
	assert.Equal(t, res.ClassId, "C01")

	ci, err := s.stateStore.ClassInfoStore().GetByName(s.ctx, res.ClassId)
	assert.NilError(t, err)
	assert.Equal(t, res.ClassId, ci.Name)
	_, err = s.stateStore.ClassIssuerStore().Get(s.ctx, ci.Id, s.addr)
	assert.NilError(t, err)
}

func TestNotAuthorized(t *testing.T) {
	t.Parallel()
	s := setupBase(t)


	s.paramsKeeper.EXPECT().GetParamSet(gomock.Any(), gomock.Any()).Do(func(any, p *ecocredit.Params)  {
		p.AllowlistEnabled = true
		p.AllowedClassCreators = append(p.AllowedClassCreators, "foo")
	}).AnyTimes()
	_, err := s.k.CreateClass(s.ctx, &v1beta1.MsgCreateClass{
		Admin:          s.addr.String(),
		Issuers:        []string{s.addr.String()},
		Metadata:       nil,
		CreditTypeName: "carbon",
	})
	assert.ErrorContains(t, err, "is not allowed to create credit classes")

	s.paramsKeeper.EXPECT().GetParamSet(gomock.Any(), gomock.Any()).Do(func(any, p *ecocredit.Params)  {
		p.AllowlistEnabled = false
	}).AnyTimes()
	_, err = s.k.CreateClass(s.ctx, &v1beta1.MsgCreateClass{
		Admin:          s.addr.String(),
		Issuers:        []string{s.addr.String()},
		Metadata:       nil,
		CreditTypeName: "carbon",
	})
	assert.ErrorContains(t, err, "is not allowed to create credit classes")
}

func TestSequences(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any, p *ecocredit.Params)  {
		p.AllowlistEnabled = false
		p.CreditClassFee = sdk.NewCoins(sdk.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).AnyTimes()

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(any, any,any, any).Return(nil).AnyTimes()
	s.bankKeeper.EXPECT().BurnCoins(any, any, any).Return(nil).AnyTimes()

	res, err := s.k.CreateClass(s.ctx, &v1beta1.MsgCreateClass{
		Admin:          s.addr.String(),
		Issuers:        []string{s.addr.String()},
		Metadata:       nil,
		CreditTypeName: "carbon",
	})
	assert.NilError(t, err, "error creating class: %+w", err)

	res2, err := s.k.CreateClass(s.ctx, &v1beta1.MsgCreateClass{
		Admin:          s.addr.String(),
		Issuers:        []string{s.addr.String()},
		Metadata:       nil,
		CreditTypeName: "carbon",
	})
	assert.NilError(t, err, "error creating class: %+w", err)

	assert.Equal(t, res.ClassId, "C01")
	assert.Equal(t, res2.ClassId, "C02")
}
