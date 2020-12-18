package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
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

func (s serverImpl) getGroupInfo(ctx sdk.Context, id group.ID) (group.GroupInfo, error) {
	var obj group.GroupInfo
	return obj, s.groupTable.GetOne(ctx, id.Bytes(), &obj)
}

func (s serverImpl) GroupAccountInfo(ctx context.Context, request *group.QueryGroupAccountInfoRequest) (*group.QueryGroupAccountInfoResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.GroupAccount)
	if err != nil {
		return nil, err
	}
	groupAccountInfo, err := s.getGroupAccountInfo(sdk.UnwrapSDKContext(ctx), addr)
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
	it, err := s.getGroupMembers(sdk.UnwrapSDKContext(ctx), request.GroupId, request.Pagination)
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

func (s serverImpl) getGroupMembers(ctx sdk.Context, id group.ID, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.groupMemberByGroupIndex.Get(ctx, id.Uint64(), pageRequest)
}

func (s serverImpl) GroupsByAdmin(ctx context.Context, request *group.QueryGroupsByAdminRequest) (*group.QueryGroupsByAdminResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Admin)
	if err != nil {
		return nil, err
	}
	it, err := s.getGroupsByAdmin(sdk.UnwrapSDKContext(ctx), addr, request.Pagination)
	if err != nil {
		return nil, err
	}

	var groups []*group.GroupInfo
	pageRes, err := orm.Paginate(it, request.Pagination, &groups)
	if err != nil {
		return nil, err
	}

	return &group.QueryGroupsByAdminResponse{
		Groups:     groups,
		Pagination: pageRes,
	}, nil
}

func (s serverImpl) getGroupsByAdmin(ctx sdk.Context, admin sdk.AccAddress, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.groupByAdminIndex.Get(ctx, admin.Bytes(), pageRequest)
}

func (s serverImpl) GroupAccountsByGroup(ctx context.Context, request *group.QueryGroupAccountsByGroupRequest) (*group.QueryGroupAccountsByGroupResponse, error) {
	it, err := s.getGroupAccountsByGroup(sdk.UnwrapSDKContext(ctx), request.GroupId, request.Pagination)
	if err != nil {
		return nil, err
	}

	var accounts []*group.GroupAccountInfo
	pageRes, err := orm.Paginate(it, request.Pagination, &accounts)
	if err != nil {
		return nil, err
	}

	return &group.QueryGroupAccountsByGroupResponse{
		GroupAccounts: accounts,
		Pagination:    pageRes,
	}, nil
}

func (s serverImpl) getGroupAccountsByGroup(ctx sdk.Context, id group.ID, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.groupAccountByGroupIndex.Get(ctx, id.Uint64(), pageRequest)
}

func (s serverImpl) GroupAccountsByAdmin(ctx context.Context, request *group.QueryGroupAccountsByAdminRequest) (*group.QueryGroupAccountsByAdminResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Admin)
	if err != nil {
		return nil, err
	}
	it, err := s.getGroupAccountsByAdmin(sdk.UnwrapSDKContext(ctx), addr, request.Pagination)
	if err != nil {
		return nil, err
	}

	var accounts []*group.GroupAccountInfo
	pageRes, err := orm.Paginate(it, request.Pagination, &accounts)
	if err != nil {
		return nil, err
	}

	return &group.QueryGroupAccountsByAdminResponse{
		GroupAccounts: accounts,
		Pagination:    pageRes,
	}, nil
}

func (s serverImpl) getGroupAccountsByAdmin(ctx sdk.Context, admin sdk.AccAddress, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.groupAccountByAdminIndex.Get(ctx, admin.Bytes(), pageRequest)
}

func (s serverImpl) Proposal(ctx context.Context, request *group.QueryProposalRequest) (*group.QueryProposalResponse, error) {
	proposal, err := s.getProposal(sdk.UnwrapSDKContext(ctx), request.ProposalId)
	if err != nil {
		return nil, err
	}

	return &group.QueryProposalResponse{Proposal: &proposal}, nil
}

func (s serverImpl) ProposalsByGroupAccount(ctx context.Context, request *group.QueryProposalsByGroupAccountRequest) (*group.QueryProposalsByGroupAccountResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.GroupAccount)
	if err != nil {
		return nil, err
	}
	it, err := s.getProposalsByGroupAccount(sdk.UnwrapSDKContext(ctx), addr, request.Pagination)
	if err != nil {
		return nil, err
	}

	var proposals []*group.Proposal
	pageRes, err := orm.Paginate(it, request.Pagination, &proposals)
	if err != nil {
		return nil, err
	}

	return &group.QueryProposalsByGroupAccountResponse{
		Proposals:  proposals,
		Pagination: pageRes,
	}, nil
}

func (s serverImpl) getProposalsByGroupAccount(ctx sdk.Context, account sdk.AccAddress, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.proposalByGroupAccountIndex.Get(ctx, account.Bytes(), pageRequest)
}

func (s serverImpl) getProposal(ctx sdk.Context, id group.ProposalID) (group.Proposal, error) {
	var p group.Proposal
	if _, err := s.proposalTable.GetOne(ctx, id.Uint64(), &p); err != nil {
		return group.Proposal{}, sdkerrors.Wrap(err, "load proposal")
	}
	return p, nil
}

func (s serverImpl) VoteByProposalVoter(ctx context.Context, request *group.QueryVoteByProposalVoterRequest) (*group.QueryVoteByProposalVoterResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Voter)
	if err != nil {
		return nil, err
	}
	vote, err := s.getVote(sdk.UnwrapSDKContext(ctx), request.ProposalId, addr)
	if err != nil {
		return nil, err
	}
	return &group.QueryVoteByProposalVoterResponse{
		Vote: &vote,
	}, nil
}

func (s serverImpl) VotesByProposal(ctx context.Context, request *group.QueryVotesByProposalRequest) (*group.QueryVotesByProposalResponse, error) {
	it, err := s.getVotesByProposal(sdk.UnwrapSDKContext(ctx), request.ProposalId, request.Pagination)
	if err != nil {
		return nil, err
	}

	var votes []*group.Vote
	pageRes, err := orm.Paginate(it, request.Pagination, &votes)
	if err != nil {
		return nil, err
	}

	return &group.QueryVotesByProposalResponse{
		Votes:      votes,
		Pagination: pageRes,
	}, nil
}

func (s serverImpl) VotesByVoter(ctx context.Context, request *group.QueryVotesByVoterRequest) (*group.QueryVotesByVoterResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Voter)
	if err != nil {
		return nil, err
	}
	it, err := s.getVotesByVoter(sdk.UnwrapSDKContext(ctx), addr, request.Pagination)
	if err != nil {
		return nil, err
	}

	var votes []*group.Vote
	pageRes, err := orm.Paginate(it, request.Pagination, &votes)
	if err != nil {
		return nil, err
	}

	return &group.QueryVotesByVoterResponse{
		Votes:      votes,
		Pagination: pageRes,
	}, nil
}

func (s serverImpl) getVote(ctx sdk.Context, id group.ProposalID, voter sdk.AccAddress) (group.Vote, error) {
	var v group.Vote
	return v, s.voteTable.GetOne(ctx, group.Vote{ProposalId: id, Voter: voter.String()}.NaturalKey(), &v)
}

func (s serverImpl) getVotesByProposal(ctx sdk.Context, id group.ProposalID, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.voteByProposalIndex.Get(ctx, id.Uint64(), pageRequest)
}

func (s serverImpl) getVotesByVoter(ctx sdk.Context, voter sdk.AccAddress, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.voteByVoterIndex.Get(ctx, voter.Bytes(), pageRequest)
}
