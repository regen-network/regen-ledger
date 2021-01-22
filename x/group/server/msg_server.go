package server

import (
	"fmt"
	"reflect"

	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/regen-network/regen-ledger/math"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/util"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) CreateGroup(ctx types.Context, req *group.MsgCreateGroupRequest) (*group.MsgCreateGroupResponse, error) {
	comment := req.Comment
	members := group.Members(req.Members)
	admin := req.Admin

	if err := members.ValidateBasic(); err != nil {
		return nil, err
	}

	maxCommentSize := s.maxCommentSize(ctx)
	if err := assertCommentSize(comment, maxCommentSize, "group comment"); err != nil {
		return nil, err
	}

	totalWeight := apd.New(0, 0)
	for i := range members {
		m := members[i]
		if err := assertCommentSize(m.Comment, maxCommentSize, "member comment"); err != nil {
			return nil, err
		}

		power, err := math.ParseNonNegativeDecimal(m.Power)
		if err != nil {
			return nil, err
		}

		if !power.IsZero() {
			err = math.Add(totalWeight, totalWeight, power)
			if err != nil {
				return nil, err
			}
		}
	}

	groupID := group.ID(s.groupSeq.NextVal(ctx))
	err := s.groupTable.Create(ctx, groupID.Bytes(), &group.GroupInfo{
		GroupId:     groupID,
		Admin:       admin,
		Comment:     comment,
		Version:     1,
		TotalWeight: math.DecimalString(totalWeight),
	})
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not create group")
	}

	for i := range members {
		m := members[i]
		err := s.groupMemberTable.Create(ctx, &group.GroupMember{
			GroupId: groupID,
			Member:  m.Address,
			Weight:  m.Power,
			Comment: m.Comment,
		})
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "could not store member %d", i)
		}
	}

	groupIDStr := util.Uint64ToBase58Check(groupID.Uint64())
	err = ctx.EventManager().EmitTypedEvent(&group.EventCreateGroup{GroupId: groupIDStr})
	if err != nil {
		return nil, err
	}

	return &group.MsgCreateGroupResponse{GroupId: groupID}, nil
}

func (s serverImpl) UpdateGroupMembers(ctx types.Context, req *group.MsgUpdateGroupMembersRequest) (*group.MsgUpdateGroupMembersResponse, error) {
	action := func(g *group.GroupInfo) error {
		for i := range req.MemberUpdates {
			member := group.GroupMember{GroupId: req.GroupId,
				Member:  req.MemberUpdates[i].Address,
				Weight:  req.MemberUpdates[i].Power,
				Comment: req.MemberUpdates[i].Comment,
			}
			var found bool
			var previousMember group.GroupMember
			switch err := s.groupMemberTable.GetOne(ctx, member.NaturalKey(), &previousMember); {
			case err == nil:
				found = true
			case orm.ErrNotFound.Is(err):
				found = false
			default:
				return sdkerrors.Wrap(err, "get group member")
			}

			totalWeight, err := math.ParseNonNegativeDecimal(g.TotalWeight)
			if err != nil {
				return err
			}
			newMemberWeight, err := math.ParseNonNegativeDecimal(member.Weight)
			if err != nil {
				return err
			}

			// handle delete
			if newMemberWeight.Cmp(apd.New(0, 0)) == 0 {
				if !found {
					return sdkerrors.Wrap(orm.ErrNotFound, "unknown member")
				}

				previousMemberWeight, err := math.ParseNonNegativeDecimal(previousMember.Weight)
				if err != nil {
					return err
				}

				err = math.SafeSub(totalWeight, totalWeight, previousMemberWeight)
				if err != nil {
					return err
				}

				g.TotalWeight = math.DecimalString(totalWeight)
				if err := s.groupMemberTable.Delete(ctx, &member); err != nil {
					return sdkerrors.Wrap(err, "delete member")
				}
				continue
			}
			// handle add + update
			if found {
				previousMemberWeight, err := math.ParseNonNegativeDecimal(previousMember.Weight)
				if err != nil {
					return err
				}
				err = math.SafeSub(totalWeight, totalWeight, previousMemberWeight)
				if err != nil {
					return err
				}
				if err := s.groupMemberTable.Save(ctx, &member); err != nil {
					return sdkerrors.Wrap(err, "add member")
				}
			} else if err := s.groupMemberTable.Create(ctx, &member); err != nil {
				return sdkerrors.Wrap(err, "add member")
			}
			err = math.Add(totalWeight, totalWeight, newMemberWeight)
			if err != nil {
				return err
			}
			g.TotalWeight = math.DecimalString(totalWeight)
		}
		g.Version++
		return s.groupTable.Save(ctx, g.GroupId.Bytes(), g)
	}

	err := s.doUpdateGroup(ctx, req, action, "members updated")
	if err != nil {
		return nil, err
	}

	return &group.MsgUpdateGroupMembersResponse{}, nil
}

