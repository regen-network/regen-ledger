package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

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
	metadata := req.Metadata
	members := group.Members(req.Members)
	admin := req.Admin

	if err := members.ValidateBasic(); err != nil {
		return nil, err
	}

	maxMetadataLength := s.maxMetadataLength(ctx)
	if err := assertMetadataLength(metadata, maxMetadataLength, "group metadata"); err != nil {
		return nil, err
	}

	totalWeight := apd.New(0, 0)
	for i := range members {
		m := members[i]
		if err := assertMetadataLength(m.Metadata, maxMetadataLength, "member metadata"); err != nil {
			return nil, err
		}

		// Members of a group must have a positive weight.
		weight, err := math.ParsePositiveDecimal(m.Weight)
		if err != nil {
			return nil, err
		}

		// Adding up members weights to compute group total weight.
		err = math.Add(totalWeight, totalWeight, weight)
		if err != nil {
			return nil, err
		}
	}

	// Create a new group in the groupTable.
	groupID := group.ID(s.groupSeq.NextVal(ctx))
	err := s.groupTable.Create(ctx, groupID.Bytes(), &group.GroupInfo{
		GroupId:     groupID,
		Admin:       admin,
		Metadata:    metadata,
		Version:     1,
		TotalWeight: math.DecimalString(totalWeight),
	})
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not create group")
	}

	// Create new group members in the groupMemberTable.
	for i := range members {
		m := members[i]
		err := s.groupMemberTable.Create(ctx, &group.GroupMember{
			GroupId: groupID,
			Member: &group.Member{
				Address:  m.Address,
				Weight:   m.Weight,
				Metadata: m.Metadata,
			},
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
		totalWeight, err := math.ParseNonNegativeDecimal(g.TotalWeight)
		if err != nil {
			return err
		}
		for i := range req.MemberUpdates {
			groupMember := group.GroupMember{GroupId: req.GroupId,
				Member: &group.Member{
					Address:  req.MemberUpdates[i].Address,
					Weight:   req.MemberUpdates[i].Weight,
					Metadata: req.MemberUpdates[i].Metadata,
				},
			}

			// Checking if the group member is already part of the group.
			var found bool
			var prevGroupMember group.GroupMember
			switch err := s.groupMemberTable.GetOne(ctx, groupMember.NaturalKey(), &prevGroupMember); {
			case err == nil:
				found = true
			case orm.ErrNotFound.Is(err):
				found = false
			default:
				return sdkerrors.Wrap(err, "get group member")
			}

			newMemberWeight, err := math.ParseNonNegativeDecimal(groupMember.Member.Weight)
			if err != nil {
				return err
			}

			// Handle delete for members with zero weight.
			if newMemberWeight.IsZero() {
				// We can't delete a group member that doesn't already exist.
				if !found {
					return sdkerrors.Wrap(orm.ErrNotFound, "unknown member")
				}

				previousMemberWeight, err := math.ParseNonNegativeDecimal(prevGroupMember.Member.Weight)
				if err != nil {
					return err
				}

				// Subtract the weight of the group member to delete from the group total weight.
				err = math.SafeSub(totalWeight, totalWeight, previousMemberWeight)
				if err != nil {
					return err
				}

				// Delete group member in the groupMemberTable.
				if err := s.groupMemberTable.Delete(ctx, &groupMember); err != nil {
					return sdkerrors.Wrap(err, "delete member")
				}
				continue
			}
			// If group member already exists, handle update
			if found {
				previousMemberWeight, err := math.ParseNonNegativeDecimal(prevGroupMember.Member.Weight)
				if err != nil {
					return err
				}
				// Subtract previous weight from the group total weight.
				err = math.SafeSub(totalWeight, totalWeight, previousMemberWeight)
				if err != nil {
					return err
				}
				// Save updated group member in the groupMemberTable.
				if err := s.groupMemberTable.Save(ctx, &groupMember); err != nil {
					return sdkerrors.Wrap(err, "add member")
				}
				// else handle create.
			} else if err := s.groupMemberTable.Create(ctx, &groupMember); err != nil {
				return sdkerrors.Wrap(err, "add member")
			}
			// In both cases (handle + update), we need to add the new member's weight to the group total weight.
			err = math.Add(totalWeight, totalWeight, newMemberWeight)
			if err != nil {
				return err
			}
		}
		// Update group in the groupTable.
		g.TotalWeight = math.DecimalString(totalWeight)
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

func (s serverImpl) UpdateGroupMetadata(ctx types.Context, req *group.MsgUpdateGroupMetadataRequest) (*group.MsgUpdateGroupMetadataResponse, error) {
	action := func(g *group.GroupInfo) error {
		g.Metadata = req.Metadata
		g.Version++
		return s.groupTable.Save(ctx, g.GroupId.Bytes(), g)
	}

	err := s.doUpdateGroup(ctx, req, action, "metadata updated")
	if err != nil {
		return nil, err
	}

	return &group.MsgUpdateGroupMetadataResponse{}, nil
}

func (s serverImpl) CreateGroupAccount(ctx types.Context, req *group.MsgCreateGroupAccountRequest) (*group.MsgCreateGroupAccountResponse, error) {
	admin, err := sdk.AccAddressFromBech32(req.GetAdmin())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "request admin")
	}
	policy := req.GetDecisionPolicy()
	groupID := req.GetGroupID()
	metadata := req.GetMetadata()

	if err := assertMetadataLength(metadata, s.maxMetadataLength(ctx), "group account metadata"); err != nil {
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
	// Only current group admin is authorized to create a group account for this group.
	if !groupAdmin.Equals(admin) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not group admin")
	}

	// Generate group account address.
	var accountAddr sdk.AccAddress
	// loop here in the rare case of a collision
	for {
		nextAccVal := s.groupAccountSeq.NextVal(ctx)
		buf := bytes.NewBuffer(nil)
		err = binary.Write(buf, binary.LittleEndian, nextAccVal)
		if err != nil {
			return nil, err
		}

		accountID := s.key.Derive(buf.Bytes())
		accountAddr = accountID.Address()

		if s.accKeeper.GetAccount(ctx.Context, accountAddr) != nil {
			// handle a rare collision
			continue
		}

		acc := s.accKeeper.NewAccount(ctx.Context, &authtypes.ModuleAccount{
			BaseAccount: &authtypes.BaseAccount{
				Address: accountAddr.String(),
			},
			Name: accountAddr.String(),
		})
		s.accKeeper.SetAccount(ctx.Context, acc)

		break
	}

	groupAccount, err := group.NewGroupAccountInfo(
		accountAddr,
		groupID,
		admin,
		metadata,
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

func (s serverImpl) UpdateGroupAccountMetadata(ctx types.Context, req *group.MsgUpdateGroupAccountMetadataRequest) (*group.MsgUpdateGroupAccountMetadataResponse, error) {
	// TODO #224
	return &group.MsgUpdateGroupAccountMetadataResponse{}, nil
}

func (s serverImpl) CreateProposal(ctx types.Context, req *group.MsgCreateProposalRequest) (*group.MsgCreateProposalResponse, error) {
	accountAddress, err := sdk.AccAddressFromBech32(req.GroupAccount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "request group account")
	}
	metadata := req.Metadata
	proposers := req.Proposers
	msgs := req.GetMsgs()

	if err := assertMetadataLength(metadata, s.maxMetadataLength(ctx), "metadata"); err != nil {
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

	// Only members of the group can submit a new proposal.
	for i := range proposers {
		if !s.groupMemberTable.Has(ctx, group.GroupMember{GroupId: g.GroupId, Member: &group.Member{Address: proposers[i]}}.NaturalKey()) {
			return nil, sdkerrors.Wrapf(group.ErrUnauthorized, "not in group: %s", proposers[i])
		}
	}

	// Check that if the messages require signers, they are all equal to the given group account.
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

	// Prevent proposal that can not succeed.
	err = policy.Validate(g)
	if err != nil {
		return nil, err
	}

	// Define proposal timout.
	// The voting window begins as soon as the proposal is submitted.
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
		Metadata:            metadata,
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
	choice := req.Choice
	metadata := req.Metadata

	if err := assertMetadataLength(metadata, s.maxMetadataLength(ctx), "metadata"); err != nil {
		return nil, err
	}

	blockTime, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		return nil, err
	}
	proposal, err := s.getProposal(ctx, id)
	if err != nil {
		return nil, err
	}
	// Ensure that we can still accept votes for this proposal.
	if proposal.Status != group.ProposalStatusSubmitted {
		return nil, sdkerrors.Wrap(group.ErrInvalid, "proposal not open for voting")
	}
	votingPeriodEnd, err := gogotypes.TimestampFromProto(&proposal.Timeout)
	if err != nil {
		return nil, err
	}
	if votingPeriodEnd.Before(ctx.BlockTime()) || votingPeriodEnd.Equal(ctx.BlockTime()) {
		return nil, sdkerrors.Wrap(group.ErrExpired, "voting period has ended already")
	}

	var accountInfo group.GroupAccountInfo

	// Ensure that group account hasn't been modified since the proposal submission.
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

	// Ensure that group hasn't been modified since the proposal submission.
	electorate, err := s.getGroupInfo(ctx, accountInfo.GroupId)
	if err != nil {
		return nil, err
	}
	if electorate.Version != proposal.GroupVersion {
		return nil, sdkerrors.Wrap(group.ErrModified, "group was modified")
	}

	// Count and store votes.
	voterAddr := req.Voter
	voter := group.GroupMember{GroupId: electorate.GroupId, Member: &group.Member{Address: voterAddr}}
	if err := s.groupMemberTable.GetOne(ctx, voter.NaturalKey(), &voter); err != nil {
		return nil, sdkerrors.Wrapf(err, "address: %s", voterAddr)
	}
	newVote := group.Vote{
		ProposalId:  id,
		Voter:       voterAddr,
		Choice:      choice,
		Metadata:    metadata,
		SubmittedAt: *blockTime,
	}
	if err := proposal.VoteState.Add(newVote, voter.Member.Weight); err != nil {
		return nil, sdkerrors.Wrap(err, "add new vote")
	}

	// The ORM will return an error if the vote already exists,
	// making sure than a voter hasn't already voted.
	if err := s.voteTable.Create(ctx, &newVote); err != nil {
		return nil, sdkerrors.Wrap(err, "store vote")
	}

	// Run tally with new votes to close early.
	if err := doTally(ctx, &proposal, electorate, accountInfo); err != nil {
		return nil, err
	}

	if err = s.proposalTable.Save(ctx, id.Uint64(), &proposal); err != nil {
		return nil, err
	}

	// TODO: add event #215

	return &group.MsgVoteResponse{}, nil
}

// doTally updates the proposal status and tally if necessary based on the group account's decision policy.
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

// Exec executes the messages from a proposal.
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
		// Ensure that group account hasn't been modified before tally.
		if proposal.GroupAccountVersion != accountInfo.Version {
			proposal.Result = group.ProposalResultUnfinalized
			proposal.Status = group.ProposalStatusAborted
			return storeUpdates()
		}

		electorate, err := s.getGroupInfo(ctx, accountInfo.GroupId)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "load group")
		}

		// Ensure that group hasn't been modified before tally.
		if electorate.Version != proposal.GroupVersion {
			proposal.Result = group.ProposalResultUnfinalized
			proposal.Status = group.ProposalStatusAborted
			return storeUpdates()
		}
		if err := doTally(ctx, &proposal, electorate, accountInfo); err != nil {
			return nil, err
		}
	}

	// Execute proposal payload.
	if proposal.Status == group.ProposalStatusClosed && proposal.Result == group.ProposalResultAccepted && proposal.ExecutorResult != group.ProposalExecutorResultSuccess {
		logger := ctx.Logger().With("module", fmt.Sprintf("x/%s", group.ModuleName))
		// Cashing context so that we don't update the store in case of failure.
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

	// Update proposal in proposalTable
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

// doUpdateGroup first makes sure that the group admin initiated the group update,
// before performing the group update and emitting an event.
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

// doAuthenticated makes sure that the group admin initiated the request,
// and perform the provided action on the group.
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

// maxMetadataLength returns the maximum length of a metadata field.
func (s serverImpl) maxMetadataLength(ctx types.Context) int {
	return group.MaxMetadataLength
}

// assertMetadataLength returns an error if given metadata length
// is greater than a fixed maxMetadataLength.
func assertMetadataLength(metadata []byte, maxMetadataLength int, description string) error {
	if len(metadata) > maxMetadataLength {
		return sdkerrors.Wrap(group.ErrMaxLimit, description)
	}
	return nil
}
