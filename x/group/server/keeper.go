package server

import (
	"fmt"
	"reflect"

	"github.com/cockroachdb/apd/v2"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/regen-network/regen-ledger/math"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
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

type Keeper struct {
	storeKey sdk.StoreKey

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

	paramSpace paramstypes.Subspace
	router     sdk.Router
}

func NewGroupKeeper(storeKey sdk.StoreKey, paramSpace paramstypes.Subspace, router sdk.Router, cdc codec.Marshaler) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(paramstypes.NewKeyTable().RegisterParamSet(&group.Params{}))
	}
	if storeKey == nil {
		panic("storeKey must not be nil")
	}
	if router == nil {
		panic("router must not be nil")
	}

	k := Keeper{storeKey: storeKey, paramSpace: paramSpace, router: router}

	// Group Table
	groupTableBuilder := orm.NewTableBuilder(GroupTablePrefix, storeKey, &group.GroupInfo{}, orm.FixLengthIndexKeys(orm.EncodedSeqLength), cdc)
	k.groupSeq = orm.NewSequence(storeKey, GroupTableSeqPrefix)
	k.groupByAdminIndex = orm.NewIndex(groupTableBuilder, GroupByAdminIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		return []orm.RowID{val.(*group.GroupInfo).Admin.Bytes()}, nil
	})
	k.groupTable = groupTableBuilder.Build()

	// Group Member Table
	groupMemberTableBuilder := orm.NewNaturalKeyTableBuilder(GroupMemberTablePrefix, storeKey, &group.GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	k.groupMemberByGroupIndex = orm.NewUInt64Index(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]uint64, error) {
		group := val.(*group.GroupMember).GroupId
		return []uint64{uint64(group)}, nil
	})
	k.groupMemberByMemberIndex = orm.NewIndex(groupMemberTableBuilder, GroupMemberByMemberIndexPrefix, func(val interface{}) ([]orm.RowID, error) {
		member := val.(*group.GroupMember).Member
		return []orm.RowID{member.Bytes()}, nil
	})
	k.groupMemberTable = groupMemberTableBuilder.Build()

	// Group Account Table
	k.groupAccountSeq = orm.NewSequence(storeKey, GroupAccountTableSeqPrefix)
	groupAccountTableBuilder := orm.NewNaturalKeyTableBuilder(GroupAccountTablePrefix, storeKey, &group.GroupAccountInfo{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	k.groupAccountByGroupIndex = orm.NewUInt64Index(groupAccountTableBuilder, GroupAccountByGroupIndexPrefix, func(value interface{}) ([]uint64, error) {
		group := value.(*group.GroupAccountInfo).GroupId
		return []uint64{uint64(group)}, nil
	})
	k.groupAccountByAdminIndex = orm.NewIndex(groupAccountTableBuilder, GroupAccountByAdminIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		admin := value.(*group.GroupAccountInfo).Admin
		return []orm.RowID{admin.Bytes()}, nil
	})
	k.groupAccountTable = groupAccountTableBuilder.Build()

	// Proposal Table
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, storeKey, &group.Proposal{}, cdc)
	// proposalTableBuilder := orm.NewNaturalKeyTableBuilder(ProposalTablePrefix, storeKey, &group.Proposal{}, orm.Max255DynamicLengthIndexKeyCodec{})
	k.ProposalGroupAccountIndex = orm.NewIndex(proposalTableBuilder, ProposalByGroupAccountIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		account := value.(*group.Proposal).GroupAccount
		return []orm.RowID{account.Bytes()}, nil
	})
	k.ProposalByProposerIndex = orm.NewIndex(proposalTableBuilder, ProposalByProposerIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		proposers := value.(*group.Proposal).Proposers
		r := make([]orm.RowID, len(proposers))
		for i := range proposers {
			r[i] = proposers[i].Bytes()
		}
		return r, nil
	})
	k.proposalTable = proposalTableBuilder.Build()

	// Vote Table
	voteTableBuilder := orm.NewNaturalKeyTableBuilder(VoteTablePrefix, storeKey, &group.Vote{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	k.voteByProposalIndex = orm.NewUInt64Index(voteTableBuilder, VoteByProposalIndexPrefix, func(value interface{}) ([]uint64, error) {
		return []uint64{uint64(value.(*group.Vote).ProposalId)}, nil
	})
	k.voteByVoterIndex = orm.NewIndex(voteTableBuilder, VoteByVoterIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		return []orm.RowID{value.(*group.Vote).Voter.Bytes()}, nil
	})
	k.voteTable = voteTableBuilder.Build()

	return k
}

// MaxCommentSize returns the maximum length of a comment
func (k Keeper) MaxCommentSize(ctx sdk.Context) int {
	var result uint32
	k.paramSpace.Get(ctx, group.ParamMaxCommentLength, &result)
	return int(result)
}

