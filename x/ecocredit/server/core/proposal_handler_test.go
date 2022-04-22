package core

import (
	"testing"

	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	coretypes "github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestCreditTypeProposal_BasicValid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	handler := NewCreditTypeProposalHandler(s.k)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Name:         "biodiversity",
		Unit:         "meters squared",
		Precision:    6,
	}
	err := handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "my credit type",
		Description: "yay",
		CreditType:  ct,
	})
	assert.NilError(t, err)
	ct2, err := s.stateStore.CreditTypeTable().Get(s.ctx, "BIO")
	assert.NilError(t, err)
	assertCreditTypesEqual(t, ct, ct2)
}

func TestCreditTypeProposal_InvalidPrecision(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	handler := NewCreditTypeProposalHandler(s.k)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Name:         "biodiversity",
		Unit:         "meters squared",
		Precision:    3,
	}
	err := handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "My New Credit Type",
		Description: "its very cool and awesome",
		CreditType:  ct,
	})
	assert.ErrorContains(t, err, "credit type precision is currently locked to 6")
}

func TestCreditTypeProposal_InvalidAbbreviation(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	handler := NewCreditTypeProposalHandler(s.k)
	ct := &coretypes.CreditType{
		Abbreviation: "biO",
		Name:         "biodiversity",
		Unit:         "meters squared",
		Precision:    6,
	}
	err := handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "My New Credit Type",
		Description: "its very cool and awesome",
		CreditType:  ct,
	})
	assert.ErrorContains(t, err, "credit type abbreviation must be 1-3 uppercase latin letters")
}

func TestCreditTypeProposal_NoName(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	handler := NewCreditTypeProposalHandler(s.k)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Unit:         "meters squared",
		Precision:    6,
	}
	err := handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "My New Credit Type",
		Description: "its very cool and awesome",
		CreditType:  ct,
	})
	assert.ErrorContains(t, err, "name cannot be empty")
}

func TestCreditTypeProposal_NoUnit(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	handler := NewCreditTypeProposalHandler(s.k)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Name:         "FooBar",
		Precision:    6,
	}
	err := handler(s.sdkCtx, &coretypes.CreditTypeProposal{
		Title:       "My New Credit Type",
		Description: "its very cool and awesome",
		CreditType:  ct,
	})
	assert.ErrorContains(t, err, "unit cannot be empty")
}

func TestCreditTypeProposal_Duplicate(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	handler := NewCreditTypeProposalHandler(s.k)
	ct := &coretypes.CreditType{
		Abbreviation: "BIO",
		Name:         "FooBar",
		Unit:         "many",
		Precision:    6,
	}
	var pulsarCreditType api.CreditType
	assert.NilError(t, ormutil.GogoToPulsarSlow(ct, &pulsarCreditType))
	assert.NilError(t, s.stateStore.CreditTypeTable().Insert(s.ctx, &pulsarCreditType))
	err := handler(s.sdkCtx, &coretypes.CreditTypeProposal{
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
