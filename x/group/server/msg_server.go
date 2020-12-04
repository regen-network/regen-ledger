package server

import (
	"context"

	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/math"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/util"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) CreateGroup(goCtx context.Context, req *group.MsgCreateGroupRequest) (*group.MsgCreateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	groupID, err := s.Keeper.CreateGroup(ctx, req.Admin, req.Members, req.Comment)
	if err != nil {
		return nil, err
	}

	groupIDStr := util.Uint64ToBase58Check(groupID.Uint64())
	err = ctx.EventManager().EmitTypedEvent(&group.EventCreateGroup{GroupId: groupIDStr, Admin: req.Admin.String()})
	if err != nil {
		return nil, err
	}

	return &group.MsgCreateGroupResponse{GroupId: groupID}, nil
}

func (s serverImpl) UpdateGroupMembers(goCtx context.Context, req *group.MsgUpdateGroupMembersRequest) (*group.MsgUpdateGroupMembersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	action := func(m *group.GroupInfo) error {
		for i := range req.MemberUpdates {
			member := group.GroupMember{GroupId: req.GroupId,
				Member:  req.MemberUpdates[i].Address,
				Weight:  req.MemberUpdates[i].Power,
				Comment: req.MemberUpdates[i].Comment,
			}
			var found bool
			var previousMemberStatus group.GroupMember
			switch err := s.groupMemberTable.GetOne(ctx, member.NaturalKey(), &previousMemberStatus); {
			case err == nil:
				found = true
			case orm.ErrNotFound.Is(err):
				found = false
			default:
				return sdkerrors.Wrap(err, "get group member")
			}

			totalWeight, err := math.ParseNonNegativeDecimal(m.TotalWeight)
			if err != nil {
				return err
			}
			weight, err := math.ParseNonNegativeDecimal(member.Weight)
			if err != nil {
				return err
			}

			// handle delete
			if weight.Cmp(apd.New(0, 0)) == 0 {
				if !found {
					return sdkerrors.Wrap(orm.ErrNotFound, "unknown member")
				}

				previousWeight, err := math.ParseNonNegativeDecimal(previousMemberStatus.Weight)
				if err != nil {
					return err
				}

				err = math.SafeSub(totalWeight, totalWeight, previousWeight)
				if err != nil {
					return err
				}

				m.TotalWeight = math.DecimalString(totalWeight)
				if err := s.groupMemberTable.Delete(ctx, &member); err != nil {
					return sdkerrors.Wrap(err, "delete member")
				}
				continue
			}
			// handle add + update
			if found {
				previousWeight, err := math.ParseNonNegativeDecimal(previousMemberStatus.Weight)
				if err != nil {
					return err
				}
				err = math.SafeSub(totalWeight, totalWeight, previousWeight)
				if err != nil {
					return err
				}
				if err := s.groupMemberTable.Save(ctx, &member); err != nil {
					return sdkerrors.Wrap(err, "add member")
				}
			} else if err := s.groupMemberTable.Create(ctx, &member); err != nil {
				return sdkerrors.Wrap(err, "add member")
			}
			err = math.Add(totalWeight, totalWeight, weight)
			if err != nil {
				return err
			}
			m.TotalWeight = math.DecimalString(totalWeight)
		}
		return s.UpdateGroup(ctx, m)
	}

	err := s.doUpdateGroup(ctx, req, action, "members updated")
	if err != nil {
		return nil, err
	}

	return &group.MsgUpdateGroupMembersResponse{}, nil
}

func (s serverImpl) UpdateGroupAdmin(goCtx context.Context, req *group.MsgUpdateGroupAdminRequest) (*group.MsgUpdateGroupAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	action := func(m *group.GroupInfo) error {
		m.Admin = req.NewAdmin
		return s.UpdateGroup(ctx, m)
	}

	err := s.doUpdateGroup(ctx, req, action, "admin updated")
	if err != nil {
		return nil, err
	}

	return &group.MsgUpdateGroupAdminResponse{}, nil
}

