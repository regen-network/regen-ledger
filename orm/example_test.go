package orm_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/orm/testdata"
)

type GroupKeeper struct {
	key                      sdk.StoreKey
	groupTable               orm.AutoUInt64Table
	groupByAdminIndex        orm.Index
	groupMemberTable         orm.PrimaryKeyTable
	groupMemberByGroupIndex  orm.Index
	groupMemberByMemberIndex orm.Index
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

func NewGroupKeeper(storeKey sdk.StoreKey, cdc codec.Codec) GroupKeeper {
	k := GroupKeeper{key: storeKey}

	groupTableBuilder, err := orm.NewAutoUInt64TableBuilder(GroupTablePrefix, GroupTableSeqPrefix, storeKey, &testdata.GroupInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	// note: quite easy to mess with Index prefixes when managed outside. no fail fast on duplicates
	k.groupByAdminIndex, err = orm.NewIndex(groupTableBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		return []orm.RowID{[]byte(val.(*testdata.GroupInfo).Admin)}, nil
	})
	if err != nil {
		panic(err.Error())
	}
	k.groupTable = groupTableBuilder.Build()

	groupMemberTableBuilder, err := orm.NewPrimaryKeyTableBuilder(GroupMemberTablePrefix, storeKey, &testdata.GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	if err != nil {
		panic(err.Error())
	}

	k.groupMemberByGroupIndex, err = orm.NewIndex(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		group := val.(*testdata.GroupMember).Group
		return []orm.RowID{[]byte(group)}, nil
	})
	if err != nil {
		panic(err.Error())
	}
	k.groupMemberByMemberIndex, err = orm.NewIndex(groupMemberTableBuilder, GroupMemberByMemberIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		return []orm.RowID{[]byte(val.(*testdata.GroupMember).Member)}, nil
	})
	if err != nil {
		panic(err.Error())
	}
	k.groupMemberTable = groupMemberTableBuilder.Build()

	return k
}
