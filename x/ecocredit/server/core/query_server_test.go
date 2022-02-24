package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/mock/gomock"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
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

func TestBatches(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	err := s.stateStore.ProjectInfoStore().Insert(s.ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            "P01",
		ClassId:         1,
		ProjectLocation: "US-CA",
		Metadata:        nil,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1beta1.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C01-20200101-20220101-001",
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1beta1.BatchInfo{
		ProjectId:  1,
		BatchDenom: "C01-20200101-20220101-002",
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))

	_, err = s.k.Batches(s.ctx, &v1beta1.QueryBatchesRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	res, err := s.k.Batches(s.ctx, &v1beta1.QueryBatchesRequest{ProjectId: "P01"})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Batches))
	assert.Equal(t, "C01-20200101-20220101-001", res.Batches[0].BatchDenom)

	res, err = s.k.Batches(s.ctx, &v1beta1.QueryBatchesRequest{ProjectId: "P01", Pagination: &query.PageRequest{Limit: 1, CountTotal: true}})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assert.Equal(t, uint64(2), res.Pagination.Total)
}

func TestBatchInfo(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-20200101-20220101-001"
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1beta1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom,
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))

	_, err := s.k.BatchInfo(s.ctx, &v1beta1.QueryBatchInfoRequest{BatchDenom: "A00-00000000-00000000-000"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	res, err := s.k.BatchInfo(s.ctx, &v1beta1.QueryBatchInfoRequest{BatchDenom: batchDenom})
	assert.NilError(t, err)
	assert.Equal(t, uint64(1), res.Info.ProjectId)
}

func TestBalance(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	noBalanceAddr := genAddrs(1)[0]
	batchDenom := "C01-20200101-20220101-001"
	tradable := "10.54321"
	retired := "50.3214"
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1beta1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom,
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchBalanceStore().Insert(s.ctx, &ecocreditv1beta1.BatchBalance{
		Address:  s.addr,
		BatchId:  1,
		Tradable: tradable,
		Retired:  retired,
	}))
	res, err := s.k.Balance(s.ctx, &v1beta1.QueryBalanceRequest{
		Account:    s.addr.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)
	assert.Equal(t, tradable, res.TradableAmount)
	assert.Equal(t, retired, res.RetiredAmount)

	res, err = s.k.Balance(s.ctx, &v1beta1.QueryBalanceRequest{
		Account:    noBalanceAddr.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)
	assert.Equal(t, "0", res.TradableAmount)
	assert.Equal(t, "0", res.RetiredAmount)

	_, err = s.k.Balance(s.ctx, &v1beta1.QueryBalanceRequest{
		Account:    s.addr.String(),
		BatchDenom: "A00-00000000-00000000-001",
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func TestSupply(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-20200101-20220101-001"
	tradable := "10.54321"
	retired := "50.3214"
	cancelled := "0.3215"
	assert.NilError(t, s.stateStore.BatchInfoStore().Insert(s.ctx, &ecocreditv1beta1.BatchInfo{
		ProjectId:  1,
		BatchDenom: batchDenom,
		Metadata:   nil,
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchSupplyStore().Insert(s.ctx, &ecocreditv1beta1.BatchSupply{
		BatchId:         1,
		TradableAmount:  tradable,
		RetiredAmount:   retired,
		CancelledAmount: cancelled,
	}))

	res, err := s.k.Supply(s.ctx, &v1beta1.QuerySupplyRequest{BatchDenom: batchDenom})
	assert.NilError(t, err)
	assert.Equal(t, tradable, res.TradableSupply)
	assert.Equal(t, retired, res.RetiredSupply)
	assert.Equal(t, cancelled, res.CancelledAmount)

	_, err = s.k.Supply(s.ctx, &v1beta1.QuerySupplyRequest{BatchDenom: "A00-00000000-00000000-001"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func TestCreditTypes(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	assert.NilError(t, s.stateStore.CreditTypeStore().Insert(s.ctx, &ecocreditv1beta1.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "a ton",
		Precision:    6,
	}))
	assert.NilError(t, s.stateStore.CreditTypeStore().Insert(s.ctx, &ecocreditv1beta1.CreditType{
		Abbreviation: "F",
		Name:         "foobar",
		Unit:         "foo per inch",
		Precision:    18,
	}))

	res, err := s.k.CreditTypes(s.ctx, &v1beta1.QueryCreditTypesRequest{})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.CreditTypes))
	assert.Equal(t, uint32(6), res.CreditTypes[0].Precision)
	assert.Equal(t, "foobar", res.CreditTypes[1].Name)
}

func TestParams(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	assert.NilError(t, s.stateStore.CreditTypeStore().Insert(s.ctx, &ecocreditv1beta1.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "a ton",
		Precision:    6,
	}))

	s.paramsKeeper.EXPECT().GetParamSet(gomock.Any(), gomock.Any()).SetArg(1, ecocredit.Params{
		CreditClassFee:       types.NewCoins(types.NewInt64Coin("foo", 30)),
		AllowedClassCreators: []string{s.addr.String()},
		AllowlistEnabled:     false,
		CreditTypes:          []*ecocredit.CreditType{{
			Abbreviation: "C",
			Name: "carbon",
			Unit: "a ton",
			Precision: 6,
		}},
	})

	res, err := s.k.Params(s.ctx, &v1beta1.QueryParamsRequest{})
	assert.NilError(t, err)
	assert.Equal(t,false, res.Params.AllowlistEnabled)
	assert.Equal(t, s.addr.String(), res.Params.AllowedClassCreators[0])
}