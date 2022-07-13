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
	groupMemberByWeightIndex orm.Index
}

var (
	GroupTablePrefix               byte = 0x0
	GroupTableSeqPrefix            byte = 0x1
	GroupByAdminIndexPrefix        byte = 0x2
	GroupByDescriptionIndexPrefix  byte = 0x3
	GroupMemberTablePrefix         byte = 0x4
	GroupMemberTableSeqPrefix      byte = 0x5
	GroupMemberTableIndexPrefix    byte = 0x6
	GroupMemberByGroupIndexPrefix  byte = 0x7
	GroupMemberByMemberIndexPrefix byte = 0x8
	GroupMemberByWeightIndexPrefix byte = 0x9
)

func NewGroupKeeper(storeKey sdk.StoreKey, cdc codec.Codec) GroupKeeper {
	k := GroupKeeper{key: storeKey}

	groupTableBuilder, err := orm.NewAutoUInt64TableBuilder(GroupTablePrefix, GroupTableSeqPrefix, storeKey, &testdata.GroupInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	// note: quite easy to mess with Index prefixes when managed outside. no fail fast on duplicates
	k.groupByAdminIndex, err = orm.NewIndex(groupTableBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]interface{}, error) {
		return []interface{}{val.(*testdata.GroupInfo).Admin.Bytes()}, nil
	}, testdata.GroupInfo{}.Admin.Bytes())
	if err != nil {
		panic(err.Error())
	}
	k.groupTable = groupTableBuilder.Build()

	groupMemberTableBuilder, err := orm.NewPrimaryKeyTableBuilder(GroupMemberTablePrefix, storeKey, &testdata.GroupMember{}, cdc)
	if err != nil {
		panic(err.Error())
	}

	k.groupMemberByGroupIndex, err = orm.NewIndex(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]interface{}, error) {
		group := val.(*testdata.GroupMember).Group
		return []interface{}{group.Bytes()}, nil
	}, testdata.GroupMember{}.Group.Bytes())
	if err != nil {
		panic(err.Error())
	}
	k.groupMemberByMemberIndex, err = orm.NewIndex(groupMemberTableBuilder, GroupMemberByMemberIndexPrefix, func(val interface{}) ([]interface{}, error) {
		return []interface{}{val.(*testdata.GroupMember).Member.Bytes()}, nil
	}, testdata.GroupMember{}.Member.Bytes())
	if err != nil {
		panic(err.Error())
	}
	k.groupMemberByWeightIndex, err = orm.NewIndex(groupMemberTableBuilder, GroupMemberByWeightIndexPrefix, func(val interface{}) ([]interface{}, error) {
		return []interface{}{val.(*testdata.GroupMember).Weight}, nil
	}, testdata.GroupMember{}.Weight)
	if err != nil {
		panic(err.Error())
	}
	k.groupMemberTable = groupMemberTableBuilder.Build()

	return k
}