func (k Keeper) CreateGroup(ctx sdk.Context, admin sdk.AccAddress, members group.Members, comment string) (group.GroupID, error) {
	if err := members.ValidateBasic(); err != nil {
		return 0, err
	}

	maxCommentSize := k.MaxCommentSize(ctx)
	if err := assertCommentSize(comment, maxCommentSize, "group comment"); err != nil {
		return 0, err
	}

	totalWeight := apd.New(0, 0)
	for i := range members {
		m := members[i]
		if err := assertCommentSize(m.Comment, maxCommentSize, "member comment"); err != nil {
			return 0, err
		}

		power, err := math.ParseNonNegativeDecimal(m.Power)
		if err != nil {
			return 0, err
		}

		if !power.IsZero() {
			err = math.Add(totalWeight, totalWeight, power)
			if err != nil {
				return 0, err
			}
		}
	}

	groupID := group.GroupID(k.groupSeq.NextVal(ctx))
	err := k.groupTable.Create(ctx, groupID.Bytes(), &group.GroupInfo{
		GroupId:     groupID,
		Admin:       admin,
		Comment:     comment,
		Version:     1,
		TotalWeight: math.DecimalString(totalWeight),
	})
	if err != nil {
		return 0, sdkerrors.Wrap(err, "could not create group")
	}

	for i := range members {
		m := members[i]
		err := k.groupMemberTable.Create(ctx, &group.GroupMember{
			GroupId: groupID,
			Member:  m.Address,
			Weight:  m.Power,
			Comment: m.Comment,
		})
		if err != nil {
			return 0, sdkerrors.Wrapf(err, "could not store member %d", i)
		}
	}
	return groupID, nil
}

func (k Keeper) GetGroup(ctx sdk.Context, id group.GroupID) (group.GroupInfo, error) {
	var obj group.GroupInfo
	return obj, k.groupTable.GetOne(ctx, id.Bytes(), &obj)
}

func (k Keeper) HasGroup(ctx sdk.Context, rowID orm.RowID) bool {
	return k.groupTable.Has(ctx, rowID)
}

func (k Keeper) UpdateGroup(ctx sdk.Context, g *group.GroupInfo) error {
	g.Version++
	return k.groupTable.Save(ctx, g.GroupId.Bytes(), g)
}

func (k Keeper) GetParams(ctx sdk.Context) group.Params {
	var p group.Params
	k.paramSpace.GetParamSet(ctx, &p)
	return p
}

func (k Keeper) SetParams(ctx sdk.Context, params group.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// CreateGroupAccount creates and persists a `GroupAccountInfo`
func (k Keeper) CreateGroupAccount(ctx sdk.Context, admin sdk.AccAddress, groupID group.GroupID, policy group.DecisionPolicy, comment string) (sdk.AccAddress, error) {
	if err := assertCommentSize(comment, k.MaxCommentSize(ctx), "group account comment"); err != nil {
		return nil, err
	}

	g, err := k.GetGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if !g.Admin.Equals(admin) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not group admin")
	}
	accountAddr := group.AccountCondition(k.groupAccountSeq.NextVal(ctx)).Address()
	groupAccount, err := group.NewGroupAccountInfo(
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
		return nil, sdkerrors.Wrap(err, "could not create group account")
	}
	return accountAddr, nil
}

func (k Keeper) HasGroupAccount(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.groupAccountTable.Has(ctx, address.Bytes())
}

func (k Keeper) GetGroupAccount(ctx sdk.Context, accountAddress sdk.AccAddress) (group.GroupAccountInfo, error) {
	var obj group.GroupAccountInfo
	return obj, k.groupAccountTable.GetOne(ctx, accountAddress.Bytes(), &obj)
}

func (k Keeper) UpdateGroupAccount(ctx sdk.Context, obj *group.GroupAccountInfo) error {
	obj.Version++
	return k.groupAccountTable.Save(ctx, obj)
}

func (k Keeper) GetGroupByGroupAccount(ctx sdk.Context, accountAddress sdk.AccAddress) (group.GroupInfo, error) {
	obj, err := k.GetGroupAccount(ctx, accountAddress)
	if err != nil {
		return group.GroupInfo{}, sdkerrors.Wrap(err, "load group account")
	}
	return k.GetGroup(ctx, obj.GroupId)
}

func (k Keeper) GetGroupMembers(ctx sdk.Context, id group.GroupID) (orm.Iterator, error) {
	return k.groupMemberByGroupIndex.Get(ctx, id.Uint64())
}