func (s serverImpl) UpdateGroupAdmin(ctx types.Context, req *group.MsgUpdateGroupAdminRequest) (*group.MsgUpdateGroupAdminResponse, error) {
	action := func(g *group.GroupInfo) error {
		g.Admin = req.NewAdmin
		g.Version++
		return s.groupTable.Save(ctx, g.GroupId.Bytes(), g)
	}

	err := s.doUpdateGroup(ctx, req, action, "admin updated")
	if err != nil {
		return nil, err
	}

	return &group.MsgUpdateGroupAdminResponse{}, nil
}

func (s serverImpl) UpdateGroupComment(ctx types.Context, req *group.MsgUpdateGroupCommentRequest) (*group.MsgUpdateGroupCommentResponse, error) {
	action := func(g *group.GroupInfo) error {
		g.Comment = req.Comment
		g.Version++
		return s.groupTable.Save(ctx, g.GroupId.Bytes(), g)
	}

	err := s.doUpdateGroup(ctx, req, action, "comment updated")
	if err != nil {
		return nil, err
	}

	return &group.MsgUpdateGroupCommentResponse{}, nil
}

func (s serverImpl) CreateGroupAccount(ctx types.Context, req *group.MsgCreateGroupAccountRequest) (*group.MsgCreateGroupAccountResponse, error) {
	admin, err := sdk.AccAddressFromBech32(req.GetAdmin())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "request admin")
	}
	policy := req.GetDecisionPolicy()
	groupID := req.GetGroupID()
	comment := req.GetComment()

	if err := assertCommentSize(comment, s.maxCommentSize(ctx), "group account comment"); err != nil {
		return nil, err
	}

	g, err := s.getGroupInfo(ctx, groupID)
	if err != nil {
		return nil, err
	}
	groupAdmin, err := sdk.AccAddressFromBech32(g.Admin)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "group admin")
	}
	if !groupAdmin.Equals(admin) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not group admin")
	}
	accountAddr := group.AccountCondition(s.groupAccountSeq.NextVal(ctx)).Address()
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

	if err := s.groupAccountTable.Create(ctx, &groupAccount); err != nil {
		return nil, sdkerrors.Wrap(err, "could not create group account")
	}

	err = ctx.EventManager().EmitTypedEvent(&group.EventCreateGroupAccount{GroupAccount: accountAddr.String()})
	if err != nil {
		return nil, err
	}

	return &group.MsgCreateGroupAccountResponse{GroupAccount: accountAddr.String()}, nil
}

func (s serverImpl) UpdateGroupAccountAdmin(ctx types.Context, req *group.MsgUpdateGroupAccountAdminRequest) (*group.MsgUpdateGroupAccountAdminResponse, error) {
	// TODO #224
	return &group.MsgUpdateGroupAccountAdminResponse{}, nil
}

func (s serverImpl) UpdateGroupAccountDecisionPolicy(ctx types.Context, req *group.MsgUpdateGroupAccountDecisionPolicyRequest) (*group.MsgUpdateGroupAccountDecisionPolicyResponse, error) {
	// TODO #224
	return &group.MsgUpdateGroupAccountDecisionPolicyResponse{}, nil
}

func (s serverImpl) UpdateGroupAccountComment(ctx types.Context, req *group.MsgUpdateGroupAccountCommentRequest) (*group.MsgUpdateGroupAccountCommentResponse, error) {
	// TODO #224
	return &group.MsgUpdateGroupAccountCommentResponse{}, nil
}