func (s serverImpl) UpdateGroupComment(goCtx context.Context, req *group.MsgUpdateGroupCommentRequest) (*group.MsgUpdateGroupCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	action := func(m *group.GroupInfo) error {
		m.Comment = req.Comment
		return s.UpdateGroup(ctx, m)
	}

	err := s.doUpdateGroup(ctx, req, action, "comment updated")
	if err != nil {
		return nil, err
	}

	return &group.MsgUpdateGroupCommentResponse{}, nil
}

func (s serverImpl) CreateGroupAccount(goCtx context.Context, req *group.MsgCreateGroupAccountRequest) (*group.MsgCreateGroupAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	decisionPolicy := req.GetDecisionPolicy()
	acc, err := s.Keeper.CreateGroupAccount(ctx, req.GetAdmin(), req.GetGroupID(), decisionPolicy, req.GetComment())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "create group account")
	}

	err = ctx.EventManager().EmitTypedEvent(&group.EventCreateGroupAccount{GroupAccount: acc.String(), Admin: req.Admin.String()})
	if err != nil {
		return nil, err
	}

	return &group.MsgCreateGroupAccountResponse{GroupAccount: acc}, nil
}

func (s serverImpl) UpdateGroupAccountAdmin(goCtx context.Context, req *group.MsgUpdateGroupAccountAdminRequest) (*group.MsgUpdateGroupAccountAdminResponse, error) {
	// TODO
	return &group.MsgUpdateGroupAccountAdminResponse{}, nil
}

func (s serverImpl) UpdateGroupAccountDecisionPolicy(goCtx context.Context, req *group.MsgUpdateGroupAccountDecisionPolicyRequest) (*group.MsgUpdateGroupAccountDecisionPolicyResponse, error) {
	// TODO
	return &group.MsgUpdateGroupAccountDecisionPolicyResponse{}, nil
}

func (s serverImpl) UpdateGroupAccountComment(goCtx context.Context, req *group.MsgUpdateGroupAccountCommentRequest) (*group.MsgUpdateGroupAccountCommentResponse, error) {
	// TODO
	return &group.MsgUpdateGroupAccountCommentResponse{}, nil
}

func (s serverImpl) CreateProposal(goCtx context.Context, req *group.MsgCreateProposalRequest) (*group.MsgCreateProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposalID, err := s.Keeper.CreateProposal(ctx, req.GroupAccount, req.Comment, req.Proposers, req.GetMsgs())
	if err != nil {
		return nil, err
	}

	// TODO: add event?

	return &group.MsgCreateProposalResponse{ProposalId: proposalID}, nil
}

func (s serverImpl) Vote(goCtx context.Context, req *group.MsgVoteRequest) (*group.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := s.Keeper.Vote(ctx, req.ProposalId, req.Voters, req.Choice, req.Comment); err != nil {
		return nil, err
	}

	// TODO: add event?

	return &group.MsgVoteResponse{}, nil
}

func (s serverImpl) Exec(goCtx context.Context, req *group.MsgExecRequest) (*group.MsgExecResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := s.Keeper.ExecProposal(ctx, req.ProposalId); err != nil {
		return nil, err
	}

	// TODO: add event?

	return &group.MsgExecResponse{}, nil
}

type authNGroupReq interface {
	GetGroupID() group.GroupID
	GetAdmin() sdk.AccAddress // equal GetSigners()
}

type actionFn func(m *group.GroupInfo) error

func (s serverImpl) doUpdateGroup(ctx sdk.Context, req authNGroupReq, action actionFn, note string) error {
	err := s.doAuthenticated(ctx, req, action, note)
	if err != nil {
		return err
	}

	groupIDStr := util.Uint64ToBase58Check(req.GetGroupID().Uint64())
	err = ctx.EventManager().EmitTypedEvent(&group.EventUpdateGroup{GroupId: groupIDStr, Admin: req.GetAdmin().String()})
	if err != nil {
		return err
	}

	return nil
}

func (s serverImpl) doAuthenticated(ctx sdk.Context, req authNGroupReq, action actionFn, note string) error {
	group, err := s.Keeper.GetGroup(ctx, req.GetGroupID())
	if err != nil {
		return err
	}
	if !group.Admin.Equals(req.GetAdmin()) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not group admin")
	}
	if err := action(&group); err != nil {
		return sdkerrors.Wrap(err, note)
	}
	return nil
}