func (k Keeper) Vote(ctx sdk.Context, id group.ProposalID, voters []sdk.AccAddress, choice group.Choice, comment string) error {
	if err := assertCommentSize(comment, k.MaxCommentSize(ctx), "comment"); err != nil {
		return err
	}
	if len(voters) == 0 {
		return sdkerrors.Wrap(group.ErrEmpty, "voters")
	}

	blockTime, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return err
	}
	proposal, err := k.GetProposal(ctx, id)
	if err != nil {
		return err
	}
	if proposal.Status != group.ProposalStatusSubmitted {
		return sdkerrors.Wrap(group.ErrInvalid, "proposal not open")
	}
	votingPeriodEnd, err := gogotypes.TimestampFromProto(&proposal.Timeout)
	if err != nil {
		return err
	}
	if votingPeriodEnd.Before(ctx.BlockTime()) || votingPeriodEnd.Equal(ctx.BlockTime()) {
		return sdkerrors.Wrap(group.ErrExpired, "voting period has ended already")
	}
	var accountInfo group.GroupAccountInfo
	if err := k.groupAccountTable.GetOne(ctx, proposal.GroupAccount.Bytes(), &accountInfo); err != nil {
		return sdkerrors.Wrap(err, "load group account")
	}
	if proposal.GroupAccountVersion != accountInfo.Version {
		return sdkerrors.Wrap(group.ErrModified, "group account was modified")
	}

	electorate, err := k.GetGroup(ctx, accountInfo.GroupId)
	if err != nil {
		return err
	}
	if electorate.Version != proposal.GroupVersion {
		return sdkerrors.Wrap(group.ErrModified, "group was modified")
	}

	// count and store votes
	for _, voterAddr := range voters {
		voter := group.GroupMember{GroupId: electorate.GroupId, Member: voterAddr}
		if err := k.groupMemberTable.GetOne(ctx, voter.NaturalKey(), &voter); err != nil {
			return sdkerrors.Wrapf(err, "address: %s", voterAddr)
		}
		newVote := group.Vote{
			ProposalId:  id,
			Voter:       voterAddr,
			Choice:      choice,
			Comment:     comment,
			SubmittedAt: *blockTime,
		}
		if err := proposal.VoteState.Add(newVote, voter.Weight); err != nil {
			return sdkerrors.Wrap(err, "add new vote")
		}

		if err := k.voteTable.Create(ctx, &newVote); err != nil {
			return sdkerrors.Wrap(err, "store vote")
		}
	}

	// run tally with new votes to close early
	if err := doTally(ctx, &proposal, electorate, accountInfo); err != nil {
		return err
	}

	return k.proposalTable.Save(ctx, id.Uint64(), &proposal)
}

func doTally(ctx sdk.Context, p *group.Proposal, electorate group.GroupInfo, accountInfo group.GroupAccountInfo) error {
	policy := accountInfo.GetDecisionPolicy()
	submittedAt, err := gogotypes.TimestampFromProto(&p.SubmittedAt)
	if err != nil {
		return err
	}
	switch result, err := policy.Allow(p.VoteState, electorate.TotalWeight, ctx.BlockTime().Sub(submittedAt)); {
	case err != nil:
		return sdkerrors.Wrap(err, "policy execution")
	case result.Allow && result.Final:
		p.Result = group.ProposalResultAccepted
		p.Status = group.ProposalStatusClosed
	case !result.Allow && result.Final:
		p.Result = group.ProposalResultRejected
		p.Status = group.ProposalStatusClosed
	}
	return nil
}

// ExecProposal can be executed n times before the timeout. It will update the proposal status and executes the msg payload.
// There are no separate transactions for the payload messages so that it is a full atomic operation that
// would either succeed or fail.
func (k Keeper) ExecProposal(ctx sdk.Context, id group.ProposalID) error {
	proposal, err := k.GetProposal(ctx, id)
	if err != nil {
		return err
	}

	if proposal.Status != group.ProposalStatusSubmitted && proposal.Status != group.ProposalStatusClosed {
		return sdkerrors.Wrapf(group.ErrInvalid, "not possible with proposal status %s", proposal.Status.String())
	}

	var accountInfo group.GroupAccountInfo
	if err := k.groupAccountTable.GetOne(ctx, proposal.GroupAccount.Bytes(), &accountInfo); err != nil {
		return sdkerrors.Wrap(err, "load group account")
	}

	storeUpdates := func() error {
		return k.proposalTable.Save(ctx, id.Uint64(), &proposal)
	}

	if proposal.Status == group.ProposalStatusSubmitted {
		if proposal.GroupAccountVersion != accountInfo.Version {
			proposal.Result = group.ProposalResultUnfinalized
			proposal.Status = group.ProposalStatusAborted
			return storeUpdates()
		}

		electorate, err := k.GetGroup(ctx, accountInfo.GroupId)
		if err != nil {
			return sdkerrors.Wrap(err, "load group")
		}

		if electorate.Version != proposal.GroupVersion {
			proposal.Result = group.ProposalResultUnfinalized
			proposal.Status = group.ProposalStatusAborted
			return storeUpdates()
		}
		if err := doTally(ctx, &proposal, electorate, accountInfo); err != nil {
			return err
		}
	}

	// execute proposal payload
	if proposal.Status == group.ProposalStatusClosed && proposal.Result == group.ProposalResultAccepted && proposal.ExecutorResult != group.ProposalExecutorResultSuccess {
		logger := ctx.Logger().With("module", fmt.Sprintf("x/%s", group.ModuleName))
		ctx, flush := ctx.CacheContext()
		_, err := DoExecuteMsgs(ctx, k.router, accountInfo.GroupAccount, proposal.GetMsgs())
		if err != nil {
			proposal.ExecutorResult = group.ProposalExecutorResultFailure
			proposalType := reflect.TypeOf(proposal).String()
			logger.Info("proposal execution failed", "cause", err, "type", proposalType, "proposalID", id)
		} else {
			proposal.ExecutorResult = group.ProposalExecutorResultSuccess
			flush()
		}
	}
	return storeUpdates()
}

