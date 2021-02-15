package server

import (
	"github.com/regen-network/regen-ledger/orm"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/group"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Group Table
	GroupTablePrefix        byte = 0x0
	GroupTableSeqPrefix     byte = 0x1
	GroupByAdminIndexPrefix byte = 0x2

	// Group Member Table
	GroupMemberTablePrefix         byte = 0x10
	GroupMemberByGroupIndexPrefix  byte = 0x11
	GroupMemberByMemberIndexPrefix byte = 0x12

	// Group Account Table
	GroupAccountTablePrefix        byte = 0x20
	GroupAccountTableSeqPrefix     byte = 0x21
	GroupAccountByGroupIndexPrefix byte = 0x22
	GroupAccountByAdminIndexPrefix byte = 0x23

	// Proposal Table
	ProposalTablePrefix               byte = 0x30
	ProposalTableSeqPrefix            byte = 0x31
	ProposalByGroupAccountIndexPrefix byte = 0x32
	ProposalByProposerIndexPrefix     byte = 0x33

	// Vote Table
	VoteTablePrefix           byte = 0x40
	VoteByProposalIndexPrefix byte = 0x41
	VoteByVoterIndexPrefix    byte = 0x42
)

type serverImpl struct {
	key    servermodule.RootModuleKey
	router sdk.Router

	accKeeper AccountKeeper

	// Group Table
	groupSeq          orm.Sequence
	groupTable        orm.Table
	groupByAdminIndex orm.Index

	// Group Member Table
	groupMemberTable         orm.NaturalKeyTable
	groupMemberByGroupIndex  orm.UInt64Index
	groupMemberByMemberIndex orm.Index

	// Group Account Table
	groupAccountSeq          orm.Sequence
	groupAccountTable        orm.NaturalKeyTable
	groupAccountByGroupIndex orm.UInt64Index
	groupAccountByAdminIndex orm.Index

	// Proposal Table
	proposalTable               orm.AutoUInt64Table
	proposalByGroupAccountIndex orm.Index
	proposalByProposerIndex     orm.Index

	// Vote Table
	voteTable           orm.NaturalKeyTable
	voteByProposalIndex orm.UInt64Index
	voteByVoterIndex    orm.Index
}

func newServer(storeKey servermodule.RootModuleKey, router sdk.Router, accKeeper AccountKeeper, cdc codec.Marshaler) serverImpl {
	s := serverImpl{key: storeKey, router: router, accKeeper: accKeeper}

	// Group Table
	groupTableBuilder := orm.NewTableBuilder(GroupTablePrefix, storeKey, &group.GroupInfo{}, orm.FixLengthIndexKeys(orm.EncodedSeqLength), cdc)
	s.groupSeq = orm.NewSequence(storeKey, GroupTableSeqPrefix)
	s.groupByAdminIndex = orm.NewIndex(groupTableBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		addr, err := sdk.AccAddressFromBech32(val.(*group.GroupInfo).Admin)
		if err != nil {
			return nil, err
		}
		return []orm.RowID{addr.Bytes()}, nil
	})
	s.groupTable = groupTableBuilder.Build()

	// Group Member Table
	groupMemberTableBuilder := orm.NewNaturalKeyTableBuilder(GroupMemberTablePrefix, storeKey, &group.GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	s.groupMemberByGroupIndex = orm.NewUInt64Index(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]uint64, error) {
		group := val.(*group.GroupMember).GroupId
		return []uint64{uint64(group)}, nil
	})
	s.groupMemberByMemberIndex = orm.NewIndex(groupMemberTableBuilder, GroupMemberByMemberIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		memberAddr := val.(*group.GroupMember).Member.Address
		addr, err := sdk.AccAddressFromBech32(memberAddr)
		if err != nil {
			return nil, err
		}
		return []orm.RowID{addr.Bytes()}, nil
	})
	s.groupMemberTable = groupMemberTableBuilder.Build()

	// Group Account Table
	s.groupAccountSeq = orm.NewSequence(storeKey, GroupAccountTableSeqPrefix)
	groupAccountTableBuilder := orm.NewNaturalKeyTableBuilder(GroupAccountTablePrefix, storeKey, &group.GroupAccountInfo{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	s.groupAccountByGroupIndex = orm.NewUInt64Index(groupAccountTableBuilder, GroupAccountByGroupIndexPrefix, func(value interface{}) ([]uint64, error) {
		group := value.(*group.GroupAccountInfo).GroupId
		return []uint64{uint64(group)}, nil
	})
	s.groupAccountByAdminIndex = orm.NewIndex(groupAccountTableBuilder, GroupAccountByAdminIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		admin := value.(*group.GroupAccountInfo).Admin
		addr, err := sdk.AccAddressFromBech32(admin)
		if err != nil {
			return nil, err
		}
		return []orm.RowID{addr.Bytes()}, nil
	})
	s.groupAccountTable = groupAccountTableBuilder.Build()

	// Proposal Table
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, storeKey, &group.Proposal{}, cdc)
	// proposalTableBuilder := orm.NewNaturalKeyTableBuilder(ProposalTablePrefix, key, &group.Proposal{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.proposalByGroupAccountIndex = orm.NewIndex(proposalTableBuilder, ProposalByGroupAccountIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		account := value.(*group.Proposal).GroupAccount
		addr, err := sdk.AccAddressFromBech32(account)
		if err != nil {
			return nil, err
		}
		return []orm.RowID{addr.Bytes()}, nil
	})
	s.proposalByProposerIndex = orm.NewIndex(proposalTableBuilder, ProposalByProposerIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		proposers := value.(*group.Proposal).Proposers
		r := make([]orm.RowID, len(proposers))
		for i := range proposers {
			addr, err := sdk.AccAddressFromBech32(proposers[i])
			if err != nil {
				return nil, err
			}
			r[i] = addr.Bytes()
		}
		return r, nil
	})
	s.proposalTable = proposalTableBuilder.Build()

	// Vote Table
	voteTableBuilder := orm.NewNaturalKeyTableBuilder(VoteTablePrefix, storeKey, &group.Vote{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	s.voteByProposalIndex = orm.NewUInt64Index(voteTableBuilder, VoteByProposalIndexPrefix, func(value interface{}) ([]uint64, error) {
		return []uint64{uint64(value.(*group.Vote).ProposalId)}, nil
	})
	s.voteByVoterIndex = orm.NewIndex(voteTableBuilder, VoteByVoterIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		addr, err := sdk.AccAddressFromBech32(value.(*group.Vote).Voter)
		if err != nil {
			return nil, err
		}
		return []orm.RowID{addr.Bytes()}, nil
	})
	s.voteTable = voteTableBuilder.Build()

	return s
}

func RegisterServices(configurator servermodule.Configurator, accountKeeper AccountKeeper) {
	impl := newServer(configurator.ModuleKey(), configurator.Router(), accountKeeper, configurator.Marshaler())
	group.RegisterMsgServer(configurator.MsgServer(), impl)
	group.RegisterQueryServer(configurator.QueryServer(), impl)
}
