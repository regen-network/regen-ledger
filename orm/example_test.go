package orm

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/testutil/testdata"
)

type GroupKeeper struct {
	key                      sdk.StoreKey
	groupTable               AutoUInt64Table
	groupByAdminIndex        Index
	groupMemberTable         NaturalKeyTable
	groupMemberByGroupIndex  Index
	groupMemberByMemberIndex Index
}

var (
	GroupTablePrefix               byte = 0x0
	GroupTableSeqPrefix            byte = 0x1
	GroupByAdminIndexPrefix        byte = 0x2
	GroupMemberTablePrefix         byte = 0x3
	GroupMemberTableSeqPrefix      byte = 0x4
	GroupMemberTableIndexPrefix    byte = 0x5
	GroupMemberByGroupIndexPrefix  byte = 0x6
	GroupMemberByMemberIndexPrefix byte = 0x7
)

func NewGroupKeeper(storeKey sdk.StoreKey, cdc codec.Marshaler) GroupKeeper {
	k := GroupKeeper{key: storeKey}

	groupTableBuilder := NewAutoUInt64TableBuilder(GroupTablePrefix, GroupTableSeqPrefix, storeKey, &testdata.GroupInfo{}, cdc)
	// note: quite easy to mess with Index prefixes when managed outside. no fail fast on duplicates
	k.groupByAdminIndex = NewIndex(groupTableBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]RowID, error) {
		return []RowID{[]byte(val.(*testdata.GroupInfo).Admin)}, nil
	})
	k.groupTable = groupTableBuilder.Build()

	groupMemberTableBuilder := NewNaturalKeyTableBuilder(GroupMemberTablePrefix, storeKey, &testdata.GroupMember{}, Max255DynamicLengthIndexKeyCodec{}, cdc)

	k.groupMemberByGroupIndex = NewIndex(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]RowID, error) {
		group := val.(*testdata.GroupMember).Group
		return []RowID{[]byte(group)}, nil
	})
	k.groupMemberByMemberIndex = NewIndex(groupMemberTableBuilder, GroupMemberByMemberIndexPrefix, func(val interface{}) ([]RowID, error) {
		return []RowID{[]byte(val.(*testdata.GroupMember).Member)}, nil
	})
	k.groupMemberTable = groupMemberTableBuilder.Build()

	return k
}