func (k Keeper) GetProposal(ctx sdk.Context, id group.ProposalID) (group.Proposal, error) {
	var p group.Proposal
	if _, err := k.proposalTable.GetOne(ctx, id.Uint64(), &p); err != nil {
		return group.Proposal{}, sdkerrors.Wrap(err, "load proposal")
	}
	return p, nil
}

func (k Keeper) CreateProposal(ctx sdk.Context, accountAddress sdk.AccAddress, comment string, proposers []sdk.AccAddress, msgs []sdk.Msg) (group.ProposalID, error) {
	if err := assertCommentSize(comment, k.MaxCommentSize(ctx), "comment"); err != nil {
		return 0, err
	}

	account, err := k.GetGroupAccount(ctx, accountAddress.Bytes())

	if err != nil {
		return 0, sdkerrors.Wrap(err, "load group account")
	}

	g, err := k.GetGroup(ctx, account.GroupId)
	if err != nil {
		return 0, sdkerrors.Wrap(err, "get group by account")
	}

	// only members can propose
	for i := range proposers {
		if !k.groupMemberTable.Has(ctx, group.GroupMember{GroupId: g.GroupId, Member: proposers[i]}.NaturalKey()) {
			return 0, sdkerrors.Wrapf(group.ErrUnauthorized, "not in group: %s", proposers[i])
		}
	}

	if err := ensureMsgAuthZ(msgs, account.GroupAccount); err != nil {
		return 0, err
	}

	blockTime, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return 0, sdkerrors.Wrap(err, "block time conversion")
	}

	policy := account.GetDecisionPolicy()

	if policy == nil {
		return 0, sdkerrors.Wrap(group.ErrEmpty, "nil policy")
	}

	// prevent proposal that can not succeed
	err = policy.Validate(g)
	if err != nil {
		return 0, err
	}

	timeout := policy.GetTimeout()
	window, err := gogotypes.DurationFromProto(&timeout)
	if err != nil {
		return 0, sdkerrors.Wrap(err, "maxVotingWindow time conversion")
	}
	endTime, err := gogotypes.TimestampProto(ctx.BlockTime().Add(window))
	if err != nil {
		return 0, sdkerrors.Wrap(err, "end time conversion")
	}

	m := &group.Proposal{
		GroupAccount:        accountAddress,
		Comment:             comment,
		Proposers:           proposers,
		SubmittedAt:         *blockTime,
		GroupVersion:        g.Version,
		GroupAccountVersion: account.Version,
		Result:              group.ProposalResultUnfinalized,
		Status:              group.ProposalStatusSubmitted,
		ExecutorResult:      group.ProposalExecutorResultNotRun,
		Timeout:             *endTime,
		VoteState: group.Tally{
			YesCount:     "0",
			NoCount:      "0",
			AbstainCount: "0",
			VetoCount:    "0",
		},
	}
	if err := m.SetMsgs(msgs); err != nil {
		return 0, sdkerrors.Wrap(err, "create proposal")
	}

	id, err := k.proposalTable.Create(ctx, m)
	if err != nil {
		return 0, sdkerrors.Wrap(err, "create proposal")
	}
	return group.ProposalID(id), nil
}

func (k Keeper) GetVote(ctx sdk.Context, id group.ProposalID, voter sdk.AccAddress) (group.Vote, error) {
	var v group.Vote
	return v, k.voteTable.GetOne(ctx, group.Vote{ProposalId: id, Voter: voter}.NaturalKey(), &v)
}

func assertCommentSize(comment string, maxCommentSize int, description string) error {
	if len(comment) > maxCommentSize {
		return sdkerrors.Wrap(group.ErrMaxLimit, description)
	}
	return nil
}
