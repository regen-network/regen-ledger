package group

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateGroup(t *testing.T) {
	myAdmin := []byte("valid--admin-address")

	specs := map[string]struct {
		src        MsgCreateGroup
		expErr     *errors.Error
		expGroup   GroupMetadata
		expMembers []GroupMember
	}{
		"happy path": {
			src: MsgCreateGroup{
				Admin:   myAdmin,
				Comment: "test",
				Members: []Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   sdk.NewDec(1),
					Comment: "first",
				}},
			},
			expGroup: GroupMetadata{
				Group:       1,
				Admin:       myAdmin,
				Comment:     "test",
				Version:     1,
				TotalWeight: sdk.OneDec(),
			},
			expMembers: []GroupMember{
				{
					Member:  sdk.AccAddress([]byte("valid-member-address")),
					Group:   1,
					Weight:  sdk.NewDec(1),
					Comment: "first",
				},
			},
		},
		"invalid message": {
			src:    MsgCreateGroup{},
			expErr: ErrEmpty,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			k, ctx := createGroupKeeper()
			res, err := NewHandler(k)(ctx, spec.src)
			require.True(t, spec.expErr.Is(err), err)
			if spec.expErr != nil {
				return
			}
			// then
			groupID := orm.DecodeSequence(res.Data)
			loaded, err := k.GetGroup(ctx, GroupID(groupID))
			require.NoError(t, err)
			assert.Equal(t, spec.expGroup, loaded)

			// and members persisted
			it, err := k.groupMemberByGroupIndex.Get(ctx, groupID)
			require.NoError(t, err)
			var loadedMembers []GroupMember
			_, err = orm.ReadAll(it, &loadedMembers)
			require.NoError(t, err)
			assert.Equal(t, spec.expMembers, loadedMembers)
		})
	}
}

func TestMsgUpdateGroupAdmin(t *testing.T) {
	k, pCtx := createGroupKeeper()

	members := []Member{{
		Address: sdk.AccAddress([]byte("valid-member-address")),
		Power:   sdk.NewDec(1),
		Comment: "first member",
	}}
	oldAdmin := []byte("my-old-admin-address")
	groupID, err := k.CreateGroup(pCtx, oldAdmin, members, "test")
	require.NoError(t, err)

	specs := map[string]struct {
		src       MsgUpdateGroupAdmin
		expStored GroupMetadata
		expErr    *errors.Error
	}{
		"with correct admin": {
			src: MsgUpdateGroupAdmin{
				Group:    groupID,
				Admin:    oldAdmin,
				NewAdmin: []byte("my-new-admin-address"),
			},
			expStored: GroupMetadata{
				Group:       groupID,
				Admin:       []byte("my-new-admin-address"),
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     2,
			},
		},
		"with wrong admin": {
			src: MsgUpdateGroupAdmin{
				Group:    groupID,
				Admin:    []byte("unknown-address"),
				NewAdmin: []byte("my-new-admin-address"),
			},
			expErr: ErrUnauthorized,
			expStored: GroupMetadata{
				Group:       groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     1,
			},
		},
		"with unknown groupID": {
			src: MsgUpdateGroupAdmin{
				Group:    999,
				Admin:    oldAdmin,
				NewAdmin: []byte("my-new-admin-address"),
			},
			expErr: orm.ErrNotFound,
			expStored: GroupMetadata{
				Group:       groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     1,
			},
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			ctx, _ := pCtx.CacheContext()
			_, err := NewHandler(k)(ctx, spec.src)
			require.True(t, spec.expErr.Is(err), err)
			// then
			loaded, err := k.GetGroup(ctx, groupID)
			require.NoError(t, err)
			assert.Equal(t, spec.expStored, loaded)
		})
	}
}