func (s serverImpl) CreateProposal(ctx types.Context, req *group.MsgCreateProposalRequest) (*group.MsgCreateProposalResponse, error) {
	accountAddress, err := sdk.AccAddressFromBech32(req.GroupAccount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "request group account")
	}
	comment := req.Comment
	proposers := req.Proposers
	msgs := req.GetMsgs()

	if err := assertCommentSize(comment, s.maxCommentSize(ctx), "comment"); err != nil {
		return nil, err
	}

	account, err := s.getGroupAccountInfo(ctx, accountAddress.Bytes())

	if err != nil {
		return nil, sdkerrors.Wrap(err, "load group account")
	}

	g, err := s.getGroupInfo(ctx, account.GroupId)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "get group by account")
	}

	// only members can propose
	for i := range proposers {
		if !s.groupMemberTable.Has(ctx, group.GroupMember{GroupId: g.GroupId, Member: proposers[i]}.NaturalKey()) {
			return nil, sdkerrors.Wrapf(group.ErrUnauthorized, "not in group: %s", proposers[i])
		}
	}

	if err := ensureMsgAuthZ(msgs, accountAddress); err != nil {
		return nil, err
	}

	blockTime, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "block time conversion")
	}

	policy := account.GetDecisionPolicy()

	if policy == nil {
		return nil, sdkerrors.Wrap(group.ErrEmpty, "nil policy")
	}

	// prevent proposal that can not succeed
	err = policy.Validate(g)
	if err != nil {
		return nil, err
	}

	timeout := policy.GetTimeout()
	window, err := gogotypes.DurationFromProto(&timeout)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "maxVotingWindow time conversion")
	}
	endTime, err := gogotypes.TimestampProto(ctx.BlockTime().Add(window))
	if err != nil {
		return nil, sdkerrors.Wrap(err, "end time conversion")
	}

	m := &group.Proposal{
		GroupAccount:        req.GroupAccount,
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
		return nil, sdkerrors.Wrap(err, "create proposal")
	}

	id, err := s.proposalTable.Create(ctx, m)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "create proposal")
	}

	// TODO: add event #215

	return &group.MsgCreateProposalResponse{ProposalId: group.ProposalID(id)}, nil
}

func (s serverImpl) Vote(ctx types.Context, req *group.MsgVoteRequest) (*group.MsgVoteResponse, error) {
	id := req.ProposalId
	voters := req.Voters
	choice := req.Choice
	comment := req.Comment

	if err := assertCommentSize(comment, s.maxCommentSize(ctx), "comment"); err != nil {
		return nil, err
	}
	if len(voters) == 0 {
		return nil, sdkerrors.Wrap(group.ErrEmpty, "voters")
	}

	blockTime, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return nil, err
	}
	proposal, err := s.getProposal(ctx, id)
	if err != nil {
		return nil, err
	}
	if proposal.Status != group.ProposalStatusSubmitted {
		return nil, sdkerrors.Wrap(group.ErrInvalid, "proposal not open")
	}
	votingPeriodEnd, err := gogotypes.TimestampFromProto(&proposal.Timeout)
	if err != nil {
		return nil, err
	}
	if votingPeriodEnd.Before(ctx.BlockTime()) || votingPeriodEnd.Equal(ctx.BlockTime()) {
		return nil, sdkerrors.Wrap(group.ErrExpired, "voting period has ended already")
	}

	var accountInfo group.GroupAccountInfo

	address, err := sdk.AccAddressFromBech32(proposal.GroupAccount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "group account")
	}
	if err := s.groupAccountTable.GetOne(ctx, address.Bytes(), &accountInfo); err != nil {
		return nil, sdkerrors.Wrap(err, "load group account")
	}
	if proposal.GroupAccountVersion != accountInfo.Version {
		return nil, sdkerrors.Wrap(group.ErrModified, "group account was modified")
	}

	electorate, err := s.getGroupInfo(ctx, accountInfo.GroupId)
	if err != nil {
		return nil, err
	}
	if electorate.Version != proposal.GroupVersion {
		return nil, sdkerrors.Wrap(group.ErrModified, "group was modified")
	}

	// count and store votes
	for _, voterAddr := range voters {
		voter := group.GroupMember{GroupId: electorate.GroupId, Member: voterAddr}
		if err := s.groupMemberTable.GetOne(ctx, voter.NaturalKey(), &voter); err != nil {
			return nil, sdkerrors.Wrapf(err, "address: %s", voterAddr)
		}
		newVote := group.Vote{
			ProposalId:  id,
			Voter:       voterAddr,
			Choice:      choice,
			Comment:     comment,
			SubmittedAt: *blockTime,
		}
		if err := proposal.VoteState.Add(newVote, voter.Weight); err != nil {
			return nil, sdkerrors.Wrap(err, "add new vote")
		}

		// The ORM will return an error if the vote already exists,
		// making sure than a voter hasn't already voted.
		if err := s.voteTable.Create(ctx, &newVote); err != nil {
			return nil, sdkerrors.Wrap(err, "store vote")
		}
	}

	// run tally with new votes to close early
	if err := doTally(ctx, &proposal, electorate, accountInfo); err != nil {
		return nil, err
	}

	if err = s.proposalTable.Save(ctx, id.Uint64(), &proposal); err != nil {
		return nil, err
	}

	// TODO: add event #215

	return &group.MsgVoteResponse{}, nil
}

