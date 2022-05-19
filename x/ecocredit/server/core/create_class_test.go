package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

func TestCreateClass_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	ccFee := core.DefaultParams().CreditClassFee[0]

	allowListEnabled := true
	creditClassFees := core.DefaultParams().CreditClassFee
	allowList := []string{s.addr.String()}
	utils.ExpectParamGet(&allowListEnabled, s.paramsKeeper, core.KeyAllowlistEnabled, 1)
	utils.ExpectParamGet(&allowList, s.paramsKeeper, core.KeyAllowedClassCreators, 1)
	utils.ExpectParamGet(&creditClassFees, s.paramsKeeper, core.KeyCreditClassFee, 1)
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(1)
	s.bankKeeper.EXPECT().BurnCoins(gmAny, gmAny, gmAny).Return(nil).Times(1)

	res, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              &ccFee,
	})
	assert.NilError(t, err, "error creating class: %+w", err)
	assert.Equal(t, res.ClassId, "C01")

	// check class info
	ci, err := s.stateStore.ClassTable().GetById(s.ctx, res.ClassId)
	assert.NilError(t, err)
	assert.Equal(t, res.ClassId, ci.Id)

	// check class issuer
	_, err = s.stateStore.ClassIssuerTable().Get(s.ctx, ci.Key, s.addr)
	assert.NilError(t, err)

	// check sequence number
	seq, err := s.stateStore.ClassSequenceTable().Get(s.ctx, "C")
	assert.NilError(t, err)
	assert.Equal(t, uint64(2), seq.NextSequence)
}

func TestCreateClass_Unauthorized(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// allowlist = true and sender is not in allowlist
	allowListEnabled := true
	allowList := []string{sdk.AccAddress("foobar").String()}
	utils.ExpectParamGet(&allowListEnabled, s.paramsKeeper, core.KeyAllowlistEnabled, 1)
	utils.ExpectParamGet(&allowList, s.paramsKeeper, core.KeyAllowedClassCreators, 1)
	_, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
	})
	assert.ErrorContains(t, err, "is not allowed to create credit classes")
}

func TestCreateClass_Sequence(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	ccFee := core.DefaultParams().CreditClassFee[0]
	gmAny := gomock.Any()

	allowListEnabled := false
	classFees := core.DefaultParams().CreditClassFee
	utils.ExpectParamGet(&allowListEnabled, s.paramsKeeper, core.KeyAllowlistEnabled, 1)
	utils.ExpectParamGet(&classFees, s.paramsKeeper, core.KeyCreditClassFee, 1)

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(2)
	s.bankKeeper.EXPECT().BurnCoins(gmAny, gmAny, gmAny).Return(nil).Times(2)

	res, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              &ccFee,
	})
	assert.NilError(t, err, "error creating class: %+w", err)

	utils.ExpectParamGet(&allowListEnabled, s.paramsKeeper, core.KeyAllowlistEnabled, 1)
	utils.ExpectParamGet(&classFees, s.paramsKeeper, core.KeyCreditClassFee, 1)
	res2, err := s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              &ccFee,
	})
	assert.NilError(t, err, "error creating class: %+w", err)

	assert.Equal(t, res.ClassId, "C01")
	assert.Equal(t, res2.ClassId, "C02")
}

func TestCreateClass_Fees(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	allowListEnabled := false
	classFees := core.DefaultParams().CreditClassFee
	utils.ExpectParamGet(&allowListEnabled, s.paramsKeeper, core.KeyAllowlistEnabled, 1)
	utils.ExpectParamGet(&classFees, s.paramsKeeper, core.KeyCreditClassFee, 1)

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
	utils.ExpectParamGet(&allowListEnabled, s.paramsKeeper, core.KeyAllowlistEnabled, 1)
	utils.ExpectParamGet(&classFees, s.paramsKeeper, core.KeyCreditClassFee, 1)
	_, err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.addr.String(),
		Issuers:          []string{s.addr.String()},
		Metadata:         "",
		CreditTypeAbbrev: "C",
		Fee:              &sdk.Coin{Denom: sdk.DefaultBondDenom, Amount: sdk.NewInt(1)},
	})
	assert.ErrorContains(t, err, "expected 20000000stake for fee, got 1stake")
}