func TestMsgUpdateGroupComment(t *testing.T) {
	k, pCtx := createGroupKeeper()

	oldComment := "first"
	members := []Member{{
		Address: sdk.AccAddress([]byte("valid-member-address")),
		Power:   sdk.NewDec(1),
		Comment: oldComment,
	}}

	oldAdmin := []byte("my-old-admin-address")
	groupID, err := k.CreateGroup(pCtx, oldAdmin, members, "test")
	require.NoError(t, err)

	specs := map[string]struct {
		src       MsgUpdateGroupComment
		expErr    *errors.Error
		expStored GroupMetadata
	}{
		"with correct admin": {
			src: MsgUpdateGroupComment{
				Group:   groupID,
				Admin:   oldAdmin,
				Comment: "new comment",
			},
			expStored: GroupMetadata{
				Group:       groupID,
				Admin:       oldAdmin,
				Comment:     "new comment",
				TotalWeight: sdk.NewDec(1),
				Version:     2,
			},
		},
		"with wrong admin": {
			src: MsgUpdateGroupComment{
				Group:   groupID,
				Admin:   []byte("unknown-address"),
				Comment: "new comment",
			},
			expErr: ErrUnauthorized,
			expStored: GroupMetadata{
				Group:       groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     1,
			},
		},
		"with unknown groupid": {
			src: MsgUpdateGroupComment{
				Group:   999,
				Admin:   []byte("unknown-address"),
				Comment: "new comment",
			},
			expErr: orm.ErrNotFound,
			expStored: GroupMetadata{
				Group:       groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     1,
			},
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			ctx, _ := pCtx.CacheContext()
			_, err := NewHandler(k)(ctx, spec.src)
			require.True(t, spec.expErr.Is(err), err)
			// then
			loaded, err := k.GetGroup(ctx, groupID)
			require.NoError(t, err)
			assert.Equal(t, spec.expStored, loaded)
		})
	}
}

