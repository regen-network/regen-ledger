package group

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gogo/protobuf/types"
	"github.com/regen-network/regen-ledger/orm"
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

	// ProposalBase Table
	ProposalBaseTablePrefix               byte = 0x30
	ProposalBaseTableSeqPrefix            byte = 0x31
	ProposalBaseByGroupAccountIndexPrefix byte = 0x32
	ProposalBaseByProposerIndexPrefix     byte = 0x33

	// Vote Table
	VoteTablePrefix               byte = 0x40
	VoteByProposalBaseIndexPrefix byte = 0x41
	VoteByVoterIndexPrefix        byte = 0x42
)

type ProposalI interface {
	orm.Persistent
	GetBase() ProposalBase
	SetBase(ProposalBase)
	GetMsgs() []sdk.Msg
	SetMsgs([]sdk.Msg) error
}

type Keeper struct {
	key               sdk.StoreKey
	proposalModelType reflect.Type

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

	// ProposalBase Table
	proposalTable             orm.AutoUInt64Table
	ProposalGroupAccountIndex orm.Index
	ProposalByProposerIndex   orm.Index

	// Vote Table
	voteTable               orm.NaturalKeyTable
	voteByProposalBaseIndex orm.UInt64Index
	voteByVoterIndex        orm.Index

	paramSpace paramstypes.Subspace
	router     sdk.Router
}

func NewGroupKeeper(storeKey sdk.StoreKey, paramSpace paramstypes.Subspace, router sdk.Router, proposalModel ProposalI) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(paramstypes.NewKeyTable().RegisterParamSet(&Params{}))
	}
	if storeKey == nil {
		panic("storeKey must not be nil")
	}

	if proposalModel == nil {
		panic("proposalModel must not be nil")
	}
	if router == nil {
		panic("router must not be nil")
	}
	tp := reflect.TypeOf(proposalModel)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}

	k := Keeper{key: storeKey, paramSpace: paramSpace, proposalModelType: tp, router: router}

	//
	// Group Table
	//
	groupTableBuilder := orm.NewTableBuilder(GroupTablePrefix, storeKey, &GroupMetadata{}, orm.FixLengthIndexKeys(orm.EncodedSeqLength))
	k.groupSeq = orm.NewSequence(storeKey, GroupTableSeqPrefix)
	k.groupByAdminIndex = orm.NewIndex(groupTableBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		return []orm.RowID{val.(*GroupMetadata).Admin.Bytes()}, nil
	})
	k.groupTable = groupTableBuilder.Build()

	//
	// Group Member Table
	//
	groupMemberTableBuilder := orm.NewNaturalKeyTableBuilder(GroupMemberTablePrefix, storeKey, &GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{})
	k.groupMemberByGroupIndex = orm.NewUInt64Index(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]uint64, error) {
		group := val.(*GroupMember).Group
		return []uint64{uint64(group)}, nil
	})
	k.groupMemberByMemberIndex = orm.NewIndex(groupMemberTableBuilder, GroupMemberByMemberIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		member := val.(*GroupMember).Member
		return []orm.RowID{member.Bytes()}, nil
	})
	k.groupMemberTable = groupMemberTableBuilder.Build()

	//
	// Group Account Table
	//
	k.groupAccountSeq = orm.NewSequence(storeKey, GroupAccountTableSeqPrefix)
	groupAccountTableBuilder := orm.NewNaturalKeyTableBuilder(GroupAccountTablePrefix, storeKey, &GroupAccountMetadata{}, orm.Max255DynamicLengthIndexKeyCodec{})
	k.groupAccountByGroupIndex = orm.NewUInt64Index(groupAccountTableBuilder, GroupAccountByGroupIndexPrefix, func(value interface{}) ([]uint64, error) {
		group := value.(*GroupAccountMetadata).Group
		return []uint64{uint64(group)}, nil
	})
	k.groupAccountByAdminIndex = orm.NewIndex(groupAccountTableBuilder, GroupAccountByAdminIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		admin := value.(*GroupAccountMetadata).Admin
		return []orm.RowID{admin.Bytes()}, nil
	})
	k.groupAccountTable = groupAccountTableBuilder.Build()

	// Proposal Table
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalBaseTablePrefix, ProposalBaseTableSeqPrefix, storeKey, proposalModel)
	k.ProposalGroupAccountIndex = orm.NewIndex(proposalTableBuilder, ProposalBaseByGroupAccountIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		account := value.(ProposalI).GetBase().GroupAccount
		return []orm.RowID{account.Bytes()}, nil

	})
	k.ProposalByProposerIndex = orm.NewIndex(proposalTableBuilder, ProposalBaseByProposerIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		proposers := value.(ProposalI).GetBase().Proposers
		r := make([]orm.RowID, len(proposers))
		for i := range proposers {
			r[i] = proposers[i].Bytes()
		}
		return r, nil
	})
	k.proposalTable = proposalTableBuilder.Build()

	//
	// Vote Table
	//
	voteTableBuilder := orm.NewNaturalKeyTableBuilder(VoteTablePrefix, storeKey, &Vote{}, orm.Max255DynamicLengthIndexKeyCodec{})
	k.voteByProposalBaseIndex = orm.NewUInt64Index(voteTableBuilder, VoteByProposalBaseIndexPrefix, func(value interface{}) ([]uint64, error) {
		return []uint64{uint64(value.(*Vote).Proposal)}, nil
	})
	k.voteByVoterIndex = orm.NewIndex(voteTableBuilder, VoteByVoterIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		return []orm.RowID{value.(*Vote).Voter.Bytes()}, nil
	})
	k.voteTable = voteTableBuilder.Build()

	return k
}

