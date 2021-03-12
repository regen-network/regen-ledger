package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) GroupInfo(ctx types.Context, request *group.QueryGroupInfoRequest) (*group.QueryGroupInfoResponse, error) {
	groupID := request.GroupId
	groupInfo, err := s.getGroupInfo(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return &group.QueryGroupInfoResponse{Info: &groupInfo}, nil
}

func (s serverImpl) getGroupInfo(ctx types.Context, id uint64) (group.GroupInfo, error) {
	var obj group.GroupInfo
	return obj, s.groupTable.GetOne(ctx, group.ID(id).Bytes(), &obj)
}

func (s serverImpl) GroupAccountInfo(ctx types.Context, request *group.QueryGroupAccountInfoRequest) (*group.QueryGroupAccountInfoResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.GroupAccount)
	if err != nil {
		return nil, err
	}
	groupAccountInfo, err := s.getGroupAccountInfo(ctx, addr)
	if err != nil {
		return nil, err
	}

	return &group.QueryGroupAccountInfoResponse{Info: &groupAccountInfo}, nil
}

func (s serverImpl) getGroupAccountInfo(ctx types.Context, accountAddress sdk.AccAddress) (group.GroupAccountInfo, error) {
	var obj group.GroupAccountInfo
	return obj, s.groupAccountTable.GetOne(ctx, accountAddress.Bytes(), &obj)
}

func (s serverImpl) GroupMembers(ctx types.Context, request *group.QueryGroupMembersRequest) (*group.QueryGroupMembersResponse, error) {
	groupID := request.GroupId
	it, err := s.getGroupMembers(ctx, groupID, request.Pagination)
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

func (s serverImpl) getGroupMembers(ctx types.Context, id uint64, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.groupMemberByGroupIndex.GetPaginated(ctx, id, pageRequest)
}

func (s serverImpl) GroupsByAdmin(ctx types.Context, request *group.QueryGroupsByAdminRequest) (*group.QueryGroupsByAdminResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Admin)
	if err != nil {
		return nil, err
	}
	it, err := s.getGroupsByAdmin(ctx, addr, request.Pagination)
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

func (s serverImpl) getGroupsByAdmin(ctx types.Context, admin sdk.AccAddress, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.groupByAdminIndex.GetPaginated(ctx, admin.Bytes(), pageRequest)
}

func (s serverImpl) GroupAccountsByGroup(ctx types.Context, request *group.QueryGroupAccountsByGroupRequest) (*group.QueryGroupAccountsByGroupResponse, error) {
	groupID := request.GroupId
	it, err := s.getGroupAccountsByGroup(ctx, groupID, request.Pagination)
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

func (s serverImpl) getGroupAccountsByGroup(ctx types.Context, id uint64, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.groupAccountByGroupIndex.GetPaginated(ctx, id, pageRequest)
}

func (s serverImpl) GroupAccountsByAdmin(ctx types.Context, request *group.QueryGroupAccountsByAdminRequest) (*group.QueryGroupAccountsByAdminResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Admin)
	if err != nil {
		return nil, err
	}
	it, err := s.getGroupAccountsByAdmin(ctx, addr, request.Pagination)
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

func (s serverImpl) getGroupAccountsByAdmin(ctx types.Context, admin sdk.AccAddress, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.groupAccountByAdminIndex.GetPaginated(ctx, admin.Bytes(), pageRequest)
}

func (s serverImpl) Proposal(ctx types.Context, request *group.QueryProposalRequest) (*group.QueryProposalResponse, error) {
	proposalID := request.ProposalId
	proposal, err := s.getProposal(ctx, proposalID)
	if err != nil {
		return nil, err
	}

	return &group.QueryProposalResponse{Proposal: &proposal}, nil
}

func (s serverImpl) ProposalsByGroupAccount(ctx types.Context, request *group.QueryProposalsByGroupAccountRequest) (*group.QueryProposalsByGroupAccountResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.GroupAccount)
	if err != nil {
		return nil, err
	}
	it, err := s.getProposalsByGroupAccount(ctx, addr, request.Pagination)
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

func (s serverImpl) getProposalsByGroupAccount(ctx types.Context, account sdk.AccAddress, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.proposalByGroupAccountIndex.GetPaginated(ctx, account.Bytes(), pageRequest)
}

func (s serverImpl) getProposal(ctx types.Context, proposalID uint64) (group.Proposal, error) {
	var p group.Proposal
	if _, err := s.proposalTable.GetOne(ctx, proposalID, &p); err != nil {
		return group.Proposal{}, sdkerrors.Wrap(err, "load proposal")
	}
	return p, nil
}

func (s serverImpl) VoteByProposalVoter(ctx types.Context, request *group.QueryVoteByProposalVoterRequest) (*group.QueryVoteByProposalVoterResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Voter)
	if err != nil {
		return nil, err
	}
	proposalID := request.ProposalId
	vote, err := s.getVote(ctx, proposalID, addr)
	if err != nil {
		return nil, err
	}
	return &group.QueryVoteByProposalVoterResponse{
		Vote: &vote,
	}, nil
}

func (s serverImpl) VotesByProposal(ctx types.Context, request *group.QueryVotesByProposalRequest) (*group.QueryVotesByProposalResponse, error) {
	proposalID := request.ProposalId
	it, err := s.getVotesByProposal(ctx, proposalID, request.Pagination)
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

func (s serverImpl) VotesByVoter(ctx types.Context, request *group.QueryVotesByVoterRequest) (*group.QueryVotesByVoterResponse, error) {
	addr, err := sdk.AccAddressFromBech32(request.Voter)
	if err != nil {
		return nil, err
	}
	it, err := s.getVotesByVoter(ctx, addr, request.Pagination)
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

func (s serverImpl) getVote(ctx types.Context, proposalID uint64, voter sdk.AccAddress) (group.Vote, error) {
	var v group.Vote
	return v, s.voteTable.GetOne(ctx, group.Vote{ProposalId: proposalID, Voter: voter.String()}.NaturalKey(), &v)
}

func (s serverImpl) getVotesByProposal(ctx types.Context, proposalID uint64, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.voteByProposalIndex.GetPaginated(ctx, proposalID, pageRequest)
}

func (s serverImpl) getVotesByVoter(ctx types.Context, voter sdk.AccAddress, pageRequest *query.PageRequest) (orm.Iterator, error) {
	return s.voteByVoterIndex.GetPaginated(ctx, voter.Bytes(), pageRequest)
}
