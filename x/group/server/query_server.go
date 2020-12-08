package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) GroupInfo(ctx context.Context, request *group.QueryGroupInfoRequest) (*group.QueryGroupInfoResponse, error) {
	groupInfo, err := s.getGroupInfo(sdk.UnwrapSDKContext(ctx), request.GroupId)
	if err != nil {
		return nil, err
	}

	return &group.QueryGroupInfoResponse{Info: &groupInfo}, nil
}

func (s serverImpl) getGroupInfo(ctx sdk.Context, id group.GroupID) (group.GroupInfo, error) {
	var obj group.GroupInfo
	return obj, s.groupTable.GetOne(ctx, id.Bytes(), &obj)
}

func (s serverImpl) GroupAccountInfo(ctx context.Context, request *group.QueryGroupAccountInfoRequest) (*group.QueryGroupAccountInfoResponse, error) {
	groupAccountInfo, err := s.getGroupAccountInfo(sdk.UnwrapSDKContext(ctx), request.GroupAccount)
	if err != nil {
		return nil, err
	}

	return &group.QueryGroupAccountInfoResponse{Info: &groupAccountInfo}, nil
}

func (s serverImpl) getGroupAccountInfo(ctx sdk.Context, accountAddress sdk.AccAddress) (group.GroupAccountInfo, error) {
	var obj group.GroupAccountInfo
	return obj, s.groupAccountTable.GetOne(ctx, accountAddress.Bytes(), &obj)
}

func (s serverImpl) GroupMembers(ctx context.Context, request *group.QueryGroupMembersRequest) (*group.QueryGroupMembersResponse, error) {
	it, err := s.getGroupMembers(sdk.UnwrapSDKContext(ctx), request.GroupId)
	if err != nil {
		return nil, err
	}

	var members []*group.GroupMember
	pageRes, err := orm.Paginate(it, request.Pagination, &members)
	if err != nil {
		return nil, err
	}

	return &group.QueryGroupMembersResponse{
		Members:    members,
		Pagination: pageRes,
	}, nil
}

func (s serverImpl) getGroupMembers(ctx sdk.Context, id group.GroupID) (orm.Iterator, error) {
	return s.groupMemberByGroupIndex.Get(ctx, id.Uint64())
}

// TODO
func (s serverImpl) GroupsByAdmin(ctx context.Context, request *group.QueryGroupsByAdminRequest) (*group.QueryGroupsByAdminResponse, error) {
	return &group.QueryGroupsByAdminResponse{}, nil
}

// TODO
func (s serverImpl) GroupAccountsByGroup(ctx context.Context, request *group.QueryGroupAccountsByGroupRequest) (*group.QueryGroupAccountsByGroupResponse, error) {
	return &group.QueryGroupAccountsByGroupResponse{}, nil
}

// TODO
func (s serverImpl) GroupAccountsByAdmin(ctx context.Context, request *group.QueryGroupAccountsByAdminRequest) (*group.QueryGroupAccountsByAdminResponse, error) {
	return &group.QueryGroupAccountsByAdminResponse{}, nil
}

func (s serverImpl) Proposal(ctx context.Context, request *group.QueryProposalRequest) (*group.QueryProposalResponse, error) {
	proposal, err := s.getProposal(sdk.UnwrapSDKContext(ctx), request.ProposalId)
	if err != nil {
		return nil, err
	}

	return &group.QueryProposalResponse{Proposal: &proposal}, nil
}

// TODO
func (s serverImpl) ProposalsByGroupAccount(ctx context.Context, request *group.QueryProposalsByGroupAccountRequest) (*group.QueryProposalsByGroupAccountResponse, error) {
	return &group.QueryProposalsByGroupAccountResponse{}, nil
}

func (s serverImpl) getProposal(ctx sdk.Context, id group.ProposalID) (group.Proposal, error) {
	var p group.Proposal
	if _, err := s.proposalTable.GetOne(ctx, id.Uint64(), &p); err != nil {
		return group.Proposal{}, sdkerrors.Wrap(err, "load proposal")
	}
	return p, nil
}

// TODO
func (s serverImpl) Votes(ctx context.Context, request *group.QueryVotesRequest) (*group.QueryVotesResponse, error) {
	return &group.QueryVotesResponse{}, nil
}

func (s serverImpl) getVote(ctx sdk.Context, id group.ProposalID, voter sdk.AccAddress) (group.Vote, error) {
	var v group.Vote
	return v, s.voteTable.GetOne(ctx, group.Vote{ProposalId: id, Voter: voter}.NaturalKey(), &v)
}
