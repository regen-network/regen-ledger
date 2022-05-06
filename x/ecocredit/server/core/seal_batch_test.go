package core

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestSealBatch_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, batchDenom := s.setupClassProjectBatch(t)
	setupSealBatchTest(s, batchDenom)

	_, err := s.k.SealBatch(s.ctx, &core.MsgSealBatch{
		Issuer:     s.addr.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)

	batchAfter, err := s.stateStore.BatchTable().GetByDenom(s.ctx, batchDenom)
	assert.NilError(t, err)
	assert.Equal(t, false, batchAfter.Open)
}

func TestSealBatch_NoOp(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, batchDenom := s.setupClassProjectBatch(t)

	_, err := s.k.SealBatch(s.ctx, &core.MsgSealBatch{
		Issuer:     s.addr.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)
}

func TestSealBatch_Unauthorized(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, batchDenom := s.setupClassProjectBatch(t)
	setupSealBatchTest(s, batchDenom)

	notIssuer := sdk.AccAddress("foobar")

	_, err := s.k.SealBatch(s.ctx, &core.MsgSealBatch{
		Issuer:     notIssuer.String(),
		BatchDenom: batchDenom,
	})

	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())
}

func TestSealBatch_NotFound(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.setupClassProjectBatch(t)

	_, err := s.k.SealBatch(s.ctx, &core.MsgSealBatch{
		Issuer:     s.addr.String(),
		BatchDenom: "C02-00000000-00000000-001",
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func setupSealBatchTest(s *baseSuite, batchDenom string) {
	batchBefore, err := s.stateStore.BatchTable().GetByDenom(s.ctx, batchDenom)
	assert.NilError(s.t, err)
	batchBefore.Open = true
	assert.NilError(s.t, s.stateStore.BatchTable().Update(s.ctx, batchBefore))
}
