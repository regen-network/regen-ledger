package core

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestAddIssuers(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addrs := genAddrs(3)

	newAddrs := genAddrs(3)

	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo {
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	for _, addr := range addrs {
		err := s.stateStore.ClassIssuerStore().Insert(s.ctx, &ecocreditv1beta1.ClassIssuer{
			ClassId: 1,
			Issuer:  addr,
		})
		assert.NilError(t, err)
	}

	bzz := make([][]byte, 3)
	for i, newAddr := range newAddrs {
		bzz[i] = newAddr.Bytes()
	}
	_, err = s.k.UpdateClassIssuers(s.ctx, &v1beta1.MsgUpdateClassIssuers{
		Admin:         s.addr.String(),
		ClassId:       "C01",
		AddIssuers:    bzz,
		RemoveIssuers: nil,
	})
	assert.NilError(t, err)

	matched := make(map[string]struct{})
	it, err := s.stateStore.ClassIssuerStore().List(s.ctx, ecocreditv1beta1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(1))
	assert.NilError(t, err)
	for it.Next() {
		val, err := it.Value()
		assert.NilError(t, err)
		addr := types.AccAddress(val.Issuer).String()
		_, ok := matched[addr]
		assert.Equal(t, false, ok, "duplicate address in class issuer store %s", addr)
		matched[addr] = struct{}{}
	}
	assert.Equal(t, 6, len(matched), "expected to get 6 address matches, got %d", len(matched))
}


func genAddrs(x int) []types.AccAddress {
	addrs := make([]types.AccAddress, x)
	for i := 0; i < x; i++ {
		_,_,addr := testdata.KeyTestPubAddr()
		addrs[i] = addr
	}
	return addrs
}