// MaxCommentSize returns the maximum length of a comment
func (k Keeper) MaxCommentSize(ctx sdk.Context) int {
	var result uint32
	k.paramSpace.Get(ctx, ParamMaxCommentLength, &result)
	return int(result)
}

func (k Keeper) CreateGroup(ctx sdk.Context, admin sdk.AccAddress, members Members, comment string) (GroupID, error) {
	if err := members.ValidateBasic(); err != nil {
		return 0, err
	}

	maxCommentSize := k.MaxCommentSize(ctx)
	if len(comment) > maxCommentSize {
		return 0, errors.Wrap(ErrMaxLimit, "group comment")
	}

	totalWeight := sdk.ZeroDec()
	for i := range members {
		m := members[i]
		if len(m.Comment) > maxCommentSize {
			return 0, errors.Wrap(ErrMaxLimit, "member comment")
		}
		totalWeight = totalWeight.Add(m.Power)
	}

	groupID := GroupID(k.groupSeq.NextVal(ctx))
	err := k.groupTable.Create(ctx, groupID.Bytes(), &GroupMetadata{
		Group:       groupID,
		Admin:       admin,
		Comment:     comment,
		Version:     1,
		TotalWeight: totalWeight,
	})
	if err != nil {
		return 0, errors.Wrap(err, "could not create group")
	}

	for i := range members {
		m := members[i]
		err := k.groupMemberTable.Create(ctx, &GroupMember{
			Group:   groupID,
			Member:  m.Address,
			Weight:  m.Power,
			Comment: m.Comment,
		})
		if err != nil {
			return 0, errors.Wrapf(err, "could not store member %d", i)
		}
	}
	return groupID, nil
}

func (k Keeper) GetGroup(ctx sdk.Context, id GroupID) (GroupMetadata, error) {
	var obj GroupMetadata
	return obj, k.groupTable.GetOne(ctx, id.Bytes(), &obj)
}

func (k Keeper) HasGroup(ctx sdk.Context, rowID orm.RowID) bool {
	return k.groupTable.Has(ctx, rowID)
}

func (k Keeper) UpdateGroup(ctx sdk.Context, g *GroupMetadata) error {
	g.Version++
	return k.groupTable.Save(ctx, g.Group.Bytes(), g)
}

func (k Keeper) GetParams(ctx sdk.Context) Params {
	var p Params
	k.paramSpace.GetParamSet(ctx, &p)
	return p
}