func TestMsgUpdateGroupMembers(t *testing.T) {
	k, pCtx := createGroupKeeper()

	members := []Member{{
		Address: sdk.AccAddress([]byte("valid-member-address")),
		Power:   sdk.NewDec(1),
		Comment: "first",
	}}

	myAdmin := []byte("valid--admin-address")
	groupID, err := k.CreateGroup(pCtx, myAdmin, members, "test")
	require.NoError(t, err)

	specs := map[string]struct {
		src        MsgUpdateGroupMembers
		expErr     *errors.Error
		expGroup   GroupMetadata
		expMembers []GroupMember
	}{
		"add new member": {
			src: MsgUpdateGroupMembers{
				Group: groupID,
				Admin: myAdmin,
				MemberUpdates: []Member{{
					Address: sdk.AccAddress([]byte("other-member-address")),
					Power:   sdk.NewDec(2),
					Comment: "second",
				}},
			},
			expGroup: GroupMetadata{
				Group:       groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(3),
				Version:     2,
			},
			expMembers: []GroupMember{
				{
					Member:  sdk.AccAddress([]byte("other-member-address")),
					Group:   groupID,
					Weight:  sdk.NewDec(2),
					Comment: "second",
				},
				{
					Member:  sdk.AccAddress([]byte("valid-member-address")),
					Group:   groupID,
					Weight:  sdk.NewDec(1),
					Comment: "first",
				},
			},
		},
		"update member": {
			src: MsgUpdateGroupMembers{
				Group: groupID,
				Admin: myAdmin,
				MemberUpdates: []Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   sdk.NewDec(2),
					Comment: "updated",
				}},
			},
			expGroup: GroupMetadata{
				Group:       groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(2),
				Version:     2,
			},
			expMembers: []GroupMember{
				{
					Member:  sdk.AccAddress([]byte("valid-member-address")),
					Group:   groupID,
					Weight:  sdk.NewDec(2),
					Comment: "updated",
				},
			},
		},
		"update member with same data": {
			src: MsgUpdateGroupMembers{
				Group: groupID,
				Admin: myAdmin,
				MemberUpdates: []Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   sdk.NewDec(1),
					Comment: "first",
				}},
			},
			expGroup: GroupMetadata{
				Group:       groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     2,
			},
			expMembers: []GroupMember{
				{
					Member:  sdk.AccAddress([]byte("valid-member-address")),
					Group:   groupID,
					Weight:  sdk.NewDec(1),
					Comment: "first",
				},
			},
		},
		"replace member": {
			src: MsgUpdateGroupMembers{
				Group: groupID,
				Admin: myAdmin,
				MemberUpdates: []Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   sdk.NewDec(0),
					Comment: "good bye",
				},
					{
						Address: sdk.AccAddress([]byte("my-new-member-addres")),
						Power:   sdk.NewDec(1),
						Comment: "welcome",
					}},
			},
			expGroup: GroupMetadata{
				Group:       groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     2,
			},
			expMembers: []GroupMember{{
				Member:  sdk.AccAddress([]byte("my-new-member-addres")),
				Group:   groupID,
				Weight:  sdk.NewDec(1),
				Comment: "welcome",
			}},
		},
		"remove existing member": {
			src: MsgUpdateGroupMembers{
				Group: groupID,
				Admin: myAdmin,
				MemberUpdates: []Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   sdk.NewDec(0),
					Comment: "good bye",
				}},
			},
			expGroup: GroupMetadata{
				Group:       groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(0),
				Version:     2,
			},
			expMembers: []GroupMember{},
		},
		"remove unknown member": {
			src: MsgUpdateGroupMembers{
				Group: groupID,
				Admin: myAdmin,
				MemberUpdates: []Member{{
					Address: sdk.AccAddress([]byte("unknown-member-addrs")),
					Power:   sdk.NewDec(0),
					Comment: "good bye",
				}},
			},
			expErr: orm.ErrNotFound,
			expGroup: GroupMetadata{
				Group:       groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     1,
			},
			expMembers: []GroupMember{{
				Member:  sdk.AccAddress([]byte("valid-member-address")),
				Group:   groupID,
				Weight:  sdk.NewDec(1),
				Comment: "first",
			}},
		},
		"with wrong admin": {
			src: MsgUpdateGroupMembers{
				Group: groupID,
				Admin: []byte("unknown-address"),
				MemberUpdates: []Member{{
					Address: sdk.AccAddress([]byte("other-member-address")),
					Power:   sdk.NewDec(2),
					Comment: "second",
				}},
			},
			expErr: ErrUnauthorized,
			expGroup: GroupMetadata{
				Group:       groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     1,
			},
			expMembers: []GroupMember{{
				Member:  sdk.AccAddress([]byte("valid-member-address")),
				Group:   groupID,
				Weight:  sdk.NewDec(1),
				Comment: "first",
			}},
		},
		"with unknown groupID": {
			src: MsgUpdateGroupMembers{
				Group: 999,
				Admin: myAdmin,
				MemberUpdates: []Member{{
					Address: sdk.AccAddress([]byte("other-member-address")),
					Power:   sdk.NewDec(2),
					Comment: "second",
				}},
			},
			expErr: orm.ErrNotFound,
			expGroup: GroupMetadata{
				Group:       groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: sdk.NewDec(1),
				Version:     1,
			},
			expMembers: []GroupMember{{
				Member:  sdk.AccAddress([]byte("valid-member-address")),
				Group:   groupID,
				Weight:  sdk.NewDec(1),
				Comment: "first",
			}},
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			ctx, _ := pCtx.CacheContext()
			_, err := NewHandler(k)(ctx, spec.src)
			require.True(t, spec.expErr.Is(err), err)
			// then
			loaded, err := k.GetGroup(ctx, groupID)
			require.NoError(t, err)
			assert.Equal(t, spec.expGroup, loaded)

			// and members persisted
			it, err := k.groupMemberByGroupIndex.Get(ctx, uint64(groupID))
			require.NoError(t, err)
			var loadedMembers []GroupMember
			_, err = orm.ReadAll(it, &loadedMembers)
			require.NoError(t, err)
			assert.Equal(t, spec.expMembers, loadedMembers)
		})
	}
}
