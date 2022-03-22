package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParams_CreditType(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	govAddr := sdk.AccAddress("foo")
	s.accountKeeper.EXPECT().GetModuleAddress(gomock.Any()).Return(govAddr).Times(3)
	_, err := s.k.AddCreditType(s.ctx, &core.MsgAddCreditType{
		CreditTypes: []*core.CreditType{
			{Abbreviation: "C", Name: "carbon", Unit: "tonnes", Precision: 6},
			{Abbreviation: "BIO", Name: "biodiversity", Unit: "acres", Precision: 1},
		},
		RootAddress: govAddr.String(),
	})
	assert.NilError(t, err)

	ct, err := s.stateStore.CreditTypeTable().Get(s.ctx, "C")
	assert.NilError(t, err)
	assert.Equal(t, "carbon", ct.Name)

	ct, err = s.stateStore.CreditTypeTable().Get(s.ctx, "BIO")
	assert.NilError(t, err)
	assert.Equal(t, "biodiversity", ct.Name)

	// cannot have duplicate abbreviations
	_, err = s.k.AddCreditType(s.ctx, &core.MsgAddCreditType{
		CreditTypes: []*core.CreditType{
			{Abbreviation: "C", Name: "carbon", Unit: "tonnes", Precision: 6},
		},
		RootAddress: govAddr.String(),
	})
	assert.ErrorContains(t, err, ormerrors.PrimaryKeyConstraintViolation.Error())

	// only gov
	_, err = s.k.AddCreditType(s.ctx, &core.MsgAddCreditType{RootAddress: s.addr.String()})
	assert.ErrorContains(t, err, "params can only be updated via governance")
}

func TestParams_Allowlist(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	govAddr := sdk.AccAddress("foo")
	s.accountKeeper.EXPECT().GetModuleAddress(gomock.Any()).Return(govAddr).Times(2)

	res, err := s.stateStore.AllowlistEnabledTable().Get(s.ctx)
	assert.NilError(t, err)

	_, err = s.k.ToggleAllowList(s.ctx, &core.MsgToggleAllowListRequest{
		RootAddress: govAddr.String(),
		Toggle:      !res.Enabled,
	})
	assert.NilError(t, err)

	res2, err := s.stateStore.AllowlistEnabledTable().Get(s.ctx)
	assert.NilError(t, err)
	assert.Equal(t, !res.Enabled, res2.Enabled)

	// only gov
	_, err = s.k.ToggleAllowList(s.ctx, &core.MsgToggleAllowListRequest{
		RootAddress: s.addr.String(),
		Toggle:      false,
	})
	assert.ErrorContains(t, err, "params can only be updated via governance")
}

func TestParams_UpdateAllowedClassCreators(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	govAddr := sdk.AccAddress("foo")
	s.accountKeeper.EXPECT().GetModuleAddress(gomock.Any()).Return(govAddr).Times(3)

	addr1, addr2 := sdk.AccAddress("bar"), sdk.AccAddress("baz")
	err := s.stateStore.AllowedClassCreatorsTable().Insert(s.ctx, &api.AllowedClassCreators{Address: addr1})
	assert.NilError(t, err)
	err = s.stateStore.AllowedClassCreatorsTable().Insert(s.ctx, &api.AllowedClassCreators{Address: addr2})
	assert.NilError(t, err)

	add1, add2 := sdk.AccAddress("add1"), sdk.AccAddress("add2")
	_, err = s.k.UpdateAllowedCreditClassCreators(s.ctx, &core.MsgUpdateAllowedCreditClassCreatorsRequest{
		RootAddress:    govAddr.String(),
		AddCreators:    []string{add1.String(), add2.String()},
		RemoveCreators: []string{addr1.String(), addr2.String()},
	})
	assert.NilError(t, err)

	it, err := s.stateStore.AllowedClassCreatorsTable().List(s.ctx, api.AllowedClassCreatorsAddressIndexKey{})
	assert.NilError(t, err)
	count := 0
	for it.Next() {
		val, err := it.Value()
		assert.NilError(t, err)

		acc := sdk.AccAddress(val.Address)
		assert.Equal(t, true, acc.Equals(add1) || acc.Equals(add2))
		count++
	}
	assert.Equal(t, 2, count)

	// no duplicates
	_, err = s.k.UpdateAllowedCreditClassCreators(s.ctx, &core.MsgUpdateAllowedCreditClassCreatorsRequest{
		RootAddress: govAddr.String(),
		AddCreators: []string{add1.String()},
	})
	assert.ErrorContains(t, err, ormerrors.PrimaryKeyConstraintViolation.Error())

	// only gov
	_, err = s.k.UpdateAllowedCreditClassCreators(s.ctx, &core.MsgUpdateAllowedCreditClassCreatorsRequest{RootAddress: s.addr.String()})
	assert.ErrorContains(t, err, "params can only be updated via governance")
}

func TestParams_UpdateClassFee(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	fee := sdk.NewInt64Coin("foo", 50)

	err := s.stateStore.CreditClassFeeTable().Insert(s.ctx, &api.CreditClassFee{
		Denom:  fee.Denom,
		Amount: fee.Amount.String(),
	})
	assert.NilError(t, err)

	addFee := sdk.NewInt64Coin("bar", 20)

	govAddr := sdk.AccAddress("foo")
	s.accountKeeper.EXPECT().GetModuleAddress(gomock.Any()).Return(govAddr).Times(3)

	_, err = s.k.UpdateCreditClassFee(s.ctx, &core.MsgUpdateCreditClassFeeRequest{
		RootAddress: govAddr.String(),
		AddFees: []*core.MsgUpdateCreditClassFeeRequest_Fee{
			{Denom: addFee.Denom, Amount: addFee.Amount.String()},
		},
		RemoveFees: []*core.MsgUpdateCreditClassFeeRequest_Fee{
			{Denom: fee.Denom, Amount: fee.Amount.String()},
		},
	})
	assert.NilError(t, err)

	res, err := s.stateStore.CreditClassFeeTable().Get(s.ctx, addFee.Denom)
	assert.NilError(t, err)
	assert.Equal(t, res.Amount, addFee.Amount.String())

	_, err = s.stateStore.CreditClassFeeTable().Get(s.ctx, fee.Denom)
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// no duplicates
	_, err = s.k.UpdateCreditClassFee(s.ctx, &core.MsgUpdateCreditClassFeeRequest{
		RootAddress: govAddr.String(),
		AddFees: []*core.MsgUpdateCreditClassFeeRequest_Fee{
			{Denom: addFee.Denom, Amount: addFee.Amount.String()},
		}})
	assert.ErrorContains(t, err, ormerrors.PrimaryKeyConstraintViolation.Error())

	// only gov
	_, err = s.k.UpdateCreditClassFee(s.ctx, &core.MsgUpdateCreditClassFeeRequest{RootAddress: s.addr.String()})
	assert.ErrorContains(t, err, "params can only be updated via governance")
}
