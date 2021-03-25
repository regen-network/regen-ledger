package testsuite

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	proto "github.com/gogo/protobuf/types"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s *IntegrationTestSuite) TestInitExportGenesis() {
	ctx := s.genesisCtx
	cdc := s.fixture.Codec()

	now := time.Now()
	submittedAt, err := proto.TimestampProto(now)
	s.Require().NoError(err)
	timeout, err := proto.TimestampProto(now.Add(time.Second * 1))
	s.Require().NoError(err)

	groupAccount := &group.GroupAccountInfo{
		Address:  s.groupAccountAddr.String(),
		GroupId:  1,
		Admin:    s.addr1.String(),
		Version:  1,
		Metadata: []byte("account metadata"),
	}
	err = groupAccount.SetDecisionPolicy(&group.ThresholdDecisionPolicy{
		Threshold: "1",
		Timeout:   proto.Duration{Seconds: 1},
	})
	s.Require().NoError(err)

	proposal := &group.Proposal{
		ProposalId:          1,
		GroupAccount:        s.groupAccountAddr.String(),
		Metadata:            []byte("proposal metadata"),
		GroupVersion:        1,
		GroupAccountVersion: 1,
		Proposers: []string{
			s.addr1.String(),
		},
		SubmittedAt: *submittedAt,
		Status:      group.ProposalStatusClosed,
		Result:      group.ProposalResultAccepted,
		VoteState: group.Tally{
			YesCount:     "1",
			NoCount:      "0",
			AbstainCount: "0",
			VetoCount:    "0",
		},
		Timeout:        *timeout,
		ExecutorResult: group.ProposalExecutorResultSuccess,
	}
	err = proposal.SetMsgs([]sdk.Msg{&banktypes.MsgSend{
		FromAddress: s.groupAccountAddr.String(),
		ToAddress:   s.addr2.String(),
		Amount:      sdk.Coins{sdk.NewInt64Coin("test", 100)},
	}})
	s.Require().NoError(err)

	genesisState := &group.GenesisState{
		GroupSeq:        2,
		Groups:          []*group.GroupInfo{{GroupId: 1, Admin: s.addr1.String(), Metadata: []byte("1"), Version: 1, TotalWeight: "1"}, {GroupId: 2, Admin: s.addr2.String(), Metadata: []byte("2"), Version: 2, TotalWeight: "2"}},
		GroupMembers:    []*group.GroupMember{{GroupId: 1, Member: &group.Member{Address: s.addr1.String(), Weight: "1", Metadata: []byte("member metadata")}}, {GroupId: 2, Member: &group.Member{Address: s.addr1.String(), Weight: "2", Metadata: []byte("member metadata")}}},
		GroupAccountSeq: 1,
		GroupAccounts:   []*group.GroupAccountInfo{groupAccount},
		ProposalSeq:     1,
		Proposals:       []*group.Proposal{proposal},
		Votes:           []*group.Vote{{ProposalId: proposal.ProposalId, Voter: s.addr1.String(), SubmittedAt: *submittedAt, Choice: group.Choice_CHOICE_YES}},
	}

	genesisBytes, err := cdc.MarshalJSON(genesisState)
	s.Require().NoError(err)

	_, err = s.initGenesisHandler(ctx, cdc, genesisBytes)
	s.Require().NoError(err)

	for i, g := range genesisState.Groups {
		res, err := s.queryClient.GroupInfo(ctx, &group.QueryGroupInfoRequest{
			GroupId: g.GroupId,
		})
		s.Require().NoError(err)
		s.Require().Equal(g, res.Info)

		membersRes, err := s.queryClient.GroupMembers(ctx, &group.QueryGroupMembersRequest{
			GroupId: g.GroupId,
		})
		s.Require().NoError(err)
		s.Require().Equal(len(membersRes.Members), 1)
		s.Require().Equal(membersRes.Members[0], genesisState.GroupMembers[i])
	}

	for _, g := range genesisState.GroupAccounts {
		res, err := s.queryClient.GroupAccountInfo(ctx, &group.QueryGroupAccountInfoRequest{
			Address: g.Address,
		})
		s.Require().NoError(err)
		s.assertGroupAccountsEqual(g, res.Info)
	}

	for _, g := range genesisState.Proposals {
		res, err := s.queryClient.Proposal(ctx, &group.QueryProposalRequest{
			ProposalId: g.ProposalId,
		})
		s.Require().NoError(err)
		s.assertProposalsEqual(g, res.Proposal)

		votesRes, err := s.queryClient.VotesByProposal(ctx, &group.QueryVotesByProposalRequest{
			ProposalId: g.ProposalId,
		})
		s.Require().NoError(err)
		s.Require().Equal(len(votesRes.Votes), 1)
		s.Require().Equal(votesRes.Votes[0], genesisState.Votes[0])
	}

	exported, err := s.exportGenesisHandler(ctx, cdc)
	s.Require().NoError(err)

	var exportedGenesisState group.GenesisState
	err = cdc.UnmarshalJSON(exported, &exportedGenesisState)
	s.Require().NoError(err)

	s.Require().Equal(genesisState.Groups, exportedGenesisState.Groups)
	s.Require().Equal(genesisState.GroupMembers, exportedGenesisState.GroupMembers)

	s.Require().Equal(len(genesisState.GroupAccounts), len(exportedGenesisState.GroupAccounts))
	for i, g := range genesisState.GroupAccounts {
		res := exportedGenesisState.GroupAccounts[i]
		s.Require().NoError(err)
		s.assertGroupAccountsEqual(g, res)
	}

	s.Require().Equal(len(genesisState.Proposals), len(exportedGenesisState.Proposals))
	for i, g := range genesisState.Proposals {
		res := exportedGenesisState.Proposals[i]
		s.Require().NoError(err)
		s.assertProposalsEqual(g, res)
	}
	s.Require().Equal(genesisState.Votes, exportedGenesisState.Votes)

	s.Require().Equal(genesisState.GroupSeq, exportedGenesisState.GroupSeq)
	s.Require().Equal(genesisState.GroupAccountSeq, exportedGenesisState.GroupAccountSeq)
	s.Require().Equal(genesisState.ProposalSeq, exportedGenesisState.ProposalSeq)

}