func doTally(ctx types.Context, p *group.Proposal, electorate group.GroupInfo, accountInfo group.GroupAccountInfo) error {
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

func (s serverImpl) Exec(ctx types.Context, req *group.MsgExecRequest) (*group.MsgExecResponse, error) {
	id := req.ProposalId

	proposal, err := s.getProposal(ctx, id)
	if err != nil {
		return nil, err
	}

	if proposal.Status != group.ProposalStatusSubmitted && proposal.Status != group.ProposalStatusClosed {
		return nil, sdkerrors.Wrapf(group.ErrInvalid, "not possible with proposal status %s", proposal.Status.String())
	}

	var accountInfo group.GroupAccountInfo
	address, err := sdk.AccAddressFromBech32(proposal.GroupAccount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "group account")
	}
	if err := s.groupAccountTable.GetOne(ctx, address.Bytes(), &accountInfo); err != nil {
		return nil, sdkerrors.Wrap(err, "load group account")
	}

	storeUpdates := func() (*group.MsgExecResponse, error) {
		if err := s.proposalTable.Save(ctx, id.Uint64(), &proposal); err != nil {
			return nil, err
		}
		return &group.MsgExecResponse{}, nil
	}

	if proposal.Status == group.ProposalStatusSubmitted {
		if proposal.GroupAccountVersion != accountInfo.Version {
			proposal.Result = group.ProposalResultUnfinalized
			proposal.Status = group.ProposalStatusAborted
			return storeUpdates()
		}

		electorate, err := s.getGroupInfo(ctx, accountInfo.GroupId)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "load group")
		}

		if electorate.Version != proposal.GroupVersion {
			proposal.Result = group.ProposalResultUnfinalized
			proposal.Status = group.ProposalStatusAborted
			return storeUpdates()
		}
		if err := doTally(ctx, &proposal, electorate, accountInfo); err != nil {
			return nil, err
		}
	}

	// execute proposal payload
	if proposal.Status == group.ProposalStatusClosed && proposal.Result == group.ProposalResultAccepted && proposal.ExecutorResult != group.ProposalExecutorResultSuccess {
		logger := ctx.Logger().With("module", fmt.Sprintf("x/%s", group.ModuleName))
		ctx, flush := ctx.CacheContext()
		address, err := sdk.AccAddressFromBech32(accountInfo.GroupAccount)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "group account")
		}
		_, err = DoExecuteMsgs(ctx, s.router, address, proposal.GetMsgs())
		if err != nil {
			proposal.ExecutorResult = group.ProposalExecutorResultFailure
			proposalType := reflect.TypeOf(proposal).String()
			logger.Info("proposal execution failed", "cause", err, "type", proposalType, "proposalID", id)
		} else {
			proposal.ExecutorResult = group.ProposalExecutorResultSuccess
			flush()
		}
	}

	res, err := storeUpdates()
	if err != nil {
		return nil, err
	}
	// TODO: add event #215

	return res, nil
}

type authNGroupReq interface {
	GetGroupID() group.ID
	GetAdmin() string
}

type actionFn func(m *group.GroupInfo) error

func (s serverImpl) doUpdateGroup(ctx types.Context, req authNGroupReq, action actionFn, note string) error {
	err := s.doAuthenticated(ctx, req, action, note)
	if err != nil {
		return err
	}

	groupIDStr := util.Uint64ToBase58Check(req.GetGroupID().Uint64())
	err = ctx.EventManager().EmitTypedEvent(&group.EventUpdateGroup{GroupId: groupIDStr})
	if err != nil {
		return err
	}

	return nil
}

func (s serverImpl) doAuthenticated(ctx types.Context, req authNGroupReq, action actionFn, note string) error {
	group, err := s.getGroupInfo(ctx, req.GetGroupID())
	if err != nil {
		return err
	}
	admin, err := sdk.AccAddressFromBech32(group.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "group admin")
	}
	reqAdmin, err := sdk.AccAddressFromBech32(req.GetAdmin())
	if err != nil {
		return sdkerrors.Wrap(err, "request admin")
	}
	if !admin.Equals(reqAdmin) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not group admin")
	}
	if err := action(&group); err != nil {
		return sdkerrors.Wrap(err, note)
	}
	return nil
}

// maxCommentSize returns the maximum length of a comment
func (s serverImpl) maxCommentSize(ctx types.Context) int {
	var result uint32
	s.paramSpace.Get(ctx.Context, group.ParamMaxCommentLength, &result)
	return int(result)
}

func assertCommentSize(comment string, maxCommentSize int, description string) error {
	if len(comment) > maxCommentSize {
		return sdkerrors.Wrap(group.ErrMaxLimit, description)
	}
	return nil
}
