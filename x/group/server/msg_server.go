package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/util"
	"github.com/regen-network/regen-ledger/x/group/types"
)

func (s serverImpl) CreateGroup(goCtx context.Context, req *types.MsgCreateGroupRequest) (*types.MsgCreateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	groupID, err := s.Keeper.CreateGroup(ctx, req.Admin, req.Members, req.Comment)
	if err != nil {
		return nil, err
	}

	groupIDStr := util.Uint64ToBase58Check(groupID.Uint64())
	err = ctx.EventManager().EmitTypedEvent(&types.EventCreateGroup{Group: groupIDStr, Admin: req.Admin.String()})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateGroupResponse{Group: groupID}, nil
}

func (s serverImpl) UpdateGroupMembers(goCtx context.Context, req *types.MsgUpdateGroupMembersRequest) (*types.MsgUpdateGroupMembersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	action := func(m *types.GroupMetadata) error {
		for i := range req.MemberUpdates {
			member := types.GroupMember{Group: req.Group,
				Member:  req.MemberUpdates[i].Address,
				Weight:  req.MemberUpdates[i].Power,
				Comment: req.MemberUpdates[i].Comment,
			}
			var found bool
			var previousMemberStatus types.GroupMember
			switch err := s.groupMemberTable.GetOne(ctx, member.NaturalKey(), &previousMemberStatus); {
			case err == nil:
				found = true
			case orm.ErrNotFound.Is(err):
				found = false
			default:
				return sdkerrors.Wrap(err, "get group member")
			}

			// handle delete
			if member.Weight.Equal(sdk.ZeroDec()) {
				if !found {
					return sdkerrors.Wrap(orm.ErrNotFound, "unknown member")
				}
				m.TotalWeight = m.TotalWeight.Sub(previousMemberStatus.Weight)
				if err := s.groupMemberTable.Delete(ctx, &member); err != nil {
					return sdkerrors.Wrap(err, "delete member")
				}
				continue
			}
			// handle add + update
			if found {
				m.TotalWeight = m.TotalWeight.Sub(previousMemberStatus.Weight)
				if err := s.groupMemberTable.Save(ctx, &member); err != nil {
					return sdkerrors.Wrap(err, "add member")
				}
			} else if err := s.groupMemberTable.Create(ctx, &member); err != nil {
				return sdkerrors.Wrap(err, "add member")
			}
			m.TotalWeight = m.TotalWeight.Add(member.Weight)
		}
		return s.UpdateGroup(ctx, m)
	}

	err := s.doUpdateGroup(ctx, req, action, "members updated")
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateGroupMembersResponse{}, nil
}

func (s serverImpl) UpdateGroupAdmin(goCtx context.Context, req *types.MsgUpdateGroupAdminRequest) (*types.MsgUpdateGroupAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	action := func(m *types.GroupMetadata) error {
		m.Admin = req.NewAdmin
		return s.UpdateGroup(ctx, m)
	}

	err := s.doUpdateGroup(ctx, req, action, "admin updated")
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateGroupAdminResponse{}, nil
}

func (s serverImpl) UpdateGroupComment(goCtx context.Context, req *types.MsgUpdateGroupCommentRequest) (*types.MsgUpdateGroupCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	action := func(m *types.GroupMetadata) error {
		m.Comment = req.Comment
		return s.UpdateGroup(ctx, m)
	}

	err := s.doUpdateGroup(ctx, req, action, "comment updated")
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateGroupCommentResponse{}, nil
}

func (s serverImpl) CreateGroupAccount(goCtx context.Context, req *types.MsgCreateGroupAccountRequest) (*types.MsgCreateGroupAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	decisionPolicy := req.GetDecisionPolicy()
	acc, err := s.Keeper.CreateGroupAccount(ctx, req.GetAdmin(), req.GetGroup(), decisionPolicy, req.GetComment())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "create group account")
	}

	err = ctx.EventManager().EmitTypedEvent(&types.EventCreateGroupAccount{GroupAccount: acc.String(), Admin: req.Admin.String()})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateGroupAccountResponse{GroupAccount: acc}, nil
}

func (s serverImpl) UpdateGroupAccountAdmin(goCtx context.Context, req *types.MsgUpdateGroupAccountAdminRequest) (*types.MsgUpdateGroupAccountAdminResponse, error) {
	// TODO
	return &types.MsgUpdateGroupAccountAdminResponse{}, nil
}

func (s serverImpl) UpdateGroupAccountDecisionPolicy(goCtx context.Context, req *types.MsgUpdateGroupAccountDecisionPolicyRequest) (*types.MsgUpdateGroupAccountDecisionPolicyResponse, error) {
	// TODO
	return &types.MsgUpdateGroupAccountDecisionPolicyResponse{}, nil
}

func (s serverImpl) UpdateGroupAccountComment(goCtx context.Context, req *types.MsgUpdateGroupAccountCommentRequest) (*types.MsgUpdateGroupAccountCommentResponse, error) {
	// TODO
	return &types.MsgUpdateGroupAccountCommentResponse{}, nil
}

func (s serverImpl) CreateProposal(goCtx context.Context, req *types.MsgCreateProposalRequest) (*types.MsgCreateProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposalID, err := s.Keeper.CreateProposal(ctx, req.GroupAccount, req.Comment, req.Proposers, req.GetMsgs())
	if err != nil {
		return nil, err
	}

	// TODO: add event?

	return &types.MsgCreateProposalResponse{Proposal: proposalID}, nil
}

func (s serverImpl) Vote(goCtx context.Context, req *types.MsgVoteRequest) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := s.Keeper.Vote(ctx, req.Proposal, req.Voters, req.Choice, req.Comment); err != nil {
		return nil, err
	}

	// TODO: add event?

	return &types.MsgVoteResponse{}, nil
}

func (s serverImpl) Exec(goCtx context.Context, req *types.MsgExecRequest) (*types.MsgExecResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := s.Keeper.ExecProposal(ctx, req.Proposal); err != nil {
		return nil, err
	}

	// TODO: add event?

	return &types.MsgExecResponse{}, nil
}

type authNGroupReq interface {
	GetGroup() types.GroupID
	GetAdmin() sdk.AccAddress // equal GetSigners()
}

type actionFn func(m *types.GroupMetadata) error

func (s serverImpl) doUpdateGroup(ctx sdk.Context, req authNGroupReq, action actionFn, note string) error {
	err := s.doAuthenticated(ctx, req, action, note)
	if err != nil {
		return err
	}

	groupIDStr := util.Uint64ToBase58Check(req.GetGroup().Uint64())
	err = ctx.EventManager().EmitTypedEvent(&types.EventCreateGroup{Group: groupIDStr, Admin: req.GetAdmin().String()})
	if err != nil {
		return err
	}

	return nil
}

func (s serverImpl) doAuthenticated(ctx sdk.Context, req authNGroupReq, action actionFn, note string) error {
	group, err := s.Keeper.GetGroup(ctx, req.GetGroup())
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