func (k Keeper) SetParams(ctx sdk.Context, params Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// CreateGroupAccount creates and persists a `GroupAccountMetadata`
func (k Keeper) CreateGroupAccount(ctx sdk.Context, admin sdk.AccAddress, groupID GroupID, policy DecisionPolicy, comment string) (sdk.AccAddress, error) {
	maxCommentSize := k.MaxCommentSize(ctx)
	if len(comment) > maxCommentSize {
		return nil, errors.Wrap(ErrMaxLimit,
			"group account comment")
	}

	g, err := k.GetGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if !g.Admin.Equals(admin) {
		return nil, errors.Wrap(errors.ErrUnauthorized, "not group admin")
	}
	accountAddr := AccountCondition(k.groupAccountSeq.NextVal(ctx)).Address()
	groupAccount, err := NewGroupAccountMetadata(
		accountAddr,
		groupID,
		admin,
		comment,
		1,
		policy,
	)
	if err != nil {
		return nil, err
	}

	if err := k.groupAccountTable.Create(ctx, &groupAccount); err != nil {
		return nil, errors.Wrap(err, "could not create group account")
	}
	return accountAddr, nil
}

func (k Keeper) HasGroupAccount(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.groupAccountTable.Has(ctx, address.Bytes())
}

func (k Keeper) GetGroupAccount(ctx sdk.Context, accountAddress sdk.AccAddress) (GroupAccountMetadata, error) {
	var obj GroupAccountMetadata
	return obj, k.groupAccountTable.GetOne(ctx, accountAddress.Bytes(), &obj)
}

func (k Keeper) UpdateGroupAccount(ctx sdk.Context, obj *GroupAccountMetadata) error {
	obj.Version++
	return k.groupAccountTable.Save(ctx, obj)
}

func (k Keeper) GetGroupByGroupAccount(ctx sdk.Context, accountAddress sdk.AccAddress) (GroupMetadata, error) {
	obj, err := k.GetGroupAccount(ctx, accountAddress)
	if err != nil {
		return GroupMetadata{}, errors.Wrap(err, "load group account")
	}
	return k.GetGroup(ctx, obj.Group)
}

func (k Keeper) GetGroupMembersByGroup(ctx sdk.Context, id GroupID) (orm.Iterator, error) {
	return k.groupMemberByGroupIndex.Get(ctx, id.Uint64())
}

func (k Keeper) Vote(ctx sdk.Context, id ProposalID, voters []sdk.AccAddress, choice Choice, comment string) error {
	maxCommentSize := k.MaxCommentSize(ctx)
	if len(comment) > maxCommentSize {
		return errors.Wrap(ErrMaxLimit, "comment")
	}
	if len(voters) == 0 {
		return errors.Wrap(ErrEmpty, "voters")
	}

	blockTime, err := types.TimestampProto(ctx.BlockTime())
	if err != nil {
		return err
	}
	proposal, err := k.GetProposal(ctx, id)
	if err != nil {
		return err
	}
	base := proposal.GetBase()
	if base.Status != ProposalStatusSubmitted {
		return errors.Wrap(ErrInvalid, "proposal not open")
	}
	votingPeriodEnd, err := types.TimestampFromProto(&base.Timeout)
	if err != nil {
		return err
	}
	if votingPeriodEnd.Before(ctx.BlockTime()) || votingPeriodEnd.Equal(ctx.BlockTime()) {
		return errors.Wrap(ErrExpired, "voting period has ended already")
	}
	var accountMetadata GroupAccountMetadata
	if err := k.groupAccountTable.GetOne(ctx, base.GroupAccount.Bytes(), &accountMetadata); err != nil {
		return errors.Wrap(err, "load group account")
	}
	if base.GroupAccountVersion != accountMetadata.Version {
		return errors.Wrap(ErrModified, "group account was modified")
	}

	electorate, err := k.GetGroup(ctx, accountMetadata.Group)
	if err != nil {
		return err
	}
	if electorate.Version != base.GroupVersion {
		return errors.Wrap(ErrModified, "group was modified")
	}

	// count and store votes
	for _, voterAddr := range voters {
		voter := GroupMember{Group: electorate.Group, Member: voterAddr}
		if err := k.groupMemberTable.GetOne(ctx, voter.NaturalKey(), &voter); err != nil {
			return errors.Wrapf(err, "address: %s", voterAddr)
		}
		newVote := Vote{
			Proposal:    id,
			Voter:       voterAddr,
			Choice:      choice,
			Comment:     comment,
			SubmittedAt: *blockTime,
		}
		if err := base.VoteState.Add(newVote, voter.Weight); err != nil {
			return errors.Wrap(err, "add new vote")
		}

		if err := k.voteTable.Create(ctx, &newVote); err != nil {
			return errors.Wrap(err, "store vote")
		}
	}

	// run tally with new votes to close early
	if err := doTally(ctx, &base, electorate, accountMetadata); err != nil {
		return err
	}

	proposal.SetBase(base)
	return k.proposalTable.Save(ctx, id.Uint64(), proposal)
}

func doTally(ctx sdk.Context, base *ProposalBase, electorate GroupMetadata, accountMetadata GroupAccountMetadata) error {
	policy := accountMetadata.GetDecisionPolicy()
	submittedAt, err := types.TimestampFromProto(&base.SubmittedAt)
	if err != nil {
		return err
	}
	switch result, err := policy.Allow(base.VoteState, electorate.TotalWeight, ctx.BlockTime().Sub(submittedAt)); {
	case err != nil:
		return errors.Wrap(err, "policy execution")
	case result == DecisionPolicyResult{Allow: true, Final: true}:
		base.Result = ProposalResultAccepted
		base.Status = ProposalStatusClosed
	case result == DecisionPolicyResult{Allow: false, Final: true}:
		base.Result = ProposalResultRejected
		base.Status = ProposalStatusClosed
	}
	return nil
}

// ExecProposal can be executed n times before the timeout. It will update the proposal status and executes the msg payload.
// There are no separate transactions for the payload messages so that it is a full atomic operation that
// would either succeed or fail.
func (k Keeper) ExecProposal(ctx sdk.Context, id ProposalID) error {
	proposal, err := k.GetProposal(ctx, id)
	if err != nil {
		return err
	}
	// check constraints
	base := proposal.GetBase()

	if base.Status != ProposalStatusSubmitted && base.Status != ProposalStatusClosed {
		return errors.Wrapf(ErrInvalid, "not possible with proposal status %s", base.Status.String())
	}

	var accountMetadata GroupAccountMetadata
	if err := k.groupAccountTable.GetOne(ctx, base.GroupAccount.Bytes(), &accountMetadata); err != nil {
		return errors.Wrap(err, "load group account")
	}

	storeUpdates := func() error {
		proposal.SetBase(base)
		return k.proposalTable.Save(ctx, id.Uint64(), proposal)
	}

	if base.Status == ProposalStatusSubmitted {
		if base.GroupAccountVersion != accountMetadata.Version {
			base.Result = ProposalResultUndefined
			base.Status = ProposalStatusAborted
			return storeUpdates()
		}

		electorate, err := k.GetGroup(ctx, accountMetadata.Group)
		if err != nil {
			return errors.Wrap(err, "load group")
		}

		if electorate.Version != base.GroupVersion {
			base.Result = ProposalResultUndefined
			base.Status = ProposalStatusAborted
			return storeUpdates()
		}
		if err := doTally(ctx, &base, electorate, accountMetadata); err != nil {
			return err
		}
	}

	// execute proposal payload
	if base.Status == ProposalStatusClosed && base.Result == ProposalResultAccepted && base.ExecutorResult != ProposalExecutorResultSuccess {
		logger := ctx.Logger().With("module", fmt.Sprintf("x/%s", ModuleName))
		ctx, flush := ctx.CacheContext()
		_, err := DoExecuteMsgs(ctx, k.router, accountMetadata.GroupAccount, proposal.GetMsgs())
		if err != nil {
			base.ExecutorResult = ProposalExecutorResultFailure
			proposalType := reflect.TypeOf(proposal).String()
			logger.Info("proposal execution failed", "cause", err, "type", proposalType, "proposalID", id)
		} else {
			base.ExecutorResult = ProposalExecutorResultSuccess
			flush()
		}
	}
	return storeUpdates()
}

func (k Keeper) GetProposal(ctx sdk.Context, id ProposalID) (ProposalI, error) {
	loaded := reflect.New(k.proposalModelType).Interface().(ProposalI)
	if _, err := k.proposalTable.GetOne(ctx, id.Uint64(), loaded); err != nil {
		return nil, errors.Wrap(err, "load proposal")
	}
	return loaded, nil
}

func (k Keeper) CreateProposal(ctx sdk.Context, accountAddress sdk.AccAddress, comment string, proposers []sdk.AccAddress, msgs []sdk.Msg) (ProposalID, error) {
	maxCommentSize := k.MaxCommentSize(ctx)
	if len(comment) > maxCommentSize {
		return 0, errors.Wrap(ErrMaxLimit, "comment")
	}

	account, err := k.GetGroupAccount(ctx, accountAddress.Bytes())

	if err != nil {
		return 0, errors.Wrap(err, "load group account")
	}

	g, err := k.GetGroup(ctx, account.Group)
	if err != nil {
		return 0, errors.Wrap(err, "get group by account")
	}

	// only members can propose
	for i := range proposers {
		if !k.groupMemberTable.Has(ctx, GroupMember{Group: g.Group, Member: proposers[i]}.NaturalKey()) {
			return 0, errors.Wrapf(ErrUnauthorized, "not in group: %s", proposers[i])
		}
	}

	if err := ensureMsgAuthZ(msgs, account.GroupAccount); err != nil {
		return 0, err
	}

	blockTime, err := types.TimestampProto(ctx.BlockTime())
	if err != nil {
		return 0, errors.Wrap(err, "block time conversion")
	}

	policy := account.GetDecisionPolicy()

	if policy == nil {
		return 0, errors.Wrap(ErrEmpty, "nil policy")
	}

	// prevent proposal that can not succeed
	err = policy.Validate(g)
	if err != nil {
		return 0, err
	}

	timeout := policy.GetTimeout()
	window, err := types.DurationFromProto(&timeout)
	if err != nil {
		return 0, errors.Wrap(err, "maxVotingWindow time conversion")
	}
	endTime, err := types.TimestampProto(ctx.BlockTime().Add(window))
	if err != nil {
		return 0, errors.Wrap(err, "end time conversion")
	}

	m := reflect.New(k.proposalModelType).Interface().(ProposalI)
	m.SetBase(ProposalBase{
		GroupAccount:        accountAddress,
		Comment:             comment,
		Proposers:           proposers,
		SubmittedAt:         *blockTime,
		GroupVersion:        g.Version,
		GroupAccountVersion: account.Version,
		Result:              ProposalResultUndefined,
		Status:              ProposalStatusSubmitted,
		ExecutorResult:      ProposalExecutorResultNotRun,
		Timeout:             *endTime,
		VoteState: Tally{
			YesCount:     sdk.ZeroDec(),
			NoCount:      sdk.ZeroDec(),
			AbstainCount: sdk.ZeroDec(),
			VetoCount:    sdk.ZeroDec(),
		},
	})
	if err := m.SetMsgs(msgs); err != nil {
		return 0, errors.Wrap(err, "create proposal")
	}

	id, err := k.proposalTable.Create(ctx, m)
	if err != nil {
		return 0, errors.Wrap(err, "create proposal")
	}
	return ProposalID(id), nil
}

func (k Keeper) GetVote(ctx sdk.Context, id ProposalID, voter sdk.AccAddress) (Vote, error) {
	var v Vote
	return v, k.voteTable.GetOne(ctx, Vote{Proposal: id, Voter: voter}.NaturalKey(), &v)
}
