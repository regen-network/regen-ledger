package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestClasses(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)
	err = s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C02",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	res, err := s.k.Classes(s.ctx, &v1beta1.QueryClassesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Classes))
	assert.Equal(t, "C01", res.Classes[0].Name)

	res, err = s.k.Classes(s.ctx, &v1beta1.QueryClassesRequest{Pagination: &query.PageRequest{
		Limit:      1,
		CountTotal: true,
		Reverse:    true,
	}})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Classes))
	assert.Equal(t, uint64(2), res.Pagination.Total)
}

func TestClassInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	_, err = s.k.ClassInfo(s.ctx, &v1beta1.QueryClassInfoRequest{ClassId: "C02"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	res, err := s.k.ClassInfo(s.ctx, &v1beta1.QueryClassInfoRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, "C01", res.Info.Name)
	assert.DeepEqual(t, s.addr.Bytes(), res.Info.Admin)
}

func TestClassIssuers(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	addrs := genAddrs(2)
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.ClassIssuerStore().Insert(s.ctx, &ecocreditv1beta1.ClassIssuer{
		ClassId: 1,
		Issuer:  s.addr,
	}))
	assert.NilError(t, s.stateStore.ClassIssuerStore().Insert(s.ctx, &ecocreditv1beta1.ClassIssuer{
		ClassId: 1,
		Issuer:  addrs[0],
	}))
	assert.NilError(t, s.stateStore.ClassIssuerStore().Insert(s.ctx, &ecocreditv1beta1.ClassIssuer{
		ClassId: 1,
		Issuer:  addrs[1],
	}))

	res, err := s.k.ClassIssuers(s.ctx, &v1beta1.QueryClassIssuersRequest{ClassId:    "C01"})
	assert.NilError(t, err)
	assert.Equal(t, 3, len(res.Issuers))
	addrs = append(addrs, s.addr)
	for _, issuer := range res.Issuers {
		addr, _ := types.AccAddressFromBech32(issuer)
		found := false
		for _, addr2 := range addrs {
			if addr.Equals(addr2) {
				found = true
				break
			}
		}
		assert.Equal(t, true, found, "address %s not found in issuer list", issuer)
	}

	_, err = s.k.ClassIssuers(s.ctx, &v1beta1.QueryClassIssuersRequest{ClassId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	res, err = s.k.ClassIssuers(s.ctx, &v1beta1.QueryClassIssuersRequest{ClassId: "C01", Pagination: &query.PageRequest{Limit: 1, CountTotal: true}})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Issuers))
	assert.Equal(t, uint64(3), res.Pagination.Total)
}

func TestProjects(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)
	err= s.stateStore.ProjectInfoStore().Insert(s.ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            "P01",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	})
	assert.NilError(t, err)
	err= s.stateStore.ProjectInfoStore().Insert(s.ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            "P02",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	})
	assert.NilError(t, err)

	res, err := s.k.Projects(s.ctx, &v1beta1.QueryProjectsRequest{ClassId: "C01"})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Projects))
	assert.Equal(t, "US-CA", res.Projects[0].ProjectLocation)

	_, err = s.k.Projects(s.ctx, &v1beta1.QueryProjectsRequest{ClassId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	res, err = s.k.Projects(s.ctx, &v1beta1.QueryProjectsRequest{
		ClassId:    "C01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Projects))
	assert.Equal(t, uint64(2), res.Pagination.Total)
}

func TestProjectInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	err := s.stateStore.ProjectInfoStore().Insert(s.ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            "P01",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	})
	assert.NilError(t, err)

	res, err := s.k.ProjectInfo(s.ctx, &v1beta1.QueryProjectInfoRequest{ProjectId: "P01"})
	assert.NilError(t, err)
	assert.Equal(t, "P01", res.Info.Name)

	_, err = s.k.ProjectInfo(s.ctx, &v1beta1.QueryProjectInfoRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}