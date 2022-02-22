package core

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestUpdateAdmin(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addrs := genAddrs(1)
	newAdmin := addrs[0]

	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo {
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	_, err = s.k.UpdateClassAdmin(s.ctx, &v1beta1.MsgUpdateClassAdmin{
		Admin:    s.addr.String(),
		ClassId:  "C01",
		NewAdmin: newAdmin.String(),
	})
	assert.NilError(t, err)

	cInfo, err := s.stateStore.ClassInfoStore().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Check(t, newAdmin.Equals(types.AccAddress(cInfo.Admin)))
}

func TestUpdateAdminErrs(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addr := genAddrs(1)[0]
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo {
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	_, err = s.k.UpdateClassAdmin(s.ctx, &v1beta1.MsgUpdateClassAdmin{
		Admin:    addr.String(),
		ClassId:  "C01",
		NewAdmin: addr.String(),
	})
	assert.ErrorContains(t, err, "unauthorized")

	_, err = s.k.UpdateClassAdmin(s.ctx, &v1beta1.MsgUpdateClassAdmin{
		Admin:    addr.String(),
		ClassId:  "FOOBAR",
		NewAdmin: addr.String(),
	})
	assert.ErrorContains(t, err, "not found")

}

func TestUpdateIssuers(t *testing.T) {
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

	removeAddrs := addrs[1:]
	removeBz := make([][]byte, 3)
	for i, rmAddr := range removeAddrs {
		removeBz[i] = rmAddr.Bytes()
	}

	_, err = s.k.UpdateClassIssuers(s.ctx, &v1beta1.MsgUpdateClassIssuers{
		Admin:         s.addr.String(),
		ClassId:       "C01",
		AddIssuers:    nil,
		RemoveIssuers: removeBz,
	})
	assert.NilError(t, err)
	it, err = s.stateStore.ClassIssuerStore().List(s.ctx, ecocreditv1beta1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(1))
	assert.NilError(t, err)
	for it.Next() {
		val, err := it.Value()
		assert.NilError(t, err)
		addr := types.AccAddress(val.Issuer)
		for _, rmAddr := range removeAddrs {
			assert.Check(t, !addr.Equals(rmAddr), "%s was supposed to be deleted", rmAddr.String())
		}
	}
}

func TestUpdateIssuersErrs(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addr := genAddrs(1)[0]
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo {
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	_, err = s.k.UpdateClassIssuers(s.ctx, &v1beta1.MsgUpdateClassIssuers{
		Admin:        addr.String(),
		ClassId:       "C01",
		AddIssuers:    nil,
		RemoveIssuers: nil,
	})
	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())

	_, err = s.k.UpdateClassIssuers(s.ctx, &v1beta1.MsgUpdateClassIssuers{
		Admin:         s.addr.String(),
		ClassId:       "FOO",
		AddIssuers:    nil,
		RemoveIssuers: nil,
	})
	assert.ErrorContains(t, err, sdkerrors.ErrNotFound.Error())
}

func TestUpdateMetaData(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo {
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   []byte("foobar"),
		CreditType: "C",
	})
	assert.NilError(t, err)

	_, err = s.k.UpdateClassMetadata(s.ctx, &v1beta1.MsgUpdateClassMetadata{
		Admin:    s.addr.String(),
		ClassId:  "C01",
		Metadata: []byte("barfoo"),
	})
	assert.NilError(t, err)

	class, err := s.stateStore.ClassInfoStore().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte("barfoo"), class.Metadata)
}

func TestUpdateMetaDataErrs(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addr := genAddrs(1)[0]

	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1beta1.ClassInfo {
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	_, err = s.k.UpdateClassMetadata(s.ctx, &v1beta1.MsgUpdateClassMetadata{
		Admin:    s.addr.String(),
		ClassId:  "FOO",
		Metadata: nil,
	})
	assert.ErrorContains(t, err, sdkerrors.ErrNotFound.Error())


	_, err = s.k.UpdateClassMetadata(s.ctx, &v1beta1.MsgUpdateClassMetadata{
		Admin:   addr.String(),
		ClassId:  "C01",
		Metadata: []byte("FOO"),
	})
	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())

}

func genAddrs(x int) []types.AccAddress {
	addrs := make([]types.AccAddress, x)
	for i := 0; i < x; i++ {
		_,_,addr := testdata.KeyTestPubAddr()
		addrs[i] = addr
	}
	return addrs
}