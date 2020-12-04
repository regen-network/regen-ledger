package server

import (
	"github.com/regen-network/regen-ledger/orm"
	regenmodule "github.com/regen-network/regen-ledger/types/module"
	"github.com/regen-network/regen-ledger/x/group"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type serverImpl struct {
	storeKey   sdk.StoreKey
	paramSpace paramstypes.Subspace
	router     sdk.Router

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
	proposalTable             orm.AutoUInt64Table
	ProposalGroupAccountIndex orm.Index
	ProposalByProposerIndex   orm.Index

	// Vote Table
	voteTable           orm.NaturalKeyTable
	voteByProposalIndex orm.UInt64Index
	voteByVoterIndex    orm.Index
}

func newServer(storeKey sdk.StoreKey, paramSpace paramstypes.Subspace, router sdk.Router, cdc codec.Marshaler) serverImpl {
	s := serverImpl{storeKey: storeKey, paramSpace: paramSpace, router: router}

	// Group Table
	groupTableBuilder := orm.NewTableBuilder(GroupTablePrefix, storeKey, &group.GroupInfo{}, orm.FixLengthIndexKeys(orm.EncodedSeqLength), cdc)
	s.groupSeq = orm.NewSequence(storeKey, GroupTableSeqPrefix)
	s.groupByAdminIndex = orm.NewIndex(groupTableBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		return []orm.RowID{val.(*group.GroupInfo).Admin.Bytes()}, nil
	})
	s.groupTable = groupTableBuilder.Build()

	// Group Member Table
	groupMemberTableBuilder := orm.NewNaturalKeyTableBuilder(GroupMemberTablePrefix, storeKey, &group.GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	s.groupMemberByGroupIndex = orm.NewUInt64Index(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]uint64, error) {
		group := val.(*group.GroupMember).GroupId
		return []uint64{uint64(group)}, nil
	})
	s.groupMemberByMemberIndex = orm.NewIndex(groupMemberTableBuilder, GroupMemberByMemberIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		member := val.(*group.GroupMember).Member
		return []orm.RowID{member.Bytes()}, nil
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
		return []orm.RowID{admin.Bytes()}, nil
	})
	s.groupAccountTable = groupAccountTableBuilder.Build()

	// Proposal Table
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, storeKey, &group.Proposal{}, cdc)
	// proposalTableBuilder := orm.NewNaturalKeyTableBuilder(ProposalTablePrefix, storeKey, &group.Proposal{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.ProposalGroupAccountIndex = orm.NewIndex(proposalTableBuilder, ProposalByGroupAccountIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		account := value.(*group.Proposal).GroupAccount
		return []orm.RowID{account.Bytes()}, nil
	})
	s.ProposalByProposerIndex = orm.NewIndex(proposalTableBuilder, ProposalByProposerIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		proposers := value.(*group.Proposal).Proposers
		r := make([]orm.RowID, len(proposers))
		for i := range proposers {
			r[i] = proposers[i].Bytes()
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
		return []orm.RowID{value.(*group.Vote).Voter.Bytes()}, nil
	})
	s.voteTable = voteTableBuilder.Build()

	return s
}

func RegisterServices(
	storeKey sdk.StoreKey, paramSpace paramstypes.Subspace,
	router sdk.Router, configurator module.Configurator) {
	cfg, ok := configurator.(regenmodule.Configurator)
	// We need regen configurator's Marshaler in order to
	// instantiate new table builders so panicking if it's not the case
	// until we use this upgraded configurator in the cosmos sdk
	if !ok {
		panic("configurator should implement regenmodule.Configurator")
	}
	group.RegisterMsgServer(configurator.MsgServer(), newServer(storeKey, paramSpace, router, cfg.Marshaler()))
}