func (s *IntegrationTestSuite) assertGroupAccountsEqual(g *group.GroupAccountInfo, other *group.GroupAccountInfo) {
	s.Require().Equal(g.Address, other.Address)
	s.Require().Equal(g.GroupId, other.GroupId)
	s.Require().Equal(g.Admin, other.Admin)
	s.Require().Equal(g.Metadata, other.Metadata)
	s.Require().Equal(g.Version, other.Version)
	s.Require().Equal(g.GetDecisionPolicy(), other.GetDecisionPolicy())
}

func (s *IntegrationTestSuite) assertProposalsEqual(g *group.Proposal, other *group.Proposal) {
	s.Require().Equal(g.ProposalId, other.ProposalId)
	s.Require().Equal(g.GroupAccount, other.GroupAccount)
	s.Require().Equal(g.Metadata, other.Metadata)
	s.Require().Equal(g.Proposers, other.Proposers)
	s.Require().Equal(g.SubmittedAt, other.SubmittedAt)
	s.Require().Equal(g.GroupVersion, other.GroupVersion)
	s.Require().Equal(g.GroupAccountVersion, other.GroupAccountVersion)
	s.Require().Equal(g.Status, other.Status)
	s.Require().Equal(g.Result, other.Result)
	s.Require().Equal(g.VoteState, other.VoteState)
	s.Require().Equal(g.Timeout, other.Timeout)
	s.Require().Equal(g.ExecutorResult, other.ExecutorResult)
	s.Require().Equal(g.GetMsgs(), other.GetMsgs())
}
